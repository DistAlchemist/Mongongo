// Code generated from Mql.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 18, 114,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5,
	3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 6, 9,
	59, 10, 9, 13, 9, 14, 9, 60, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3,
	11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16,
	3, 17, 3, 17, 3, 17, 3, 17, 7, 17, 84, 10, 17, 12, 17, 14, 17, 87, 11,
	17, 3, 18, 3, 18, 7, 18, 91, 10, 18, 12, 18, 14, 18, 94, 11, 18, 3, 18,
	3, 18, 3, 18, 7, 18, 99, 10, 18, 12, 18, 14, 18, 102, 11, 18, 3, 18, 7,
	18, 105, 10, 18, 12, 18, 14, 18, 108, 11, 18, 3, 19, 6, 19, 111, 10, 19,
	13, 19, 14, 19, 112, 2, 2, 20, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15,
	9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 2, 31, 2, 33, 16,
	35, 17, 37, 18, 3, 2, 5, 5, 2, 11, 12, 15, 15, 34, 34, 4, 2, 67, 92, 99,
	124, 3, 2, 41, 41, 2, 119, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3,
	2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15,
	3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2,
	23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2,
	2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 3, 39, 3, 2, 2, 2, 5, 41, 3, 2, 2,
	2, 7, 43, 3, 2, 2, 2, 9, 45, 3, 2, 2, 2, 11, 47, 3, 2, 2, 2, 13, 49, 3,
	2, 2, 2, 15, 53, 3, 2, 2, 2, 17, 58, 3, 2, 2, 2, 19, 64, 3, 2, 2, 2, 21,
	67, 3, 2, 2, 2, 23, 69, 3, 2, 2, 2, 25, 71, 3, 2, 2, 2, 27, 73, 3, 2, 2,
	2, 29, 75, 3, 2, 2, 2, 31, 77, 3, 2, 2, 2, 33, 79, 3, 2, 2, 2, 35, 88,
	3, 2, 2, 2, 37, 110, 3, 2, 2, 2, 39, 40, 7, 65, 2, 2, 40, 4, 3, 2, 2, 2,
	41, 42, 7, 63, 2, 2, 42, 6, 3, 2, 2, 2, 43, 44, 7, 48, 2, 2, 44, 8, 3,
	2, 2, 2, 45, 46, 7, 93, 2, 2, 46, 10, 3, 2, 2, 2, 47, 48, 7, 95, 2, 2,
	48, 12, 3, 2, 2, 2, 49, 50, 7, 73, 2, 2, 50, 51, 7, 71, 2, 2, 51, 52, 7,
	86, 2, 2, 52, 14, 3, 2, 2, 2, 53, 54, 7, 85, 2, 2, 54, 55, 7, 71, 2, 2,
	55, 56, 7, 86, 2, 2, 56, 16, 3, 2, 2, 2, 57, 59, 9, 2, 2, 2, 58, 57, 3,
	2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 60, 61, 3, 2, 2, 2, 61,
	62, 3, 2, 2, 2, 62, 63, 8, 9, 2, 2, 63, 18, 3, 2, 2, 2, 64, 65, 7, 63,
	2, 2, 65, 66, 7, 64, 2, 2, 66, 20, 3, 2, 2, 2, 67, 68, 7, 46, 2, 2, 68,
	22, 3, 2, 2, 2, 69, 70, 7, 125, 2, 2, 70, 24, 3, 2, 2, 2, 71, 72, 7, 127,
	2, 2, 72, 26, 3, 2, 2, 2, 73, 74, 7, 61, 2, 2, 74, 28, 3, 2, 2, 2, 75,
	76, 9, 3, 2, 2, 76, 30, 3, 2, 2, 2, 77, 78, 4, 50, 59, 2, 78, 32, 3, 2,
	2, 2, 79, 85, 5, 29, 15, 2, 80, 84, 5, 29, 15, 2, 81, 84, 5, 31, 16, 2,
	82, 84, 7, 97, 2, 2, 83, 80, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 83, 82, 3,
	2, 2, 2, 84, 87, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 85, 86, 3, 2, 2, 2, 86,
	34, 3, 2, 2, 2, 87, 85, 3, 2, 2, 2, 88, 92, 7, 41, 2, 2, 89, 91, 10, 4,
	2, 2, 90, 89, 3, 2, 2, 2, 91, 94, 3, 2, 2, 2, 92, 90, 3, 2, 2, 2, 92, 93,
	3, 2, 2, 2, 93, 95, 3, 2, 2, 2, 94, 92, 3, 2, 2, 2, 95, 106, 7, 41, 2,
	2, 96, 100, 7, 41, 2, 2, 97, 99, 10, 4, 2, 2, 98, 97, 3, 2, 2, 2, 99, 102,
	3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 103, 3, 2,
	2, 2, 102, 100, 3, 2, 2, 2, 103, 105, 7, 41, 2, 2, 104, 96, 3, 2, 2, 2,
	105, 108, 3, 2, 2, 2, 106, 104, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107,
	36, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 109, 111, 5, 31, 16, 2, 110, 109,
	3, 2, 2, 2, 111, 112, 3, 2, 2, 2, 112, 110, 3, 2, 2, 2, 112, 113, 3, 2,
	2, 2, 113, 38, 3, 2, 2, 2, 10, 2, 60, 83, 85, 92, 100, 106, 112, 3, 8,
	2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'?'", "'='", "'.'", "'['", "']'", "'GET'", "'SET'", "", "'=>'", "','",
	"'{'", "'}'", "';'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "GET", "SET", "WHITESPACE", "ASSOC", "COMMA", "LEFT_BRACE",
	"RIGHT_BRACE", "SEMICOLON", "Identifier", "StringLiteral", "IntegerLiteral",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "GET", "SET", "WHITESPACE", "ASSOC",
	"COMMA", "LEFT_BRACE", "RIGHT_BRACE", "SEMICOLON", "Letter", "Digit", "Identifier",
	"StringLiteral", "IntegerLiteral",
}

type MqlLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewMqlLexer(input antlr.CharStream) *MqlLexer {

	l := new(MqlLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Mql.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MqlLexer tokens.
const (
	MqlLexerT__0           = 1
	MqlLexerT__1           = 2
	MqlLexerT__2           = 3
	MqlLexerT__3           = 4
	MqlLexerT__4           = 5
	MqlLexerGET            = 6
	MqlLexerSET            = 7
	MqlLexerWHITESPACE     = 8
	MqlLexerASSOC          = 9
	MqlLexerCOMMA          = 10
	MqlLexerLEFT_BRACE     = 11
	MqlLexerRIGHT_BRACE    = 12
	MqlLexerSEMICOLON      = 13
	MqlLexerIdentifier     = 14
	MqlLexerStringLiteral  = 15
	MqlLexerIntegerLiteral = 16
)
