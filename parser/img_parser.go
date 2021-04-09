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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 22, 133,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 3, 2, 6, 2, 34, 10, 2,
	13, 2, 14, 2, 35, 3, 2, 7, 2, 39, 10, 2, 12, 2, 14, 2, 42, 11, 2, 3, 2,
	3, 2, 3, 2, 3, 3, 3, 3, 5, 3, 49, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10,
	90, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5,
	15, 121, 10, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 2, 2, 17, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
	24, 26, 28, 30, 2, 3, 3, 2, 3, 4, 2, 125, 2, 33, 3, 2, 2, 2, 4, 48, 3,
	2, 2, 2, 6, 50, 3, 2, 2, 2, 8, 58, 3, 2, 2, 2, 10, 67, 3, 2, 2, 2, 12,
	73, 3, 2, 2, 2, 14, 75, 3, 2, 2, 2, 16, 79, 3, 2, 2, 2, 18, 89, 3, 2, 2,
	2, 20, 91, 3, 2, 2, 2, 22, 102, 3, 2, 2, 2, 24, 107, 3, 2, 2, 2, 26, 111,
	3, 2, 2, 2, 28, 115, 3, 2, 2, 2, 30, 124, 3, 2, 2, 2, 32, 34, 5, 4, 3,
	2, 33, 32, 3, 2, 2, 2, 34, 35, 3, 2, 2, 2, 35, 33, 3, 2, 2, 2, 35, 36,
	3, 2, 2, 2, 36, 40, 3, 2, 2, 2, 37, 39, 5, 16, 9, 2, 38, 37, 3, 2, 2, 2,
	39, 42, 3, 2, 2, 2, 40, 38, 3, 2, 2, 2, 40, 41, 3, 2, 2, 2, 41, 43, 3,
	2, 2, 2, 42, 40, 3, 2, 2, 2, 43, 44, 5, 30, 16, 2, 44, 45, 7, 2, 2, 3,
	45, 3, 3, 2, 2, 2, 46, 49, 5, 6, 4, 2, 47, 49, 5, 8, 5, 2, 48, 46, 3, 2,
	2, 2, 48, 47, 3, 2, 2, 2, 49, 5, 3, 2, 2, 2, 50, 51, 7, 21, 2, 2, 51, 52,
	7, 11, 2, 2, 52, 53, 7, 19, 2, 2, 53, 54, 7, 16, 2, 2, 54, 55, 5, 10, 6,
	2, 55, 56, 7, 17, 2, 2, 56, 57, 7, 15, 2, 2, 57, 7, 3, 2, 2, 2, 58, 59,
	7, 21, 2, 2, 59, 60, 7, 18, 2, 2, 60, 61, 7, 11, 2, 2, 61, 62, 7, 19, 2,
	2, 62, 63, 7, 16, 2, 2, 63, 64, 5, 14, 8, 2, 64, 65, 7, 17, 2, 2, 65, 66,
	7, 15, 2, 2, 66, 9, 3, 2, 2, 2, 67, 68, 7, 14, 2, 2, 68, 69, 7, 21, 2,
	2, 69, 70, 7, 12, 2, 2, 70, 71, 5, 12, 7, 2, 71, 72, 7, 14, 2, 2, 72, 11,
	3, 2, 2, 2, 73, 74, 9, 2, 2, 2, 74, 13, 3, 2, 2, 2, 75, 76, 7, 14, 2, 2,
	76, 77, 7, 21, 2, 2, 77, 78, 7, 14, 2, 2, 78, 15, 3, 2, 2, 2, 79, 80, 7,
	21, 2, 2, 80, 81, 7, 12, 2, 2, 81, 82, 5, 18, 10, 2, 82, 83, 7, 15, 2,
	2, 83, 17, 3, 2, 2, 2, 84, 90, 5, 20, 11, 2, 85, 90, 5, 22, 12, 2, 86,
	90, 5, 24, 13, 2, 87, 90, 5, 26, 14, 2, 88, 90, 5, 28, 15, 2, 89, 84, 3,
	2, 2, 2, 89, 85, 3, 2, 2, 2, 89, 86, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 89,
	88, 3, 2, 2, 2, 90, 19, 3, 2, 2, 2, 91, 92, 7, 5, 2, 2, 92, 93, 7, 16,
	2, 2, 93, 94, 7, 20, 2, 2, 94, 95, 7, 13, 2, 2, 95, 96, 7, 20, 2, 2, 96,
	97, 7, 13, 2, 2, 97, 98, 7, 20, 2, 2, 98, 99, 7, 13, 2, 2, 99, 100, 7,
	20, 2, 2, 100, 101, 7, 17, 2, 2, 101, 21, 3, 2, 2, 2, 102, 103, 7, 6, 2,
	2, 103, 104, 7, 16, 2, 2, 104, 105, 7, 20, 2, 2, 105, 106, 7, 17, 2, 2,
	106, 23, 3, 2, 2, 2, 107, 108, 7, 7, 2, 2, 108, 109, 7, 16, 2, 2, 109,
	110, 7, 17, 2, 2, 110, 25, 3, 2, 2, 2, 111, 112, 7, 8, 2, 2, 112, 113,
	7, 16, 2, 2, 113, 114, 7, 17, 2, 2, 114, 27, 3, 2, 2, 2, 115, 116, 7, 9,
	2, 2, 116, 117, 7, 16, 2, 2, 117, 120, 7, 20, 2, 2, 118, 119, 7, 13, 2,
	2, 119, 121, 7, 20, 2, 2, 120, 118, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121,
	122, 3, 2, 2, 2, 122, 123, 7, 17, 2, 2, 123, 29, 3, 2, 2, 2, 124, 125,
	7, 21, 2, 2, 125, 126, 7, 12, 2, 2, 126, 127, 7, 10, 2, 2, 127, 128, 7,
	16, 2, 2, 128, 129, 5, 10, 6, 2, 129, 130, 7, 17, 2, 2, 130, 131, 7, 15,
	2, 2, 131, 31, 3, 2, 2, 2, 7, 35, 40, 48, 89, 120,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'png'", "'jpg'", "'crop'", "'rotate'", "'flipX'", "'flipY'", "'resize'",
	"'save'", "'='", "'.'", "','", "'\"'", "';'", "'('", "')'", "'[]'", "'open'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "ASSIGN", "DOT", "COMMA", "APOS", "SEMICOLON",
	"OPEN_PARAN", "CLOSED_PARAN", "FOLDER", "OPEN", "NUMBER", "IDENT", "WHITESPACE",
}

var ruleNames = []string{
	"start", "declaration", "openFile", "openFolder", "fileName", "imageType",
	"folderName", "action_", "actionType", "crop", "rotate", "flipX", "flipY",
	"resize", "export",
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
	ImgParserASSIGN       = 9
	ImgParserDOT          = 10
	ImgParserCOMMA        = 11
	ImgParserAPOS         = 12
	ImgParserSEMICOLON    = 13
	ImgParserOPEN_PARAN   = 14
	ImgParserCLOSED_PARAN = 15
	ImgParserFOLDER       = 16
	ImgParserOPEN         = 17
	ImgParserNUMBER       = 18
	ImgParserIDENT        = 19
	ImgParserWHITESPACE   = 20
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
	ImgParserRULE_export      = 14
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

func (s *StartContext) Export() IExportContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExportContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExportContext)
}

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

func (s *StartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ImgVisitor:
		return t.VisitStart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ImgParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ImgParserRULE_start)

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
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(30)
				p.Declaration()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(35)
				p.Action_()
			}

		}
		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}
	{
		p.SetState(41)
		p.Export()
	}
	{
		p.SetState(42)
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

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ImgVisitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
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

	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(44)
			p.OpenFile()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(45)
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

func (s *OpenFileContext) IDENT() antlr.TerminalNode {
	return s.GetToken(ImgParserIDENT, 0)
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

func (s *OpenFileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ImgVisitor:
		return t.VisitOpenFile(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(48)
		p.Match(ImgParserIDENT)
	}
	{
		p.SetState(49)
		p.Match(ImgParserASSIGN)
	}
	{
		p.SetState(50)
		p.Match(ImgParserOPEN)
	}
	{
		p.SetState(51)
		p.Match(ImgParserOPEN_PARAN)
	}
	{
		p.SetState(52)
		p.FileName()
	}
	{
		p.SetState(53)
		p.Match(ImgParserCLOSED_PARAN)
	}
	{
		p.SetState(54)
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

func (s *OpenFolderContext) IDENT() antlr.TerminalNode {
	return s.GetToken(ImgParserIDENT, 0)
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

func (s *OpenFolderContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ImgVisitor:
		return t.VisitOpenFolder(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(56)
		p.Match(ImgParserIDENT)
	}
	{
		p.SetState(57)
		p.Match(ImgParserFOLDER)
	}
	{
		p.SetState(58)
		p.Match(ImgParserASSIGN)
	}
	{
		p.SetState(59)
		p.Match(ImgParserOPEN)
	}
	{
		p.SetState(60)
		p.Match(ImgParserOPEN_PARAN)
	}
	{
		p.SetState(61)
		p.FolderName()
	}
	{
		p.SetState(62)
		p.Match(ImgParserCLOSED_PARAN)
	}
	{
		p.SetState(63)
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

func (s *FileNameContext) IDENT() antlr.TerminalNode {
	return s.GetToken(ImgParserIDENT, 0)
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

func (s *FileNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ImgVisitor:
		return t.VisitFileName(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(65)
		p.Match(ImgParserAPOS)
	}
	{
		p.SetState(66)
		p.Match(ImgParserIDENT)
	}
	{
		p.SetState(67)
		p.Match(ImgParserDOT)
	}
	{
		p.SetState(68)
		p.ImageType()
	}
	{
		p.SetState(69)
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

func (s *ImageTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ImgVisitor:
		return t.VisitImageType(s)

	default:
		return t.VisitChildren(s)
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
	p.SetState(71)
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

func (s *FolderNameContext) IDENT() antlr.TerminalNode {
	return s.GetToken(ImgParserIDENT, 0)
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

func (s *FolderNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ImgVisitor:
		return t.VisitFolderName(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(73)
		p.Match(ImgParserAPOS)
	}
	{
		p.SetState(74)
		p.Match(ImgParserIDENT)
	}
	{
		p.SetState(75)
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

func (s *Action_Context) IDENT() antlr.TerminalNode {
	return s.GetToken(ImgParserIDENT, 0)
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

func (s *Action_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ImgVisitor:
		return t.VisitAction_(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(77)
		p.Match(ImgParserIDENT)
	}
	{
		p.SetState(78)
		p.Match(ImgParserDOT)
	}
	{
		p.SetState(79)
		p.ActionType()
	}
	{
		p.SetState(80)
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

func (s *ActionTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ImgVisitor:
		return t.VisitActionType(s)

	default:
		return t.VisitChildren(s)
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

	p.SetState(87)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ImgParserT__2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(82)
			p.Crop()
		}

	case ImgParserT__3:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(83)
			p.Rotate()
		}

	case ImgParserT__4:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(84)
			p.FlipX()
		}

	case ImgParserT__5:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(85)
			p.FlipY()
		}

	case ImgParserT__6:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(86)
			p.Resize()
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

func (s *CropContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ImgVisitor:
		return t.VisitCrop(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(89)
		p.Match(ImgParserT__2)
	}
	{
		p.SetState(90)
		p.Match(ImgParserOPEN_PARAN)
	}
	{
		p.SetState(91)
		p.Match(ImgParserNUMBER)
	}
	{
		p.SetState(92)
		p.Match(ImgParserCOMMA)
	}
	{
		p.SetState(93)
		p.Match(ImgParserNUMBER)
	}
	{
		p.SetState(94)
		p.Match(ImgParserCOMMA)
	}
	{
		p.SetState(95)
		p.Match(ImgParserNUMBER)
	}
	{
		p.SetState(96)
		p.Match(ImgParserCOMMA)
	}
	{
		p.SetState(97)
		p.Match(ImgParserNUMBER)
	}
	{
		p.SetState(98)
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

func (s *RotateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ImgVisitor:
		return t.VisitRotate(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(100)
		p.Match(ImgParserT__3)
	}
	{
		p.SetState(101)
		p.Match(ImgParserOPEN_PARAN)
	}
	{
		p.SetState(102)
		p.Match(ImgParserNUMBER)
	}
	{
		p.SetState(103)
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

func (s *FlipXContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ImgVisitor:
		return t.VisitFlipX(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(105)
		p.Match(ImgParserT__4)
	}
	{
		p.SetState(106)
		p.Match(ImgParserOPEN_PARAN)
	}
	{
		p.SetState(107)
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

func (s *FlipYContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ImgVisitor:
		return t.VisitFlipY(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(109)
		p.Match(ImgParserT__5)
	}
	{
		p.SetState(110)
		p.Match(ImgParserOPEN_PARAN)
	}
	{
		p.SetState(111)
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

func (s *ResizeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ImgVisitor:
		return t.VisitResize(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(113)
		p.Match(ImgParserT__6)
	}
	{
		p.SetState(114)
		p.Match(ImgParserOPEN_PARAN)
	}
	{
		p.SetState(115)
		p.Match(ImgParserNUMBER)
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ImgParserCOMMA {
		{
			p.SetState(116)
			p.Match(ImgParserCOMMA)
		}
		{
			p.SetState(117)
			p.Match(ImgParserNUMBER)
		}

	}
	{
		p.SetState(120)
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

func (s *ExportContext) IDENT() antlr.TerminalNode {
	return s.GetToken(ImgParserIDENT, 0)
}

func (s *ExportContext) DOT() antlr.TerminalNode {
	return s.GetToken(ImgParserDOT, 0)
}

func (s *ExportContext) OPEN_PARAN() antlr.TerminalNode {
	return s.GetToken(ImgParserOPEN_PARAN, 0)
}

func (s *ExportContext) FileName() IFileNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFileNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFileNameContext)
}

func (s *ExportContext) CLOSED_PARAN() antlr.TerminalNode {
	return s.GetToken(ImgParserCLOSED_PARAN, 0)
}

func (s *ExportContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ImgParserSEMICOLON, 0)
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

func (s *ExportContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ImgVisitor:
		return t.VisitExport(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ImgParser) Export() (localctx IExportContext) {
	localctx = NewExportContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ImgParserRULE_export)

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
		p.SetState(122)
		p.Match(ImgParserIDENT)
	}
	{
		p.SetState(123)
		p.Match(ImgParserDOT)
	}
	{
		p.SetState(124)
		p.Match(ImgParserT__7)
	}
	{
		p.SetState(125)
		p.Match(ImgParserOPEN_PARAN)
	}
	{
		p.SetState(126)
		p.FileName()
	}
	{
		p.SetState(127)
		p.Match(ImgParserCLOSED_PARAN)
	}
	{
		p.SetState(128)
		p.Match(ImgParserSEMICOLON)
	}

	return localctx
}
