// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package utils

import (
	"encoding/binary"
	"log"
	"os"
)

func readInt32(f *os.File) int32 {
	b4 := make([]byte, 4)
	n, err := f.Read(b4)
	if err != nil {
		log.Fatal(err)
	}
	if n != len(b4) {
		log.Fatalf("insufficient read: %v, expected: %v\n", n, len(b4))
	}
	return int32(binary.BigEndian.Uint32(b4))
}

func readInt32B(buf []byte) int32 {
	return int32(binary.BigEndian.Uint32(buf))
}

func writeInt32(f *os.File, num int32) (n int, err error) {
	b4 := make([]byte, 4)
	binary.BigEndian.PutUint32(b4, uint32(num))
	return f.Write(b4)
}

func writeInt32B(buf []byte, num int32) (n int, err error) {
	b4 := make([]byte, 4)
	binary.BigEndian.PutUint32(b4, uint32(num))
	buf = append(buf, b4...)
	return 4, nil
}

func bytesLeft(file *os.File) int64 {
	curPos, err := file.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	file.Seek(curPos, 0)
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	totalBytes := fileInfo.Size()
	return totalBytes - curPos
}

func restBytes(file *os.File) []byte {
	n := bytesLeft(file)
	b := make([]byte, n)
	file.Read(b)
	return b
}
