// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package mql

import (
	"fmt"

	"github.com/DistAlchemist/Mongongo/config"
)

type setSuperColumnMap struct {
	cfMetaData config.CFMetaData
	rowKey string 
	superColumnMapExpr []superMapPair
}

func (p setSuperColumnMap) execute() {
	fmt.Println(p.explainPlan())
	return 
}

func (p *setSuperColumnMap) explainPlan() string {
	res := fmt.Sprintf("%s Column Family: Batch SET a set of Super Columns: \n", p.cfMetaData.ColumnType) + 
		fmt.Sprintf("\tTable Name:     %s\n", p.cfMetaData.TableName) + 
		fmt.Sprintf("\tColumn Family:  %s\n", p.cfMetaData.CFName) + 
		fmt.Sprintf("\tRowKey:         %s\n", p.rowKey)
	for _, superPair := range p.superColumnMapExpr {
		superKey := superPair.key
		pairs := superPair.mapPair
		for _, pair := range pairs {
			columnKey := pair.key
			value := pair.value
			res += 
			fmt.Sprintf("\tSuperColumnKey: %s\n", superKey) + 
			fmt.Sprintf("\tColumnKey:      %s\n", columnKey) + 
			fmt.Sprintf("\tValue:          %s\n", value)
		}
	}
	return res 
}