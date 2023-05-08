package term

import (
	"testing"

	"github.com/alecthomas/participle"
	"github.com/stretchr/testify/assert"
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
			bound: &Bound{LeftValue: &RangeValue{InfinityVal: "*", flag: false}, RightValue: &RangeValue{PhraseValue: []string{`dsada`, ` `, `78`}, flag: true}, LeftInclude: false, RightInclude: true},
		},
		{
			name:  "TestRangeTerm02",
			input: `<="dsada 78"^8.9`,
			want:  &RangeTerm{SRangeTerm: &SRangeTerm{Symbol: "<=", Value: &RangeValue{PhraseValue: []string{`dsada`, ` `, `78`}}}, BoostSymbol: "^8.9"},
			boost: 8.9,
			bound: &Bound{LeftValue: &RangeValue{InfinityVal: "*", flag: false}, RightValue: &RangeValue{PhraseValue: []string{`dsada`, ` `, `78`}, flag: true}, LeftInclude: false, RightInclude: true},
		},
		{
			name:  "TestRangeTerm03",
			input: `<=dsada\ 78`,
			want:  &RangeTerm{SRangeTerm: &SRangeTerm{Symbol: "<=", Value: &RangeValue{SingleValue: []string{`dsada`, `\ `, `78`}}}},
			boost: 1.0,
			bound: &Bound{LeftValue: &RangeValue{InfinityVal: "*", flag: false}, RightValue: &RangeValue{SingleValue: []string{`dsada`, `\ `, `78`}, flag: true}, LeftInclude: false, RightInclude: true},
		},
		{
			name:  "TestRangeTerm04",
			input: `<=dsada\ 78^0.5`,
			want:  &RangeTerm{SRangeTerm: &SRangeTerm{Symbol: "<=", Value: &RangeValue{SingleValue: []string{`dsada`, `\ `, `78`}}}, BoostSymbol: "^0.5"},
			boost: 0.5,
			bound: &Bound{LeftValue: &RangeValue{InfinityVal: "*", flag: false}, RightValue: &RangeValue{SingleValue: []string{`dsada`, `\ `, `78`}, flag: true}, LeftInclude: false, RightInclude: true},
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
			bound: &Bound{LeftValue: &RangeValue{SingleValue: []string{`1`}, flag: false}, RightValue: &RangeValue{SingleValue: []string{"2"}, flag: true}, LeftInclude: true, RightInclude: true},
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
			bound: &Bound{LeftValue: &RangeValue{SingleValue: []string{`1`}, flag: false}, RightValue: &RangeValue{SingleValue: []string{"2"}, flag: true}, LeftInclude: true, RightInclude: true},
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
			bound: &Bound{LeftValue: &RangeValue{SingleValue: []string{`1`}, flag: false}, RightValue: &RangeValue{SingleValue: []string{"2"}, flag: true}, LeftInclude: true, RightInclude: false},
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
			bound: &Bound{LeftValue: &RangeValue{SingleValue: []string{`1`}, flag: false}, RightValue: &RangeValue{SingleValue: []string{"2"}, flag: true}, LeftInclude: true, RightInclude: false},
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
			bound: &Bound{LeftValue: &RangeValue{SingleValue: []string{`1`}, flag: false}, RightValue: &RangeValue{SingleValue: []string{"2"}, flag: true}, LeftInclude: false, RightInclude: false},
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
			bound: &Bound{LeftValue: &RangeValue{SingleValue: []string{`1`}, flag: false}, RightValue: &RangeValue{SingleValue: []string{"2"}, flag: true}, LeftInclude: false, RightInclude: true},
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
			bound: &Bound{LeftValue: &RangeValue{SingleValue: []string{`10`}, flag: false}, RightValue: &RangeValue{InfinityVal: "*", flag: true}, LeftInclude: true, RightInclude: false},
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
			bound: &Bound{LeftValue: &RangeValue{InfinityVal: "*", flag: false}, RightValue: &RangeValue{SingleValue: []string{"2012", "-", "01", "-", "01"}, flag: true}, LeftInclude: false, RightInclude: false},
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
			bound: &Bound{LeftValue: &RangeValue{InfinityVal: "*", flag: false}, RightValue: &RangeValue{PhraseValue: []string{"2012", "-", "01", "-", "01", " ", "09", ":", "08", ":", "16"}, flag: true}, LeftInclude: false, RightInclude: false},
		},
		{
			name:  `TestRangeTerm14`,
			input: `>2012-01-01^9.8`,
			want: &RangeTerm{SRangeTerm: &SRangeTerm{
				Symbol: ">",
				Value:  &RangeValue{SingleValue: []string{"2012", "-", "01", "-", "01"}},
			}, BoostSymbol: "^9.8"},
			boost: 9.8,
			bound: &Bound{LeftValue: &RangeValue{SingleValue: []string{"2012", "-", "01", "-", "01"}, flag: false}, RightValue: &RangeValue{InfinityVal: "*", flag: true}, LeftInclude: false, RightInclude: false},
		},
		{
			name:  "TestRangeTerm15",
			input: `<="dsada 78"^.9`,
			want:  &RangeTerm{SRangeTerm: &SRangeTerm{Symbol: "<=", Value: &RangeValue{PhraseValue: []string{`dsada`, ` `, `78`}}}, BoostSymbol: "^.9"},
			boost: 0.9,
			bound: &Bound{LeftValue: &RangeValue{InfinityVal: "*", flag: false}, RightValue: &RangeValue{PhraseValue: []string{`dsada`, ` `, `78`}, flag: true}, LeftInclude: false, RightInclude: true},
		},
		{
			name:  "TestRangeTerm16",
			input: `<="dsada 78"^`,
			want:  &RangeTerm{SRangeTerm: &SRangeTerm{Symbol: "<=", Value: &RangeValue{PhraseValue: []string{`dsada`, ` `, `78`}}}, BoostSymbol: "^"},
			boost: 1.0,
			bound: &Bound{LeftValue: &RangeValue{InfinityVal: "*", flag: false}, RightValue: &RangeValue{PhraseValue: []string{`dsada`, ` `, `78`}, flag: true}, LeftInclude: false, RightInclude: true},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var out = &RangeTerm{}
			err := rangesTermParser.ParseString(tt.input, out)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, out)
			assert.Equal(t, tt.boost, out.Boost())
			assert.Equal(t, tt.bound, out.GetBound())
		})
	}

	var out *RangeTerm
	assert.Empty(t, "", out.String())
	assert.Nil(t, out.GetBound())
	assert.Equal(t, 0.0, out.Boost())
	assert.Equal(t, UNKNOWN_TERM_TYPE, out.GetTermType())

	out = &RangeTerm{}
	assert.Empty(t, "", out.String())
	assert.Nil(t, out.GetBound())
	assert.Equal(t, 0.0, out.Boost())
	assert.Equal(t, UNKNOWN_TERM_TYPE, out.GetTermType())
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
			want:     &FuzzyTerm{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, `\*`, ` `, `78`}}},
			valueS:   `"dsada\* 78"`,
			wildcard: false,
			fuzzy:    0,
			boost:    1.0,
		},
		{
			name:     "TestFuzzyTerm02",
			input:    `"dsada* 78"`,
			want:     &FuzzyTerm{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, `*`, ` `, `78`}}},
			valueS:   `"dsada* 78"`,
			wildcard: true,
			fuzzy:    0,
			boost:    1.0,
		},
		{
			name:     "TestFuzzyTerm03",
			input:    `"dsada\* 78"^08`,
			want:     &FuzzyTerm{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, `\*`, ` `, `78`}}, BoostSymbol: "^08"},
			valueS:   `"dsada\* 78"^08`,
			wildcard: false,
			fuzzy:    0,
			boost:    8.0,
		},
		{
			name:     "TestFuzzyTerm04",
			input:    `"dsada* 78"^08`,
			want:     &FuzzyTerm{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, `*`, ` `, `78`}}, BoostSymbol: "^08"},
			valueS:   `"dsada* 78"^08`,
			wildcard: true,
			fuzzy:    0,
			boost:    8.0,
		},
		{
			name:     "TestFuzzyTerm05",
			input:    `"dsada\* 78"~8`,
			want:     &FuzzyTerm{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, `\*`, ` `, `78`}}, FuzzySymbol: "~8"},
			valueS:   `"dsada\* 78"~8`,
			wildcard: false,
			fuzzy:    8,
			boost:    1.0,
		},
		{
			name:     "TestFuzzyTerm06",
			input:    `"dsada* 78"~8`,
			want:     &FuzzyTerm{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, `*`, ` `, `78`}}, FuzzySymbol: "~8"},
			valueS:   `"dsada* 78"~8`,
			wildcard: true,
			fuzzy:    8,
			boost:    1.0,
		},
		{
			name:     "TestFuzzyTerm06_01",
			input:    `"dsada* 78"~8.1`,
			want:     &FuzzyTerm{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, `*`, ` `, `78`}}, FuzzySymbol: "~8.1"},
			valueS:   `"dsada* 78"~8.1`,
			wildcard: true,
			fuzzy:    8,
			boost:    1.0,
		},
		{
			name:     "TestFuzzyTerm06_02",
			input:    `"dsada* 78"~8.6`,
			want:     &FuzzyTerm{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, `*`, ` `, `78`}}, FuzzySymbol: "~8.6"},
			valueS:   `"dsada* 78"~8.6`,
			wildcard: true,
			fuzzy:    9,
			boost:    1.0,
		},
		{
			name:     "TestFuzzyTerm07",
			input:    `"dsada 78"~`,
			want:     &FuzzyTerm{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, ` `, `78`}}, FuzzySymbol: "~"},
			valueS:   `"dsada 78"~`,
			wildcard: false,
			fuzzy:    -1,
			boost:    1.0,
		},
		{
			name:     "TestFuzzyTerm08",
			input:    `"dsada* 78"~`,
			want:     &FuzzyTerm{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, `*`, ` `, `78`}}, FuzzySymbol: "~"},
			valueS:   `"dsada* 78"~`,
			wildcard: true,
			fuzzy:    -1,
			boost:    1.0,
		},
		{
			name:     "TestFuzzyTerm09",
			input:    `\/dsada\/\ dasda80980?*`,
			want:     &FuzzyTerm{SingleTerm: &SingleTerm{Begin: `\/`, Chars: []string{`dsada`, `\/\ `, `dasda`, `80980`, `?`, `*`}}},
			valueS:   `\/dsada\/\ dasda80980?*`,
			wildcard: true,
			fuzzy:    0,
			boost:    1.0,
		},
		{
			name:     "TestFuzzyTerm10",
			input:    `\/dsada\/\ dasda80980?*\^\^^08`,
			want:     &FuzzyTerm{SingleTerm: &SingleTerm{Begin: `\/`, Chars: []string{`dsada`, `\/\ `, `dasda`, `80980`, `?`, `*`, `\^\^`}}, BoostSymbol: `^08`},
			valueS:   `\/dsada\/\ dasda80980?*\^\^^08`,
			wildcard: true,
			fuzzy:    0,
			boost:    8.0,
		},
		{
			name:     "TestFuzzyTerm10_01",
			input:    `\/dsada\/\ dasda80980?*\^\^^`,
			want:     &FuzzyTerm{SingleTerm: &SingleTerm{Begin: `\/`, Chars: []string{`dsada`, `\/\ `, `dasda`, `80980`, `?`, `*`, `\^\^`}}, BoostSymbol: `^`},
			valueS:   `\/dsada\/\ dasda80980?*\^\^^`,
			wildcard: true,
			fuzzy:    0,
			boost:    1.0,
		},
		{
			name:     "TestFuzzyTerm11",
			input:    `\/dsada\/\ dasda80980?*\^\^~8`,
			want:     &FuzzyTerm{SingleTerm: &SingleTerm{Begin: `\/`, Chars: []string{`dsada`, `\/\ `, `dasda`, `80980`, `?`, `*`, `\^\^`}}, FuzzySymbol: `~8`},
			valueS:   `\/dsada\/\ dasda80980?*\^\^~8`,
			wildcard: true,
			fuzzy:    8,
			boost:    1.0,
		},
		{
			name:     "TestFuzzyTerm12",
			input:    `\/dsada\/\ dasda80980?*\^\^~`,
			want:     &FuzzyTerm{SingleTerm: &SingleTerm{Begin: `\/`, Chars: []string{`dsada`, `\/\ `, `dasda`, `80980`, `?`, `*`, `\^\^`}}, FuzzySymbol: `~`},
			valueS:   `\/dsada\/\ dasda80980?*\^\^~`,
			wildcard: true,
			fuzzy:    -1,
			boost:    1.0,
		},
		{
			name:     "TestFuzzyTerm13",
			input:    `\/dsada\/\ dasda80980\?\*`,
			want:     &FuzzyTerm{SingleTerm: &SingleTerm{Begin: `\/`, Chars: []string{`dsada`, `\/\ `, `dasda`, `80980`, `\?\*`}}},
			valueS:   `\/dsada\/\ dasda80980\?\*`,
			wildcard: false,
			fuzzy:    0,
			boost:    1.0,
		},
		{
			name:     "TestFuzzyTerm14",
			input:    `\/dsada\/\ dasda80980\?\*\^\^^08`,
			want:     &FuzzyTerm{SingleTerm: &SingleTerm{Begin: `\/`, Chars: []string{`dsada`, `\/\ `, `dasda`, `80980`, `\?\*\^\^`}}, BoostSymbol: `^08`},
			valueS:   `\/dsada\/\ dasda80980\?\*\^\^^08`,
			wildcard: false,
			fuzzy:    0,
			boost:    8.0,
		},
		{
			name:     "TestFuzzyTerm15",
			input:    `\/dsada\/\ dasda80980\?\*\^\^~8`,
			want:     &FuzzyTerm{SingleTerm: &SingleTerm{Begin: `\/`, Chars: []string{`dsada`, `\/\ `, `dasda`, `80980`, `\?\*\^\^`}}, FuzzySymbol: `~8`},
			valueS:   `\/dsada\/\ dasda80980\?\*\^\^~8`,
			wildcard: false,
			fuzzy:    8,
			boost:    1.0,
		},
		{
			name:     "TestFuzzyTerm16",
			input:    `\/dsada\/\ dasda80980\?\*\^\^~`,
			want:     &FuzzyTerm{SingleTerm: &SingleTerm{Begin: `\/`, Chars: []string{`dsada`, `\/\ `, `dasda`, `80980`, `\?\*\^\^`}}, FuzzySymbol: `~`},
			valueS:   `\/dsada\/\ dasda80980\?\*\^\^~`,
			wildcard: false,
			fuzzy:    -1,
			boost:    1.0,
		},
		{
			name:     "TestFuzzyTerm17",
			input:    `\/dsada\/\ dasda80980?*\^\^^.8`,
			want:     &FuzzyTerm{SingleTerm: &SingleTerm{Begin: `\/`, Chars: []string{`dsada`, `\/\ `, `dasda`, `80980`, `?`, `*`, `\^\^`}}, BoostSymbol: `^.8`},
			valueS:   `\/dsada\/\ dasda80980?*\^\^^.8`,
			wildcard: true,
			fuzzy:    0,
			boost:    0.8,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var out = &FuzzyTerm{}
			err := termParser.ParseString(tt.input, out)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, out)
			assert.Equal(t, tt.valueS, out.String())
			assert.Equal(t, tt.wildcard, out.haveWildcard())
			assert.Equal(t, tt.fuzzy, out.Fuzziness())
			assert.Equal(t, tt.boost, out.Boost())
		})
	}
	var out *FuzzyTerm
	assert.Empty(t, out.String())
	assert.Equal(t, 0, out.Fuzziness())
	assert.Equal(t, 0.0, out.Boost())
	assert.Equal(t, false, out.haveWildcard())
	assert.Equal(t, UNKNOWN_TERM_TYPE, out.GetTermType())
	_, err := out.Value(func(s string) (interface{}, error) { return s, nil })
	assert.Equal(t, ErrEmptyFuzzyTerm, err)

	out = &FuzzyTerm{}
	assert.Empty(t, out.String())
	assert.Equal(t, 0, out.Fuzziness())
	assert.Equal(t, 0.0, out.Boost())
	assert.Equal(t, false, out.haveWildcard())
	assert.Equal(t, UNKNOWN_TERM_TYPE, out.GetTermType())
	v, _ := out.Value(func(s string) (interface{}, error) { return s, nil })
	assert.Equal(t, "", v)
}

func TestTermGroupElem(t *testing.T) {
	type test struct {
		name    string
		input   *TermGroupElem
		tType   TermType
		valueS  interface{}
		wantErr error
	}

	for _, tt := range []test{
		{
			name:    "test_empty_case01",
			input:   nil,
			tType:   UNKNOWN_TERM_TYPE,
			valueS:  nil,
			wantErr: ErrEmptyTermGroupElem,
		},
		{
			name:    "test_empty_case02",
			input:   &TermGroupElem{},
			tType:   UNKNOWN_TERM_TYPE,
			valueS:  "",
			wantErr: nil,
		},
		{
			name: "test_single",
			input: &TermGroupElem{
				SingleTerm: &SingleTerm{Begin: "123"},
			},
			tType:   SINGLE_TERM_TYPE,
			valueS:  "123",
			wantErr: nil,
		},
		{
			name: "test_phrase",
			input: &TermGroupElem{
				PhraseTerm: &PhraseTerm{Chars: []string{"123"}},
			},
			tType:   PHRASE_TERM_TYPE,
			valueS:  "123",
			wantErr: nil,
		},
		{
			name: "test_s_range",
			input: &TermGroupElem{
				SRangeTerm: &SRangeTerm{
					Symbol: ">",
					Value:  &RangeValue{SingleValue: []string{"123"}},
				},
			},
			tType: RANGE_TERM_TYPE,
			valueS: &Bound{
				LeftValue:    &RangeValue{SingleValue: []string{"123"}, flag: false},
				RightValue:   &RangeValue{InfinityVal: "*", flag: true},
				LeftInclude:  false,
				RightInclude: false,
			},
			wantErr: nil,
		},
		{
			name: "test_d_range",
			input: &TermGroupElem{
				DRangeTerm: &DRangeTerm{
					LBRACKET: "{",
					RBRACKET: "}",
					LValue:   &RangeValue{SingleValue: []string{"123"}},
					RValue:   &RangeValue{SingleValue: []string{"456"}},
				},
			},
			tType: RANGE_TERM_TYPE,
			valueS: &Bound{
				LeftValue:    &RangeValue{SingleValue: []string{"123"}, flag: false},
				RightValue:   &RangeValue{SingleValue: []string{"456"}, flag: true},
				LeftInclude:  false,
				RightInclude: false,
			},
			wantErr: nil,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.tType, tt.input.GetTermType())
			v, e := tt.input.Value(func(s string) (interface{}, error) { return s, nil })
			assert.Equal(t, tt.wantErr, e)
			assert.Equal(t, tt.valueS, v)
		})
	}
	var out *TermGroupElem
	assert.Empty(t, out.String())
	out = &TermGroupElem{}
	assert.Empty(t, out.String())
}
