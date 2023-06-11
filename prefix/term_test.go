package prefix

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhuliquan/lucene_parser/term"
)

func TestTerm(t *testing.T) {
	type testCase struct {
		name      string
		input     *Term
		want1     string
		want2     interface{}
		boost     term.BoostValue
		termType  term.TermType
		bound     *term.Bound
		fuzziness term.Fuzziness
	}

	for _, tt := range []testCase{
		{
			name:      "test_nil_term",
			input:     nil,
			want1:     "",
			want2:     nil,
			boost:     term.NoBoost,
			termType:  term.UNKNOWN_TERM_TYPE,
			bound:     nil,
			fuzziness: term.NoFuzzy,
		},
		{
			name:      "test_empty_term",
			input:     &Term{},
			want1:     "",
			want2:     "",
			boost:     term.NoBoost,
			termType:  term.UNKNOWN_TERM_TYPE,
			bound:     nil,
			fuzziness: term.NoFuzzy,
		},
		{
			name:      "test_single",
			input:     &Term{FuzzyTerm: &term.FuzzyTerm{SingleTerm: &term.SingleTerm{Begin: "x"}, BoostSymbol: ""}},
			want1:     "x",
			want2:     "x",
			boost:     term.DefaultBoost,
			termType:  term.SINGLE_TERM_TYPE,
			bound:     nil,
			fuzziness: term.NoFuzzy,
		},
		{
			name:      "test_single_with_boost",
			input:     &Term{FuzzyTerm: &term.FuzzyTerm{SingleTerm: &term.SingleTerm{Begin: "x"}, BoostSymbol: "^8"}},
			want1:     "x^8",
			want2:     "x",
			boost:     term.BoostValue(8),
			termType:  term.SINGLE_TERM_TYPE | term.BOOST_TERM_TYPE,
			bound:     nil,
			fuzziness: term.NoFuzzy,
		},
		{
			name:      "test_single_with_fuzzy",
			input:     &Term{FuzzyTerm: &term.FuzzyTerm{SingleTerm: &term.SingleTerm{Begin: "x"}, FuzzySymbol: "~8"}},
			want1:     "x~8",
			want2:     "x",
			boost:     term.DefaultBoost,
			termType:  term.SINGLE_TERM_TYPE | term.FUZZY_TERM_TYPE,
			bound:     nil,
			fuzziness: term.Fuzziness(8),
		},
		{
			name:      "test_phrase",
			input:     &Term{FuzzyTerm: &term.FuzzyTerm{PhraseTerm: &term.PhraseTerm{Chars: []string{"x"}}}},
			want1:     "\"x\"",
			want2:     "x",
			boost:     term.DefaultBoost,
			termType:  term.PHRASE_TERM_TYPE,
			bound:     nil,
			fuzziness: term.NoFuzzy,
		},
		{
			name:      "test_phrase_with_boost",
			input:     &Term{FuzzyTerm: &term.FuzzyTerm{PhraseTerm: &term.PhraseTerm{Chars: []string{"x"}}, BoostSymbol: "^8"}},
			want1:     "\"x\"^8",
			want2:     "x",
			boost:     term.BoostValue(8),
			termType:  term.PHRASE_TERM_TYPE | term.BOOST_TERM_TYPE,
			bound:     nil,
			fuzziness: term.NoFuzzy,
		},
		{
			name:      "test_phrase_with_fuzzy",
			input:     &Term{FuzzyTerm: &term.FuzzyTerm{PhraseTerm: &term.PhraseTerm{Chars: []string{"x"}}, FuzzySymbol: "~8"}},
			want1:     "\"x\"~8",
			want2:     "x",
			boost:     term.DefaultBoost,
			termType:  term.PHRASE_TERM_TYPE | term.FUZZY_TERM_TYPE,
			bound:     nil,
			fuzziness: term.Fuzziness(8),
		},
		{
			name: "test_range",
			input: &Term{RangeTerm: &term.RangeTerm{DRangeTerm: &term.DRangeTerm{
				LBRACKET: "{",
				RBRACKET: "]",
				LValue:   &term.RangeValue{SingleValue: []string{"x"}},
				RValue:   &term.RangeValue{SingleValue: []string{"y"}},
			}}},
			want1: "{ x TO y ]",
			want2: &term.Bound{
				LeftValue:    &term.RangeValue{SingleValue: []string{"x"}},
				RightValue:   &term.RangeValue{SingleValue: []string{"y"}, SideFlag: true},
				LeftInclude:  false,
				RightInclude: true,
			},
			boost:    term.DefaultBoost,
			termType: term.RANGE_TERM_TYPE,
			bound: &term.Bound{
				LeftValue:    &term.RangeValue{SingleValue: []string{"x"}},
				RightValue:   &term.RangeValue{SingleValue: []string{"y"}, SideFlag: true},
				LeftInclude:  false,
				RightInclude: true,
			},
			fuzziness: term.NoFuzzy,
		},
		{
			name: "test_range_with_boost",
			input: &Term{RangeTerm: &term.RangeTerm{DRangeTerm: &term.DRangeTerm{
				LBRACKET: "{",
				RBRACKET: "]",
				LValue:   &term.RangeValue{SingleValue: []string{"x"}},
				RValue:   &term.RangeValue{SingleValue: []string{"y"}},
			}, BoostSymbol: "^8"}},
			want1: "{ x TO y ]^8",
			want2: &term.Bound{
				LeftValue:    &term.RangeValue{SingleValue: []string{"x"}},
				RightValue:   &term.RangeValue{SingleValue: []string{"y"}, SideFlag: true},
				LeftInclude:  false,
				RightInclude: true,
			},
			boost:    term.BoostValue(8),
			termType: term.RANGE_TERM_TYPE | term.BOOST_TERM_TYPE,
			bound: &term.Bound{
				LeftValue:    &term.RangeValue{SingleValue: []string{"x"}},
				RightValue:   &term.RangeValue{SingleValue: []string{"y"}, SideFlag: true},
				LeftInclude:  false,
				RightInclude: true,
			},
			fuzziness: term.NoFuzzy,
		},
		{
			name:      "test_regex",
			input:     &Term{RegexpTerm: &term.RegexpTerm{Chars: []string{"x"}}},
			want1:     "/x/",
			want2:     "x",
			boost:     term.DefaultBoost,
			termType:  term.REGEXP_TERM_TYPE,
			bound:     nil,
			fuzziness: term.NoFuzzy,
		},
		{
			name: "test_group",
			input: &Term{
				TermGroup: &TermGroup{
					PrefixTermGroup: &PrefixTermGroup{
						PrefixTerms: []*PrefixOperatorTerm{
							{
								Symbol: "+",
								FieldTermGroup: &term.FieldTermGroup{
									SingleTerm: &term.SingleTerm{Begin: "x"},
								},
							},
						},
					},
					BoostSymbol: "",
				},
			},
			want1:     "( +x )",
			want2:     "+x",
			boost:     term.DefaultBoost,
			termType:  term.GROUP_TERM_TYPE,
			bound:     nil,
			fuzziness: term.NoFuzzy,
		},
		{
			name: "test_group_with_boost",
			input: &Term{
				TermGroup: &TermGroup{
					PrefixTermGroup: &PrefixTermGroup{
						PrefixTerms: []*PrefixOperatorTerm{
							{
								Symbol: "+",
								FieldTermGroup: &term.FieldTermGroup{
									SingleTerm: &term.SingleTerm{Begin: "x"},
								},
							},
							{
								Symbol: "-",
								FieldTermGroup: &term.FieldTermGroup{
									SingleTerm: &term.SingleTerm{Begin: "y"},
								},
							},
						},
					},
					BoostSymbol: "^8",
				},
			},
			want1:     "( +x -y )^8",
			want2:     "+x -y",
			boost:     term.BoostValue(8),
			termType:  term.GROUP_TERM_TYPE | term.BOOST_TERM_TYPE,
			bound:     nil,
			fuzziness: term.NoFuzzy,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want1, tt.input.String())
			out, _ := tt.input.Value(func(s string) (interface{}, error) { return s, nil })
			assert.Equal(t, tt.want2, out)
			assert.Equal(t, tt.boost, tt.input.Boost())
			assert.Equal(t, tt.termType, tt.input.GetTermType())
			assert.Equal(t, tt.bound, tt.input.GetBound())
			assert.Equal(t, tt.fuzziness, tt.input.Fuzziness())
		})
	}
}
