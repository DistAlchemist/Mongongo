// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import "sync"

// BinaryMemtable is the binary version of memtable
type BinaryMemtable struct {
	threshold      int
	currentSize    int32
	tableName      string
	cfName         string
	isFrozen       bool
	columnFamilies map[string][]byte
	mu             sync.Mutex
	cond           *sync.Cond
}

// NewBinaryMemtable initializes a BinaryMemtable
func NewBinaryMemtable(table, cfName string) *BinaryMemtable {
	b := &BinaryMemtable{}
	b.threshold = 512 * 1024 * 1024
	b.currentSize = 0
	b.tableName = table
	b.cfName = cfName
	b.isFrozen = false
	b.columnFamilies = make(map[string][]byte)
	b.cond = sync.NewCond(&b.mu)
	return b
}
