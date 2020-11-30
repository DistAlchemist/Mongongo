// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package mql

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/DistAlchemist/Mongongo/mql/parser"
)

// Result embeds error message and results
type Result struct {
	ErrorCode int
	ErrorText string
	ResultSet map[string]string
}

// ExecuteQuery first compile query and execute it
func ExecuteQuery(query string) Result {
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
	plan := doSemanticAnalysis(queryTree.children[0])
	plan.execute()
	var res Result
	res.ErrorCode = 0
	// var action, setResult, getResult string
	// res.ResultSet = make(map[string]string)
	// if listener.stmtType == parser.MqlParserRULE_setStmt {
	// 	setResult = executeSetUniqueKey(listener.tableName, listener.columnFamilyName, listener.rowKey,
	// 		listener.columnOrSuperColumnKeys[0], listener.cellValue)
	// 	action = fmt.Sprintf("%v %v.%v[%v][%v] = %v", listener.action, listener.tableName,
	// 		listener.columnFamilyName, listener.rowKey, listener.columnOrSuperColumnKeys[0],
	// 		listener.cellValue)
	// } else if listener.stmtType == parser.MqlParserRULE_getStmt {
	// 	getResult = executeGetUniqueKey(listener.tableName, listener.columnFamilyName, listener.rowKey,
	// 		listener.columnOrSuperColumnKeys[0])
	// 	action = fmt.Sprintf("%v %v.%v[%v][%v]", listener.action, listener.tableName,
	// 		listener.columnFamilyName, listener.rowKey, listener.columnOrSuperColumnKeys[0])
	// }
	// res.ResultSet["action"] = listener.action + action
	// res.ResultSet["setresult"] = setResult
	// res.ResultSet["getresult"] = getResult

	return res
}
