// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import (
	"log"
	"os"
	"strings"

	"github.com/DistAlchemist/Mongongo/utils"
)

// SSTableWriter ...
type SSTableWriter struct {
	*SSTable
	dataFile  *os.File
	indexFile *os.File
}

// NewSSTableWriter ...
func NewSSTableWriter(filename string, keyCount int) *SSTableWriter {
	s := &SSTableWriter{}
	s.SSTable = NewSSTable(filename)
	var err error
	s.dataFile, err = os.Open(s.dataFileName)
	if err != nil {
		log.Fatal(err)
	}
	s.indexFile, err = os.Open(s.indexFilename(s.dataFileName))
	if err != nil {
		log.Fatal(err)
	}
	s.bf = utils.NewBloomFilter(keyCount, 15)
	return s
}

func compare(s1, s2 string) bool {
	// currently I only use direct compare,
	// which corresponds to random partition strategy
	return s1 < s2
}

func (s *SSTableWriter) beforeAppend(decoratedKey string) int64 {
	if decoratedKey == "" {
		log.Fatal("key must not be empty")
	}
	if s.lastWrittenKey != "" && compare(s.lastWrittenKey, decoratedKey) == false {
		log.Printf("Last written key: %v\n", s.lastWrittenKey)
		log.Printf("Current key: %v\n", decoratedKey)
		log.Printf("Writing into file %v\n", s.dataFileName)
		log.Fatal("keys must be written in ascending order")
	}
	if s.lastWrittenKey == "" {
		return 0
	}
	return getCurrentPos(s.dataFile)
}

func (s *SSTableWriter) afterAppend(decoratedKey string, position int64) {
	s.bf.Fill(decoratedKey)
	s.lastWrittenKey = decoratedKey
	indexPosition := getCurrentPos(s.indexFile)
	// write index file
	writeString(s.indexFile, decoratedKey)
	writeInt64(s.indexFile, position)
	if s.indexKeysWritten%SSTIndexInterval != 0 {
		s.indexKeysWritten++
		return
	}
	s.indexKeysWritten++
	if s.indexPositions == nil {
		s.indexPositions = make([]*KeyPositionInfo, 0)
	}
	s.indexPositions = append(s.indexPositions, NewKeyPositionInfo(decoratedKey, indexPosition))
}

func (s *SSTableWriter) append(decoratedKey string, buf []byte) {
	currentPos := s.beforeAppend(decoratedKey)
	writeString(s.dataFile, decoratedKey)
	writeBytes(s.dataFile, buf)
	s.afterAppend(decoratedKey, currentPos)
}

func (s *SSTableWriter) closeAndOpenReader() *SSTableReader {
	// renames temp SSTable files to valid data, index and bloom filter files
	// bloom filter file
	fos, err := os.Open(s.filterFilename(s.dataFileName))
	if err != nil {
		log.Fatal(err)
	}
	utils.BFSerializer.Serialize(s.bf, fos)
	fos.Sync()
	fos.Close()
	// index file
	s.indexFile.Sync()
	s.indexFile.Close()
	// main data
	s.dataFile.Sync()
	s.dataFile.Close()

	s.rename(s.indexFilename(s.dataFileName))
	s.rename(s.filterFilename(s.dataFileName))
	s.dataFileName = s.rename(s.dataFileName)
	return NewSSTableReaderI(s.dataFileName, s.indexPositions, s.bf)
}

func (s *SSTableWriter) rename(tmpFilename string) string {
	filename := strings.Replace(tmpFilename, "-"+SSTableTmpFile, "", 1)
	os.Rename(tmpFilename, filename)
	return filename
}
