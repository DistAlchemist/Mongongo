// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package config

import (
	"log"
	"os"
)

const (
	// Batch for commitlog sync option
	Batch = iota
	// Periodic for commitlog sync option
	Periodic
)

const (
	// Random is one of hashing strategy
	Random = "RANDOM"
	// Ophf is one of hashing strategy
	Ophf = "OPHF"
)

var (
	// ClusterName ...
	ClusterName = "Test"
	// StoragePort ...
	StoragePort = "11170"
	// ControlPort ...
	ControlPort = "21170"
	// HTTPPort ...
	HTTPPort = "31170"
	// ReplicationFactor ...
	ReplicationFactor = 3
	// RPCTimeoutInMillis set 5s by default
	RPCTimeoutInMillis = 5000
	// GcGraceInSeconds defaults to 10 days
	GcGraceInSeconds = 10 * 24 * 3600
	// Seeds is a set of nodes to connect to when a new node join the cluster
	Seeds = map[string]bool{
		"thumm01": true,
	}
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

	// FastSync defaults to false
	FastSync = false
	// CommitLogSync can be either periodic or batch
	CommitLogSync = Periodic
	// CommitLogSyncPeriodInMS defaults to 1000 i.e. 1s
	CommitLogSyncPeriodInMS = 1000
	// InitialToken defaults to empty
	InitialToken = ""
	// RackAware for replica distribution, default: false
	RackAware = false
	// SysTableName is the table name for system
	SysTableName = "system"
	// HintsCF is the cf name for hinted handoff
	HintsCF = "HintsColumnFamily"
	// Tables for list of table name
	// currently we cannot change the schema online
	// this will be improved in the future
	Tables = []string{SysTableName, "table1", "table2"} // TO BE IMPROVED

	// ApplicationColumnFamilies is a set of column family names
	ApplicationColumnFamilies = map[string]bool{
		"standardCF1": true,
		"standardCF2": true,
		"superCF1":    true,
		"superCF2":    true,
	}

	// SystemMetadata stores cf for system table
	SystemMetadata = map[string]CFMetaData{
		"LocationInfo": {
			SysTableName,   // TableName
			"LocationInfo", // CFName
			"Standard",     // ColumnType
			"Name",         // IndexProperty
			"row",          // NRowKey
			"",             // NSuperColumnMap
			"",             // NSuperColumnKey
			"column",       // NColumnMap
			"",             // NColumnKey
			"",             // NColumnValue
			""},            // NColumnTimestamp
		"HintsColumnFamily": {
			SysTableName,        // TableName
			"HintsColumnFamily", // CFName
			"Super",             // ColumnType
			"Name",              // IndexProperty
			"row0",              // NRowKey
			"superCM",           // NSuperColumnMap
			"superCK",           // NSuperColumnKey
			"column0",           // NColumnMap
			"",                  // NColumnKey
			"",                  // NColumnValue
			""},                 // NColumnTimestamp
	}

	// TableToCFMetaData map table names to column families and corresponding meta data
	// TableToCFMetaData map[string]map[string]CFMetaData
	// type CFMetaData struct {
	// 	TableName     string // name of table which has this column family
	// 	CFName        string // name of column family
	// 	ColumnType    string // standard or super
	// 	IndexProperty string // name sorted or timestamp sorted
	// 	NRowKey          string
	// 	NSuperColumnMap  string // only used in super column family
	// 	NSuperColumnKey  string // only used in super column family
	// 	NColumnMap       string
	// 	NColumnKey       string
	// 	NColumnValue     string
	// 	NColumnTimestamp string
	// }
	TableToCFMetaData = map[string]map[string]CFMetaData{
		SysTableName: SystemMetadata,
		"table1": {
			"standardCF1": {
				"table1",      // TableName
				"standardCF1", // CFName
				"Standard",    // ColumnType
				"Timestamp",   // IndexProperty
				"row1",        // NRowKey
				"",            // NSuperColumnMap
				"",            // NSuperColumnKey
				"column1",     // NColumnMap
				"",            // NColumnKey
				"",            // NColumnValue
				""},           // NColumnTimestamp
			"superCF1": {
				"table1",   // TableName
				"superCF1", // CFName
				"Super",    // ColumnType
				"Name",     // IndexProperty
				"row2",     // NRowKey
				"superCM",  // NSuperColumnMap
				"superCK",  // NSuperColumnKey
				"column2",  // NColumnMap
				"",         // NColumnKey
				"",         // NColumnValue
				""},        // NColumnTimestamp
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
	// MinCompactionThres ...
	MinCompactionThres = 4
	// MaxCompactionThres ...
	MaxCompactionThres = 32
	// LogRotationThres for log rotation threshold, defaults to 128MB
	LogRotationThres = int64(128 * 1024 * 1024)
	// ColumnIndexSizeInKB : indexing will kick in if size of column/supercolumn excedes
	ColumnIndexSizeInKB = 64 // pre read from file
	// TouchKeyCacheSize specifies sizeof touch key cache
	TouchKeyCacheSize = 1024
	// MemtableLifetime is the number of hours to keep a memtable in memory
	MemtableLifetime = 6
	// MemtableSize is the size of memtable in memory before it is dumped
	MemtableSize = 128 // 128 MB
	// MemtableObjectCount is the number of objects in millions in the memtable before it is dumped
	MemtableObjectCount = 1
	// FlushDataBufferSizeInMB is the buffer size to use when flushing memtables to disk.
	// Only one memtable is ever flushed at a time.
	FlushDataBufferSizeInMB = 32
	// FlushIndexBufferSizeInMB is the buffer size when flushing index to disk
	FlushIndexBufferSizeInMB = 8
	// DoConsistencyCheck enables or disables consistency checks.
	// false results in high read throughput at the cost of consistency
	DoConsistencyCheck = true
	// SnapshotBeforeCompaction defaults to false
	SnapshotBeforeCompaction = false
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
	// make sure all tables have directory
	for _, path := range DataFileDirs {
		mkdir(path + string(os.PathSeparator) + SysTableName)
		for _, table := range Tables {
			mkdir(path + string(os.PathSeparator) + table)
		}
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

// GetColumnTypeTableName retrieve column type from cf metadata
func GetColumnTypeTableName(table string, cfName string) string {
	cfMetadata, ok := TableToCFMetaData[table][cfName]
	if !ok {
		return ""
	}
	return cfMetadata.ColumnType
}

// GetCompactionFileLocation ...
func GetCompactionFileLocation(expectedCompactedFileSize int64) string {
	// TODO
	return ""
}

// GetDataFileLocationForTable will loop through all the disks to see
// which disk has the max free space. Return the disk with max free
// space for compactions. If the size of the expected compacted file
// is greater than the max disk space available return "", we cannot
// do compaction in this case.
func GetDataFileLocationForTable(table string, expectedCompactedFileSize int64) string {
	// maxFreeDisk := int64(0)
	// maxDiskIndex := 0
	// dataFileDirectory := ""
	dataDirectoryForTable := GetAllDataFileLocationsForTable(table)
	// for i := 0; i < len(dataDirectoryForTable); i++ {
	// 	f, err := os.Open(dataDirectoryForTable[i])
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fileInfo, err :=  f.Stat()
	// 	fileInfo.Sys()
	// }
	// Here gives a very simplified version
	return dataDirectoryForTable[0]
}

// GetAllDataFileLocationsForTable gets a list of data dirs for a given table
func GetAllDataFileLocationsForTable(table string) []string {
	tableLocations := make([]string, len(DataFileDirs))
	for i, dir := range DataFileDirs {
		tableLocations[i] = dir + string(os.PathSeparator) + table
	}
	return tableLocations
}

// GetColumnIndexSize returns size in MB
func GetColumnIndexSize() int {
	return ColumnIndexSizeInKB * 1024
}

// GetTables ...
func GetTables() []string {
	return Tables
}
