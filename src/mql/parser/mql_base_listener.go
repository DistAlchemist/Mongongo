// Code generated from Mql.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Mql

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMqlListener is a complete listener for a parse tree produced by MqlParser.
type BaseMqlListener struct{}

var _ MqlListener = &BaseMqlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMqlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMqlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMqlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMqlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStringVal is called when production stringVal is entered.
func (s *BaseMqlListener) EnterStringVal(ctx *StringValContext) {}

// ExitStringVal is called when production stringVal is exited.
func (s *BaseMqlListener) ExitStringVal(ctx *StringValContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseMqlListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseMqlListener) ExitStmt(ctx *StmtContext) {}

// EnterGetStmt is called when production getStmt is entered.
func (s *BaseMqlListener) EnterGetStmt(ctx *GetStmtContext) {}

// ExitGetStmt is called when production getStmt is exited.
func (s *BaseMqlListener) ExitGetStmt(ctx *GetStmtContext) {}

// EnterSetStmt is called when production setStmt is entered.
func (s *BaseMqlListener) EnterSetStmt(ctx *SetStmtContext) {}

// ExitSetStmt is called when production setStmt is exited.
func (s *BaseMqlListener) ExitSetStmt(ctx *SetStmtContext) {}

// EnterColumnSpec is called when production columnSpec is entered.
func (s *BaseMqlListener) EnterColumnSpec(ctx *ColumnSpecContext) {}

// ExitColumnSpec is called when production columnSpec is exited.
func (s *BaseMqlListener) ExitColumnSpec(ctx *ColumnSpecContext) {}

// EnterTableName is called when production tableName is entered.
func (s *BaseMqlListener) EnterTableName(ctx *TableNameContext) {}

// ExitTableName is called when production tableName is exited.
func (s *BaseMqlListener) ExitTableName(ctx *TableNameContext) {}

// EnterColumnFamilyName is called when production columnFamilyName is entered.
func (s *BaseMqlListener) EnterColumnFamilyName(ctx *ColumnFamilyNameContext) {}

// ExitColumnFamilyName is called when production columnFamilyName is exited.
func (s *BaseMqlListener) ExitColumnFamilyName(ctx *ColumnFamilyNameContext) {}

// EnterValueExpr is called when production valueExpr is entered.
func (s *BaseMqlListener) EnterValueExpr(ctx *ValueExprContext) {}

// ExitValueExpr is called when production valueExpr is exited.
func (s *BaseMqlListener) ExitValueExpr(ctx *ValueExprContext) {}

// EnterCellValue is called when production cellValue is entered.
func (s *BaseMqlListener) EnterCellValue(ctx *CellValueContext) {}

// ExitCellValue is called when production cellValue is exited.
func (s *BaseMqlListener) ExitCellValue(ctx *CellValueContext) {}

// EnterColumnMapValue is called when production columnMapValue is entered.
func (s *BaseMqlListener) EnterColumnMapValue(ctx *ColumnMapValueContext) {}

// ExitColumnMapValue is called when production columnMapValue is exited.
func (s *BaseMqlListener) ExitColumnMapValue(ctx *ColumnMapValueContext) {}

// EnterSuperColumnMapValue is called when production superColumnMapValue is entered.
func (s *BaseMqlListener) EnterSuperColumnMapValue(ctx *SuperColumnMapValueContext) {}

// ExitSuperColumnMapValue is called when production superColumnMapValue is exited.
func (s *BaseMqlListener) ExitSuperColumnMapValue(ctx *SuperColumnMapValueContext) {}

// EnterColumnMapEntry is called when production columnMapEntry is entered.
func (s *BaseMqlListener) EnterColumnMapEntry(ctx *ColumnMapEntryContext) {}

// ExitColumnMapEntry is called when production columnMapEntry is exited.
func (s *BaseMqlListener) ExitColumnMapEntry(ctx *ColumnMapEntryContext) {}

// EnterSuperColumnMapEntry is called when production superColumnMapEntry is entered.
func (s *BaseMqlListener) EnterSuperColumnMapEntry(ctx *SuperColumnMapEntryContext) {}

// ExitSuperColumnMapEntry is called when production superColumnMapEntry is exited.
func (s *BaseMqlListener) ExitSuperColumnMapEntry(ctx *SuperColumnMapEntryContext) {}

// EnterColumnOrSuperColumnName is called when production columnOrSuperColumnName is entered.
func (s *BaseMqlListener) EnterColumnOrSuperColumnName(ctx *ColumnOrSuperColumnNameContext) {}

// ExitColumnOrSuperColumnName is called when production columnOrSuperColumnName is exited.
func (s *BaseMqlListener) ExitColumnOrSuperColumnName(ctx *ColumnOrSuperColumnNameContext) {}

// EnterRowKey is called when production rowKey is entered.
func (s *BaseMqlListener) EnterRowKey(ctx *RowKeyContext) {}

// ExitRowKey is called when production rowKey is exited.
func (s *BaseMqlListener) ExitRowKey(ctx *RowKeyContext) {}

// EnterColumnOrSuperColumnKey is called when production columnOrSuperColumnKey is entered.
func (s *BaseMqlListener) EnterColumnOrSuperColumnKey(ctx *ColumnOrSuperColumnKeyContext) {}

// ExitColumnOrSuperColumnKey is called when production columnOrSuperColumnKey is exited.
func (s *BaseMqlListener) ExitColumnOrSuperColumnKey(ctx *ColumnOrSuperColumnKeyContext) {}

// EnterColumnKey is called when production columnKey is entered.
func (s *BaseMqlListener) EnterColumnKey(ctx *ColumnKeyContext) {}

// ExitColumnKey is called when production columnKey is exited.
func (s *BaseMqlListener) ExitColumnKey(ctx *ColumnKeyContext) {}

// EnterSuperColumnKey is called when production superColumnKey is entered.
func (s *BaseMqlListener) EnterSuperColumnKey(ctx *SuperColumnKeyContext) {}

// ExitSuperColumnKey is called when production superColumnKey is exited.
func (s *BaseMqlListener) ExitSuperColumnKey(ctx *SuperColumnKeyContext) {}
