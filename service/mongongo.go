// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package service

import (
	"log"

	"github.com/DistAlchemist/Mongongo/db"
)

// Mongongo expose the interface of operations
type Mongongo struct {
	// Mongongo struct
	Hostname string
	Port     int
	// storageService *StorageService
}

// // ExecuteArgs arguments of executeQueryOnServer
// type ExecuteArgs struct {
// 	Line string
// }

// // ExecuteReply reply format of executeQueryOnServer
// type ExecuteReply struct {
// 	Result mql.Result
// }

// // ExecuteQueryOnServer handles the rpc from cli client
// func (mg *Mongongo) ExecuteQueryOnServer(args *ExecuteArgs, reply *ExecuteReply) error {
// 	//
// 	line := args.Line
// 	log.Printf("server executing %+v\n", line)
// 	reply.Result = mql.ExecuteQuery(line)
// 	return nil
// }

// Start setup other service such as storageService
func (mg *Mongongo) Start() {
	GetInstance().Start()
}

// InsertArgs ...
type InsertArgs struct {
	Table            string
	Key              string
	CPath            ColumnPath
	Value            []byte
	Timestamp        int64
	ConsistencyLevel int
}

// InsertReply ...
type InsertReply struct {
	Result string
}

// Insert is an rpc
func (mg *Mongongo) Insert(args *InsertArgs, reply *InsertReply) error {
	log.Printf("enter mg.Insert\n")
	table := args.Table
	key := args.Key
	columnPath := args.CPath
	value := args.Value
	timestamp := args.Timestamp
	consistencyLevel := args.ConsistencyLevel
	rm := db.NewRowMutation(table, key)
	rm.AddQ(db.NewQueryPath(columnPath.ColumnFamily, columnPath.SuperColumn, columnPath.Column),
		value, timestamp)
	mg.doInsert(consistencyLevel, rm)
	return nil
}

func (mg *Mongongo) doInsert(consistencyLevel int, rm db.RowMutation) {
	if consistencyLevel != 0 {
		insertBlocking(rm, consistencyLevel)
	} else {
		insert(rm)
	}
}
