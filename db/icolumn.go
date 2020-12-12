// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import (
	"os"
)

// IColumn provide interface for Column and SuperColumn
type IColumn interface {
	addColumn(column IColumn)
	getName() string
	GetName() string
	getSize() int32
	toByteArray() []byte
	serializedSize() uint32
	getObjectCount() int
	timestamp() int64
	GetTimestamp() int64
	putColumn(IColumn) bool
	getSubColumns() map[string]IColumn
	GetSubColumns() map[string]IColumn
	isMarkedForDelete() bool
	IsMarkedForDelete() bool
	getValue() []byte
	GetValue() []byte
	getMarkedForDeleteAt() int64
	getLocalDeletionTime() int
	mostRecentChangeAt() int64
}

// IColumnSerializer ...
type IColumnSerializer interface {
	serialize(column IColumn, dos *os.File)
	serializeB(column IColumn, dos []byte)
	deserialize(dis *os.File) IColumn
}
