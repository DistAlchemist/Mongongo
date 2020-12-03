// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import (
	"sync"

	"github.com/DistAlchemist/Mongongo/config"
)

var (
	mu        sync.Mutex
	minstance *Manager
)

// Manager manages database
type Manager struct {
}

// GetManagerInstance return an instance of DBManager
func GetManagerInstance() *Manager {
	mu.Lock()
	defer mu.Unlock()
	if minstance == nil {
		minstance = &Manager{}
		minstance.init()
	}
	return minstance
}

func (d *Manager) init() {
	// read the config file
	tableToColumnFamily := config.Init()
	for table := range tableToColumnFamily {
		tbl := openTable(table)
		tbl.onStart()
	}
	recoveryMgr := GetRecoveryManager()
	recoveryMgr.doRecovery()
	// config.Init()
	// storeMetadata(tableToColumnFamily) // useless

}

// // create metadata tables. table stores tableName and columnFamilyName
// // each column family has an ID
// func storeMetadata(tableToColumnFamily map[string]map[string]config.CFMetaData) error {
// 	var cnt int32
// 	cnt = 0
// 	for _, columnFamilies := range tableToColumnFamily {
// 		tmetadata := GetTableMetadataInstance()
// 		if tmetadata.isEmpty() {
// 			for columnFamily := range columnFamilies {
// 				tmetadata.add(columnFamily, int(atomic.AddInt32(&cnt, 1)),
// 					config.GetColumnType(columnFamily))
// 			}
// 			tmetadata.add()
// 		}
// 	}
// 	return nil
// }

// Start reads the system table and retrives the metadata
// associated with this storage instance. The metadata is
// stored in a Column Family called LocationInfo which has
// two columns: "Token" and "Generation".
func (d *Manager) Start() *StorageMetadata {
	storageMetadata := &StorageMetadata{}
	return storageMetadata
}

// StorageMetadata stores id and generation
type StorageMetadata struct {
	storageID  uint64
	generation int
}

// RowMutationArgs for rm arguments
type RowMutationArgs struct {
	RM RowMutation
}

// RowMutationReply for rm reply structure
type RowMutationReply struct {
	Result string
}

// // Insert dispatches rowmutation to other replicas
// func Insert(rm RowMutation) string {
// 	// endpointMap := GetInstance().getNStorageEndPointMap(rm.rowKey)
// 	// oversimplified: should get endpoint list to write
// 	c, err := rpc.DialHTTP("tcp", "localhost"+":"+"11111")
// 	defer c.Close()
// 	if err != nil {
// 		log.Fatal("dialing:", err)
// 	}
// 	args := RowMutationArgs{}
// 	reply := RowMutationReply{}
// 	args.RM = rm
// 	err = c.Call("StorageService.DoRowMutation", &args, &reply)
// 	if err != nil {
// 		log.Fatal("calling:", err)
// 	}
// 	fmt.Printf("DoRowMutation.Result: %+v\n", reply.Result)
// 	return reply.Result
// }
