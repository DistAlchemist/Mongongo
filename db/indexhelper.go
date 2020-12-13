// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import (
	"os"
	"sort"

	"github.com/DistAlchemist/Mongongo/utils"
)

// IndexInfo ...
type IndexInfo struct {
	width     int64
	lastName  []byte
	firstName []byte
	offset    int64
}

// NewIndexInfo ...
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

func deserializeIndex(in *os.File) []*IndexInfo {
	indexList := make([]*IndexInfo, 0)
	columnIndexSize := readInt(in)
	start := getCurrentPos(in)
	for getCurrentPos(in) < start+int64(columnIndexSize) {
		indexList = append(indexList, indexInfoDeserialize(in))
	}
	return indexList
}

func indexFor(name []byte, indexList []*IndexInfo, reversed bool) int {
	// the index of the IndexInfo in which name will be found.
	// if the index is len(indexList), the name appears nowhere
	if len(name) == 0 && reversed {
		return len(indexList) - 1
	}
	// target := NewIndexInfo(name, name, 0, 0)
	index := sort.Search(len(indexList), func(i int) bool {
		return (string(name) <= string(indexList[i].lastName) &&
			string(name) >= string(indexList[i].firstName)) ||
			(string(name) <= string(indexList[i].firstName) &&
				string(name) <= string(indexList[i].lastName))
	})
	return index
}

func defreezeBloomFilter(file *os.File) *utils.BloomFilter {
	// size := readInt(file)
	bytes, _ := readBytes(file)
	return utils.BFSerializer.DeserializeB(bytes)
}
