// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

// SliceByNamesReadCommand ...
type SliceByNamesReadCommand struct {
	*AReadCommand
	columnNames [][]byte
}

// NewSliceByNamesReadCommand ...
func NewSliceByNamesReadCommand(table, key string, path QueryPath, columnNames [][]byte) *SliceByNamesReadCommand {
	r := &SliceByNamesReadCommand{}
	r.AReadCommand = NewAReadCommand(table, key, path, cmdTypeGetSliceByNames)
	r.columnNames = make([][]byte, 0)
	r.columnNames = append(r.columnNames, columnNames...)
	return r
}

// GetKey ...
func (s *SliceByNamesReadCommand) GetKey() string {
	return s.Key
}

// GetQPath ...
func (s *SliceByNamesReadCommand) GetQPath() QueryPath {
	return s.QPath
}

// GetCFName ...
func (s *SliceByNamesReadCommand) GetCFName() string {
	return s.QPath.ColumnFamilyName
}

// GetTable ...
func (s *SliceByNamesReadCommand) GetTable() string {
	return s.Table
}

// GetRow ...
func (s *SliceByNamesReadCommand) GetRow(table *Table) *Row {
	return table.getRow(NewNamesQueryFilterS(s.Key, &s.QPath, s.columnNames))
}
