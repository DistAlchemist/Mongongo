// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import (
	"bytes"
	"encoding/binary"
	"log"
	"os"
	"strconv"
)

// Column stores name and value etc.
type Column struct {
	Name       string
	Value      string
	Timestamp  int64
	size       int32
	deleteMark bool
}

func (c Column) addColumn(column IColumn) {
	log.Printf("Invalid method: Column doesn't support addColumn\n")
}

func (c Column) getMarkedForDeleteAt() int64 {
	if c.isMarkedForDelete() == false {
		log.Fatal("column is not marked for delete")
	}
	return c.Timestamp
}

func (c Column) getLocalDeletionTime() int {
	return 0 // miao miao miao ?
}

func (c Column) isMarkedForDelete() bool {
	return c.deleteMark
}

func (c Column) getValue() []byte {
	return []byte(c.Value)
}

func (c Column) getSize() int32 {
	// size of a column:
	//  4 bytes for name length
	//  # bytes for name string bytes
	//  8 bytes for timestamp
	//  4 bytes for value byte array
	//  # bytes for value bytes
	return int32(4 + 8 + 4 + len(c.Name) + len(c.Value))
}

// delete deletes a Column
/*func (c Column) delete() {
	if c.isMarkedForDelete == false {
		c.isMarkedForDelete = true
		c.Value = ""
	}
}*/

func (c Column) repair(column Column) {
	if c.Timestamp < column.Timestamp {
		c.Value = column.Value
		c.Timestamp = column.Timestamp
	}
}

func (c Column) diff(column Column) Column {
	var columnDiff Column
	if c.Timestamp < column.Timestamp {
		columnDiff.Name = column.Name
		columnDiff.Value = column.Value
		columnDiff.Timestamp = column.Timestamp
	}
	return columnDiff
}

func getObjectCount() int {
	return 1
}

func (c Column) toString() string {
	var buffer bytes.Buffer
	buffer.WriteString(c.Name)
	buffer.WriteString(":")
	//buffer.WriteString(strconv.FormatBool(c.isMarkedForDelete))
	buffer.WriteString(":")
	buffer.WriteString(strconv.FormatInt(c.Timestamp, 10))
	buffer.WriteString(":")
	buffer.WriteString(strconv.FormatInt(int64(len(c.Value)), 10))
	buffer.WriteString(":")
	buffer.WriteString(c.Value)
	buffer.WriteString(":")
	return buffer.String()
}

func (c Column) digest() string {
	var buffer bytes.Buffer
	buffer.WriteString(c.Name)
	//buffer.WriteString(c.Seperator)
	buffer.WriteString(strconv.FormatInt(c.Timestamp, 10))
	return buffer.String()
}

// NewColumn constructs a Column
func NewColumn(name, value string, timestamp int64, deleteMark bool) Column {
	c := Column{}
	c.Name = name
	c.Value = value
	c.Timestamp = timestamp
	c.deleteMark = deleteMark
	return c
}

func (c Column) toByteArray() []byte {
	buf := make([]byte, 0)
	// write column name length
	b4 := make([]byte, 4)
	binary.BigEndian.PutUint32(b4, uint32(len(c.Name)))
	buf = append(buf, b4...)
	// write column name
	buf = append(buf, []byte(c.Name)...)
	// write deleteMark
	if c.deleteMark {
		buf = append(buf, byte(1))
	} else {
		buf = append(buf, byte(0))
	}
	// write timestamp
	b8 := make([]byte, 8)
	binary.BigEndian.PutUint64(b8, uint64(c.Timestamp))
	buf = append(buf, b8...)
	// write value length
	binary.BigEndian.PutUint32(b4, uint32(len(c.Value)))
	buf = append(buf, b4...)
	// write value bytes
	buf = append(buf, []byte(c.Value)...)
	return buf
}

func (c Column) serializedSize() uint32 {
	// 4 byte: length of column name
	// # bytes: column name bytes
	// 1 byte:  deleteMark
	// 8 bytes: timestamp
	// 4 bytes: length of value
	// # bytes: value bytes
	return uint32(4 + 1 + 8 + 4 + len(c.Name) + len(c.Value))
}

func (c Column) getObjectCount() int {
	return 1
}

func (c Column) timestamp() int64 {
	return c.Timestamp
}

func (c Column) comparePriority(o Column) int64 {
	if c.isMarkedForDelete() {
		// tombstone always wins ties
		if c.Timestamp < o.Timestamp {
			return -1
		}
		return 1
	}
	return c.Timestamp - o.Timestamp
}

func (c Column) putColumn(column IColumn) bool {
	// resolve the column by comparing timestamps.
	// if a newer value is being put, take the change.
	// else ignore.
	_, ok := column.(Column)
	if !ok {
		log.Fatal("Only Column objects should be put here")
	}
	if c.Name != column.getName() {
		log.Fatal("The name should match the name of the current column")
	}
	if c.Timestamp <= column.timestamp() {
		return true
	}
	return false
}

func (c Column) getName() string {
	return c.Name
}

func (c Column) getSubColumns() map[string]IColumn {
	log.Fatal("This operation is not supported on simple columns")
	return nil
}

func (c Column) mostRecentChangeAt() int64 {
	return c.Timestamp
}

// CSerializer ...
var CSerializer = NewColumnSerializer()

// ColumnSerializer ...
type ColumnSerializer struct{}

// NewColumnSerializer ...
func NewColumnSerializer() *ColumnSerializer {
	return &ColumnSerializer{}
}

func (c *ColumnSerializer) serialize(column IColumn, dos *os.File) {
	writeString(dos, column.getName())
	writeBool(dos, column.isMarkedForDelete())
	writeInt64(dos, column.timestamp())
	writeBytes(dos, column.getValue()) // will first write byte length, the bytes
}

func (c *ColumnSerializer) serializeB(column IColumn, dos []byte) {
	writeStringB(dos, column.getName())
	writeBoolB(dos, column.isMarkedForDelete())
	writeInt64B(dos, column.timestamp())
	writeBytesB(dos, column.getValue()) // will first write byte length, the bytes
}

func (c *ColumnSerializer) deserialize(dis *os.File) IColumn {
	name, _ := readString(dis)
	deleteMark, _ := readBool(dis)
	timestamp := readInt64(dis)
	value, _ := readBytes(dis)
	return NewColumn(name, string(value), timestamp, deleteMark)
}
