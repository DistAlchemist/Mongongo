// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import (
	"log"
	"sort"
	"sync"
	"sync/atomic"
	"time"

	"github.com/davecgh/go-spew/spew"

	"github.com/DistAlchemist/Mongongo/config"
)

// Memtable specifies memtable
type Memtable struct {
	flushKey         string
	threshold        int32
	thresholdCnt     int32
	currentSize      int32
	currentObjectCnt int32
	// table and cf name are used to determine the
	// ColumnFamilyStore
	tableName string
	cfName    string
	// creation time of this memtable
	creationTime   int64
	isFrozen       bool
	isDirty        bool
	isFlushed      bool
	columnFamilies map[string]ColumnFamily
	// lock and condition for notifying new clients about Memtable switches
	mu   sync.Mutex
	cond *sync.Cond
}

// NewMemtable initializes a new memtable
func NewMemtable(table, cfName string) *Memtable {
	m := &Memtable{}
	m.flushKey = "FlushKey"
	m.threshold = int32(config.MemtableSize * 1024 * 1024)
	m.thresholdCnt = int32(config.MemtableObjectCount * 1024 * 1024)
	m.currentSize = 0
	m.currentObjectCnt = 0
	m.isFrozen = false
	m.isDirty = false
	m.isFlushed = false
	m.columnFamilies = make(map[string]ColumnFamily)
	m.cond = sync.NewCond(&m.mu)
	m.tableName = table
	m.cfName = cfName
	m.creationTime = time.Now().UnixNano() / int64(time.Millisecond)
	return m
}

func (m *Memtable) put(key string, columnFamily *ColumnFamily) {
	// should only be called by ColumnFamilyStore.apply
	if m.isFrozen {
		log.Fatal("memtable is frozen!")
	}
	m.isDirty = true
	m.runResolve(key, columnFamily)
}

func (m *Memtable) runResolve(key string, columnFamily *ColumnFamily) {
	oldCf, ok := m.columnFamilies[key]
	if ok {
		oldSize := oldCf.size
		oldObjectCount := oldCf.getColumnCount()
		oldCf.addColumns(columnFamily)
		newSize := oldCf.size
		newObjectCount := oldCf.getColumnCount()
		m.resolveSize(oldSize, newSize)
		m.resolveCount(oldObjectCount, newObjectCount)
		oldCf.deleteCF(columnFamily)
	} else {
		m.columnFamilies[key] = *columnFamily
		atomic.AddInt32(&m.currentSize, columnFamily.size+int32(len(key)))
		atomic.AddInt32(&m.currentObjectCnt, int32(columnFamily.getColumnCount()))
	}
}

func (m *Memtable) resolveSize(oldSize, newSize int32) {
	atomic.AddInt32(&m.currentSize, int32(newSize-oldSize))
}

func (m *Memtable) resolveCount(oldCount, newCount int) {
	atomic.AddInt32(&m.currentObjectCnt, int32(newCount-oldCount))
}

func (m *Memtable) isThresholdViolated() bool {
	if m.currentSize >= m.threshold || m.currentObjectCnt >= m.thresholdCnt {
		return true
	}
	return false
}

func (m *Memtable) flush(cLogCtx *CommitLogContext) {
	// flush this memtable to disk
	cfStore := OpenTable(m.tableName).columnFamilyStores[m.cfName]
	writer := NewSSTableWriter(cfStore.getTmpSSTablePath(), len(m.columnFamilies))
	// sort keys in the order they would be in when decorated
	orderedKeys := make([]string, 0)
	for cfName := range m.columnFamilies {
		orderedKeys = append(orderedKeys, writer.partitioner.DecorateKey(cfName))
	}
	sort.Sort(ByKey(orderedKeys))
	for _, key := range orderedKeys {
		k := writer.partitioner.DecorateKey(key)
		buf := make([]byte, 0)
		columnFamily, ok := m.columnFamilies[k]
		if ok {
			// serialize the cf with column indexes
			CFSerializer.serializeWithIndexes(&columnFamily, buf)
			// now write the key and value to disk
			writer.append(key, buf)
		}
	}
	ssTable := writer.closeAndOpenReader()
	cfStore.onMemtableFlush(cLogCtx)
	cfStore.storeLocation(ssTable)
	m.isFlushed = true
	log.Print("Completed flushing ", ssTable.getFilename())
}

func (m *Memtable) freeze() {
	m.isFrozen = true
}

func reverse(a []IColumn) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func (m *Memtable) getNamesIterator(filter *NamesQueryFilter) ColumnIterator {
	cf, ok := m.columnFamilies[filter.key]
	spew.Printf("\tcf get from memtable: %#+v\n\n", cf)
	var columnFamily *ColumnFamily
	if ok == false {
		columnFamily = createColumnFamily(m.tableName, filter.path.ColumnFamilyName)
	} else {
		// columnFamily = cf.cloneMeShallow()
		columnFamily = &cf
		spew.Printf("\tshould enter here, cf: %#+v\n\n", columnFamily)
	}
	return NewSColumnIterator(0, columnFamily, filter.columns)
}

// obtain an iterator of columns in this memtable in the specified
// order starting from a given column
func (m *Memtable) getSliceIterator(filter *SliceQueryFilter) ColumnIterator {
	cf, ok := m.columnFamilies[filter.key] // rowKey -> column family
	var columnFamily *ColumnFamily
	var columns []IColumn
	if ok == false {
		columnFamily = createColumnFamily(m.tableName, filter.path.ColumnFamilyName)
		columns = columnFamily.GetSortedColumns()
	} else {
		columnFamily = cf.cloneMeShallow()
		columns = cf.GetSortedColumns()
	}
	if filter.reversed == true {
		reverse(columns)
	}
	var startIColumn IColumn
	if config.GetColumnTypeTableName(m.tableName, filter.path.ColumnFamilyName) == "Standard" {
		startIColumn = NewColumn(string(filter.start), "", 0, false)
	} else {
		startIColumn = NewSuperColumn(string(filter.start))
	}
	index := 0
	if len(filter.start) == 0 && filter.reversed {
		// scan from the largest column in descending order
		index = 0
	} else {
		index = sort.Search(len(columns), func(i int) bool {
			return columns[i].getName() >= startIColumn.getName()
		})
	}
	startIndex := index
	return NewAColumnIterator(startIndex, columnFamily, columns)
}

func (m *Memtable) isClean() bool {
	return m.isDirty == false
}
