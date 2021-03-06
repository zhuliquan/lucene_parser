package token

import (
	"reflect"
	"testing"
)

func TestLexer(t *testing.T) {
	var err error
	if err != nil {
		panic(err)
	}

	type testCase struct {
		name  string
		input string
		want  []*Token
		typeS []TokenType
		wantS []string
	}

	var testCases = []testCase{
		{
			name:  "TestScan01",
			input: `\ \ \:7:>8908 8+9 x:>=90`,
			want: []*Token{
				{IDENT: `\ \ \:`},
				{NUMBER: `7`},
				{COLON: ":"},
				{COMPARE: ">"},
				{NUMBER: "8908"},
				{WHITESPACE: " "},
				{NUMBER: "8"},
				{PLUS: "+"},
				{NUMBER: "9"},
				{WHITESPACE: " "},
				{IDENT: "x"},
				{COLON: ":"},
				{COMPARE: ">="},
				{NUMBER: "90"},
			},
			typeS: []TokenType{
				IDENT_TOKEN_TYPE,
				NUMBER_TOKEN_TYPE,
				COLON_TOKEN_TYPE,
				COMPARE_TOKEN_TYPE,
				NUMBER_TOKEN_TYPE,
				WHITESPACE_TOKEN_TYPE,
				NUMBER_TOKEN_TYPE,
				PLUS_TOKEN_TYPE,
				NUMBER_TOKEN_TYPE,
				WHITESPACE_TOKEN_TYPE,
				IDENT_TOKEN_TYPE,
				COLON_TOKEN_TYPE,
				COMPARE_TOKEN_TYPE,
				NUMBER_TOKEN_TYPE,
			},
			wantS: []string{
				`\ \ \:`,
				`7`,
				":",
				">",
				"8908",
				" ",
				"8",
				"+",
				"9",
				" ",
				"x",
				":",
				">=",
				"90",
			},
		},
		{
			name:  "TestScan02",
			input: `now-8d x:/[\d\s]+/ y:"dasda 8\ : +"`,
			want: []*Token{
				{IDENT: "now"},
				{MINUS: "-"},
				{NUMBER: "8"},
				{IDENT: "d"},
				{WHITESPACE: " "},
				{IDENT: "x"},
				{COLON: ":"},
				{SLASH: "/"},
				{LBRACK: "["},
				{REVERSE: "\\"},
				{IDENT: "d"},
				{REVERSE: "\\"},
				{IDENT: "s"},
				{RBRACK: "]"},
				{PLUS: "+"},
				{SLASH: "/"},
				{WHITESPACE: " "},
				{IDENT: "y"},
				{COLON: ":"},
				{QUOTE: "\""},
				{IDENT: "dasda"},
				{WHITESPACE: " "},
				{NUMBER: "8"},
				{IDENT: "\\ "},
				{COLON: ":"},
				{WHITESPACE: " "},
				{PLUS: "+"},
				{QUOTE: "\""},
			},
			typeS: []TokenType{
				IDENT_TOKEN_TYPE,
				MINUS_TOKEN_TYPE,
				NUMBER_TOKEN_TYPE,
				IDENT_TOKEN_TYPE,
				WHITESPACE_TOKEN_TYPE,
				IDENT_TOKEN_TYPE,
				COLON_TOKEN_TYPE,
				SLASH_TOKEN_TYPE,
				LBRACK_TOKEN_TYPE,
				REVERSE_TOKEN_TYPE,
				IDENT_TOKEN_TYPE,
				REVERSE_TOKEN_TYPE,
				IDENT_TOKEN_TYPE,
				RBRACK_TOKEN_TYPE,
				PLUS_TOKEN_TYPE,
				SLASH_TOKEN_TYPE,
				WHITESPACE_TOKEN_TYPE,
				IDENT_TOKEN_TYPE,
				COLON_TOKEN_TYPE,
				QUOTE_TOKEN_TYPE,
				IDENT_TOKEN_TYPE,
				WHITESPACE_TOKEN_TYPE,
				NUMBER_TOKEN_TYPE,
				IDENT_TOKEN_TYPE,
				COLON_TOKEN_TYPE,
				WHITESPACE_TOKEN_TYPE,
				PLUS_TOKEN_TYPE,
				QUOTE_TOKEN_TYPE,
			},
			wantS: []string{
				"now",
				"-",
				"8",
				"d",
				" ",
				"x",
				":",
				"/",
				"[",
				"\\",
				"d",
				"\\",
				"s",
				"]",
				"+",
				"/",
				" ",
				"y",
				":",
				"\"",
				"dasda",
				" ",
				"8",
				"\\ ",
				":",
				" ",
				"+",
				"\"",
			},
		},
		{
			name:  "TestScan03",
			input: `\!\:.\ \\:<=<(you OR !& \!\&*\** [{ you\[\]+ you?}])^090~9~ouo |!!&&`,
			want: []*Token{
				{IDENT: `\!\:`},
				{DOT: `.`},
				{IDENT: `\ \\`},
				{COLON: ":"},
				{COMPARE: "<="},
				{COMPARE: "<"},
				{LPAREN: "("},
				{IDENT: "you"},
				{WHITESPACE: " "},
				{IDENT: "OR"},
				{WHITESPACE: " "},
				{NOT: "!"},
				{AND: "&"},
				{WHITESPACE: " "},
				{IDENT: `\!\&`},
				{WILDCARD: "*"},
				{IDENT: `\*`},
				{WILDCARD: "*"},
				{WHITESPACE: " "},
				{LBRACK: "["},
				{LBRACE: "{"},
				{WHITESPACE: " "},
				{IDENT: `you\[\]`},
				{PLUS: `+`},
				{WHITESPACE: " "},
				{IDENT: "you"},
				{WILDCARD: "?"},
				{RBRACE: "}"},
				{RBRACK: "]"},
				{RPAREN: ")"},
				{BOOST: `^`},
				{NUMBER: `090`},
				{FUZZY: `~`},
				{NUMBER: `9`},
				{FUZZY: `~`},
				{IDENT: "ouo"},
				{WHITESPACE: " "},
				{SOR: "|"},
				{NOT: "!"},
				{NOT: "!"},
				{AND: "&"},
				{AND: "&"},
			},
			typeS: []TokenType{
				IDENT_TOKEN_TYPE,
				DOT_TOKEN_TYPE,
				IDENT_TOKEN_TYPE,
				COLON_TOKEN_TYPE,
				COMPARE_TOKEN_TYPE,
				COMPARE_TOKEN_TYPE,
				LPAREN_TOKEN_TYPE,
				IDENT_TOKEN_TYPE,
				WHITESPACE_TOKEN_TYPE,
				IDENT_TOKEN_TYPE,
				WHITESPACE_TOKEN_TYPE,
				NOT_TOKEN_TYPE,
				AND_TOKEN_TYPE,
				WHITESPACE_TOKEN_TYPE,
				IDENT_TOKEN_TYPE,
				WILDCARD_TOKEN_TYPE,
				IDENT_TOKEN_TYPE,
				WILDCARD_TOKEN_TYPE,
				WHITESPACE_TOKEN_TYPE,
				LBRACK_TOKEN_TYPE,
				LBRACE_TOKEN_TYPE,
				WHITESPACE_TOKEN_TYPE,
				IDENT_TOKEN_TYPE,
				PLUS_TOKEN_TYPE,
				WHITESPACE_TOKEN_TYPE,
				IDENT_TOKEN_TYPE,
				WILDCARD_TOKEN_TYPE,
				RBRACE_TOKEN_TYPE,
				RBRACK_TOKEN_TYPE,
				RPAREN_TOKEN_TYPE,
				BOOST_TOKEN_TYPE,
				NUMBER_TOKEN_TYPE,
				FUZZY_TOKEN_TYPE,
				NUMBER_TOKEN_TYPE,
				FUZZY_TOKEN_TYPE,
				IDENT_TOKEN_TYPE,
				WHITESPACE_TOKEN_TYPE,
				SOR_TOKEN_TYPE,
				NOT_TOKEN_TYPE,
				NOT_TOKEN_TYPE,
				AND_TOKEN_TYPE,
				AND_TOKEN_TYPE,
			},
			wantS: []string{
				`\!\:`,
				`.`,
				`\ \\`,
				":",
				"<=",
				"<",
				"(",
				"you",
				" ",
				"OR",
				" ",
				"!",
				"&",
				" ",
				`\!\&`,
				"*",
				`\*`,
				"*",
				" ",
				"[",
				"{",
				" ",
				`you\[\]`,
				`+`,
				" ",
				"you",
				"?",
				"}",
				"]",
				")",
				`^`,
				`090`,
				`~`,
				`9`,
				`~`,
				"ouo",
				" ",
				"|",
				"!",
				"!",
				"&",
				"&",
			},
		},
		{
			name:  "TestScan04",
			input: `x:2021-09/d y:/89\/\d+\d*/`,
			want: []*Token{
				{IDENT: `x`},
				{COLON: ":"},
				{NUMBER: "2021"},
				{MINUS: "-"},
				{NUMBER: "09"},
				{SLASH: "/"},
				{IDENT: "d"},
				{WHITESPACE: " "},
				{IDENT: "y"},
				{COLON: ":"},
				{SLASH: "/"},
				{NUMBER: "89"},
				{IDENT: "\\/"},
				{REVERSE: `\`},
				{IDENT: "d"},
				{PLUS: `+`},
				{REVERSE: "\\"},
				{IDENT: "d"},
				{WILDCARD: "*"},
				{SLASH: "/"},
			},
			typeS: []TokenType{
				IDENT_TOKEN_TYPE,
				COLON_TOKEN_TYPE,
				NUMBER_TOKEN_TYPE,
				MINUS_TOKEN_TYPE,
				NUMBER_TOKEN_TYPE,
				SLASH_TOKEN_TYPE,
				IDENT_TOKEN_TYPE,
				WHITESPACE_TOKEN_TYPE,
				IDENT_TOKEN_TYPE,
				COLON_TOKEN_TYPE,
				SLASH_TOKEN_TYPE,
				NUMBER_TOKEN_TYPE,
				IDENT_TOKEN_TYPE,
				REVERSE_TOKEN_TYPE,
				IDENT_TOKEN_TYPE,
				PLUS_TOKEN_TYPE,
				REVERSE_TOKEN_TYPE,
				IDENT_TOKEN_TYPE,
				WILDCARD_TOKEN_TYPE,
				SLASH_TOKEN_TYPE,
			},
			wantS: []string{
				`x`,
				":",
				"2021",
				"-",
				"09",
				"/",
				"d",
				" ",
				"y",
				":",
				"/",
				"89",
				"\\/",
				`\`,
				"d",
				`+`,
				"\\",
				"d",
				"*",
				"/",
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if out := Scan(tt.input); !reflect.DeepEqual(out, tt.want) {
				t.Errorf("Scan ( %+v ) = %+v, but want: %+v", tt.input, out, tt.want)
			} else {
				for i := 0; i < len(out); i++ {
					if out[i].GetTokenType() != tt.typeS[i] {
						t.Errorf("expect get type: %+v, but get type: %+v", tt.typeS[i], out[i].GetTokenType())
					}
				}
				for i := 0; i < len(out); i++ {
					if out[i].String() != tt.wantS[i] {
						t.Errorf("expect get string: %+v, but get string: %+v", tt.wantS[i], out[i].String())
					}
				}
			}
		})
	}

	var x *Token
	if x.String() != "" {
		t.Error("expect empty")
	}
	if x.GetTokenType() != UNKNOWN_TOKEN_TYPE {
		t.Error("expect unknown type")
	}
	x = &Token{}
	if x.String() != "" {
		t.Error("expect empty")
	}
	if x.GetTokenType() != UNKNOWN_TOKEN_TYPE {
		t.Error("expect unknown type")
	}
	x = &Token{EOL: "\n"}
	if x.String() != "\n" {
		t.Errorf("expect \n")
	}
	if x.GetTokenType() != EOL_TOKEN_TYPE {
		t.Error("expect eol token")
	}
}
