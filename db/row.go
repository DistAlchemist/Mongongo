// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import (
	"encoding/binary"
	"sync/atomic"
)

// Row for key and cf
type Row struct {
	table          string
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

// NewRowT init a Row with given table name and key name
func NewRowT(table, key string) *Row {
	r := &Row{}
	r.table = table
	r.key = key
	r.columnFamilies = make(map[string]*ColumnFamily)
	return r
}

func (r *Row) getColumnFamilies() map[string]*ColumnFamily {
	return r.columnFamilies
}

func (r *Row) addColumnFamily(cf *ColumnFamily) {
	r.columnFamilies[cf.ColumnFamilyName] = cf
	atomic.AddInt32(&r.size, cf.getSize())
}

func (r *Row) toByteArray() []byte {
	buf := make([]byte, 0)
	// write key length
	b4 := make([]byte, 4)
	binary.BigEndian.PutUint32(b4, uint32(len(r.key)))
	buf = append(buf, b4...)
	// write key string
	buf = append(buf, []byte(r.key)...)
	// write cf size
	binary.BigEndian.PutUint32(b4, uint32(len(r.columnFamilies)))
	buf = append(buf, b4...)
	// write cf bytes
	if r.size > 0 {
		for _, columnFamily := range r.columnFamilies {
			buf = append(buf, columnFamily.toByteArray()...)
		}
	}
	return buf
}

func (r *Row) clear() {
	r.columnFamilies = make(map[string]*ColumnFamily)
}

func rowSerialize(row *Row, dos []byte) {
	writeStringB(dos, row.table)
	writeStringB(dos, row.key)
	columnFamilies := row.getColumnFamilies()
	size := len(columnFamilies)
	writeInt32B(dos, int32(size))
	for _, cf := range columnFamilies {
		CFSerializer.serialize(cf, dos)
	}
}
