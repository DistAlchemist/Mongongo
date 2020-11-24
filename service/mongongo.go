package service

import (
	"log"

	"github.com/DistAlchemist/Mongongo/mql"
)

// Mongongo expose the interface of operations
type Mongongo struct {
	// Mongongo struct
	Hostname       string
	Port           int
	storageService *StorageService
}

// ExecuteArgs arguments of executeQueryOnServer
type ExecuteArgs struct {
	Line string
}

// ExecuteReply reply format of executeQueryOnServer
type ExecuteReply struct {
	Result mql.Result
}

// Insert updates tableNmae.key.columnFamilyColumn with (cellData, timeStamp)
func (mg *Mongongo) insert(tableName, key, columnFamilyColumn, cellData string, timeStamp int64) error {
	//
	return nil
}

// GetColumn get the value of tableName.key.columnFamilyColumn
func (mg *Mongongo) getColumn(tableName, key, columnFamilyColumn string) error {
	//
	return nil
}

// Remove delete the value of tableName.key.columnFamilyColumn (lazily)
func (mg *Mongongo) remove(tableName, key, columnFamilyColumn string) error {
	//
	return nil
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
	mg.storageService = GetInstance()
	mg.storageService.Start()
}
