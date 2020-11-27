// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import (
	"log"
	"strings"
)

// RowMutation definition
type RowMutation struct {
	TableName    string
	RowKey       string
	Modification map[string]*ColumnFamily
}

// NewRowMutation creates a new row mutation
func NewRowMutation(tableName, rowKey string) RowMutation {
	rm := RowMutation{}
	rm.TableName = tableName
	rm.RowKey = rowKey
	return rm
}

// Add store columnFamilyName and columnName inside rowMutation
func (rm *RowMutation) Add(columnFamilyColumn, value string, timestamp int64) {
	cfColumn := strings.Split(columnFamilyColumn, ":")
	sz := len(cfColumn)
	if sz == 0 || sz == 1 || sz > 3 {
		log.Printf("Invalid format: %v. Must be <column family>:<column>\n", cfColumn)
	}
	columnFamilyName := cfColumn[0]
	var columnFamily *ColumnFamily
	if sz == 2 {
		columnName := cfColumn[1]
		columnFamily = NewColumnFamily(columnFamilyName, "Standard")
		columnFamily.CreateColumn(columnName, value, timestamp)
	} else if sz == 3 {
		columnName := cfColumn[1] + ":" + cfColumn[2]
		columnFamily = NewColumnFamily(columnFamilyName, "Super")
		columnFamily.CreateColumn(columnName, value, timestamp)

	}
	if rm.Modification == nil {
		rm.Modification = make(map[string]*ColumnFamily)
	}
	rm.Modification[columnFamilyName] = columnFamily
}
