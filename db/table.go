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
)

var (
	tableInstances map[string]*Table
	tCreateLock    sync.Mutex
)

// Table ...
type Table struct {
	//
	tableMetadata      *TableMetadata
	tableName          string
	columnFamilyStores map[string]*ColumnFamilyStore
}

func openTable(table string) *Table {
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
	t.tableMetadata = t.getTableMetadataInstance()
	cfIDMap := t.tableMetadata.cfIDMap
	for columnFamily := range cfIDMap {
		t.columnFamilyStores[columnFamily] = NewColumnFamilyStore(tableName, columnFamily)
	}
	return t
}

func (t *Table) getTableMetadataInstance() *TableMetadata {
	if t.tableMetadata == nil {
		fileName := getFileName()
		if _, err := os.Stat(fileName); err == nil {
			// file exists
			t.loadTableMetadata(fileName)
		} else if os.IsNotExist(err) {
			return NewTableMetadata()
		} else {
			log.Fatal(err)
		}
	}
	return t.tableMetadata
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
		if cfStore != nil {
			cfStore.onStart()
		}
	}
}
