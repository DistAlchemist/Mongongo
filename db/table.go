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

// GetTableMetadataInstance will create an instance if not exists
func GetTableMetadataInstance() *TableMetadata {
	if tableMetadata == nil {
		// file := getFileName()
		tableMetadata = &TableMetadata{"TableMetadata", "PrimaryCardinality",
			nil, nil, nil}
	}
	return tableMetadata
}

func (t *TableMetadata) isEmpty() bool {
	return t.cfIDMap == nil
}

func (t *TableMetadata) add(cf string, id int, tp string) {
	t.cfIDMap[cf] = id
	t.idCfMap[id] = cf
	t.cfTypeMap[cf] = tp
}

func getFileName() string {
	table := config.Tables[0]
	return config.MetadataDir + string(os.PathSeparator) +
		table + "-Metadata.db"
}
