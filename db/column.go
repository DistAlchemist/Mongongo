// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import (
	"bytes"
	"log"
	"strconv"
)

// Column stores name and value etc.
type Column struct {
	Name      string
	Value     string
	Timestamp int64
	//isMarkedForDelete bool
}

func (c Column) addColumn(name string, column IColumn) {
	log.Printf("Invalid method: Column doesn't support addColumn\n")
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
func (c Column) putColumn(column Column) bool {
	if c.Name != column.Name {
		log.Printf("The name should match the name of the current column or super column")
	}
	if c.Timestamp <= column.Timestamp {
		return true
	}
	return false
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
func NewColumn(name, value string, timestamp int64) Column {
	c := Column{}
	c.Name = name
	c.Value = value
	c.Timestamp = timestamp
	//c.isMarkedForDelete = false
	return c
}
