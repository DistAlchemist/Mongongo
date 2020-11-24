// Generated from /data/Mongongo/src/mql/parser/Mql.g4 by ANTLR 4.8
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class MqlParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, GET=6, SET=7, WHITESPACE=8, ASSOC=9, 
		COMMA=10, LEFT_BRACE=11, RIGHT_BRACE=12, SEMICOLON=13, Identifier=14, 
		StringLiteral=15, IntegerLiteral=16;
	public static final int
		RULE_stringVal = 0, RULE_stmt = 1, RULE_getStmt = 2, RULE_setStmt = 3, 
		RULE_columnSpec = 4, RULE_tableName = 5, RULE_columnFamilyName = 6, RULE_valueExpr = 7, 
		RULE_cellValue = 8, RULE_columnMapValue = 9, RULE_superColumnMapValue = 10, 
		RULE_columnMapEntry = 11, RULE_superColumnMapEntry = 12, RULE_columnOrSuperColumnName = 13, 
		RULE_rowKey = 14, RULE_columnOrSuperColumnKey = 15, RULE_columnKey = 16, 
		RULE_superColumnKey = 17;
	private static String[] makeRuleNames() {
		return new String[] {
			"stringVal", "stmt", "getStmt", "setStmt", "columnSpec", "tableName", 
			"columnFamilyName", "valueExpr", "cellValue", "columnMapValue", "superColumnMapValue", 
			"columnMapEntry", "superColumnMapEntry", "columnOrSuperColumnName", "rowKey", 
			"columnOrSuperColumnKey", "columnKey", "superColumnKey"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'?'", "'='", "'.'", "'['", "']'", "'GET'", "'SET'", null, "'=>'", 
			"','", "'{'", "'}'", "';'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, "GET", "SET", "WHITESPACE", "ASSOC", 
			"COMMA", "LEFT_BRACE", "RIGHT_BRACE", "SEMICOLON", "Identifier", "StringLiteral", 
			"IntegerLiteral"
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
	public String getGrammarFileName() { return "Mql.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public MqlParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class StringValContext extends ParserRuleContext {
		public TerminalNode StringLiteral() { return getToken(MqlParser.StringLiteral, 0); }
		public StringValContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stringVal; }
	}

	public final StringValContext stringVal() throws RecognitionException {
		StringValContext _localctx = new StringValContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_stringVal);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(36);
			_la = _input.LA(1);
			if ( !(_la==T__0 || _la==StringLiteral) ) {
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

	public static class StmtContext extends ParserRuleContext {
		public GetStmtContext getStmt() {
			return getRuleContext(GetStmtContext.class,0);
		}
		public SetStmtContext setStmt() {
			return getRuleContext(SetStmtContext.class,0);
		}
		public StmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stmt; }
	}

	public final StmtContext stmt() throws RecognitionException {
		StmtContext _localctx = new StmtContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_stmt);
		try {
			setState(40);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case GET:
				enterOuterAlt(_localctx, 1);
				{
				setState(38);
				getStmt();
				}
				break;
			case SET:
				enterOuterAlt(_localctx, 2);
				{
				setState(39);
				setStmt();
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

	public static class GetStmtContext extends ParserRuleContext {
		public TerminalNode GET() { return getToken(MqlParser.GET, 0); }
		public ColumnSpecContext columnSpec() {
			return getRuleContext(ColumnSpecContext.class,0);
		}
		public GetStmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_getStmt; }
	}

	public final GetStmtContext getStmt() throws RecognitionException {
		GetStmtContext _localctx = new GetStmtContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_getStmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(42);
			match(GET);
			setState(43);
			columnSpec();
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

	public static class SetStmtContext extends ParserRuleContext {
		public TerminalNode SET() { return getToken(MqlParser.SET, 0); }
		public ColumnSpecContext columnSpec() {
			return getRuleContext(ColumnSpecContext.class,0);
		}
		public ValueExprContext valueExpr() {
			return getRuleContext(ValueExprContext.class,0);
		}
		public SetStmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_setStmt; }
	}

	public final SetStmtContext setStmt() throws RecognitionException {
		SetStmtContext _localctx = new SetStmtContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_setStmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(45);
			match(SET);
			setState(46);
			columnSpec();
			setState(47);
			match(T__1);
			setState(48);
			valueExpr();
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

	public static class ColumnSpecContext extends ParserRuleContext {
		public ColumnOrSuperColumnKeyContext columnOrSuperColumnKey;
		public List<ColumnOrSuperColumnKeyContext> a = new ArrayList<ColumnOrSuperColumnKeyContext>();
		public TableNameContext tableName() {
			return getRuleContext(TableNameContext.class,0);
		}
		public ColumnFamilyNameContext columnFamilyName() {
			return getRuleContext(ColumnFamilyNameContext.class,0);
		}
		public RowKeyContext rowKey() {
			return getRuleContext(RowKeyContext.class,0);
		}
		public List<ColumnOrSuperColumnKeyContext> columnOrSuperColumnKey() {
			return getRuleContexts(ColumnOrSuperColumnKeyContext.class);
		}
		public ColumnOrSuperColumnKeyContext columnOrSuperColumnKey(int i) {
			return getRuleContext(ColumnOrSuperColumnKeyContext.class,i);
		}
		public ColumnSpecContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_columnSpec; }
	}

	public final ColumnSpecContext columnSpec() throws RecognitionException {
		ColumnSpecContext _localctx = new ColumnSpecContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_columnSpec);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(50);
			tableName();
			setState(51);
			match(T__2);
			setState(52);
			columnFamilyName();
			setState(53);
			match(T__3);
			setState(54);
			rowKey();
			setState(55);
			match(T__4);
			setState(65);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__3) {
				{
				setState(56);
				match(T__3);
				setState(57);
				((ColumnSpecContext)_localctx).columnOrSuperColumnKey = columnOrSuperColumnKey();
				((ColumnSpecContext)_localctx).a.add(((ColumnSpecContext)_localctx).columnOrSuperColumnKey);
				setState(58);
				match(T__4);
				setState(63);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__3) {
					{
					setState(59);
					match(T__3);
					setState(60);
					((ColumnSpecContext)_localctx).columnOrSuperColumnKey = columnOrSuperColumnKey();
					((ColumnSpecContext)_localctx).a.add(((ColumnSpecContext)_localctx).columnOrSuperColumnKey);
					setState(61);
					match(T__4);
					}
				}

				}
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

	public static class TableNameContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(MqlParser.Identifier, 0); }
		public TableNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tableName; }
	}

	public final TableNameContext tableName() throws RecognitionException {
		TableNameContext _localctx = new TableNameContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_tableName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(67);
			match(Identifier);
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

	public static class ColumnFamilyNameContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(MqlParser.Identifier, 0); }
		public ColumnFamilyNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_columnFamilyName; }
	}

	public final ColumnFamilyNameContext columnFamilyName() throws RecognitionException {
		ColumnFamilyNameContext _localctx = new ColumnFamilyNameContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_columnFamilyName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(69);
			match(Identifier);
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

	public static class ValueExprContext extends ParserRuleContext {
		public CellValueContext cellValue() {
			return getRuleContext(CellValueContext.class,0);
		}
		public ColumnMapValueContext columnMapValue() {
			return getRuleContext(ColumnMapValueContext.class,0);
		}
		public SuperColumnMapValueContext superColumnMapValue() {
			return getRuleContext(SuperColumnMapValueContext.class,0);
		}
		public ValueExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_valueExpr; }
	}

	public final ValueExprContext valueExpr() throws RecognitionException {
		ValueExprContext _localctx = new ValueExprContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_valueExpr);
		try {
			setState(74);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(71);
				cellValue();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(72);
				columnMapValue();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(73);
				superColumnMapValue();
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

	public static class CellValueContext extends ParserRuleContext {
		public StringValContext stringVal() {
			return getRuleContext(StringValContext.class,0);
		}
		public CellValueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cellValue; }
	}

	public final CellValueContext cellValue() throws RecognitionException {
		CellValueContext _localctx = new CellValueContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_cellValue);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(76);
			stringVal();
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

	public static class ColumnMapValueContext extends ParserRuleContext {
		public TerminalNode LEFT_BRACE() { return getToken(MqlParser.LEFT_BRACE, 0); }
		public List<ColumnMapEntryContext> columnMapEntry() {
			return getRuleContexts(ColumnMapEntryContext.class);
		}
		public ColumnMapEntryContext columnMapEntry(int i) {
			return getRuleContext(ColumnMapEntryContext.class,i);
		}
		public TerminalNode RIGHT_BRACE() { return getToken(MqlParser.RIGHT_BRACE, 0); }
		public List<TerminalNode> COMMA() { return getTokens(MqlParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(MqlParser.COMMA, i);
		}
		public ColumnMapValueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_columnMapValue; }
	}

	public final ColumnMapValueContext columnMapValue() throws RecognitionException {
		ColumnMapValueContext _localctx = new ColumnMapValueContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_columnMapValue);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(78);
			match(LEFT_BRACE);
			setState(79);
			columnMapEntry();
			setState(84);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(80);
				match(COMMA);
				setState(81);
				columnMapEntry();
				}
				}
				setState(86);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(87);
			match(RIGHT_BRACE);
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

	public static class SuperColumnMapValueContext extends ParserRuleContext {
		public TerminalNode LEFT_BRACE() { return getToken(MqlParser.LEFT_BRACE, 0); }
		public List<SuperColumnMapEntryContext> superColumnMapEntry() {
			return getRuleContexts(SuperColumnMapEntryContext.class);
		}
		public SuperColumnMapEntryContext superColumnMapEntry(int i) {
			return getRuleContext(SuperColumnMapEntryContext.class,i);
		}
		public TerminalNode RIGHT_BRACE() { return getToken(MqlParser.RIGHT_BRACE, 0); }
		public List<TerminalNode> COMMA() { return getTokens(MqlParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(MqlParser.COMMA, i);
		}
		public SuperColumnMapValueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_superColumnMapValue; }
	}

	public final SuperColumnMapValueContext superColumnMapValue() throws RecognitionException {
		SuperColumnMapValueContext _localctx = new SuperColumnMapValueContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_superColumnMapValue);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(89);
			match(LEFT_BRACE);
			setState(90);
			superColumnMapEntry();
			setState(95);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(91);
				match(COMMA);
				setState(92);
				superColumnMapEntry();
				}
				}
				setState(97);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(98);
			match(RIGHT_BRACE);
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

	public static class ColumnMapEntryContext extends ParserRuleContext {
		public ColumnKeyContext columnKey() {
			return getRuleContext(ColumnKeyContext.class,0);
		}
		public TerminalNode ASSOC() { return getToken(MqlParser.ASSOC, 0); }
		public CellValueContext cellValue() {
			return getRuleContext(CellValueContext.class,0);
		}
		public ColumnMapEntryContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_columnMapEntry; }
	}

	public final ColumnMapEntryContext columnMapEntry() throws RecognitionException {
		ColumnMapEntryContext _localctx = new ColumnMapEntryContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_columnMapEntry);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(100);
			columnKey();
			setState(101);
			match(ASSOC);
			setState(102);
			cellValue();
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

	public static class SuperColumnMapEntryContext extends ParserRuleContext {
		public SuperColumnKeyContext superColumnKey() {
			return getRuleContext(SuperColumnKeyContext.class,0);
		}
		public TerminalNode ASSOC() { return getToken(MqlParser.ASSOC, 0); }
		public ColumnMapValueContext columnMapValue() {
			return getRuleContext(ColumnMapValueContext.class,0);
		}
		public SuperColumnMapEntryContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_superColumnMapEntry; }
	}

	public final SuperColumnMapEntryContext superColumnMapEntry() throws RecognitionException {
		SuperColumnMapEntryContext _localctx = new SuperColumnMapEntryContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_superColumnMapEntry);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(104);
			superColumnKey();
			setState(105);
			match(ASSOC);
			setState(106);
			columnMapValue();
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

	public static class ColumnOrSuperColumnNameContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(MqlParser.Identifier, 0); }
		public ColumnOrSuperColumnNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_columnOrSuperColumnName; }
	}

	public final ColumnOrSuperColumnNameContext columnOrSuperColumnName() throws RecognitionException {
		ColumnOrSuperColumnNameContext _localctx = new ColumnOrSuperColumnNameContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_columnOrSuperColumnName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(108);
			match(Identifier);
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

	public static class RowKeyContext extends ParserRuleContext {
		public StringValContext stringVal() {
			return getRuleContext(StringValContext.class,0);
		}
		public RowKeyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_rowKey; }
	}

	public final RowKeyContext rowKey() throws RecognitionException {
		RowKeyContext _localctx = new RowKeyContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_rowKey);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(110);
			stringVal();
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

	public static class ColumnOrSuperColumnKeyContext extends ParserRuleContext {
		public StringValContext stringVal() {
			return getRuleContext(StringValContext.class,0);
		}
		public ColumnOrSuperColumnKeyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_columnOrSuperColumnKey; }
	}

	public final ColumnOrSuperColumnKeyContext columnOrSuperColumnKey() throws RecognitionException {
		ColumnOrSuperColumnKeyContext _localctx = new ColumnOrSuperColumnKeyContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_columnOrSuperColumnKey);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(112);
			stringVal();
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

	public static class ColumnKeyContext extends ParserRuleContext {
		public StringValContext stringVal() {
			return getRuleContext(StringValContext.class,0);
		}
		public ColumnKeyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_columnKey; }
	}

	public final ColumnKeyContext columnKey() throws RecognitionException {
		ColumnKeyContext _localctx = new ColumnKeyContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_columnKey);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(114);
			stringVal();
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

	public static class SuperColumnKeyContext extends ParserRuleContext {
		public StringValContext stringVal() {
			return getRuleContext(StringValContext.class,0);
		}
		public SuperColumnKeyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_superColumnKey; }
	}

	public final SuperColumnKeyContext superColumnKey() throws RecognitionException {
		SuperColumnKeyContext _localctx = new SuperColumnKeyContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_superColumnKey);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(116);
			stringVal();
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\22y\4\2\t\2\4\3\t"+
		"\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t\13\4"+
		"\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22\4\23"+
		"\t\23\3\2\3\2\3\3\3\3\5\3+\n\3\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\6\3\6"+
		"\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6B\n\6\5\6D\n\6\3\7\3\7"+
		"\3\b\3\b\3\t\3\t\3\t\5\tM\n\t\3\n\3\n\3\13\3\13\3\13\3\13\7\13U\n\13\f"+
		"\13\16\13X\13\13\3\13\3\13\3\f\3\f\3\f\3\f\7\f`\n\f\f\f\16\fc\13\f\3\f"+
		"\3\f\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\17\3\17\3\20\3\20\3\21\3\21"+
		"\3\22\3\22\3\23\3\23\3\23\2\2\24\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36"+
		" \"$\2\3\4\2\3\3\21\21\2m\2&\3\2\2\2\4*\3\2\2\2\6,\3\2\2\2\b/\3\2\2\2"+
		"\n\64\3\2\2\2\fE\3\2\2\2\16G\3\2\2\2\20L\3\2\2\2\22N\3\2\2\2\24P\3\2\2"+
		"\2\26[\3\2\2\2\30f\3\2\2\2\32j\3\2\2\2\34n\3\2\2\2\36p\3\2\2\2 r\3\2\2"+
		"\2\"t\3\2\2\2$v\3\2\2\2&\'\t\2\2\2\'\3\3\2\2\2(+\5\6\4\2)+\5\b\5\2*(\3"+
		"\2\2\2*)\3\2\2\2+\5\3\2\2\2,-\7\b\2\2-.\5\n\6\2.\7\3\2\2\2/\60\7\t\2\2"+
		"\60\61\5\n\6\2\61\62\7\4\2\2\62\63\5\20\t\2\63\t\3\2\2\2\64\65\5\f\7\2"+
		"\65\66\7\5\2\2\66\67\5\16\b\2\678\7\6\2\289\5\36\20\29C\7\7\2\2:;\7\6"+
		"\2\2;<\5 \21\2<A\7\7\2\2=>\7\6\2\2>?\5 \21\2?@\7\7\2\2@B\3\2\2\2A=\3\2"+
		"\2\2AB\3\2\2\2BD\3\2\2\2C:\3\2\2\2CD\3\2\2\2D\13\3\2\2\2EF\7\20\2\2F\r"+
		"\3\2\2\2GH\7\20\2\2H\17\3\2\2\2IM\5\22\n\2JM\5\24\13\2KM\5\26\f\2LI\3"+
		"\2\2\2LJ\3\2\2\2LK\3\2\2\2M\21\3\2\2\2NO\5\2\2\2O\23\3\2\2\2PQ\7\r\2\2"+
		"QV\5\30\r\2RS\7\f\2\2SU\5\30\r\2TR\3\2\2\2UX\3\2\2\2VT\3\2\2\2VW\3\2\2"+
		"\2WY\3\2\2\2XV\3\2\2\2YZ\7\16\2\2Z\25\3\2\2\2[\\\7\r\2\2\\a\5\32\16\2"+
		"]^\7\f\2\2^`\5\32\16\2_]\3\2\2\2`c\3\2\2\2a_\3\2\2\2ab\3\2\2\2bd\3\2\2"+
		"\2ca\3\2\2\2de\7\16\2\2e\27\3\2\2\2fg\5\"\22\2gh\7\13\2\2hi\5\22\n\2i"+
		"\31\3\2\2\2jk\5$\23\2kl\7\13\2\2lm\5\24\13\2m\33\3\2\2\2no\7\20\2\2o\35"+
		"\3\2\2\2pq\5\2\2\2q\37\3\2\2\2rs\5\2\2\2s!\3\2\2\2tu\5\2\2\2u#\3\2\2\2"+
		"vw\5\2\2\2w%\3\2\2\2\b*ACLVa";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}