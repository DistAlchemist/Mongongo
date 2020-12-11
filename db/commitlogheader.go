// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import (
	"log"

	"github.com/willf/bitset"
)

// CommitLogHeader represents the header of commit log
type CommitLogHeader struct {
	// header        []byte
	// position      []int
	dirty         *bitset.BitSet
	lastFlushedAt []int
}

// NewCommitLogHeader creates a new commit log header
// size is the number of column families
func NewCommitLogHeader(size int) *CommitLogHeader {
	c := &CommitLogHeader{}
	// c.header = make([]byte, size)
	// c.position = make([]int, size)
	c.dirty = bitset.New(uint(size))
	c.lastFlushedAt = make([]int, size)
	return c
}

// NewCommitLogHeaderD used in deserializing
func NewCommitLogHeaderD(dirty *bitset.BitSet, lastFlushedAt []int) *CommitLogHeader {
	c := &CommitLogHeader{}
	c.dirty = dirty
	c.lastFlushedAt = lastFlushedAt
	return c
}

// NewCommitLogHeaderC used in copy
func NewCommitLogHeaderC(clHeader *CommitLogHeader) *CommitLogHeader {
	c := &CommitLogHeader{}
	c.dirty = clHeader.dirty.Clone()
	c.lastFlushedAt = make([]int, len(clHeader.lastFlushedAt))
	for _, x := range clHeader.lastFlushedAt {
		c.lastFlushedAt = append(c.lastFlushedAt, x)
	}
	return c
}

func (c *CommitLogHeader) getPosition(index int) int {
	return c.lastFlushedAt[index]
}

func (c *CommitLogHeader) turnOn(index int, position int64) {
	c.dirty.Set(uint(index))
	c.lastFlushedAt[index] = int(position)
}

func (c *CommitLogHeader) turnOff(index int) {
	c.dirty.Clear(uint(index))
	c.lastFlushedAt[index] = 0
}

func (c *CommitLogHeader) isDirty(index int) bool {
	return c.dirty.Test(uint(index))
}

func (c *CommitLogHeader) isSafeToDelete() bool {
	return c.dirty.Any()
}

func (c *CommitLogHeader) clear() {
	c.dirty.ClearAll()
	c.lastFlushedAt = make([]int, 0)
}

func (c *CommitLogHeader) toByteArray() []byte {
	bos := make([]byte, 0)
	clhSerialize(c, bos)
	return bos
}

func clhSerialize(clHeader *CommitLogHeader, dos []byte) {
	dbytes, err := clHeader.dirty.MarshalBinary()
	if err != nil {
		log.Fatal(err)
	}
	writeIntB(dos, len(dbytes))
	dos = append(dos, dbytes...)
	writeIntB(dos, len(clHeader.lastFlushedAt))
	for _, position := range clHeader.lastFlushedAt {
		writeIntB(dos, position)
	}
}

// func clhDeserialize(dis []byte) *CommitLogHeader {
// 	dirtyLen := readInt(dis)
// }
