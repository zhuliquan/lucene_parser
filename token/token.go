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
		Pattern: `([^-!\s:\|&"\?\*\\\^~\(\)\{\}\[\]\+\/><=0-9\.])+`,
	},
	{
		Name:    "ESCAPE",
		Pattern: `(\\(\s|:|&|\||\?|\*|\\|\^|~|\(|\)|!|\[|\]|\{|\}|\+|-|\/|>|<|=))+`,
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

func init() {
	Lexer, _ = stateful.NewSimple(rules)
}

var Scanner *participle.Parser

type Token struct {
	EOL        *string `parser:"  @EOL" json:"eol"`
	WHITESPACE *string `parser:"| @WHITESPACE" json:"whitespace"`
	IDENT      *string `parser:"| @IDENT" json:"ident"`
	ESCAPE     *string `parser:"| @ESCAPE" json:"escape"`
	DOT        *string `parser:"| @DOT" json:"dot"`
	NUMBER     *string `parser:"| @NUMBER" json:"number"`
	QUOTE      *string `parser:"| @QUOTE" json:"quote"`
	SLASH      *string `parser:"| @SLASH" json:"slash"`
	REVERSE    *string `parser:"| @REVERSE" json:"reverse"`
	COLON      *string `parser:"| @COLON" json:"colon"`
	COMPARE    *string `parser:"| @COMPARE" json:"compare"`
	PLUS       *string `parser:"| @PLUS" json:"plus"`
	MINUS      *string `parser:"| @MINUS" json:"minus"`
	FUZZY      *string `parser:"| @FUZZY" json:"fuzzy"`
	BOOST      *string `parser:"| @BOOST" json:"boost"`
	WILDCARD   *string `parser:"| @WILDCARD" json:"wildcard"`
	LPAREN     *string `parser:"| @LPAREN" json:"lparen"`
	RPAREN     *string `parser:"| @RPAREN" json:"rparen"`
	LBRACK     *string `parser:"| @LBRACK" json:"lbrack"`
	RBRACK     *string `parser:"| @RBRACK" json:"rbrack"`
	LBRACE     *string `parser:"| @LBRACE" json:"lbrace"`
	RBRACE     *string `parser:"| @RBRACE" json:"rbrace"`
	AND        *string `parser:"| @AND" json:"and"`
	SOR        *string `parser:"| @SOR" json:"sor"`
	NOT        *string `parser:"| @NOT" json:"not"`
}

func (t *Token) String() string {
	var res = ""
	if t == nil {
		res = ""
	} else if t.EOL != nil {
		res = *t.EOL
	} else if t.WHITESPACE != nil {
		res = *t.WHITESPACE
	} else if t.IDENT != nil {
		res = *t.IDENT
	} else if t.ESCAPE != nil {
		res = *t.ESCAPE
	} else if t.DOT != nil {
		res = *t.DOT
	} else if t.NUMBER != nil {
		res = *t.NUMBER
	} else if t.QUOTE != nil {
		res = *t.QUOTE
	} else if t.SLASH != nil {
		res = *t.SLASH
	} else if t.REVERSE != nil {
		res = *t.REVERSE
	} else if t.COLON != nil {
		res = *t.COLON
	} else if t.COMPARE != nil {
		res = *t.COMPARE
	} else if t.PLUS != nil {
		res = *t.PLUS
	} else if t.MINUS != nil {
		res = *t.MINUS
	} else if t.FUZZY != nil {
		res = *t.FUZZY
	} else if t.BOOST != nil {
		res = *t.BOOST
	} else if t.WILDCARD != nil {
		res = *t.WILDCARD
	} else if t.LPAREN != nil {
		res = *t.LPAREN
	} else if t.RPAREN != nil {
		res = *t.RPAREN
	} else if t.LBRACK != nil {
		res = *t.LBRACK
	} else if t.RBRACK != nil {
		res = *t.RBRACK
	} else if t.LBRACE != nil {
		res = *t.LBRACE
	} else if t.RBRACE != nil {
		res = *t.RBRACE
	} else if t.AND != nil {
		res = *t.AND
	} else if t.SOR != nil {
		res = *t.SOR
	} else if t.NOT != nil {
		res = *t.NOT
	} else {
		return ""
	}
	return res
}

func (t *Token) getTokenType() TokenType {
	if t == nil {
		return UNKNOWN_TOKEN_TYPE
	} else if t.EOL != nil {
		return EOL_TOKEN_TYPE
	} else if t.WHITESPACE != nil {
		return WHITESPACE_TOKEN_TYPE
	} else if t.IDENT != nil {
		return IDENT_TOKEN_TYPE
	} else if t.ESCAPE != nil {
		return ESCAPE_TOKEN_TYPE
	} else if t.DOT != nil {
		return DOT_TOKEN_TYPE
	} else if t.NUMBER != nil {
		return NUMBER_TOKEN_TYPE
	} else if t.QUOTE != nil {
		return QUOTE_TOKEN_TYPE
	} else if t.SLASH != nil {
		return SLASH_TOKEN_TYPE
	} else if t.REVERSE != nil {
		return REVERSE_TOKEN_TYPE
	} else if t.COLON != nil {
		return COLON_TOKEN_TYPE
	} else if t.COMPARE != nil {
		return COMPARE_TOKEN_TYPE
	} else if t.PLUS != nil {
		return PLUS_TOKEN_TYPE
	} else if t.MINUS != nil {
		return MINUS_TOKEN_TYPE
	} else if t.FUZZY != nil {
		return FUZZY_TOKEN_TYPE
	} else if t.BOOST != nil {
		return BOOST_TOKEN_TYPE
	} else if t.WILDCARD != nil {
		return WILDCARD_TOKEN_TYPE
	} else if t.LPAREN != nil {
		return LPAREN_TOKEN_TYPE
	} else if t.RPAREN != nil {
		return RPAREN_TOKEN_TYPE
	} else if t.LBRACK != nil {
		return LBRACK_TOKEN_TYPE
	} else if t.RBRACK != nil {
		return RBRACK_TOKEN_TYPE
	} else if t.LBRACE != nil {
		return LBRACE_TOKEN_TYPE
	} else if t.RBRACE != nil {
		return RBRACE_TOKEN_TYPE
	} else if t.AND != nil {
		return AND_TOKEN_TYPE
	} else if t.SOR != nil {
		return SOR_TOKEN_TYPE
	} else if t.NOT != nil {
		return NOT_TOKEN_TYPE
	} else {
		return UNKNOWN_TOKEN_TYPE
	}
}

func init() {
	Scanner = participle.MustBuild(
		&Token{},
		participle.Lexer(Lexer),
	)
}

func GetTokenType(c string) TokenType {
	t := &Token{}
	Scanner.ParseString(c, t)
	return t.getTokenType()
}
