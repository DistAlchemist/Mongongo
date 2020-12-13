// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	tableInstances   = map[string]*Table{}
	tableMetadataMap = map[string]*TableMetadata{}
	idCFMap          = map[int]string{}
	tCreateLock      sync.Mutex
)

// Table ...
type Table struct {
	//
	tableMetadata      *TableMetadata
	tableName          string
	columnFamilyStores map[string]*ColumnFamilyStore
}

func OpenTable(table string) *Table {
	tableInstance, ok := tableInstances[table]
	if !ok {
		// read config to know the column families for
		// this table.
		tCreateLock.Lock()
		defer tCreateLock.Unlock()
		tableInstance = NewTable(table)
		tableInstances[table] = tableInstance
	}
	return tableInstance
}

// NewTable create a Table
func NewTable(tableName string) *Table {
	t := &Table{}
	t.tableName = tableName
	t.tableMetadata = getTableMetadataInstance(t.tableName)
	cfIDMap := t.tableMetadata.cfIDMap
	for columnFamily := range cfIDMap {
		t.columnFamilyStores[columnFamily] = NewColumnFamilyStore(tableName, columnFamily)
	}
	return t
}

func (t *Table) get(key string) *Row {
	// selects the row associated with the given key
	row := NewRowT(t.tableName, key)
	for columnFamily := range t.getColumnFamilies() {
		cf := t.getCF(key, columnFamily)
		if cf != nil {
			row.addColumnFamily(cf)
		}
	}
	return row
}

func (t *Table) getCF(key, cfName string) *ColumnFamily {
	cfStore, ok := t.columnFamilyStores[cfName]
	if ok == false {
		log.Fatal("Column family" + cfName + " has not been defined")
	}
	return cfStore.getColumnFamily(NewIdentityQueryFilter(key, NewQueryPathCF(cfName)))
}

func (t *Table) getColumnFamilies() map[string]int {
	return t.tableMetadata.cfIDMap
}

func (t *Table) getColumnFamilyStore(cfName string) *ColumnFamilyStore {
	return t.columnFamilyStores[cfName]
}

func getTableMetadataInstance(tableName string) *TableMetadata {
	tableMetadata, ok := tableMetadataMap[tableName]
	if !ok {
		tableMetadata = NewTableMetadata()
		tableMetadataMap[tableName] = tableMetadata
	}
	return tableMetadata
}

func (t *Table) loadTableMetadata(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(f)
	sizeStr, err := reader.ReadString(' ')
	if err != nil {
		log.Fatal(err)
	}
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		log.Fatal(err)
	}
	tmetadata := NewTableMetadata()
	for i := 0; i < size; i++ {
		cfName, err := reader.ReadString(' ')
		if err != nil {
			log.Fatal(err)
		}
		idStr, err := reader.ReadString(' ')
		if err != nil {
			log.Fatal(err)
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Fatal(err)
		}
		typeName, err := reader.ReadString(' ')
		if err != nil {
			log.Fatal(err)
		}
		tmetadata.Add(cfName, id, typeName)
	}
	t.tableMetadata = tmetadata
}

func (t *Table) onStart() {
	cfIDMap := t.tableMetadata.cfIDMap
	for columnFamily := range cfIDMap {
		cfStore := t.columnFamilyStores[columnFamily]
		cfStore.onStart()
	}
}

func (t *Table) isValidColumnFamily(columnFamily string) bool {
	return t.tableMetadata.isValidColumnFamily(columnFamily)
}

func (t *Table) getNumberOfColumnFamilies() int {
	return t.tableMetadata.getSize()
}

// First adds the row to the commit log associated with this
// table. Then the data associated with the individual column
// families is also written to the column family store's memtable
func (t *Table) apply(row *Row) {
	key := row.Key
	// add row to commit log
	start := time.Now().UnixNano() / int64(time.Millisecond)
	// cLogCtx := openCommitLog(t.tableName).add(row) // first write to commitlog
	cLogCtx := openCommitLogE().add(row) // first write to commitlog
	for cName, columnFamily := range row.ColumnFamilies {
		cfStore := t.columnFamilyStores[cName]
		cfStore.apply(key, columnFamily, cLogCtx) // then write to memtable
	}
	// row.clear()
	timeTaken := time.Now().UnixNano()/int64(time.Millisecond) - start
	log.Printf("table.apply(row) took %v ms\n", timeTaken)
}

func (t *Table) getColumnFamilyID(cfName string) int {
	return t.tableMetadata.getColumnFamilyID(cfName)
}

func (t *Table) getRow(filter QueryFilter) *Row {
	cfStore := t.columnFamilyStores[filter.getPath().ColumnFamilyName]
	row := NewRowT(t.tableName, filter.getKey())
	columnFamily := cfStore.getColumnFamily(filter)
	if columnFamily != nil {
		row.addColumnFamily(columnFamily)
	}
	return row
}
