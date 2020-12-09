// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/DistAlchemist/Mongongo/utils"
)

var (
	// openedFiles for SSTableReader
	openedFiles = NewFileSSTableMap()
	srmu        sync.Mutex
)

// SSTableReader ...
type SSTableReader struct {
	*SSTable
}

// filename is the full path name with dir
func openSSTableReader(dataFilename string) *SSTableReader {
	sstable, ok := openedFiles.get(dataFilename)
	if !ok {
		sstable = NewSSTableReader(dataFilename)
		start := time.Now().UnixNano() / int64(time.Millisecond.Milliseconds())
		sstable.loadIndexFile()
		sstable.loadBloomFilter()
		log.Printf("index load time for %v: %v ms.",
			dataFilename, time.Now().UnixNano()/int64(time.Millisecond)-start)
		openedFiles.put(dataFilename, sstable)
	}
	return sstable
}

func getSSTableReader(dataFileName string) *SSTableReader {
	srmu.Lock()
	defer srmu.Unlock()
	sstable, _ := openedFiles.get(dataFileName)
	return sstable
}

// NewSSTableReader ...
func NewSSTableReader(filename string) *SSTableReader {
	s := &SSTableReader{}
	s.SSTable = NewSSTable(filename)
	return s
}

// NewSSTableReaderI ...
func NewSSTableReaderI(filename string, indexPositions []*KeyPositionInfo, bf *utils.BloomFilter) *SSTableReader {
	s := &SSTableReader{}
	s.SSTable = NewSSTable(filename)
	s.bf = bf
	srmu.Lock()
	defer srmu.Unlock()
	openedFiles.put(filename, s)
	return s
}

func (s *SSTableReader) loadIndexFile() {
	/** Index file structure:
	 * decoreatedKey (int32+string)
	 * index (int64)
	 * (repeat above two)
	 * */
	s.indexPositions = make([]*KeyPositionInfo, 0)
	input, err := os.Open(s.indexFilename(s.dataFileName))
	if err != nil {
		log.Fatal(err)
	}
	fileInfo, err := input.Stat()
	if err != nil {
		log.Fatal(err)
	}
	// length in bytes
	indexSize := fileInfo.Size()
	i := 0
	for {
		indexPosition := getCurrentPos(input)
		if indexPosition == indexSize {
			break
		}
		decoratedKey, _ := readString(input)
		readInt64(input)
		if i%s.indexInterval == 0 {
			s.indexPositions = append(s.indexPositions,
				NewKeyPositionInfo(decoratedKey, indexPosition))
		}
	}
}

func (s *SSTableReader) loadBloomFilter() {
	stream, err := os.Open(s.filterFilename(s.dataFileName))
	if err != nil {
		log.Fatal(err)
	}
	s.bf = utils.BFSerializer.Deserialize(stream)
}

func (s *SSTableReader) getFileStruct() *FileStruct {
	return NewFileStruct(s)
}

func (s *SSTableReader) getTableName() string {
	return s.parseTableName(s.dataFileName)
}

func (s *SSTableReader) makeColumnFamily() *ColumnFamily {
	return createColumnFamily(s.getTableName(), s.getColumnFamilyName())
}

func (s *SSTableReader) getIndexPositions() []*KeyPositionInfo {
	return s.indexPositions
}

func (s *SSTableReader) delete() {
	os.Remove(s.dataFileName)
	os.Remove(s.indexFilename(s.dataFileName))
	os.Remove(s.filterFilename(s.dataFileName))
	srmu.Lock()
	defer srmu.Unlock()
	openedFiles.remove(s.dataFileName)
}

// FileSSTableMap ...
type FileSSTableMap struct {
	m map[string]*SSTableReader
}

// NewFileSSTableMap ...
func NewFileSSTableMap() *FileSSTableMap {
	f := &FileSSTableMap{}
	f.m = make(map[string]*SSTableReader)
	return f
}

// Caution: the key is always full filename with dir
func (f *FileSSTableMap) get(filename string) (*SSTableReader, bool) {
	res, ok := f.m[filename]
	return res, ok
}

// Caution: the key is always full filename with dir
func (f *FileSSTableMap) put(filename string, value *SSTableReader) {
	f.m[filename] = value
}

func (f *FileSSTableMap) values() []*SSTableReader {
	res := make([]*SSTableReader, len(f.m))
	for _, value := range f.m {
		res = append(res, value)
	}
	return res
}

func (f *FileSSTableMap) clear() {
	f.m = make(map[string]*SSTableReader)
}

func (f *FileSSTableMap) remove(filename string) {
	delete(f.m, filename)
}
