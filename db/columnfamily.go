// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import (
	"encoding/binary"
	"sort"
	"sync/atomic"

	"github.com/DistAlchemist/Mongongo/config"
)

// ColumnFamily definition
type ColumnFamily struct {
	ColumnFamilyName  string
	ColumnType        string
	Factory           AColumnFactory
	Columns           map[string]IColumn
	size              int32
	deleteMark        bool
	localDeletionTime int
	markedForDeleteAt int64
	columnSerializer  IColumnSerializer
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
	if "Standard" == cf.ColumnType {
		cf.columnSerializer = NewColumnSerializer()
	} else {
		cf.columnSerializer = NewSuperColumnSerializer()
	}
	return cf
}

func createColumnFamily(tableName, cfName string) *ColumnFamily {
	columnType := config.GetColumnTypeTableName(tableName, cfName)
	return NewColumnFamily(cfName, columnType)
}

// CreateColumn setup a new column in columnFamily
func (cf *ColumnFamily) CreateColumn(columnName, value string, timestamp int64) {
	column := cf.Factory.createColumn(columnName, value, timestamp)
	cf.addColumn(column)
}

// If we find and old column that has the same
// name, then ask it to resolve itself, else
// we add the new column
func (cf *ColumnFamily) addColumn(column IColumn) {
	name := column.getName()
	oldColumn, ok := cf.Columns[name]
	if ok {
		_, yes := oldColumn.(SuperColumn)
		if yes { // is SuperColumn
			oldSize := oldColumn.getSize()
			oldColumn.putColumn(column)
			atomic.AddInt32(&cf.size, int32(oldColumn.getSize()-oldSize))
		} else {
			if oldColumn.(Column).comparePriority(column.(Column)) <= 0 {
				cf.Columns[name] = column
				atomic.AddInt32(&cf.size, int32(column.getSize()))
			}
		}
	} else {
		atomic.AddInt32(&cf.size, column.getSize())
		cf.Columns[name] = column
	}
}

// with query path as argument
// in most places the cf must be part of a query path but
// here it is ignored.
func (cf *ColumnFamily) addColumnQP(path *QueryPath, value string, timestamp int64, deleted bool) {
	var column IColumn
	if path.SuperColumnName == nil {
		column = NewColumn(string(path.ColumnName), value, timestamp, deleted)
	} else {
		column = NewSuperColumn(string(path.SuperColumnName))
		column.addColumn(NewColumn(string(path.ColumnFamilyName), value, timestamp, deleted))
	}
	cf.addColumn(column)
}

func (cf *ColumnFamily) isSuper() bool {
	return cf.ColumnType == "Super"
}

// IsSuper ...
func (cf *ColumnFamily) IsSuper() bool {
	return cf.ColumnType == "Super"
}

// GetColumn ...
func (cf *ColumnFamily) GetColumn(name string) IColumn {
	return cf.Columns[name]
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
	for _, column := range columns {
		cf.addColumn(column)
	}
}

func (cf *ColumnFamily) getMarkedForDeleteAt() int64 {
	return cf.markedForDeleteAt
}

func (cf *ColumnFamily) getLocalDeletionTime() int {
	return cf.localDeletionTime
}

func (cf *ColumnFamily) deleteCF(cf2 *ColumnFamily) {
	t := cf.localDeletionTime
	if t < cf2.localDeletionTime {
		t = cf2.localDeletionTime
	}
	m := cf.getMarkedForDeleteAt()
	if m < cf2.getMarkedForDeleteAt() {
		m = cf2.getMarkedForDeleteAt()
	}
	cf.delete(t, m)
}

func (cf *ColumnFamily) remove(columnName string) {
	delete(cf.Columns, columnName)
}

// GetSortedColumns ...
func (cf *ColumnFamily) GetSortedColumns() []IColumn {
	cnames := make([]string, 0)
	for name := range cf.Columns {
		cnames = append(cnames, name)
	}
	sort.Sort(ByKey(cnames))
	res := make([]IColumn, 0)
	for _, name := range cnames {
		res = append(res, cf.Columns[name])
	}
	return res
}

func (cf *ColumnFamily) cloneMeShallow() *ColumnFamily {
	c := NewColumnFamily(cf.ColumnFamilyName, cf.ColumnType)
	c.markedForDeleteAt = cf.markedForDeleteAt
	c.localDeletionTime = cf.localDeletionTime
	return c
}

func (cf *ColumnFamily) clear() {
	cf.Columns = make(map[string]IColumn)
}

func (cf *ColumnFamily) delete(localtime int, timestamp int64) {
	cf.localDeletionTime = localtime
	cf.markedForDeleteAt = timestamp
}

func (cf *ColumnFamily) getColumnSerializer() IColumnSerializer {
	return cf.columnSerializer
}
