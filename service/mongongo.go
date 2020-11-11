package service

// Mongongo expose the interface of operations
type Mongongo struct {
	// Mongongo struct
}

// Insert updates tableNmae.key.columnFamilyColumn with (cellData, timeStamp)
func (mg *Mongongo) Insert(tableName, key, columnFamilyColumn, cellData string, timeStamp int64) {
	//
}

// GetColumn get the value of tableName.key.columnFamilyColumn
func (mg *Mongongo) GetColumn(tableName, key, columnFamilyColumn string) {
	//
}

// Remove delete the value of tableName.key.columnFamilyColumn (lazily)
func (mg *Mongongo) Remove(tableName, key, columnFamilyColumn string) {
	//
}
