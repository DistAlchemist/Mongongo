// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import "os"

// IndexInfo ...
type IndexInfo struct {
	width     int64
	lastName  []byte
	firstName []byte
	offset    int64
}

func NewIndexInfo(firstName []byte, lastName []byte, offset int64, width int64) *IndexInfo {
	r := &IndexInfo{}
	r.firstName = firstName
	r.lastName = lastName
	r.offset = offset
	r.width = width
	return r
}

func (r *IndexInfo) serialize(dos []byte) {
	writeBytesB(dos, r.firstName)
	writeBytesB(dos, r.lastName)
	writeInt64B(dos, r.offset)
	writeInt64B(dos, r.width)
}

func (r *IndexInfo) serializedSize() int {
	return 4 + len(r.firstName) + 4 + len(r.lastName) + 8 + 8
}

func indexInfoDeserialize(dis *os.File) *IndexInfo {
	r := &IndexInfo{}
	str, _ := readString(dis)
	r.firstName = []byte(str)
	str, _ = readString(dis)
	r.lastName = []byte(str)
	r.offset = readInt64(dis)
	r.width = readInt64(dis)
	return r
}
