// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package server

import (
	"log"

	"github.com/DistAlchemist/Mongongo/mql"
	"github.com/DistAlchemist/Mongongo/service"
)

// Mongongo expose the interface of operations
type Mongongo struct {
	// Mongongo struct
	Hostname string
	Port     int
	// storageService *StorageService
}

// ExecuteArgs arguments of executeQueryOnServer
type ExecuteArgs struct {
	Line string
}

// ExecuteReply reply format of executeQueryOnServer
type ExecuteReply struct {
	Result mql.Result
}

// ExecuteQueryOnServer handles the rpc from cli client
func (mg *Mongongo) ExecuteQueryOnServer(args *ExecuteArgs, reply *ExecuteReply) error {
	//
	line := args.Line
	log.Printf("server executing %+v\n", line)
	reply.Result = mql.ExecuteQuery(line)
	return nil
}

// Start setup other service such as storageService
func (mg *Mongongo) Start() {
	service.GetInstance().Start()
}
