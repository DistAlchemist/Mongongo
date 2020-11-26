// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package mql

import (
	"fmt"

	"github.com/DistAlchemist/Mongongo/config"
)

type setColumnMap struct {
	cfMetaData config.CFMetaData 
	rowKey string 
	superColumnKey string 
	columnMapExpr []mapPair
}

func (p setColumnMap) execute() {
	fmt.Println(p.explainPlan())
	return 
}

func (p *setColumnMap) explainPlan() string {
	res := fmt.Sprintf("%s Column Family: Batch SET a set of Columns: \n", p.cfMetaData.ColumnType) + 
		fmt.Sprintf("\tTable Name:     %s\n", p.cfMetaData.TableName) + 
		fmt.Sprintf("\tColumn Family:  %s\n", p.cfMetaData.CFName) + 
		fmt.Sprintf("\tRowKey:         %s\n", p.rowKey)
	if p.superColumnKey != "" {
		res += 
		fmt.Sprintf("\tSuperColumnKey: %s\n", p.superColumnKey)
	}
	for _, pair := range p.columnMapExpr {
		columnKey := pair.key
		value := pair.value
		res += 
		fmt.Sprintf("\tColumnKey:      %s\n", columnKey) + 
		fmt.Sprintf("\tValue:          %s\n", value)
	}
	return res 
}