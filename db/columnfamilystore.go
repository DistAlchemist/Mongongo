// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

/*
Package db implements ...
*/
package db

import (
	"container/heap"
	"encoding/binary"
	"io"
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
	"github.com/DistAlchemist/Mongongo/dht"
	"github.com/DistAlchemist/Mongongo/utils"
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

func (c *ColumnFamilyStore) stageOrderedCompaction(files []string) map[int][]string {
	// stage the compactions, compact similar size files.
	// this function figures out the files close enough by
	// size and if they are greater than the threshold then
	// compact
	// sort the files based on the generation ID
	sort.Sort(ByFileName(files))
	buckets := make(map[int][]string)
	maxBuckets := 1000
	averages := make([]int64, maxBuckets)
	min := int64(50 * 1024 * 1024)
	i := 0
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		fileInfo, err := f.Stat()
		if err != nil {
			log.Fatal(err)
		}
		size := fileInfo.Size()
		if (size > averages[i]/2 && size < 3*averages[i]/2) ||
			(size < min && averages[i] < min) {
			averages[i] = (averages[i] + size) / 2
			fileList, ok := buckets[i]
			if !ok {
				fileList = make([]string, 0)
				buckets[i] = fileList
			}
			fileList = append(fileList, file)
		} else {
			if i >= maxBuckets {
				break
			}
			i++
			fileList := make([]string, 0)
			buckets[i] = fileList
			fileList = append(fileList, file)
			averages[i] = size
		}
	}
	return buckets
}

// ByFileName ...
type ByFileName []string

// Len ...
func (p ByFileName) Len() int {
	return len(p)
}

// Less ...
func (p ByFileName) Less(i, j int) bool {
	return getIndexFromFileName(p[i]) < getIndexFromFileName(p[j])
}

// Swap ...
func (p ByFileName) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func getIndexFromFileName(filename string) int {
	// filename is of form <table>-<column family>-<index>-Data.db
	tokens := strings.Split(filename, "-")
	res, err := strconv.Atoi(tokens[len(tokens)-2])
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func getExpectedCompactedFileSize(files []string) int64 {
	// TODO
	return 0
}

func getMaxSizeFile(files []string) string {
	// TODO
	return ""
}

func removeFromList(files []string, file string) {
	var i int
	var f string
	for i, f = range files {
		if f == file {
			break
		}
	}
	files = append(files[:i], files[i+1:]...)
}

// FPQ is a priority queue of FileStruct
type FPQ []*FileStruct

// Len ...
func (pq FPQ) Len() int {
	return len(pq)
}

// Less ...
func (pq FPQ) Less(i, j int) bool {
	switch config.HashingStrategy {
	case config.Ophf:
		return pq[i].key < pq[j].key
	default:
		lhs := strings.Split(pq[i].key, ":")[0]
		rhs := strings.Split(pq[j].key, ":")[0]
		return lhs < rhs
	}
}

// Swap ...
func (pq FPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push ...
func (pq *FPQ) Push(x interface{}) {
	item := x.(*FileStruct)
	*pq = append(*pq, item)
}

// Pop ...
func (pq *FPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[:n-1]
	return item
}

func (c *ColumnFamilyStore) initPriorityQueue(files []string, ranges []*dht.Range, minBufferSize int) *FPQ {
	pq := &FPQ{}
	if len(files) > 1 || (ranges != nil && len(files) > 0) {
		bufferSize := c.compactionMemoryThres / len(files)
		if minBufferSize < bufferSize {
			bufferSize = minBufferSize
		}
		for _, file := range files {
			fs := &FileStruct{}
			fs.key = ""
			reader, err := os.Open(file)
			fs.reader = reader
			fs.buf = make([]byte, 0)
			if err != nil {
				log.Fatal(err)
			}
			fs = getNextKey(fs)
			if fs == nil {
				continue
			}
			heap.Push(pq, fs)
		}
	}
	return pq
}

func getNextKey(fs *FileStruct) *FileStruct {
	// Read the next key from the data file, this function will
	// skip then block index and read the next available key into
	// the filestruct that is passed. If it cannot read or a end
	// of file is reached it will return nil.
	_, key, ok := readKV(fs.reader, fs.buf)
	if !ok {
		fs.reader.Close()
		return nil
	}
	fs.key = key
	if fs.key == SSTBlkIdxKey {
		fs.reader.Close()
		return nil
	}
	return fs
}

func readKV(file *os.File, buf []byte) (int, string, bool) {
	// read key length
	b4 := make([]byte, 4)
	n, err := file.Read(b4)
	if err != nil {
		log.Print(err)
		return 0, "", false
	}
	if n < 4 {
		return 0, "", false
	}
	buf = append(buf, b4...)
	keySize := int(binary.BigEndian.Uint32(b4))
	// read key bytes as key string
	bs := make([]byte, keySize)
	file.Read(bs)
	key := string(bs)
	buf = append(buf, bs...)
	// read value length
	file.Read(b4)
	buf = append(buf, b4...)
	valueSize := int(binary.BigEndian.Uint32(b4))
	// read value bytes
	bv := make([]byte, valueSize)
	file.Read(bv)
	buf = append(buf, bv...)
	return 4 + 4 + keySize + valueSize, key, true
}

func (c *ColumnFamilyStore) getTmpFileName(files []string) string {
	// TODO
	return ""
}

func getApproximateKeyCount(files []string) int {
	// TODO
	return 0
}

func (c *ColumnFamilyStore) doFileCompaction(files []string, minBufferSize int) {
	// TODO
	// newfile := ""
	// startTime := time.Now().UnixNano() / int64(time.Millisecond)
	// totalBytesRead := int64(0)
	// totalByteWritten := int64(0)
	// totalkeysRead := int64(0)
	// totalkeysWritten := int64(0)
	// // calculate the expected compacted filesize
	// expectedCompactedFileSize := getExpectedCompactedFileSize(files)
	// compactionFileLocation := config.GetCompactionFileLocation(expectedCompactedFileSize)
	// // if the compaction file path is empty, that
	// // means we have no space left for this compaction
	// if compactionFileLocation == "" {
	// 	maxFile := getMaxSizeFile(files)
	// 	removeFromList(files, maxFile)
	// 	c.doFileCompaction(files, minBufferSize)
	// 	return
	// }
	// pq := c.initPriorityQueue(files, nil, minBufferSize)
	// if pq.Len() > 0 {
	// 	mergedFileName := c.getTmpFileName(files)
	// 	lastkey := ""
	// 	lfs := make([]*FileStruct, 0)
	// 	expectedBloomFilterSize := getApproximateKeyCount(files)
	// 	if expectedBloomFilterSize <= 0 {
	// 		expectedBloomFilterSize = SSTIndexInterval
	// 	}
	// 	log.Printf("Expeected bloom filter size: %v\n", expectedBloomFilterSize)
	// 	// create the bloom filter for the compacted file
	// 	compactedBloomFilter := utils.NewBloomFilter(expectedBloomFilterSize, 15)
	// 	columnFamilies := make([]*ColumnFamily, 0)
	// 	for pq.Len() > 0 || len(lfs) > 0 {
	// 		var fs *FileStruct
	// 		if pq.Len() > 0 {
	// 			fs = pq.Pop().(*FileStruct)
	// 		}
	// 		if fs != nil && (lastkey == "" || lastkey == fs.key) {
	// 			// The keys are the same so we need to add this to
	// 			// the lfs list
	// 			lastkey = fs.key
	// 			lfs = append(lfs, fs)
	// 		} else {
	// 			sort.Sort(ByName(lfs))
	// 			var columnFamily *ColumnFamily
	// 			if len(lfs) > 1 {
	// 				for _, filestruct := range lfs {
	// 					// read the length although we don't need it
	// 					r := bytes.NewReader(filestruct.buf)
	// 					readInt(r)
	// 					// TODO
	// 				}
	// 			}
	// 		}
	// 	}
	// }

}

func readInt(r io.Reader) int {
	b4 := make([]byte, 4)
	r.Read(b4)
	res := binary.BigEndian.Uint32(b4)
	return int(res)
}

// ByName ...
type ByName []*FileStruct

// Len ...
func (p ByName) Len() int {
	return len(p)
}

// Less ...
func (p ByName) Less(i, j int) bool {
	return p[i].reader.Name() < p[j].reader.Name()
}

// Swap ...
func (p ByName) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (c *ColumnFamilyStore) doCompaction() {
	// break the files into buckets and then compact
	c.rwmu.Lock()
	c.isCompacting = true
	c.rwmu.Unlock()
	files := make([]string, 0)
	for file := range c.ssTables {
		files = append(files, file)
	}
	buckets := c.stageOrderedCompaction(files)
	for _, fileList := range buckets {
		sort.Sort(ByFileName(fileList))
		if len(fileList) >= c.threshold {
			files = make([]string, 0)
			count := 0
			for _, file := range fileList {
				files = append(files, file)
				count++
				if count == c.threshold {
					break
				}
			}
			// for each becket if it has crossed the threshold
			// do the compaction. Incase of range compaction,
			// merge the counting bloom filters also.
			if count == c.threshold {
				c.doFileCompaction(files, c.bufSize)
			}
		}
	}
	c.rwmu.Lock()
	c.isCompacting = false
	c.rwmu.Unlock()
}

func (c *ColumnFamilyStore) apply(key string, columnFamily *ColumnFamily, cLogCtx *CommitLogContext) {
	// c.memtable.mu.Lock()
	// defer c.memtable.mu.Unlock()
	c.memtable.put(key, columnFamily, cLogCtx)
}

func (c *ColumnFamilyStore) switchMemtable(key string, columnFamily *ColumnFamily, cLogCtx *CommitLogContext) {
	// Used on start up when we are recovering from logs
	c.memtable.mu.Lock()
	c.memtable = NewMemtable(c.tableName, c.columnFamilyName)
	c.memtable.mu.Unlock()
	if key != c.memtable.flushKey {
		c.memtable.mu.Lock()
		c.memtable.put(key, columnFamily, cLogCtx)
		c.memtable.mu.Unlock()
	}
}

func (c *ColumnFamilyStore) getNextFileName() string {
	// increment twice to generate non-consecutive numbers
	atomic.AddInt32(&c.fileIdxGenerator, 1)
	atomic.AddInt32(&c.fileIdxGenerator, 1)
	name := c.tableName + "-" + c.columnFamilyName + "-" +
		strconv.Itoa(int(c.fileIdxGenerator))
	return name
}

func (c *ColumnFamilyStore) onMemtableFlush(cLogCtx *CommitLogContext) {
	// Called when the memtable is frozen and ready to be flushed
	// to disk. This method informs the commitlog that a particular
	// columnFamily is being flushed to disk.
	if cLogCtx.isValidContext() {
		openCommitLog(c.tableName).onMemtableFlush(c.columnFamilyName, cLogCtx)
	}
}

func (c *ColumnFamilyStore) storeLocation(filename string, bf *utils.BloomFilter) {
	// Called after the memtable flushes its inmemory data.
	// This information is cached in the ColumnFamilyStore.
	// This is useful for reads because the ColumnFamilyStore first
	// looks in the inmemory store and then into the disk to find
	// the key. If invoked during recoveryMode the onMemtableFlush()
	// need not be invoked.
	doCompaction := false
	ssTableSize := 0
	c.rwmu.Lock()
	c.ssTables[filename] = true
	storeBloomFilter(filename, bf)
	ssTableSize = len(c.ssTables)
	c.rwmu.Unlock()
	if ssTableSize >= c.threshold && !c.isCompacting {
		doCompaction = true
	}
	if c.isCompacting {
		if ssTableSize%c.threshold == 0 {
			doCompaction = true
		}
	}
	if doCompaction {
		log.Printf("Submitting for compaction...")
		go c.doCompaction()
	}

}

func storeBloomFilter(filename string, bf *utils.BloomFilter) {
	SSTbfs[filename] = bf
}
