// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package service

// SlicePredicate ...
type SlicePredicate struct {
	ColumnNames [][]byte
	SRange      SliceRange
}

// NewSlicePredicate ...
func NewSlicePredicate(columnNames [][]byte, sRange SliceRange) SlicePredicate {
	res := SlicePredicate{}
	res.ColumnNames = columnNames
	res.SRange = sRange
	return res
}
