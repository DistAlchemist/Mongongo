// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package config

var (
	// Random is one of hashing strategy
	Random = "RANDOM"
	// Ophf is one of hashing strategy
	Ophf                      = "OPHF"
	// StoragePort ...
	StoragePort               = 7000
	// ControlPort ...
	ControlPort               = 7001
	// HTTPPort ...
	HTTPPort                  = 7002
	// ClusterName ...
	ClusterName               = "Test"
	// ReplicationFactor ...
	ReplicationFactor         = 3
	// RPCTimeoutInMillis set 2s by default
	RPCTimeoutInMillis        = 2000
	// Seeds is a set of nodes to connect to when a new node join the cluster
	Seeds                     map[string]bool
	// MetadataDir is the dir for store meta data
	MetadataDir               string
	// SnapshotDir for snapshot
	SnapshotDir               string
	// MapOutputDirs keep the list of map output dirs
	MapOutputDirs             []string
	// DataFileDirs keep the list of data file dirs
	DataFileDirs              []string
	// CurIndex stores the current index into the above list of dirs
	CurIndex                  = 0 
	// LogFileDir for log file
	LogFileDir                string
	// BootstrapFileDir for bootstrap file
	BootstrapFileDir          string
	// LogRotationThres for log rotation threshold, defaults to 128MB
	LogRotationThres          = 128 * 1024 * 1024
	// FastSync defaults to false
	FastSync                  = false
	// RackAware for replica distribution, default: false
	RackAware                 = false
	// Tables for list of table name
	Tables                    []string
	// ApplicationColumnFamilies is a set of column family names
	ApplicationColumnFamilies []map[string]bool

	// TableToCFMetaData map table names to column families and corresponding meta data
	TableToCFMetaData map[string]map[string]CFMetaData
	// HashingStrategy : Random or OPHF
	HashingStrategy = Random
	// ColumnIndexSizeInKB : indexing will kick in if size of column/supercolumn excedes
	ColumnIndexSizeInKB int 
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
)

// DatabaseDescriptor contains meta data for the underlying storage system
type DatabaseDescriptor struct {
}

// GetTableMetaData read CFMetaData through map using table name
func GetTableMetaData(tableName string) map[string]CFMetaData {
	return TableToCFMetaData[tableName]
}
