package db

// ColumnFamily definition
type ColumnFamily struct {
	ColumnFamilyName string
}

// CreateColumn setup a new column in columnFamily
func (cf *ColumnFamily) CreateColumn(columnName, value string, timestamp int64) {
	//
	column := Column{columnName, value, timestamp}
	cf.addColumn(columnName, column)
}

func (cf *ColumnFamily) addColumn(columnName string, column Column) {
	// TODO
}
