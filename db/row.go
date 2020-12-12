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
	Table          string
	Key            string
	ColumnFamilies map[string]*ColumnFamily
	Size           int32
}

// NewRow init a Row with given key
func NewRow(key string) *Row {
	r := &Row{}
	r.Key = key
	r.ColumnFamilies = make(map[string]*ColumnFamily)
	r.Size = 0
	return r
}

// NewRowT init a Row with given table name and key name
func NewRowT(table, key string) *Row {
	r := &Row{}
	r.Table = table
	r.Key = key
	r.ColumnFamilies = make(map[string]*ColumnFamily)
	return r
}

func (r *Row) getColumnFamilies() map[string]*ColumnFamily {
	return r.ColumnFamilies
}

func (r *Row) addColumnFamily(cf *ColumnFamily) {
	r.ColumnFamilies[cf.ColumnFamilyName] = cf
	atomic.AddInt32(&r.Size, cf.getSize())
}

func (r *Row) toByteArray() []byte {
	buf := make([]byte, 0)
	// write key length
	b4 := make([]byte, 4)
	binary.BigEndian.PutUint32(b4, uint32(len(r.Key)))
	buf = append(buf, b4...)
	// write key string
	buf = append(buf, []byte(r.Key)...)
	// write cf size
	binary.BigEndian.PutUint32(b4, uint32(len(r.ColumnFamilies)))
	buf = append(buf, b4...)
	// write cf bytes
	if r.Size > 0 {
		for _, columnFamily := range r.ColumnFamilies {
			buf = append(buf, columnFamily.toByteArray()...)
		}
	}
	return buf
}

func (r *Row) clear() {
	r.ColumnFamilies = make(map[string]*ColumnFamily)
}

func rowSerialize(row *Row, dos []byte) {
	writeStringB(dos, row.Table)
	writeStringB(dos, row.Key)
	columnFamilies := row.getColumnFamilies()
	size := len(columnFamilies)
	writeInt32B(dos, int32(size))
	for _, cf := range columnFamilies {
		CFSerializer.serialize(cf, dos)
	}
}
