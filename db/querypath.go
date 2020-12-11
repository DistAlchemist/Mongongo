// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

// QueryPath ...
type QueryPath struct {
	columnFamilyName string
	superColumnName  []byte
	columnName       []byte
}

// NewQueryPath ...
func NewQueryPath(columnFamilyName string, superColumnName, columnName []byte) *QueryPath {
	q := &QueryPath{}
	q.columnFamilyName = columnFamilyName
	q.superColumnName = superColumnName
	q.columnName = columnName
	return q
}

// NewQueryPathCF ...
func NewQueryPathCF(columnFamilyName string) *QueryPath {
	return NewQueryPath(columnFamilyName, nil, nil)
}
