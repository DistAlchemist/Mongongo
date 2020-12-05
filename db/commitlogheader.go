// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import "encoding/binary"

// CommitLogHeader represents the header of commit log
type CommitLogHeader struct {
	header   []byte
	position []int
}

// NewCommitLogHeader creates a new commit log header
// size is the number of column families
func NewCommitLogHeader(size int) *CommitLogHeader {
	c := &CommitLogHeader{}
	c.header = make([]byte, size)
	c.position = make([]int, size)
	return c
}

func (c *CommitLogHeader) copy() *CommitLogHeader {
	n := &CommitLogHeader{}
	n.header = make([]byte, 0)
	n.position = make([]int, 0)
	for _, h := range c.header {
		n.header = append(n.header, h)
	}
	for _, p := range c.position {
		n.position = append(n.position, p)
	}
	return n
}

func (c *CommitLogHeader) zeroPositions() {
	size := len(c.position)
	c.position = make([]int, size)
}

func (c *CommitLogHeader) turnOn(idx int, position int64) {
	c.header[idx] = byte(1)
	c.position[idx] = int(position)
}

func (c *CommitLogHeader) turnOff(idx int) {
	c.header[idx] = byte(0)
	c.position[idx] = 0
}

func (c *CommitLogHeader) and(commitLogHeader *CommitLogHeader) {
	clh2 := commitLogHeader.header
	for i := 0; i < len(c.header); i++ {
		c.header[i] = c.header[i] & clh2[i]
	}
}

func (c *CommitLogHeader) isSafeToDelete() bool {
	for _, b := range c.header {
		if b == 1 {
			return false
		}
	}
	return true
}

func (c *CommitLogHeader) toByteArray() []byte {
	// of the format:
	//  headerByteLength: uint32
	//  header          : []byte
	//  position        : []uint32
	buf := make([]byte, 0)
	b4 := make([]byte, 4)
	binary.BigEndian.PutUint32(b4, uint32(len(c.header)))
	// put headByteLength uint32
	buf = append(buf, b4...)
	// put header
	buf = append(buf, c.header...)
	// put position
	for _, p := range c.position {
		binary.BigEndian.PutUint32(b4, uint32(p))
		buf = append(buf, b4...)
	}
	return buf
}
