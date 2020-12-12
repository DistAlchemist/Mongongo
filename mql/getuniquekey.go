// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package mql

import (
	"fmt"

	"github.com/DistAlchemist/Mongongo/config"
)

type getUniqueKey struct {
	cfMetaData     config.CFMetaData
	rowKey         string
	superColumnKey string
	columnKey      string
}

func (p getUniqueKey) execute() {
	fmt.Println(p.explainPlan())
}

func executeGetUniqueKey(tableName, columnFamilyName, rowKey, columnKeye string) string {
	//
	return ""
}

func (p *getUniqueKey) explainPlan() string {
	res := fmt.Sprintf("%s Column Family: Unique Key GET: \n", p.cfMetaData.ColumnType) +
		fmt.Sprintf("\tTable Name:     %s\n", p.cfMetaData.TableName) +
		fmt.Sprintf("\tColumn Family:  %s\n", p.cfMetaData.CFName) +
		fmt.Sprintf("\tRowKey:         %s\n", p.rowKey)
	if p.superColumnKey != "" {
		res +=
			fmt.Sprintf("\tSuperColumnKey: %s\n", p.superColumnKey)
	}
	res +=
		fmt.Sprintf("\tColumnKey:      %s\n", p.columnKey)
	return res
}
