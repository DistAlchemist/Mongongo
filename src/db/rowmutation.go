package db

import "strings"

// RowMutation definition
type RowMutation struct {
	TableName    string
	RowKey       string
	Modification map[string]ColumnFamily
}

// Add store columnFamilyName and columnName inside rowMutation
func (rm *RowMutation) Add(columnFamilyColumn, value string, timestamp int64) {
	cfColumn := strings.Split(columnFamilyColumn, ":")
	columnFamilyName, columnName := cfColumn[0], cfColumn[1]
	columnFamily := ColumnFamily{columnFamilyName}
	columnFamily.CreateColumn(columnName, value, timestamp)
	if rm.Modification == nil {
		rm.Modification = make(map[string]ColumnFamily)
	}
	rm.Modification[columnFamilyName] = columnFamily
}
