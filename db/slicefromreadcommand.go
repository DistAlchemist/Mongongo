// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

// SliceFromReadCommand ...
type SliceFromReadCommand struct {
	*AReadCommand
	Start    []byte
	Finish   []byte
	Reversed bool
	Count    int
}

// NewSliceFromReadCommand ...
func NewSliceFromReadCommand(table, key string, path QueryPath, start, finish []byte,
	reversed bool, count int) *SliceFromReadCommand {
	r := &SliceFromReadCommand{}
	r.AReadCommand = NewAReadCommand(table, key, path, cmdTypeGetSlice)
	r.Start = start
	r.Finish = finish
	r.Reversed = reversed
	r.Count = count
	return r
}

// GetKey ...
func (s *SliceFromReadCommand) GetKey() string {
	return s.Key
}

// GetQPath ...
func (s *SliceFromReadCommand) GetQPath() QueryPath {
	return s.QPath
}

// GetCFName ...
func (s *SliceFromReadCommand) GetCFName() string {
	return s.QPath.ColumnFamilyName
}

// GetTable ...
func (s *SliceFromReadCommand) GetTable() string {
	return s.Table
}

// GetRow ...
func (s *SliceFromReadCommand) GetRow(table *Table) *Row {
	return table.getRow(NewSliceQueryFilter(s.Key, &s.QPath,
		s.Start, s.Finish, s.Reversed, s.Count))
}
