// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

import "log"

// Column stores name and value etc.
type Column struct {
	Name      string
	Value     string
	Timestamp int64
}

func (c Column) addColumn(name string, column IColumn) {
	log.Printf("Invalid method: Column doesn't support addColumn\n")
}

// NewColumn constructs a Column
func NewColumn(name, value string, timestamp int64) Column {
	c := Column{}
	c.Name = name
	c.Value = value
	c.Timestamp = timestamp
	return c
}
