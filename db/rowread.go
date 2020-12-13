// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import "github.com/DistAlchemist/Mongongo/network"

// RowReadArgs ...
type RowReadArgs struct {
	HeaderKey   string
	HeaderValue network.EndPoint
	From        network.EndPoint
	RCommand    ReadCommand
	// RM          RowMutation
}

// RowReadReply ...
type RowReadReply struct {
	Result string
	Status bool
	R      *Row
}

// DoRowRead ...
func DoRowRead(args *RowReadArgs, reply *RowReadReply) error {
	readCommand := args.RCommand
	table := OpenTable(readCommand.GetTable())
	row := readCommand.GetRow(table)
	reply.R = row
	reply.Result = "SUCESS"
	return nil
}
