// Generated from c:\Go\src\img\chem.g4 by ANTLR 4.8
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class chemParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		UPPERCASE=10, LOWERCASE=11, DIGIT=12, DIGITS=13, EQUAL=14, PLUS=15, WS=16;
	public static final int
		RULE_varName = 0, RULE_program = 1, RULE_commandStatementList = 2, RULE_commandStatement = 3, 
		RULE_printCommand = 4, RULE_balanceCommand = 5, RULE_solveCommand = 6, 
		RULE_showCommand = 7, RULE_varCommand = 8, RULE_declReaction = 9, RULE_declPartial = 10, 
		RULE_reaction = 11, RULE_partialReaction = 12, RULE_side = 13, RULE_molecule = 14, 
		RULE_element = 15, RULE_number = 16;
	private static String[] makeRuleNames() {
		return new String[] {
			"varName", "program", "commandStatementList", "commandStatement", "printCommand", 
			"balanceCommand", "solveCommand", "showCommand", "varCommand", "declReaction", 
			"declPartial", "reaction", "partialReaction", "side", "molecule", "element", 
			"number"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'PRINT ('", "')'", "'BALANCE'", "'('", "'SOLVE'", "'SHOW'", "'VAR'", 
			"'{'", "'}'", null, null, null, null, "'='", "'+'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, "UPPERCASE", 
			"LOWERCASE", "DIGIT", "DIGITS", "EQUAL", "PLUS", "WS"
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
	public String getGrammarFileName() { return "chem.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public chemParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class VarNameContext extends ParserRuleContext {
		public List<TerminalNode> LOWERCASE() { return getTokens(chemParser.LOWERCASE); }
		public TerminalNode LOWERCASE(int i) {
			return getToken(chemParser.LOWERCASE, i);
		}
		public VarNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_varName; }
	}

	public final VarNameContext varName() throws RecognitionException {
		VarNameContext _localctx = new VarNameContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_varName);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(35); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(34);
				match(LOWERCASE);
				}
				}
				setState(37); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==LOWERCASE );
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

	public static class ProgramContext extends ParserRuleContext {
		public CommandStatementListContext commandStatementList() {
			return getRuleContext(CommandStatementListContext.class,0);
		}
		public ProgramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_program; }
	}

	public final ProgramContext program() throws RecognitionException {
		ProgramContext _localctx = new ProgramContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_program);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(39);
			commandStatementList();
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

	public static class CommandStatementListContext extends ParserRuleContext {
		public List<CommandStatementContext> commandStatement() {
			return getRuleContexts(CommandStatementContext.class);
		}
		public CommandStatementContext commandStatement(int i) {
			return getRuleContext(CommandStatementContext.class,i);
		}
		public CommandStatementListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_commandStatementList; }
	}

	public final CommandStatementListContext commandStatementList() throws RecognitionException {
		CommandStatementListContext _localctx = new CommandStatementListContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_commandStatementList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(42); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(41);
				commandStatement();
				}
				}
				setState(44); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__2) | (1L << T__4) | (1L << T__5) | (1L << T__6))) != 0) );
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

	public static class CommandStatementContext extends ParserRuleContext {
		public PrintCommandContext printCommand() {
			return getRuleContext(PrintCommandContext.class,0);
		}
		public BalanceCommandContext balanceCommand() {
			return getRuleContext(BalanceCommandContext.class,0);
		}
		public SolveCommandContext solveCommand() {
			return getRuleContext(SolveCommandContext.class,0);
		}
		public ShowCommandContext showCommand() {
			return getRuleContext(ShowCommandContext.class,0);
		}
		public VarCommandContext varCommand() {
			return getRuleContext(VarCommandContext.class,0);
		}
		public CommandStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_commandStatement; }
	}

	public final CommandStatementContext commandStatement() throws RecognitionException {
		CommandStatementContext _localctx = new CommandStatementContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_commandStatement);
		try {
			setState(51);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__0:
				enterOuterAlt(_localctx, 1);
				{
				setState(46);
				printCommand();
				}
				break;
			case T__2:
				enterOuterAlt(_localctx, 2);
				{
				setState(47);
				balanceCommand();
				}
				break;
			case T__4:
				enterOuterAlt(_localctx, 3);
				{
				setState(48);
				solveCommand();
				}
				break;
			case T__5:
				enterOuterAlt(_localctx, 4);
				{
				setState(49);
				showCommand();
				}
				break;
			case T__6:
				enterOuterAlt(_localctx, 5);
				{
				setState(50);
				varCommand();
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

	public static class PrintCommandContext extends ParserRuleContext {
		public VarNameContext varName() {
			return getRuleContext(VarNameContext.class,0);
		}
		public PrintCommandContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_printCommand; }
	}

	public final PrintCommandContext printCommand() throws RecognitionException {
		PrintCommandContext _localctx = new PrintCommandContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_printCommand);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(53);
			match(T__0);
			setState(54);
			varName();
			setState(55);
			match(T__1);
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

	public static class BalanceCommandContext extends ParserRuleContext {
		public VarNameContext varName() {
			return getRuleContext(VarNameContext.class,0);
		}
		public BalanceCommandContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_balanceCommand; }
	}

	public final BalanceCommandContext balanceCommand() throws RecognitionException {
		BalanceCommandContext _localctx = new BalanceCommandContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_balanceCommand);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(57);
			match(T__2);
			setState(58);
			match(T__3);
			setState(59);
			varName();
			setState(60);
			match(T__1);
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

	public static class SolveCommandContext extends ParserRuleContext {
		public VarNameContext varName() {
			return getRuleContext(VarNameContext.class,0);
		}
		public SolveCommandContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_solveCommand; }
	}

	public final SolveCommandContext solveCommand() throws RecognitionException {
		SolveCommandContext _localctx = new SolveCommandContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_solveCommand);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(62);
			match(T__4);
			setState(63);
			match(T__3);
			setState(64);
			varName();
			setState(65);
			match(T__1);
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

	public static class ShowCommandContext extends ParserRuleContext {
		public VarNameContext varName() {
			return getRuleContext(VarNameContext.class,0);
		}
		public ShowCommandContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_showCommand; }
	}

	public final ShowCommandContext showCommand() throws RecognitionException {
		ShowCommandContext _localctx = new ShowCommandContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_showCommand);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(67);
			match(T__5);
			setState(68);
			match(T__3);
			setState(69);
			varName();
			setState(70);
			match(T__1);
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

	public static class VarCommandContext extends ParserRuleContext {
		public DeclReactionContext declReaction() {
			return getRuleContext(DeclReactionContext.class,0);
		}
		public DeclPartialContext declPartial() {
			return getRuleContext(DeclPartialContext.class,0);
		}
		public VarCommandContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_varCommand; }
	}

	public final VarCommandContext varCommand() throws RecognitionException {
		VarCommandContext _localctx = new VarCommandContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_varCommand);
		try {
			setState(74);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(72);
				declReaction();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(73);
				declPartial();
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

	public static class DeclReactionContext extends ParserRuleContext {
		public VarNameContext varName() {
			return getRuleContext(VarNameContext.class,0);
		}
		public TerminalNode EQUAL() { return getToken(chemParser.EQUAL, 0); }
		public ReactionContext reaction() {
			return getRuleContext(ReactionContext.class,0);
		}
		public DeclReactionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declReaction; }
	}

	public final DeclReactionContext declReaction() throws RecognitionException {
		DeclReactionContext _localctx = new DeclReactionContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_declReaction);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(76);
			match(T__6);
			setState(77);
			varName();
			setState(78);
			match(EQUAL);
			setState(79);
			match(T__7);
			setState(80);
			reaction();
			setState(81);
			match(T__8);
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

	public static class DeclPartialContext extends ParserRuleContext {
		public VarNameContext varName() {
			return getRuleContext(VarNameContext.class,0);
		}
		public TerminalNode EQUAL() { return getToken(chemParser.EQUAL, 0); }
		public PartialReactionContext partialReaction() {
			return getRuleContext(PartialReactionContext.class,0);
		}
		public DeclPartialContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declPartial; }
	}

	public final DeclPartialContext declPartial() throws RecognitionException {
		DeclPartialContext _localctx = new DeclPartialContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_declPartial);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(83);
			match(T__6);
			setState(84);
			varName();
			setState(85);
			match(EQUAL);
			setState(86);
			match(T__7);
			setState(87);
			partialReaction();
			setState(88);
			match(T__8);
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

	public static class ReactionContext extends ParserRuleContext {
		public List<SideContext> side() {
			return getRuleContexts(SideContext.class);
		}
		public SideContext side(int i) {
			return getRuleContext(SideContext.class,i);
		}
		public TerminalNode EQUAL() { return getToken(chemParser.EQUAL, 0); }
		public ReactionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_reaction; }
	}

	public final ReactionContext reaction() throws RecognitionException {
		ReactionContext _localctx = new ReactionContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_reaction);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(90);
			side(0);
			setState(91);
			match(EQUAL);
			setState(92);
			side(0);
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

	public static class PartialReactionContext extends ParserRuleContext {
		public SideContext side() {
			return getRuleContext(SideContext.class,0);
		}
		public PartialReactionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_partialReaction; }
	}

	public final PartialReactionContext partialReaction() throws RecognitionException {
		PartialReactionContext _localctx = new PartialReactionContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_partialReaction);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(94);
			side(0);
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

	public static class SideContext extends ParserRuleContext {
		public MoleculeContext molecule() {
			return getRuleContext(MoleculeContext.class,0);
		}
		public NumberContext number() {
			return getRuleContext(NumberContext.class,0);
		}
		public List<SideContext> side() {
			return getRuleContexts(SideContext.class);
		}
		public SideContext side(int i) {
			return getRuleContext(SideContext.class,i);
		}
		public TerminalNode PLUS() { return getToken(chemParser.PLUS, 0); }
		public SideContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_side; }
	}

	public final SideContext side() throws RecognitionException {
		return side(0);
	}

	private SideContext side(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		SideContext _localctx = new SideContext(_ctx, _parentState);
		SideContext _prevctx = _localctx;
		int _startState = 26;
		enterRecursionRule(_localctx, 26, RULE_side, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(98);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DIGIT) {
				{
				setState(97);
				number();
				}
			}

			setState(100);
			molecule(0);
			}
			_ctx.stop = _input.LT(-1);
			setState(107);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new SideContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_side);
					setState(102);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(103);
					match(PLUS);
					setState(104);
					side(3);
					}
					} 
				}
				setState(109);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class MoleculeContext extends ParserRuleContext {
		public List<ElementContext> element() {
			return getRuleContexts(ElementContext.class);
		}
		public ElementContext element(int i) {
			return getRuleContext(ElementContext.class,i);
		}
		public List<NumberContext> number() {
			return getRuleContexts(NumberContext.class);
		}
		public NumberContext number(int i) {
			return getRuleContext(NumberContext.class,i);
		}
		public List<MoleculeContext> molecule() {
			return getRuleContexts(MoleculeContext.class);
		}
		public MoleculeContext molecule(int i) {
			return getRuleContext(MoleculeContext.class,i);
		}
		public MoleculeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_molecule; }
	}

	public final MoleculeContext molecule() throws RecognitionException {
		return molecule(0);
	}

	private MoleculeContext molecule(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		MoleculeContext _localctx = new MoleculeContext(_ctx, _parentState);
		MoleculeContext _prevctx = _localctx;
		int _startState = 28;
		enterRecursionRule(_localctx, 28, RULE_molecule, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(115); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(111);
					element();
					setState(113);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
					case 1:
						{
						setState(112);
						number();
						}
						break;
					}
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(117); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
			} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
			}
			_ctx.stop = _input.LT(-1);
			setState(127);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new MoleculeContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_molecule);
					setState(119);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(120);
					match(T__3);
					setState(121);
					molecule(0);
					setState(122);
					match(T__1);
					setState(123);
					number();
					}
					} 
				}
				setState(129);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class ElementContext extends ParserRuleContext {
		public TerminalNode UPPERCASE() { return getToken(chemParser.UPPERCASE, 0); }
		public TerminalNode LOWERCASE() { return getToken(chemParser.LOWERCASE, 0); }
		public ElementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_element; }
	}

	public final ElementContext element() throws RecognitionException {
		ElementContext _localctx = new ElementContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_element);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(130);
			match(UPPERCASE);
			setState(132);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				{
				setState(131);
				match(LOWERCASE);
				}
				break;
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

	public static class NumberContext extends ParserRuleContext {
		public TerminalNode DIGIT() { return getToken(chemParser.DIGIT, 0); }
		public List<TerminalNode> DIGITS() { return getTokens(chemParser.DIGITS); }
		public TerminalNode DIGITS(int i) {
			return getToken(chemParser.DIGITS, i);
		}
		public NumberContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_number; }
	}

	public final NumberContext number() throws RecognitionException {
		NumberContext _localctx = new NumberContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_number);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(134);
			match(DIGIT);
			setState(138);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(135);
					match(DIGITS);
					}
					} 
				}
				setState(140);
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

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 13:
			return side_sempred((SideContext)_localctx, predIndex);
		case 14:
			return molecule_sempred((MoleculeContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean side_sempred(SideContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean molecule_sempred(MoleculeContext _localctx, int predIndex) {
		switch (predIndex) {
		case 1:
			return precpred(_ctx, 1);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\22\u0090\4\2\t\2"+
		"\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\3\2\6\2&\n\2\r\2\16\2\'\3\3\3\3\3\4\6\4-\n\4\r\4\16\4.\3\5\3\5\3\5\3"+
		"\5\3\5\5\5\66\n\5\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b"+
		"\3\b\3\t\3\t\3\t\3\t\3\t\3\n\3\n\5\nM\n\n\3\13\3\13\3\13\3\13\3\13\3\13"+
		"\3\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\16\3\16\3\17\3\17"+
		"\5\17e\n\17\3\17\3\17\3\17\3\17\3\17\7\17l\n\17\f\17\16\17o\13\17\3\20"+
		"\3\20\3\20\5\20t\n\20\6\20v\n\20\r\20\16\20w\3\20\3\20\3\20\3\20\3\20"+
		"\3\20\7\20\u0080\n\20\f\20\16\20\u0083\13\20\3\21\3\21\5\21\u0087\n\21"+
		"\3\22\3\22\7\22\u008b\n\22\f\22\16\22\u008e\13\22\3\22\2\4\34\36\23\2"+
		"\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \"\2\2\2\u008c\2%\3\2\2\2\4)\3\2"+
		"\2\2\6,\3\2\2\2\b\65\3\2\2\2\n\67\3\2\2\2\f;\3\2\2\2\16@\3\2\2\2\20E\3"+
		"\2\2\2\22L\3\2\2\2\24N\3\2\2\2\26U\3\2\2\2\30\\\3\2\2\2\32`\3\2\2\2\34"+
		"b\3\2\2\2\36p\3\2\2\2 \u0084\3\2\2\2\"\u0088\3\2\2\2$&\7\r\2\2%$\3\2\2"+
		"\2&\'\3\2\2\2\'%\3\2\2\2\'(\3\2\2\2(\3\3\2\2\2)*\5\6\4\2*\5\3\2\2\2+-"+
		"\5\b\5\2,+\3\2\2\2-.\3\2\2\2.,\3\2\2\2./\3\2\2\2/\7\3\2\2\2\60\66\5\n"+
		"\6\2\61\66\5\f\7\2\62\66\5\16\b\2\63\66\5\20\t\2\64\66\5\22\n\2\65\60"+
		"\3\2\2\2\65\61\3\2\2\2\65\62\3\2\2\2\65\63\3\2\2\2\65\64\3\2\2\2\66\t"+
		"\3\2\2\2\678\7\3\2\289\5\2\2\29:\7\4\2\2:\13\3\2\2\2;<\7\5\2\2<=\7\6\2"+
		"\2=>\5\2\2\2>?\7\4\2\2?\r\3\2\2\2@A\7\7\2\2AB\7\6\2\2BC\5\2\2\2CD\7\4"+
		"\2\2D\17\3\2\2\2EF\7\b\2\2FG\7\6\2\2GH\5\2\2\2HI\7\4\2\2I\21\3\2\2\2J"+
		"M\5\24\13\2KM\5\26\f\2LJ\3\2\2\2LK\3\2\2\2M\23\3\2\2\2NO\7\t\2\2OP\5\2"+
		"\2\2PQ\7\20\2\2QR\7\n\2\2RS\5\30\r\2ST\7\13\2\2T\25\3\2\2\2UV\7\t\2\2"+
		"VW\5\2\2\2WX\7\20\2\2XY\7\n\2\2YZ\5\32\16\2Z[\7\13\2\2[\27\3\2\2\2\\]"+
		"\5\34\17\2]^\7\20\2\2^_\5\34\17\2_\31\3\2\2\2`a\5\34\17\2a\33\3\2\2\2"+
		"bd\b\17\1\2ce\5\"\22\2dc\3\2\2\2de\3\2\2\2ef\3\2\2\2fg\5\36\20\2gm\3\2"+
		"\2\2hi\f\4\2\2ij\7\21\2\2jl\5\34\17\5kh\3\2\2\2lo\3\2\2\2mk\3\2\2\2mn"+
		"\3\2\2\2n\35\3\2\2\2om\3\2\2\2pu\b\20\1\2qs\5 \21\2rt\5\"\22\2sr\3\2\2"+
		"\2st\3\2\2\2tv\3\2\2\2uq\3\2\2\2vw\3\2\2\2wu\3\2\2\2wx\3\2\2\2x\u0081"+
		"\3\2\2\2yz\f\3\2\2z{\7\6\2\2{|\5\36\20\2|}\7\4\2\2}~\5\"\22\2~\u0080\3"+
		"\2\2\2\177y\3\2\2\2\u0080\u0083\3\2\2\2\u0081\177\3\2\2\2\u0081\u0082"+
		"\3\2\2\2\u0082\37\3\2\2\2\u0083\u0081\3\2\2\2\u0084\u0086\7\f\2\2\u0085"+
		"\u0087\7\r\2\2\u0086\u0085\3\2\2\2\u0086\u0087\3\2\2\2\u0087!\3\2\2\2"+
		"\u0088\u008c\7\16\2\2\u0089\u008b\7\17\2\2\u008a\u0089\3\2\2\2\u008b\u008e"+
		"\3\2\2\2\u008c\u008a\3\2\2\2\u008c\u008d\3\2\2\2\u008d#\3\2\2\2\u008e"+
		"\u008c\3\2\2\2\r\'.\65Ldmsw\u0081\u0086\u008c";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}