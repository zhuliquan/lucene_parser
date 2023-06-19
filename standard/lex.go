package standard

import (
	"github.com/alecthomas/participle/lexer/stateful"
)

var rules = []stateful.Rule{
	{
		Name:    "EOF",
		Pattern: `\n`,
	},
	{
		Name:    "WHITESPACE",
		Pattern: `[\t\r\f 0x3000]+`,
	},
	{
		Name:    "IDENT",
		Pattern: `[^-!:\|&"\?\*\\\^@~\(\)\{\}\[\]\+\/><=0-9\.\t\r\f\n 　]+`,
	},
	{
		Name:    "ESCAPE",
		Pattern: `(\\.)+`, // Every character that follows a backslash is considered as an escaped character
	},
	{
		Name:    "NUMBER",
		Pattern: `[0-9]+`,
	},
	{
		Name:    "DOT",
		Pattern: `\.`,
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
		Name:    "TILDE",
		Pattern: `~`,
	},
	{
		Name:    "CARAT",
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
		Name:    "OR",
		Pattern: `\|`,
	},
	{
		Name:    "NOT",
		Pattern: `!`,
	},
}

var Lexer *stateful.Definition

func init() {
	var err error
	Lexer, err = stateful.NewSimple(rules)
	if err != nil {
		panic(err)
	}
}
