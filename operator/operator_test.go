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
		name    string
		input   string
		want    *AndSymbol
		wantStr string
	}
	var testCases = []testCase{
		{
			name:    "TestAndSymbol01",
			input:   ` AND   `,
			want:    &AndSymbol{Symbol: "AND"},
			wantStr: " AND ",
		},
		{
			name:    "TestAndSymbol02",
			input:   ` and `,
			want:    &AndSymbol{Symbol: "and"},
			wantStr: " AND ",
		},
		{
			name:    "TestAndSymbol03",
			input:   ` && `,
			want:    &AndSymbol{Symbol: "&&"},
			wantStr: " AND ",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var symbol = &AndSymbol{}
			if err := operatorParser.ParseString(tt.input, symbol); err != nil {
				t.Errorf("failed to parse input: %s, err: %+v", tt.input, err)
			} else if !reflect.DeepEqual(symbol, tt.want) {
				t.Errorf("ParseString( %s ) = %+v, want: %+v", tt.input, symbol, tt.want)
			} else if symbol.String() != tt.wantStr {
				t.Errorf("expect %s, but %s", symbol.String(), tt.wantStr)
			} else if symbol.GetLogicType() != AND_LOGIC_TYPE {
				t.Errorf("expect got AND LOGIC TYPE")
			}
		})
	}

	var o *AndSymbol
	if o.String() != "" {
		t.Error("expect empty")
	}
	if o.GetLogicType() != UNKNOWN_LOGIC_TYPE {
		t.Error("expect UNKNOWN")
	}
	o = &AndSymbol{Symbol: ""}
	if o.String() != "" {
		t.Error("expect empty")
	}
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
			name:    "TestOrSymbol01",
			input:   ` OR  `,
			want:    &OrSymbol{Symbol: "OR"},
			wantStr: " OR ",
		},
		{
			name:    "TestOrSymbol02",
			input:   ` or  `,
			want:    &OrSymbol{Symbol: "or"},
			wantStr: " OR ",
		},
		{
			name:    "TestOrSymbol03",
			input:   ` ||  `,
			want:    &OrSymbol{Symbol: "||"},
			wantStr: " OR ",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var symbol = &OrSymbol{}
			if err := operatorParser.ParseString(tt.input, symbol); err != nil {
				t.Errorf("failed to parse input: %s, err: %+v", tt.input, err)
			} else if !reflect.DeepEqual(symbol, tt.want) {
				t.Errorf("ParseString( %s ) = %+v, want: %+v", tt.input, symbol, tt.want)
			} else if symbol.String() != tt.wantStr {
				t.Errorf("expect %s, but %s", symbol.String(), tt.wantStr)
			} else if symbol.GetLogicType() != OR_LOGIC_TYPE {
				t.Errorf("expect got OR LOGIC TYPE")
			}
		})
	}

	var o *OrSymbol
	if o.String() != "" {
		t.Error("expect empty")
	}
	if o.GetLogicType() != UNKNOWN_LOGIC_TYPE {
		t.Error("expect UNKNOWN")
	}
	o = &OrSymbol{Symbol: ""}
	if o.String() != "" {
		t.Error("expect empty")
	}
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
			name:    "TestNotSymbol01",
			input:   `NOT `,
			want:    &NotSymbol{Symbol: "NOT"},
			wantStr: "NOT ",
		},
		{
			name:    "TestNotSymbol02",
			input:   `not `,
			want:    &NotSymbol{Symbol: "not"},
			wantStr: "NOT ",
		},
		{
			name:    "TestNotSymbol03",
			input:   `! `,
			want:    &NotSymbol{Symbol: "!"},
			wantStr: "NOT ",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var symbol = &NotSymbol{}
			if err := operatorParser.ParseString(tt.input, symbol); err != nil {
				t.Errorf("failed to parse input: %s, err: %+v", tt.input, err)
			} else if !reflect.DeepEqual(symbol, tt.want) {
				t.Errorf("ParseString( %s ) = %+v, want: %+v", tt.input, symbol, tt.want)
			} else if symbol.String() != tt.wantStr {
				t.Errorf("expect %s, but %s", symbol.String(), tt.wantStr)
			} else if symbol.GetLogicType() != NOT_LOGIC_TYPE {
				t.Errorf("expect NOT LOGIC TYPE")
			}
		})
	}

	var o *NotSymbol
	if o.String() != "" {
		t.Error("expect empty")
	}
	if o.GetLogicType() != UNKNOWN_LOGIC_TYPE {
		t.Error("expect UNKNOWN")
	}
	o = &NotSymbol{Symbol: ""}
	if o.String() != "" {
		t.Error("expect empty")
	}
}

func TestPrefixOperator(t *testing.T) {
	

}
