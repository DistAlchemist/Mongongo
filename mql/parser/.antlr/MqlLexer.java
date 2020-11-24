// Generated from /home/gqz/projects/go/src/github.com/DistAlchemist/Mongongo/mql/parser/Mql.g4 by ANTLR 4.8
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class MqlLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, GET=6, SET=7, WHITESPACE=8, ASSOC=9, 
		COMMA=10, LEFT_BRACE=11, RIGHT_BRACE=12, SEMICOLON=13, Identifier=14, 
		StringLiteral=15, IntegerLiteral=16;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "T__4", "GET", "SET", "WHITESPACE", "ASSOC", 
			"COMMA", "LEFT_BRACE", "RIGHT_BRACE", "SEMICOLON", "Letter", "Digit", 
			"Identifier", "StringLiteral", "IntegerLiteral"
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


	public MqlLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Mql.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\22r\b\1\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3\5\3\6\3\6\3\7\3\7\3\7\3\7\3\b"+
		"\3\b\3\b\3\b\3\t\6\t;\n\t\r\t\16\t<\3\t\3\t\3\n\3\n\3\n\3\13\3\13\3\f"+
		"\3\f\3\r\3\r\3\16\3\16\3\17\3\17\3\20\3\20\3\21\3\21\3\21\3\21\7\21T\n"+
		"\21\f\21\16\21W\13\21\3\22\3\22\7\22[\n\22\f\22\16\22^\13\22\3\22\3\22"+
		"\3\22\7\22c\n\22\f\22\16\22f\13\22\3\22\7\22i\n\22\f\22\16\22l\13\22\3"+
		"\23\6\23o\n\23\r\23\16\23p\2\2\24\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23"+
		"\13\25\f\27\r\31\16\33\17\35\2\37\2!\20#\21%\22\3\2\5\5\2\13\f\17\17\""+
		"\"\4\2C\\c|\3\2))\2w\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2"+
		"\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3"+
		"\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2"+
		"%\3\2\2\2\3\'\3\2\2\2\5)\3\2\2\2\7+\3\2\2\2\t-\3\2\2\2\13/\3\2\2\2\r\61"+
		"\3\2\2\2\17\65\3\2\2\2\21:\3\2\2\2\23@\3\2\2\2\25C\3\2\2\2\27E\3\2\2\2"+
		"\31G\3\2\2\2\33I\3\2\2\2\35K\3\2\2\2\37M\3\2\2\2!O\3\2\2\2#X\3\2\2\2%"+
		"n\3\2\2\2\'(\7A\2\2(\4\3\2\2\2)*\7?\2\2*\6\3\2\2\2+,\7\60\2\2,\b\3\2\2"+
		"\2-.\7]\2\2.\n\3\2\2\2/\60\7_\2\2\60\f\3\2\2\2\61\62\7I\2\2\62\63\7G\2"+
		"\2\63\64\7V\2\2\64\16\3\2\2\2\65\66\7U\2\2\66\67\7G\2\2\678\7V\2\28\20"+
		"\3\2\2\29;\t\2\2\2:9\3\2\2\2;<\3\2\2\2<:\3\2\2\2<=\3\2\2\2=>\3\2\2\2>"+
		"?\b\t\2\2?\22\3\2\2\2@A\7?\2\2AB\7@\2\2B\24\3\2\2\2CD\7.\2\2D\26\3\2\2"+
		"\2EF\7}\2\2F\30\3\2\2\2GH\7\177\2\2H\32\3\2\2\2IJ\7=\2\2J\34\3\2\2\2K"+
		"L\t\3\2\2L\36\3\2\2\2MN\4\62;\2N \3\2\2\2OU\5\35\17\2PT\5\35\17\2QT\5"+
		"\37\20\2RT\7a\2\2SP\3\2\2\2SQ\3\2\2\2SR\3\2\2\2TW\3\2\2\2US\3\2\2\2UV"+
		"\3\2\2\2V\"\3\2\2\2WU\3\2\2\2X\\\7)\2\2Y[\n\4\2\2ZY\3\2\2\2[^\3\2\2\2"+
		"\\Z\3\2\2\2\\]\3\2\2\2]_\3\2\2\2^\\\3\2\2\2_j\7)\2\2`d\7)\2\2ac\n\4\2"+
		"\2ba\3\2\2\2cf\3\2\2\2db\3\2\2\2de\3\2\2\2eg\3\2\2\2fd\3\2\2\2gi\7)\2"+
		"\2h`\3\2\2\2il\3\2\2\2jh\3\2\2\2jk\3\2\2\2k$\3\2\2\2lj\3\2\2\2mo\5\37"+
		"\20\2nm\3\2\2\2op\3\2\2\2pn\3\2\2\2pq\3\2\2\2q&\3\2\2\2\n\2<SU\\djp\3"+
		"\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}