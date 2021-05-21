// Generated from Img.g4 by ANTLR 4.7.

package parser // Img

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 28, 203,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 3, 2, 7, 2,
	46, 10, 2, 12, 2, 14, 2, 49, 11, 2, 3, 2, 6, 2, 52, 10, 2, 13, 2, 14, 2,
	53, 3, 2, 7, 2, 57, 10, 2, 12, 2, 14, 2, 60, 11, 2, 3, 2, 7, 2, 63, 10,
	2, 12, 2, 14, 2, 66, 11, 2, 3, 2, 7, 2, 69, 10, 2, 12, 2, 14, 2, 72, 11,
	2, 3, 2, 6, 2, 75, 10, 2, 13, 2, 14, 2, 76, 3, 2, 3, 2, 3, 3, 3, 3, 5,
	3, 83, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10,
	128, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5,
	15, 159, 10, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20,
	189, 10, 20, 3, 20, 3, 20, 3, 20, 3, 21, 7, 21, 195, 10, 21, 12, 21, 14,
	21, 198, 11, 21, 3, 22, 3, 22, 3, 22, 3, 22, 2, 2, 23, 2, 4, 6, 8, 10,
	12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 2, 3, 3,
	2, 3, 4, 2, 199, 2, 51, 3, 2, 2, 2, 4, 82, 3, 2, 2, 2, 6, 84, 3, 2, 2,
	2, 8, 92, 3, 2, 2, 2, 10, 101, 3, 2, 2, 2, 12, 107, 3, 2, 2, 2, 14, 109,
	3, 2, 2, 2, 16, 113, 3, 2, 2, 2, 18, 127, 3, 2, 2, 2, 20, 129, 3, 2, 2,
	2, 22, 140, 3, 2, 2, 2, 24, 145, 3, 2, 2, 2, 26, 149, 3, 2, 2, 2, 28, 153,
	3, 2, 2, 2, 30, 162, 3, 2, 2, 2, 32, 167, 3, 2, 2, 2, 34, 172, 3, 2, 2,
	2, 36, 177, 3, 2, 2, 2, 38, 182, 3, 2, 2, 2, 40, 196, 3, 2, 2, 2, 42, 199,
	3, 2, 2, 2, 44, 46, 5, 42, 22, 2, 45, 44, 3, 2, 2, 2, 46, 49, 3, 2, 2,
	2, 47, 45, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 50, 3, 2, 2, 2, 49, 47,
	3, 2, 2, 2, 50, 52, 5, 4, 3, 2, 51, 47, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2,
	53, 51, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 58, 3, 2, 2, 2, 55, 57, 5,
	42, 22, 2, 56, 55, 3, 2, 2, 2, 57, 60, 3, 2, 2, 2, 58, 56, 3, 2, 2, 2,
	58, 59, 3, 2, 2, 2, 59, 64, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 61, 63, 5,
	16, 9, 2, 62, 61, 3, 2, 2, 2, 63, 66, 3, 2, 2, 2, 64, 62, 3, 2, 2, 2, 64,
	65, 3, 2, 2, 2, 65, 70, 3, 2, 2, 2, 66, 64, 3, 2, 2, 2, 67, 69, 5, 42,
	22, 2, 68, 67, 3, 2, 2, 2, 69, 72, 3, 2, 2, 2, 70, 68, 3, 2, 2, 2, 70,
	71, 3, 2, 2, 2, 71, 74, 3, 2, 2, 2, 72, 70, 3, 2, 2, 2, 73, 75, 5, 38,
	20, 2, 74, 73, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 74, 3, 2, 2, 2, 76,
	77, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 79, 7, 2, 2, 3, 79, 3, 3, 2, 2,
	2, 80, 83, 5, 6, 4, 2, 81, 83, 5, 8, 5, 2, 82, 80, 3, 2, 2, 2, 82, 81,
	3, 2, 2, 2, 83, 5, 3, 2, 2, 2, 84, 85, 7, 27, 2, 2, 85, 86, 7, 16, 2, 2,
	86, 87, 7, 24, 2, 2, 87, 88, 7, 21, 2, 2, 88, 89, 5, 10, 6, 2, 89, 90,
	7, 22, 2, 2, 90, 91, 7, 20, 2, 2, 91, 7, 3, 2, 2, 2, 92, 93, 7, 27, 2,
	2, 93, 94, 7, 23, 2, 2, 94, 95, 7, 16, 2, 2, 95, 96, 7, 24, 2, 2, 96, 97,
	7, 21, 2, 2, 97, 98, 5, 14, 8, 2, 98, 99, 7, 22, 2, 2, 99, 100, 7, 20,
	2, 2, 100, 9, 3, 2, 2, 2, 101, 102, 7, 19, 2, 2, 102, 103, 7, 27, 2, 2,
	103, 104, 7, 17, 2, 2, 104, 105, 5, 12, 7, 2, 105, 106, 7, 19, 2, 2, 106,
	11, 3, 2, 2, 2, 107, 108, 9, 2, 2, 2, 108, 13, 3, 2, 2, 2, 109, 110, 7,
	19, 2, 2, 110, 111, 7, 27, 2, 2, 111, 112, 7, 19, 2, 2, 112, 15, 3, 2,
	2, 2, 113, 114, 7, 27, 2, 2, 114, 115, 7, 17, 2, 2, 115, 116, 5, 18, 10,
	2, 116, 117, 7, 20, 2, 2, 117, 17, 3, 2, 2, 2, 118, 128, 5, 20, 11, 2,
	119, 128, 5, 22, 12, 2, 120, 128, 5, 24, 13, 2, 121, 128, 5, 26, 14, 2,
	122, 128, 5, 28, 15, 2, 123, 128, 5, 30, 16, 2, 124, 128, 5, 32, 17, 2,
	125, 128, 5, 34, 18, 2, 126, 128, 5, 36, 19, 2, 127, 118, 3, 2, 2, 2, 127,
	119, 3, 2, 2, 2, 127, 120, 3, 2, 2, 2, 127, 121, 3, 2, 2, 2, 127, 122,
	3, 2, 2, 2, 127, 123, 3, 2, 2, 2, 127, 124, 3, 2, 2, 2, 127, 125, 3, 2,
	2, 2, 127, 126, 3, 2, 2, 2, 128, 19, 3, 2, 2, 2, 129, 130, 7, 5, 2, 2,
	130, 131, 7, 21, 2, 2, 131, 132, 7, 25, 2, 2, 132, 133, 7, 18, 2, 2, 133,
	134, 7, 25, 2, 2, 134, 135, 7, 18, 2, 2, 135, 136, 7, 25, 2, 2, 136, 137,
	7, 18, 2, 2, 137, 138, 7, 25, 2, 2, 138, 139, 7, 22, 2, 2, 139, 21, 3,
	2, 2, 2, 140, 141, 7, 6, 2, 2, 141, 142, 7, 21, 2, 2, 142, 143, 7, 25,
	2, 2, 143, 144, 7, 22, 2, 2, 144, 23, 3, 2, 2, 2, 145, 146, 7, 7, 2, 2,
	146, 147, 7, 21, 2, 2, 147, 148, 7, 22, 2, 2, 148, 25, 3, 2, 2, 2, 149,
	150, 7, 8, 2, 2, 150, 151, 7, 21, 2, 2, 151, 152, 7, 22, 2, 2, 152, 27,
	3, 2, 2, 2, 153, 154, 7, 9, 2, 2, 154, 155, 7, 21, 2, 2, 155, 158, 7, 25,
	2, 2, 156, 157, 7, 18, 2, 2, 157, 159, 7, 25, 2, 2, 158, 156, 3, 2, 2,
	2, 158, 159, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 160, 161, 7, 22, 2, 2, 161,
	29, 3, 2, 2, 2, 162, 163, 7, 10, 2, 2, 163, 164, 7, 21, 2, 2, 164, 165,
	7, 25, 2, 2, 165, 166, 7, 22, 2, 2, 166, 31, 3, 2, 2, 2, 167, 168, 7, 11,
	2, 2, 168, 169, 7, 21, 2, 2, 169, 170, 7, 25, 2, 2, 170, 171, 7, 22, 2,
	2, 171, 33, 3, 2, 2, 2, 172, 173, 7, 12, 2, 2, 173, 174, 7, 21, 2, 2, 174,
	175, 7, 25, 2, 2, 175, 176, 7, 22, 2, 2, 176, 35, 3, 2, 2, 2, 177, 178,
	7, 13, 2, 2, 178, 179, 7, 21, 2, 2, 179, 180, 7, 25, 2, 2, 180, 181, 7,
	22, 2, 2, 181, 37, 3, 2, 2, 2, 182, 183, 7, 27, 2, 2, 183, 184, 7, 17,
	2, 2, 184, 185, 7, 14, 2, 2, 185, 188, 7, 21, 2, 2, 186, 189, 5, 10, 6,
	2, 187, 189, 5, 14, 8, 2, 188, 186, 3, 2, 2, 2, 188, 187, 3, 2, 2, 2, 189,
	190, 3, 2, 2, 2, 190, 191, 7, 22, 2, 2, 191, 192, 7, 20, 2, 2, 192, 39,
	3, 2, 2, 2, 193, 195, 7, 27, 2, 2, 194, 193, 3, 2, 2, 2, 195, 198, 3, 2,
	2, 2, 196, 194, 3, 2, 2, 2, 196, 197, 3, 2, 2, 2, 197, 41, 3, 2, 2, 2,
	198, 196, 3, 2, 2, 2, 199, 200, 7, 15, 2, 2, 200, 201, 5, 40, 21, 2, 201,
	43, 3, 2, 2, 2, 13, 47, 53, 58, 64, 70, 76, 82, 127, 158, 188, 196,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'png'", "'jpg'", "'crop'", "'rotate'", "'flipX'", "'flipY'", "'resize'",
	"'brightness'", "'contrast'", "'saturation'", "'pixelate'", "'save'", "'--'",
	"'='", "'.'", "','", "'\"'", "';'", "'('", "')'", "'[]'", "'open'", "",
	"'-'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "ASSIGN", "DOT",
	"COMMA", "APOS", "SEMICOLON", "OPEN_PARAN", "CLOSED_PARAN", "FOLDER", "OPEN",
	"NUMBER", "MINUS", "ID", "WHITESPACE",
}

var ruleNames = []string{
	"start", "declaration", "openFile", "openFolder", "fileName", "imageType",
	"folderName", "action_", "actionType", "crop", "rotate", "flipX", "flipY",
	"resize", "brightness", "contrast", "saturation", "pixelate", "export",
	"text", "comment",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ImgParser struct {
	*antlr.BaseParser
}

func NewImgParser(input antlr.TokenStream) *ImgParser {
	this := new(ImgParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Img.g4"

	return this
}

// ImgParser tokens.
const (
	ImgParserEOF          = antlr.TokenEOF
	ImgParserT__0         = 1
	ImgParserT__1         = 2
	ImgParserT__2         = 3
	ImgParserT__3         = 4
	ImgParserT__4         = 5
	ImgParserT__5         = 6
	ImgParserT__6         = 7
	ImgParserT__7         = 8
	ImgParserT__8         = 9
	ImgParserT__9         = 10
	ImgParserT__10        = 11
	ImgParserT__11        = 12
	ImgParserT__12        = 13
	ImgParserASSIGN       = 14
	ImgParserDOT          = 15
	ImgParserCOMMA        = 16
	ImgParserAPOS         = 17
	ImgParserSEMICOLON    = 18
	ImgParserOPEN_PARAN   = 19
	ImgParserCLOSED_PARAN = 20
	ImgParserFOLDER       = 21
	ImgParserOPEN         = 22
	ImgParserNUMBER       = 23
	ImgParserMINUS        = 24
	ImgParserID           = 25
	ImgParserWHITESPACE   = 26
)

// ImgParser rules.
const (
	ImgParserRULE_start       = 0
	ImgParserRULE_declaration = 1
	ImgParserRULE_openFile    = 2
	ImgParserRULE_openFolder  = 3
	ImgParserRULE_fileName    = 4
	ImgParserRULE_imageType   = 5
	ImgParserRULE_folderName  = 6
	ImgParserRULE_action_     = 7
	ImgParserRULE_actionType  = 8
	ImgParserRULE_crop        = 9
	ImgParserRULE_rotate      = 10
	ImgParserRULE_flipX       = 11
	ImgParserRULE_flipY       = 12
	ImgParserRULE_resize      = 13
	ImgParserRULE_brightness  = 14
	ImgParserRULE_contrast    = 15
	ImgParserRULE_saturation  = 16
	ImgParserRULE_pixelate    = 17
	ImgParserRULE_export      = 18
	ImgParserRULE_text        = 19
	ImgParserRULE_comment     = 20
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ImgParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ImgParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(ImgParserEOF, 0)
}

func (s *StartContext) AllDeclaration() []IDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDeclarationContext)(nil)).Elem())
	var tst = make([]IDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeclarationContext)
		}
	}

	return tst
}

func (s *StartContext) Declaration(i int) IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *StartContext) AllComment() []ICommentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommentContext)(nil)).Elem())
	var tst = make([]ICommentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommentContext)
		}
	}

	return tst
}

func (s *StartContext) Comment(i int) ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *StartContext) AllAction_() []IAction_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAction_Context)(nil)).Elem())
	var tst = make([]IAction_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAction_Context)
		}
	}

	return tst
}

func (s *StartContext) Action_(i int) IAction_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAction_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAction_Context)
}

func (s *StartContext) AllExport() []IExportContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExportContext)(nil)).Elem())
	var tst = make([]IExportContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExportContext)
		}
	}

	return tst
}

func (s *StartContext) Export(i int) IExportContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExportContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExportContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *ImgParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ImgParserRULE_start)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(45)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ImgParserT__12 {
				{
					p.SetState(42)
					p.Comment()
				}

				p.SetState(47)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(48)
				p.Declaration()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(53)
				p.Comment()
			}

		}
		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(59)
				p.Action_()
			}

		}
		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ImgParserT__12 {
		{
			p.SetState(65)
			p.Comment()
		}

		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ImgParserID {
		{
			p.SetState(71)
			p.Export()
		}

		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(76)
		p.Match(ImgParserEOF)
	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ImgParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ImgParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) OpenFile() IOpenFileContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpenFileContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOpenFileContext)
}

func (s *DeclarationContext) OpenFolder() IOpenFolderContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpenFolderContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOpenFolderContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (p *ImgParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ImgParserRULE_declaration)

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

	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(78)
			p.OpenFile()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(79)
			p.OpenFolder()
		}

	}

	return localctx
}

// IOpenFileContext is an interface to support dynamic dispatch.
type IOpenFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOpenFileContext differentiates from other interfaces.
	IsOpenFileContext()
}

type OpenFileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpenFileContext() *OpenFileContext {
	var p = new(OpenFileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ImgParserRULE_openFile
	return p
}

func (*OpenFileContext) IsOpenFileContext() {}

func NewOpenFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpenFileContext {
	var p = new(OpenFileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ImgParserRULE_openFile

	return p
}

func (s *OpenFileContext) GetParser() antlr.Parser { return s.parser }

func (s *OpenFileContext) ID() antlr.TerminalNode {
	return s.GetToken(ImgParserID, 0)
}

func (s *OpenFileContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(ImgParserASSIGN, 0)
}

func (s *OpenFileContext) OPEN() antlr.TerminalNode {
	return s.GetToken(ImgParserOPEN, 0)
}

func (s *OpenFileContext) OPEN_PARAN() antlr.TerminalNode {
	return s.GetToken(ImgParserOPEN_PARAN, 0)
}

func (s *OpenFileContext) FileName() IFileNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFileNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFileNameContext)
}

func (s *OpenFileContext) CLOSED_PARAN() antlr.TerminalNode {
	return s.GetToken(ImgParserCLOSED_PARAN, 0)
}

func (s *OpenFileContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ImgParserSEMICOLON, 0)
}

func (s *OpenFileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpenFileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpenFileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.EnterOpenFile(s)
	}
}

func (s *OpenFileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.ExitOpenFile(s)
	}
}

func (p *ImgParser) OpenFile() (localctx IOpenFileContext) {
	localctx = NewOpenFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ImgParserRULE_openFile)

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
		p.SetState(82)
		p.Match(ImgParserID)
	}
	{
		p.SetState(83)
		p.Match(ImgParserASSIGN)
	}
	{
		p.SetState(84)
		p.Match(ImgParserOPEN)
	}
	{
		p.SetState(85)
		p.Match(ImgParserOPEN_PARAN)
	}
	{
		p.SetState(86)
		p.FileName()
	}
	{
		p.SetState(87)
		p.Match(ImgParserCLOSED_PARAN)
	}
	{
		p.SetState(88)
		p.Match(ImgParserSEMICOLON)
	}

	return localctx
}

// IOpenFolderContext is an interface to support dynamic dispatch.
type IOpenFolderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOpenFolderContext differentiates from other interfaces.
	IsOpenFolderContext()
}

type OpenFolderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpenFolderContext() *OpenFolderContext {
	var p = new(OpenFolderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ImgParserRULE_openFolder
	return p
}

func (*OpenFolderContext) IsOpenFolderContext() {}

func NewOpenFolderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpenFolderContext {
	var p = new(OpenFolderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ImgParserRULE_openFolder

	return p
}

func (s *OpenFolderContext) GetParser() antlr.Parser { return s.parser }

func (s *OpenFolderContext) ID() antlr.TerminalNode {
	return s.GetToken(ImgParserID, 0)
}

func (s *OpenFolderContext) FOLDER() antlr.TerminalNode {
	return s.GetToken(ImgParserFOLDER, 0)
}

func (s *OpenFolderContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(ImgParserASSIGN, 0)
}

func (s *OpenFolderContext) OPEN() antlr.TerminalNode {
	return s.GetToken(ImgParserOPEN, 0)
}

func (s *OpenFolderContext) OPEN_PARAN() antlr.TerminalNode {
	return s.GetToken(ImgParserOPEN_PARAN, 0)
}

func (s *OpenFolderContext) FolderName() IFolderNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFolderNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFolderNameContext)
}

func (s *OpenFolderContext) CLOSED_PARAN() antlr.TerminalNode {
	return s.GetToken(ImgParserCLOSED_PARAN, 0)
}

func (s *OpenFolderContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ImgParserSEMICOLON, 0)
}

func (s *OpenFolderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpenFolderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpenFolderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.EnterOpenFolder(s)
	}
}

func (s *OpenFolderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.ExitOpenFolder(s)
	}
}

func (p *ImgParser) OpenFolder() (localctx IOpenFolderContext) {
	localctx = NewOpenFolderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ImgParserRULE_openFolder)

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
		p.SetState(90)
		p.Match(ImgParserID)
	}
	{
		p.SetState(91)
		p.Match(ImgParserFOLDER)
	}
	{
		p.SetState(92)
		p.Match(ImgParserASSIGN)
	}
	{
		p.SetState(93)
		p.Match(ImgParserOPEN)
	}
	{
		p.SetState(94)
		p.Match(ImgParserOPEN_PARAN)
	}
	{
		p.SetState(95)
		p.FolderName()
	}
	{
		p.SetState(96)
		p.Match(ImgParserCLOSED_PARAN)
	}
	{
		p.SetState(97)
		p.Match(ImgParserSEMICOLON)
	}

	return localctx
}

// IFileNameContext is an interface to support dynamic dispatch.
type IFileNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFileNameContext differentiates from other interfaces.
	IsFileNameContext()
}

type FileNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileNameContext() *FileNameContext {
	var p = new(FileNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ImgParserRULE_fileName
	return p
}

func (*FileNameContext) IsFileNameContext() {}

func NewFileNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileNameContext {
	var p = new(FileNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ImgParserRULE_fileName

	return p
}

func (s *FileNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FileNameContext) AllAPOS() []antlr.TerminalNode {
	return s.GetTokens(ImgParserAPOS)
}

func (s *FileNameContext) APOS(i int) antlr.TerminalNode {
	return s.GetToken(ImgParserAPOS, i)
}

func (s *FileNameContext) ID() antlr.TerminalNode {
	return s.GetToken(ImgParserID, 0)
}

func (s *FileNameContext) DOT() antlr.TerminalNode {
	return s.GetToken(ImgParserDOT, 0)
}

func (s *FileNameContext) ImageType() IImageTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImageTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImageTypeContext)
}

func (s *FileNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.EnterFileName(s)
	}
}

func (s *FileNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.ExitFileName(s)
	}
}

func (p *ImgParser) FileName() (localctx IFileNameContext) {
	localctx = NewFileNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ImgParserRULE_fileName)

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
		p.SetState(99)
		p.Match(ImgParserAPOS)
	}
	{
		p.SetState(100)
		p.Match(ImgParserID)
	}
	{
		p.SetState(101)
		p.Match(ImgParserDOT)
	}
	{
		p.SetState(102)
		p.ImageType()
	}
	{
		p.SetState(103)
		p.Match(ImgParserAPOS)
	}

	return localctx
}

// IImageTypeContext is an interface to support dynamic dispatch.
type IImageTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImageTypeContext differentiates from other interfaces.
	IsImageTypeContext()
}

type ImageTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImageTypeContext() *ImageTypeContext {
	var p = new(ImageTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ImgParserRULE_imageType
	return p
}

func (*ImageTypeContext) IsImageTypeContext() {}

func NewImageTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImageTypeContext {
	var p = new(ImageTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ImgParserRULE_imageType

	return p
}

func (s *ImageTypeContext) GetParser() antlr.Parser { return s.parser }
func (s *ImageTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImageTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImageTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.EnterImageType(s)
	}
}

func (s *ImageTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.ExitImageType(s)
	}
}

func (p *ImgParser) ImageType() (localctx IImageTypeContext) {
	localctx = NewImageTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ImgParserRULE_imageType)
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
	p.SetState(105)
	_la = p.GetTokenStream().LA(1)

	if !(_la == ImgParserT__0 || _la == ImgParserT__1) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IFolderNameContext is an interface to support dynamic dispatch.
type IFolderNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFolderNameContext differentiates from other interfaces.
	IsFolderNameContext()
}

type FolderNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFolderNameContext() *FolderNameContext {
	var p = new(FolderNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ImgParserRULE_folderName
	return p
}

func (*FolderNameContext) IsFolderNameContext() {}

func NewFolderNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FolderNameContext {
	var p = new(FolderNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ImgParserRULE_folderName

	return p
}

func (s *FolderNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FolderNameContext) AllAPOS() []antlr.TerminalNode {
	return s.GetTokens(ImgParserAPOS)
}

func (s *FolderNameContext) APOS(i int) antlr.TerminalNode {
	return s.GetToken(ImgParserAPOS, i)
}

func (s *FolderNameContext) ID() antlr.TerminalNode {
	return s.GetToken(ImgParserID, 0)
}

func (s *FolderNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FolderNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FolderNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.EnterFolderName(s)
	}
}

func (s *FolderNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.ExitFolderName(s)
	}
}

func (p *ImgParser) FolderName() (localctx IFolderNameContext) {
	localctx = NewFolderNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ImgParserRULE_folderName)

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
		p.SetState(107)
		p.Match(ImgParserAPOS)
	}
	{
		p.SetState(108)
		p.Match(ImgParserID)
	}
	{
		p.SetState(109)
		p.Match(ImgParserAPOS)
	}

	return localctx
}

// IAction_Context is an interface to support dynamic dispatch.
type IAction_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAction_Context differentiates from other interfaces.
	IsAction_Context()
}

type Action_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAction_Context() *Action_Context {
	var p = new(Action_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ImgParserRULE_action_
	return p
}

func (*Action_Context) IsAction_Context() {}

func NewAction_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Action_Context {
	var p = new(Action_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ImgParserRULE_action_

	return p
}

func (s *Action_Context) GetParser() antlr.Parser { return s.parser }

func (s *Action_Context) ID() antlr.TerminalNode {
	return s.GetToken(ImgParserID, 0)
}

func (s *Action_Context) DOT() antlr.TerminalNode {
	return s.GetToken(ImgParserDOT, 0)
}

func (s *Action_Context) ActionType() IActionTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IActionTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IActionTypeContext)
}

func (s *Action_Context) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ImgParserSEMICOLON, 0)
}

func (s *Action_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Action_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Action_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.EnterAction_(s)
	}
}

func (s *Action_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.ExitAction_(s)
	}
}

func (p *ImgParser) Action_() (localctx IAction_Context) {
	localctx = NewAction_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ImgParserRULE_action_)

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
		p.SetState(111)
		p.Match(ImgParserID)
	}
	{
		p.SetState(112)
		p.Match(ImgParserDOT)
	}
	{
		p.SetState(113)
		p.ActionType()
	}
	{
		p.SetState(114)
		p.Match(ImgParserSEMICOLON)
	}

	return localctx
}

// IActionTypeContext is an interface to support dynamic dispatch.
type IActionTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsActionTypeContext differentiates from other interfaces.
	IsActionTypeContext()
}

type ActionTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionTypeContext() *ActionTypeContext {
	var p = new(ActionTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ImgParserRULE_actionType
	return p
}

func (*ActionTypeContext) IsActionTypeContext() {}

func NewActionTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionTypeContext {
	var p = new(ActionTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ImgParserRULE_actionType

	return p
}

func (s *ActionTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionTypeContext) Crop() ICropContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICropContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICropContext)
}

func (s *ActionTypeContext) Rotate() IRotateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRotateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRotateContext)
}

func (s *ActionTypeContext) FlipX() IFlipXContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFlipXContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFlipXContext)
}

func (s *ActionTypeContext) FlipY() IFlipYContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFlipYContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFlipYContext)
}

func (s *ActionTypeContext) Resize() IResizeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IResizeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IResizeContext)
}

func (s *ActionTypeContext) Brightness() IBrightnessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBrightnessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBrightnessContext)
}

func (s *ActionTypeContext) Contrast() IContrastContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContrastContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContrastContext)
}

func (s *ActionTypeContext) Saturation() ISaturationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISaturationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISaturationContext)
}

func (s *ActionTypeContext) Pixelate() IPixelateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPixelateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPixelateContext)
}

func (s *ActionTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.EnterActionType(s)
	}
}

func (s *ActionTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.ExitActionType(s)
	}
}

func (p *ImgParser) ActionType() (localctx IActionTypeContext) {
	localctx = NewActionTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ImgParserRULE_actionType)

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

	p.SetState(125)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ImgParserT__2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(116)
			p.Crop()
		}

	case ImgParserT__3:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(117)
			p.Rotate()
		}

	case ImgParserT__4:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(118)
			p.FlipX()
		}

	case ImgParserT__5:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(119)
			p.FlipY()
		}

	case ImgParserT__6:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(120)
			p.Resize()
		}

	case ImgParserT__7:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(121)
			p.Brightness()
		}

	case ImgParserT__8:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(122)
			p.Contrast()
		}

	case ImgParserT__9:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(123)
			p.Saturation()
		}

	case ImgParserT__10:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(124)
			p.Pixelate()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICropContext is an interface to support dynamic dispatch.
type ICropContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCropContext differentiates from other interfaces.
	IsCropContext()
}

type CropContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCropContext() *CropContext {
	var p = new(CropContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ImgParserRULE_crop
	return p
}

func (*CropContext) IsCropContext() {}

func NewCropContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CropContext {
	var p = new(CropContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ImgParserRULE_crop

	return p
}

func (s *CropContext) GetParser() antlr.Parser { return s.parser }

func (s *CropContext) OPEN_PARAN() antlr.TerminalNode {
	return s.GetToken(ImgParserOPEN_PARAN, 0)
}

func (s *CropContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(ImgParserNUMBER)
}

func (s *CropContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(ImgParserNUMBER, i)
}

func (s *CropContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ImgParserCOMMA)
}

func (s *CropContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ImgParserCOMMA, i)
}

func (s *CropContext) CLOSED_PARAN() antlr.TerminalNode {
	return s.GetToken(ImgParserCLOSED_PARAN, 0)
}

func (s *CropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CropContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.EnterCrop(s)
	}
}

func (s *CropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.ExitCrop(s)
	}
}

func (p *ImgParser) Crop() (localctx ICropContext) {
	localctx = NewCropContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ImgParserRULE_crop)

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
		p.SetState(127)
		p.Match(ImgParserT__2)
	}
	{
		p.SetState(128)
		p.Match(ImgParserOPEN_PARAN)
	}
	{
		p.SetState(129)
		p.Match(ImgParserNUMBER)
	}
	{
		p.SetState(130)
		p.Match(ImgParserCOMMA)
	}
	{
		p.SetState(131)
		p.Match(ImgParserNUMBER)
	}
	{
		p.SetState(132)
		p.Match(ImgParserCOMMA)
	}
	{
		p.SetState(133)
		p.Match(ImgParserNUMBER)
	}
	{
		p.SetState(134)
		p.Match(ImgParserCOMMA)
	}
	{
		p.SetState(135)
		p.Match(ImgParserNUMBER)
	}
	{
		p.SetState(136)
		p.Match(ImgParserCLOSED_PARAN)
	}

	return localctx
}

// IRotateContext is an interface to support dynamic dispatch.
type IRotateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRotateContext differentiates from other interfaces.
	IsRotateContext()
}

type RotateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRotateContext() *RotateContext {
	var p = new(RotateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ImgParserRULE_rotate
	return p
}

func (*RotateContext) IsRotateContext() {}

func NewRotateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RotateContext {
	var p = new(RotateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ImgParserRULE_rotate

	return p
}

func (s *RotateContext) GetParser() antlr.Parser { return s.parser }

func (s *RotateContext) OPEN_PARAN() antlr.TerminalNode {
	return s.GetToken(ImgParserOPEN_PARAN, 0)
}

func (s *RotateContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ImgParserNUMBER, 0)
}

func (s *RotateContext) CLOSED_PARAN() antlr.TerminalNode {
	return s.GetToken(ImgParserCLOSED_PARAN, 0)
}

func (s *RotateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RotateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RotateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.EnterRotate(s)
	}
}

func (s *RotateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.ExitRotate(s)
	}
}

func (p *ImgParser) Rotate() (localctx IRotateContext) {
	localctx = NewRotateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ImgParserRULE_rotate)

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
		p.SetState(138)
		p.Match(ImgParserT__3)
	}
	{
		p.SetState(139)
		p.Match(ImgParserOPEN_PARAN)
	}
	{
		p.SetState(140)
		p.Match(ImgParserNUMBER)
	}
	{
		p.SetState(141)
		p.Match(ImgParserCLOSED_PARAN)
	}

	return localctx
}

// IFlipXContext is an interface to support dynamic dispatch.
type IFlipXContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFlipXContext differentiates from other interfaces.
	IsFlipXContext()
}

type FlipXContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFlipXContext() *FlipXContext {
	var p = new(FlipXContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ImgParserRULE_flipX
	return p
}

func (*FlipXContext) IsFlipXContext() {}

func NewFlipXContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FlipXContext {
	var p = new(FlipXContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ImgParserRULE_flipX

	return p
}

func (s *FlipXContext) GetParser() antlr.Parser { return s.parser }

func (s *FlipXContext) OPEN_PARAN() antlr.TerminalNode {
	return s.GetToken(ImgParserOPEN_PARAN, 0)
}

func (s *FlipXContext) CLOSED_PARAN() antlr.TerminalNode {
	return s.GetToken(ImgParserCLOSED_PARAN, 0)
}

func (s *FlipXContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FlipXContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FlipXContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.EnterFlipX(s)
	}
}

func (s *FlipXContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.ExitFlipX(s)
	}
}

func (p *ImgParser) FlipX() (localctx IFlipXContext) {
	localctx = NewFlipXContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ImgParserRULE_flipX)

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
		p.SetState(143)
		p.Match(ImgParserT__4)
	}
	{
		p.SetState(144)
		p.Match(ImgParserOPEN_PARAN)
	}
	{
		p.SetState(145)
		p.Match(ImgParserCLOSED_PARAN)
	}

	return localctx
}

// IFlipYContext is an interface to support dynamic dispatch.
type IFlipYContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFlipYContext differentiates from other interfaces.
	IsFlipYContext()
}

type FlipYContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFlipYContext() *FlipYContext {
	var p = new(FlipYContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ImgParserRULE_flipY
	return p
}

func (*FlipYContext) IsFlipYContext() {}

func NewFlipYContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FlipYContext {
	var p = new(FlipYContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ImgParserRULE_flipY

	return p
}

func (s *FlipYContext) GetParser() antlr.Parser { return s.parser }

func (s *FlipYContext) OPEN_PARAN() antlr.TerminalNode {
	return s.GetToken(ImgParserOPEN_PARAN, 0)
}

func (s *FlipYContext) CLOSED_PARAN() antlr.TerminalNode {
	return s.GetToken(ImgParserCLOSED_PARAN, 0)
}

func (s *FlipYContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FlipYContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FlipYContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.EnterFlipY(s)
	}
}

func (s *FlipYContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.ExitFlipY(s)
	}
}

func (p *ImgParser) FlipY() (localctx IFlipYContext) {
	localctx = NewFlipYContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ImgParserRULE_flipY)

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
		p.SetState(147)
		p.Match(ImgParserT__5)
	}
	{
		p.SetState(148)
		p.Match(ImgParserOPEN_PARAN)
	}
	{
		p.SetState(149)
		p.Match(ImgParserCLOSED_PARAN)
	}

	return localctx
}

// IResizeContext is an interface to support dynamic dispatch.
type IResizeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsResizeContext differentiates from other interfaces.
	IsResizeContext()
}

type ResizeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResizeContext() *ResizeContext {
	var p = new(ResizeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ImgParserRULE_resize
	return p
}

func (*ResizeContext) IsResizeContext() {}

func NewResizeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResizeContext {
	var p = new(ResizeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ImgParserRULE_resize

	return p
}

func (s *ResizeContext) GetParser() antlr.Parser { return s.parser }

func (s *ResizeContext) OPEN_PARAN() antlr.TerminalNode {
	return s.GetToken(ImgParserOPEN_PARAN, 0)
}

func (s *ResizeContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(ImgParserNUMBER)
}

func (s *ResizeContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(ImgParserNUMBER, i)
}

func (s *ResizeContext) CLOSED_PARAN() antlr.TerminalNode {
	return s.GetToken(ImgParserCLOSED_PARAN, 0)
}

func (s *ResizeContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ImgParserCOMMA, 0)
}

func (s *ResizeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResizeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResizeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.EnterResize(s)
	}
}

func (s *ResizeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.ExitResize(s)
	}
}

func (p *ImgParser) Resize() (localctx IResizeContext) {
	localctx = NewResizeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ImgParserRULE_resize)
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
		p.SetState(151)
		p.Match(ImgParserT__6)
	}
	{
		p.SetState(152)
		p.Match(ImgParserOPEN_PARAN)
	}
	{
		p.SetState(153)
		p.Match(ImgParserNUMBER)
	}
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ImgParserCOMMA {
		{
			p.SetState(154)
			p.Match(ImgParserCOMMA)
		}
		{
			p.SetState(155)
			p.Match(ImgParserNUMBER)
		}

	}
	{
		p.SetState(158)
		p.Match(ImgParserCLOSED_PARAN)
	}

	return localctx
}

// IBrightnessContext is an interface to support dynamic dispatch.
type IBrightnessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBrightnessContext differentiates from other interfaces.
	IsBrightnessContext()
}

type BrightnessContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBrightnessContext() *BrightnessContext {
	var p = new(BrightnessContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ImgParserRULE_brightness
	return p
}

func (*BrightnessContext) IsBrightnessContext() {}

func NewBrightnessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BrightnessContext {
	var p = new(BrightnessContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ImgParserRULE_brightness

	return p
}

func (s *BrightnessContext) GetParser() antlr.Parser { return s.parser }

func (s *BrightnessContext) OPEN_PARAN() antlr.TerminalNode {
	return s.GetToken(ImgParserOPEN_PARAN, 0)
}

func (s *BrightnessContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ImgParserNUMBER, 0)
}

func (s *BrightnessContext) CLOSED_PARAN() antlr.TerminalNode {
	return s.GetToken(ImgParserCLOSED_PARAN, 0)
}

func (s *BrightnessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BrightnessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BrightnessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.EnterBrightness(s)
	}
}

func (s *BrightnessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.ExitBrightness(s)
	}
}

func (p *ImgParser) Brightness() (localctx IBrightnessContext) {
	localctx = NewBrightnessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ImgParserRULE_brightness)

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
		p.SetState(160)
		p.Match(ImgParserT__7)
	}
	{
		p.SetState(161)
		p.Match(ImgParserOPEN_PARAN)
	}
	{
		p.SetState(162)
		p.Match(ImgParserNUMBER)
	}
	{
		p.SetState(163)
		p.Match(ImgParserCLOSED_PARAN)
	}

	return localctx
}

// IContrastContext is an interface to support dynamic dispatch.
type IContrastContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContrastContext differentiates from other interfaces.
	IsContrastContext()
}

type ContrastContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContrastContext() *ContrastContext {
	var p = new(ContrastContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ImgParserRULE_contrast
	return p
}

func (*ContrastContext) IsContrastContext() {}

func NewContrastContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContrastContext {
	var p = new(ContrastContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ImgParserRULE_contrast

	return p
}

func (s *ContrastContext) GetParser() antlr.Parser { return s.parser }

func (s *ContrastContext) OPEN_PARAN() antlr.TerminalNode {
	return s.GetToken(ImgParserOPEN_PARAN, 0)
}

func (s *ContrastContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ImgParserNUMBER, 0)
}

func (s *ContrastContext) CLOSED_PARAN() antlr.TerminalNode {
	return s.GetToken(ImgParserCLOSED_PARAN, 0)
}

func (s *ContrastContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContrastContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContrastContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.EnterContrast(s)
	}
}

func (s *ContrastContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.ExitContrast(s)
	}
}

func (p *ImgParser) Contrast() (localctx IContrastContext) {
	localctx = NewContrastContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ImgParserRULE_contrast)

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
		p.SetState(165)
		p.Match(ImgParserT__8)
	}
	{
		p.SetState(166)
		p.Match(ImgParserOPEN_PARAN)
	}
	{
		p.SetState(167)
		p.Match(ImgParserNUMBER)
	}
	{
		p.SetState(168)
		p.Match(ImgParserCLOSED_PARAN)
	}

	return localctx
}

// ISaturationContext is an interface to support dynamic dispatch.
type ISaturationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSaturationContext differentiates from other interfaces.
	IsSaturationContext()
}

type SaturationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySaturationContext() *SaturationContext {
	var p = new(SaturationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ImgParserRULE_saturation
	return p
}

func (*SaturationContext) IsSaturationContext() {}

func NewSaturationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SaturationContext {
	var p = new(SaturationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ImgParserRULE_saturation

	return p
}

func (s *SaturationContext) GetParser() antlr.Parser { return s.parser }

func (s *SaturationContext) OPEN_PARAN() antlr.TerminalNode {
	return s.GetToken(ImgParserOPEN_PARAN, 0)
}

func (s *SaturationContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ImgParserNUMBER, 0)
}

func (s *SaturationContext) CLOSED_PARAN() antlr.TerminalNode {
	return s.GetToken(ImgParserCLOSED_PARAN, 0)
}

func (s *SaturationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SaturationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SaturationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.EnterSaturation(s)
	}
}

func (s *SaturationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.ExitSaturation(s)
	}
}

func (p *ImgParser) Saturation() (localctx ISaturationContext) {
	localctx = NewSaturationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ImgParserRULE_saturation)

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
		p.SetState(170)
		p.Match(ImgParserT__9)
	}
	{
		p.SetState(171)
		p.Match(ImgParserOPEN_PARAN)
	}
	{
		p.SetState(172)
		p.Match(ImgParserNUMBER)
	}
	{
		p.SetState(173)
		p.Match(ImgParserCLOSED_PARAN)
	}

	return localctx
}

// IPixelateContext is an interface to support dynamic dispatch.
type IPixelateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPixelateContext differentiates from other interfaces.
	IsPixelateContext()
}

type PixelateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPixelateContext() *PixelateContext {
	var p = new(PixelateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ImgParserRULE_pixelate
	return p
}

func (*PixelateContext) IsPixelateContext() {}

func NewPixelateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PixelateContext {
	var p = new(PixelateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ImgParserRULE_pixelate

	return p
}

func (s *PixelateContext) GetParser() antlr.Parser { return s.parser }

func (s *PixelateContext) OPEN_PARAN() antlr.TerminalNode {
	return s.GetToken(ImgParserOPEN_PARAN, 0)
}

func (s *PixelateContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ImgParserNUMBER, 0)
}

func (s *PixelateContext) CLOSED_PARAN() antlr.TerminalNode {
	return s.GetToken(ImgParserCLOSED_PARAN, 0)
}

func (s *PixelateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PixelateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PixelateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.EnterPixelate(s)
	}
}

func (s *PixelateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.ExitPixelate(s)
	}
}

func (p *ImgParser) Pixelate() (localctx IPixelateContext) {
	localctx = NewPixelateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ImgParserRULE_pixelate)

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
		p.SetState(175)
		p.Match(ImgParserT__10)
	}
	{
		p.SetState(176)
		p.Match(ImgParserOPEN_PARAN)
	}
	{
		p.SetState(177)
		p.Match(ImgParserNUMBER)
	}
	{
		p.SetState(178)
		p.Match(ImgParserCLOSED_PARAN)
	}

	return localctx
}

// IExportContext is an interface to support dynamic dispatch.
type IExportContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExportContext differentiates from other interfaces.
	IsExportContext()
}

type ExportContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExportContext() *ExportContext {
	var p = new(ExportContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ImgParserRULE_export
	return p
}

func (*ExportContext) IsExportContext() {}

func NewExportContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExportContext {
	var p = new(ExportContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ImgParserRULE_export

	return p
}

func (s *ExportContext) GetParser() antlr.Parser { return s.parser }

func (s *ExportContext) ID() antlr.TerminalNode {
	return s.GetToken(ImgParserID, 0)
}

func (s *ExportContext) DOT() antlr.TerminalNode {
	return s.GetToken(ImgParserDOT, 0)
}

func (s *ExportContext) OPEN_PARAN() antlr.TerminalNode {
	return s.GetToken(ImgParserOPEN_PARAN, 0)
}

func (s *ExportContext) CLOSED_PARAN() antlr.TerminalNode {
	return s.GetToken(ImgParserCLOSED_PARAN, 0)
}

func (s *ExportContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ImgParserSEMICOLON, 0)
}

func (s *ExportContext) FileName() IFileNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFileNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFileNameContext)
}

func (s *ExportContext) FolderName() IFolderNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFolderNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFolderNameContext)
}

func (s *ExportContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExportContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExportContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.EnterExport(s)
	}
}

func (s *ExportContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.ExitExport(s)
	}
}

func (p *ImgParser) Export() (localctx IExportContext) {
	localctx = NewExportContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ImgParserRULE_export)

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
		p.SetState(180)
		p.Match(ImgParserID)
	}
	{
		p.SetState(181)
		p.Match(ImgParserDOT)
	}
	{
		p.SetState(182)
		p.Match(ImgParserT__11)
	}
	{
		p.SetState(183)
		p.Match(ImgParserOPEN_PARAN)
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(184)
			p.FileName()
		}

	case 2:
		{
			p.SetState(185)
			p.FolderName()
		}

	}
	{
		p.SetState(188)
		p.Match(ImgParserCLOSED_PARAN)
	}
	{
		p.SetState(189)
		p.Match(ImgParserSEMICOLON)
	}

	return localctx
}

// ITextContext is an interface to support dynamic dispatch.
type ITextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTextContext differentiates from other interfaces.
	IsTextContext()
}

type TextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTextContext() *TextContext {
	var p = new(TextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ImgParserRULE_text
	return p
}

func (*TextContext) IsTextContext() {}

func NewTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TextContext {
	var p = new(TextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ImgParserRULE_text

	return p
}

func (s *TextContext) GetParser() antlr.Parser { return s.parser }

func (s *TextContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ImgParserID)
}

func (s *TextContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ImgParserID, i)
}

func (s *TextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.EnterText(s)
	}
}

func (s *TextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.ExitText(s)
	}
}

func (p *ImgParser) Text() (localctx ITextContext) {
	localctx = NewTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ImgParserRULE_text)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(191)
				p.Match(ImgParserID)
			}

		}
		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}

	return localctx
}

// ICommentContext is an interface to support dynamic dispatch.
type ICommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommentContext differentiates from other interfaces.
	IsCommentContext()
}

type CommentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentContext() *CommentContext {
	var p = new(CommentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ImgParserRULE_comment
	return p
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ImgParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) Text() ITextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITextContext)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImgListener); ok {
		listenerT.ExitComment(s)
	}
}

func (p *ImgParser) Comment() (localctx ICommentContext) {
	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ImgParserRULE_comment)

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
		p.SetState(197)
		p.Match(ImgParserT__12)
	}
	{
		p.SetState(198)
		p.Text()
	}

	return localctx
}
