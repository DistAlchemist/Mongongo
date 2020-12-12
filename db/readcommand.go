// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

var (
	cmdTypeGetSliceByNames byte = 1
	cmdTypeGetSlice        byte = 2
)

// ReadCommand ...
type ReadCommand interface {
	GetKey() string
	GetQPath() QueryPath
	GetCFName() string
	GetTable() string
	GetRow(table *Table) *Row
}

// AReadCommand ...
type AReadCommand struct {
	Table       string
	Key         string
	QPath       QueryPath
	CommandType byte
}

// NewAReadCommand ...
func NewAReadCommand(table, key string, queryPath QueryPath, cmdType byte) *AReadCommand {
	r := &AReadCommand{}
	r.Table = table
	r.Key = key
	r.QPath = queryPath
	r.CommandType = cmdType
	return r
}
