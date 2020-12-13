// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package mql

import (
	"fmt"
	"log"
	"net/rpc"
	"time"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/davecgh/go-spew/spew"

	"github.com/DistAlchemist/Mongongo/mql/parser"
	"github.com/DistAlchemist/Mongongo/service"
)

var cc *rpc.Client

// Result embeds error message and results
type Result struct {
	ErrorCode int
	ErrorText string
	ResultSet map[string]string
}

// ExecuteQuery first compile query and execute it
func ExecuteQuery(c *rpc.Client, query string) Result {
	cc = c // somewhat ugly workaround
	// setup the input
	is := antlr.NewInputStream(query)
	// create the lexer
	lexer := parser.NewMqlLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// create the parser
	p := parser.NewMqlParser(stream)
	// finally parse the expression (by walking the tree)
	var listener mqlListener
	listener.init()
	// during the Walk, we build the abstract syntax tree
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Stmt())

	// do semantic phase
	queryTree := listener.root.children[0] // root -> stmt -> setStmt/getStmt...
	executeCLIStmt(queryTree.children[0])  // stmt -> setStmt/getStmt
	// plan := doSemanticAnalysis(queryTree.children[0])
	// plan.execute()
	var res Result
	res.ErrorCode = 0
	return res
}

func executeCLIStmt(ast *node) {
	switch ast.id {
	case parser.MqlParserRULE_setStmt:
		executeSet(ast)
	case parser.MqlParserRULE_getStmt:
		executeGet(ast)
	default:
		log.Printf("Invalid statement\n")
	}
}

func executeSet(ast *node) {
	// execute set statement
	childCount := len(ast.children)
	if childCount != 2 {
		log.Printf("should be set columnSpec = valueExpr\n")
	}
	columnFamilySpec := ast.children[0]
	tableName := getTableName(columnFamilySpec)
	key := getKey(columnFamilySpec)
	columnFamily := getColumnFamily(columnFamilySpec)
	columnSpecCnt := numColumnSpecifiers(columnFamilySpec)
	// setStmt.valueExpr.cellValue.stringVal.text
	value := ast.children[1].children[0].children[0].text
	// assume simple columnFamily for now
	if columnSpecCnt == 1 {
		// set table.cf['key']['column'] = 'value'
		// get the column name
		columnName := getColumn(columnFamilySpec, 0)
		// do the insert
		// service.InsertN(tableName, key, service.NewColumnPath(columnFamily))
		args := service.InsertArgs{}
		reply := service.InsertReply{}
		args.Table = tableName
		args.Key = key
		args.CPath = service.NewColumnPath(columnFamily, nil, []byte(columnName))
		args.Value = []byte(value)
		args.Timestamp = currentTimeMillis()
		args.ConsistencyLevel = 0
		err := cc.Call("Mongongo.Insert", &args, &reply)
		if err != nil {
			log.Fatal("calling:", err)
		}
		log.Printf("reply.result: %+v\n", reply.Result)
	} else {
		log.Printf("currently only support set table.cf['key']['column']='value'\n")
	}
}

func currentTimeMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// func getColumn(ast *node, idx int) string {
// 	return ast.children[idx+3].children[0].text
// }

func numColumnSpecifiers(ast *node) int {
	// skip table, column family and rowkey
	return len(ast.children) - 3
}

func getColumnFamily(ast *node) string {
	// columnSpec.columnFamilyName.text
	return ast.children[1].text
}

func getKey(ast *node) string {
	// columnSpec.rowKey.stringVal.text
	return ast.children[2].children[0].text
}

func getTableName(ast *node) string {
	// columnSpec.tableName.text
	return ast.children[0].text
}

func executeGet(ast *node) {
	// execute get statement
	// getStmt.columnSpec
	childCount := len(ast.children)
	if childCount != 1 {
		log.Fatalf("GET should only has columnSpec\n")
	}
	columnFamilySpec := ast.children[0]
	tableName := getTableName(columnFamilySpec)
	key := getKey(columnFamilySpec)
	columnFamily := getColumnFamily(columnFamilySpec)
	columnSpecCnt := numColumnSpecifiers(columnFamilySpec)
	// assume simple columnFamily for now
	if columnSpecCnt == 0 {
		// get table.cf['key']
		srange := service.NewSliceRange(nil, nil, true, 1000000)
		args := service.GetSliceArgs{}
		args.Keyspace = tableName
		args.Key = key
		args.ColumnParent = service.NewColumnParent(columnFamily, nil)
		args.Predicate = service.NewSlicePredicate(nil, srange)
		args.ConsistencyLevel = 1
		reply := service.GetSliceReply{}
		err := cc.Call("Mongongo.GetSlice", &args, &reply)
		if err != nil {
			log.Fatal("calling:", err)
		}
		columns := reply.Columns
		size := len(columns)
		for _, cosc := range columns {
			column := cosc.Column
			fmt.Printf("column=%v, value=%v, timestamp=%v\n", column.Name,
				column.Value, column.Timestamp)
		}
		fmt.Printf("returned %v rows.\n", size)
	} else {
		// get table.cf['key']['column']
		columnName := getColumn(columnFamilySpec, 0)
		args := service.GetArgs{}
		args.Keyspace = tableName
		args.Key = key
		args.ColumnPath = service.NewColumnPath(columnFamily, nil, []byte(columnName))
		args.ConsistencyLevel = 1
		reply := service.GetReply{}
		err := cc.Call("Mongongo.Get", &args, &reply)
		if err != nil {
			log.Fatal("calling:", err)
		}
		column := reply.Cosc.Column
		spew.Printf("get column: %#+v\n\n", column)
		if column == nil {
			return
		}
		fmt.Printf("name=%v, value=%v, timestamp=%v\n", column.Name,
			column.Value, column.Timestamp)
	}
}
