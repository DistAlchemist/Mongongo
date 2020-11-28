// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package config

import (
	"log"
	"os"
)

var (
	// Random is one of hashing strategy
	Random = "RANDOM"
	// Ophf is one of hashing strategy
	Ophf = "OPHF"
	// StoragePort ...
	StoragePort = 7000
	// ControlPort ...
	ControlPort = 7001
	// HTTPPort ...
	HTTPPort = 7002
	// ClusterName ...
	ClusterName = "Test"
	// ReplicationFactor ...
	ReplicationFactor = 3
	// RPCTimeoutInMillis set 2s by default
	RPCTimeoutInMillis = 2000
	// Seeds is a set of nodes to connect to when a new node join the cluster
	Seeds map[string]bool
	// MetadataDir is the dir for store meta data
	MetadataDir = "var/storage/system" // pre read from file
	// SnapshotDir for snapshot
	SnapshotDir = MetadataDir + string(os.PathSeparator) + "snapshot" // pre
	// MapOutputDirs keep the list of map output dirs
	MapOutputDirs []string
	// DataFileDirs keep the list of data file dirs
	DataFileDirs = []string{"var/storage/data"} // pre
	// CurIndex stores the current index into the above list of dirs
	CurIndex = 0
	// LogFileDir for log file
	LogFileDir = "var/storage/commitlog" // pre
	// BootstrapFileDir for bootstrap file
	BootstrapFileDir = "var/storage/bootstrap" // pre

	// LogRotationThres for log rotation threshold, defaults to 128MB
	LogRotationThres = 128 * 1024 * 1024
	// FastSync defaults to false
	FastSync = false
	// RackAware for replica distribution, default: false
	RackAware = false

	// Tables for list of table name
	// currently we cannot change the schema online
	// this will be improved in the future
	Tables = []string{"table1", "table2"} // TO BE IMPROVED

	// ApplicationColumnFamilies is a set of column family names
	ApplicationColumnFamilies = map[string]bool{
		"standardCF1": true,
		"standardCF2": true,
		"superCF1":    true,
		"superCF2":    true,
	}

	// TableToCFMetaData map table names to column families and corresponding meta data
	// TableToCFMetaData map[string]map[string]CFMetaData
	TableToCFMetaData = map[string]map[string]CFMetaData{
		"table1": {
			"standardCF1": {"table1", "standardCF1", "Standard", "Timestamp",
				"row1", "", "", "column1", "", "", ""},
			"superCF1": {"table1", "superCF1", "Super", "Name",
				"row2", "superCM", "superCK", "column2", "", "", ""},
		},
		"table2": {
			"standardCF2": {"table2", "standardCF2", "Standard", "Name",
				"row1", "", "", "column2", "", "", ""},
			"superCF2": {"table2", "superCF2", "Super", "Timestamp",
				"row2", "superCM", "superCK", "column2", "", "", ""},
		},
	}

	// HashingStrategy : Random or OPHF
	HashingStrategy = Random
	// ColumnIndexSizeInKB : indexing will kick in if size of column/supercolumn excedes
	ColumnIndexSizeInKB = 64 // pre read from file
	// TouchKeyCacheSize specifies sizeof touch key cache
	TouchKeyCacheSize = 1024
	// MemtableLifetime is the number of hours to keep a memtable in memory
	MemtableLifetime = 6
	// MemtableSize is the size of memtable in memory before it is dumped
	MemtableSize = 128
	// MemtableObjectCount is the number of objects in millions in the memtable before it is dumped
	MemtableObjectCount = 1
	// DoConsistencyCheck enables or disables consistency checks.
	// false results in high read throughput at the cost of consistency
	DoConsistencyCheck = true
	// JobTrackerHost is the address where to run the job tracker
	JobTrackerHost string
	// ConfigFileName is the path to config file
	ConfigFileName string
	// RingRange is the size of consistent hashing ring
	RingRange = uint64(1 << 32)
)

// DatabaseDescriptor contains meta data for the underlying storage system
type DatabaseDescriptor struct {
}

// GetTableMetaData read CFMetaData through map using table name
func GetTableMetaData(tableName string) map[string]CFMetaData {
	return TableToCFMetaData[tableName]
}

// Init read the configuration file to retrieve DB related properties
func Init() map[string]map[string]CFMetaData {
	ConfigFileName = os.Getenv("storage-config") + string(os.PathSeparator) + "storage-conf.json"
	return initInternal(ConfigFileName)
}

func mkdir(path string) {
	err := os.MkdirAll(path, 0700)
	/* +-----+---+--------------------------+
	   | rwx | 7 | Read, write and execute  |
	   | rw- | 6 | Read, write              |
	   | r-x | 5 | Read, and execute        |
	   | r-- | 4 | Read,                    |
	   | -wx | 3 | Write and execute        |
	   | -w- | 2 | Write                    |
	   | --x | 1 | Execute                  |
	   | --- | 0 | no permissions           |
	   +------------------------------------+

	   +------------+------+-------+
	   | Permission | Octal| Field |
	   +------------+------+-------+
	   | rwx------  | 0700 | User  |
	   | ---rwx---  | 0070 | Group |
	   | ------rwx  | 0007 | Other |
	   +------------+------+-------+*/
	if err != nil {
		log.Printf("error when mkdir %v\n", MetadataDir)
	}
}

func initInternal(file string) map[string]map[string]CFMetaData {
	// I will not use file, however..
	mkdir(MetadataDir)
	mkdir(SnapshotDir)
	for _, path := range DataFileDirs {
		mkdir(path)
	}
	mkdir(LogFileDir)
	mkdir(BootstrapFileDir)
	return TableToCFMetaData
}

// GetColumnType retrieve column type from cf metadata
func GetColumnType(cfName string) string {
	table := Tables[0]
	cfMetadata, ok := TableToCFMetaData[table][cfName]
	if !ok {
		return ""
	}
	return cfMetadata.ColumnType
}
