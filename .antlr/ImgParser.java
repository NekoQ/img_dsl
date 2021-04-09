// Generated from c:\Go\src\img\Img.g4 by ANTLR 4.8
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class ImgParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, ASSIGN=9, 
		DOT=10, COMMA=11, APOS=12, SEMICOLON=13, OPEN_PARAN=14, CLOSED_PARAN=15, 
		FOLDER=16, OPEN=17, NUMBER=18, ID=19, WHITESPACE=20;
	public static final int
		RULE_start = 0, RULE_declaration = 1, RULE_openFile = 2, RULE_openFolder = 3, 
		RULE_fileName = 4, RULE_imageType = 5, RULE_folderName = 6, RULE_action_ = 7, 
		RULE_actionType = 8, RULE_crop = 9, RULE_rotate = 10, RULE_flipX = 11, 
		RULE_flipY = 12, RULE_resize = 13, RULE_export = 14;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "declaration", "openFile", "openFolder", "fileName", "imageType", 
			"folderName", "action_", "actionType", "crop", "rotate", "flipX", "flipY", 
			"resize", "export"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'png'", "'jpg'", "'crop'", "'rotate'", "'flipX'", "'flipY'", "'resize'", 
			"'save'", "'='", "'.'", "','", "'\"'", "';'", "'('", "')'", "'[]'", "'open'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, "ASSIGN", "DOT", 
			"COMMA", "APOS", "SEMICOLON", "OPEN_PARAN", "CLOSED_PARAN", "FOLDER", 
			"OPEN", "NUMBER", "ID", "WHITESPACE"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Img.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public ImgParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class StartContext extends ParserRuleContext {
		public ExportContext export() {
			return getRuleContext(ExportContext.class,0);
		}
		public TerminalNode EOF() { return getToken(ImgParser.EOF, 0); }
		public List<DeclarationContext> declaration() {
			return getRuleContexts(DeclarationContext.class);
		}
		public DeclarationContext declaration(int i) {
			return getRuleContext(DeclarationContext.class,i);
		}
		public List<Action_Context> action_() {
			return getRuleContexts(Action_Context.class);
		}
		public Action_Context action_(int i) {
			return getRuleContext(Action_Context.class,i);
		}
		public StartContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_start; }
	}

	public final StartContext start() throws RecognitionException {
		StartContext _localctx = new StartContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_start);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(31); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(30);
					declaration();
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(33); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
			setState(38);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,1,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(35);
					action_();
					}
					} 
				}
				setState(40);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,1,_ctx);
			}
			setState(41);
			export();
			setState(42);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DeclarationContext extends ParserRuleContext {
		public OpenFileContext openFile() {
			return getRuleContext(OpenFileContext.class,0);
		}
		public OpenFolderContext openFolder() {
			return getRuleContext(OpenFolderContext.class,0);
		}
		public DeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaration; }
	}

	public final DeclarationContext declaration() throws RecognitionException {
		DeclarationContext _localctx = new DeclarationContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_declaration);
		try {
			setState(46);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(44);
				openFile();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(45);
				openFolder();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class OpenFileContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(ImgParser.ID, 0); }
		public TerminalNode ASSIGN() { return getToken(ImgParser.ASSIGN, 0); }
		public TerminalNode OPEN() { return getToken(ImgParser.OPEN, 0); }
		public TerminalNode OPEN_PARAN() { return getToken(ImgParser.OPEN_PARAN, 0); }
		public FileNameContext fileName() {
			return getRuleContext(FileNameContext.class,0);
		}
		public TerminalNode CLOSED_PARAN() { return getToken(ImgParser.CLOSED_PARAN, 0); }
		public TerminalNode SEMICOLON() { return getToken(ImgParser.SEMICOLON, 0); }
		public OpenFileContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_openFile; }
	}

	public final OpenFileContext openFile() throws RecognitionException {
		OpenFileContext _localctx = new OpenFileContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_openFile);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(48);
			match(ID);
			setState(49);
			match(ASSIGN);
			setState(50);
			match(OPEN);
			setState(51);
			match(OPEN_PARAN);
			setState(52);
			fileName();
			setState(53);
			match(CLOSED_PARAN);
			setState(54);
			match(SEMICOLON);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class OpenFolderContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(ImgParser.ID, 0); }
		public TerminalNode FOLDER() { return getToken(ImgParser.FOLDER, 0); }
		public TerminalNode ASSIGN() { return getToken(ImgParser.ASSIGN, 0); }
		public TerminalNode OPEN() { return getToken(ImgParser.OPEN, 0); }
		public TerminalNode OPEN_PARAN() { return getToken(ImgParser.OPEN_PARAN, 0); }
		public FolderNameContext folderName() {
			return getRuleContext(FolderNameContext.class,0);
		}
		public TerminalNode CLOSED_PARAN() { return getToken(ImgParser.CLOSED_PARAN, 0); }
		public TerminalNode SEMICOLON() { return getToken(ImgParser.SEMICOLON, 0); }
		public OpenFolderContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_openFolder; }
	}

	public final OpenFolderContext openFolder() throws RecognitionException {
		OpenFolderContext _localctx = new OpenFolderContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_openFolder);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(56);
			match(ID);
			setState(57);
			match(FOLDER);
			setState(58);
			match(ASSIGN);
			setState(59);
			match(OPEN);
			setState(60);
			match(OPEN_PARAN);
			setState(61);
			folderName();
			setState(62);
			match(CLOSED_PARAN);
			setState(63);
			match(SEMICOLON);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FileNameContext extends ParserRuleContext {
		public List<TerminalNode> APOS() { return getTokens(ImgParser.APOS); }
		public TerminalNode APOS(int i) {
			return getToken(ImgParser.APOS, i);
		}
		public TerminalNode ID() { return getToken(ImgParser.ID, 0); }
		public TerminalNode DOT() { return getToken(ImgParser.DOT, 0); }
		public ImageTypeContext imageType() {
			return getRuleContext(ImageTypeContext.class,0);
		}
		public FileNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fileName; }
	}

	public final FileNameContext fileName() throws RecognitionException {
		FileNameContext _localctx = new FileNameContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_fileName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(65);
			match(APOS);
			setState(66);
			match(ID);
			setState(67);
			match(DOT);
			setState(68);
			imageType();
			setState(69);
			match(APOS);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ImageTypeContext extends ParserRuleContext {
		public ImageTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_imageType; }
	}

	public final ImageTypeContext imageType() throws RecognitionException {
		ImageTypeContext _localctx = new ImageTypeContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_imageType);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(71);
			_la = _input.LA(1);
			if ( !(_la==T__0 || _la==T__1) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FolderNameContext extends ParserRuleContext {
		public List<TerminalNode> APOS() { return getTokens(ImgParser.APOS); }
		public TerminalNode APOS(int i) {
			return getToken(ImgParser.APOS, i);
		}
		public TerminalNode ID() { return getToken(ImgParser.ID, 0); }
		public FolderNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_folderName; }
	}

	public final FolderNameContext folderName() throws RecognitionException {
		FolderNameContext _localctx = new FolderNameContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_folderName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(73);
			match(APOS);
			setState(74);
			match(ID);
			setState(75);
			match(APOS);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Action_Context extends ParserRuleContext {
		public TerminalNode ID() { return getToken(ImgParser.ID, 0); }
		public TerminalNode DOT() { return getToken(ImgParser.DOT, 0); }
		public ActionTypeContext actionType() {
			return getRuleContext(ActionTypeContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(ImgParser.SEMICOLON, 0); }
		public Action_Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_action_; }
	}

	public final Action_Context action_() throws RecognitionException {
		Action_Context _localctx = new Action_Context(_ctx, getState());
		enterRule(_localctx, 14, RULE_action_);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(77);
			match(ID);
			setState(78);
			match(DOT);
			setState(79);
			actionType();
			setState(80);
			match(SEMICOLON);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ActionTypeContext extends ParserRuleContext {
		public CropContext crop() {
			return getRuleContext(CropContext.class,0);
		}
		public RotateContext rotate() {
			return getRuleContext(RotateContext.class,0);
		}
		public FlipXContext flipX() {
			return getRuleContext(FlipXContext.class,0);
		}
		public FlipYContext flipY() {
			return getRuleContext(FlipYContext.class,0);
		}
		public ResizeContext resize() {
			return getRuleContext(ResizeContext.class,0);
		}
		public ActionTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_actionType; }
	}

	public final ActionTypeContext actionType() throws RecognitionException {
		ActionTypeContext _localctx = new ActionTypeContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_actionType);
		try {
			setState(87);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__2:
				enterOuterAlt(_localctx, 1);
				{
				setState(82);
				crop();
				}
				break;
			case T__3:
				enterOuterAlt(_localctx, 2);
				{
				setState(83);
				rotate();
				}
				break;
			case T__4:
				enterOuterAlt(_localctx, 3);
				{
				setState(84);
				flipX();
				}
				break;
			case T__5:
				enterOuterAlt(_localctx, 4);
				{
				setState(85);
				flipY();
				}
				break;
			case T__6:
				enterOuterAlt(_localctx, 5);
				{
				setState(86);
				resize();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CropContext extends ParserRuleContext {
		public TerminalNode OPEN_PARAN() { return getToken(ImgParser.OPEN_PARAN, 0); }
		public List<TerminalNode> NUMBER() { return getTokens(ImgParser.NUMBER); }
		public TerminalNode NUMBER(int i) {
			return getToken(ImgParser.NUMBER, i);
		}
		public List<TerminalNode> COMMA() { return getTokens(ImgParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(ImgParser.COMMA, i);
		}
		public TerminalNode CLOSED_PARAN() { return getToken(ImgParser.CLOSED_PARAN, 0); }
		public CropContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_crop; }
	}

	public final CropContext crop() throws RecognitionException {
		CropContext _localctx = new CropContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_crop);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(89);
			match(T__2);
			setState(90);
			match(OPEN_PARAN);
			setState(91);
			match(NUMBER);
			setState(92);
			match(COMMA);
			setState(93);
			match(NUMBER);
			setState(94);
			match(COMMA);
			setState(95);
			match(NUMBER);
			setState(96);
			match(COMMA);
			setState(97);
			match(NUMBER);
			setState(98);
			match(CLOSED_PARAN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class RotateContext extends ParserRuleContext {
		public TerminalNode OPEN_PARAN() { return getToken(ImgParser.OPEN_PARAN, 0); }
		public TerminalNode NUMBER() { return getToken(ImgParser.NUMBER, 0); }
		public TerminalNode CLOSED_PARAN() { return getToken(ImgParser.CLOSED_PARAN, 0); }
		public RotateContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_rotate; }
	}

	public final RotateContext rotate() throws RecognitionException {
		RotateContext _localctx = new RotateContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_rotate);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(100);
			match(T__3);
			setState(101);
			match(OPEN_PARAN);
			setState(102);
			match(NUMBER);
			setState(103);
			match(CLOSED_PARAN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FlipXContext extends ParserRuleContext {
		public TerminalNode OPEN_PARAN() { return getToken(ImgParser.OPEN_PARAN, 0); }
		public TerminalNode CLOSED_PARAN() { return getToken(ImgParser.CLOSED_PARAN, 0); }
		public FlipXContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_flipX; }
	}

	public final FlipXContext flipX() throws RecognitionException {
		FlipXContext _localctx = new FlipXContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_flipX);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(105);
			match(T__4);
			setState(106);
			match(OPEN_PARAN);
			setState(107);
			match(CLOSED_PARAN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FlipYContext extends ParserRuleContext {
		public TerminalNode OPEN_PARAN() { return getToken(ImgParser.OPEN_PARAN, 0); }
		public TerminalNode CLOSED_PARAN() { return getToken(ImgParser.CLOSED_PARAN, 0); }
		public FlipYContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_flipY; }
	}

	public final FlipYContext flipY() throws RecognitionException {
		FlipYContext _localctx = new FlipYContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_flipY);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(109);
			match(T__5);
			setState(110);
			match(OPEN_PARAN);
			setState(111);
			match(CLOSED_PARAN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ResizeContext extends ParserRuleContext {
		public TerminalNode OPEN_PARAN() { return getToken(ImgParser.OPEN_PARAN, 0); }
		public List<TerminalNode> NUMBER() { return getTokens(ImgParser.NUMBER); }
		public TerminalNode NUMBER(int i) {
			return getToken(ImgParser.NUMBER, i);
		}
		public TerminalNode CLOSED_PARAN() { return getToken(ImgParser.CLOSED_PARAN, 0); }
		public TerminalNode COMMA() { return getToken(ImgParser.COMMA, 0); }
		public ResizeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_resize; }
	}

	public final ResizeContext resize() throws RecognitionException {
		ResizeContext _localctx = new ResizeContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_resize);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(113);
			match(T__6);
			setState(114);
			match(OPEN_PARAN);
			setState(115);
			match(NUMBER);
			setState(118);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMMA) {
				{
				setState(116);
				match(COMMA);
				setState(117);
				match(NUMBER);
				}
			}

			setState(120);
			match(CLOSED_PARAN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExportContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(ImgParser.ID, 0); }
		public TerminalNode DOT() { return getToken(ImgParser.DOT, 0); }
		public TerminalNode OPEN_PARAN() { return getToken(ImgParser.OPEN_PARAN, 0); }
		public FileNameContext fileName() {
			return getRuleContext(FileNameContext.class,0);
		}
		public TerminalNode CLOSED_PARAN() { return getToken(ImgParser.CLOSED_PARAN, 0); }
		public TerminalNode SEMICOLON() { return getToken(ImgParser.SEMICOLON, 0); }
		public ExportContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_export; }
	}

	public final ExportContext export() throws RecognitionException {
		ExportContext _localctx = new ExportContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_export);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(122);
			match(ID);
			setState(123);
			match(DOT);
			setState(124);
			match(T__7);
			setState(125);
			match(OPEN_PARAN);
			setState(126);
			fileName();
			setState(127);
			match(CLOSED_PARAN);
			setState(128);
			match(SEMICOLON);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\26\u0085\4\2\t\2"+
		"\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\3\2\6\2\"\n\2\r\2"+
		"\16\2#\3\2\7\2\'\n\2\f\2\16\2*\13\2\3\2\3\2\3\2\3\3\3\3\5\3\61\n\3\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\6\3"+
		"\6\3\6\3\6\3\6\3\6\3\7\3\7\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\n\3\n"+
		"\3\n\3\n\3\n\5\nZ\n\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13"+
		"\3\13\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\17\3\17"+
		"\3\17\3\17\3\17\5\17y\n\17\3\17\3\17\3\20\3\20\3\20\3\20\3\20\3\20\3\20"+
		"\3\20\3\20\2\2\21\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36\2\3\3\2\3\4\2"+
		"}\2!\3\2\2\2\4\60\3\2\2\2\6\62\3\2\2\2\b:\3\2\2\2\nC\3\2\2\2\fI\3\2\2"+
		"\2\16K\3\2\2\2\20O\3\2\2\2\22Y\3\2\2\2\24[\3\2\2\2\26f\3\2\2\2\30k\3\2"+
		"\2\2\32o\3\2\2\2\34s\3\2\2\2\36|\3\2\2\2 \"\5\4\3\2! \3\2\2\2\"#\3\2\2"+
		"\2#!\3\2\2\2#$\3\2\2\2$(\3\2\2\2%\'\5\20\t\2&%\3\2\2\2\'*\3\2\2\2(&\3"+
		"\2\2\2()\3\2\2\2)+\3\2\2\2*(\3\2\2\2+,\5\36\20\2,-\7\2\2\3-\3\3\2\2\2"+
		".\61\5\6\4\2/\61\5\b\5\2\60.\3\2\2\2\60/\3\2\2\2\61\5\3\2\2\2\62\63\7"+
		"\25\2\2\63\64\7\13\2\2\64\65\7\23\2\2\65\66\7\20\2\2\66\67\5\n\6\2\67"+
		"8\7\21\2\289\7\17\2\29\7\3\2\2\2:;\7\25\2\2;<\7\22\2\2<=\7\13\2\2=>\7"+
		"\23\2\2>?\7\20\2\2?@\5\16\b\2@A\7\21\2\2AB\7\17\2\2B\t\3\2\2\2CD\7\16"+
		"\2\2DE\7\25\2\2EF\7\f\2\2FG\5\f\7\2GH\7\16\2\2H\13\3\2\2\2IJ\t\2\2\2J"+
		"\r\3\2\2\2KL\7\16\2\2LM\7\25\2\2MN\7\16\2\2N\17\3\2\2\2OP\7\25\2\2PQ\7"+
		"\f\2\2QR\5\22\n\2RS\7\17\2\2S\21\3\2\2\2TZ\5\24\13\2UZ\5\26\f\2VZ\5\30"+
		"\r\2WZ\5\32\16\2XZ\5\34\17\2YT\3\2\2\2YU\3\2\2\2YV\3\2\2\2YW\3\2\2\2Y"+
		"X\3\2\2\2Z\23\3\2\2\2[\\\7\5\2\2\\]\7\20\2\2]^\7\24\2\2^_\7\r\2\2_`\7"+
		"\24\2\2`a\7\r\2\2ab\7\24\2\2bc\7\r\2\2cd\7\24\2\2de\7\21\2\2e\25\3\2\2"+
		"\2fg\7\6\2\2gh\7\20\2\2hi\7\24\2\2ij\7\21\2\2j\27\3\2\2\2kl\7\7\2\2lm"+
		"\7\20\2\2mn\7\21\2\2n\31\3\2\2\2op\7\b\2\2pq\7\20\2\2qr\7\21\2\2r\33\3"+
		"\2\2\2st\7\t\2\2tu\7\20\2\2ux\7\24\2\2vw\7\r\2\2wy\7\24\2\2xv\3\2\2\2"+
		"xy\3\2\2\2yz\3\2\2\2z{\7\21\2\2{\35\3\2\2\2|}\7\25\2\2}~\7\f\2\2~\177"+
		"\7\n\2\2\177\u0080\7\20\2\2\u0080\u0081\5\n\6\2\u0081\u0082\7\21\2\2\u0082"+
		"\u0083\7\17\2\2\u0083\37\3\2\2\2\7#(\60Yx";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}