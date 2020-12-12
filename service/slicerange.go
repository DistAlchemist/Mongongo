// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package service

// SliceRange ...
type SliceRange struct {
	Start    []byte
	Finish   []byte
	Reversed bool
	Count    int
}

// NewSliceRange ...
func NewSliceRange(start, finish []byte, reversed bool, count int) SliceRange {
	res := SliceRange{}
	res.Start = start
	res.Finish = finish
	res.Reversed = reversed
	res.Count = count
	return res
}
