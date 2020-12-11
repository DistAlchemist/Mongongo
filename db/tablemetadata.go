// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import (
	"os"

	"github.com/DistAlchemist/Mongongo/config"
)

var tableMetadata *TableMetadata

// TableMetadata stores infos about table and its columnFamilies
type TableMetadata struct {
	cfIDMap   map[string]int
	cfTypeMap map[string]string
}

// NewTableMetadata initializes a TableMetadata
func NewTableMetadata() *TableMetadata {
	t := &TableMetadata{}
	t.cfIDMap = make(map[string]int)
	t.cfTypeMap = make(map[string]string)
	return t
}

func (t *TableMetadata) isEmpty() bool {
	return t.cfIDMap == nil
}

// Add adds column family, id and typename to table metadata
func (t *TableMetadata) Add(cf string, id int, tp string) {
	t.cfIDMap[cf] = id
	idCFMap[id] = cf
	t.cfTypeMap[cf] = tp
}

func getFileName() string {
	table := config.Tables[0]
	return config.MetadataDir + string(os.PathSeparator) +
		table + "-Metadata.db"
}

func (t *TableMetadata) isValidColumnFamily(cfName string) bool {
	_, ok := t.cfIDMap[cfName]
	return ok
}

func (t *TableMetadata) getSize() int {
	return len(t.cfIDMap)
}

func (t *TableMetadata) getColumnFamilyID(cfName string) int {
	return t.cfIDMap[cfName]
}
