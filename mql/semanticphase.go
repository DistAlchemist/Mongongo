// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package mql

import (
	"log"

	"github.com/DistAlchemist/Mongongo/config"
	"github.com/DistAlchemist/Mongongo/mql/parser"
)

type mapPair struct {
	key   string
	value string
}

type superMapPair struct {
	key     string
	mapPair []mapPair
}

func doSemanticAnalysis(ast *node) Plan {
	var plan Plan
	switch ast.id {
	case parser.MqlParserRULE_setStmt:
		plan = compileSet(ast)
	case parser.MqlParserRULE_getStmt:
		plan = compileGet(ast)
	default:
		log.Printf("Unsupported stmt type: %v\n", ast.id)
	}
	return plan
}

func getSimpleExpr(ast *node) string {
	if ast.id != parser.MqlParserRULE_cellValue &&
		ast.id != parser.MqlParserRULE_columnKey &&
		ast.id != parser.MqlParserRULE_rowKey &&
		ast.id != parser.MqlParserRULE_superColumnKey &&
		ast.id != parser.MqlParserRULE_columnOrSuperColumnKey {
		log.Printf("Invalid type id: %v\n", ast.id)
	}
	return ast.children[0].text
}

func getColumnMapExpr(ast *node) []mapPair {
	if ast.id != parser.MqlParserRULE_columnMapValue {
		log.Printf("Invalid type id: %v, should be columnMapValue id: %v\n", ast.id,
			parser.MqlParserRULE_columnMapValue)
	}
	res := make([]mapPair, 0)
	for _, entryNode := range ast.children {
		columnKey := getSimpleExpr(entryNode.children[0])
		columnValue := getSimpleExpr(entryNode.children[1])
		res = append(res, mapPair{columnKey, columnValue})
	}
	return res
}

func getSuperColumnMapExpr(ast *node) []superMapPair {
	if ast.id != parser.MqlParserRULE_superColumnMapValue {
		log.Printf("Invalid type id: %v, should be superColumnMapValue id: %v\n", ast.id,
			parser.MqlParserRULE_columnMapValue)
	}
	res := make([]superMapPair, 0)
	for _, entryNode := range ast.children {
		superColumnKey := getSimpleExpr(entryNode.children[0])
		columnMapExpr := getColumnMapExpr(entryNode.children[1])
		entry := superMapPair{superColumnKey, columnMapExpr}
		res = append(res, entry)
	}
	return res
}

func getColumn(ast *node, pos int) string {
	return ast.children[pos+3].children[0].text
}

func compileSet(ast *node) Plan {
	columnSpec := ast.children[0]
	cfMetaData := getColumnFamilyInfo(columnSpec)
	rowKey := columnSpec.children[2].children[0].text
	// skip over tableName, columnFamily and rowKey
	dimensions := len(columnSpec.children) - 3
	valueNode := ast.children[1]
	var plan Plan
	if cfMetaData.ColumnType == "Super" { // Super column family
		if dimensions == 2 {
			// set table.superCF['rowKey']['superColumnKey']['columnKey']='value'
			value := getSimpleExpr(valueNode.children[0])
			superColumnKey := getColumn(columnSpec, 0)
			columnKey := getColumn(columnSpec, 1)
			plan = setUniqueKey{cfMetaData, rowKey, superColumnKey, columnKey, value}
		} else if dimensions == 1 {
			// set table.superCF['rowKey']['superColumnKey']={'columnKey'=>'value',...}
			columnMapExpr := getColumnMapExpr(valueNode.children[0])
			superColumnKey := getColumn(columnSpec, 0)
			plan = setColumnMap{cfMetaData, rowKey, superColumnKey, columnMapExpr}
		} else {
			// set table.superCF['rowKey'] = {'superColumnKey'=>{columnMapValue},...}
			if dimensions != 0 {
				log.Printf("Invalid dimension: %v\n", dimensions)
			}
			superColumnMapExpr := getSuperColumnMapExpr(valueNode.children[0])
			plan = setSuperColumnMap{cfMetaData, rowKey, superColumnMapExpr}
		}
	} else { // Standard column family
		if dimensions == 1 {
			// set table.standardCF['key']['column']='value'
			value := getSimpleExpr(valueNode.children[0])
			columnKey := getColumn(columnSpec, 0)
			plan = setUniqueKey{cfMetaData, rowKey, "", columnKey, value}
		} else {
			// set table.standardCF['key']={'columnKey'=>'value',...}
			if dimensions != 0 {
				log.Printf("Invalid dimensions: %v\n", dimensions)
			}
			columnMapExpr := getColumnMapExpr(valueNode.children[0])
			plan = setColumnMap{cfMetaData, rowKey, "", columnMapExpr}
		}
	}
	return plan
}

func getColumnFamilyInfo(ast *node) config.CFMetaData {
	tableNode := ast.children[0]
	columnFamilyNode := ast.children[1]
	tableName := tableNode.text
	columnFamilyName := columnFamilyNode.text
	columnFamilies := config.GetTableMetaData(tableName)
	if columnFamilies == nil {
		log.Printf("Invalid table: %v\n", tableName)
	}
	cfMetaData, ok := columnFamilies[columnFamilyName]
	if !ok {
		log.Printf("Invalid column family: %v.%v\n", tableName, columnFamilyName)
	}
	// skip over tablName, columnFamily and rowKey
	dimensions := len(ast.children) - 3
	if (cfMetaData.ColumnType == "Super" && dimensions > 2) || (cfMetaData.ColumnType == "Standard" && dimensions > 1) {
		log.Printf("Too many dimensions: %v for %v\n", dimensions, cfMetaData.ColumnType)
	}
	return cfMetaData
}

func compileGet(ast *node) Plan {
	plan := setUniqueKey{}
	return plan
}
