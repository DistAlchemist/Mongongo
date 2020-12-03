// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import "sync/atomic"

// Row for key and cf
type Row struct {
	key            string
	columnFamilies map[string]*ColumnFamily
	size           int32
}

// NewRow init a Row with given key
func NewRow(key string) *Row {
	r := &Row{}
	r.key = key
	r.columnFamilies = make(map[string]*ColumnFamily)
	r.size = 0
	return r
}

func (r *Row) addColumnFamily(cf *ColumnFamily) {
	r.columnFamilies[cf.ColumnFamilyName] = cf
	atomic.AddInt32(&r.size, cf.getSize())
}
