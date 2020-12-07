// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import (
	"encoding/binary"
	"sync/atomic"

	"github.com/DistAlchemist/Mongongo/config"
)

// ColumnFamily definition
type ColumnFamily struct {
	ColumnFamilyName string
	ColumnType       string
	Factory          AColumnFactory
	Columns          map[string]IColumn
	size             int32
	deleteMark       bool
}

var typeToColumnFactory = map[string]AColumnFactory{
	"Standard": ColumnFactory{},
	"Super":    SuperColumnFactory{},
}

// NewColumnFamily create a new column family, set columnfactory according to its type
func NewColumnFamily(columnFamilyName, columnType string) *ColumnFamily {
	cf := &ColumnFamily{}
	cf.ColumnFamilyName = columnFamilyName
	cf.ColumnType = columnType
	cf.Factory = typeToColumnFactory[columnType]
	cf.deleteMark = false
	return cf
}

// CreateColumn setup a new column in columnFamily
func (cf *ColumnFamily) CreateColumn(columnName, value string, timestamp int64) {
	column := cf.Factory.createColumn(columnName, value, timestamp)
	cf.addColumn(columnName, column)
}

func (cf *ColumnFamily) addColumn(name string, column IColumn) {
	newSize := int32(0)
	oldColumn, ok := cf.Columns[name]
	if !ok {
		oldSize := oldColumn.getSize()
		if oldColumn.putColumn(column) {
			// This will never be called for super column as put column
			// always returns false
			cf.Columns[name] = column
			newSize = column.getSize()
		} else {
			newSize = oldColumn.getSize()
		}
		atomic.AddInt32(&cf.size, int32(newSize-oldSize))
	} else {
		newSize = column.getSize()
		atomic.AddInt32(&cf.size, newSize)
		cf.Columns[name] = column
	}
}

func (cf *ColumnFamily) getSize() int32 {
	if cf.size == 0 {
		for cfName := range cf.Columns {
			atomic.AddInt32(&cf.size, cf.Columns[cfName].getSize())
		}
	}
	return cf.size
}

func (cf *ColumnFamily) isMarkedForDelete() bool {
	return cf.deleteMark
}

func (cf *ColumnFamily) toByteArray() []byte {
	buf := make([]byte, 0)
	// write cf name length
	b4 := make([]byte, 4)
	binary.BigEndian.PutUint32(b4, uint32(len(cf.ColumnFamilyName)))
	buf = append(buf, b4...)
	// write cf name bytes
	buf = append(buf, []byte(cf.ColumnFamilyName)...)
	// write if this cf is marked for delete
	if cf.deleteMark {
		buf = append(buf, byte(1))
	} else {
		buf = append(buf, byte(0))
	}
	// write column size
	binary.BigEndian.PutUint32(b4, uint32(len(cf.Columns)))
	buf = append(buf, b4...)
	// write column bytes
	for _, column := range cf.Columns {
		buf = append(buf, column.toByteArray()...)
	}
	return buf
}

func (cf *ColumnFamily) getColumnCount() int {
	count := 0
	columns := cf.Columns
	if columns != nil {
		if config.GetColumnType(cf.ColumnFamilyName) != "Super" {
			count = len(columns)
		} else {
			for _, column := range columns {
				count += column.getObjectCount()
			}
		}
	}
	return count
}

func (cf *ColumnFamily) addColumns(columnFamily *ColumnFamily) {
	columns := cf.Columns
	for cName, column := range columns {
		cf.addColumn(cName, column)
	}
}

func (cf *ColumnFamily) clear() {
	cf.Columns = make(map[string]IColumn)
}
