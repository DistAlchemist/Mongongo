// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import (
	"encoding/binary"
	"log"
	"os"
	"strconv"
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
	clInstance = map[string]*CommitLog{}
	clHeaders  = map[string]*CommitLogHeader{}
	clmu       sync.Mutex
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

func (c *CommitLog) setNextFileName() {
	c.logFile = config.LogFileDir + string(os.PathSeparator) +
		"CommitLog-" + c.table + "-" +
		strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10) +
		".log"
}

func (c *CommitLog) createWriter(file string) *os.File {
	f, err := os.Create(file)
	if err != nil {
		log.Print(err)
	}
	return f
}

func (c *CommitLog) writeCommitLogHeader() {
	// writes a header with all bits set to zero
	table := openTable(c.table)
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

// NewCommitLog creates a new commit log
func NewCommitLog(table string, recoveryMode bool) *CommitLog {
	c := &CommitLog{}
	c.table = table
	c.forcedRollOver = false
	if !recoveryMode {
		c.setNextFileName()
		c.logWriter = c.createWriter(c.logFile)
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

// add the specified row to the commit log. This method will
// reset the file offset to what it is before the start of
// the operation in case of any problems. This way we can
// assume that the subsequent commit log entry will override
// the garbage left over by the previous write.
func (c *CommitLog) add(row *Row) *CommitLogContext {
	// serialize the row
	buf := row.toByteArray()
	//
	currentPos, err := c.logWriter.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	c.logWriter.Seek(currentPos, 0)
	cLogCtx := NewCommitLogContext(c.logFile, currentPos)
	// update the header
	c.updateHeader(row)
	// write key (table name)
	writeString(c.logWriter, c.table)
	// write row bytes
	writeBytes(c.logWriter, buf)
	// flush to disk
	err = c.logWriter.Sync()
	if err != nil {
		log.Print(err)
	}
	fileInfo, err := c.logWriter.Stat()
	if err != nil {
		log.Print(err)
	}
	// length in bytes
	fileSize := fileInfo.Size()
	c.checkThresholdAndRollLog(fileSize)
	// reset file offset if any error occurs TODO
	return cLogCtx
}

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

func (c *CommitLog) checkThresholdAndRollLog(fileSize int64) {
	if fileSize >= config.LogRotationThres || c.forcedRollOver {
		// rolls the current log file over to a new one
		c.setNextFileName()
		oldLogFile := c.logWriter.Name()
		c.logWriter.Close()
		// change logWriter to new log file
		c.logWriter = c.createWriter(c.logFile)
		// save old log header
		clHeaders[oldLogFile] = c.clHeader.copy()
		// zero out positions in old file log header
		c.clHeader.zeroPositions()
		c.writeCommitLogHeaderB(c.clHeader.toByteArray(), false)
		// Get the list of files in commit log dir if it is greater than a
		// certain number. Force flush all the column families that way we
		// ensure that a slowly populated column family is not screwing up
		// by accumulating the commit log. TODO
	}
}

func (c *CommitLog) updateHeader(row *Row) {
	// update the header of the commit log if
	// a new column family is encounter for the
	// first time
	table := openTable(c.table)
	for cName := range row.columnFamilies {
		id := table.tableMetadata.cfIDMap[cName]
		if c.clHeader.header[id] == 0 || (c.clHeader.header[id] == 1 &&
			c.clHeader.position[id] == 0) {
			// really ugly workaround for getting file current position
			// but I cannot find other way :(
			currentPos, err := c.logWriter.Seek(0, 0)
			if err != nil {
				log.Fatal(err)
			}
			c.logWriter.Seek(currentPos, 0)
			c.clHeader.turnOn(id, currentPos)
			c.writeCommitLogHeaderB(c.clHeader.toByteArray(), true)
		}
	}
}
