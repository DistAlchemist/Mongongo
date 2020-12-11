// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import (
	"encoding/binary"
	"log"
	"math"
	"os"
	"sync/atomic"
)

// SuperColumn implements IColumn interface
type SuperColumn struct {
	Name              string
	Columns           map[string]IColumn
	deleteMark        bool
	size              int32
	Timestamp         int64
	localDeletionTime int
	markedForDeleteAt int64
}

func (sc SuperColumn) markForDeleteAt(localDeletionTime int, timestamp int64) {
	sc.localDeletionTime = localDeletionTime
	sc.markedForDeleteAt = timestamp
}

func (sc SuperColumn) getLocalDeletionTime() int {
	return sc.localDeletionTime
}

func (sc SuperColumn) getMarkedForDeleteAt() int64 {
	return sc.markedForDeleteAt
}

func (sc SuperColumn) isMarkedForDelete() bool {
	return sc.deleteMark
}

func (sc SuperColumn) getValue() []byte {
	log.Fatal("super column doesn't support getValue")
	return []byte{}
}

func (sc SuperColumn) addColumn(column IColumn) {
	name := column.getName()
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
	sc.localDeletionTime = math.MinInt32
	sc.markedForDeleteAt = math.MinInt64
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

// Go through each subComlun. If it exists then resolve.
// Else create
func (sc SuperColumn) putColumn(column IColumn) bool {
	_, ok := column.(SuperColumn)
	if !ok {
		log.Fatal("Only Super column objects should be put here")
	}
	if sc.Name != column.getName() {
		log.Fatal("The name should match the name of the current super column")
	}
	columns := column.getSubColumns()
	for _, subColumn := range columns {
		sc.addColumn(subColumn)
	}
	if column.getMarkedForDeleteAt() > sc.markedForDeleteAt {
		sc.markForDeleteAt(column.getLocalDeletionTime(), column.getMarkedForDeleteAt())
	}
	return false
}

func (sc SuperColumn) cloneMeShallow() SuperColumn {
	s := NewSuperColumn(sc.Name)
	s.markForDeleteAt(sc.localDeletionTime, sc.markedForDeleteAt)
	return s
}

func (sc SuperColumn) getSubColumns() map[string]IColumn {
	return sc.Columns
}

func (sc SuperColumn) mostRecentChangeAt() int64 {
	res := int64(math.MinInt64)
	for _, column := range sc.Columns {
		if column.mostRecentChangeAt() > res {
			res = column.mostRecentChangeAt()
		}
	}
	return res
}

// SCSerializer ...
var SCSerializer = NewSuperColumnSerializer()

// SuperColumnSerializer ...
type SuperColumnSerializer struct{}

// NewSuperColumnSerializer ...
func NewSuperColumnSerializer() *SuperColumnSerializer {
	return &SuperColumnSerializer{}
}

func (s *SuperColumnSerializer) serialize(column IColumn, dos *os.File) {
	superColumn := column.(SuperColumn)
	writeString(dos, column.getName())
	writeInt(dos, superColumn.getLocalDeletionTime())
	writeInt64(dos, superColumn.getMarkedForDeleteAt())
	columns := column.getSubColumns()
	for _, subColumn := range columns {
		CSerializer.serialize(subColumn, dos)
	}
}

func (s *SuperColumnSerializer) serializeB(column IColumn, dos []byte) {
	superColumn := column.(SuperColumn)
	writeStringB(dos, column.getName())
	writeIntB(dos, superColumn.getLocalDeletionTime())
	writeInt64B(dos, superColumn.getMarkedForDeleteAt())
	columns := column.getSubColumns()
	for _, subColumn := range columns {
		CSerializer.serializeB(subColumn, dos)
	}
}

func (s *SuperColumnSerializer) deserialize(dis *os.File) IColumn {
	name, _ := readString(dis)
	superColumn := NewSuperColumn(name)
	localDeletionTime := readInt(dis)
	timestamp := readInt64(dis)
	superColumn.markForDeleteAt(localDeletionTime, timestamp)
	// read the number of columns
	size := readInt(dis)
	for i := 0; i < size; i++ {
		subColumn := CSerializer.deserialize(dis)
		superColumn.addColumn(subColumn)
	}
	return superColumn
}
