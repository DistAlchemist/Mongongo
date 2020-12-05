// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import (
	"sync"
	"sync/atomic"
	"time"

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
	m.columnFamilies = make(map[string]ColumnFamily)
	m.cond = sync.NewCond(&m.mu)
	m.tableName = table
	m.cfName = cfName
	m.creationTime = time.Now().UnixNano() / int64(time.Millisecond)
	return m
}

// put data into the memtable
// flush memtable to disk when the size exceeds the threshold
func (m *Memtable) put(key string, columnFamily *ColumnFamily, cLogCtx *CommitLogContext) {
	if m.isThresholdViolated(key) {
		// flush memtable to disk as SSTable if size excedes the limit
		m.mu.Lock()
		defer m.mu.Unlock()
		cfStore := openTable(m.tableName).columnFamilyStores[m.cfName]
		if !m.isFrozen {
			m.isFrozen = true
			// submit memtable flush
			GetMemtableManager().submit(cfStore.columnFamilyName, m, cLogCtx)
			cfStore.switchMemtable(key, columnFamily, cLogCtx)
		} else {
			cfStore.apply(key, columnFamily, cLogCtx)
		}
	} else {
		// submit task to put key-cf to memtable
		go m.runResolve(key, columnFamily)
	}
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

func (m *Memtable) isThresholdViolated(key string) bool {
	bVal := false
	if m.currentSize >= m.threshold || m.currentObjectCnt >= m.thresholdCnt || bVal || key == m.flushKey {
		return true
	}
	return false
}

func (m *Memtable) flush(cLogCtx *CommitLogContext) {
	// flush this memtable to disk
	cfStore := openTable(m.tableName).columnFamilyStores[m.cfName]
	if len(m.columnFamilies) == 0 {
		// This should be called even if size is 0
		// Because we should try to delete the useless commitlogs
		// even though there is nothing to flush in memtables for
		// a given family like Hints etc.
		cfStore.onMemtableFlush(cLogCtx)
		return
	}
	// partitioner type: OrderPreserving or Random
	pType := config.HashingStrategy
	dir := config.DataFileDirs[0]
	filename := cfStore.getNextFileName()
	ssTable := NewSSTableP(dir, filename, pType)
	switch pType {
	case config.Ophf:
		m.flushForOrderPreservingPartitioner(ssTable, cfStore, cLogCtx)
	default:
		m.flushForRandomPartitioner(ssTable, cfStore, cLogCtx)
	}
	m.columnFamilies = make(map[string]ColumnFamily)
}

func (m *Memtable) flushForOrderPreservingPartitioner(ssTable *SSTable, cfStore *ColumnFamilyStore, cLogCtx *CommitLogContext) {
	// TODO
}

func (m *Memtable) flushForRandomPartitioner(ssTable *SSTable, cfStore *ColumnFamilyStore, cLogCtx *CommitLogContext) {
	// TODO
}
