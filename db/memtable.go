// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import (
	"sync"
	"time"

	"github.com/DistAlchemist/Mongongo/config"
)

// Memtable specifies memtable
type Memtable struct {
	flushKey         string
	threshold        int32
	thresholdCnt     int32
	currentSize      int32
	currentObjectCnt int32
	// table and cf name are used to determine the
	// ColumnFamilyStore
	tableName string
	cfName    string
	// creation time of this memtable
	creationTime   int64
	isFrozen       bool
	columnFamilies map[string]ColumnFamily
	// lock and condition for notifying new clients about Memtable switches
	mu   sync.Mutex
	cond *sync.Cond
}

// NewMemtable initializes a new memtable
func NewMemtable(table, cfName string) *Memtable {
	m := &Memtable{}
	m.flushKey = "FlushKey"
	m.threshold = int32(config.MemtableSize * 1024 * 1024)
	m.thresholdCnt = int32(config.MemtableObjectCount * 1024 * 1024)
	m.currentSize = 0
	m.currentObjectCnt = 0
	m.isFrozen = false
	m.columnFamilies = make(map[string]ColumnFamily)
	m.cond = sync.NewCond(&m.mu)
	m.tableName = table
	m.cfName = cfName
	m.creationTime = time.Now().UnixNano() / int64(time.Millisecond)
	return m
}

// put data into the memtable
// flush memtable to disk when the size exceeds the threshold
func (m *Memtable) put(key string, columnFamily *ColumnFamily, cLogCtx *CommitLogContext) {
	if m.isThresholdViolated(key) {
		m.mu.Lock()
		defer m.mu.Unlock()
		cfStore := openTable(m.tableName).columnFamilyStores[m.cfName]
		if !m.isFrozen {
			m.isFrozen = true
			// submit memtable flush TODO
			cfStore.switchMemtable(key, columnFamily, cLogCtx)
		} else {
			cfStore.apply(key, columnFamily, cLogCtx)
		}
	} else {
		// submit task to put key-cf to memtable TODO
	}
}

func (m *Memtable) isThresholdViolated(key string) bool {
	bVal := false
	if m.currentSize >= m.threshold || m.currentObjectCnt >= m.thresholdCnt || bVal || key == m.flushKey {
		return true
	}
	return false
}
