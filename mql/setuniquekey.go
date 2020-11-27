// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package mql

import (
	"fmt"
	"time"

	"github.com/DistAlchemist/Mongongo/config"
	"github.com/DistAlchemist/Mongongo/db"
	"github.com/DistAlchemist/Mongongo/service"
)

type setUniqueKey struct {
	cfMetaData     config.CFMetaData
	rowKey         string
	columnKey      string
	superColumnKey string
	value          string
}

func (p setUniqueKey) execute() {
	fmt.Println(p.explainPlan())
	columnKey := p.columnKey
	var columnFamilyColumn string
	if p.superColumnKey != "" {
		superColumnKey := p.superColumnKey
		columnFamilyColumn = p.cfMetaData.CFName + ":" + superColumnKey + ":" +
			columnKey
	} else {
		columnFamilyColumn = p.cfMetaData.CFName + ":" + columnKey
	}
	rm := db.NewRowMutation(p.cfMetaData.TableName, p.rowKey)
	rm.Add(columnFamilyColumn, p.value, makeTimestamp())
	service.Insert(rm)
	return
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func (p *setUniqueKey) explainPlan() string {
	res := fmt.Sprintf("%s Column Family: Unique Key SET: \n", p.cfMetaData.ColumnType) +
		fmt.Sprintf("\tTable Name:     %s\n", p.cfMetaData.TableName) +
		fmt.Sprintf("\tColumn Family:  %s\n", p.cfMetaData.CFName) +
		fmt.Sprintf("\tRowKey:         %s\n", p.rowKey)
	if p.superColumnKey != "" {
		res +=
			fmt.Sprintf("\tSuperColumnKey: %s\n", p.superColumnKey)
	}
	res +=
		fmt.Sprintf("\tColumnKey:      %s\n", p.columnKey) +
			fmt.Sprintf("\tValue:          %s\n", p.value)
	return res
}

// func executeSetUniqueKey(tableName, columnFamilyName, rowKey, columnKey, value string) string {
// 	columnFamilyColumn := columnFamilyName + ":" + columnKey
// 	rm := db.RowMutation{tableName, rowKey, nil}
// 	rm.Add(columnFamilyColumn, value, time.Now().UnixNano()/int64(time.Millisecond))
// 	return db.Insert(rm)
// }
