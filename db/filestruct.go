// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import (
	"log"
	"os"
)

// FileStruct ...
type FileStruct struct {
	key       string
	row       *IteratingRow
	exhausted bool
	file      *os.File
	sstable   *SSTableReader
}

// NewFileStruct ...
func NewFileStruct(s *SSTableReader) *FileStruct {
	f := &FileStruct{}
	f.exhausted = false
	var err error
	f.file, err = os.Open(s.dataFileName)
	if err != nil {
		log.Fatal(err)
	}
	f.sstable = s
	return f
}

func (f *FileStruct) advance(materialize bool) {
	// Read the next key from the data file.
	if f.exhausted {
		log.Fatal("index out of bounds!")
	}
	if getCurrentPos(f.file) == getFileSize(f.file) {
		f.file.Close()
		f.exhausted = true
		return
	}
	f.row = NewIteratingRow(f.file, f.sstable)
	if materialize {
		for f.row.hasNext() {
			column := f.row.next()
			f.row.getEmptyColumnFamily().addColumn(column)
		}
	} else {
		f.row.skipRemaining()
	}
}

func (f *FileStruct) getFileName() string {
	return f.file.Name()
}

func (f *FileStruct) getColumnFamily() *ColumnFamily {
	return f.row.getEmptyColumnFamily()
}

func (f *FileStruct) isExhausted() bool {
	return f.exhausted
}

func getFileSizeFromName(name string) int64 {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	return getFileSize(file)
}

func getFileSize(file *os.File) int64 {
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	return fileInfo.Size()
}
