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
		boost BoostValue
		bound *Bound
	}
	var testCases = []testCase{
		{
			name:  "test_lte_phrase",
			input: `<="dsada 78"`,
			want:  &RangeTerm{SRangeTerm: &SRangeTerm{Symbol: "<=", Value: &RangeValue{PhraseValue: []string{`dsada`, ` `, `78`}}}},
			boost: DefaultBoost,
			bound: &Bound{LeftValue: &RangeValue{InfinityVal: "*", SideFlag: false}, RightValue: &RangeValue{PhraseValue: []string{`dsada`, ` `, `78`}, SideFlag: true}, LeftInclude: false, RightInclude: true},
		},
		{
			name:  "test_lte_phrase_with_boost",
			input: `<="dsada 78"^8.9`,
			want:  &RangeTerm{SRangeTerm: &SRangeTerm{Symbol: "<=", Value: &RangeValue{PhraseValue: []string{`dsada`, ` `, `78`}}}, BoostSymbol: "^8.9"},
			boost: BoostValue(8.9),
			bound: &Bound{LeftValue: &RangeValue{InfinityVal: "*", SideFlag: false}, RightValue: &RangeValue{PhraseValue: []string{`dsada`, ` `, `78`}, SideFlag: true}, LeftInclude: false, RightInclude: true},
		},
		{
			name:  "test_lte_single",
			input: `<=dsada\ 78`,
			want:  &RangeTerm{SRangeTerm: &SRangeTerm{Symbol: "<=", Value: &RangeValue{SingleValue: []string{`dsada`, `\ `, `78`}}}},
			boost: DefaultBoost,
			bound: &Bound{LeftValue: &RangeValue{InfinityVal: "*", SideFlag: false}, RightValue: &RangeValue{SingleValue: []string{`dsada`, `\ `, `78`}, SideFlag: true}, LeftInclude: false, RightInclude: true},
		},
		{
			name:  "test_lte_single_with_boost",
			input: `<=dsada\ 78^0.5`,
			want:  &RangeTerm{SRangeTerm: &SRangeTerm{Symbol: "<=", Value: &RangeValue{SingleValue: []string{`dsada`, `\ `, `78`}}}, BoostSymbol: "^0.5"},
			boost: BoostValue(0.5),
			bound: &Bound{LeftValue: &RangeValue{InfinityVal: "*", SideFlag: false}, RightValue: &RangeValue{SingleValue: []string{`dsada`, `\ `, `78`}, SideFlag: true}, LeftInclude: false, RightInclude: true},
		},
		{
			name:  "test_range_two_include",
			input: `[1 TO 2]`,
			want: &RangeTerm{DRangeTerm: &DRangeTerm{
				LBRACKET: "[",
				LValue:   &RangeValue{SingleValue: []string{"1"}},
				RValue:   &RangeValue{SingleValue: []string{"2"}},
				RBRACKET: "]",
			}},
			boost: DefaultBoost,
			bound: &Bound{LeftValue: &RangeValue{SingleValue: []string{`1`}, SideFlag: false}, RightValue: &RangeValue{SingleValue: []string{"2"}, SideFlag: true}, LeftInclude: true, RightInclude: true},
		},
		{
			name:  "test_range_with_boost",
			input: `[1 TO 2]^0.7`,
			want: &RangeTerm{DRangeTerm: &DRangeTerm{
				LBRACKET: "[",
				LValue:   &RangeValue{SingleValue: []string{"1"}},
				RValue:   &RangeValue{SingleValue: []string{"2"}},
				RBRACKET: "]",
			}, BoostSymbol: "^0.7"},
			boost: BoostValue(0.7),
			bound: &Bound{LeftValue: &RangeValue{SingleValue: []string{`1`}, SideFlag: false}, RightValue: &RangeValue{SingleValue: []string{"2"}, SideFlag: true}, LeftInclude: true, RightInclude: true},
		},
		{
			name:  "test_range_left_include_and_right_exclude",
			input: `[1 TO 2 }`,
			want: &RangeTerm{DRangeTerm: &DRangeTerm{
				LBRACKET: "[",
				LValue:   &RangeValue{SingleValue: []string{"1"}},
				RValue:   &RangeValue{SingleValue: []string{"2"}},
				RBRACKET: "}",
			}},
			boost: DefaultBoost,
			bound: &Bound{LeftValue: &RangeValue{SingleValue: []string{`1`}, SideFlag: false}, RightValue: &RangeValue{SingleValue: []string{"2"}, SideFlag: true}, LeftInclude: true, RightInclude: false},
		},
		{
			name:  "test_range_left_include_and_right_exclude_with_boost",
			input: `[1 TO 2 }^0.9`,
			want: &RangeTerm{DRangeTerm: &DRangeTerm{
				LBRACKET: "[",
				LValue:   &RangeValue{SingleValue: []string{"1"}},
				RValue:   &RangeValue{SingleValue: []string{"2"}},
				RBRACKET: "}",
			}, BoostSymbol: "^0.9"},
			boost: BoostValue(0.9),
			bound: &Bound{LeftValue: &RangeValue{SingleValue: []string{`1`}, SideFlag: false}, RightValue: &RangeValue{SingleValue: []string{"2"}, SideFlag: true}, LeftInclude: true, RightInclude: false},
		},
		{
			name:  `test_range_two_exclude_with_boost`,
			input: `{ 1 TO 2}^7`,
			want: &RangeTerm{DRangeTerm: &DRangeTerm{
				LBRACKET: "{",
				LValue:   &RangeValue{SingleValue: []string{"1"}},
				RValue:   &RangeValue{SingleValue: []string{"2"}},
				RBRACKET: "}",
			}, BoostSymbol: "^7"},
			boost: BoostValue(7.0),
			bound: &Bound{LeftValue: &RangeValue{SingleValue: []string{`1`}, SideFlag: false}, RightValue: &RangeValue{SingleValue: []string{"2"}, SideFlag: true}, LeftInclude: false, RightInclude: false},
		},
		{
			name:  `test_left_exclude_and_right_include`,
			input: `{ 1 TO 2]`,
			want: &RangeTerm{DRangeTerm: &DRangeTerm{
				LBRACKET: "{",
				LValue:   &RangeValue{SingleValue: []string{"1"}},
				RValue:   &RangeValue{SingleValue: []string{"2"}},
				RBRACKET: "]",
			}},
			boost: DefaultBoost,
			bound: &Bound{LeftValue: &RangeValue{SingleValue: []string{`1`}, SideFlag: false}, RightValue: &RangeValue{SingleValue: []string{"2"}, SideFlag: true}, LeftInclude: false, RightInclude: true},
		},
		{
			name:  `test_left_include_and_right_inf`,
			input: `[10 TO *]`,
			want: &RangeTerm{DRangeTerm: &DRangeTerm{
				LBRACKET: "[",
				LValue:   &RangeValue{SingleValue: []string{"10"}},
				RValue:   &RangeValue{InfinityVal: "*"},
				RBRACKET: "]",
			}},
			boost: DefaultBoost,
			bound: &Bound{LeftValue: &RangeValue{SingleValue: []string{`10`}, SideFlag: false}, RightValue: &RangeValue{InfinityVal: "*", SideFlag: true}, LeftInclude: true, RightInclude: false},
		},
		{
			name:  `test_left_inf_and_right_date`,
			input: `{* TO 2012-01-01}`,
			want: &RangeTerm{DRangeTerm: &DRangeTerm{
				LBRACKET: "{",
				LValue:   &RangeValue{InfinityVal: "*"},
				RValue:   &RangeValue{SingleValue: []string{"2012", "-", "01", "-", "01"}},
				RBRACKET: "}",
			}},
			boost: DefaultBoost,
			bound: &Bound{LeftValue: &RangeValue{InfinityVal: "*", SideFlag: false}, RightValue: &RangeValue{SingleValue: []string{"2012", "-", "01", "-", "01"}, SideFlag: true}, LeftInclude: false, RightInclude: false},
		},
		{
			name:  `test_left_inf_right_phrase_date`,
			input: `[* TO "2012-01-01 09:08:16"}`,
			want: &RangeTerm{DRangeTerm: &DRangeTerm{
				LBRACKET: "[",
				LValue:   &RangeValue{InfinityVal: "*"},
				RValue:   &RangeValue{PhraseValue: []string{"2012", "-", "01", "-", "01", " ", "09", ":", "08", ":", "16"}},
				RBRACKET: "}",
			}},
			boost: DefaultBoost,
			bound: &Bound{LeftValue: &RangeValue{InfinityVal: "*", SideFlag: false}, RightValue: &RangeValue{PhraseValue: []string{"2012", "-", "01", "-", "01", " ", "09", ":", "08", ":", "16"}, SideFlag: true}, LeftInclude: false, RightInclude: false},
		},
		{
			name:  `test_single_range_gt_with_boost`,
			input: `>2012-01-01^9.8`,
			want: &RangeTerm{SRangeTerm: &SRangeTerm{
				Symbol: ">",
				Value:  &RangeValue{SingleValue: []string{"2012", "-", "01", "-", "01"}},
			}, BoostSymbol: "^9.8"},
			boost: BoostValue(9.8),
			bound: &Bound{LeftValue: &RangeValue{SingleValue: []string{"2012", "-", "01", "-", "01"}, SideFlag: false}, RightValue: &RangeValue{InfinityVal: "*", SideFlag: true}, LeftInclude: false, RightInclude: false},
		},
		{
			name:  "test_single_range_lte_with_boost",
			input: `<="dsada 78"^.9`,
			want:  &RangeTerm{SRangeTerm: &SRangeTerm{Symbol: "<=", Value: &RangeValue{PhraseValue: []string{`dsada`, ` `, `78`}}}, BoostSymbol: "^.9"},
			boost: BoostValue(0.9),
			bound: &Bound{LeftValue: &RangeValue{InfinityVal: "*", SideFlag: false}, RightValue: &RangeValue{PhraseValue: []string{`dsada`, ` `, `78`}, SideFlag: true}, LeftInclude: false, RightInclude: true},
		},
		{
			name:  "test_single_range_lte_with_default_boost",
			input: `<="dsada 78"^`,
			want:  &RangeTerm{SRangeTerm: &SRangeTerm{Symbol: "<=", Value: &RangeValue{PhraseValue: []string{`dsada`, ` `, `78`}}}, BoostSymbol: "^"},
			boost: DefaultBoost,
			bound: &Bound{LeftValue: &RangeValue{InfinityVal: "*", SideFlag: false}, RightValue: &RangeValue{PhraseValue: []string{`dsada`, ` `, `78`}, SideFlag: true}, LeftInclude: false, RightInclude: true},
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
	assert.Equal(t, NoBoost, out.Boost())
	assert.Equal(t, UNKNOWN_TERM_TYPE, out.GetTermType())

	out = &RangeTerm{}
	assert.Empty(t, "", out.String())
	assert.Nil(t, out.GetBound())
	assert.Equal(t, NoBoost, out.Boost())
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
		boost    BoostValue
		fuzzy    Fuzziness
	}
	var testCases = []testCase{
		{
			name:     "test_phrase_with_space_and_escape",
			input:    `"dsada\* 78"`,
			want:     &FuzzyTerm{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, `\*`, ` `, `78`}}},
			valueS:   `"dsada\* 78"`,
			wildcard: false,
			fuzzy:    NoFuzzy,
			boost:    DefaultBoost,
		},
		{
			name:     "test_phrase_with_space_and_star_symbol",
			input:    `"dsada* 78"`,
			want:     &FuzzyTerm{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, `*`, ` `, `78`}}},
			valueS:   `"dsada* 78"`,
			wildcard: false,
			fuzzy:    NoFuzzy,
			boost:    DefaultBoost,
		},
		{
			name:     "test_phrase_with_space_and_escape_and_int_boost",
			input:    `"dsada\* 78"^08`,
			want:     &FuzzyTerm{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, `\*`, ` `, `78`}}, BoostSymbol: "^08"},
			valueS:   `"dsada\* 78"^08`,
			wildcard: false,
			fuzzy:    NoFuzzy,
			boost:    BoostValue(8.0),
		},
		{
			name:     "test_phrase_with_space_and_star_symbol_and_int_boost",
			input:    `"dsada* 78"^08`,
			want:     &FuzzyTerm{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, `*`, ` `, `78`}}, BoostSymbol: "^08"},
			valueS:   `"dsada* 78"^08`,
			wildcard: false,
			fuzzy:    NoFuzzy,
			boost:    BoostValue(8.0),
		},
		{
			name:     "test_phrase_with_space_and_escape_and_fuzzy",
			input:    `"dsada\* 78"~8`,
			want:     &FuzzyTerm{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, `\*`, ` `, `78`}}, FuzzySymbol: "~8"},
			valueS:   `"dsada\* 78"~8`,
			wildcard: false,
			fuzzy:    Fuzziness(8),
			boost:    DefaultBoost,
		},
		{
			name:     "test_phrase_with_space_and_star_symbol_and_fuzzy",
			input:    `"dsada* 78"~8`,
			want:     &FuzzyTerm{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, `*`, ` `, `78`}}, FuzzySymbol: "~8"},
			valueS:   `"dsada* 78"~8`,
			wildcard: false,
			fuzzy:    Fuzziness(8),
			boost:    DefaultBoost,
		},
		{
			name:     "test_phrase_with_space_and_escape_and_float_fuzzy",
			input:    `"dsada* 78"~8.1`,
			want:     &FuzzyTerm{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, `*`, ` `, `78`}}, FuzzySymbol: "~8.1"},
			valueS:   `"dsada* 78"~8.1`,
			wildcard: false,
			fuzzy:    Fuzziness(8.1),
			boost:    DefaultBoost,
		},
		{
			name:     "test_phrase_with_space_and_star_symbol_and_float_fuzzy",
			input:    `"dsada* 78"~8.6`,
			want:     &FuzzyTerm{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, `*`, ` `, `78`}}, FuzzySymbol: "~8.6"},
			valueS:   `"dsada* 78"~8.6`,
			wildcard: false,
			fuzzy:    Fuzziness(8.6),
			boost:    DefaultBoost,
		},
		{
			name:     "test_phrase_with_none_number_fuzzy",
			input:    `"dsada 78"~`,
			want:     &FuzzyTerm{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, ` `, `78`}}, FuzzySymbol: "~"},
			valueS:   `"dsada 78"~`,
			wildcard: false,
			fuzzy:    AutoFuzzy,
			boost:    DefaultBoost,
		},
		{
			name:     "test_phrase_with_star_symbol_and_none_number_fuzzy",
			input:    `"dsada* 78"~`,
			want:     &FuzzyTerm{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, `*`, ` `, `78`}}, FuzzySymbol: "~"},
			valueS:   `"dsada* 78"~`,
			wildcard: false,
			fuzzy:    AutoFuzzy,
			boost:    DefaultBoost,
		},
		{
			name:     "test_single_with_escape_and_wildcard",
			input:    `\/dsada\/\ dasda80980?*`,
			want:     &FuzzyTerm{SingleTerm: &SingleTerm{Begin: `\/`, Chars: []string{`dsada`, `\/\ `, `dasda`, `80980`, `?`, `*`}}},
			valueS:   `\/dsada\/\ dasda80980?*`,
			wildcard: true,
			fuzzy:    NoFuzzy,
			boost:    DefaultBoost,
		},
		{
			name:     "test_single_with_escape_and_wildcard_and_boost",
			input:    `\/dsada\/\ dasda80980?*\^\^^08`,
			want:     &FuzzyTerm{SingleTerm: &SingleTerm{Begin: `\/`, Chars: []string{`dsada`, `\/\ `, `dasda`, `80980`, `?`, `*`, `\^\^`}}, BoostSymbol: `^08`},
			valueS:   `\/dsada\/\ dasda80980?*\^\^^08`,
			wildcard: true,
			fuzzy:    NoFuzzy,
			boost:    BoostValue(8),
		},
		{
			name:     "test_single_with_escape_and_wildcard_and_none_number_boost",
			input:    `\/dsada\/\ dasda80980?*\^\^^`,
			want:     &FuzzyTerm{SingleTerm: &SingleTerm{Begin: `\/`, Chars: []string{`dsada`, `\/\ `, `dasda`, `80980`, `?`, `*`, `\^\^`}}, BoostSymbol: `^`},
			valueS:   `\/dsada\/\ dasda80980?*\^\^^`,
			wildcard: true,
			fuzzy:    NoFuzzy,
			boost:    DefaultBoost,
		},
		{
			name:     "test_single_with_escape_and_wildcard_and_fuzzy",
			input:    `\/dsada\/\ dasda80980?*\^\^~8`,
			want:     &FuzzyTerm{SingleTerm: &SingleTerm{Begin: `\/`, Chars: []string{`dsada`, `\/\ `, `dasda`, `80980`, `?`, `*`, `\^\^`}}, FuzzySymbol: `~8`},
			valueS:   `\/dsada\/\ dasda80980?*\^\^~8`,
			wildcard: true,
			fuzzy:    Fuzziness(8),
			boost:    DefaultBoost,
		},
		{
			name:     "test_single_with_escape_and_wildcard_and_none_number_fuzzy",
			input:    `\/dsada\/\ dasda80980?*\^\^~`,
			want:     &FuzzyTerm{SingleTerm: &SingleTerm{Begin: `\/`, Chars: []string{`dsada`, `\/\ `, `dasda`, `80980`, `?`, `*`, `\^\^`}}, FuzzySymbol: `~`},
			valueS:   `\/dsada\/\ dasda80980?*\^\^~`,
			wildcard: true,
			fuzzy:    AutoFuzzy,
			boost:    DefaultBoost,
		},
		{
			name:     "test_single_with_escape",
			input:    `\/dsada\/\ dasda80980\?\*`,
			want:     &FuzzyTerm{SingleTerm: &SingleTerm{Begin: `\/`, Chars: []string{`dsada`, `\/\ `, `dasda`, `80980`, `\?\*`}}},
			valueS:   `\/dsada\/\ dasda80980\?\*`,
			wildcard: false,
			fuzzy:    NoFuzzy,
			boost:    DefaultBoost,
		},
		{
			name:     "test_single_with_escape_and_number_boost",
			input:    `\/dsada\/\ dasda80980\?\*\^\^^08`,
			want:     &FuzzyTerm{SingleTerm: &SingleTerm{Begin: `\/`, Chars: []string{`dsada`, `\/\ `, `dasda`, `80980`, `\?\*\^\^`}}, BoostSymbol: `^08`},
			valueS:   `\/dsada\/\ dasda80980\?\*\^\^^08`,
			wildcard: false,
			fuzzy:    NoFuzzy,
			boost:    BoostValue(8.0),
		},
		{
			name:     "test_single_with_escape_and_fuzzy",
			input:    `\/dsada\/\ dasda80980\?\*\^\^~8`,
			want:     &FuzzyTerm{SingleTerm: &SingleTerm{Begin: `\/`, Chars: []string{`dsada`, `\/\ `, `dasda`, `80980`, `\?\*\^\^`}}, FuzzySymbol: `~8`},
			valueS:   `\/dsada\/\ dasda80980\?\*\^\^~8`,
			wildcard: false,
			fuzzy:    Fuzziness(8),
			boost:    DefaultBoost,
		},
		{
			name:     "test_single_with_escape_and_none_number_fuzzy",
			input:    `\/dsada\/\ dasda80980\?\*\^\^~`,
			want:     &FuzzyTerm{SingleTerm: &SingleTerm{Begin: `\/`, Chars: []string{`dsada`, `\/\ `, `dasda`, `80980`, `\?\*\^\^`}}, FuzzySymbol: `~`},
			valueS:   `\/dsada\/\ dasda80980\?\*\^\^~`,
			wildcard: false,
			fuzzy:    AutoFuzzy,
			boost:    DefaultBoost,
		},
		{
			name:     "test_single_with_escape_and_boost",
			input:    `\/dsada\/\ dasda80980?*\^\^^.8`,
			want:     &FuzzyTerm{SingleTerm: &SingleTerm{Begin: `\/`, Chars: []string{`dsada`, `\/\ `, `dasda`, `80980`, `?`, `*`, `\^\^`}}, BoostSymbol: `^.8`},
			valueS:   `\/dsada\/\ dasda80980?*\^\^^.8`,
			wildcard: true,
			fuzzy:    NoFuzzy,
			boost:    BoostValue(0.8),
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
			assert.Equal(t, tt.fuzzy, out.Fuzzy())
			assert.Equal(t, tt.boost, out.Boost())
		})
	}
	var out *FuzzyTerm
	assert.Empty(t, out.String())
	assert.Equal(t, NoFuzzy, out.Fuzzy())
	assert.Equal(t, NoBoost, out.Boost())
	assert.Equal(t, false, out.haveWildcard())
	assert.Equal(t, UNKNOWN_TERM_TYPE, out.GetTermType())
	_, err := out.Value(func(s string) (interface{}, error) { return s, nil })
	assert.Equal(t, ErrEmptyFuzzyTerm, err)

	out = &FuzzyTerm{}
	assert.Empty(t, out.String())
	assert.Equal(t, NoFuzzy, out.Fuzzy())
	assert.Equal(t, NoBoost, out.Boost())
	assert.Equal(t, false, out.haveWildcard())
	assert.Equal(t, UNKNOWN_TERM_TYPE, out.GetTermType())
	v, _ := out.Value(func(s string) (interface{}, error) { return s, nil })
	assert.Equal(t, "", v)
}

func TestFieldTermGroup(t *testing.T) {
	type test struct {
		name    string
		input   *FieldTermGroup
		tType   TermType
		valueS  interface{}
		wantErr error
	}

	for _, tt := range []test{
		{
			name:    "test_nil_field_term_group",
			input:   nil,
			tType:   UNKNOWN_TERM_TYPE,
			valueS:  nil,
			wantErr: ErrEmptyTermGroupElem,
		},
		{
			name:    "test_nil_logic_term_group",
			input:   &FieldTermGroup{},
			tType:   UNKNOWN_TERM_TYPE,
			valueS:  "",
			wantErr: nil,
		},
		{
			name: "test_single_term_group",
			input: &FieldTermGroup{
				SingleTerm: &SingleTerm{Begin: "123"},
			},
			tType:   SINGLE_TERM_TYPE,
			valueS:  "123",
			wantErr: nil,
		},
		{
			name: "test_phrase_term_group",
			input: &FieldTermGroup{
				PhraseTerm: &PhraseTerm{Chars: []string{"123"}},
			},
			tType:   PHRASE_TERM_TYPE,
			valueS:  "123",
			wantErr: nil,
		},
		{
			name: "test_single_range_term_group",
			input: &FieldTermGroup{
				SRangeTerm: &SRangeTerm{
					Symbol: ">",
					Value:  &RangeValue{SingleValue: []string{"123"}},
				},
			},
			tType: RANGE_TERM_TYPE,
			valueS: &Bound{
				LeftValue:    &RangeValue{SingleValue: []string{"123"}, SideFlag: false},
				RightValue:   &RangeValue{InfinityVal: "*", SideFlag: true},
				LeftInclude:  false,
				RightInclude: false,
			},
			wantErr: nil,
		},
		{
			name: "test_double_range_term_group",
			input: &FieldTermGroup{
				DRangeTerm: &DRangeTerm{
					LBRACKET: "{",
					RBRACKET: "}",
					LValue:   &RangeValue{SingleValue: []string{"123"}},
					RValue:   &RangeValue{SingleValue: []string{"456"}},
				},
			},
			tType: RANGE_TERM_TYPE,
			valueS: &Bound{
				LeftValue:    &RangeValue{SingleValue: []string{"123"}, SideFlag: false},
				RightValue:   &RangeValue{SingleValue: []string{"456"}, SideFlag: true},
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
	var out *FieldTermGroup
	assert.Empty(t, out.String())
	out = &FieldTermGroup{}
	assert.Empty(t, out.String())
}
