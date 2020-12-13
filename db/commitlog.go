// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import (
	"encoding/binary"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/DistAlchemist/Mongongo/config"
)

// CommitLog tracks every write operation into the system.
// The aim of the commit log is to be able to successfully
// recover data that was not stored to disk via the memtable.
// Every commit log maintains a header represented by the
// abstraction CommitLogHeader. The header contains a bit
// array and an array of int64 and both the arrays are of
// size: # column families. Whenever a ColumnFamily is
// written to, for the first time its bit flag is set to
// one in the CommitLogHeader. When it is flushed to disk
// by the Memtable its corresponding bit in the header is
// set to zero. This helps track which CommitLog can be thrown
// away as a result of Memtable flushes. However if a ColumnFamily
// is flushed and again written to disk then its entry in the
// array of int64 is updated with the offset in the CommitLog
// file where it was written. This helps speed up recovery since
// we can seek to these offsets and start processing the commit
// log. Every Commit Log is rolled over everytime it reaches its
// threshold in size. Over time there could be a number of
// commit logs that would be generated. However whenever we flush
// a column family disk and update its bit flag we take this bit
// array and bitwise & it with the headers of the other commit logs
// that are older.
type CommitLog struct {
	bufSize              int
	table                string
	logFile              string
	clHeader             *CommitLogHeader
	commitHeaderStartPos int64
	forcedRollOver       bool
	logWriter            *os.File
}

var (
	clInstance  = map[string]*CommitLog{}
	clSInstance *CommitLog // stands for Single Instance
	clHeaders   = map[string]*CommitLogHeader{}
	clmu        sync.Mutex
)

// CommitLogContext represents the context of commit log
type CommitLogContext struct {
	file string
	// offset within the Commit Log where this row was added
	position int64
}

// NewCommitLogContext creates a new commitLogContext
func NewCommitLogContext(file string, position int64) *CommitLogContext {
	c := &CommitLogContext{}
	c.file = file
	c.position = position
	return c
}

func (c *CommitLogContext) isValidContext() bool {
	return c.position != -1
}

func (c *CommitLog) setNextFileName() {
	c.logFile = config.LogFileDir + string(os.PathSeparator) +
		"CommitLog-" +
		strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10) +
		".log"
}

func createCLWriter(file string) *os.File {
	f, err := os.Create(file)
	if err != nil {
		log.Print(err)
	}
	return f
}

func (c *CommitLog) writeCommitLogHeader() {
	// writes a header with all bits set to zero
	table := OpenTable(c.table)
	cfSize := table.getNumberOfColumnFamilies() // number of cf
	c.commitHeaderStartPos = 0
	// write the commit log header
	c.clHeader = NewCommitLogHeader(cfSize)
	c.writeCLH(c.clHeader.toByteArray(), false)
}

func (c *CommitLog) writeCommitLogHeaderB(bytes []byte, reset bool) {
	// record the current position
	currentPos, err := c.logWriter.Seek(c.commitHeaderStartPos, 0)
	if err != nil {
		log.Fatal(err)
	}
	currentPos += c.commitHeaderStartPos
	// write the commit log header
	_, err = c.logWriter.Write(bytes)
	if err != nil {
		log.Print(err)
	}
	if reset {
		// seek back to the old position
		c.logWriter.Seek(currentPos, 0)
	}
}

func (c *CommitLog) writeOldCommitLogHeader(oldFile string, header *CommitLogHeader) {
	logWriter := createCLWriter(oldFile)
	writeCommitLogHeader(logWriter, header.toByteArray())
	logWriter.Close()
}

func (c *CommitLog) writeCLH(bytes []byte, reset bool) {
	currentPos, err := c.logWriter.Seek(c.commitHeaderStartPos, 0)
	if err != nil {
		log.Print(err)
	}
	currentPos += c.commitHeaderStartPos
	// write the commit log header
	c.logWriter.Write(bytes)
	if reset {
		c.logWriter.Seek(currentPos, 0)
	}
}

func (c *CommitLog) getContext() *CommitLogContext {
	ctx := NewCommitLogContext(c.logFile, getCurrentPos(c.logWriter))
	return ctx
}

// NewCommitLog creates a new commit log
func NewCommitLog(table string, recoveryMode bool) *CommitLog {
	c := &CommitLog{}
	c.table = table
	c.forcedRollOver = false
	if !recoveryMode {
		c.setNextFileName()
		c.logWriter = createCLWriter(c.logFile)
		c.writeCommitLogHeader()
	}
	return c
}

// NewCommitLogE creates a new commit log
func NewCommitLogE(recoveryMode bool) *CommitLog {
	c := &CommitLog{}
	// c.table = table
	c.forcedRollOver = false
	if !recoveryMode {
		c.setNextFileName()
		c.logWriter = createCLWriter(c.logFile)
		c.writeCommitLogHeader()
	}
	return c
}

func openCommitLog(table string) *CommitLog {
	commitLog, ok := clInstance[table]
	if !ok {
		clmu.Lock()
		defer clmu.Unlock()
		commitLog = NewCommitLog(table, false)
		clInstance[table] = commitLog
	}
	return commitLog
}

func openCommitLogE() *CommitLog {
	clmu.Lock()
	defer clmu.Unlock()
	if clSInstance == nil {
		clSInstance = NewCommitLogE(false)
	}
	return clSInstance
}

func (c *CommitLog) maybeUpdateHeader(row *Row) {
	// update the header of the commit log if a
	// new column family is encountered for the
	// first time
	table := OpenTable(row.Table)
	for cfName := range row.getColumnFamilies() {
		id := table.getColumnFamilyID(cfName)
		if c.clHeader.isDirty(id) == false {
			c.clHeader.turnOn(id, getCurrentPos(c.logWriter))
			c.seekAndWriteCommitLogHeader(c.clHeader.toByteArray())
		}
	}
}

func (c *CommitLog) seekAndWriteCommitLogHeader(bytes []byte) {
	// writes header at the beginning of the file, then seeks
	// back to current position
	currentPos := getCurrentPos(c.logWriter)
	c.logWriter.Seek(0, 0)
	writeCommitLogHeader(c.logWriter, bytes)
	c.logWriter.Seek(currentPos, 0)
}

func writeCommitLogHeader(logWriter *os.File, bytes []byte) {
	writeInt64(logWriter, int64(len(bytes)))
	writeBytes(logWriter, bytes)
}

func (c *CommitLog) maybeRollLog() bool {
	if getFileSize(c.logWriter) >= config.LogRotationThres {
		// rolls the current log file over to a new one
		c.setNextFileName()
		oldLogFile := c.logWriter.Name()
		c.logWriter.Close()
		// point reader/writer to a new commit log file
		c.logWriter = createCLWriter(c.logFile)
		// squirrel away the old commit log header
		clHeaders[oldLogFile] = NewCommitLogHeaderC(c.clHeader)
		c.clHeader.clear()
		writeCommitLogHeader(c.logWriter, c.clHeader.toByteArray())
		return true
	}
	return false
}

// add the specified row to the commit log. This method will
// reset the file offset to what it is before the start of
// the operation in case of any problems. This way we can
// assume that the subsequent commit log entry will override
// the garbage left over by the previous write.
func (c *CommitLog) add(row *Row) *CommitLogContext {
	curPos := int64(-1)
	buf := make([]byte, 0)
	// serialize the row
	rowSerialize(row, buf)
	curPos = getCurrentPos(c.logWriter)
	cLogCtx := NewCommitLogContext(c.logFile, curPos)
	// update header
	c.maybeUpdateHeader(row)
	writeInt64(c.logWriter, int64(len(buf)))
	writeBytes(c.logWriter, buf)
	c.maybeRollLog()
	return cLogCtx
}

// writeString will first write string length(int32)
// and then write string in bytes
func writeString(file *os.File, s string) int {
	// write string length
	b4 := make([]byte, 4)
	binary.BigEndian.PutUint32(b4, uint32(len(s)))
	file.Write(b4)
	// write string bytes
	file.Write([]byte(s))
	// return total bytes written
	return 4 + len(s)
}
func writeStringB(file []byte, s string) int {
	// write string length
	b4 := make([]byte, 4)
	binary.BigEndian.PutUint32(b4, uint32(len(s)))
	file = append(file, b4...)
	// write string bytes
	file = append(file, []byte(s)...)
	// return total bytes written
	return 4 + len(s)
}

func writeInt(file *os.File, num int) int {
	return writeInt32(file, int32(num))
}

func writeIntB(buf []byte, num int) int {
	return writeInt32B(buf, int32(num))
}

func writeInt32(file *os.File, num int32) int {
	b4 := make([]byte, 4)
	binary.BigEndian.PutUint32(b4, uint32(num))
	file.Write(b4)
	return 4
}

func writeInt32B(buf []byte, num int32) int {
	b4 := make([]byte, 4)
	binary.BigEndian.PutUint32(b4, uint32(num))
	buf = append(buf, b4...)
	return 4
}

func writeInt64(file *os.File, num int64) int {
	b8 := make([]byte, 8)
	binary.BigEndian.PutUint64(b8, uint64(num))
	file.Write(b8)
	return 8
}

func writeInt64B(buf []byte, num int64) int {
	b8 := make([]byte, 8)
	binary.BigEndian.PutUint64(b8, uint64(num))
	buf = append(buf, b8...)
	return 8
}

func writeBool(file *os.File, b bool) int {
	if b == true {
		file.Write([]byte{1})
	} else {
		file.Write([]byte{0})
	}
	return 1
}

func writeBoolB(file []byte, b bool) int {
	if b == true {
		file = append(file, byte(1))
	} else {
		file = append(file, byte(0))
	}
	return 1
}

func writeBytes(file *os.File, b []byte) int {
	// write byte length
	b4 := make([]byte, 4)
	binary.BigEndian.PutUint32(b4, uint32(len(b)))
	file.Write(b4)
	// write bytes
	file.Write(b)
	// return total bytes written
	return 4 + len(b)
}

func writeBytesB(buf []byte, b []byte) int {
	// write byte length
	b4 := make([]byte, 4)
	binary.BigEndian.PutUint32(b4, uint32(len(b)))
	buf = append(buf, b4...)
	// write bytes
	buf = append(buf, b...)
	// return total bytes written
	return 4 + len(b)
}

// func (c *CommitLog) checkThresholdAndRollLog(fileSize int64) {
// 	if fileSize >= config.LogRotationThres || c.forcedRollOver {
// 		// rolls the current log file over to a new one
// 		c.setNextFileName()
// 		oldLogFile := c.logWriter.Name()
// 		c.logWriter.Close()
// 		// change logWriter to new log file
// 		c.logWriter = c.createWriter(c.logFile)
// 		// save old log header
// 		clHeaders[oldLogFile] = c.clHeader.copy()
// 		// zero out positions in old file log header
// 		c.clHeader.zeroPositions()
// 		c.writeCommitLogHeaderB(c.clHeader.toByteArray(), false)
// 		// Get the list of files in commit log dir if it is greater than a
// 		// certain number. Force flush all the column families that way we
// 		// ensure that a slowly populated column family is not screwing up
// 		// by accumulating the commit log. TODO
// 	}
// }

// func (c *CommitLog) updateHeader(row *Row) {
// 	// update the header of the commit log if
// 	// a new column family is encounter for the
// 	// first time
// 	table := openTable(c.table)
// 	for cName := range row.columnFamilies {
// 		id := table.tableMetadata.cfIDMap[cName]
// 		if c.clHeader.header[id] == 0 || (c.clHeader.header[id] == 1 &&
// 			c.clHeader.position[id] == 0) {
// 			// really ugly workaround for getting file current position
// 			// but I cannot find other way :(
// 			currentPos, err := c.logWriter.Seek(0, 0)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			c.logWriter.Seek(currentPos, 0)
// 			c.clHeader.turnOn(id, currentPos)
// 			c.writeCommitLogHeaderB(c.clHeader.toByteArray(), true)
// 		}
// 	}
// }

func (c *CommitLog) onMemtableFlush(tableName, cf string, cLogCtx *CommitLogContext) {
	// Called on memtable flush to add to the commit log a token
	// indicating that this column family has been flushed.
	// The bit flag associated with this column family is set
	// in the header and this is used to decide if the log
	// file can be deleted.
	table := OpenTable(tableName)
	id := table.tableMetadata.cfIDMap[cf]
	c.discard(cLogCtx, id)
}

// ByTime provide struct to sort file by timestamp
type ByTime []string

// Len implements the length of the slice
func (a ByTime) Len() int {
	return len(a)
}

// Less implements less comparator
func (a ByTime) Less(i, j int) bool {
	return getCreationTime(a[i]) < getCreationTime(a[j])
}

// Swap implements swap method
func (a ByTime) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func getCreationTime(name string) int64 {
	arr := strings.FieldsFunc(name, func(r rune) bool {
		return r == '-' || r == '.'
	})
	num, err := strconv.ParseInt(arr[len(arr)-2], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

// delete log segments whose contents have
// been turned into SSTables
func (c *CommitLog) discard(cLogCtx *CommitLogContext, id int) {
	// Check if old commit logs can be deleted.
	header, ok := clHeaders[cLogCtx.file]
	if !ok {
		if c.logFile == cLogCtx.file {
			// we are dealing with the current commit log
			header = c.clHeader
			clHeaders[cLogCtx.file] = c.clHeader
		} else {
			return
		}
	}
	//
	// commitLogHeader.turnOff(id)
	oldFiles := make([]string, 0)
	for key := range clHeaders {
		oldFiles = append(oldFiles, key)
	}
	sort.Sort(ByTime(oldFiles))
	// Loop through all the commit log files in the history.
	// Process the files that are older than the one in the
	// context. For each of these files the header needs to
	// modified by performing a bitwise & of the header with
	// the header of the file in the context. If we encounter
	// file in the context in our list of old commit log files
	// then we update the header and write it back to the commit
	// log.
	for _, oldFile := range oldFiles {
		if oldFile == cLogCtx.file {
			// Need to turn on again. Because we always keep
			// the bit turned on and the position indicates
			// from where the commit log needs to be read.
			// When a flush occurs we turn off perform &
			// operation and then turn on with the new position.
			header.turnOn(id, cLogCtx.position)
			if oldFile == c.logFile {
				c.seekAndWriteCommitLogHeader(header.toByteArray())
			} else {
				c.writeOldCommitLogHeader(cLogCtx.file, header)
			}
			break
		}
		header.turnOff(id)
		if header.isSafeToDelete() {
			log.Printf("Deleting commit log: %v\n", oldFile)
			err := os.Remove(oldFile)
			if err != nil {
				log.Fatal(err)
			}
			delete(clHeaders, oldFile)
		} else {
			c.writeOldCommitLogHeader(oldFile, header)
		}
	}
}
