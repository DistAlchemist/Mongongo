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

	"github.com/DistAlchemist/Mongongo/config"
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

func insertBlocking(rm db.RowMutation, consistencyLevel int) {
	// TODO
}

func insert(rm db.RowMutation) {
	// use this method to have this RowMutation applied
	// across all replicas. This method will take care
	// of the possiblity of a replica being down and
	// hint the data across to some other replica.

	// 1. get the N nodes from storage service where the
	// data needs to be replicated
	// 2. construct a message for write
	// 3. send them asynchronously to the replicas
	// startTime := utils.CurrentTimeMillis()
	// this is the ZERO consistency level, so user doesn't
	// care if we don't really have N destinations available.
	endpointMap := GetInstance().getHintedStorageEndpointMap(rm.RowKey)
	messageMap := createWriteMessage(rm, endpointMap)
	reply := db.RowMutationReply{}
	for endpoint, message := range messageMap {
		log.Printf("insert writing key %v to %v\n", rm.RowKey, endpoint)
		to := endpoint
		client, err := rpc.DialHTTP("tcp", to.HostName+":"+config.StoragePort)
		if err != nil {
			log.Fatal("dialing: ", err)
		}
		client.Call("StorageService.DoRowMutation", &message, &reply)
		log.Printf("row mutation status for %v: %v\n", to, reply)
	}
}

// WriteMessage ...
// type WriteMessage struct {
// 	HeaderKey   string
// 	HeaderValue network.EndPoint
// 	From        network.EndPoint
// 	RM          db.RowMutation
// }

func createWriteMessage(rm db.RowMutation, endpointMap map[network.EndPoint]network.EndPoint) map[network.EndPoint]db.RowMutationArgs {
	messageMap := make(map[network.EndPoint]db.RowMutationArgs)
	message := db.RowMutationArgs{}
	message.RM = rm
	message.From = *GetInstance().tcpAddr
	for target, hint := range endpointMap {
		if target != hint {
			hintedMessage := db.RowMutationArgs{}
			hintedMessage.HeaderKey = db.HINT
			hintedMessage.HeaderValue = hint
			hintedMessage.From = *GetInstance().tcpAddr
			hintedMessage.RM = rm
			log.Printf("sending the hint of %v to %v \n", hint.HostName, target.HostName)
			messageMap[target] = hintedMessage
		} else {
			messageMap[target] = message
		}
	}
	return messageMap
}
