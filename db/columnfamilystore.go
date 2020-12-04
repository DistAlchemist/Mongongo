// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/DistAlchemist/Mongongo/config"
)

// ColumnFamilyStore provides storage specification of
// column family
type ColumnFamilyStore struct {
	threshold             int
	bufSize               int
	compactionMemoryThres int
	tableName             string
	columnFamilyName      string
	// to generate the next index for a SSTable
	fileIdxGenerator int32
	// memtables associated with this cfStore
	memtable       *Memtable
	binaryMemtable *BinaryMemtable
	// SSTable on disk for this cf
	ssTables map[string]bool
	// modification lock used for protecting reads
	// from compactions
	rwmu sync.RWMutex
	// flag indicates if a compaction is in process
	isCompacting bool
}

// NewColumnFamilyStore initializes a new ColumnFamilyStore
func NewColumnFamilyStore(table, cfName string) *ColumnFamilyStore {
	c := &ColumnFamilyStore{}
	c.threshold = 4
	c.bufSize = 128 * 1024 * 1024
	c.compactionMemoryThres = 1 << 30
	c.tableName = table
	c.columnFamilyName = cfName
	c.fileIdxGenerator = 0
	c.ssTables = make(map[string]bool)
	c.isCompacting = false
	// Get all data files associated with old Memtables for this table.
	// The names are <Table>-<CfName>-1.db, ..., <Table>-<CfName>-n.db.
	// The max is n and increment it to be used as the next index.
	indices := make([]int, 0)
	dataFileDirs := config.DataFileDirs
	for _, dir := range dataFileDirs {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			log.Fatal(err)
		}
		for _, fileInfo := range files {
			filename := fileInfo.Name()
			tblCfName := getTableAndColumnFamilyName(filename)
			if tblCfName[0] == table && tblCfName[0] == cfName {
				idx := getIdxFromFileName(filename)
				indices = append(indices, idx)
			}
		}
	}
	sort.Ints(indices)
	sz := len(indices)
	value := 0
	if sz > 0 {
		value = indices[sz-1]
	}
	atomic.StoreInt32(&c.fileIdxGenerator, int32(value))
	c.memtable = NewMemtable(table, cfName)
	c.binaryMemtable = NewBinaryMemtable(table, cfName)
	return c
}

func getTableAndColumnFamilyName(filename string) []string {
	// filename is of format:
	//   <table>-<column family>-<index>-Data.db
	values := strings.Split(filename, "-")
	return values[:2] // tableName and cfName
}

func getIdxFromFileName(filename string) int {
	// filename if of format:
	//   <table>-<column family>-<index>-Data.db
	values := strings.Split(filename, "-")
	if len(values) < 3 {
		log.Fatal("Invalid filename")
	}
	res, err := strconv.Atoi(values[2])
	if err != nil {
		log.Fatal(err)
	}
	return res
}

// fileInfoList encapsulates os.FileInfo for comparison needs.
type fileInfoList []os.FileInfo

func (f fileInfoList) Len() int {
	return len(f)
}

func (f fileInfoList) Less(i, j int) bool {
	return f[i].ModTime().UnixNano() > f[j].ModTime().UnixNano()
}

func (f fileInfoList) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (c *ColumnFamilyStore) onStart() {
	// do major compaction
	ssTables := make([]os.FileInfo, 0)
	dataFileDirs := config.DataFileDirs
	filenames := make([]string, 0)
	for _, dir := range dataFileDirs {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			log.Fatal(err)
		}
		for _, fileInfo := range files {
			filename := fileInfo.Name()
			if strings.Contains(filename, c.columnFamilyName) &&
				(fileInfo.Size() == 0 || strings.Contains(filename, SSTableTmpFile)) {
				err := os.Remove(path.Join(dir, filename))
				if err != nil {
					log.Print(err)
				}
				continue
			}
			tblCfName := getTableAndColumnFamilyName(filename)
			if tblCfName[0] == c.tableName && tblCfName[1] == c.columnFamilyName &&
				strings.Contains(filename, "-Data.db") {
				ssTables = append(ssTables, fileInfo)
				filenames = append(filenames, path.Join(dir, fileInfo.Name()))
			}
		}
	}
	// sort.Sort(fileInfoList(ssTables))
	// filenames := make([]string, len(ssTables))
	// for _, ssTable := range ssTables {
	// 	filenames = append(filenames, ssTable.Name())
	// }
	// filename of the type:
	//  var/storage/data/<tableName>-<columnFamilyName>-<index>-Data.db
	for _, filename := range filenames {
		c.ssTables[filename] = true
	}
	onSSTableStart(filenames)
	log.Println("Submitting a major compaction task")
	go c.doCompaction()
	// TODO should also submit periodic minor compaction
}

func (c *ColumnFamilyStore) doCompaction() {
	// TODO
}

func (c *ColumnFamilyStore) apply(key string, columnFamily *ColumnFamily, cLogCtx *CommitLogContext) {
	// c.memtable.mu.Lock()
	// defer c.memtable.mu.Unlock()
	c.memtable.put(key, columnFamily, cLogCtx)
}

func (c *ColumnFamilyStore) switchMemtable(key string, columnFamily *ColumnFamily, cLogCtx *CommitLogContext) {
	// TODO
}
