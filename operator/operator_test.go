package operator

import (
	"testing"

	"github.com/alecthomas/participle"
	"github.com/stretchr/testify/assert"
	"github.com/zhuliquan/lucene_parser/token"
)

func TestAndSymbol(t *testing.T) {
	var operatorParser = participle.MustBuild(
		&AndSymbol{},
		participle.Lexer(token.Lexer),
	)
	type testCase struct {
		name    string
		input   string
		want    *AndSymbol
		wantStr string
	}
	var testCases = []testCase{
		{
			name:    "test_AND_space_symbol",
			input:   ` AND   `,
			want:    &AndSymbol{Symbol: "AND"},
			wantStr: " AND ",
		},
		{
			name:    "Test_and_space_symbol",
			input:   ` and `,
			want:    &AndSymbol{Symbol: "and"},
			wantStr: " AND ",
		},
		{
			name:    "Test_&&_space_symbol",
			input:   ` && `,
			want:    &AndSymbol{Symbol: "&&"},
			wantStr: " AND ",
		},
		{
			name:    "Test_&&_symbol",
			input:   `&&`,
			want:    &AndSymbol{Symbol: "&&"},
			wantStr: " AND ",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var symbol = &AndSymbol{}
			err := operatorParser.ParseString(tt.input, symbol)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, symbol)
			assert.Equal(t, tt.wantStr, symbol.String())
			assert.Equal(t, AND_LOGIC_TYPE, symbol.GetLogicType())
		})
	}

	var o *AndSymbol
	assert.Empty(t, o.String())
	assert.Equal(t, UNKNOWN_LOGIC_TYPE, o.GetLogicType())
	o = &AndSymbol{Symbol: ""}
	assert.Empty(t, o.String())
}

func TestOrSymbol(t *testing.T) {
	var operatorParser = participle.MustBuild(
		&OrSymbol{},
		participle.Lexer(token.Lexer),
	)

	type testCase struct {
		name    string
		input   string
		want    *OrSymbol
		wantStr string
	}
	var testCases = []testCase{
		{
			name:    "test_OR_space_symbol",
			input:   ` OR  `,
			want:    &OrSymbol{Symbol: "OR"},
			wantStr: " OR ",
		},
		{
			name:    "test_or_space_symbol",
			input:   ` or  `,
			want:    &OrSymbol{Symbol: "or"},
			wantStr: " OR ",
		},
		{
			name:    "test_||_space_symbol",
			input:   ` ||  `,
			want:    &OrSymbol{Symbol: "||"},
			wantStr: " OR ",
		},
		{
			name:    "test_||_symbol",
			input:   `||`,
			want:    &OrSymbol{Symbol: "||"},
			wantStr: " OR ",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var symbol = &OrSymbol{}
			err := operatorParser.ParseString(tt.input, symbol)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, symbol)
			assert.Equal(t, tt.wantStr, symbol.String())
			assert.Equal(t, OR_LOGIC_TYPE, symbol.GetLogicType())
		})
	}

	var o *OrSymbol
	assert.Empty(t, o.String())
	assert.Equal(t, UNKNOWN_LOGIC_TYPE, o.GetLogicType())
	o = &OrSymbol{Symbol: ""}
	assert.Empty(t, o.String())
}

func TestNotSymbol(t *testing.T) {
	var operatorParser = participle.MustBuild(
		&NotSymbol{},
		participle.Lexer(token.Lexer),
	)
	type testCase struct {
		name    string
		input   string
		want    *NotSymbol
		wantStr string
	}
	var testCases = []testCase{
		{
			name:    "test_NOT_SPACE_symbol",
			input:   `NOT `,
			want:    &NotSymbol{Symbol: "NOT"},
			wantStr: "NOT ",
		},
		{
			name:    "test_not_SPACE_symbol",
			input:   `not `,
			want:    &NotSymbol{Symbol: "not"},
			wantStr: "NOT ",
		},
		{
			name:    "test_!_SPACE_symbol",
			input:   `! `,
			want:    &NotSymbol{Symbol: "!"},
			wantStr: "NOT ",
		},
		{
			name:    "test_!_symbol",
			input:   `!`,
			want:    &NotSymbol{Symbol: "!"},
			wantStr: "NOT ",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var symbol = &NotSymbol{}
			err := operatorParser.ParseString(tt.input, symbol)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, symbol)
			assert.Equal(t, tt.wantStr, symbol.String())
			assert.Equal(t, NOT_LOGIC_TYPE, symbol.GetLogicType())
		})
	}

	var o *NotSymbol
	assert.Empty(t, o.String())
	assert.Equal(t, UNKNOWN_LOGIC_TYPE, o.GetLogicType())
	o = &NotSymbol{Symbol: ""}
	assert.Empty(t, o.String())
}

func TestPrefixOperator(t *testing.T) {

}
