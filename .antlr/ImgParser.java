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
		T__9=10, T__10=11, T__11=12, T__12=13, ASSIGN=14, DOT=15, COMMA=16, APOS=17, 
		SEMICOLON=18, OPEN_PARAN=19, CLOSED_PARAN=20, FOLDER=21, OPEN=22, NUMBER=23, 
		MINUS=24, ID=25, WHITESPACE=26;
	public static final int
		RULE_start = 0, RULE_declaration = 1, RULE_openFile = 2, RULE_openFolder = 3, 
		RULE_fileName = 4, RULE_imageType = 5, RULE_folderName = 6, RULE_action_ = 7, 
		RULE_actionType = 8, RULE_crop = 9, RULE_rotate = 10, RULE_flipX = 11, 
		RULE_flipY = 12, RULE_resize = 13, RULE_brightness = 14, RULE_contrast = 15, 
		RULE_saturation = 16, RULE_pixelate = 17, RULE_export = 18, RULE_text = 19, 
		RULE_comment = 20;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "declaration", "openFile", "openFolder", "fileName", "imageType", 
			"folderName", "action_", "actionType", "crop", "rotate", "flipX", "flipY", 
			"resize", "brightness", "contrast", "saturation", "pixelate", "export", 
			"text", "comment"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'png'", "'jpg'", "'crop'", "'rotate'", "'flipX'", "'flipY'", "'resize'", 
			"'brightness'", "'contrast'", "'saturation'", "'pixelate'", "'save'", 
			"'--'", "'='", "'.'", "','", "'\"'", "';'", "'('", "')'", "'[]'", "'open'", 
			null, "'-'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, "ASSIGN", "DOT", "COMMA", "APOS", "SEMICOLON", "OPEN_PARAN", 
			"CLOSED_PARAN", "FOLDER", "OPEN", "NUMBER", "MINUS", "ID", "WHITESPACE"
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
		public List<CommentContext> comment() {
			return getRuleContexts(CommentContext.class);
		}
		public CommentContext comment(int i) {
			return getRuleContext(CommentContext.class,i);
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
			setState(49); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(45);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==T__12) {
						{
						{
						setState(42);
						comment();
						}
						}
						setState(47);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(48);
					declaration();
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(51); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,1,_ctx);
			} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
			setState(56);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(53);
					comment();
					}
					} 
				}
				setState(58);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			}
			setState(62);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(59);
					action_();
					}
					} 
				}
				setState(64);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
			}
			setState(68);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__12) {
				{
				{
				setState(65);
				comment();
				}
				}
				setState(70);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(72); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(71);
				export();
				}
				}
				setState(74); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==ID );
			setState(76);
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
			setState(80);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(78);
				openFile();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(79);
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
			setState(82);
			match(ID);
			setState(83);
			match(ASSIGN);
			setState(84);
			match(OPEN);
			setState(85);
			match(OPEN_PARAN);
			setState(86);
			fileName();
			setState(87);
			match(CLOSED_PARAN);
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
			setState(90);
			match(ID);
			setState(91);
			match(FOLDER);
			setState(92);
			match(ASSIGN);
			setState(93);
			match(OPEN);
			setState(94);
			match(OPEN_PARAN);
			setState(95);
			folderName();
			setState(96);
			match(CLOSED_PARAN);
			setState(97);
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
			setState(99);
			match(APOS);
			setState(100);
			match(ID);
			setState(101);
			match(DOT);
			setState(102);
			imageType();
			setState(103);
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
			setState(105);
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
			setState(107);
			match(APOS);
			setState(108);
			match(ID);
			setState(109);
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
			setState(111);
			match(ID);
			setState(112);
			match(DOT);
			setState(113);
			actionType();
			setState(114);
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
			setState(125);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__2:
				enterOuterAlt(_localctx, 1);
				{
				setState(116);
				crop();
				}
				break;
			case T__3:
				enterOuterAlt(_localctx, 2);
				{
				setState(117);
				rotate();
				}
				break;
			case T__4:
				enterOuterAlt(_localctx, 3);
				{
				setState(118);
				flipX();
				}
				break;
			case T__5:
				enterOuterAlt(_localctx, 4);
				{
				setState(119);
				flipY();
				}
				break;
			case T__6:
				enterOuterAlt(_localctx, 5);
				{
				setState(120);
				resize();
				}
				break;
			case T__7:
				enterOuterAlt(_localctx, 6);
				{
				setState(121);
				brightness();
				}
				break;
			case T__8:
				enterOuterAlt(_localctx, 7);
				{
				setState(122);
				contrast();
				}
				break;
			case T__9:
				enterOuterAlt(_localctx, 8);
				{
				setState(123);
				saturation();
				}
				break;
			case T__10:
				enterOuterAlt(_localctx, 9);
				{
				setState(124);
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
			setState(127);
			match(T__2);
			setState(128);
			match(OPEN_PARAN);
			setState(129);
			match(NUMBER);
			setState(130);
			match(COMMA);
			setState(131);
			match(NUMBER);
			setState(132);
			match(COMMA);
			setState(133);
			match(NUMBER);
			setState(134);
			match(COMMA);
			setState(135);
			match(NUMBER);
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
			setState(138);
			match(T__3);
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
			setState(143);
			match(T__4);
			setState(144);
			match(OPEN_PARAN);
			setState(145);
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
			setState(147);
			match(T__5);
			setState(148);
			match(OPEN_PARAN);
			setState(149);
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
			setState(151);
			match(T__6);
			setState(152);
			match(OPEN_PARAN);
			setState(153);
			match(NUMBER);
			setState(156);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMMA) {
				{
				setState(154);
				match(COMMA);
				setState(155);
				match(NUMBER);
				}
			}

			setState(158);
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
			setState(160);
			match(T__7);
			setState(161);
			match(OPEN_PARAN);
			setState(162);
			match(NUMBER);
			setState(163);
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
			setState(165);
			match(T__8);
			setState(166);
			match(OPEN_PARAN);
			setState(167);
			match(NUMBER);
			setState(168);
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
			setState(170);
			match(T__9);
			setState(171);
			match(OPEN_PARAN);
			setState(172);
			match(NUMBER);
			setState(173);
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
			setState(175);
			match(T__10);
			setState(176);
			match(OPEN_PARAN);
			setState(177);
			match(NUMBER);
			setState(178);
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
		public TerminalNode CLOSED_PARAN() { return getToken(ImgParser.CLOSED_PARAN, 0); }
		public TerminalNode SEMICOLON() { return getToken(ImgParser.SEMICOLON, 0); }
		public FileNameContext fileName() {
			return getRuleContext(FileNameContext.class,0);
		}
		public FolderNameContext folderName() {
			return getRuleContext(FolderNameContext.class,0);
		}
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
			setState(180);
			match(ID);
			setState(181);
			match(DOT);
			setState(182);
			match(T__11);
			setState(183);
			match(OPEN_PARAN);
			setState(186);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				{
				setState(184);
				fileName();
				}
				break;
			case 2:
				{
				setState(185);
				folderName();
				}
				break;
			}
			setState(188);
			match(CLOSED_PARAN);
			setState(189);
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

	public static class TextContext extends ParserRuleContext {
		public List<TerminalNode> ID() { return getTokens(ImgParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(ImgParser.ID, i);
		}
		public TextContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_text; }
	}

	public final TextContext text() throws RecognitionException {
		TextContext _localctx = new TextContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_text);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(194);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(191);
					match(ID);
					}
					} 
				}
				setState(196);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
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

	public static class CommentContext extends ParserRuleContext {
		public TextContext text() {
			return getRuleContext(TextContext.class,0);
		}
		public CommentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_comment; }
	}

	public final CommentContext comment() throws RecognitionException {
		CommentContext _localctx = new CommentContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_comment);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(197);
			match(T__12);
			setState(198);
			text();
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\34\u00cb\4\2\t\2"+
		"\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\3\2\7\2.\n\2\f\2\16\2\61\13\2"+
		"\3\2\6\2\64\n\2\r\2\16\2\65\3\2\7\29\n\2\f\2\16\2<\13\2\3\2\7\2?\n\2\f"+
		"\2\16\2B\13\2\3\2\7\2E\n\2\f\2\16\2H\13\2\3\2\6\2K\n\2\r\2\16\2L\3\2\3"+
		"\2\3\3\3\3\5\3S\n\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3"+
		"\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\b\3\b\3\b\3\b\3\t"+
		"\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\5\n\u0080\n\n\3\13"+
		"\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\f\3\f"+
		"\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3\17\5\17\u009f"+
		"\n\17\3\17\3\17\3\20\3\20\3\20\3\20\3\20\3\21\3\21\3\21\3\21\3\21\3\22"+
		"\3\22\3\22\3\22\3\22\3\23\3\23\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3\24"+
		"\3\24\5\24\u00bd\n\24\3\24\3\24\3\24\3\25\7\25\u00c3\n\25\f\25\16\25\u00c6"+
		"\13\25\3\26\3\26\3\26\3\26\2\2\27\2\4\6\b\n\f\16\20\22\24\26\30\32\34"+
		"\36 \"$&(*\2\3\3\2\3\4\2\u00c7\2\63\3\2\2\2\4R\3\2\2\2\6T\3\2\2\2\b\\"+
		"\3\2\2\2\ne\3\2\2\2\fk\3\2\2\2\16m\3\2\2\2\20q\3\2\2\2\22\177\3\2\2\2"+
		"\24\u0081\3\2\2\2\26\u008c\3\2\2\2\30\u0091\3\2\2\2\32\u0095\3\2\2\2\34"+
		"\u0099\3\2\2\2\36\u00a2\3\2\2\2 \u00a7\3\2\2\2\"\u00ac\3\2\2\2$\u00b1"+
		"\3\2\2\2&\u00b6\3\2\2\2(\u00c4\3\2\2\2*\u00c7\3\2\2\2,.\5*\26\2-,\3\2"+
		"\2\2.\61\3\2\2\2/-\3\2\2\2/\60\3\2\2\2\60\62\3\2\2\2\61/\3\2\2\2\62\64"+
		"\5\4\3\2\63/\3\2\2\2\64\65\3\2\2\2\65\63\3\2\2\2\65\66\3\2\2\2\66:\3\2"+
		"\2\2\679\5*\26\28\67\3\2\2\29<\3\2\2\2:8\3\2\2\2:;\3\2\2\2;@\3\2\2\2<"+
		":\3\2\2\2=?\5\20\t\2>=\3\2\2\2?B\3\2\2\2@>\3\2\2\2@A\3\2\2\2AF\3\2\2\2"+
		"B@\3\2\2\2CE\5*\26\2DC\3\2\2\2EH\3\2\2\2FD\3\2\2\2FG\3\2\2\2GJ\3\2\2\2"+
		"HF\3\2\2\2IK\5&\24\2JI\3\2\2\2KL\3\2\2\2LJ\3\2\2\2LM\3\2\2\2MN\3\2\2\2"+
		"NO\7\2\2\3O\3\3\2\2\2PS\5\6\4\2QS\5\b\5\2RP\3\2\2\2RQ\3\2\2\2S\5\3\2\2"+
		"\2TU\7\33\2\2UV\7\20\2\2VW\7\30\2\2WX\7\25\2\2XY\5\n\6\2YZ\7\26\2\2Z["+
		"\7\24\2\2[\7\3\2\2\2\\]\7\33\2\2]^\7\27\2\2^_\7\20\2\2_`\7\30\2\2`a\7"+
		"\25\2\2ab\5\16\b\2bc\7\26\2\2cd\7\24\2\2d\t\3\2\2\2ef\7\23\2\2fg\7\33"+
		"\2\2gh\7\21\2\2hi\5\f\7\2ij\7\23\2\2j\13\3\2\2\2kl\t\2\2\2l\r\3\2\2\2"+
		"mn\7\23\2\2no\7\33\2\2op\7\23\2\2p\17\3\2\2\2qr\7\33\2\2rs\7\21\2\2st"+
		"\5\22\n\2tu\7\24\2\2u\21\3\2\2\2v\u0080\5\24\13\2w\u0080\5\26\f\2x\u0080"+
		"\5\30\r\2y\u0080\5\32\16\2z\u0080\5\34\17\2{\u0080\5\36\20\2|\u0080\5"+
		" \21\2}\u0080\5\"\22\2~\u0080\5$\23\2\177v\3\2\2\2\177w\3\2\2\2\177x\3"+
		"\2\2\2\177y\3\2\2\2\177z\3\2\2\2\177{\3\2\2\2\177|\3\2\2\2\177}\3\2\2"+
		"\2\177~\3\2\2\2\u0080\23\3\2\2\2\u0081\u0082\7\5\2\2\u0082\u0083\7\25"+
		"\2\2\u0083\u0084\7\31\2\2\u0084\u0085\7\22\2\2\u0085\u0086\7\31\2\2\u0086"+
		"\u0087\7\22\2\2\u0087\u0088\7\31\2\2\u0088\u0089\7\22\2\2\u0089\u008a"+
		"\7\31\2\2\u008a\u008b\7\26\2\2\u008b\25\3\2\2\2\u008c\u008d\7\6\2\2\u008d"+
		"\u008e\7\25\2\2\u008e\u008f\7\31\2\2\u008f\u0090\7\26\2\2\u0090\27\3\2"+
		"\2\2\u0091\u0092\7\7\2\2\u0092\u0093\7\25\2\2\u0093\u0094\7\26\2\2\u0094"+
		"\31\3\2\2\2\u0095\u0096\7\b\2\2\u0096\u0097\7\25\2\2\u0097\u0098\7\26"+
		"\2\2\u0098\33\3\2\2\2\u0099\u009a\7\t\2\2\u009a\u009b\7\25\2\2\u009b\u009e"+
		"\7\31\2\2\u009c\u009d\7\22\2\2\u009d\u009f\7\31\2\2\u009e\u009c\3\2\2"+
		"\2\u009e\u009f\3\2\2\2\u009f\u00a0\3\2\2\2\u00a0\u00a1\7\26\2\2\u00a1"+
		"\35\3\2\2\2\u00a2\u00a3\7\n\2\2\u00a3\u00a4\7\25\2\2\u00a4\u00a5\7\31"+
		"\2\2\u00a5\u00a6\7\26\2\2\u00a6\37\3\2\2\2\u00a7\u00a8\7\13\2\2\u00a8"+
		"\u00a9\7\25\2\2\u00a9\u00aa\7\31\2\2\u00aa\u00ab\7\26\2\2\u00ab!\3\2\2"+
		"\2\u00ac\u00ad\7\f\2\2\u00ad\u00ae\7\25\2\2\u00ae\u00af\7\31\2\2\u00af"+
		"\u00b0\7\26\2\2\u00b0#\3\2\2\2\u00b1\u00b2\7\r\2\2\u00b2\u00b3\7\25\2"+
		"\2\u00b3\u00b4\7\31\2\2\u00b4\u00b5\7\26\2\2\u00b5%\3\2\2\2\u00b6\u00b7"+
		"\7\33\2\2\u00b7\u00b8\7\21\2\2\u00b8\u00b9\7\16\2\2\u00b9\u00bc\7\25\2"+
		"\2\u00ba\u00bd\5\n\6\2\u00bb\u00bd\5\16\b\2\u00bc\u00ba\3\2\2\2\u00bc"+
		"\u00bb\3\2\2\2\u00bd\u00be\3\2\2\2\u00be\u00bf\7\26\2\2\u00bf\u00c0\7"+
		"\24\2\2\u00c0\'\3\2\2\2\u00c1\u00c3\7\33\2\2\u00c2\u00c1\3\2\2\2\u00c3"+
		"\u00c6\3\2\2\2\u00c4\u00c2\3\2\2\2\u00c4\u00c5\3\2\2\2\u00c5)\3\2\2\2"+
		"\u00c6\u00c4\3\2\2\2\u00c7\u00c8\7\17\2\2\u00c8\u00c9\5(\25\2\u00c9+\3"+
		"\2\2\2\r/\65:@FLR\177\u009e\u00bc\u00c4";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}