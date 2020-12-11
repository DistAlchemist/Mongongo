// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import (
	"encoding/gob"
	"log"
	"math"
	"net/rpc"
	"sync"

	"github.com/DistAlchemist/Mongongo/config"
	"github.com/DistAlchemist/Mongongo/gms"
	"github.com/DistAlchemist/Mongongo/network"
)

// HHOMInstance ...
var (
	HHOMInstance *HintedHandOffManager
	hmu          sync.Mutex
)

// HintedHandOffManager ...
type HintedHandOffManager struct{}

// GetHintedHandOffManagerInstance ...
func GetHintedHandOffManagerInstance() *HintedHandOffManager {
	if HHOMInstance == nil {
		hmu.Lock()
		defer hmu.Unlock()
		HHOMInstance = &HintedHandOffManager{}
	}
	return HHOMInstance
}

// DeliverHintsToEndpoint ...
func DeliverHintsToEndpoint(endpoint *network.EndPoint) {
	log.Printf("started hinted handoff for endpoint %v\n", endpoint.HostName)
	targetEPBytes := endpoint.HostName
	// 1. scan through all the keys that we need to handoff
	// 2. for each key read the list of recipients if the endpoint matches send
	// 3. delete that recipient from theke if write was successful
	systemTable := openTable(config.SysTableName)
	for _, tableName := range config.GetTables() {
		hintedColumnFamily := systemTable.getCF(tableName, config.HintsCF)
		if hintedColumnFamily == nil {
			continue
		}
		keys := hintedColumnFamily.getSortedColumns()
		for _, keyColumn := range keys {
			keyStr := keyColumn.getName()
			endpoints := keyColumn.getSubColumns()
			for _, hintEndPoint := range endpoints {
				if hintEndPoint.getName() == targetEPBytes && sendMessage(endpoint.HostName, "", keyStr) {
					deleteEndPoint(hintEndPoint.getName(), tableName, keyColumn.getName(), keyColumn.timestamp())
					if len(endpoints) == 1 {
						deleteHintedData(tableName, keyStr)
					}
				}
			}
		}
	}
	log.Printf("finished hinted handoff for endpoint %v\n", endpoint.HostName)
}

// DeliverHints ...
func (h *HintedHandOffManager) DeliverHints(to *network.EndPoint) {
	go DeliverHintsToEndpoint(to)
}

func (h *HintedHandOffManager) submit(columnFamilyStore *ColumnFamilyStore) {
	go h.deliverAllHints(columnFamilyStore)
}

func (h *HintedHandOffManager) deliverAllHints(hintStore *ColumnFamilyStore) {
	log.Printf("start deliverAllHints\n")
	// 1. Scan through all the keys that we need to handoff
	// 2. For each key read the list of recipients and send
	// 3. Delete that recipient from the key if write was successful
	// 4. If all writes were success for a given key we can even delete the key
	// 5. Now force a flush
	// 6. Do major compaction to clean up all deletes etc.
	for _, tableName := range config.GetTables() {
		hintColumnFamily := removeDeleted(hintStore.getColumnFamily(NewIdentityQueryFilter(tableName, NewQueryPathCF(config.HintsCF))), math.MaxInt32)
		if hintColumnFamily == nil {
			continue
		}
		keys := hintColumnFamily.getSortedColumns()
		for _, keyColumn := range keys {
			endpoints := keyColumn.getSubColumns()
			keyStr := keyColumn.getName()
			deleted := 0
			for endpointStr := range endpoints {
				if sendMessage(endpointStr, tableName, keyStr) {
					deleteEndPoint(endpointStr, tableName, keyStr, keyColumn.timestamp())
					deleted++
				}
			}
			if deleted == len(endpoints) {
				deleteHintedData(tableName, keyStr)
			}
		}
	}
	hintStore.forceFlush()
	hintStore.forceCompaction(nil, nil, 0, nil)
	log.Print("Finished deliverAllHints")
}

func deleteEndPoint(endpointAddr, tableName, key string, timestamp int64) {
	rm := NewRowMutation(config.SysTableName, tableName)
	rm.Delete(NewQueryPath(config.HintsCF, []byte(key), []byte(endpointAddr)), timestamp)
	rm.ApplyE()
}

func deleteHintedData(tableName, key string) {
	// delete the row from application cfs: find
	// the largest timestamp in any of the data columns,
	// and delete the entire cf with that value for
	// the tombstone.
	// Note that we delete all data associated with the
	// key: this may be more than we sent earlier in
	// sendMessage, since HH is not serialized with
	// writes. This is sub-optimal but okay, sin HH
	// is just an effort to make a recovering node
	// more consistent than it would have been; we can
	// rely on the other consistency mechanisms to
	// finish the job in this corner case.
	rm := NewRowMutation(tableName, key)
	table := openTable(tableName)
	row := table.get(key) // not necessary to do removeDeleted here
	cfs := row.getColumnFamilies()
	for _, cf := range cfs {
		maxTS := int64(math.MinInt64)
		if cf.isSuper() == false {
			for _, col := range cf.getSortedColumns() {
				if col.timestamp() > maxTS {
					maxTS = col.timestamp()
				}
			}
		} else {
			for _, col := range cf.getSortedColumns() {
				if col.timestamp() > maxTS {
					maxTS = col.timestamp()
				}
				subColumns := col.getSubColumns()
				for _, subCol := range subColumns {
					if subCol.timestamp() > maxTS {
						maxTS = subCol.timestamp()
					}
				}
			}
		}
		rm.Delete(NewQueryPathCF(cf.ColumnFamilyName), maxTS)
	}
	rm.ApplyE()
}

func sendMessage(endpointAddr, tableName, key string) bool {
	endPoint := network.NewEndPointH(endpointAddr, config.StoragePort)
	if !gms.GetFailureDetector().IsAlive(*endPoint) {
		return false
	}
	table := openTable(tableName)
	row := table.get(key)
	purgedRow := NewRow(key)
	for _, cf := range row.getColumnFamilies() {
		purgedRow.addColumnFamily(removeDeletedGC(cf))
	}
	rm := NewRowMutationR(tableName, purgedRow)
	return sendEndPointRM(endPoint, rm)
}

func sendEndPointRM(end *network.EndPoint, rm *RowMutation) bool {
	gob.Register(SuperColumnFactory{})
	gob.Register(SuperColumn{})
	c, err := rpc.DialHTTP("tcp", end.HostName+":"+end.Port)
	defer c.Close()
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := RowMutationArgs{*rm}
	reply := RowMutationReply{}
	err = c.Call("StorageService.DoRowMutation", &args, &reply)
	if err != nil {
		log.Fatal("calling:", err)
	}
	// fmt.Printf("DoRowMutation.Result for %v:%v: %+v\n",
	// 	end.HostName, end.Port, reply.Result)
	return reply.Status
}
