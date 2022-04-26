package token

import (
	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer/stateful"
)

var rules = []stateful.Rule{
	{
		Name:    "EOL",
		Pattern: `\n`,
	},
	{
		Name:    "WHITESPACE",
		Pattern: `[\t\r\f ]+`,
	},
	{
		Name:    "IDENT",
		Pattern: `([^-!\s:\|&"\?\*\\\^~\(\)\{\}\[\]\+\/><=0-9\.]|(\\(\s|:|&|\||\?|\*|\\|\^|~|\(|\)|!|\[|\]|\{|\}|\+|-|\/|>|<|=)))+`,
	},
	{
		Name:    "DOT",
		Pattern: `\.`,
	},
	{
		Name:    "NUMBER",
		Pattern: `[0-9]+`,
	},
	{
		Name:    "QUOTE",
		Pattern: `"`,
	},
	{
		Name:    "SLASH",
		Pattern: `\/`,
	},
	{
		Name:    "REVERSE",
		Pattern: `\\`,
	},
	{
		Name:    "COLON",
		Pattern: `:`,
	},
	{
		Name:    "COMPARE",
		Pattern: `[<>]=?`,
	},
	{
		Name:    "PLUS",
		Pattern: `\+`,
	},
	{
		Name:    "MINUS",
		Pattern: `-`,
	},
	{
		Name:    "FUZZY",
		Pattern: `~`,
	},
	{
		Name:    "BOOST",
		Pattern: `\^`,
	},
	{
		Name:    "WILDCARD",
		Pattern: `[\?\*]`,
	},
	{
		Name:    "LPAREN",
		Pattern: `\(`,
	},
	{
		Name:    "RPAREN",
		Pattern: `\)`,
	},
	{
		Name:    "LBRACK",
		Pattern: `\[`,
	},
	{
		Name:    "RBRACK",
		Pattern: `\]`,
	},
	{
		Name:    "LBRACE",
		Pattern: `\{`,
	},
	{
		Name:    "RBRACE",
		Pattern: `\}`,
	},
	{
		Name:    "AND",
		Pattern: `&`,
	},
	{
		Name:    "SOR",
		Pattern: `\|`,
	},
	{
		Name:    "NOT",
		Pattern: `!`,
	},
}

var Lexer *stateful.Definition
var Scanner *participle.Parser

func Scan(exp string) []*Token {
	var tokens = []*Token{}
	var ch = make(chan *Token, 100)
	if err := Scanner.ParseString(exp, ch); err != nil {
		return nil
	} else {
		for c := range ch {
			tokens = append(tokens, c)
		}
		return tokens
	}
}

func init() {
	Lexer, _ = stateful.NewSimple(rules)
	Scanner = participle.MustBuild(
		&Token{},
		participle.Lexer(Lexer),
	)
}

type Token struct {
	EOL        string `parser:"  @EOL" json:"eol"`
	WHITESPACE string `parser:"| @WHITESPACE" json:"whitespace"`
	IDENT      string `parser:"| @IDENT" json:"ident"`
	DOT        string `parser:"| @DOT" json:"dot"`
	NUMBER     string `parser:"| @NUMBER" json:"number"`
	QUOTE      string `parser:"| @QUOTE" json:"quote"`
	SLASH      string `parser:"| @SLASH" json:"slash"`
	REVERSE    string `parser:"| @REVERSE" json:"reverse"`
	COLON      string `parser:"| @COLON" json:"colon"`
	COMPARE    string `parser:"| @COMPARE" json:"compare"`
	PLUS       string `parser:"| @PLUS" json:"plus"`
	MINUS      string `parser:"| @MINUS" json:"minus"`
	FUZZY      string `parser:"| @FUZZY" json:"fuzzy"`
	BOOST      string `parser:"| @BOOST" json:"boost"`
	WILDCARD   string `parser:"| @WILDCARD" json:"wildcard"`
	LPAREN     string `parser:"| @LPAREN" json:"lparen"`
	RPAREN     string `parser:"| @RPAREN" json:"rparen"`
	LBRACK     string `parser:"| @LBRACK" json:"lbrack"`
	RBRACK     string `parser:"| @RBRACK" json:"rbrack"`
	LBRACE     string `parser:"| @LBRACE" json:"lbrace"`
	RBRACE     string `parser:"| @RBRACE" json:"rbrace"`
	AND        string `parser:"| @AND" json:"and"`
	SOR        string `parser:"| @SOR" json:"sor"`
	NOT        string `parser:"| @NOT" json:"not"`
}

func (t *Token) String() string {
	var res = ""
	if t == nil {
		res = ""
	} else if t.EOL != "" {
		res = t.EOL
	} else if t.WHITESPACE != "" {
		res = t.WHITESPACE
	} else if t.IDENT != "" {
		res = t.IDENT
	} else if t.DOT != "" {
		res = t.DOT
	} else if t.NUMBER != "" {
		res = t.NUMBER
	} else if t.QUOTE != "" {
		res = t.QUOTE
	} else if t.SLASH != "" {
		res = t.SLASH
	} else if t.REVERSE != "" {
		res = t.REVERSE
	} else if t.COLON != "" {
		res = t.COLON
	} else if t.COMPARE != "" {
		res = t.COMPARE
	} else if t.PLUS != "" {
		res = t.PLUS
	} else if t.MINUS != "" {
		res = t.MINUS
	} else if t.FUZZY != "" {
		res = t.FUZZY
	} else if t.BOOST != "" {
		res = t.BOOST
	} else if t.WILDCARD != "" {
		res = t.WILDCARD
	} else if t.LPAREN != "" {
		res = t.LPAREN
	} else if t.RPAREN != "" {
		res = t.RPAREN
	} else if t.LBRACK != "" {
		res = t.LBRACK
	} else if t.RBRACK != "" {
		res = t.RBRACK
	} else if t.LBRACE != "" {
		res = t.LBRACE
	} else if t.RBRACE != "" {
		res = t.RBRACE
	} else if t.AND != "" {
		res = t.AND
	} else if t.SOR != "" {
		res = t.SOR
	} else if t.NOT != "" {
		res = t.NOT
	} else {
		return ""
	}
	return res
}

func (t *Token) GetTokenType() TokenType {
	if t == nil {
		return UNKNOWN_TOKEN_TYPE
	} else if t.EOL != "" {
		return EOL_TOKEN_TYPE
	} else if t.WHITESPACE != "" {
		return WHITESPACE_TOKEN_TYPE
	} else if t.IDENT != "" {
		return IDENT_TOKEN_TYPE
	} else if t.DOT != "" {
		return DOT_TOKEN_TYPE
	} else if t.NUMBER != "" {
		return NUMBER_TOKEN_TYPE
	} else if t.QUOTE != "" {
		return QUOTE_TOKEN_TYPE
	} else if t.SLASH != "" {
		return SLASH_TOKEN_TYPE
	} else if t.REVERSE != "" {
		return REVERSE_TOKEN_TYPE
	} else if t.COLON != "" {
		return COLON_TOKEN_TYPE
	} else if t.COMPARE != "" {
		return COMPARE_TOKEN_TYPE
	} else if t.PLUS != "" {
		return PLUS_TOKEN_TYPE
	} else if t.MINUS != "" {
		return MINUS_TOKEN_TYPE
	} else if t.FUZZY != "" {
		return FUZZY_TOKEN_TYPE
	} else if t.BOOST != "" {
		return BOOST_TOKEN_TYPE
	} else if t.WILDCARD != "" {
		return WILDCARD_TOKEN_TYPE
	} else if t.LPAREN != "" {
		return LPAREN_TOKEN_TYPE
	} else if t.RPAREN != "" {
		return RPAREN_TOKEN_TYPE
	} else if t.LBRACK != "" {
		return LBRACK_TOKEN_TYPE
	} else if t.RBRACK != "" {
		return RBRACK_TOKEN_TYPE
	} else if t.LBRACE != "" {
		return LBRACE_TOKEN_TYPE
	} else if t.RBRACE != "" {
		return RBRACE_TOKEN_TYPE
	} else if t.AND != "" {
		return AND_TOKEN_TYPE
	} else if t.SOR != "" {
		return SOR_TOKEN_TYPE
	} else if t.NOT != "" {
		return NOT_TOKEN_TYPE
	} else {
		return UNKNOWN_TOKEN_TYPE
	}
}
