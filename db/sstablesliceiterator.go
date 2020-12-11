// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import (
	"log"
	"os"

	"gopkg.in/karalabe/cookiejar.v1/collections/deque"
)

// SSTableSliceIterator ...
type SSTableSliceIterator struct {
	reversed    bool
	startColumn []byte
	reader      *ColumnGroupReader
	nextValue   IColumn
	nextRead    bool
}

// NewSSTableSliceIterator ...
func NewSSTableSliceIterator(ssTable *SSTableReader, key string, startColumn []byte, reversed bool) *SSTableSliceIterator {
	s := &SSTableSliceIterator{}
	s.reversed = reversed
	decoratedKey := ssTable.partitioner.DecorateKey(key)
	// get key position in the data file
	position := ssTable.getPosition(decoratedKey)
	s.startColumn = startColumn
	if position >= 0 {
		s.reader = NewColumnGroupReader(ssTable, decoratedKey, position, startColumn, reversed)
	}
	s.nextValue = nil
	s.nextRead = false
	return s
}

func (s *SSTableSliceIterator) hasNext() bool {
	if s.nextRead == true {
		return s.nextValue != nil
	}
	s.nextRead = true
	s.nextValue = s.computeNext()
	return s.nextValue != nil
}

func (s *SSTableSliceIterator) next() IColumn {
	if s.nextRead == true {
		s.nextRead = false
		return s.nextValue
	}
	return s.computeNext()
}

func (s *SSTableSliceIterator) computeNext() IColumn {
	if s.reader == nil {
		return nil
	}
	for {
		column := s.reader.pollColumn()
		if column == nil {
			return nil
		}
		if s.isColumnNeeded(column) {
			return column
		}
	}
}

func (s *SSTableSliceIterator) isColumnNeeded(column IColumn) bool {
	if s.reversed {
		return len(s.startColumn) == 0 || column.getName() <= string(s.startColumn)
	}
	return column.getName() >= string(s.startColumn)
}

func (s *SSTableSliceIterator) close() {
	if s.reader != nil {
		s.reader.close()
	}
}

func (s *SSTableSliceIterator) getColumnFamily() *ColumnFamily {
	return s.reader.getEmptyColumnFamily()
}

// ColumnGroupReader finds the block for a starting
// column and returns blocks before/after it for each
// next call. This function assumes that the CF is
// sorted by name and exploits the name index
type ColumnGroupReader struct {
	emptyColumnFamily   *ColumnFamily
	indices             []*IndexInfo
	columnStartPosition int64
	file                *os.File
	curRangeIndex       int
	blockColumns        *deque.Deque
	reversed            bool
}

// NewColumnGroupReader ...
func NewColumnGroupReader(ssTable *SSTableReader, key string, position int64, startColumn []byte, reversed bool) *ColumnGroupReader {
	c := &ColumnGroupReader{}
	var err error
	c.file, err = os.Open(ssTable.getFilename())
	if err != nil {
		log.Fatal(err)
	}
	c.file.Seek(position, 0)
	keyInDisk, _ := readString(c.file)
	if key != keyInDisk {
		log.Fatal("key should equals to keyInDisk")
	}
	readInt(c.file)         // read row size
	skipBloomFilter(c.file) // skip bloom filter
	c.indices = deserializeIndex(c.file)

	c.emptyColumnFamily = CFSerializer.deserializeFromSSTableNoColumns(ssTable.makeColumnFamily(), c.file)
	readInt(c.file) // column count
	c.columnStartPosition = getCurrentPos(c.file)
	c.curRangeIndex = indexFor(startColumn, c.indices, reversed)
	c.reversed = reversed
	if reversed && c.curRangeIndex == len(c.indices) {
		c.curRangeIndex--
	}
	return c
}

func (c *ColumnGroupReader) getEmptyColumnFamily() *ColumnFamily {
	return c.emptyColumnFamily
}

func (c *ColumnGroupReader) close() {
	c.file.Close()
}

func (c *ColumnGroupReader) pollColumn() IColumn {
	if c.blockColumns.Size() == 0 {
		if c.getNextBlock() {
			return c.blockColumns.PopLeft().(IColumn)
		}
	}
	return c.blockColumns.PopLeft().(IColumn)
}

func (c *ColumnGroupReader) getNextBlock() bool {
	if c.curRangeIndex < 0 || c.curRangeIndex >= len(c.indices) {
		return false
	}
	// seek to the correct offset to the data, and
	// calculate the data size
	curColPosition := c.indices[c.curRangeIndex]
	c.file.Seek(c.columnStartPosition+curColPosition.offset, 0)
	for getCurrentPos(c.file) < c.columnStartPosition+curColPosition.offset+curColPosition.width {
		column := c.emptyColumnFamily.getColumnSerializer().deserialize(c.file)
		if c.reversed {
			c.blockColumns.PushLeft(column)
		} else {
			c.blockColumns.PushRight(column)
		}
	}
	if c.reversed {
		c.curRangeIndex--
	} else {
		c.curRangeIndex++
	}
	return true
}
