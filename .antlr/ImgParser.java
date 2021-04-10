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
			setState(49);
			export();
			setState(50);
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
			setState(54);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(52);
				openFile();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(53);
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
			setState(56);
			match(ID);
			setState(57);
			match(ASSIGN);
			setState(58);
			match(OPEN);
			setState(59);
			match(OPEN_PARAN);
			setState(60);
			fileName();
			setState(61);
			match(CLOSED_PARAN);
			setState(62);
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
			setState(64);
			match(ID);
			setState(65);
			match(FOLDER);
			setState(66);
			match(ASSIGN);
			setState(67);
			match(OPEN);
			setState(68);
			match(OPEN_PARAN);
			setState(69);
			folderName();
			setState(70);
			match(CLOSED_PARAN);
			setState(71);
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
			setState(73);
			match(APOS);
			setState(74);
			match(ID);
			setState(75);
			match(DOT);
			setState(76);
			imageType();
			setState(77);
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
			setState(79);
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
			setState(81);
			match(APOS);
			setState(82);
			match(ID);
			setState(83);
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
			setState(85);
			match(ID);
			setState(86);
			match(DOT);
			setState(87);
			actionType();
			setState(88);
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
			setState(99);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__2:
				enterOuterAlt(_localctx, 1);
				{
				setState(90);
				crop();
				}
				break;
			case T__3:
				enterOuterAlt(_localctx, 2);
				{
				setState(91);
				rotate();
				}
				break;
			case T__4:
				enterOuterAlt(_localctx, 3);
				{
				setState(92);
				flipX();
				}
				break;
			case T__5:
				enterOuterAlt(_localctx, 4);
				{
				setState(93);
				flipY();
				}
				break;
			case T__6:
				enterOuterAlt(_localctx, 5);
				{
				setState(94);
				resize();
				}
				break;
			case T__7:
				enterOuterAlt(_localctx, 6);
				{
				setState(95);
				brightness();
				}
				break;
			case T__8:
				enterOuterAlt(_localctx, 7);
				{
				setState(96);
				contrast();
				}
				break;
			case T__9:
				enterOuterAlt(_localctx, 8);
				{
				setState(97);
				saturation();
				}
				break;
			case T__10:
				enterOuterAlt(_localctx, 9);
				{
				setState(98);
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
			setState(101);
			match(T__2);
			setState(102);
			match(OPEN_PARAN);
			setState(103);
			match(NUMBER);
			setState(104);
			match(COMMA);
			setState(105);
			match(NUMBER);
			setState(106);
			match(COMMA);
			setState(107);
			match(NUMBER);
			setState(108);
			match(COMMA);
			setState(109);
			match(NUMBER);
			setState(110);
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
			setState(112);
			match(T__3);
			setState(113);
			match(OPEN_PARAN);
			setState(114);
			match(NUMBER);
			setState(115);
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
			setState(117);
			match(T__4);
			setState(118);
			match(OPEN_PARAN);
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
			setState(121);
			match(T__5);
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
			setState(125);
			match(T__6);
			setState(126);
			match(OPEN_PARAN);
			setState(127);
			match(NUMBER);
			setState(130);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMMA) {
				{
				setState(128);
				match(COMMA);
				setState(129);
				match(NUMBER);
				}
			}

			setState(132);
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
			setState(134);
			match(T__7);
			setState(135);
			match(OPEN_PARAN);
			setState(136);
			match(NUMBER);
			setState(137);
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
			setState(139);
			match(T__8);
			setState(140);
			match(OPEN_PARAN);
			setState(141);
			match(NUMBER);
			setState(142);
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
			setState(144);
			match(T__9);
			setState(145);
			match(OPEN_PARAN);
			setState(146);
			match(NUMBER);
			setState(147);
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
			setState(149);
			match(T__10);
			setState(150);
			match(OPEN_PARAN);
			setState(151);
			match(NUMBER);
			setState(152);
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
			setState(154);
			match(ID);
			setState(155);
			match(DOT);
			setState(156);
			match(T__11);
			setState(157);
			match(OPEN_PARAN);
			setState(158);
			fileName();
			setState(159);
			match(CLOSED_PARAN);
			setState(160);
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\32\u00a5\4\2\t\2"+
		"\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\3\2\6\2*\n\2\r\2\16\2+\3\2\7\2/\n\2\f\2\16\2\62\13"+
		"\2\3\2\3\2\3\2\3\3\3\3\5\39\n\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\5\3"+
		"\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\b\3\b"+
		"\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\5\nf"+
		"\n\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3"+
		"\f\3\f\3\f\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3\17"+
		"\5\17\u0085\n\17\3\17\3\17\3\20\3\20\3\20\3\20\3\20\3\21\3\21\3\21\3\21"+
		"\3\21\3\22\3\22\3\22\3\22\3\22\3\23\3\23\3\23\3\23\3\23\3\24\3\24\3\24"+
		"\3\24\3\24\3\24\3\24\3\24\3\24\2\2\25\2\4\6\b\n\f\16\20\22\24\26\30\32"+
		"\34\36 \"$&\2\3\3\2\3\4\2\u009d\2)\3\2\2\2\48\3\2\2\2\6:\3\2\2\2\bB\3"+
		"\2\2\2\nK\3\2\2\2\fQ\3\2\2\2\16S\3\2\2\2\20W\3\2\2\2\22e\3\2\2\2\24g\3"+
		"\2\2\2\26r\3\2\2\2\30w\3\2\2\2\32{\3\2\2\2\34\177\3\2\2\2\36\u0088\3\2"+
		"\2\2 \u008d\3\2\2\2\"\u0092\3\2\2\2$\u0097\3\2\2\2&\u009c\3\2\2\2(*\5"+
		"\4\3\2)(\3\2\2\2*+\3\2\2\2+)\3\2\2\2+,\3\2\2\2,\60\3\2\2\2-/\5\20\t\2"+
		".-\3\2\2\2/\62\3\2\2\2\60.\3\2\2\2\60\61\3\2\2\2\61\63\3\2\2\2\62\60\3"+
		"\2\2\2\63\64\5&\24\2\64\65\7\2\2\3\65\3\3\2\2\2\669\5\6\4\2\679\5\b\5"+
		"\28\66\3\2\2\28\67\3\2\2\29\5\3\2\2\2:;\7\31\2\2;<\7\17\2\2<=\7\27\2\2"+
		"=>\7\24\2\2>?\5\n\6\2?@\7\25\2\2@A\7\23\2\2A\7\3\2\2\2BC\7\31\2\2CD\7"+
		"\26\2\2DE\7\17\2\2EF\7\27\2\2FG\7\24\2\2GH\5\16\b\2HI\7\25\2\2IJ\7\23"+
		"\2\2J\t\3\2\2\2KL\7\22\2\2LM\7\31\2\2MN\7\20\2\2NO\5\f\7\2OP\7\22\2\2"+
		"P\13\3\2\2\2QR\t\2\2\2R\r\3\2\2\2ST\7\22\2\2TU\7\31\2\2UV\7\22\2\2V\17"+
		"\3\2\2\2WX\7\31\2\2XY\7\20\2\2YZ\5\22\n\2Z[\7\23\2\2[\21\3\2\2\2\\f\5"+
		"\24\13\2]f\5\26\f\2^f\5\30\r\2_f\5\32\16\2`f\5\34\17\2af\5\36\20\2bf\5"+
		" \21\2cf\5\"\22\2df\5$\23\2e\\\3\2\2\2e]\3\2\2\2e^\3\2\2\2e_\3\2\2\2e"+
		"`\3\2\2\2ea\3\2\2\2eb\3\2\2\2ec\3\2\2\2ed\3\2\2\2f\23\3\2\2\2gh\7\5\2"+
		"\2hi\7\24\2\2ij\7\30\2\2jk\7\21\2\2kl\7\30\2\2lm\7\21\2\2mn\7\30\2\2n"+
		"o\7\21\2\2op\7\30\2\2pq\7\25\2\2q\25\3\2\2\2rs\7\6\2\2st\7\24\2\2tu\7"+
		"\30\2\2uv\7\25\2\2v\27\3\2\2\2wx\7\7\2\2xy\7\24\2\2yz\7\25\2\2z\31\3\2"+
		"\2\2{|\7\b\2\2|}\7\24\2\2}~\7\25\2\2~\33\3\2\2\2\177\u0080\7\t\2\2\u0080"+
		"\u0081\7\24\2\2\u0081\u0084\7\30\2\2\u0082\u0083\7\21\2\2\u0083\u0085"+
		"\7\30\2\2\u0084\u0082\3\2\2\2\u0084\u0085\3\2\2\2\u0085\u0086\3\2\2\2"+
		"\u0086\u0087\7\25\2\2\u0087\35\3\2\2\2\u0088\u0089\7\n\2\2\u0089\u008a"+
		"\7\24\2\2\u008a\u008b\7\30\2\2\u008b\u008c\7\25\2\2\u008c\37\3\2\2\2\u008d"+
		"\u008e\7\13\2\2\u008e\u008f\7\24\2\2\u008f\u0090\7\30\2\2\u0090\u0091"+
		"\7\25\2\2\u0091!\3\2\2\2\u0092\u0093\7\f\2\2\u0093\u0094\7\24\2\2\u0094"+
		"\u0095\7\30\2\2\u0095\u0096\7\25\2\2\u0096#\3\2\2\2\u0097\u0098\7\r\2"+
		"\2\u0098\u0099\7\24\2\2\u0099\u009a\7\30\2\2\u009a\u009b\7\25\2\2\u009b"+
		"%\3\2\2\2\u009c\u009d\7\31\2\2\u009d\u009e\7\20\2\2\u009e\u009f\7\16\2"+
		"\2\u009f\u00a0\7\24\2\2\u00a0\u00a1\5\n\6\2\u00a1\u00a2\7\25\2\2\u00a2"+
		"\u00a3\7\23\2\2\u00a3\'\3\2\2\2\7+\608e\u0084";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}