package mql

import (
	"fmt"

	"github.com/DistAlchemist/Mongongo/mql/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
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
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Stmt())
	var res Result
	var action string
	res.ResultSet = make(map[string]string)
	if listener.stmtType == parser.MqlParserRULE_setStmt {
		action = fmt.Sprintf("%v %v.%v[%q][%q] = %q", listener.action, listener.tableName,
			listener.columnFamilyName, listener.rowKey, listener.columnOrSuperColumnKeys[0],
			listener.cellValue)
	} else if listener.stmtType == parser.MqlParserRULE_getStmt {
		action = fmt.Sprintf("%v %v.%v[%q][%q]", listener.action, listener.tableName,
			listener.columnFamilyName, listener.rowKey, listener.columnOrSuperColumnKeys[0])
	}
	res.ResultSet["action"] = action
	return res
}
