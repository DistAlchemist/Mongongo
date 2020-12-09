// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import "os"

// CFSerializer ...
var CFSerializer = NewCFSerializer()

// ColumnFamilySerializer ...
type ColumnFamilySerializer struct{}

// NewCFSerializer ...
func NewCFSerializer() *ColumnFamilySerializer {
	c := &ColumnFamilySerializer{}
	return c
}

func (c *ColumnFamilySerializer) deserializeFromSSTableNoColumns(cf *ColumnFamily, input *os.File) *ColumnFamily {
	localtime := readInt(input)
	timestamp := readInt64(input)
	cf.delete(localtime, timestamp)
	return cf
}

func (c *ColumnFamilySerializer) serializeWithIndexes(columnFamily *ColumnFamily, dos []byte) {
	CIndexer.serialize(columnFamily, dos)
	c.serializeForSSTable(columnFamily, dos)
}

func (c *ColumnFamilySerializer) serializeForSSTable(columnFamily *ColumnFamily, dos []byte) {
	writeIntB(dos, columnFamily.localDeletionTime)
	writeInt64B(dos, columnFamily.markedForDeleteAt)
	columns := columnFamily.getSortedColumns()
	writeIntB(dos, len(columns))
	for _, column := range columns {
		columnFamily.getColumnSerializer().serializeB(column, dos)
	}
}
