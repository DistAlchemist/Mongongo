// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

// QueryPath ...
type QueryPath struct {
	ColumnFamilyName string
	SuperColumnName  []byte
	ColumnName       []byte
}

// NewQueryPath ...
func NewQueryPath(columnFamilyName string, superColumnName, columnName []byte) *QueryPath {
	q := &QueryPath{}
	q.ColumnFamilyName = columnFamilyName
	q.SuperColumnName = superColumnName
	q.ColumnName = columnName
	return q
}

// NewQueryPathCF ...
func NewQueryPathCF(columnFamilyName string) *QueryPath {
	return NewQueryPath(columnFamilyName, nil, nil)
}
