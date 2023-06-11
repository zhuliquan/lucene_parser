package term

import (
	"testing"

	"github.com/alecthomas/participle"
	"github.com/stretchr/testify/assert"
	"github.com/zhuliquan/lucene_parser/operator"
	"github.com/zhuliquan/lucene_parser/token"
)

func TestTerm(t *testing.T) {
	var termParser = participle.MustBuild(
		&Term{},
		participle.Lexer(token.Lexer),
	)

	type testCase struct {
		name      string
		input     string
		want      *Term
		wantStr   string
		boost     BoostValue
		valueS    interface{}
		fuzziness Fuzziness
		bound     *Bound
		tType     TermType
	}
	var testCases = []testCase{
		{
			name:      "test_phrase_with_space",
			input:     `"dsada 78"`,
			want:      &Term{FuzzyTerm: &FuzzyTerm{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, ` `, `78`}}}},
			wantStr:   "\"dsada 78\"",
			boost:     DefaultBoost,
			valueS:    "dsada 78",
			fuzziness: NoFuzzy,
			bound:     nil,
			tType:     PHRASE_TERM_TYPE,
		},
		{
			name:      "test_phrase_with_space_and_boost",
			input:     `"dsada 78"^08`,
			want:      &Term{FuzzyTerm: &FuzzyTerm{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, ` `, `78`}}, BoostSymbol: "^08"}},
			wantStr:   `"dsada 78"^08`,
			boost:     BoostValue(8.0),
			valueS:    "dsada 78",
			fuzziness: NoFuzzy,
			bound:     nil,
			tType:     PHRASE_TERM_TYPE | BOOST_TERM_TYPE,
		},
		{
			name:      "test_phrase_with_space_and_fuzzy_01",
			input:     `"dsada 78"~8`,
			want:      &Term{FuzzyTerm: &FuzzyTerm{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, ` `, `78`}}, FuzzySymbol: "~8"}},
			wantStr:   `"dsada 78"~8`,
			boost:     DefaultBoost,
			valueS:    "dsada 78",
			fuzziness: Fuzziness(8.0),
			bound:     nil,
			tType:     PHRASE_TERM_TYPE | FUZZY_TERM_TYPE,
		},
		{
			name:      "test_phrase_with_space_and_fuzzy_02",
			input:     `"dsada 78"~8.1`,
			want:      &Term{FuzzyTerm: &FuzzyTerm{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, ` `, `78`}}, FuzzySymbol: "~8.1"}},
			wantStr:   `"dsada 78"~8.1`,
			boost:     DefaultBoost,
			valueS:    "dsada 78",
			fuzziness: Fuzziness(8.1),
			bound:     nil,
			tType:     PHRASE_TERM_TYPE | FUZZY_TERM_TYPE,
		},
		{
			name:      "test_phrase_with_space_and_fuzzy_03",
			input:     `"dsada 78"~8.6`,
			want:      &Term{FuzzyTerm: &FuzzyTerm{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, ` `, `78`}}, FuzzySymbol: "~8.6"}},
			wantStr:   `"dsada 78"~8.6`,
			boost:     DefaultBoost,
			valueS:    "dsada 78",
			fuzziness: Fuzziness(8.6),
			bound:     nil,
			tType:     PHRASE_TERM_TYPE | FUZZY_TERM_TYPE,
		},
		{
			name:      "test_regex",
			input:     `/dsada 78/`,
			want:      &Term{RegexpTerm: &RegexpTerm{Chars: []string{`dsada`, ` `, `78`}}},
			wantStr:   `/dsada 78/`,
			boost:     DefaultBoost,
			valueS:    "dsada 78",
			fuzziness: NoFuzzy,
			bound:     nil,
			tType:     REGEXP_TERM_TYPE,
		},
		{
			name:      "test_single_with_escape_and_wildcard",
			input:     `\/dsada\/\ dasda80980?*`,
			want:      &Term{FuzzyTerm: &FuzzyTerm{SingleTerm: &SingleTerm{Begin: `\/`, Chars: []string{`dsada`, `\/\ `, `dasda`, `80980`, `?`, `*`}}}},
			wantStr:   `\/dsada\/\ dasda80980?*`,
			boost:     DefaultBoost,
			valueS:    `\/dsada\/\ dasda80980?*`,
			fuzziness: NoFuzzy,
			bound:     nil,
			tType:     SINGLE_TERM_TYPE | WILDCARD_TERM_TYPE,
		},
		{
			name:      "test_single_with_escape_and_wildcard_and_boost",
			input:     `\/dsada\/\ dasda80980?*\^\^^08`,
			want:      &Term{FuzzyTerm: &FuzzyTerm{SingleTerm: &SingleTerm{Begin: `\/`, Chars: []string{`dsada`, `\/\ `, `dasda`, `80980`, `?`, `*`, `\^\^`}}, BoostSymbol: `^08`}},
			wantStr:   `\/dsada\/\ dasda80980?*\^\^^08`,
			boost:     BoostValue(8),
			valueS:    `\/dsada\/\ dasda80980?*\^\^`,
			fuzziness: NoFuzzy,
			bound:     nil,
			tType:     SINGLE_TERM_TYPE | BOOST_TERM_TYPE | WILDCARD_TERM_TYPE,
		},
		{
			name:      "test_single_with_escape_and_wildcard_and_fuzzy",
			input:     `\/dsada\/\ dasda80980?*\^\^~8`,
			want:      &Term{FuzzyTerm: &FuzzyTerm{SingleTerm: &SingleTerm{Begin: `\/`, Chars: []string{`dsada`, `\/\ `, `dasda`, `80980`, `?`, `*`, `\^\^`}}, FuzzySymbol: `~8`}},
			wantStr:   `\/dsada\/\ dasda80980?*\^\^~8`,
			boost:     DefaultBoost,
			valueS:    `\/dsada\/\ dasda80980?*\^\^`,
			fuzziness: Fuzziness(8),
			bound:     nil,
			tType:     SINGLE_TERM_TYPE | FUZZY_TERM_TYPE | WILDCARD_TERM_TYPE,
		},
		{
			name:  "test_double_range_left_include_and_right_include_with_boost",
			input: `[1 TO 2]^7`,
			want: &Term{RangeTerm: &RangeTerm{
				DRangeTerm: &DRangeTerm{
					LBRACKET: "[",
					LValue:   &RangeValue{SingleValue: []string{"1"}},
					RValue:   &RangeValue{SingleValue: []string{"2"}},
					RBRACKET: "]",
				},
				BoostSymbol: "^7",
			}},
			wantStr: `[ 1 TO 2 ]^7`,
			boost:   BoostValue(7),
			valueS: &Bound{
				LeftValue:    &RangeValue{SingleValue: []string{"1"}, SideFlag: false},
				RightValue:   &RangeValue{SingleValue: []string{"2"}, SideFlag: true},
				LeftInclude:  true,
				RightInclude: true,
			},
			fuzziness: NoFuzzy,
			bound: &Bound{
				LeftValue:    &RangeValue{SingleValue: []string{"1"}, SideFlag: false},
				RightValue:   &RangeValue{SingleValue: []string{"2"}, SideFlag: true},
				LeftInclude:  true,
				RightInclude: true,
			},
			tType: RANGE_TERM_TYPE | BOOST_TERM_TYPE,
		},
		{
			name:  "test_double_range_left_include_and_right_exclude_with_boost",
			input: `[1 TO 2 }^7`,
			want: &Term{RangeTerm: &RangeTerm{
				DRangeTerm: &DRangeTerm{
					LBRACKET: "[",
					LValue:   &RangeValue{SingleValue: []string{"1"}},
					RValue:   &RangeValue{SingleValue: []string{"2"}},
					RBRACKET: "}",
				},
				BoostSymbol: "^7",
			}},
			wantStr: `[ 1 TO 2 }^7`,
			boost:   BoostValue(7),
			valueS: &Bound{
				LeftValue:    &RangeValue{SingleValue: []string{"1"}, SideFlag: false},
				RightValue:   &RangeValue{SingleValue: []string{"2"}, SideFlag: true},
				LeftInclude:  true,
				RightInclude: false,
			},
			fuzziness: NoFuzzy,
			bound: &Bound{
				LeftValue:    &RangeValue{SingleValue: []string{"1"}, SideFlag: false},
				RightValue:   &RangeValue{SingleValue: []string{"2"}, SideFlag: true},
				LeftInclude:  true,
				RightInclude: false,
			},
			tType: RANGE_TERM_TYPE | BOOST_TERM_TYPE,
		},
		{
			name:  `test_double_range_left_exclude_and_right_exclude_with_boost`,
			input: `{ 1 TO 2}^7`,
			want: &Term{RangeTerm: &RangeTerm{
				DRangeTerm: &DRangeTerm{
					LBRACKET: "{",
					LValue:   &RangeValue{SingleValue: []string{"1"}},
					RValue:   &RangeValue{SingleValue: []string{"2"}},
					RBRACKET: "}",
				},
				BoostSymbol: "^7",
			}},
			wantStr: `{ 1 TO 2 }^7`,
			boost:   BoostValue(7),
			valueS: &Bound{
				LeftValue:    &RangeValue{SingleValue: []string{"1"}, SideFlag: false},
				RightValue:   &RangeValue{SingleValue: []string{"2"}, SideFlag: true},
				LeftInclude:  false,
				RightInclude: false,
			},
			fuzziness: NoFuzzy,
			bound: &Bound{
				LeftValue:    &RangeValue{SingleValue: []string{"1"}, SideFlag: false},
				RightValue:   &RangeValue{SingleValue: []string{"2"}, SideFlag: true},
				LeftInclude:  false,
				RightInclude: false,
			},
			tType: RANGE_TERM_TYPE | BOOST_TERM_TYPE,
		},
		{
			name:  `test_double_range_left_exclude_and_right_exclude_with_boost`,
			input: `{ 1 TO 2]^7`,
			want: &Term{RangeTerm: &RangeTerm{
				DRangeTerm: &DRangeTerm{
					LBRACKET: "{",
					LValue:   &RangeValue{SingleValue: []string{"1"}},
					RValue:   &RangeValue{SingleValue: []string{"2"}},
					RBRACKET: "]",
				},
				BoostSymbol: "^7",
			}},
			wantStr: `{ 1 TO 2 ]^7`,
			boost:   BoostValue(7),
			valueS: &Bound{
				LeftValue:    &RangeValue{SingleValue: []string{"1"}, SideFlag: false},
				RightValue:   &RangeValue{SingleValue: []string{"2"}, SideFlag: true},
				LeftInclude:  false,
				RightInclude: true,
			},
			fuzziness: NoFuzzy,
			bound: &Bound{
				LeftValue:    &RangeValue{SingleValue: []string{"1"}, SideFlag: false},
				RightValue:   &RangeValue{SingleValue: []string{"2"}, SideFlag: true},
				LeftInclude:  false,
				RightInclude: true,
			},
			tType: RANGE_TERM_TYPE | BOOST_TERM_TYPE,
		},
		{
			name:  `test_double_range_left_exclude_and_right_inf_with_boost`,
			input: `[10 TO *]^7`,
			want: &Term{RangeTerm: &RangeTerm{
				DRangeTerm: &DRangeTerm{
					LBRACKET: "[",
					LValue:   &RangeValue{SingleValue: []string{"10"}},
					RValue:   &RangeValue{InfinityVal: "*"},
					RBRACKET: "]",
				},
				BoostSymbol: "^7",
			}},
			wantStr: `[ 10 TO * ]^7`,
			boost:   BoostValue(7),
			valueS: &Bound{
				LeftValue:    &RangeValue{SingleValue: []string{"10"}, SideFlag: false},
				RightValue:   &RangeValue{InfinityVal: "*", SideFlag: true},
				LeftInclude:  true,
				RightInclude: false,
			},
			fuzziness: NoFuzzy,
			bound: &Bound{
				LeftValue:    &RangeValue{SingleValue: []string{"10"}, SideFlag: false},
				RightValue:   &RangeValue{InfinityVal: "*", SideFlag: true},
				LeftInclude:  true,
				RightInclude: false,
			},
			tType: RANGE_TERM_TYPE | BOOST_TERM_TYPE,
		},
		{
			name:  `test_double_range_left_inf_and_right_exclude_with_boost`,
			input: `{* TO 2012-01-01}`,
			want: &Term{RangeTerm: &RangeTerm{
				DRangeTerm: &DRangeTerm{
					LBRACKET: "{",
					LValue:   &RangeValue{InfinityVal: "*"},
					RValue:   &RangeValue{SingleValue: []string{"2012", "-", "01", "-", "01"}},
					RBRACKET: "}",
				},
			}},
			wantStr: `{ * TO 2012-01-01 }`,
			boost:   DefaultBoost,
			valueS: &Bound{
				LeftValue:    &RangeValue{InfinityVal: "*", SideFlag: false},
				RightValue:   &RangeValue{SingleValue: []string{"2012", "-", "01", "-", "01"}, SideFlag: true},
				LeftInclude:  false,
				RightInclude: false,
			},
			fuzziness: NoFuzzy,
			bound: &Bound{
				LeftValue:    &RangeValue{InfinityVal: "*", SideFlag: false},
				RightValue:   &RangeValue{SingleValue: []string{"2012", "-", "01", "-", "01"}, SideFlag: true},
				LeftInclude:  false,
				RightInclude: false,
			},
			tType: RANGE_TERM_TYPE,
		},
		{
			name:  `test_double_range_left_inf_and_right_exclude`,
			input: `{* TO "2012-01-01 09:08:16"}`,
			want: &Term{RangeTerm: &RangeTerm{
				DRangeTerm: &DRangeTerm{
					LBRACKET: "{",
					LValue:   &RangeValue{InfinityVal: "*"},
					RValue:   &RangeValue{PhraseValue: []string{"2012", "-", "01", "-", "01", " ", "09", ":", "08", ":", "16"}},
					RBRACKET: "}",
				},
			}},
			wantStr: `{ * TO "2012-01-01 09:08:16" }`,
			boost:   DefaultBoost,
			valueS: &Bound{
				LeftValue:    &RangeValue{InfinityVal: "*", SideFlag: false},
				RightValue:   &RangeValue{PhraseValue: []string{"2012", "-", "01", "-", "01", " ", "09", ":", "08", ":", "16"}, SideFlag: true},
				LeftInclude:  false,
				RightInclude: false,
			},
			fuzziness: NoFuzzy,
			bound: &Bound{
				LeftValue:    &RangeValue{InfinityVal: "*", SideFlag: false},
				RightValue:   &RangeValue{PhraseValue: []string{"2012", "-", "01", "-", "01", " ", "09", ":", "08", ":", "16"}, SideFlag: true},
				LeftInclude:  false,
				RightInclude: false,
			},
			tType: RANGE_TERM_TYPE,
		},
		{
			name:  "test_single_range_lte_phrase",
			input: `<="dsada 78"`,
			want: &Term{
				RangeTerm: &RangeTerm{
					SRangeTerm: &SRangeTerm{
						Symbol: "<=",
						Value:  &RangeValue{PhraseValue: []string{`dsada`, ` `, `78`}},
					},
				},
			},
			wantStr: `{ * TO "dsada 78" ]`,
			boost:   DefaultBoost,
			valueS: &Bound{
				LeftValue:    &RangeValue{InfinityVal: "*", SideFlag: false},
				RightValue:   &RangeValue{PhraseValue: []string{`dsada`, ` `, `78`}, SideFlag: true},
				LeftInclude:  false,
				RightInclude: true,
			},
			fuzziness: NoFuzzy,
			bound: &Bound{
				LeftValue:    &RangeValue{InfinityVal: "*", SideFlag: false},
				RightValue:   &RangeValue{PhraseValue: []string{`dsada`, ` `, `78`}, SideFlag: true},
				LeftInclude:  false,
				RightInclude: true,
			},
			tType: RANGE_TERM_TYPE,
		},
		{
			name:  "test_single_range_lt_phrase_with_boost",
			input: `<"dsada 78"^08`,
			want: &Term{
				RangeTerm: &RangeTerm{
					SRangeTerm: &SRangeTerm{
						Symbol: "<",
						Value:  &RangeValue{PhraseValue: []string{`dsada`, ` `, `78`}},
					},
					BoostSymbol: "^08",
				},
			},
			wantStr: `{ * TO "dsada 78" }^08`,
			boost:   BoostValue(8),
			valueS: &Bound{
				LeftValue:    &RangeValue{InfinityVal: "*", SideFlag: false},
				RightValue:   &RangeValue{PhraseValue: []string{`dsada`, ` `, `78`}, SideFlag: true},
				LeftInclude:  false,
				RightInclude: false,
			},
			fuzziness: NoFuzzy,
			bound: &Bound{
				LeftValue:    &RangeValue{InfinityVal: "*", SideFlag: false},
				RightValue:   &RangeValue{PhraseValue: []string{`dsada`, ` `, `78`}, SideFlag: true},
				LeftInclude:  false,
				RightInclude: false,
			},
			tType: RANGE_TERM_TYPE | BOOST_TERM_TYPE,
		},
		{
			name:  "test_single_range_gt_phrase_with_boost",
			input: `>"dsada 78"^080`,
			want: &Term{
				RangeTerm: &RangeTerm{
					SRangeTerm: &SRangeTerm{
						Symbol: ">",
						Value:  &RangeValue{PhraseValue: []string{`dsada`, ` `, `78`}},
					},
					BoostSymbol: "^080",
				},
			},
			wantStr: `{ "dsada 78" TO * }^080`,
			boost:   BoostValue(80),
			valueS: &Bound{
				LeftValue:    &RangeValue{PhraseValue: []string{`dsada`, ` `, `78`}, SideFlag: false},
				RightValue:   &RangeValue{InfinityVal: "*", SideFlag: true},
				LeftInclude:  false,
				RightInclude: false,
			},
			fuzziness: NoFuzzy,
			bound: &Bound{
				LeftValue:    &RangeValue{PhraseValue: []string{`dsada`, ` `, `78`}, SideFlag: false},
				RightValue:   &RangeValue{InfinityVal: "*", SideFlag: true},
				LeftInclude:  false,
				RightInclude: false,
			},
			tType: RANGE_TERM_TYPE | BOOST_TERM_TYPE,
		},
		{
			name:  "test_single_range_lte_single",
			input: `<=dsada\ 78`,
			want: &Term{
				RangeTerm: &RangeTerm{
					SRangeTerm: &SRangeTerm{
						Symbol: "<=",
						Value:  &RangeValue{SingleValue: []string{`dsada`, `\ `, `78`}},
					},
				},
			},
			wantStr: `{ * TO dsada\ 78 ]`,
			boost:   DefaultBoost,
			valueS: &Bound{
				LeftValue:    &RangeValue{InfinityVal: "*", SideFlag: false},
				RightValue:   &RangeValue{SingleValue: []string{`dsada`, `\ `, `78`}, SideFlag: true},
				LeftInclude:  false,
				RightInclude: true,
			},
			fuzziness: NoFuzzy,
			bound: &Bound{
				LeftValue:    &RangeValue{InfinityVal: "*", SideFlag: false},
				RightValue:   &RangeValue{SingleValue: []string{`dsada`, `\ `, `78`}, SideFlag: true},
				LeftInclude:  false,
				RightInclude: true,
			},
			tType: RANGE_TERM_TYPE,
		},
		{
			name:  "test_single_range_lt_single_with_boost",
			input: `<dsada\ 78^08`,
			want: &Term{
				RangeTerm: &RangeTerm{
					SRangeTerm: &SRangeTerm{
						Symbol: "<",
						Value:  &RangeValue{SingleValue: []string{`dsada`, `\ `, `78`}},
					},
					BoostSymbol: "^08",
				},
			},
			wantStr: `{ * TO dsada\ 78 }^08`,
			boost:   BoostValue(8),
			valueS: &Bound{
				LeftValue:    &RangeValue{InfinityVal: "*", SideFlag: false},
				RightValue:   &RangeValue{SingleValue: []string{`dsada`, `\ `, `78`}, SideFlag: true},
				LeftInclude:  false,
				RightInclude: false,
			},
			fuzziness: NoFuzzy,
			bound: &Bound{
				LeftValue:    &RangeValue{InfinityVal: "*", SideFlag: false},
				RightValue:   &RangeValue{SingleValue: []string{`dsada`, `\ `, `78`}, SideFlag: true},
				LeftInclude:  false,
				RightInclude: false,
			},
			tType: RANGE_TERM_TYPE | BOOST_TERM_TYPE,
		},
		{
			name:  "test_single_range_gt_single_with_boost",
			input: `>dsada\ 78^080`,
			want: &Term{
				RangeTerm: &RangeTerm{
					SRangeTerm: &SRangeTerm{
						Symbol: ">",
						Value:  &RangeValue{SingleValue: []string{`dsada`, `\ `, `78`}},
					},
					BoostSymbol: "^080",
				},
			},
			wantStr: `{ dsada\ 78 TO * }^080`,
			boost:   BoostValue(80.0),
			valueS: &Bound{
				LeftValue:    &RangeValue{SingleValue: []string{`dsada`, `\ `, `78`}, SideFlag: false},
				RightValue:   &RangeValue{InfinityVal: "*", SideFlag: true},
				LeftInclude:  false,
				RightInclude: false,
			},
			fuzziness: NoFuzzy,
			bound: &Bound{
				LeftValue:    &RangeValue{SingleValue: []string{`dsada`, `\ `, `78`}, SideFlag: false},
				RightValue:   &RangeValue{InfinityVal: "*", SideFlag: true},
				LeftInclude:  false,
				RightInclude: false,
			},
			tType: RANGE_TERM_TYPE | BOOST_TERM_TYPE,
		},
		{
			name:      "test_regex",
			input:     `/\d+\d+\.\d+.+/`,
			want:      &Term{RegexpTerm: &RegexpTerm{Chars: []string{`\`, `d`, `+`, `\`, `d`, `+`, `\`, `.`, `\`, `d`, `+`, `.`, `+`}}},
			wantStr:   `/\d+\d+\.\d+.+/`,
			boost:     DefaultBoost,
			valueS:    `\d+\d+\.\d+.+`,
			fuzziness: NoFuzzy,
			bound:     nil,
			tType:     REGEXP_TERM_TYPE,
		},
		{
			name:  "test_or_group",
			input: `(foo or bar)`,
			want: &Term{TermGroup: &TermGroup{
				LogicTermGroup: &LogicTermGroup{
					OrTermGroup: &OrTermGroup{
						AndTermGroup: &AndTermGroup{
							FieldTermGroup: &FieldTermGroup{SingleTerm: &SingleTerm{Begin: "foo"}},
						},
					},
					OSTermGroup: []*OSTermGroup{
						{
							OrSymbol: &operator.OrSymbol{Symbol: "or"},
							OrTermGroup: &OrTermGroup{
								AndTermGroup: &AndTermGroup{
									FieldTermGroup: &FieldTermGroup{SingleTerm: &SingleTerm{Begin: "bar"}},
								},
							},
						},
					},
				},
			}},
			wantStr:   `( foo OR bar )`,
			boost:     DefaultBoost,
			valueS:    `foo OR bar`,
			fuzziness: NoFuzzy,
			bound:     nil,
			tType:     GROUP_TERM_TYPE,
		},
		{
			name:  "test_term_group_with_boost",
			input: `(foo OR bar or [1 TO 2] AND >10 AND <20 and "abc")^7`,
			want: &Term{TermGroup: &TermGroup{
				LogicTermGroup: &LogicTermGroup{
					OrTermGroup: &OrTermGroup{
						AndTermGroup: &AndTermGroup{
							FieldTermGroup: &FieldTermGroup{SingleTerm: &SingleTerm{Begin: "foo"}},
						},
					},
					OSTermGroup: []*OSTermGroup{
						{
							OrSymbol: &operator.OrSymbol{Symbol: "OR"},
							OrTermGroup: &OrTermGroup{
								AndTermGroup: &AndTermGroup{
									FieldTermGroup: &FieldTermGroup{SingleTerm: &SingleTerm{Begin: "bar"}},
								},
							},
						},
						{
							OrSymbol: &operator.OrSymbol{Symbol: "or"},
							OrTermGroup: &OrTermGroup{
								AndTermGroup: &AndTermGroup{
									FieldTermGroup: &FieldTermGroup{DRangeTerm: &DRangeTerm{
										LBRACKET: "[",
										LValue:   &RangeValue{SingleValue: []string{"1"}},
										RValue:   &RangeValue{SingleValue: []string{"2"}},
										RBRACKET: "]",
									}},
								},
								AnSTermGroup: []*AnSTermGroup{
									{
										AndSymbol: &operator.AndSymbol{Symbol: "AND"},
										AndTermGroup: &AndTermGroup{
											FieldTermGroup: &FieldTermGroup{
												SRangeTerm: &SRangeTerm{
													Symbol: ">",
													Value:  &RangeValue{SingleValue: []string{"10"}},
												},
											},
										},
									},
									{
										AndSymbol: &operator.AndSymbol{Symbol: "AND"},
										AndTermGroup: &AndTermGroup{
											FieldTermGroup: &FieldTermGroup{
												SRangeTerm: &SRangeTerm{
													Symbol: "<",
													Value:  &RangeValue{SingleValue: []string{"20"}},
												},
											},
										},
									},
									{
										AndSymbol: &operator.AndSymbol{Symbol: "and"},
										AndTermGroup: &AndTermGroup{
											FieldTermGroup: &FieldTermGroup{
												PhraseTerm: &PhraseTerm{
													Chars: []string{"abc"},
												},
											},
										},
									},
								},
							},
						},
					},
				},
				BoostSymbol: "^7",
			}},
			wantStr:   `( foo OR bar OR [ 1 TO 2 ] AND { 10 TO * } AND { * TO 20 } AND "abc" )^7`,
			boost:     BoostValue(7),
			valueS:    `foo OR bar OR [ 1 TO 2 ] AND { 10 TO * } AND { * TO 20 } AND "abc"`,
			fuzziness: NoFuzzy,
			bound:     nil,
			tType:     GROUP_TERM_TYPE | BOOST_TERM_TYPE,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var out = &Term{}
			err := termParser.ParseString(tt.input, out)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, out)
			assert.Equal(t, tt.wantStr, out.String())
			assert.Equal(t, tt.boost, out.Boost())
			v, _ := out.Value(func(s string) (interface{}, error) { return s, nil })
			assert.Equal(t, tt.valueS, v)
			assert.Equal(t, tt.fuzziness, out.Fuzziness())
			assert.Equal(t, tt.bound, out.GetBound())
			assert.Equal(t, tt.tType, out.GetTermType())
		})
	}
	var out *Term
	assert.Equal(t, "", out.String())
	assert.Equal(t, NoBoost, out.Boost())
	assert.Equal(t, NoFuzzy, out.Fuzziness())
	assert.Equal(t, UNKNOWN_TERM_TYPE, out.GetTermType())
	_, err := out.Value(func(s string) (interface{}, error) { return s, nil })
	assert.Equal(t, ErrEmptyTerm, err)
	out = &Term{}
	assert.Equal(t, "", out.String())
	assert.Equal(t, NoBoost, out.Boost())
	assert.Equal(t, NoFuzzy, out.Fuzziness())
	assert.Equal(t, UNKNOWN_TERM_TYPE, out.GetTermType())
	v, _ := out.Value(func(s string) (interface{}, error) { return s, nil })
	assert.Equal(t, "", v)
}

func TestTermIsRegexp(t *testing.T) {
	var termParser = participle.MustBuild(
		&Term{},
		participle.Lexer(token.Lexer),
	)

	type testCase struct {
		name  string
		input string
		want  bool
	}

	var testCases = []testCase{
		{
			name:  "test_no_regex",
			input: `12313\+90`,
			want:  false,
		},
		{
			name:  "tes_regex",
			input: `/[1-9]+\.\d+/`,
			want:  true,
		},
		{
			name:  "test_phrase",
			input: `"dsad 7089"`,
			want:  false,
		},
		{
			name:  "test_range",
			input: `[1 TO 454 ]`,
			want:  false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var out = &Term{}
			err := termParser.ParseString(tt.input, out)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, out.GetTermType()&REGEXP_TERM_TYPE == REGEXP_TERM_TYPE)
		})
	}
}

func TestTermIsWildcard(t *testing.T) {

	var termParser = participle.MustBuild(
		&Term{},
		participle.Lexer(token.Lexer),
	)

	type testCase struct {
		name  string
		input string
		want  bool
	}

	var testCases = []testCase{
		{
			name:  "test_?_wildcard",
			input: `12313?`,
			want:  true,
		},
		{
			name:  "test_escape_?",
			input: `12313\?`,
			want:  false,
		},
		{
			name:  "test_*_wildcard",
			input: `12313*`,
			want:  true,
		},
		{
			name:  "test_escape_*",
			input: `12313\*`,
			want:  false,
		},
		{
			name:  "test_regex",
			input: `/[1-9]+\.\d+/`,
			want:  false,
		},
		{
			name:  "test_?_wildcard_and_?_wildcard",
			input: `"dsad?\? 7089*"`,
			want:  false,
		},
		{
			name:  "test_phrase",
			input: `"dsadad 789"`,
			want:  false,
		},
		{
			name:  "test_range",
			input: `[1 TO 2]`,
			want:  false,
		},
		{
			name:  "test_single",
			input: `"178"`,
			want:  false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var out = &Term{}
			err := termParser.ParseString(tt.input, out)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, out.GetTermType()&WILDCARD_TERM_TYPE == WILDCARD_TERM_TYPE)
			// 利用自身缓冲再次尝试
			assert.Equal(t, tt.want, out.GetTermType()&WILDCARD_TERM_TYPE == WILDCARD_TERM_TYPE)
		})
	}
}

func TestTermIsRange(t *testing.T) {
	var termParser = participle.MustBuild(
		&Term{},
		participle.Lexer(token.Lexer),
	)

	type testCase struct {
		name  string
		input string
		want  bool
	}

	var testCases = []testCase{
		{
			name:  "test_single",
			input: `12313\+90`,
			want:  false,
		},
		{
			name:  "test_regex",
			input: `/[1-9]+\.\d+/`,
			want:  false,
		},
		{
			name:  "rest_phrase",
			input: `"dsad 7089"`,
			want:  false,
		},
		{
			name:  "test_range",
			input: `[1 TO 454 ]`,
			want:  true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var out = &Term{}
			err := termParser.ParseString(tt.input, out)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, out.GetTermType()&RANGE_TERM_TYPE == RANGE_TERM_TYPE)
		})
	}
}

func TestTermFuzziness(t *testing.T) {

	var termParser = participle.MustBuild(
		&Term{},
		participle.Lexer(token.Lexer),
	)

	type testCase struct {
		name  string
		input string
		want  Fuzziness
	}

	var testCases = []testCase{
		{
			name:  "test_no_fuzzy",
			input: `12313\+90`,
			want:  NoFuzzy,
		},
		{
			name:  "test_regex",
			input: `/[1-9]+\.\d+/`,
			want:  NoFuzzy,
		},
		{
			name:  "test_phrase",
			input: `"dsad 7089"`,
			want:  NoFuzzy,
		},
		{
			name:  "test_range",
			input: `[1 TO 454 ]`,
			want:  NoFuzzy,
		},
		{
			name:  "test_single_with_fuzzy",
			input: `12313\+90~3`,
			want:  Fuzziness(3),
		},
		{
			name:  "test_phrase_with_fuzzy",
			input: `"dsad 7089"~3`,
			want:  Fuzziness(3),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var out = &Term{}
			err := termParser.ParseString(tt.input, out)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, out.Fuzziness())
		})
	}

}

func TestTermBoost(t *testing.T) {

	var termParser = participle.MustBuild(
		&Term{},
		participle.Lexer(token.Lexer),
	)

	type testCase struct {
		name  string
		input string
		want  BoostValue
	}

	var testCases = []testCase{
		{
			name:  "test_single_default_boost",
			input: `12313\+90`,
			want:  DefaultBoost,
		},
		{
			name:  "test_regex_default_boost",
			input: `/[1-9]+\.\d+/`,
			want:  DefaultBoost,
		},
		{
			name:  "test_phrase_default_boost",
			input: `"dsad 7089"`,
			want:  DefaultBoost,
		},
		{
			name:  "test_range_default_boost",
			input: `[1 TO 454 ]`,
			want:  DefaultBoost,
		},
		{
			name:  "test_single_with_float_boost",
			input: `12313\+90^1.2`,
			want:  BoostValue(1.2),
		},
		{
			name:  "test_single_with_less_one_boost",
			input: `12313\+90^0.2`,
			want:  BoostValue(0.2),
		},
		{
			name:  "test_single_with_great_one_boost",
			input: `"dsad 7089"^3.8`,
			want:  BoostValue(3.8),
		},
		{
			name:  "test_phrase_with_less_one_boost",
			input: `"dsad 7089"^0.8`,
			want:  BoostValue(0.8),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var out = &Term{}
			err := termParser.ParseString(tt.input, out)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, out.Boost())
		})
	}
}
