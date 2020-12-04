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
	cfName      string
	cardinality string
	cfIDMap     map[string]int
	idCfMap     map[int]string
	cfTypeMap   map[string]string
}

// // GetTableMetadataInstance will create an instance if not exists
// func GetTableMetadataInstance() *TableMetadata {
// 	if tableMetadata == nil {
// 		// file := getFileName()
// 		tableMetadata = &TableMetadata{"TableMetadata", "PrimaryCardinality",
// 			nil, nil, nil}
// 	}
// 	return tableMetadata
// }

// NewTableMetadata initializes a TableMetadata
func NewTableMetadata() *TableMetadata {
	t := &TableMetadata{}
	t.cfName = "TableMetadata"
	t.cardinality = "PrimaryCardinality"
	t.cfIDMap = make(map[string]int)
	t.idCfMap = make(map[int]string)
	t.cfTypeMap = make(map[string]string)
	return t
}

func (t *TableMetadata) isEmpty() bool {
	return t.cfIDMap == nil
}

// Add adds column family, id and typename to table metadata
func (t *TableMetadata) Add(cf string, id int, tp string) {
	t.cfIDMap[cf] = id
	t.idCfMap[id] = cf
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
