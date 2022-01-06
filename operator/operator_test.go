package operator

import (
	"reflect"
	"testing"

	"github.com/alecthomas/participle"
	"github.com/zhuliquan/lucene_parser/token"
)

func TestAndSymbol(t *testing.T) {
	var operatorParser = participle.MustBuild(
		&AndSymbol{},
		participle.Lexer(token.Lexer),
	)
	type testCase struct {
		name  string
		input string
		want  *AndSymbol
	}
	var testCases = []testCase{
		{
			name:  "TestAndSymbol01",
			input: ` AND   `,
			want:  &AndSymbol{Symbol: "AND"},
		},
		{
			name:  "TestAndSymbol02",
			input: ` and `,
			want:  &AndSymbol{Symbol: "and"},
		},
		{
			name:  "TestAndSymbol03",
			input: ` && `,
			want:  &AndSymbol{Symbol: "&&"},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var symbol = &AndSymbol{}
			if err := operatorParser.ParseString(tt.input, symbol); err != nil {
				t.Errorf("failed to parse input: %s, err: %+v", tt.input, err)
			} else if !reflect.DeepEqual(symbol, tt.want) {
				t.Errorf("ParseString( %s ) = %+v, want: %+v", tt.input, symbol, tt.want)
			}
		})
	}
}

func TestOrSymbol(t *testing.T) {
	var operatorParser = participle.MustBuild(
		&OrSymbol{},
		participle.Lexer(token.Lexer),
	)

	type testCase struct {
		name  string
		input string
		want  *OrSymbol
	}
	var testCases = []testCase{
		{
			name:  "TestOrSymbol01",
			input: ` OR  `,
			want:  &OrSymbol{Symbol: "OR"},
		},
		{
			name:  "TestOrSymbol02",
			input: ` or  `,
			want:  &OrSymbol{Symbol: "or"},
		},
		{
			name:  "TestOrSymbol03",
			input: ` ||  `,
			want:  &OrSymbol{Symbol: "||"},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var symbol = &OrSymbol{}
			if err := operatorParser.ParseString(tt.input, symbol); err != nil {
				t.Errorf("failed to parse input: %s, err: %+v", tt.input, err)
			} else if !reflect.DeepEqual(symbol, tt.want) {
				t.Errorf("ParseString( %s ) = %+v, want: %+v", tt.input, symbol, tt.want)
			}
		})
	}
}

func TestNotSymbol(t *testing.T) {
	var operatorParser = participle.MustBuild(
		&NotSymbol{},
		participle.Lexer(token.Lexer),
	)
	type testCase struct {
		name  string
		input string
		want  *NotSymbol
	}
	var testCases = []testCase{
		{
			name:  "TestNotSymbol01",
			input: `NOT `,
			want:  &NotSymbol{Symbol: "NOT"},
		},
		{
			name:  "TestNotSymbol02",
			input: `not `,
			want:  &NotSymbol{Symbol: "not"},
		},
		{
			name:  "TestNotSymbol03",
			input: `! `,
			want:  &NotSymbol{Symbol: "!"},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var symbol = &NotSymbol{}
			if err := operatorParser.ParseString(tt.input, symbol); err != nil {
				t.Errorf("failed to parse input: %s, err: %+v", tt.input, err)
			} else if !reflect.DeepEqual(symbol, tt.want) {
				t.Errorf("ParseString( %s ) = %+v, want: %+v", tt.input, symbol, tt.want)
			}
		})
	}
}
