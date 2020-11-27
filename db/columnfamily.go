// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

// ColumnFamily definition
type ColumnFamily struct {
	ColumnFamilyName string
	ColumnType       string
	Factory          AColumnFactory
	Columns          map[string]IColumn
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
	return cf
}

// CreateColumn setup a new column in columnFamily
func (cf *ColumnFamily) CreateColumn(columnName, value string, timestamp int64) {
	column := cf.Factory.createColumn(columnName, value, timestamp)
	cf.addColumn(columnName, column)
}

func (cf *ColumnFamily) addColumn(columnName string, column IColumn) {
	if cf.Columns == nil {
		cf.Columns = make(map[string]IColumn)
	}
	cf.Columns[columnName] = column
}
