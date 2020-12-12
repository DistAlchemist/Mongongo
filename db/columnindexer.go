// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import (
	"log"

	"github.com/DistAlchemist/Mongongo/config"
	"github.com/DistAlchemist/Mongongo/utils"
)

// CIndexer ...
var CIndexer = &ColumnIndexer{}

// ColumnIndexer ...
type ColumnIndexer struct{}

func (c *ColumnIndexer) serialize(columnFamily *ColumnFamily, dos []byte) {
	// currently it is sorted by key string
	columns := columnFamily.GetSortedColumns()
	bf := c.createColumnBloomFilter(columns)
	// write out the bloom filter
	buf := make([]byte, 0)
	utils.BFSerializer.SerializeB(bf, buf)
	// write the length of the serialized bloom filter
	// and write the serialized bytes. 2 in 1 :)
	writeBytesB(dos, buf)
	// do the indexing
	c.doIndexing(columns, dos)
}

func (c *ColumnIndexer) createColumnBloomFilter(columns []IColumn) *utils.BloomFilter {
	columnCount := 0
	for _, column := range columns {
		columnCount += column.getObjectCount()
	}
	bf := utils.NewBloomFilter(columnCount, 4)
	for _, column := range columns {
		bf.Fill(column.getName())
		// if this is super column type
		// we need to get the subColumns too
		_, ok := column.(SuperColumn)
		if ok {
			subColumns := column.GetSubColumns()
			for _, subColumn := range subColumns {
				bf.Fill(subColumn.getName())
			}
		}
	}
	return bf
}

func (c *ColumnIndexer) doIndexing(columns []IColumn, dos []byte) {
	// Given the collection of columns in the column family,
	// the name index is generated and written into the provided
	// stream
	if len(columns) == 0 {
		// empty write index size 0
		writeIntB(dos, 0)
		return
	}
	// indexList maintains a list of ColumnIndexInfo objects
	// for the columns in this column family. The key is the
	// column name and the position is the relative offset of
	// that column name from the start of the list. Doing this
	// so that we don't read all the columns into memory.
	indexList := make([]*IndexInfo, 0)
	endPosition, startPosition := 0, -1
	indexSizeInBytes := 0
	var column, firstColumn IColumn
	// column offsets at the right thresholds into the index map
	for _, column = range columns {
		if firstColumn == nil {
			firstColumn = column
			startPosition = endPosition
		}
		endPosition += int(column.serializedSize())
		// if we hit the column index size that
		// we have to index after, go ahead and index it
		if endPosition-startPosition >= config.GetColumnIndexSize() {
			cIndexInfo := NewIndexInfo([]byte(firstColumn.getName()),
				[]byte(column.getName()),
				int64(startPosition),
				int64(endPosition-startPosition))
			indexList = append(indexList, cIndexInfo)
			indexSizeInBytes += cIndexInfo.serializedSize()
			firstColumn = nil
		}
	}
	// the last column may have fallen on an index boundary
	// already. if not, index it explicitly.
	if len(indexList) == 0 || string(indexList[len(indexList)-1].lastName) != column.getName() {
		cIndexInfo := NewIndexInfo([]byte(firstColumn.getName()),
			[]byte(column.getName()),
			int64(startPosition),
			int64(endPosition-startPosition))
		indexList = append(indexList, cIndexInfo)
	}
	if indexSizeInBytes <= 0 {
		log.Fatal("index size should > 0")
	}
	writeIntB(dos, indexSizeInBytes)
	for _, cIndexInfo := range indexList {
		cIndexInfo.serialize(dos)
	}
}
