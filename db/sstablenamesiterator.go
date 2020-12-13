// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import (
	"log"
	"os"
)

// SSTableNamesIterator ...
type SSTableNamesIterator struct {
	cf       *ColumnFamily
	curIndex int
	columns  [][]byte
	iter     []IColumn
}

// NewSSTableNamesIterator ...
func NewSSTableNamesIterator(sstable *SSTableReader, key string, columns [][]byte) *SSTableNamesIterator {
	r := &SSTableNamesIterator{}
	r.columns = columns
	r.curIndex = 0
	decoratedKey := sstable.partitioner.DecorateKey(key)
	position := sstable.getPosition(decoratedKey)
	if position < 0 {
		return r
	}
	file, err := os.Open(sstable.dataFileName)
	if err != nil {
		log.Fatal(err)
	}
	keyInDisk, _ := readString(file)
	if keyInDisk != decoratedKey {
		log.Fatal("keyInDisk should == decoratedKey")
	}
	readInt(file)
	// read the bloom filter that summarizing the columns
	bf := defreezeBloomFilter(file)
	filteredColumnNames := make([][]byte, len(columns))
	for _, name := range columns {
		if bf.IsPresent(string(name)) {
			filteredColumnNames = append(filteredColumnNames, name)
		}
	}
	if len(filteredColumnNames) == 0 {
		return r
	}
	indexList := deserializeIndex(file)
	cf := CFSerializer.deserializeFromSSTableNoColumns(sstable.makeColumnFamily(), file)
	readInt(file) // columncount
	ranges := make([]*IndexInfo, 0)
	// get the various column ranges we have to read
	for _, name := range filteredColumnNames {
		index := indexFor(name, indexList, false)
		if index == len(indexList) {
			continue
		}
		indexInfo := indexList[index]
		if string(name) < string(indexInfo.firstName) {
			continue
		}
		ranges = append(ranges, indexInfo)
	}
	// seek to the correct offset to the data
	columnBegin := getCurrentPos(file)
	// now read all the columns from the ranges
	for _, indexInfo := range ranges {
		file.Seek(columnBegin+indexInfo.offset, 0)
		for getCurrentPos(file) < columnBegin+indexInfo.offset+indexInfo.width {
			column := cf.getColumnSerializer().deserialize(file)
			// we check vs the origin list, not the filtered list
			// for efficiency
			if containsC(columns, column.getName()) {
				cf.addColumn(column)
			}
		}
	}
	file.Close()
	r.iter = cf.GetSortedColumns()
	return r
}

func (r *SSTableNamesIterator) getColumnFamily() *ColumnFamily {
	return r.cf
}

func (r *SSTableNamesIterator) computeNext() IColumn {
	if r.iter == nil || !r.hasNext() {
		return nil
	}
	r.curIndex++
	return r.iter[r.curIndex-1]
}

func (r *SSTableNamesIterator) hasNext() bool {
	return r.curIndex < len(r.iter)
}

func (r *SSTableNamesIterator) next() IColumn {
	return r.computeNext()
}

func (r *SSTableNamesIterator) close() {}
