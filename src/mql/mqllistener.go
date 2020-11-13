package mql

import (
	"log"

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
}

func (l *mqlListener) EnterSetStmt(c *parser.SetStmtContext) {
	l.stmtType = parser.MqlParserRULE_setStmt
	l.action = "SET"
	log.Println("entering EnterSetStmt")
}

func (l *mqlListener) EnterGetStmt(c *parser.SetStmtContext) {
	l.stmtType = parser.MqlParserRULE_getStmt
	l.action = "GET"
	log.Println("entering EnterGetStmt")
}

func (l *mqlListener) EnterColumnSpec(c *parser.ColumnSpecContext) {
	// skip tableName, columnFamilyName and rowKey
	l.dimensionCnt = c.GetChildCount() - 3

	if l.dimensionCnt > 1 {
		log.Fatal("currenty only support tableName.columnFamilyName[rowkey][column]=cellValue")
	}
	// if l.dimensionCnt > 2 {
	// 	log.Fatal("too many dimensions")
	// }
}

func (l *mqlListener) EnterTableName(c *parser.TableNameContext) {
	l.tableName = c.GetText()
}

func (l *mqlListener) EnterColumnFamilyName(c *parser.ColumnFamilyNameContext) {
	l.columnFamilyName = c.GetText()
}

func (l *mqlListener) EnterRowKey(c *parser.RowKeyContext) {
	l.rowKey = c.GetText()
}

func (l *mqlListener) EnterColumnOrSuperColumnKey(c *parser.ColumnOrSuperColumnKeyContext) {
	l.columnOrSuperColumnKeys = append(l.columnOrSuperColumnKeys, c.GetText())
}

func (l *mqlListener) EnterValueExpr(c *parser.ValueExprContext) {
	//
}

func (l *mqlListener) EnterCellValue(c *parser.CellValueContext) {
	l.cellValue = c.GetText()
}
