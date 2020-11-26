package mql

import (
	"log"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/DistAlchemist/Mongongo/mql/parser"
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
	root                    *node
	curNode                 *node
}

type node struct {
	id       int
	text     string
	parent   *node
	children []*node
}

func (n *node) addChild(id int, text string) *node {
	child := &node{id: id, text: text, parent: n}
	n.children = append(n.children, child)
	return child
}

func (n *node) dfs() {
	log.Printf("enter id: %2v, text: %v\n", n.id, n.text)
	for _, c := range n.children {
		c.dfs()
	}
	log.Printf("exit  id: %2v, text: %v\n", n.id, n.text)
}

func (l *mqlListener) init() {
	l.root = &node{-1, "", nil, nil}
	l.curNode = l.root
}

// EnterEveryRule is called when any rule is entered.
func (l *mqlListener) EnterEveryRule(c antlr.ParserRuleContext) {
	log.Printf("Enter: GetRuleIndex: %v, ", c.GetRuleIndex())
	log.Printf("GetText %v\n", c.GetText())
	l.curNode = l.curNode.addChild(c.GetRuleIndex(), c.GetText())
}

// ExitEveryRule is called when any rule is exited.
func (l *mqlListener) ExitEveryRule(c antlr.ParserRuleContext) {
	log.Printf("Enter: GetRuleIndex: %v, ", c.GetRuleIndex())
	log.Printf("GetText %v\n", c.GetText())
	l.curNode = l.curNode.parent
}

// func (l *mqlListener) EnterStmt(c *parser.StmtContext) {
// 	log.Println("entering EnterStmt")
// }

// func (l *mqlListener) EnterSetStmt(c *parser.SetStmtContext) {
// 	l.stmtType = parser.MqlParserRULE_setStmt
// 	l.action = "SET"
// 	log.Println("entering EnterSetStmt")
// }

// func (l *mqlListener) EnterGetStmt(c *parser.GetStmtContext) {
// 	l.stmtType = parser.MqlParserRULE_getStmt
// 	l.action = "GET"
// 	log.Println("entering EnterGetStmt")
// }

// func (l *mqlListener) EnterColumnSpec(c *parser.ColumnSpecContext) {
// 	fmt.Println("enter columnSpec")
// 	// skip tableName, columnFamilyName and rowKey
// 	l.dimensionCnt = c.GetChildCount() - 3
// 	fmt.Println("ColumnSpec childCount: %v", c.GetChildCount())
// 	fmt.Println("dimensionCnt: %v", l.dimensionCnt)

// 	// if l.dimensionCnt > 1 {
// 	// 	log.Fatal("currenty only support tableName.columnFamilyName[rowkey][column]=cellValue")
// 	// }
// 	// if l.dimensionCnt > 2 {
// 	// 	log.Fatal("too many dimensions")
// 	// }
// }

// func (l *mqlListener) ExitColumnSpec(c *parser.ColumnSpecContext) {
// 	fmt.Println("exit columnSpec")
// 	// skip tableName, columnFamilyName and rowKey
// 	l.dimensionCnt = c.GetChildCount() - 3
// 	fmt.Println("exit ColumnSpec childCount: %v", c.GetChildCount())
// 	fmt.Println("exit dimensionCnt: %v", l.dimensionCnt)

// 	// if l.dimensionCnt > 1 {
// 	// 	log.Fatal("currenty only support tableName.columnFamilyName[rowkey][column]=cellValue")
// 	// }
// 	// if l.dimensionCnt > 2 {
// 	// 	log.Fatal("too many dimensions")
// 	// }
// }
// func (l *mqlListener) EnterTableName(c *parser.TableNameContext) {
// 	fmt.Println("enter tableName")
// 	l.tableName = c.GetText()
// }

// func (l *mqlListener) EnterColumnFamilyName(c *parser.ColumnFamilyNameContext) {
// 	fmt.Println("enter columnFamilyName")
// 	l.columnFamilyName = c.GetText()
// }

// func (l *mqlListener) EnterRowKey(c *parser.RowKeyContext) {
// 	fmt.Println("enter rowKey")
// 	l.rowKey = c.GetText()
// }

// func (l *mqlListener) EnterColumnOrSuperColumnKey(c *parser.ColumnOrSuperColumnKeyContext) {
// 	fmt.Println("enter columnOrSuperColumnKey")
// 	l.columnOrSuperColumnKeys = append(l.columnOrSuperColumnKeys, c.GetText())
// }

// // func (l *mqlListener) EnterValueExpr(c *parser.ValueExprContext) {
// // 	//
// // }

// func (l *mqlListener) EnterCellValue(c *parser.CellValueContext) {
// 	fmt.Println("enter cellValue")
// 	l.cellValue = c.GetText()
// }
