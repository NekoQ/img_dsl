// Generated from Img.g4 by ANTLR 4.7.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 22, 126,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3,
	17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 6, 19, 111, 10, 19, 13, 19,
	14, 19, 112, 3, 20, 6, 20, 116, 10, 20, 13, 20, 14, 20, 117, 3, 21, 6,
	21, 121, 10, 21, 13, 21, 14, 21, 122, 3, 21, 3, 21, 2, 2, 22, 3, 3, 5,
	4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25,
	14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 3,
	2, 5, 3, 2, 50, 59, 5, 2, 67, 92, 97, 97, 99, 124, 5, 2, 11, 12, 15, 15,
	34, 34, 2, 128, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2,
	9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2,
	2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2,
	2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2,
	2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3,
	2, 2, 2, 2, 41, 3, 2, 2, 2, 3, 43, 3, 2, 2, 2, 5, 47, 3, 2, 2, 2, 7, 51,
	3, 2, 2, 2, 9, 56, 3, 2, 2, 2, 11, 63, 3, 2, 2, 2, 13, 69, 3, 2, 2, 2,
	15, 75, 3, 2, 2, 2, 17, 82, 3, 2, 2, 2, 19, 87, 3, 2, 2, 2, 21, 89, 3,
	2, 2, 2, 23, 91, 3, 2, 2, 2, 25, 93, 3, 2, 2, 2, 27, 95, 3, 2, 2, 2, 29,
	97, 3, 2, 2, 2, 31, 99, 3, 2, 2, 2, 33, 101, 3, 2, 2, 2, 35, 104, 3, 2,
	2, 2, 37, 110, 3, 2, 2, 2, 39, 115, 3, 2, 2, 2, 41, 120, 3, 2, 2, 2, 43,
	44, 7, 114, 2, 2, 44, 45, 7, 112, 2, 2, 45, 46, 7, 105, 2, 2, 46, 4, 3,
	2, 2, 2, 47, 48, 7, 108, 2, 2, 48, 49, 7, 114, 2, 2, 49, 50, 7, 105, 2,
	2, 50, 6, 3, 2, 2, 2, 51, 52, 7, 101, 2, 2, 52, 53, 7, 116, 2, 2, 53, 54,
	7, 113, 2, 2, 54, 55, 7, 114, 2, 2, 55, 8, 3, 2, 2, 2, 56, 57, 7, 116,
	2, 2, 57, 58, 7, 113, 2, 2, 58, 59, 7, 118, 2, 2, 59, 60, 7, 99, 2, 2,
	60, 61, 7, 118, 2, 2, 61, 62, 7, 103, 2, 2, 62, 10, 3, 2, 2, 2, 63, 64,
	7, 104, 2, 2, 64, 65, 7, 110, 2, 2, 65, 66, 7, 107, 2, 2, 66, 67, 7, 114,
	2, 2, 67, 68, 7, 90, 2, 2, 68, 12, 3, 2, 2, 2, 69, 70, 7, 104, 2, 2, 70,
	71, 7, 110, 2, 2, 71, 72, 7, 107, 2, 2, 72, 73, 7, 114, 2, 2, 73, 74, 7,
	91, 2, 2, 74, 14, 3, 2, 2, 2, 75, 76, 7, 116, 2, 2, 76, 77, 7, 103, 2,
	2, 77, 78, 7, 117, 2, 2, 78, 79, 7, 107, 2, 2, 79, 80, 7, 124, 2, 2, 80,
	81, 7, 103, 2, 2, 81, 16, 3, 2, 2, 2, 82, 83, 7, 117, 2, 2, 83, 84, 7,
	99, 2, 2, 84, 85, 7, 120, 2, 2, 85, 86, 7, 103, 2, 2, 86, 18, 3, 2, 2,
	2, 87, 88, 7, 63, 2, 2, 88, 20, 3, 2, 2, 2, 89, 90, 7, 48, 2, 2, 90, 22,
	3, 2, 2, 2, 91, 92, 7, 46, 2, 2, 92, 24, 3, 2, 2, 2, 93, 94, 7, 36, 2,
	2, 94, 26, 3, 2, 2, 2, 95, 96, 7, 61, 2, 2, 96, 28, 3, 2, 2, 2, 97, 98,
	7, 42, 2, 2, 98, 30, 3, 2, 2, 2, 99, 100, 7, 43, 2, 2, 100, 32, 3, 2, 2,
	2, 101, 102, 7, 93, 2, 2, 102, 103, 7, 95, 2, 2, 103, 34, 3, 2, 2, 2, 104,
	105, 7, 113, 2, 2, 105, 106, 7, 114, 2, 2, 106, 107, 7, 103, 2, 2, 107,
	108, 7, 112, 2, 2, 108, 36, 3, 2, 2, 2, 109, 111, 9, 2, 2, 2, 110, 109,
	3, 2, 2, 2, 111, 112, 3, 2, 2, 2, 112, 110, 3, 2, 2, 2, 112, 113, 3, 2,
	2, 2, 113, 38, 3, 2, 2, 2, 114, 116, 9, 3, 2, 2, 115, 114, 3, 2, 2, 2,
	116, 117, 3, 2, 2, 2, 117, 115, 3, 2, 2, 2, 117, 118, 3, 2, 2, 2, 118,
	40, 3, 2, 2, 2, 119, 121, 9, 4, 2, 2, 120, 119, 3, 2, 2, 2, 121, 122, 3,
	2, 2, 2, 122, 120, 3, 2, 2, 2, 122, 123, 3, 2, 2, 2, 123, 124, 3, 2, 2,
	2, 124, 125, 8, 21, 2, 2, 125, 42, 3, 2, 2, 2, 6, 2, 112, 117, 122, 3,
	8, 2, 2,
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
	"", "'png'", "'jpg'", "'crop'", "'rotate'", "'flipX'", "'flipY'", "'resize'",
	"'save'", "'='", "'.'", "','", "'\"'", "';'", "'('", "')'", "'[]'", "'open'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "ASSIGN", "DOT", "COMMA", "APOS", "SEMICOLON",
	"OPEN_PARAN", "CLOSED_PARAN", "FOLDER", "OPEN", "NUMBER", "IDENT", "WHITESPACE",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "ASSIGN",
	"DOT", "COMMA", "APOS", "SEMICOLON", "OPEN_PARAN", "CLOSED_PARAN", "FOLDER",
	"OPEN", "NUMBER", "IDENT", "WHITESPACE",
}

type ImgLexer struct {
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

func NewImgLexer(input antlr.CharStream) *ImgLexer {

	l := new(ImgLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Img.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ImgLexer tokens.
const (
	ImgLexerT__0         = 1
	ImgLexerT__1         = 2
	ImgLexerT__2         = 3
	ImgLexerT__3         = 4
	ImgLexerT__4         = 5
	ImgLexerT__5         = 6
	ImgLexerT__6         = 7
	ImgLexerT__7         = 8
	ImgLexerASSIGN       = 9
	ImgLexerDOT          = 10
	ImgLexerCOMMA        = 11
	ImgLexerAPOS         = 12
	ImgLexerSEMICOLON    = 13
	ImgLexerOPEN_PARAN   = 14
	ImgLexerCLOSED_PARAN = 15
	ImgLexerFOLDER       = 16
	ImgLexerOPEN         = 17
	ImgLexerNUMBER       = 18
	ImgLexerIDENT        = 19
	ImgLexerWHITESPACE   = 20
)
