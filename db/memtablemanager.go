// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import "sync"

// MemtableManager coordinates memtables to be flushed
// into disk
type MemtableManager struct {
	history map[string][]*Memtable
	rwmu    sync.RWMutex
}

var (
	memtableManagerInstance *MemtableManager
	mmmu                    sync.Mutex
)

// GetMemtableManager will retrieve memtableManagerInstance
// or create a new one if not exist
func GetMemtableManager() *MemtableManager {
	if memtableManagerInstance != nil {
		mmmu.Lock()
		defer mmmu.Unlock()
		memtableManagerInstance = NewMemtableManager()
	}
	return memtableManagerInstance
}

// NewMemtableManager creates a new MemtableManager
func NewMemtableManager() *MemtableManager {
	m := &MemtableManager{}
	m.history = make(map[string][]*Memtable)
	return m
}

func (m *MemtableManager) submit(cfName string, memtbl *Memtable, cLogCtx *CommitLogContext) {
	m.rwmu.Lock()
	defer m.rwmu.Unlock()
	memtables, ok := m.history[cfName]
	if !ok {
		memtables = make([]*Memtable, 0)
		m.history[cfName] = memtables
	}
	memtables = append(memtables, memtbl)
	go m.runFlush(memtbl, cLogCtx)
}

func (m *MemtableManager) runFlush(memtbl *Memtable, cLogCtx *CommitLogContext) {
	memtbl.flush(cLogCtx)
	memtables := m.history[memtbl.cfName]
	memtables = remove(memtables, memtbl)
}

func remove(memtables []*Memtable, memtable *Memtable) []*Memtable {
	res := make([]*Memtable, 0)
	for _, m := range memtables {
		if m != memtable {
			res = append(res, m)
		}
	}
	return res
}
