package db

import (
	"fmt"
	"log"
	"net/rpc"
)

type RowMutationArgs struct {
	RM RowMutation
}

type RowMutationReply struct {
	Result string
}

// Insert dispatches rowmutation to other replicas
func Insert(rm RowMutation) string {
	//
	// endpointMap := GetInstance().getNStorageEndPointMap(rm.rowKey)
	// oversimplified: should get endpoint list to write
	c, err := rpc.DialHTTP("tcp", "localhost"+":"+"11111")
	defer c.Close()
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := RowMutationArgs{}
	reply := RowMutationReply{}
	args.RM = rm
	err = c.Call("StorageService.DoRowMutation", &args, &reply)
	if err != nil {
		log.Fatal("calling:", err)
	}
	fmt.Printf("DoRowMutation.Result: %+v\n", reply.Result)
	return reply.Result
}
