// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import (
	"log"
	"os"
)

// IteratingRow ...
type IteratingRow struct {
	key               string
	finishedAt        int64
	emptyColumnFamily *ColumnFamily
	file              *os.File
}

// NewIteratingRow ...
func NewIteratingRow(file *os.File, sstable *SSTableReader) *IteratingRow {
	r := &IteratingRow{}
	r.file = file
	r.key, _ = readString(file)
	dataSize := int64(readInt(file))
	dataStart := getCurrentPos(file)
	r.finishedAt = dataStart + dataSize
	skipBloomFilterAndIndex(file)
	r.emptyColumnFamily = CFSerializer.deserializeFromSSTableNoColumns(sstable.makeColumnFamily(), file)
	readInt(file) // read column count
	return r
}

func (r *IteratingRow) hasNext() bool {
	return r.finishedAt != getCurrentPos(r.file)
}

func (r *IteratingRow) next() IColumn {
	if r.finishedAt == getCurrentPos(r.file) {
		log.Fatal("reach end of row")
		return Column{}
	}
	return r.emptyColumnFamily.columnSerializer.deserialize(r.file)
}

func (r *IteratingRow) getEmptyColumnFamily() *ColumnFamily {
	return r.emptyColumnFamily
}

func (r *IteratingRow) skipRemaining() {
	r.file.Seek(r.finishedAt, 0)
}

func skipBloomFilterAndIndex(in *os.File) int {
	return skipBloomFilter(in) + skipIndex(in)
}

func skipBloomFilter(in *os.File) int {
	totalBytesRead := 0
	// size of the bloom filter
	size := readInt(in)
	totalBytesRead += 4
	// skip the serialized bloom filter
	curPos := getCurrentPos(in)
	n, err := in.Seek(curPos+int64(size), 0)
	if err != nil {
		log.Fatal(err)
	}
	if int64(size) != n {
		log.Fatal("reach EOF")
	}
	totalBytesRead += size
	return totalBytesRead
}

func skipIndex(file *os.File) int {
	// read only the column index list
	columnIndexSize := readInt(file)
	totalBytesRead := 4
	// skip the column index data
	curPos := getCurrentPos(file)
	n, err := file.Seek(curPos+int64(columnIndexSize), 0)
	if err != nil {
		log.Fatal(err)
	}
	if int64(columnIndexSize) != n {
		log.Fatal("read EOF")
	}
	totalBytesRead += columnIndexSize
	return totalBytesRead
}
