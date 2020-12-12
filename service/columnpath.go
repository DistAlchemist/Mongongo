// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package service

// ColumnPath ...
type ColumnPath struct {
	ColumnFamily string
	SuperColumn  []byte
	Column       []byte
}

// NewColumnPath ...
func NewColumnPath(cf string, sc, c []byte) ColumnPath {
	res := ColumnPath{}
	res.ColumnFamily = cf
	res.SuperColumn = sc
	res.Column = c
	return res
}
