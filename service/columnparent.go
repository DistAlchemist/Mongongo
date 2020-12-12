// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package service

// ColumnParent ...
type ColumnParent struct {
	ColumnFamily string
	SuperColumn  []byte
}

// NewColumnParent ...
func NewColumnParent(columnFamily string, superColumn []byte) ColumnParent {
	res := ColumnParent{}
	res.ColumnFamily = columnFamily
	res.SuperColumn = superColumn
	return res
}
