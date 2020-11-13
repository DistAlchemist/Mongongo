package mql

import (
	"fmt"
	"log"

	"github.com/DistAlchemist/Mongongo/mql/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type mqlListener struct {
	*parser.BaseMqlListener
	stmtType                int
	action                  string
	dimensionCnt            int
	tableName               string
	columnFamilyName        string
	rowKey                  string
	columnOrSuperColumnKeys []string
	cellValue               string
}

// // VisitTerminal is called when a terminal node is visited.
// func (l *mqlListener) VisitTerminal(node antlr.TerminalNode) {
// 	fmt.Println("visitTerminal")
// }

// // VisitErrorNode is called when an error node is visited.
// func (l *mqlListener) VisitErrorNode(node antlr.ErrorNode) {
// 	fmt.Println("visitErrorNode")
// }

// // EnterStringVal is called when production stringVal is entered.
// func (l *mqlListener) EnterStringVal(ctx *parser.StringValContext) {
// 	fmt.Println("enterstringVal")
// }

// func (l *mqlListener) EnterEveryRule(c antlr.ParserRuleContext) {
// 	fmt.Println("enterEveryRule")
// }

// // ExitStmt is called when production stmt is exited.
// func (l *mqlListener) ExitStmt(c *parser.StmtContext) {
// 	fmt.Println("exitStmt")
// }

func (l *mqlListener) EnterStmt(c *parser.StmtContext) {
	log.Println("entering EnterStmt")
}

func (l *mqlListener) EnterSetStmt(c *parser.SetStmtContext) {
	l.stmtType = parser.MqlParserRULE_setStmt
	l.action = "SET"
	log.Println("entering EnterSetStmt")
}

func (l *mqlListener) EnterGetStmt(c *parser.GetStmtContext) {
	l.stmtType = parser.MqlParserRULE_getStmt
	l.action = "GET"
	log.Println("entering EnterGetStmt")
}

func (l *mqlListener) EnterColumnSpec(c *parser.ColumnSpecContext) {
	fmt.Println("enter columnSpec")
	// skip tableName, columnFamilyName and rowKey
	l.dimensionCnt = c.GetChildCount() - 3
	fmt.Println("ColumnSpec childCount: %v", c.GetChildCount())
	fmt.Println("dimensionCnt: %v", l.dimensionCnt)

	// if l.dimensionCnt > 1 {
	// 	log.Fatal("currenty only support tableName.columnFamilyName[rowkey][column]=cellValue")
	// }
	// if l.dimensionCnt > 2 {
	// 	log.Fatal("too many dimensions")
	// }
}

func (l *mqlListener) ExitColumnSpec(c *parser.ColumnSpecContext) {
	fmt.Println("exit columnSpec")
	// skip tableName, columnFamilyName and rowKey
	l.dimensionCnt = c.GetChildCount() - 3
	fmt.Println("exit ColumnSpec childCount: %v", c.GetChildCount())
	fmt.Println("exit dimensionCnt: %v", l.dimensionCnt)

	// if l.dimensionCnt > 1 {
	// 	log.Fatal("currenty only support tableName.columnFamilyName[rowkey][column]=cellValue")
	// }
	// if l.dimensionCnt > 2 {
	// 	log.Fatal("too many dimensions")
	// }
}
func (l *mqlListener) EnterTableName(c *parser.TableNameContext) {
	fmt.Println("enter tableName")
	l.tableName = c.GetText()
}

func (l *mqlListener) EnterColumnFamilyName(c *parser.ColumnFamilyNameContext) {
	fmt.Println("enter columnFamilyName")
	l.columnFamilyName = c.GetText()
}

func (l *mqlListener) EnterRowKey(c *parser.RowKeyContext) {
	fmt.Println("enter rowKey")
	l.rowKey = c.GetText()
}

func (l *mqlListener) EnterColumnOrSuperColumnKey(c *parser.ColumnOrSuperColumnKeyContext) {
	fmt.Println("enter columnOrSuperColumnKey")
	l.columnOrSuperColumnKeys = append(l.columnOrSuperColumnKeys, c.GetText())
}

// func (l *mqlListener) EnterValueExpr(c *parser.ValueExprContext) {
// 	//
// }

func (l *mqlListener) EnterCellValue(c *parser.CellValueContext) {
	fmt.Println("enter cellValue")
	l.cellValue = c.GetText()
}

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
	fmt.Println("this is executeQuery function")
	fmt.Printf("%+v\n", listener)
	var res Result
	var action, setResult, getResult string
	res.ResultSet = make(map[string]string)
	if listener.stmtType == parser.MqlParserRULE_setStmt {
		setResult = executeSetUniqueKey(listener.tableName, listener.columnFamilyName, listener.rowKey,
			listener.columnOrSuperColumnKeys[0], listener.cellValue)
		action = fmt.Sprintf("%v %v.%v[%v][%v] = %v", listener.action, listener.tableName,
			listener.columnFamilyName, listener.rowKey, listener.columnOrSuperColumnKeys[0],
			listener.cellValue)
	} else if listener.stmtType == parser.MqlParserRULE_getStmt {
		getResult = executeGetUniqueKey(listener.tableName, listener.columnFamilyName, listener.rowKey,
			listener.columnOrSuperColumnKeys[0])
		action = fmt.Sprintf("%v %v.%v[%v][%v]", listener.action, listener.tableName,
			listener.columnFamilyName, listener.rowKey, listener.columnOrSuperColumnKeys[0])
	}
	res.ResultSet["action"] = listener.action + action
	res.ResultSet["setresult"] = setResult
	res.ResultSet["getresult"] = getResult
	return res
}
