// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import (
	"encoding/binary"
	"log"
	"sync/atomic"
)

// SuperColumn implements IColumn interface
type SuperColumn struct {
	Name       string
	Columns    map[string]IColumn
	deleteMark bool
	size       int32
	Timestamp  int64
}

func (sc SuperColumn) addColumn(name string, column IColumn) {
	oldColumn, ok := sc.Columns[name]
	if !ok {
		sc.Columns[name] = column
		atomic.AddInt32(&sc.size, column.getSize())
	} else {
		if oldColumn.timestamp() <= column.timestamp() {
			sc.Columns[name] = column
			delta := int32(-1 * oldColumn.getSize())
			// subtruct the size of the oldColumn
			atomic.AddInt32(&sc.size, delta)
			atomic.AddInt32(&sc.size, int32(column.getSize()))
		}
	}
}

// NewSuperColumn constructs a SuperColun
func NewSuperColumn(name string) SuperColumn {
	sc := SuperColumn{}
	sc.Name = name
	sc.Columns = make(map[string]IColumn)
	sc.deleteMark = false
	sc.size = 0
	return sc
}

func (sc SuperColumn) getSize() int32 {
	return sc.size
}

func (sc SuperColumn) getObjectCount() int {
	return 1 + len(sc.Columns)
}

func (sc SuperColumn) toByteArray() []byte {
	buf := make([]byte, 0)
	// write supercolumn name length
	b4 := make([]byte, 4)
	binary.BigEndian.PutUint32(b4, uint32(len(sc.Name)))
	buf = append(buf, b4...)
	// write supercolumn name bytes
	buf = append(buf, []byte(sc.Name)...)
	// write deleteMark
	if sc.deleteMark {
		buf = append(buf, byte(1))
	} else {
		buf = append(buf, byte(0))
	}
	// write column size
	binary.BigEndian.PutUint32(b4, uint32(len(sc.Columns)))
	buf = append(buf, b4...)
	// write subcolumns total size, used to skip over
	// all these columns if we are not interested in
	// this super column
	binary.BigEndian.PutUint32(b4, sc.getSizeOfAllColumns())
	buf = append(buf, b4...)
	for _, column := range sc.Columns {
		buf = append(buf, column.toByteArray()...)
	}
	return buf
}

func (sc SuperColumn) getSizeOfAllColumns() uint32 {
	res := uint32(0)
	for _, column := range sc.Columns {
		res += column.serializedSize()
	}
	return res
}

func (sc SuperColumn) serializedSize() uint32 {
	// 4 bytes: super column name length
	// # bytes: super column name bytes
	// 1 byte:  deleteMark
	// 4 bytes: number of sub-columns
	// 4 bytes: size of sub-columns
	// # bytes: size of all sub-columns
	return uint32(4+1+4+4+len(sc.Name)) + sc.getSizeOfAllColumns()
}

func (sc SuperColumn) timestamp() int64 {
	return sc.Timestamp
}

func (sc SuperColumn) getName() string {
	return sc.Name
}

func (sc SuperColumn) putColumn(column IColumn) bool {
	_, ok := column.(SuperColumn)
	if !ok {
		log.Fatal("Only Super column objects should be put here")
	}
	if sc.Name != column.getName() {
		log.Fatal("The name should match the name of the current super column")
	}
	columns := column.getSubColumns()
	for name, subColumn := range columns {
		sc.addColumn(name, subColumn)
	}
	return false
}

func (sc SuperColumn) getSubColumns() map[string]IColumn {
	return sc.Columns
}
