package term

import (
	"math"
	"reflect"
	"testing"

	"github.com/alecthomas/participle"
	"github.com/zhuliquan/lucene_parser/token"
)

func TestRangeTerm(t *testing.T) {
	var rangesTermParser = participle.MustBuild(
		&RangeTerm{},
		participle.Lexer(token.Lexer),
	)

	type testCase struct {
		name  string
		input string
		want  *RangeTerm
		boost float64
		bound *Bound
	}
	var testCases = []testCase{
		{
			name:  "TestRangeTerm01",
			input: `<="dsada 78"`,
			want:  &RangeTerm{SRangeTerm: &SRangeTerm{Symbol: "<=", Value: &RangeValue{PhraseValue: []string{`dsada`, ` `, `78`}}}},
			boost: 1.0,
			bound: &Bound{LeftValue: &RangeValue{InfinityVal: "*"}, RightValue: &RangeValue{PhraseValue: []string{`dsada`, ` `, `78`}}, LeftInclude: false, RightInclude: true},
		},
		{
			name:  "TestRangeTerm02",
			input: `<="dsada 78"^8.9`,
			want:  &RangeTerm{SRangeTerm: &SRangeTerm{Symbol: "<=", Value: &RangeValue{PhraseValue: []string{`dsada`, ` `, `78`}}}, BoostSymbol: "^8.9"},
			boost: 8.9,
			bound: &Bound{LeftValue: &RangeValue{InfinityVal: "*"}, RightValue: &RangeValue{PhraseValue: []string{`dsada`, ` `, `78`}}, LeftInclude: false, RightInclude: true},
		},
		{
			name:  "TestRangeTerm03",
			input: `<=dsada\ 78`,
			want:  &RangeTerm{SRangeTerm: &SRangeTerm{Symbol: "<=", Value: &RangeValue{SingleValue: []string{`dsada\ `, `78`}}}},
			boost: 1.0,
			bound: &Bound{LeftValue: &RangeValue{InfinityVal: "*"}, RightValue: &RangeValue{SingleValue: []string{`dsada\ `, `78`}}, LeftInclude: false, RightInclude: true},
		},
		{
			name:  "TestRangeTerm04",
			input: `<=dsada\ 78^0.5`,
			want:  &RangeTerm{SRangeTerm: &SRangeTerm{Symbol: "<=", Value: &RangeValue{SingleValue: []string{`dsada\ `, `78`}}}, BoostSymbol: "^0.5"},
			boost: 0.5,
			bound: &Bound{LeftValue: &RangeValue{InfinityVal: "*"}, RightValue: &RangeValue{SingleValue: []string{`dsada\ `, `78`}}, LeftInclude: false, RightInclude: true},
		},
		{
			name:  "TestRangeTerm05",
			input: `[1 TO 2]`,
			want: &RangeTerm{DRangeTerm: &DRangeTerm{
				LBRACKET: "[",
				LValue:   &RangeValue{SingleValue: []string{"1"}},
				RValue:   &RangeValue{SingleValue: []string{"2"}},
				RBRACKET: "]",
			}},
			boost: 1.0,
			bound: &Bound{LeftValue: &RangeValue{SingleValue: []string{`1`}}, RightValue: &RangeValue{SingleValue: []string{"2"}}, LeftInclude: true, RightInclude: true},
		},
		{
			name:  "TestRangeTerm06",
			input: `[1 TO 2]^0.7`,
			want: &RangeTerm{DRangeTerm: &DRangeTerm{
				LBRACKET: "[",
				LValue:   &RangeValue{SingleValue: []string{"1"}},
				RValue:   &RangeValue{SingleValue: []string{"2"}},
				RBRACKET: "]",
			}, BoostSymbol: "^0.7"},
			boost: 0.7,
			bound: &Bound{LeftValue: &RangeValue{SingleValue: []string{`1`}}, RightValue: &RangeValue{SingleValue: []string{"2"}}, LeftInclude: true, RightInclude: true},
		},
		{
			name:  "TestRangeTerm07",
			input: `[1 TO 2 }`,
			want: &RangeTerm{DRangeTerm: &DRangeTerm{
				LBRACKET: "[",
				LValue:   &RangeValue{SingleValue: []string{"1"}},
				RValue:   &RangeValue{SingleValue: []string{"2"}},
				RBRACKET: "}",
			}},
			boost: 1.0,
			bound: &Bound{LeftValue: &RangeValue{SingleValue: []string{`1`}}, RightValue: &RangeValue{SingleValue: []string{"2"}}, LeftInclude: true, RightInclude: false},
		},
		{
			name:  "TestRangeTerm08",
			input: `[1 TO 2 }^0.9`,
			want: &RangeTerm{DRangeTerm: &DRangeTerm{
				LBRACKET: "[",
				LValue:   &RangeValue{SingleValue: []string{"1"}},
				RValue:   &RangeValue{SingleValue: []string{"2"}},
				RBRACKET: "}",
			}, BoostSymbol: "^0.9"},
			boost: 0.9,
			bound: &Bound{LeftValue: &RangeValue{SingleValue: []string{`1`}}, RightValue: &RangeValue{SingleValue: []string{"2"}}, LeftInclude: true, RightInclude: false},
		},
		{
			name:  `TestRangeTerm09`,
			input: `{ 1 TO 2}^7`,
			want: &RangeTerm{DRangeTerm: &DRangeTerm{
				LBRACKET: "{",
				LValue:   &RangeValue{SingleValue: []string{"1"}},
				RValue:   &RangeValue{SingleValue: []string{"2"}},
				RBRACKET: "}",
			}, BoostSymbol: "^7"},
			boost: 7.0,
			bound: &Bound{LeftValue: &RangeValue{SingleValue: []string{`1`}}, RightValue: &RangeValue{SingleValue: []string{"2"}}, LeftInclude: false, RightInclude: false},
		},
		{
			name:  `TestRangeTerm10`,
			input: `{ 1 TO 2]`,
			want: &RangeTerm{DRangeTerm: &DRangeTerm{
				LBRACKET: "{",
				LValue:   &RangeValue{SingleValue: []string{"1"}},
				RValue:   &RangeValue{SingleValue: []string{"2"}},
				RBRACKET: "]",
			}},
			boost: 1.0,
			bound: &Bound{LeftValue: &RangeValue{SingleValue: []string{`1`}}, RightValue: &RangeValue{SingleValue: []string{"2"}}, LeftInclude: false, RightInclude: true},
		},
		{
			name:  `TestRangeTerm11`,
			input: `[10 TO *]`,
			want: &RangeTerm{DRangeTerm: &DRangeTerm{
				LBRACKET: "[",
				LValue:   &RangeValue{SingleValue: []string{"10"}},
				RValue:   &RangeValue{InfinityVal: "*"},
				RBRACKET: "]",
			}},
			boost: 1.0,
			bound: &Bound{LeftValue: &RangeValue{SingleValue: []string{`10`}}, RightValue: &RangeValue{InfinityVal: "*"}, LeftInclude: true, RightInclude: false},
		},
		{
			name:  `TestRangeTerm12`,
			input: `{* TO 2012-01-01}`,
			want: &RangeTerm{DRangeTerm: &DRangeTerm{
				LBRACKET: "{",
				LValue:   &RangeValue{InfinityVal: "*"},
				RValue:   &RangeValue{SingleValue: []string{"2012", "-", "01", "-", "01"}},
				RBRACKET: "}",
			}},
			boost: 1.0,
			bound: &Bound{LeftValue: &RangeValue{InfinityVal: "*"}, RightValue: &RangeValue{SingleValue: []string{"2012", "-", "01", "-", "01"}}, LeftInclude: false, RightInclude: false},
		},
		{
			name:  `TestRangeTerm13`,
			input: `[* TO "2012-01-01 09:08:16"}`,
			want: &RangeTerm{DRangeTerm: &DRangeTerm{
				LBRACKET: "[",
				LValue:   &RangeValue{InfinityVal: "*"},
				RValue:   &RangeValue{PhraseValue: []string{"2012", "-", "01", "-", "01", " ", "09", ":", "08", ":", "16"}},
				RBRACKET: "}",
			}},
			boost: 1.0,
			bound: &Bound{LeftValue: &RangeValue{InfinityVal: "*"}, RightValue: &RangeValue{PhraseValue: []string{"2012", "-", "01", "-", "01", " ", "09", ":", "08", ":", "16"}}, LeftInclude: false, RightInclude: false},
		},
		{
			name:  `TestRangeTerm14`,
			input: `>2012-01-01^9.8`,
			want: &RangeTerm{SRangeTerm: &SRangeTerm{
				Symbol: ">",
				Value:  &RangeValue{SingleValue: []string{"2012", "-", "01", "-", "01"}},
			}, BoostSymbol: "^9.8"},
			boost: 9.8,
			bound: &Bound{LeftValue: &RangeValue{SingleValue: []string{"2012", "-", "01", "-", "01"}}, RightValue: &RangeValue{InfinityVal: "*"}, LeftInclude: false, RightInclude: false},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var out = &RangeTerm{}
			if err := rangesTermParser.ParseString(tt.input, out); err != nil {
				t.Errorf("failed to parse input: %s, err: %+v", tt.input, err)
			} else if !reflect.DeepEqual(tt.want, out) {
				t.Errorf("rangesTermParser.ParseString( %s ) = %+v, want: %+v", tt.input, out, tt.want)
			} else if math.Abs(tt.boost-out.Boost()) > 1E-6 {
				t.Errorf("expect get boost: %f, but get boost: %f", tt.boost, out.Boost())
			} else if !reflect.DeepEqual(out.GetBound(), tt.bound) {
				t.Errorf("expect get bound: %+v, but get bound: %+v", tt.bound, out.GetBound())
			}
		})
	}
}

func TestFuzzyTerm(t *testing.T) {
	var termParser = participle.MustBuild(
		&FuzzyTerm{},
		participle.Lexer(token.Lexer),
	)

	type testCase struct {
		name     string
		input    string
		want     *FuzzyTerm
		valueS   string
		wildcard bool
		boost    float64
		fuzzy    int
	}
	var testCases = []testCase{
		{
			name:     "TestFuzzyTerm01",
			input:    `"dsada\* 78"`,
			want:     &FuzzyTerm{PhraseTerm: &PhraseTerm{Value: []string{`dsada\*`, ` `, `78`}}},
			valueS:   `dsada\* 78`,
			wildcard: false,
			fuzzy:    0,
			boost:    1.0,
		},
		{
			name:     "TestFuzzyTerm02",
			input:    `"dsada* 78"`,
			want:     &FuzzyTerm{PhraseTerm: &PhraseTerm{Value: []string{`dsada`, `*`, ` `, `78`}}},
			valueS:   `dsada* 78`,
			wildcard: true,
			fuzzy:    0,
			boost:    1.0,
		},
		{
			name:     "TestFuzzyTerm03",
			input:    `"dsada\* 78"^08`,
			want:     &FuzzyTerm{PhraseTerm: &PhraseTerm{Value: []string{`dsada\*`, ` `, `78`}}, BoostSymbol: "^08"},
			valueS:   `dsada\* 78`,
			wildcard: false,
			fuzzy:    0,
			boost:    8.0,
		},
		{
			name:     "TestFuzzyTerm04",
			input:    `"dsada* 78"^08`,
			want:     &FuzzyTerm{PhraseTerm: &PhraseTerm{Value: []string{`dsada`, `*`, ` `, `78`}}, BoostSymbol: "^08"},
			valueS:   `dsada* 78`,
			wildcard: true,
			fuzzy:    0,
			boost:    8.0,
		},
		{
			name:     "TestFuzzyTerm05",
			input:    `"dsada\* 78"~8`,
			want:     &FuzzyTerm{PhraseTerm: &PhraseTerm{Value: []string{`dsada\*`, ` `, `78`}}, FuzzySymbol: "~8"},
			valueS:   `dsada\* 78`,
			wildcard: false,
			fuzzy:    8,
			boost:    1.0,
		},
		{
			name:     "TestFuzzyTerm06",
			input:    `"dsada* 78"~8`,
			want:     &FuzzyTerm{PhraseTerm: &PhraseTerm{Value: []string{`dsada`, `*`, ` `, `78`}}, FuzzySymbol: "~8"},
			valueS:   `dsada* 78`,
			wildcard: true,
			fuzzy:    8,
			boost:    1.0,
		},
		{
			name:     "TestFuzzyTerm07",
			input:    `"dsada 78"~`,
			want:     &FuzzyTerm{PhraseTerm: &PhraseTerm{Value: []string{`dsada`, ` `, `78`}}, FuzzySymbol: "~"},
			valueS:   `dsada 78`,
			wildcard: false,
			fuzzy:    1,
			boost:    1.0,
		},
		{
			name:     "TestFuzzyTerm08",
			input:    `"dsada* 78"~`,
			want:     &FuzzyTerm{PhraseTerm: &PhraseTerm{Value: []string{`dsada`, `*`, ` `, `78`}}, FuzzySymbol: "~"},
			valueS:   `dsada* 78`,
			wildcard: true,
			fuzzy:    1,
			boost:    1.0,
		},
		{
			name:     "TestFuzzyTerm09",
			input:    `\/dsada\/\ dasda80980?*`,
			want:     &FuzzyTerm{SingleTerm: &SingleTerm{Value: []string{`\/dsada\/\ dasda`, `80980`, `?`, `*`}}},
			valueS:   `\/dsada\/\ dasda80980?*`,
			wildcard: true,
			fuzzy:    0,
			boost:    1.0,
		},
		{
			name:     "TestFuzzyTerm10",
			input:    `\/dsada\/\ dasda80980?*\^\^^08`,
			want:     &FuzzyTerm{SingleTerm: &SingleTerm{Value: []string{`\/dsada\/\ dasda`, `80980`, `?`, `*`, `\^\^`}}, BoostSymbol: `^08`},
			valueS:   `\/dsada\/\ dasda80980?*\^\^`,
			wildcard: true,
			fuzzy:    0,
			boost:    8.0,
		},
		{
			name:     "TestFuzzyTerm11",
			input:    `\/dsada\/\ dasda80980?*\^\^~8`,
			want:     &FuzzyTerm{SingleTerm: &SingleTerm{Value: []string{`\/dsada\/\ dasda`, `80980`, `?`, `*`, `\^\^`}}, FuzzySymbol: `~8`},
			valueS:   `\/dsada\/\ dasda80980?*\^\^`,
			wildcard: true,
			fuzzy:    8,
			boost:    1.0,
		},
		{
			name:     "TestFuzzyTerm12",
			input:    `\/dsada\/\ dasda80980?*\^\^~`,
			want:     &FuzzyTerm{SingleTerm: &SingleTerm{Value: []string{`\/dsada\/\ dasda`, `80980`, `?`, `*`, `\^\^`}}, FuzzySymbol: `~`},
			valueS:   `\/dsada\/\ dasda80980?*\^\^`,
			wildcard: true,
			fuzzy:    1,
			boost:    1.0,
		},
		{
			name:     "TestFuzzyTerm13",
			input:    `\/dsada\/\ dasda80980\?\*`,
			want:     &FuzzyTerm{SingleTerm: &SingleTerm{Value: []string{`\/dsada\/\ dasda`, `80980`, `\?\*`}}},
			valueS:   `\/dsada\/\ dasda80980\?\*`,
			wildcard: false,
			fuzzy:    0,
			boost:    1.0,
		},
		{
			name:     "TestFuzzyTerm14",
			input:    `\/dsada\/\ dasda80980\?\*\^\^^08`,
			want:     &FuzzyTerm{SingleTerm: &SingleTerm{Value: []string{`\/dsada\/\ dasda`, `80980`, `\?\*\^\^`}}, BoostSymbol: `^08`},
			valueS:   `\/dsada\/\ dasda80980\?\*\^\^`,
			wildcard: false,
			fuzzy:    0,
			boost:    8.0,
		},
		{
			name:     "TestFuzzyTerm15",
			input:    `\/dsada\/\ dasda80980\?\*\^\^~8`,
			want:     &FuzzyTerm{SingleTerm: &SingleTerm{Value: []string{`\/dsada\/\ dasda`, `80980`, `\?\*\^\^`}}, FuzzySymbol: `~8`},
			valueS:   `\/dsada\/\ dasda80980\?\*\^\^`,
			wildcard: false,
			fuzzy:    8,
			boost:    1.0,
		},
		{
			name:     "TestFuzzyTerm16",
			input:    `\/dsada\/\ dasda80980\?\*\^\^~`,
			want:     &FuzzyTerm{SingleTerm: &SingleTerm{Value: []string{`\/dsada\/\ dasda`, `80980`, `\?\*\^\^`}}, FuzzySymbol: `~`},
			valueS:   `\/dsada\/\ dasda80980\?\*\^\^`,
			wildcard: false,
			fuzzy:    1,
			boost:    1.0,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var out = &FuzzyTerm{}
			if err := termParser.ParseString(tt.input, out); err != nil {
				t.Errorf("failed to parse input: %s, err: %+v", tt.input, err)
			} else if !reflect.DeepEqual(tt.want, out) {
				t.Errorf("fuzzyTermParser.ParseString( %s ) = %+v, want: %+v", tt.input, out, tt.want)
			} else if out.ValueS() != tt.valueS {
				t.Errorf("expect get values: %s but get values: %s", tt.valueS, out.ValueS())
			} else if out.haveWildcard() != tt.wildcard {
				t.Errorf("expect wildcard: %+v, but wildcard: %+v", tt.wildcard, out.haveWildcard())
			} else if out.Fuzziness() != tt.fuzzy {
				t.Errorf("expect get fuzzy: %d, but get fuzzy: %d", tt.fuzzy, out.Fuzziness())
			} else if math.Abs(out.Boost()-tt.boost) > 1E-6 {
				t.Errorf("expect get boost: %f, but get boost: %f", tt.boost, out.Boost())
			}
		})
	}
}
