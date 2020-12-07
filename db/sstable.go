// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import (
	"encoding/binary"
	"log"
	"os"
	"sort"

	"github.com/willf/bitset"

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
 *                +------------------+
 *                |------------------|<--+
 *                |                  |   |  BLOCK-INDEX PTR
 *                |                  |   |
 *                |------------------|---+
 *                |------------------|<--+
 *                |                  |   |
 *                |                  |   |  BLOCK-INDEX PTR
 *                |                  |   |
 *                |------------------|---+
 *                |------------------|<--+
 *                |                  |   |
 *                |                  |   |
 *                |                  |   | BLOCK-INDEX PTR
 *                |                  |   |
 *                |------------------|   |
 *                |------------------|---+
 *                |------------------|---> BLOOM-FILTER
 * version-info <-+-------|----------+-> relative offset to last block index.
 */
var (
	// SSTableTmpFile is the tmp file name for sstable
	SSTableTmpFile      = "tmp"
	SSTVersion          = int64(0)
	SSTIndexMetadataMap map[string][]*KeyPositionInfo
	// every 128th key is an index
	SSTIndexInterval = 128
	// key associated with block index written to disk
	SSTBlockIndexKey = "BLOCK-INDEX"
	// position in SSTable after the first Block Index
	SSTPositionAfterFirstBlockIndex = int64(0)
	// this map has the SSTable as key and a BloomFilter
	// as value. This BloomFilter will tell us if a key/
	// column pair is in the SSTable. If not, we can avoid
	// scanning it.
	SSTbfs = make(map[string]*utils.BloomFilter)
	// maintains a touched set of keys
	SSTTouchCache = NewTouchKeyCache(config.TouchKeyCacheSize)
	bfMarker      = "Bloom-Filter"
	SSTBlkIdxKey  = "BLOCK-INDEX"
)

// KeyPositionInfo contains index key and its corresponding
// position in the data file. Binary search is performed on
// a list of these objects to lookup keys within the SSTable
// data file.
type KeyPositionInfo struct {
	key      string
	position int64
}

// NewKeyPositionInfo ...
func NewKeyPositionInfo(key string, position int64) *KeyPositionInfo {
	k := &KeyPositionInfo{}
	k.key = key
	k.position = position
	return k
}

// SSTable is the struct for SSTable
type SSTable struct {
	dataFileName       string
	dataWriter         *os.File
	blockIndex         map[string]*BlockMetadata
	blockIndexes       []map[string]*BlockMetadata
	lastWrittenKey     string
	indexKeysWritten   int
	indexInterval      int
	firstBlockPosition int64
}

// NewSSTable initializes a SSTable
func NewSSTable(filename string) *SSTable {
	s := &SSTable{}
	s.indexKeysWritten = 0
	s.lastWrittenKey = ""
	s.indexInterval = 128
	// filename of the type:
	//  var/storage/data/<tableName>-<columnFamilyName>-<index>-Data.db
	s.dataFileName = filename
	_, ok := SSTIndexMetadataMap[s.dataFileName]
	if !ok {
		s.loadIndexFile() // mainly load bloom filter and index file
	}
	return s
}

func (s *SSTable) loadIndexFile() {
	// filename of the type:
	//  var/storage/data/<tableName>-<columnFamilyName>-<index>-Data.db
	file, err := os.Open(s.dataFileName)
	if err != nil {
		log.Fatal(err)
	}
	fileInfo, err := file.Stat()
	size := fileInfo.Size()
	s.loadBloomFilter(file, size)
	// start building index
	// the first block index position is stored
	// at the 16 bytes before the end of the file
	file.Seek(size-16, 0)
	b8 := make([]byte, 8)
	firstBlockIndexPosition := readInt64(file, b8)
	keyPositionInfos := make([]*KeyPositionInfo, 0)
	SSTIndexMetadataMap[s.dataFileName] = keyPositionInfos
	nextPosition := size - 16 - firstBlockIndexPosition
	file.Seek(nextPosition, 0)
	// the structure of an index block is as follows:
	//  * key(string) -> block key "BLOCK-INDEX"
	//  * blockIndexSize int32: block index size
	//  * numKeys int32: # of keys in the block
	//  for i in range numKeys:
	//    * keySize int32: lengh of keyInBlock, work around..
	//    * keyInBlock string (if i==0, this is the largest key)
	//    * keyOffset int64: relative offset in the block
	//    * dataSize int64: size of data for that key
	// The goal is to obtain KeyPositionInfo:
	//    pair: (largestKeyInBlock, indexBlockPosition)
	// The code below is really an ugly workaround....
	b4 := make([]byte, 4)
	var currentPosition int64
	for {
		currentPosition = nextPosition
		b11 := make([]byte, 11)
		blockIdxKey := readBlockIdxKey(file, b11)
		nextPosition -= 11
		if blockIdxKey != SSTBlkIdxKey {
			log.Printf("Done reading the block indexes\n")
			break
		}
		readInt32(file, b4) // read block index size
		nextPosition -= 4
		numKeys := readInt32(file, b4)
		nextPosition -= 4
		for i := int32(0); i < numKeys; i++ {
			keyInBlock, size := readString(file)
			nextPosition -= size
			if i == 0 {
				keyPositionInfos = append(keyPositionInfos,
					&KeyPositionInfo{keyInBlock, currentPosition})
			}
			readInt64(file, b8) // read relative offset
			readInt64(file, b8) // read dataSize
			nextPosition -= 16
		}
	}
	// should also sort KeyPositionInfos, but I omit it. :)
}

func readString(file *os.File) (string, int64) {
	b4 := make([]byte, 4)
	size := int(readInt32(file, b4))
	bs := make([]byte, size)
	return readBlockIdxKey(file, bs), int64(size + 4)
}

func readBlockIdxKey(file *os.File, b []byte) string {
	n, err := file.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	if n != len(b) {
		log.Fatal("should read len(b) byte for block index key")
	}
	return string(b)
}

func readInt64(file *os.File, b8 []byte) int64 {
	n, err := file.Read(b8)
	if err != nil {
		log.Fatal(err)
	}
	if n != 8 {
		log.Fatal("should read 8 bytes")
	}
	return int64(binary.BigEndian.Uint64(b8))
}

func readInt32(file *os.File, b4 []byte) int32 {
	n, err := file.Read(b4)
	if err != nil {
		log.Fatal(err)
	}
	if n != 4 {
		log.Fatal("should read 4 bytes")
	}
	return int32(binary.BigEndian.Uint32(b4))
}

func readUint64(file *os.File, b8 []byte) uint64 {
	n, err := file.Read(b8)
	if err != nil {
		log.Fatal(err)
	}
	if n != 8 {
		log.Fatal("should read 8 bytes")
	}
	return binary.BigEndian.Uint64(b8)
}

func (s *SSTable) loadBloomFilter(file *os.File, size int64) {
	if _, ok := SSTbfs[s.dataFileName]; ok {
		return // bloom filter already exists in memory
	}
	// the last 8 bytes form a int64 denoting
	// relative position of bloom filter
	file.Seek(size-8, 0)
	b8 := make([]byte, 8)
	_, err := file.Read(b8)
	if err != nil {
		log.Fatal(err)
	}
	position := int64(binary.BigEndian.Uint64(b8))
	// seek to the position of bloom filter
	file.Seek(size-8-position, 0)
	// the contents of bf are as follows:
	// (optional) a string representing key, should be "Bloom-Filter"
	// total datasize, int64
	// count, int32 (i.e. # of elements already stored in bf)
	// hashes, int32
	// size, int32
	// bitset, BitSet, stored as []uint64
	// Start decoding!
	n, err := file.Read(b8)
	if err != nil {
		log.Fatal(err)
	}
	if n != 8 {
		log.Fatal("should read 8 bytes")
	}
	// don't need this variable
	// totalDataSize := int64(binary.BigEndian.Uint64(b8))
	b4 := make([]byte, 4)
	count := readInt32(file, b4)
	// read hashes: the number of hash functions
	hashes := readInt32(file, b4)
	// read size: the number of bits of BitSet
	bitsize := readInt32(file, b4)
	// convert to number of uint64
	num8byte := (bitsize-1)/64 + 1
	buf := make([]uint64, num8byte)
	for i := int32(0); i < num8byte; i++ {
		buf = append(buf, readUint64(file, b8))
	}
	bs := bitset.From(buf)
	SSTbfs[s.dataFileName] = utils.NewBloomFilterS(count, hashes, bitsize, bs)

	// reader := bufio.NewReader(file)
	// key, err := reader.ReadString(' ') // read key
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if key == bfMarker {
	// 	reader.ReadString(' ') // read total data size
	// 	_, ok := SSTbfs[s.dataFileName]
	// 	if !ok {
	// 		// read the count of the BloomFilter
	// 		// i.e. the number of elements
	// 		countStr, err := reader.ReadString(' ')
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		count := strconv.Atoi(countStr)
	// 		// read number of hash functions
	// 		hashesStr, err := reader.ReadString(' ')
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		hashes := strconv.Atoi(hashesStr)
	// 		// read the size of bloom filter
	// 		sizeStr, err := reader.ReadString(' ')
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		size := strconv.Atoi(sizeStr)
	// 		// read bitset
	// 		buf := readBitSet(reader)
	// 		bs := bitset.From(buf)
	// 		SSTbfs[s.dataFileName] = utils.NewBloomFilterS(count, hashes, size, bs)
	// 	}
	// }
}

func onSSTableStart(filenames []string) {
	for _, filename := range filenames {
		ssTable := NewSSTable(filename)
		ssTable.close()
	}
}

func (s *SSTable) close() {
	s.closeByte(make([]byte, 0), 0)
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

// NewSSTableP is used for DB writes into the SSTable
// Use this version to write to the SSTable
func NewSSTableP(directory, filename, pType string) *SSTable {
	s := &SSTable{}
	s.dataFileName = directory + string(os.PathSeparator) +
		filename + "-Data.db"
	var err error
	s.dataWriter, err = os.Create(s.dataFileName)
	if err != nil {
		log.Fatal(err)
	}
	SSTPositionAfterFirstBlockIndex = 0
	s.initBlockIndex(pType)
	s.blockIndexes = make([]map[string]*BlockMetadata, 0)
	return s
}

func (s *SSTable) initBlockIndex(pType string) {
	// TODO make ordered map
	switch pType {
	case config.Ophf:
		s.blockIndex = make(map[string]*BlockMetadata)
	default:
		s.blockIndex = make(map[string]*BlockMetadata)
	}
}

func (s *SSTable) beforeAppend(hash string) int64 {
	if hash == "" {
		log.Fatal("hash value shouldn't be empty")
	}
	if s.lastWrittenKey != "" {
		previousKey := s.lastWrittenKey
		if hash < previousKey {
			log.Printf("Last written key: %v\n", previousKey)
			log.Printf("Current key: %v\n", hash)
			log.Printf("Writing into file: %v\n", s.dataFileName)
			log.Fatal("Keys must be written in ascending order.")
		}
	}
	currentPos := SSTPositionAfterFirstBlockIndex
	if s.lastWrittenKey != "" {
		currentPos, err := s.dataWriter.Seek(0, 0)
		if err != nil {
			log.Fatal(err)
		}
		s.dataWriter.Seek(currentPos, 0)
	}
	return currentPos
}

func (s *SSTable) afterAppend(hash string, position, size int64) {
	s.indexKeysWritten++
	key := hash
	s.lastWrittenKey = key
	s.blockIndex[key] = NewBlockMetadata(position, size)
	if s.indexKeysWritten == s.indexInterval {
		s.blockIndexes = append(s.blockIndexes, s.blockIndex)
		s.initBlockIndex(config.HashingStrategy)
		s.indexKeysWritten = 0
	}
}

func (s *SSTable) append(key, hash string, buf []byte) {
	currentPos := s.beforeAppend(hash)
	str := hash + ":" + key
	b4 := make([]byte, 4)
	binary.BigEndian.PutUint32(b4, uint32(len(str)))
	// write string length
	s.dataWriter.Write(b4)
	// write string bytes
	s.dataWriter.WriteString(str)
	binary.BigEndian.PutUint32(b4, uint32(len(buf)))
	// write byte slice lengh
	s.dataWriter.Write(b4)
	s.dataWriter.Write(buf)
	s.afterAppend(hash, currentPos, int64(len(buf)))
}

func (s *SSTable) closeBF(bf *utils.BloomFilter) {
	// any remnants in the blockIndex should be added to the dump
	s.blockIndexes = append(s.blockIndexes, s.blockIndex)
	s.dumpBlockIndexes()
	// serialize the bloom filter
	buf := bf.ToByteArray()
	s.closeByte(buf, len(buf))
}

func (s *SSTable) dumpBlockIndexes() {
	position, err := s.dataWriter.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	s.dataWriter.Seek(position, 0)
	s.firstBlockPosition = position
	for _, block := range s.blockIndexes {
		s.dumpBlockIndex(block)
	}
}

// ByKey ...
type ByKey []string

// Len ...
func (p ByKey) Len() int {
	return len(p)
}

// Less ...
func (p ByKey) Less(i, j int) bool {
	return p[i] < p[j]
}

// Swap ...
func (p ByKey) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (s *SSTable) dumpBlockIndex(blockIndex map[string]*BlockMetadata) {
	if len(blockIndex) == 0 {
		return
	}
	// record the position where we start sriting the block index.
	// this will be used as the position of the lastWrittenKey in
	// the block in the index file.
	position, err := s.dataWriter.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	keys := make([]string, 0)
	for key := range blockIndex {
		keys = append(keys, key)
	}
	// String Sorted Table !
	sort.Sort(ByKey(keys))
	buf := make([]byte, 0)
	// write number of keys in this block
	b4 := make([]byte, 0)
	binary.BigEndian.PutUint32(b4, uint32(len(keys)))
	buf = append(buf, b4...)
	// write key info
	b8 := make([]byte, 8)
	for _, key := range keys {
		// write key string length
		binary.BigEndian.PutUint32(b4, uint32(len(key)))
		buf = append(buf, b4...)
		// write key string bytes
		buf = append(buf, []byte(key)...)
		// write position of the key as a relative offset
		blockMetadata := blockIndex[key]
		binary.BigEndian.PutUint64(b8, uint64(position-blockMetadata.position))
		buf = append(buf, b8...)
		// write block metadata size
		binary.BigEndian.PutUint64(b8, uint64(blockMetadata.size))
		buf = append(buf, b8...)
	}
	// write out the block index
	writeKV(s.dataWriter, SSTBlkIdxKey, buf)
	// load this index into the inmemory index map
	keyPositionInfos, ok := SSTIndexMetadataMap[s.dataFileName]
	if !ok {
		keyPositionInfos = make([]*KeyPositionInfo, 0)
		SSTIndexMetadataMap[s.dataFileName] = keyPositionInfos
	}
	keyPositionInfos = append(keyPositionInfos, NewKeyPositionInfo(keys[0], position))
}

func writeKV(file *os.File, key string, buf []byte) {
	length := len(buf)
	// size allocated: int32 + key length + int32(data size) + data byte size
	// write key length int32
	b4 := make([]byte, 4)
	binary.BigEndian.PutUint32(b4, uint32(len(key)))
	file.Write(b4)
	// write key bytes
	file.Write([]byte(key))
	// write data length
	binary.BigEndian.PutUint32(b4, uint32(length))
	file.Write(b4)
	// write data bytes
	file.Write(buf)
	// flush writes
	file.Sync()
}

func writeFooter(file *os.File, footer []byte, size int) {
	// size if int32(marker length) + marker data + int32(data size) + data bytes
	b4 := make([]byte, 4)
	binary.BigEndian.PutUint32(b4, uint32(len(bfMarker)))
	// write marker size
	file.Write(b4)
	// write marker bytes
	file.Write([]byte(bfMarker))
	// write footer size
	binary.BigEndian.PutUint32(b4, uint32(size))
	file.Write(b4)
	// write footer bytes
	file.Write(footer)
}

func (s *SSTable) closeByte(footer []byte, size int) {
	// write the bloom filter for this SSTable
	// then write three int64:
	//  1. version
	//  2. a pointer to the last written block index
	//  3. position of the bloom filter
	if s.dataWriter == nil {
		return
	}
	bloomFilterPosition, err := s.dataWriter.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	s.dataWriter.Seek(bloomFilterPosition, 0)
	// write footer
	writeFooter(s.dataWriter, footer, size)
	// write version field
	b8 := make([]byte, 8)
	binary.BigEndian.PutUint64(b8, uint64(SSTVersion))
	s.dataWriter.Write(b8)
	// write relative position of the first block index from
	// current position
	currentPos := getCurrentPos(s.dataWriter)
	blockPosition := currentPos - s.firstBlockPosition
	binary.BigEndian.PutUint64(b8, uint64(blockPosition))
	s.dataWriter.Write(b8)
	// write the position of the bloom filter
	bloomFilterRelativePosition := getCurrentPos(s.dataWriter) - bloomFilterPosition
	binary.BigEndian.PutUint64(b8, uint64(bloomFilterRelativePosition))
	s.dataWriter.Write(b8)
	// flush to disk
	s.dataWriter.Sync()
}

func getCurrentPos(file *os.File) int64 {
	res, err := file.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	file.Seek(res, 0)
	return res
}

// BlockMetadata ...
type BlockMetadata struct {
	position int64
	size     int64
}

// NewBlockMetadata ...
func NewBlockMetadata(position, size int64) *BlockMetadata {
	b := &BlockMetadata{}
	b.position = position
	b.size = size
	return b
}
