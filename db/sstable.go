// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import (
	"github.com/DistAlchemist/Mongongo/config"
	"github.com/DistAlchemist/Mongongo/utils"
)

/**
 * SSTable structure borrowed from Cassandra:
 * struct SSTable stores data on disk in sorted fashion.
 * However, the sorting is upto the application. This
 * class expects keys to be handed to it in sorted order.
 * SSTable is broken up into blocks where each block
 * contains 128 keys. At the end of the file, the block
 * index is written which contains the offsets to the keys
 * in the block. SSTable also maintains an index file to
 * which every 128th key is written with a pointer to the
 * block index which is the block that actually contains
 * the key. This index file is then read and maintained in
 * memory. SSTable is append only and immutable. SSTable
 * on disk looks as follows: (graph borrowed from Cassandra)
 *                  ------------------------
 *                 |------------------------|<-------|
 *                 |                        |        |  BLOCK-INDEX PTR
 *                 |                        |        |
 *                 |------------------------|--------
 *                 |------------------------|<-------|
 *                 |                        |        |
 *                 |                        |        |  BLOCK-INDEX PTR
 *                 |                        |        |
 *                 |------------------------|---------
 *                 |------------------------|<--------|
 *                 |                        |         |
 *                 |                        |         |
 *                 |                        |         | BLOCK-INDEX PTR
 *                 |                        |         |
 *                 |------------------------|         |
 *                 |------------------------|----------
 *                 |------------------------|-----------------> BLOOM-FILTER
 * version-info <--|----------|-------------|-------> relative offset to last block index.
 */
// SSTableTmpFile is the tmp file name for sstable
var (
	SSTableTmpFile      = "tmp"
	SSTIndexMetadataMap map[string][]KeyPositionInfo
	// every 128th key is an index
	SSTIndexInterval = 128
	// key associated with block index written to disk
	SSTBlockIndexKey = "BLOCK-INDEX"
	// position in SSTable after the first Block Index
	SSTPositionAfterFirstBlockIndex = 0
	// this map has the SSTable as key and a BloomFilter
	// as value. This BloomFilter will tell us if a key/
	// column pair is in the SSTable. If not, we can avoid
	// scanning it.
	SSTbfs = make(map[string]utils.BloomFilter)
	// maintains a touched set of keys
	SSTTouchCache = NewTouchKeyCache(config.TouchKeyCacheSize)
)

// KeyPositionInfo contains index key and its corresponding
// position in the data file. Binary search is performed on
// a list of these objects to lookup keys within the SSTable
// data file.
type KeyPositionInfo struct {
	key      string
	position int64
}

// SSTable is the struct for SSTable
type SSTable struct {
	dataFileName string
}

// NewSSTable initializes a SSTable
func NewSSTable(filename string) *SSTable {
	s := &SSTable{}
	// filename of the type:
	//  var/storage/data/<tableName>-<columnFamilyName>-<index>-Data.db
	s.dataFileName = filename
	_, ok := SSTIndexMetadataMap[s.dataFileName]
	if !ok {
		s.loadIndexFile()
	}
	return s
}

func (s *SSTable) loadIndexFile() {
	// TODO
}

func onSSTableStart(filenames []string) {
	for _, filename := range filenames {
		ssTable := NewSSTable(filename)
		ssTable.close()
	}
}

func (s *SSTable) close() {
	// Empty
}

// TouchKeyCache implements LRU cache
type TouchKeyCache struct {
	size int
}

// NewTouchKeyCache initializes a cache with given size
func NewTouchKeyCache(size int) *TouchKeyCache {
	t := &TouchKeyCache{}
	t.size = size
	return t
}
