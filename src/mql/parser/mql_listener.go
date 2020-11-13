// Code generated from Mql.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Mql

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MqlListener is a complete listener for a parse tree produced by MqlParser.
type MqlListener interface {
	antlr.ParseTreeListener

	// EnterStringVal is called when entering the stringVal production.
	EnterStringVal(c *StringValContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterGetStmt is called when entering the getStmt production.
	EnterGetStmt(c *GetStmtContext)

	// EnterSetStmt is called when entering the setStmt production.
	EnterSetStmt(c *SetStmtContext)

	// EnterColumnSpec is called when entering the columnSpec production.
	EnterColumnSpec(c *ColumnSpecContext)

	// EnterTableName is called when entering the tableName production.
	EnterTableName(c *TableNameContext)

	// EnterColumnFamilyName is called when entering the columnFamilyName production.
	EnterColumnFamilyName(c *ColumnFamilyNameContext)

	// EnterValueExpr is called when entering the valueExpr production.
	EnterValueExpr(c *ValueExprContext)

	// EnterCellValue is called when entering the cellValue production.
	EnterCellValue(c *CellValueContext)

	// EnterColumnMapValue is called when entering the columnMapValue production.
	EnterColumnMapValue(c *ColumnMapValueContext)

	// EnterSuperColumnMapValue is called when entering the superColumnMapValue production.
	EnterSuperColumnMapValue(c *SuperColumnMapValueContext)

	// EnterColumnMapEntry is called when entering the columnMapEntry production.
	EnterColumnMapEntry(c *ColumnMapEntryContext)

	// EnterSuperColumnMapEntry is called when entering the superColumnMapEntry production.
	EnterSuperColumnMapEntry(c *SuperColumnMapEntryContext)

	// EnterColumnOrSuperColumnName is called when entering the columnOrSuperColumnName production.
	EnterColumnOrSuperColumnName(c *ColumnOrSuperColumnNameContext)

	// EnterRowKey is called when entering the rowKey production.
	EnterRowKey(c *RowKeyContext)

	// EnterColumnOrSuperColumnKey is called when entering the columnOrSuperColumnKey production.
	EnterColumnOrSuperColumnKey(c *ColumnOrSuperColumnKeyContext)

	// EnterColumnKey is called when entering the columnKey production.
	EnterColumnKey(c *ColumnKeyContext)

	// EnterSuperColumnKey is called when entering the superColumnKey production.
	EnterSuperColumnKey(c *SuperColumnKeyContext)

	// ExitStringVal is called when exiting the stringVal production.
	ExitStringVal(c *StringValContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitGetStmt is called when exiting the getStmt production.
	ExitGetStmt(c *GetStmtContext)

	// ExitSetStmt is called when exiting the setStmt production.
	ExitSetStmt(c *SetStmtContext)

	// ExitColumnSpec is called when exiting the columnSpec production.
	ExitColumnSpec(c *ColumnSpecContext)

	// ExitTableName is called when exiting the tableName production.
	ExitTableName(c *TableNameContext)

	// ExitColumnFamilyName is called when exiting the columnFamilyName production.
	ExitColumnFamilyName(c *ColumnFamilyNameContext)

	// ExitValueExpr is called when exiting the valueExpr production.
	ExitValueExpr(c *ValueExprContext)

	// ExitCellValue is called when exiting the cellValue production.
	ExitCellValue(c *CellValueContext)

	// ExitColumnMapValue is called when exiting the columnMapValue production.
	ExitColumnMapValue(c *ColumnMapValueContext)

	// ExitSuperColumnMapValue is called when exiting the superColumnMapValue production.
	ExitSuperColumnMapValue(c *SuperColumnMapValueContext)

	// ExitColumnMapEntry is called when exiting the columnMapEntry production.
	ExitColumnMapEntry(c *ColumnMapEntryContext)

	// ExitSuperColumnMapEntry is called when exiting the superColumnMapEntry production.
	ExitSuperColumnMapEntry(c *SuperColumnMapEntryContext)

	// ExitColumnOrSuperColumnName is called when exiting the columnOrSuperColumnName production.
	ExitColumnOrSuperColumnName(c *ColumnOrSuperColumnNameContext)

	// ExitRowKey is called when exiting the rowKey production.
	ExitRowKey(c *RowKeyContext)

	// ExitColumnOrSuperColumnKey is called when exiting the columnOrSuperColumnKey production.
	ExitColumnOrSuperColumnKey(c *ColumnOrSuperColumnKeyContext)

	// ExitColumnKey is called when exiting the columnKey production.
	ExitColumnKey(c *ColumnKeyContext)

	// ExitSuperColumnKey is called when exiting the superColumnKey production.
	ExitSuperColumnKey(c *SuperColumnKeyContext)
}
