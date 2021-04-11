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
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, ASSIGN=13, DOT=14, COMMA=15, APOS=16, SEMICOLON=17, 
		OPEN_PARAN=18, CLOSED_PARAN=19, FOLDER=20, OPEN=21, NUMBER=22, ID=23, 
		WHITESPACE=24;
	public static final int
		RULE_start = 0, RULE_declaration = 1, RULE_openFile = 2, RULE_openFolder = 3, 
		RULE_fileName = 4, RULE_imageType = 5, RULE_folderName = 6, RULE_action_ = 7, 
		RULE_actionType = 8, RULE_crop = 9, RULE_rotate = 10, RULE_flipX = 11, 
		RULE_flipY = 12, RULE_resize = 13, RULE_brightness = 14, RULE_contrast = 15, 
		RULE_saturation = 16, RULE_pixelate = 17, RULE_export = 18;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "declaration", "openFile", "openFolder", "fileName", "imageType", 
			"folderName", "action_", "actionType", "crop", "rotate", "flipX", "flipY", 
			"resize", "brightness", "contrast", "saturation", "pixelate", "export"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'png'", "'jpg'", "'crop'", "'rotate'", "'flipX'", "'flipY'", "'resize'", 
			"'brightness'", "'contrast'", "'saturation'", "'pixelate'", "'save'", 
			"'='", "'.'", "','", "'\"'", "';'", "'('", "')'", "'[]'", "'open'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, "ASSIGN", "DOT", "COMMA", "APOS", "SEMICOLON", "OPEN_PARAN", "CLOSED_PARAN", 
			"FOLDER", "OPEN", "NUMBER", "ID", "WHITESPACE"
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
		public List<ExportContext> export() {
			return getRuleContexts(ExportContext.class);
		}
		public ExportContext export(int i) {
			return getRuleContext(ExportContext.class,i);
		}
		public StartContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_start; }
	}

	public final StartContext start() throws RecognitionException {
		StartContext _localctx = new StartContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_start);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(39); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(38);
					declaration();
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(41); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
			setState(46);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,1,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(43);
					action_();
					}
					} 
				}
				setState(48);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,1,_ctx);
			}
			setState(50); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(49);
				export();
				}
				}
				setState(52); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==ID );
			setState(54);
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
			setState(58);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(56);
				openFile();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(57);
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
			setState(60);
			match(ID);
			setState(61);
			match(ASSIGN);
			setState(62);
			match(OPEN);
			setState(63);
			match(OPEN_PARAN);
			setState(64);
			fileName();
			setState(65);
			match(CLOSED_PARAN);
			setState(66);
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
			setState(68);
			match(ID);
			setState(69);
			match(FOLDER);
			setState(70);
			match(ASSIGN);
			setState(71);
			match(OPEN);
			setState(72);
			match(OPEN_PARAN);
			setState(73);
			folderName();
			setState(74);
			match(CLOSED_PARAN);
			setState(75);
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
			setState(77);
			match(APOS);
			setState(78);
			match(ID);
			setState(79);
			match(DOT);
			setState(80);
			imageType();
			setState(81);
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
			setState(83);
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
			setState(85);
			match(APOS);
			setState(86);
			match(ID);
			setState(87);
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
			setState(89);
			match(ID);
			setState(90);
			match(DOT);
			setState(91);
			actionType();
			setState(92);
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
		public BrightnessContext brightness() {
			return getRuleContext(BrightnessContext.class,0);
		}
		public ContrastContext contrast() {
			return getRuleContext(ContrastContext.class,0);
		}
		public SaturationContext saturation() {
			return getRuleContext(SaturationContext.class,0);
		}
		public PixelateContext pixelate() {
			return getRuleContext(PixelateContext.class,0);
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
			setState(103);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__2:
				enterOuterAlt(_localctx, 1);
				{
				setState(94);
				crop();
				}
				break;
			case T__3:
				enterOuterAlt(_localctx, 2);
				{
				setState(95);
				rotate();
				}
				break;
			case T__4:
				enterOuterAlt(_localctx, 3);
				{
				setState(96);
				flipX();
				}
				break;
			case T__5:
				enterOuterAlt(_localctx, 4);
				{
				setState(97);
				flipY();
				}
				break;
			case T__6:
				enterOuterAlt(_localctx, 5);
				{
				setState(98);
				resize();
				}
				break;
			case T__7:
				enterOuterAlt(_localctx, 6);
				{
				setState(99);
				brightness();
				}
				break;
			case T__8:
				enterOuterAlt(_localctx, 7);
				{
				setState(100);
				contrast();
				}
				break;
			case T__9:
				enterOuterAlt(_localctx, 8);
				{
				setState(101);
				saturation();
				}
				break;
			case T__10:
				enterOuterAlt(_localctx, 9);
				{
				setState(102);
				pixelate();
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
			setState(105);
			match(T__2);
			setState(106);
			match(OPEN_PARAN);
			setState(107);
			match(NUMBER);
			setState(108);
			match(COMMA);
			setState(109);
			match(NUMBER);
			setState(110);
			match(COMMA);
			setState(111);
			match(NUMBER);
			setState(112);
			match(COMMA);
			setState(113);
			match(NUMBER);
			setState(114);
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
			setState(116);
			match(T__3);
			setState(117);
			match(OPEN_PARAN);
			setState(118);
			match(NUMBER);
			setState(119);
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
			setState(121);
			match(T__4);
			setState(122);
			match(OPEN_PARAN);
			setState(123);
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
			setState(125);
			match(T__5);
			setState(126);
			match(OPEN_PARAN);
			setState(127);
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
			setState(129);
			match(T__6);
			setState(130);
			match(OPEN_PARAN);
			setState(131);
			match(NUMBER);
			setState(134);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMMA) {
				{
				setState(132);
				match(COMMA);
				setState(133);
				match(NUMBER);
				}
			}

			setState(136);
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

	public static class BrightnessContext extends ParserRuleContext {
		public TerminalNode OPEN_PARAN() { return getToken(ImgParser.OPEN_PARAN, 0); }
		public TerminalNode NUMBER() { return getToken(ImgParser.NUMBER, 0); }
		public TerminalNode CLOSED_PARAN() { return getToken(ImgParser.CLOSED_PARAN, 0); }
		public BrightnessContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_brightness; }
	}

	public final BrightnessContext brightness() throws RecognitionException {
		BrightnessContext _localctx = new BrightnessContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_brightness);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(138);
			match(T__7);
			setState(139);
			match(OPEN_PARAN);
			setState(140);
			match(NUMBER);
			setState(141);
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

	public static class ContrastContext extends ParserRuleContext {
		public TerminalNode OPEN_PARAN() { return getToken(ImgParser.OPEN_PARAN, 0); }
		public TerminalNode NUMBER() { return getToken(ImgParser.NUMBER, 0); }
		public TerminalNode CLOSED_PARAN() { return getToken(ImgParser.CLOSED_PARAN, 0); }
		public ContrastContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_contrast; }
	}

	public final ContrastContext contrast() throws RecognitionException {
		ContrastContext _localctx = new ContrastContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_contrast);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(143);
			match(T__8);
			setState(144);
			match(OPEN_PARAN);
			setState(145);
			match(NUMBER);
			setState(146);
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

	public static class SaturationContext extends ParserRuleContext {
		public TerminalNode OPEN_PARAN() { return getToken(ImgParser.OPEN_PARAN, 0); }
		public TerminalNode NUMBER() { return getToken(ImgParser.NUMBER, 0); }
		public TerminalNode CLOSED_PARAN() { return getToken(ImgParser.CLOSED_PARAN, 0); }
		public SaturationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_saturation; }
	}

	public final SaturationContext saturation() throws RecognitionException {
		SaturationContext _localctx = new SaturationContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_saturation);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(148);
			match(T__9);
			setState(149);
			match(OPEN_PARAN);
			setState(150);
			match(NUMBER);
			setState(151);
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

	public static class PixelateContext extends ParserRuleContext {
		public TerminalNode OPEN_PARAN() { return getToken(ImgParser.OPEN_PARAN, 0); }
		public TerminalNode NUMBER() { return getToken(ImgParser.NUMBER, 0); }
		public TerminalNode CLOSED_PARAN() { return getToken(ImgParser.CLOSED_PARAN, 0); }
		public PixelateContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pixelate; }
	}

	public final PixelateContext pixelate() throws RecognitionException {
		PixelateContext _localctx = new PixelateContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_pixelate);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(153);
			match(T__10);
			setState(154);
			match(OPEN_PARAN);
			setState(155);
			match(NUMBER);
			setState(156);
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
		enterRule(_localctx, 36, RULE_export);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(158);
			match(ID);
			setState(159);
			match(DOT);
			setState(160);
			match(T__11);
			setState(161);
			match(OPEN_PARAN);
			setState(162);
			fileName();
			setState(163);
			match(CLOSED_PARAN);
			setState(164);
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\32\u00a9\4\2\t\2"+
		"\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\3\2\6\2*\n\2\r\2\16\2+\3\2\7\2/\n\2\f\2\16\2\62\13"+
		"\2\3\2\6\2\65\n\2\r\2\16\2\66\3\2\3\2\3\3\3\3\5\3=\n\3\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3"+
		"\6\3\6\3\7\3\7\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n"+
		"\3\n\3\n\3\n\3\n\5\nj\n\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13"+
		"\3\13\3\13\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\17"+
		"\3\17\3\17\3\17\3\17\5\17\u0089\n\17\3\17\3\17\3\20\3\20\3\20\3\20\3\20"+
		"\3\21\3\21\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3\22\3\23\3\23\3\23\3\23"+
		"\3\23\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\2\2\25\2\4\6\b\n\f"+
		"\16\20\22\24\26\30\32\34\36 \"$&\2\3\3\2\3\4\2\u00a2\2)\3\2\2\2\4<\3\2"+
		"\2\2\6>\3\2\2\2\bF\3\2\2\2\nO\3\2\2\2\fU\3\2\2\2\16W\3\2\2\2\20[\3\2\2"+
		"\2\22i\3\2\2\2\24k\3\2\2\2\26v\3\2\2\2\30{\3\2\2\2\32\177\3\2\2\2\34\u0083"+
		"\3\2\2\2\36\u008c\3\2\2\2 \u0091\3\2\2\2\"\u0096\3\2\2\2$\u009b\3\2\2"+
		"\2&\u00a0\3\2\2\2(*\5\4\3\2)(\3\2\2\2*+\3\2\2\2+)\3\2\2\2+,\3\2\2\2,\60"+
		"\3\2\2\2-/\5\20\t\2.-\3\2\2\2/\62\3\2\2\2\60.\3\2\2\2\60\61\3\2\2\2\61"+
		"\64\3\2\2\2\62\60\3\2\2\2\63\65\5&\24\2\64\63\3\2\2\2\65\66\3\2\2\2\66"+
		"\64\3\2\2\2\66\67\3\2\2\2\678\3\2\2\289\7\2\2\39\3\3\2\2\2:=\5\6\4\2;"+
		"=\5\b\5\2<:\3\2\2\2<;\3\2\2\2=\5\3\2\2\2>?\7\31\2\2?@\7\17\2\2@A\7\27"+
		"\2\2AB\7\24\2\2BC\5\n\6\2CD\7\25\2\2DE\7\23\2\2E\7\3\2\2\2FG\7\31\2\2"+
		"GH\7\26\2\2HI\7\17\2\2IJ\7\27\2\2JK\7\24\2\2KL\5\16\b\2LM\7\25\2\2MN\7"+
		"\23\2\2N\t\3\2\2\2OP\7\22\2\2PQ\7\31\2\2QR\7\20\2\2RS\5\f\7\2ST\7\22\2"+
		"\2T\13\3\2\2\2UV\t\2\2\2V\r\3\2\2\2WX\7\22\2\2XY\7\31\2\2YZ\7\22\2\2Z"+
		"\17\3\2\2\2[\\\7\31\2\2\\]\7\20\2\2]^\5\22\n\2^_\7\23\2\2_\21\3\2\2\2"+
		"`j\5\24\13\2aj\5\26\f\2bj\5\30\r\2cj\5\32\16\2dj\5\34\17\2ej\5\36\20\2"+
		"fj\5 \21\2gj\5\"\22\2hj\5$\23\2i`\3\2\2\2ia\3\2\2\2ib\3\2\2\2ic\3\2\2"+
		"\2id\3\2\2\2ie\3\2\2\2if\3\2\2\2ig\3\2\2\2ih\3\2\2\2j\23\3\2\2\2kl\7\5"+
		"\2\2lm\7\24\2\2mn\7\30\2\2no\7\21\2\2op\7\30\2\2pq\7\21\2\2qr\7\30\2\2"+
		"rs\7\21\2\2st\7\30\2\2tu\7\25\2\2u\25\3\2\2\2vw\7\6\2\2wx\7\24\2\2xy\7"+
		"\30\2\2yz\7\25\2\2z\27\3\2\2\2{|\7\7\2\2|}\7\24\2\2}~\7\25\2\2~\31\3\2"+
		"\2\2\177\u0080\7\b\2\2\u0080\u0081\7\24\2\2\u0081\u0082\7\25\2\2\u0082"+
		"\33\3\2\2\2\u0083\u0084\7\t\2\2\u0084\u0085\7\24\2\2\u0085\u0088\7\30"+
		"\2\2\u0086\u0087\7\21\2\2\u0087\u0089\7\30\2\2\u0088\u0086\3\2\2\2\u0088"+
		"\u0089\3\2\2\2\u0089\u008a\3\2\2\2\u008a\u008b\7\25\2\2\u008b\35\3\2\2"+
		"\2\u008c\u008d\7\n\2\2\u008d\u008e\7\24\2\2\u008e\u008f\7\30\2\2\u008f"+
		"\u0090\7\25\2\2\u0090\37\3\2\2\2\u0091\u0092\7\13\2\2\u0092\u0093\7\24"+
		"\2\2\u0093\u0094\7\30\2\2\u0094\u0095\7\25\2\2\u0095!\3\2\2\2\u0096\u0097"+
		"\7\f\2\2\u0097\u0098\7\24\2\2\u0098\u0099\7\30\2\2\u0099\u009a\7\25\2"+
		"\2\u009a#\3\2\2\2\u009b\u009c\7\r\2\2\u009c\u009d\7\24\2\2\u009d\u009e"+
		"\7\30\2\2\u009e\u009f\7\25\2\2\u009f%\3\2\2\2\u00a0\u00a1\7\31\2\2\u00a1"+
		"\u00a2\7\20\2\2\u00a2\u00a3\7\16\2\2\u00a3\u00a4\7\24\2\2\u00a4\u00a5"+
		"\5\n\6\2\u00a5\u00a6\7\25\2\2\u00a6\u00a7\7\23\2\2\u00a7\'\3\2\2\2\b+"+
		"\60\66<i\u0088";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}