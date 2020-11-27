// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package service

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/rpc"

	"github.com/DistAlchemist/Mongongo/db"
	"github.com/DistAlchemist/Mongongo/network"
)

// RowMutationArgs for rm arguments
type RowMutationArgs struct {
	RM db.RowMutation
}

// RowMutationReply for rm reply structure
type RowMutationReply struct {
	Result string
}

// Insert will apply this row mutation to
// all replicas. (TODO) It will take care of the
// possibility of a replica being down and
// hint the data across to some other replica.
func Insert(rm db.RowMutation) {
	endpointMap := GetInstance().getNStorageEndPointMap(rm.RowKey)
	gob.Register(db.SuperColumnFactory{})
	gob.Register(db.SuperColumn{})
	for endpoint := range endpointMap {
		go func(end network.EndPoint) {
			c, err := rpc.DialHTTP("tcp", end.HostName+":"+end.Port)
			defer c.Close()
			if err != nil {
				log.Fatal("dialing:", err)
			}
			args := RowMutationArgs{rm}
			reply := RowMutationReply{}
			err = c.Call("StorageService.DoRowMutation", &args, &reply)
			if err != nil {
				log.Fatal("calling:", err)
			}
			fmt.Printf("DoRowMutation.Result for %v:%v: %+v\n",
				end.HostName, end.Port, reply.Result)
		}(endpoint)
	}
	return
}
