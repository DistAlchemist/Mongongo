// Code generated from Mql.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Mql

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 18, 121,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 3, 2, 3, 2, 3, 3, 3, 3, 5, 3, 43, 10, 3, 3, 4, 3, 4,
	3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 66, 10, 6, 5, 6, 68, 10,
	6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 5, 9, 77, 10, 9, 3, 10, 3,
	10, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 85, 10, 11, 12, 11, 14, 11, 88,
	11, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 96, 10, 12, 12,
	12, 14, 12, 99, 11, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3,
	18, 3, 19, 3, 19, 3, 19, 2, 2, 20, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 30, 32, 34, 36, 2, 3, 4, 2, 3, 3, 17, 17, 2, 109, 2, 38,
	3, 2, 2, 2, 4, 42, 3, 2, 2, 2, 6, 44, 3, 2, 2, 2, 8, 47, 3, 2, 2, 2, 10,
	52, 3, 2, 2, 2, 12, 69, 3, 2, 2, 2, 14, 71, 3, 2, 2, 2, 16, 76, 3, 2, 2,
	2, 18, 78, 3, 2, 2, 2, 20, 80, 3, 2, 2, 2, 22, 91, 3, 2, 2, 2, 24, 102,
	3, 2, 2, 2, 26, 106, 3, 2, 2, 2, 28, 110, 3, 2, 2, 2, 30, 112, 3, 2, 2,
	2, 32, 114, 3, 2, 2, 2, 34, 116, 3, 2, 2, 2, 36, 118, 3, 2, 2, 2, 38, 39,
	9, 2, 2, 2, 39, 3, 3, 2, 2, 2, 40, 43, 5, 6, 4, 2, 41, 43, 5, 8, 5, 2,
	42, 40, 3, 2, 2, 2, 42, 41, 3, 2, 2, 2, 43, 5, 3, 2, 2, 2, 44, 45, 7, 8,
	2, 2, 45, 46, 5, 10, 6, 2, 46, 7, 3, 2, 2, 2, 47, 48, 7, 9, 2, 2, 48, 49,
	5, 10, 6, 2, 49, 50, 7, 4, 2, 2, 50, 51, 5, 16, 9, 2, 51, 9, 3, 2, 2, 2,
	52, 53, 5, 12, 7, 2, 53, 54, 7, 5, 2, 2, 54, 55, 5, 14, 8, 2, 55, 56, 7,
	6, 2, 2, 56, 57, 5, 30, 16, 2, 57, 67, 7, 7, 2, 2, 58, 59, 7, 6, 2, 2,
	59, 60, 5, 32, 17, 2, 60, 65, 7, 7, 2, 2, 61, 62, 7, 6, 2, 2, 62, 63, 5,
	32, 17, 2, 63, 64, 7, 7, 2, 2, 64, 66, 3, 2, 2, 2, 65, 61, 3, 2, 2, 2,
	65, 66, 3, 2, 2, 2, 66, 68, 3, 2, 2, 2, 67, 58, 3, 2, 2, 2, 67, 68, 3,
	2, 2, 2, 68, 11, 3, 2, 2, 2, 69, 70, 7, 16, 2, 2, 70, 13, 3, 2, 2, 2, 71,
	72, 7, 16, 2, 2, 72, 15, 3, 2, 2, 2, 73, 77, 5, 18, 10, 2, 74, 77, 5, 20,
	11, 2, 75, 77, 5, 22, 12, 2, 76, 73, 3, 2, 2, 2, 76, 74, 3, 2, 2, 2, 76,
	75, 3, 2, 2, 2, 77, 17, 3, 2, 2, 2, 78, 79, 5, 2, 2, 2, 79, 19, 3, 2, 2,
	2, 80, 81, 7, 13, 2, 2, 81, 86, 5, 24, 13, 2, 82, 83, 7, 12, 2, 2, 83,
	85, 5, 24, 13, 2, 84, 82, 3, 2, 2, 2, 85, 88, 3, 2, 2, 2, 86, 84, 3, 2,
	2, 2, 86, 87, 3, 2, 2, 2, 87, 89, 3, 2, 2, 2, 88, 86, 3, 2, 2, 2, 89, 90,
	7, 14, 2, 2, 90, 21, 3, 2, 2, 2, 91, 92, 7, 13, 2, 2, 92, 97, 5, 26, 14,
	2, 93, 94, 7, 12, 2, 2, 94, 96, 5, 26, 14, 2, 95, 93, 3, 2, 2, 2, 96, 99,
	3, 2, 2, 2, 97, 95, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 100, 3, 2, 2, 2,
	99, 97, 3, 2, 2, 2, 100, 101, 7, 14, 2, 2, 101, 23, 3, 2, 2, 2, 102, 103,
	5, 34, 18, 2, 103, 104, 7, 11, 2, 2, 104, 105, 5, 18, 10, 2, 105, 25, 3,
	2, 2, 2, 106, 107, 5, 36, 19, 2, 107, 108, 7, 11, 2, 2, 108, 109, 5, 20,
	11, 2, 109, 27, 3, 2, 2, 2, 110, 111, 7, 16, 2, 2, 111, 29, 3, 2, 2, 2,
	112, 113, 5, 2, 2, 2, 113, 31, 3, 2, 2, 2, 114, 115, 5, 2, 2, 2, 115, 33,
	3, 2, 2, 2, 116, 117, 5, 2, 2, 2, 117, 35, 3, 2, 2, 2, 118, 119, 5, 2,
	2, 2, 119, 37, 3, 2, 2, 2, 8, 42, 65, 67, 76, 86, 97,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'?'", "'='", "'.'", "'['", "']'", "'GET'", "'SET'", "", "'=>'", "','",
	"'{'", "'}'", "';'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "GET", "SET", "WHITESPACE", "ASSOC", "COMMA", "LEFT_BRACE",
	"RIGHT_BRACE", "SEMICOLON", "Identifier", "StringLiteral", "IntegerLiteral",
}

var ruleNames = []string{
	"stringVal", "stmt", "getStmt", "setStmt", "columnSpec", "tableName", "columnFamilyName",
	"valueExpr", "cellValue", "columnMapValue", "superColumnMapValue", "columnMapEntry",
	"superColumnMapEntry", "columnOrSuperColumnName", "rowKey", "columnOrSuperColumnKey",
	"columnKey", "superColumnKey",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type MqlParser struct {
	*antlr.BaseParser
}

func NewMqlParser(input antlr.TokenStream) *MqlParser {
	this := new(MqlParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Mql.g4"

	return this
}

// MqlParser tokens.
const (
	MqlParserEOF            = antlr.TokenEOF
	MqlParserT__0           = 1
	MqlParserT__1           = 2
	MqlParserT__2           = 3
	MqlParserT__3           = 4
	MqlParserT__4           = 5
	MqlParserGET            = 6
	MqlParserSET            = 7
	MqlParserWHITESPACE     = 8
	MqlParserASSOC          = 9
	MqlParserCOMMA          = 10
	MqlParserLEFT_BRACE     = 11
	MqlParserRIGHT_BRACE    = 12
	MqlParserSEMICOLON      = 13
	MqlParserIdentifier     = 14
	MqlParserStringLiteral  = 15
	MqlParserIntegerLiteral = 16
)

// MqlParser rules.
const (
	MqlParserRULE_stringVal               = 0
	MqlParserRULE_stmt                    = 1
	MqlParserRULE_getStmt                 = 2
	MqlParserRULE_setStmt                 = 3
	MqlParserRULE_columnSpec              = 4
	MqlParserRULE_tableName               = 5
	MqlParserRULE_columnFamilyName        = 6
	MqlParserRULE_valueExpr               = 7
	MqlParserRULE_cellValue               = 8
	MqlParserRULE_columnMapValue          = 9
	MqlParserRULE_superColumnMapValue     = 10
	MqlParserRULE_columnMapEntry          = 11
	MqlParserRULE_superColumnMapEntry     = 12
	MqlParserRULE_columnOrSuperColumnName = 13
	MqlParserRULE_rowKey                  = 14
	MqlParserRULE_columnOrSuperColumnKey  = 15
	MqlParserRULE_columnKey               = 16
	MqlParserRULE_superColumnKey          = 17
)

// IStringValContext is an interface to support dynamic dispatch.
type IStringValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringValContext differentiates from other interfaces.
	IsStringValContext()
}

type StringValContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringValContext() *StringValContext {
	var p = new(StringValContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MqlParserRULE_stringVal
	return p
}

func (*StringValContext) IsStringValContext() {}

func NewStringValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringValContext {
	var p = new(StringValContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MqlParserRULE_stringVal

	return p
}

func (s *StringValContext) GetParser() antlr.Parser { return s.parser }

func (s *StringValContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(MqlParserStringLiteral, 0)
}

func (s *StringValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MqlListener); ok {
		listenerT.EnterStringVal(s)
	}
}

func (s *StringValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MqlListener); ok {
		listenerT.ExitStringVal(s)
	}
}

func (p *MqlParser) StringVal() (localctx IStringValContext) {
	localctx = NewStringValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MqlParserRULE_stringVal)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(36)
		_la = p.GetTokenStream().LA(1)

		if !(_la == MqlParserT__0 || _la == MqlParserStringLiteral) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MqlParserRULE_stmt
	return p
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MqlParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) GetStmt() IGetStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGetStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGetStmtContext)
}

func (s *StmtContext) SetStmt() ISetStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISetStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISetStmtContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MqlListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MqlListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (p *MqlParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MqlParserRULE_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(40)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MqlParserGET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(38)
			p.GetStmt()
		}

	case MqlParserSET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(39)
			p.SetStmt()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IGetStmtContext is an interface to support dynamic dispatch.
type IGetStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGetStmtContext differentiates from other interfaces.
	IsGetStmtContext()
}

type GetStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGetStmtContext() *GetStmtContext {
	var p = new(GetStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MqlParserRULE_getStmt
	return p
}

func (*GetStmtContext) IsGetStmtContext() {}

func NewGetStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GetStmtContext {
	var p = new(GetStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MqlParserRULE_getStmt

	return p
}

func (s *GetStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *GetStmtContext) GET() antlr.TerminalNode {
	return s.GetToken(MqlParserGET, 0)
}

func (s *GetStmtContext) ColumnSpec() IColumnSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnSpecContext)
}

func (s *GetStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GetStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MqlListener); ok {
		listenerT.EnterGetStmt(s)
	}
}

func (s *GetStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MqlListener); ok {
		listenerT.ExitGetStmt(s)
	}
}

func (p *MqlParser) GetStmt() (localctx IGetStmtContext) {
	localctx = NewGetStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MqlParserRULE_getStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		p.Match(MqlParserGET)
	}
	{
		p.SetState(43)
		p.ColumnSpec()
	}

	return localctx
}

// ISetStmtContext is an interface to support dynamic dispatch.
type ISetStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSetStmtContext differentiates from other interfaces.
	IsSetStmtContext()
}

type SetStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetStmtContext() *SetStmtContext {
	var p = new(SetStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MqlParserRULE_setStmt
	return p
}

func (*SetStmtContext) IsSetStmtContext() {}

func NewSetStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetStmtContext {
	var p = new(SetStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MqlParserRULE_setStmt

	return p
}

func (s *SetStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SetStmtContext) SET() antlr.TerminalNode {
	return s.GetToken(MqlParserSET, 0)
}

func (s *SetStmtContext) ColumnSpec() IColumnSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnSpecContext)
}

func (s *SetStmtContext) ValueExpr() IValueExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueExprContext)
}

func (s *SetStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MqlListener); ok {
		listenerT.EnterSetStmt(s)
	}
}

func (s *SetStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MqlListener); ok {
		listenerT.ExitSetStmt(s)
	}
}

func (p *MqlParser) SetStmt() (localctx ISetStmtContext) {
	localctx = NewSetStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MqlParserRULE_setStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(45)
		p.Match(MqlParserSET)
	}
	{
		p.SetState(46)
		p.ColumnSpec()
	}
	{
		p.SetState(47)
		p.Match(MqlParserT__1)
	}
	{
		p.SetState(48)
		p.ValueExpr()
	}

	return localctx
}

// IColumnSpecContext is an interface to support dynamic dispatch.
type IColumnSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_columnOrSuperColumnKey returns the _columnOrSuperColumnKey rule contexts.
	Get_columnOrSuperColumnKey() IColumnOrSuperColumnKeyContext

	// Set_columnOrSuperColumnKey sets the _columnOrSuperColumnKey rule contexts.
	Set_columnOrSuperColumnKey(IColumnOrSuperColumnKeyContext)

	// GetA returns the a rule context list.
	GetA() []IColumnOrSuperColumnKeyContext

	// SetA sets the a rule context list.
	SetA([]IColumnOrSuperColumnKeyContext)

	// IsColumnSpecContext differentiates from other interfaces.
	IsColumnSpecContext()
}

type ColumnSpecContext struct {
	*antlr.BaseParserRuleContext
	parser                  antlr.Parser
	_columnOrSuperColumnKey IColumnOrSuperColumnKeyContext
	a                       []IColumnOrSuperColumnKeyContext
}

func NewEmptyColumnSpecContext() *ColumnSpecContext {
	var p = new(ColumnSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MqlParserRULE_columnSpec
	return p
}

func (*ColumnSpecContext) IsColumnSpecContext() {}

func NewColumnSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnSpecContext {
	var p = new(ColumnSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MqlParserRULE_columnSpec

	return p
}

func (s *ColumnSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnSpecContext) Get_columnOrSuperColumnKey() IColumnOrSuperColumnKeyContext {
	return s._columnOrSuperColumnKey
}

func (s *ColumnSpecContext) Set_columnOrSuperColumnKey(v IColumnOrSuperColumnKeyContext) {
	s._columnOrSuperColumnKey = v
}

func (s *ColumnSpecContext) GetA() []IColumnOrSuperColumnKeyContext { return s.a }

func (s *ColumnSpecContext) SetA(v []IColumnOrSuperColumnKeyContext) { s.a = v }

func (s *ColumnSpecContext) TableName() ITableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *ColumnSpecContext) ColumnFamilyName() IColumnFamilyNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnFamilyNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnFamilyNameContext)
}

func (s *ColumnSpecContext) RowKey() IRowKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRowKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRowKeyContext)
}

func (s *ColumnSpecContext) AllColumnOrSuperColumnKey() []IColumnOrSuperColumnKeyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IColumnOrSuperColumnKeyContext)(nil)).Elem())
	var tst = make([]IColumnOrSuperColumnKeyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IColumnOrSuperColumnKeyContext)
		}
	}

	return tst
}

func (s *ColumnSpecContext) ColumnOrSuperColumnKey(i int) IColumnOrSuperColumnKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnOrSuperColumnKeyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IColumnOrSuperColumnKeyContext)
}

func (s *ColumnSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MqlListener); ok {
		listenerT.EnterColumnSpec(s)
	}
}

func (s *ColumnSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MqlListener); ok {
		listenerT.ExitColumnSpec(s)
	}
}

func (p *MqlParser) ColumnSpec() (localctx IColumnSpecContext) {
	localctx = NewColumnSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MqlParserRULE_columnSpec)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.TableName()
	}
	{
		p.SetState(51)
		p.Match(MqlParserT__2)
	}
	{
		p.SetState(52)
		p.ColumnFamilyName()
	}
	{
		p.SetState(53)
		p.Match(MqlParserT__3)
	}
	{
		p.SetState(54)
		p.RowKey()
	}
	{
		p.SetState(55)
		p.Match(MqlParserT__4)
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == MqlParserT__3 {
		{
			p.SetState(56)
			p.Match(MqlParserT__3)
		}
		{
			p.SetState(57)

			var _x = p.ColumnOrSuperColumnKey()

			localctx.(*ColumnSpecContext)._columnOrSuperColumnKey = _x
		}
		localctx.(*ColumnSpecContext).a = append(localctx.(*ColumnSpecContext).a, localctx.(*ColumnSpecContext)._columnOrSuperColumnKey)
		{
			p.SetState(58)
			p.Match(MqlParserT__4)
		}
		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == MqlParserT__3 {
			{
				p.SetState(59)
				p.Match(MqlParserT__3)
			}
			{
				p.SetState(60)

				var _x = p.ColumnOrSuperColumnKey()

				localctx.(*ColumnSpecContext)._columnOrSuperColumnKey = _x
			}
			localctx.(*ColumnSpecContext).a = append(localctx.(*ColumnSpecContext).a, localctx.(*ColumnSpecContext)._columnOrSuperColumnKey)
			{
				p.SetState(61)
				p.Match(MqlParserT__4)
			}

		}

	}

	return localctx
}

// ITableNameContext is an interface to support dynamic dispatch.
type ITableNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableNameContext differentiates from other interfaces.
	IsTableNameContext()
}

type TableNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableNameContext() *TableNameContext {
	var p = new(TableNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MqlParserRULE_tableName
	return p
}

func (*TableNameContext) IsTableNameContext() {}

func NewTableNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableNameContext {
	var p = new(TableNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MqlParserRULE_tableName

	return p
}

func (s *TableNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TableNameContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MqlParserIdentifier, 0)
}

func (s *TableNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MqlListener); ok {
		listenerT.EnterTableName(s)
	}
}

func (s *TableNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MqlListener); ok {
		listenerT.ExitTableName(s)
	}
}

func (p *MqlParser) TableName() (localctx ITableNameContext) {
	localctx = NewTableNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MqlParserRULE_tableName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.Match(MqlParserIdentifier)
	}

	return localctx
}

// IColumnFamilyNameContext is an interface to support dynamic dispatch.
type IColumnFamilyNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnFamilyNameContext differentiates from other interfaces.
	IsColumnFamilyNameContext()
}

type ColumnFamilyNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnFamilyNameContext() *ColumnFamilyNameContext {
	var p = new(ColumnFamilyNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MqlParserRULE_columnFamilyName
	return p
}

func (*ColumnFamilyNameContext) IsColumnFamilyNameContext() {}

func NewColumnFamilyNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnFamilyNameContext {
	var p = new(ColumnFamilyNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MqlParserRULE_columnFamilyName

	return p
}

func (s *ColumnFamilyNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnFamilyNameContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MqlParserIdentifier, 0)
}

func (s *ColumnFamilyNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnFamilyNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnFamilyNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MqlListener); ok {
		listenerT.EnterColumnFamilyName(s)
	}
}

func (s *ColumnFamilyNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MqlListener); ok {
		listenerT.ExitColumnFamilyName(s)
	}
}

func (p *MqlParser) ColumnFamilyName() (localctx IColumnFamilyNameContext) {
	localctx = NewColumnFamilyNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MqlParserRULE_columnFamilyName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)
		p.Match(MqlParserIdentifier)
	}

	return localctx
}

// IValueExprContext is an interface to support dynamic dispatch.
type IValueExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueExprContext differentiates from other interfaces.
	IsValueExprContext()
}

type ValueExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueExprContext() *ValueExprContext {
	var p = new(ValueExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MqlParserRULE_valueExpr
	return p
}

func (*ValueExprContext) IsValueExprContext() {}

func NewValueExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueExprContext {
	var p = new(ValueExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MqlParserRULE_valueExpr

	return p
}

func (s *ValueExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueExprContext) CellValue() ICellValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICellValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICellValueContext)
}

func (s *ValueExprContext) ColumnMapValue() IColumnMapValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnMapValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnMapValueContext)
}

func (s *ValueExprContext) SuperColumnMapValue() ISuperColumnMapValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISuperColumnMapValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISuperColumnMapValueContext)
}

func (s *ValueExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MqlListener); ok {
		listenerT.EnterValueExpr(s)
	}
}

func (s *ValueExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MqlListener); ok {
		listenerT.ExitValueExpr(s)
	}
}

func (p *MqlParser) ValueExpr() (localctx IValueExprContext) {
	localctx = NewValueExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MqlParserRULE_valueExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(71)
			p.CellValue()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(72)
			p.ColumnMapValue()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(73)
			p.SuperColumnMapValue()
		}

	}

	return localctx
}

// ICellValueContext is an interface to support dynamic dispatch.
type ICellValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCellValueContext differentiates from other interfaces.
	IsCellValueContext()
}

type CellValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCellValueContext() *CellValueContext {
	var p = new(CellValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MqlParserRULE_cellValue
	return p
}

func (*CellValueContext) IsCellValueContext() {}

func NewCellValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CellValueContext {
	var p = new(CellValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MqlParserRULE_cellValue

	return p
}

func (s *CellValueContext) GetParser() antlr.Parser { return s.parser }

func (s *CellValueContext) StringVal() IStringValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringValContext)
}

func (s *CellValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CellValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CellValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MqlListener); ok {
		listenerT.EnterCellValue(s)
	}
}

func (s *CellValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MqlListener); ok {
		listenerT.ExitCellValue(s)
	}
}

func (p *MqlParser) CellValue() (localctx ICellValueContext) {
	localctx = NewCellValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MqlParserRULE_cellValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		p.StringVal()
	}

	return localctx
}

// IColumnMapValueContext is an interface to support dynamic dispatch.
type IColumnMapValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnMapValueContext differentiates from other interfaces.
	IsColumnMapValueContext()
}

type ColumnMapValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnMapValueContext() *ColumnMapValueContext {
	var p = new(ColumnMapValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MqlParserRULE_columnMapValue
	return p
}

func (*ColumnMapValueContext) IsColumnMapValueContext() {}

func NewColumnMapValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnMapValueContext {
	var p = new(ColumnMapValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MqlParserRULE_columnMapValue

	return p
}

func (s *ColumnMapValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnMapValueContext) LEFT_BRACE() antlr.TerminalNode {
	return s.GetToken(MqlParserLEFT_BRACE, 0)
}

func (s *ColumnMapValueContext) AllColumnMapEntry() []IColumnMapEntryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IColumnMapEntryContext)(nil)).Elem())
	var tst = make([]IColumnMapEntryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IColumnMapEntryContext)
		}
	}

	return tst
}

func (s *ColumnMapValueContext) ColumnMapEntry(i int) IColumnMapEntryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnMapEntryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IColumnMapEntryContext)
}

func (s *ColumnMapValueContext) RIGHT_BRACE() antlr.TerminalNode {
	return s.GetToken(MqlParserRIGHT_BRACE, 0)
}

func (s *ColumnMapValueContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MqlParserCOMMA)
}

func (s *ColumnMapValueContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MqlParserCOMMA, i)
}

func (s *ColumnMapValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnMapValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnMapValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MqlListener); ok {
		listenerT.EnterColumnMapValue(s)
	}
}

func (s *ColumnMapValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MqlListener); ok {
		listenerT.ExitColumnMapValue(s)
	}
}

func (p *MqlParser) ColumnMapValue() (localctx IColumnMapValueContext) {
	localctx = NewColumnMapValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MqlParserRULE_columnMapValue)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Match(MqlParserLEFT_BRACE)
	}
	{
		p.SetState(79)
		p.ColumnMapEntry()
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MqlParserCOMMA {
		{
			p.SetState(80)
			p.Match(MqlParserCOMMA)
		}
		{
			p.SetState(81)
			p.ColumnMapEntry()
		}

		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(87)
		p.Match(MqlParserRIGHT_BRACE)
	}

	return localctx
}

// ISuperColumnMapValueContext is an interface to support dynamic dispatch.
type ISuperColumnMapValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSuperColumnMapValueContext differentiates from other interfaces.
	IsSuperColumnMapValueContext()
}

type SuperColumnMapValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySuperColumnMapValueContext() *SuperColumnMapValueContext {
	var p = new(SuperColumnMapValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MqlParserRULE_superColumnMapValue
	return p
}

func (*SuperColumnMapValueContext) IsSuperColumnMapValueContext() {}

func NewSuperColumnMapValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SuperColumnMapValueContext {
	var p = new(SuperColumnMapValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MqlParserRULE_superColumnMapValue

	return p
}

func (s *SuperColumnMapValueContext) GetParser() antlr.Parser { return s.parser }

func (s *SuperColumnMapValueContext) LEFT_BRACE() antlr.TerminalNode {
	return s.GetToken(MqlParserLEFT_BRACE, 0)
}

func (s *SuperColumnMapValueContext) AllSuperColumnMapEntry() []ISuperColumnMapEntryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISuperColumnMapEntryContext)(nil)).Elem())
	var tst = make([]ISuperColumnMapEntryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISuperColumnMapEntryContext)
		}
	}

	return tst
}

func (s *SuperColumnMapValueContext) SuperColumnMapEntry(i int) ISuperColumnMapEntryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISuperColumnMapEntryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISuperColumnMapEntryContext)
}

func (s *SuperColumnMapValueContext) RIGHT_BRACE() antlr.TerminalNode {
	return s.GetToken(MqlParserRIGHT_BRACE, 0)
}

func (s *SuperColumnMapValueContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MqlParserCOMMA)
}

func (s *SuperColumnMapValueContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MqlParserCOMMA, i)
}

func (s *SuperColumnMapValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SuperColumnMapValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SuperColumnMapValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MqlListener); ok {
		listenerT.EnterSuperColumnMapValue(s)
	}
}

func (s *SuperColumnMapValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MqlListener); ok {
		listenerT.ExitSuperColumnMapValue(s)
	}
}

func (p *MqlParser) SuperColumnMapValue() (localctx ISuperColumnMapValueContext) {
	localctx = NewSuperColumnMapValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MqlParserRULE_superColumnMapValue)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
		p.Match(MqlParserLEFT_BRACE)
	}
	{
		p.SetState(90)
		p.SuperColumnMapEntry()
	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MqlParserCOMMA {
		{
			p.SetState(91)
			p.Match(MqlParserCOMMA)
		}
		{
			p.SetState(92)
			p.SuperColumnMapEntry()
		}

		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(98)
		p.Match(MqlParserRIGHT_BRACE)
	}

	return localctx
}

// IColumnMapEntryContext is an interface to support dynamic dispatch.
type IColumnMapEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnMapEntryContext differentiates from other interfaces.
	IsColumnMapEntryContext()
}

type ColumnMapEntryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnMapEntryContext() *ColumnMapEntryContext {
	var p = new(ColumnMapEntryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MqlParserRULE_columnMapEntry
	return p
}

func (*ColumnMapEntryContext) IsColumnMapEntryContext() {}

func NewColumnMapEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnMapEntryContext {
	var p = new(ColumnMapEntryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MqlParserRULE_columnMapEntry

	return p
}

func (s *ColumnMapEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnMapEntryContext) ColumnKey() IColumnKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnKeyContext)
}

func (s *ColumnMapEntryContext) ASSOC() antlr.TerminalNode {
	return s.GetToken(MqlParserASSOC, 0)
}

func (s *ColumnMapEntryContext) CellValue() ICellValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICellValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICellValueContext)
}

func (s *ColumnMapEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnMapEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnMapEntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MqlListener); ok {
		listenerT.EnterColumnMapEntry(s)
	}
}

func (s *ColumnMapEntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MqlListener); ok {
		listenerT.ExitColumnMapEntry(s)
	}
}

func (p *MqlParser) ColumnMapEntry() (localctx IColumnMapEntryContext) {
	localctx = NewColumnMapEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MqlParserRULE_columnMapEntry)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.ColumnKey()
	}
	{
		p.SetState(101)
		p.Match(MqlParserASSOC)
	}
	{
		p.SetState(102)
		p.CellValue()
	}

	return localctx
}

// ISuperColumnMapEntryContext is an interface to support dynamic dispatch.
type ISuperColumnMapEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSuperColumnMapEntryContext differentiates from other interfaces.
	IsSuperColumnMapEntryContext()
}

type SuperColumnMapEntryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySuperColumnMapEntryContext() *SuperColumnMapEntryContext {
	var p = new(SuperColumnMapEntryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MqlParserRULE_superColumnMapEntry
	return p
}

func (*SuperColumnMapEntryContext) IsSuperColumnMapEntryContext() {}

func NewSuperColumnMapEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SuperColumnMapEntryContext {
	var p = new(SuperColumnMapEntryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MqlParserRULE_superColumnMapEntry

	return p
}

func (s *SuperColumnMapEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *SuperColumnMapEntryContext) SuperColumnKey() ISuperColumnKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISuperColumnKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISuperColumnKeyContext)
}

func (s *SuperColumnMapEntryContext) ASSOC() antlr.TerminalNode {
	return s.GetToken(MqlParserASSOC, 0)
}

func (s *SuperColumnMapEntryContext) ColumnMapValue() IColumnMapValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnMapValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnMapValueContext)
}

func (s *SuperColumnMapEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SuperColumnMapEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SuperColumnMapEntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MqlListener); ok {
		listenerT.EnterSuperColumnMapEntry(s)
	}
}

func (s *SuperColumnMapEntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MqlListener); ok {
		listenerT.ExitSuperColumnMapEntry(s)
	}
}

func (p *MqlParser) SuperColumnMapEntry() (localctx ISuperColumnMapEntryContext) {
	localctx = NewSuperColumnMapEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MqlParserRULE_superColumnMapEntry)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		p.SuperColumnKey()
	}
	{
		p.SetState(105)
		p.Match(MqlParserASSOC)
	}
	{
		p.SetState(106)
		p.ColumnMapValue()
	}

	return localctx
}

// IColumnOrSuperColumnNameContext is an interface to support dynamic dispatch.
type IColumnOrSuperColumnNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnOrSuperColumnNameContext differentiates from other interfaces.
	IsColumnOrSuperColumnNameContext()
}

type ColumnOrSuperColumnNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnOrSuperColumnNameContext() *ColumnOrSuperColumnNameContext {
	var p = new(ColumnOrSuperColumnNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MqlParserRULE_columnOrSuperColumnName
	return p
}

func (*ColumnOrSuperColumnNameContext) IsColumnOrSuperColumnNameContext() {}

func NewColumnOrSuperColumnNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnOrSuperColumnNameContext {
	var p = new(ColumnOrSuperColumnNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MqlParserRULE_columnOrSuperColumnName

	return p
}

func (s *ColumnOrSuperColumnNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnOrSuperColumnNameContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MqlParserIdentifier, 0)
}

func (s *ColumnOrSuperColumnNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnOrSuperColumnNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnOrSuperColumnNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MqlListener); ok {
		listenerT.EnterColumnOrSuperColumnName(s)
	}
}

func (s *ColumnOrSuperColumnNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MqlListener); ok {
		listenerT.ExitColumnOrSuperColumnName(s)
	}
}

func (p *MqlParser) ColumnOrSuperColumnName() (localctx IColumnOrSuperColumnNameContext) {
	localctx = NewColumnOrSuperColumnNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, MqlParserRULE_columnOrSuperColumnName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(MqlParserIdentifier)
	}

	return localctx
}

// IRowKeyContext is an interface to support dynamic dispatch.
type IRowKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRowKeyContext differentiates from other interfaces.
	IsRowKeyContext()
}

type RowKeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRowKeyContext() *RowKeyContext {
	var p = new(RowKeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MqlParserRULE_rowKey
	return p
}

func (*RowKeyContext) IsRowKeyContext() {}

func NewRowKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RowKeyContext {
	var p = new(RowKeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MqlParserRULE_rowKey

	return p
}

func (s *RowKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *RowKeyContext) StringVal() IStringValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringValContext)
}

func (s *RowKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RowKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RowKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MqlListener); ok {
		listenerT.EnterRowKey(s)
	}
}

func (s *RowKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MqlListener); ok {
		listenerT.ExitRowKey(s)
	}
}

func (p *MqlParser) RowKey() (localctx IRowKeyContext) {
	localctx = NewRowKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, MqlParserRULE_rowKey)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)
		p.StringVal()
	}

	return localctx
}

// IColumnOrSuperColumnKeyContext is an interface to support dynamic dispatch.
type IColumnOrSuperColumnKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnOrSuperColumnKeyContext differentiates from other interfaces.
	IsColumnOrSuperColumnKeyContext()
}

type ColumnOrSuperColumnKeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnOrSuperColumnKeyContext() *ColumnOrSuperColumnKeyContext {
	var p = new(ColumnOrSuperColumnKeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MqlParserRULE_columnOrSuperColumnKey
	return p
}

func (*ColumnOrSuperColumnKeyContext) IsColumnOrSuperColumnKeyContext() {}

func NewColumnOrSuperColumnKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnOrSuperColumnKeyContext {
	var p = new(ColumnOrSuperColumnKeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MqlParserRULE_columnOrSuperColumnKey

	return p
}

func (s *ColumnOrSuperColumnKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnOrSuperColumnKeyContext) StringVal() IStringValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringValContext)
}

func (s *ColumnOrSuperColumnKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnOrSuperColumnKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnOrSuperColumnKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MqlListener); ok {
		listenerT.EnterColumnOrSuperColumnKey(s)
	}
}

func (s *ColumnOrSuperColumnKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MqlListener); ok {
		listenerT.ExitColumnOrSuperColumnKey(s)
	}
}

func (p *MqlParser) ColumnOrSuperColumnKey() (localctx IColumnOrSuperColumnKeyContext) {
	localctx = NewColumnOrSuperColumnKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, MqlParserRULE_columnOrSuperColumnKey)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		p.StringVal()
	}

	return localctx
}

// IColumnKeyContext is an interface to support dynamic dispatch.
type IColumnKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnKeyContext differentiates from other interfaces.
	IsColumnKeyContext()
}

type ColumnKeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnKeyContext() *ColumnKeyContext {
	var p = new(ColumnKeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MqlParserRULE_columnKey
	return p
}

func (*ColumnKeyContext) IsColumnKeyContext() {}

func NewColumnKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnKeyContext {
	var p = new(ColumnKeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MqlParserRULE_columnKey

	return p
}

func (s *ColumnKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnKeyContext) StringVal() IStringValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringValContext)
}

func (s *ColumnKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MqlListener); ok {
		listenerT.EnterColumnKey(s)
	}
}

func (s *ColumnKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MqlListener); ok {
		listenerT.ExitColumnKey(s)
	}
}

func (p *MqlParser) ColumnKey() (localctx IColumnKeyContext) {
	localctx = NewColumnKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, MqlParserRULE_columnKey)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.StringVal()
	}

	return localctx
}

// ISuperColumnKeyContext is an interface to support dynamic dispatch.
type ISuperColumnKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSuperColumnKeyContext differentiates from other interfaces.
	IsSuperColumnKeyContext()
}

type SuperColumnKeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySuperColumnKeyContext() *SuperColumnKeyContext {
	var p = new(SuperColumnKeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MqlParserRULE_superColumnKey
	return p
}

func (*SuperColumnKeyContext) IsSuperColumnKeyContext() {}

func NewSuperColumnKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SuperColumnKeyContext {
	var p = new(SuperColumnKeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MqlParserRULE_superColumnKey

	return p
}

func (s *SuperColumnKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *SuperColumnKeyContext) StringVal() IStringValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringValContext)
}

func (s *SuperColumnKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SuperColumnKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SuperColumnKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MqlListener); ok {
		listenerT.EnterSuperColumnKey(s)
	}
}

func (s *SuperColumnKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MqlListener); ok {
		listenerT.ExitSuperColumnKey(s)
	}
}

func (p *MqlParser) SuperColumnKey() (localctx ISuperColumnKeyContext) {
	localctx = NewSuperColumnKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, MqlParserRULE_superColumnKey)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		p.StringVal()
	}

	return localctx
}
