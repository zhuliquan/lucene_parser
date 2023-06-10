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
		boost     float64
		valueS    interface{}
		fuzziness int
		bound     *Bound
		tType     TermType
	}
	var testCases = []testCase{
		{
			name:      "TestTerm01",
			input:     `"dsada 78"`,
			want:      &Term{FuzzyTerm: &FuzzyTerm{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, ` `, `78`}}}},
			wantStr:   "\"dsada 78\"",
			boost:     1.0,
			valueS:    "dsada 78",
			fuzziness: 0,
			bound:     nil,
			tType:     PHRASE_TERM_TYPE,
		},
		{
			name:      "TestTerm02",
			input:     `"dsada 78"^08`,
			want:      &Term{FuzzyTerm: &FuzzyTerm{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, ` `, `78`}}, BoostSymbol: "^08"}},
			wantStr:   `"dsada 78"^08`,
			boost:     8.0,
			valueS:    "dsada 78",
			fuzziness: 0,
			bound:     nil,
			tType:     PHRASE_TERM_TYPE | BOOST_TERM_TYPE,
		},
		{
			name:      "TestTerm03",
			input:     `"dsada 78"~8`,
			want:      &Term{FuzzyTerm: &FuzzyTerm{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, ` `, `78`}}, FuzzySymbol: "~8"}},
			wantStr:   `"dsada 78"~8`,
			boost:     1.0,
			valueS:    "dsada 78",
			fuzziness: 8,
			bound:     nil,
			tType:     PHRASE_TERM_TYPE | FUZZY_TERM_TYPE,
		},
		{
			name:      "TestTerm03_01",
			input:     `"dsada 78"~8.1`,
			want:      &Term{FuzzyTerm: &FuzzyTerm{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, ` `, `78`}}, FuzzySymbol: "~8.1"}},
			wantStr:   `"dsada 78"~8.1`,
			boost:     1.0,
			valueS:    "dsada 78",
			fuzziness: 8,
			bound:     nil,
			tType:     PHRASE_TERM_TYPE | FUZZY_TERM_TYPE,
		},
		{
			name:      "TestTerm03_02",
			input:     `"dsada 78"~8.6`,
			want:      &Term{FuzzyTerm: &FuzzyTerm{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, ` `, `78`}}, FuzzySymbol: "~8.6"}},
			wantStr:   `"dsada 78"~8.6`,
			boost:     1.0,
			valueS:    "dsada 78",
			fuzziness: 9,
			bound:     nil,
			tType:     PHRASE_TERM_TYPE | FUZZY_TERM_TYPE,
		},
		{
			name:      "TestTerm05",
			input:     `/dsada 78/`,
			want:      &Term{RegexpTerm: &RegexpTerm{Chars: []string{`dsada`, ` `, `78`}}},
			wantStr:   `/dsada 78/`,
			boost:     1.0,
			valueS:    "dsada 78",
			fuzziness: 0,
			bound:     nil,
			tType:     REGEXP_TERM_TYPE,
		},
		{
			name:      "TestTerm06",
			input:     `\/dsada\/\ dasda80980?*`,
			want:      &Term{FuzzyTerm: &FuzzyTerm{SingleTerm: &SingleTerm{Begin: `\/`, Chars: []string{`dsada`, `\/\ `, `dasda`, `80980`, `?`, `*`}}}},
			wantStr:   `\/dsada\/\ dasda80980?*`,
			boost:     1.0,
			valueS:    `\/dsada\/\ dasda80980?*`,
			fuzziness: 0,
			bound:     nil,
			tType:     SINGLE_TERM_TYPE | WILDCARD_TERM_TYPE,
		},
		{
			name:      "TestTerm07",
			input:     `\/dsada\/\ dasda80980?*\^\^^08`,
			want:      &Term{FuzzyTerm: &FuzzyTerm{SingleTerm: &SingleTerm{Begin: `\/`, Chars: []string{`dsada`, `\/\ `, `dasda`, `80980`, `?`, `*`, `\^\^`}}, BoostSymbol: `^08`}},
			wantStr:   `\/dsada\/\ dasda80980?*\^\^^08`,
			boost:     8.0,
			valueS:    `\/dsada\/\ dasda80980?*\^\^`,
			fuzziness: 0,
			bound:     nil,
			tType:     SINGLE_TERM_TYPE | BOOST_TERM_TYPE | WILDCARD_TERM_TYPE,
		},
		{
			name:      "TestTerm08",
			input:     `\/dsada\/\ dasda80980?*\^\^~8`,
			want:      &Term{FuzzyTerm: &FuzzyTerm{SingleTerm: &SingleTerm{Begin: `\/`, Chars: []string{`dsada`, `\/\ `, `dasda`, `80980`, `?`, `*`, `\^\^`}}, FuzzySymbol: `~8`}},
			wantStr:   `\/dsada\/\ dasda80980?*\^\^~8`,
			boost:     1.0,
			valueS:    `\/dsada\/\ dasda80980?*\^\^`,
			fuzziness: 8,
			bound:     nil,
			tType:     SINGLE_TERM_TYPE | FUZZY_TERM_TYPE | WILDCARD_TERM_TYPE,
		},
		{
			name:  "TestTerm10",
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
			boost:   7.0,
			valueS: &Bound{
				LeftValue:    &RangeValue{SingleValue: []string{"1"}, flag: false},
				RightValue:   &RangeValue{SingleValue: []string{"2"}, flag: true},
				LeftInclude:  true,
				RightInclude: true,
			},
			fuzziness: 0,
			bound: &Bound{
				LeftValue:    &RangeValue{SingleValue: []string{"1"}, flag: false},
				RightValue:   &RangeValue{SingleValue: []string{"2"}, flag: true},
				LeftInclude:  true,
				RightInclude: true,
			},
			tType: RANGE_TERM_TYPE | BOOST_TERM_TYPE,
		},
		{
			name:  "TestTerm11",
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
			boost:   7.0,
			valueS: &Bound{
				LeftValue:    &RangeValue{SingleValue: []string{"1"}, flag: false},
				RightValue:   &RangeValue{SingleValue: []string{"2"}, flag: true},
				LeftInclude:  true,
				RightInclude: false,
			},
			fuzziness: 0,
			bound: &Bound{
				LeftValue:    &RangeValue{SingleValue: []string{"1"}, flag: false},
				RightValue:   &RangeValue{SingleValue: []string{"2"}, flag: true},
				LeftInclude:  true,
				RightInclude: false,
			},
			tType: RANGE_TERM_TYPE | BOOST_TERM_TYPE,
		},
		{
			name:  `TestTerm12`,
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
			boost:   7.0,
			valueS: &Bound{
				LeftValue:    &RangeValue{SingleValue: []string{"1"}, flag: false},
				RightValue:   &RangeValue{SingleValue: []string{"2"}, flag: true},
				LeftInclude:  false,
				RightInclude: false,
			},
			fuzziness: 0,
			bound: &Bound{
				LeftValue:    &RangeValue{SingleValue: []string{"1"}, flag: false},
				RightValue:   &RangeValue{SingleValue: []string{"2"}, flag: true},
				LeftInclude:  false,
				RightInclude: false,
			},
			tType: RANGE_TERM_TYPE | BOOST_TERM_TYPE,
		},
		{
			name:  `TestTerm13`,
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
			boost:   7.0,
			valueS: &Bound{
				LeftValue:    &RangeValue{SingleValue: []string{"1"}, flag: false},
				RightValue:   &RangeValue{SingleValue: []string{"2"}, flag: true},
				LeftInclude:  false,
				RightInclude: true,
			},
			fuzziness: 0,
			bound: &Bound{
				LeftValue:    &RangeValue{SingleValue: []string{"1"}, flag: false},
				RightValue:   &RangeValue{SingleValue: []string{"2"}, flag: true},
				LeftInclude:  false,
				RightInclude: true,
			},
			tType: RANGE_TERM_TYPE | BOOST_TERM_TYPE,
		},
		{
			name:  `TestTerm14`,
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
			boost:   7.0,
			valueS: &Bound{
				LeftValue:    &RangeValue{SingleValue: []string{"10"}, flag: false},
				RightValue:   &RangeValue{InfinityVal: "*", flag: true},
				LeftInclude:  true,
				RightInclude: false,
			},
			fuzziness: 0,
			bound: &Bound{
				LeftValue:    &RangeValue{SingleValue: []string{"10"}, flag: false},
				RightValue:   &RangeValue{InfinityVal: "*", flag: true},
				LeftInclude:  true,
				RightInclude: false,
			},
			tType: RANGE_TERM_TYPE | BOOST_TERM_TYPE,
		},
		{
			name:  `TestTerm15`,
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
			boost:   1.0,
			valueS: &Bound{
				LeftValue:    &RangeValue{InfinityVal: "*", flag: false},
				RightValue:   &RangeValue{SingleValue: []string{"2012", "-", "01", "-", "01"}, flag: true},
				LeftInclude:  false,
				RightInclude: false,
			},
			fuzziness: 0,
			bound: &Bound{
				LeftValue:    &RangeValue{InfinityVal: "*", flag: false},
				RightValue:   &RangeValue{SingleValue: []string{"2012", "-", "01", "-", "01"}, flag: true},
				LeftInclude:  false,
				RightInclude: false,
			},
			tType: RANGE_TERM_TYPE,
		},
		{
			name:  `TestTerm16`,
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
			boost:   1.0,
			valueS: &Bound{
				LeftValue:    &RangeValue{InfinityVal: "*", flag: false},
				RightValue:   &RangeValue{PhraseValue: []string{"2012", "-", "01", "-", "01", " ", "09", ":", "08", ":", "16"}, flag: true},
				LeftInclude:  false,
				RightInclude: false,
			},
			fuzziness: 0,
			bound: &Bound{
				LeftValue:    &RangeValue{InfinityVal: "*", flag: false},
				RightValue:   &RangeValue{PhraseValue: []string{"2012", "-", "01", "-", "01", " ", "09", ":", "08", ":", "16"}, flag: true},
				LeftInclude:  false,
				RightInclude: false,
			},
			tType: RANGE_TERM_TYPE,
		},
		{
			name:  "TestTerm17",
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
			boost:   1.0,
			valueS: &Bound{
				LeftValue:    &RangeValue{InfinityVal: "*", flag: false},
				RightValue:   &RangeValue{PhraseValue: []string{`dsada`, ` `, `78`}, flag: true},
				LeftInclude:  false,
				RightInclude: true,
			},
			fuzziness: 0,
			bound: &Bound{
				LeftValue:    &RangeValue{InfinityVal: "*", flag: false},
				RightValue:   &RangeValue{PhraseValue: []string{`dsada`, ` `, `78`}, flag: true},
				LeftInclude:  false,
				RightInclude: true,
			},
			tType: RANGE_TERM_TYPE,
		},
		{
			name:  "TestTerm18",
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
			boost:   8.0,
			valueS: &Bound{
				LeftValue:    &RangeValue{InfinityVal: "*", flag: false},
				RightValue:   &RangeValue{PhraseValue: []string{`dsada`, ` `, `78`}, flag: true},
				LeftInclude:  false,
				RightInclude: false,
			},
			fuzziness: 0,
			bound: &Bound{
				LeftValue:    &RangeValue{InfinityVal: "*", flag: false},
				RightValue:   &RangeValue{PhraseValue: []string{`dsada`, ` `, `78`}, flag: true},
				LeftInclude:  false,
				RightInclude: false,
			},
			tType: RANGE_TERM_TYPE | BOOST_TERM_TYPE,
		},
		{
			name:  "TestTerm20",
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
			boost:   80.0,
			valueS: &Bound{
				LeftValue:    &RangeValue{PhraseValue: []string{`dsada`, ` `, `78`}, flag: false},
				RightValue:   &RangeValue{InfinityVal: "*", flag: true},
				LeftInclude:  false,
				RightInclude: false,
			},
			fuzziness: 0,
			bound: &Bound{
				LeftValue:    &RangeValue{PhraseValue: []string{`dsada`, ` `, `78`}, flag: false},
				RightValue:   &RangeValue{InfinityVal: "*", flag: true},
				LeftInclude:  false,
				RightInclude: false,
			},
			tType: RANGE_TERM_TYPE | BOOST_TERM_TYPE,
		},
		{
			name:  "TestTerm21",
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
			boost:   1.0,
			valueS: &Bound{
				LeftValue:    &RangeValue{InfinityVal: "*", flag: false},
				RightValue:   &RangeValue{SingleValue: []string{`dsada`, `\ `, `78`}, flag: true},
				LeftInclude:  false,
				RightInclude: true,
			},
			fuzziness: 0,
			bound: &Bound{
				LeftValue:    &RangeValue{InfinityVal: "*", flag: false},
				RightValue:   &RangeValue{SingleValue: []string{`dsada`, `\ `, `78`}, flag: true},
				LeftInclude:  false,
				RightInclude: true,
			},
			tType: RANGE_TERM_TYPE,
		},
		{
			name:  "TestTerm22",
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
			boost:   8.0,
			valueS: &Bound{
				LeftValue:    &RangeValue{InfinityVal: "*", flag: false},
				RightValue:   &RangeValue{SingleValue: []string{`dsada`, `\ `, `78`}, flag: true},
				LeftInclude:  false,
				RightInclude: false,
			},
			fuzziness: 0,
			bound: &Bound{
				LeftValue:    &RangeValue{InfinityVal: "*", flag: false},
				RightValue:   &RangeValue{SingleValue: []string{`dsada`, `\ `, `78`}, flag: true},
				LeftInclude:  false,
				RightInclude: false,
			},
			tType: RANGE_TERM_TYPE | BOOST_TERM_TYPE,
		},
		{
			name:  "TestTerm24",
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
			boost:   80.0,
			valueS: &Bound{
				LeftValue:    &RangeValue{SingleValue: []string{`dsada`, `\ `, `78`}, flag: false},
				RightValue:   &RangeValue{InfinityVal: "*", flag: true},
				LeftInclude:  false,
				RightInclude: false,
			},
			fuzziness: 0,
			bound: &Bound{
				LeftValue:    &RangeValue{SingleValue: []string{`dsada`, `\ `, `78`}, flag: false},
				RightValue:   &RangeValue{InfinityVal: "*", flag: true},
				LeftInclude:  false,
				RightInclude: false,
			},
			tType: RANGE_TERM_TYPE | BOOST_TERM_TYPE,
		},
		{
			name:      "TestTerm25",
			input:     `/\d+\d+\.\d+.+/`,
			want:      &Term{RegexpTerm: &RegexpTerm{Chars: []string{`\`, `d`, `+`, `\`, `d`, `+`, `\`, `.`, `\`, `d`, `+`, `.`, `+`}}},
			wantStr:   `/\d+\d+\.\d+.+/`,
			boost:     1.0,
			valueS:    `\d+\d+\.\d+.+`,
			fuzziness: 0,
			bound:     nil,
			tType:     REGEXP_TERM_TYPE,
		},
		{
			name:  "TestTerm26",
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
			boost:     1.0,
			valueS:    `foo OR bar`,
			fuzziness: 0,
			bound:     nil,
			tType:     GROUP_TERM_TYPE,
		},
		{
			name:  "TestTerm27",
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
			boost:     7.0,
			valueS:    `foo OR bar OR [ 1 TO 2 ] AND { 10 TO * } AND { * TO 20 } AND "abc"`,
			fuzziness: 0,
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
	assert.Equal(t, 0.0, out.Boost())
	assert.Equal(t, 0, out.Fuzziness())
	assert.Equal(t, UNKNOWN_TERM_TYPE, out.GetTermType())
	_, err := out.Value(func(s string) (interface{}, error) { return s, nil })
	assert.Equal(t, ErrEmptyTerm, err)
	out = &Term{}
	assert.Equal(t, "", out.String())
	assert.Equal(t, 0.0, out.Boost())
	assert.Equal(t, 0, out.Fuzziness())
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
			name:  "TestRegexpTerm01",
			input: `12313\+90`,
			want:  false,
		},
		{
			name:  "TestRegexpTerm02",
			input: `/[1-9]+\.\d+/`,
			want:  true,
		},
		{
			name:  "TestRegexpTerm03",
			input: `"dsad 7089"`,
			want:  false,
		},
		{
			name:  "TestRegexpTerm04",
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
			name:  "TestWildcard01",
			input: `12313?`,
			want:  true,
		},
		{
			name:  "TestWildcard02",
			input: `12313\?`,
			want:  false,
		},
		{
			name:  "TestWildcard03",
			input: `12313*`,
			want:  true,
		},
		{
			name:  "TestWildcard04",
			input: `12313\*`,
			want:  false,
		},
		{
			name:  "TestWildcard05",
			input: `/[1-9]+\.\d+/`,
			want:  false,
		},
		{
			name:  "TestWildcard06",
			input: `"dsad?\? 7089*"`,
			want:  true,
		},
		{
			name:  "TestWildcard07",
			input: `"dsadad 789"`,
			want:  false,
		},
		{
			name:  "TestWildcard08",
			input: `[1 TO 2]`,
			want:  false,
		},
		{
			name:  "TestWildcard09",
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
			name:  "TestRangeTerm01",
			input: `12313\+90`,
			want:  false,
		},
		{
			name:  "TestRangeTerm02",
			input: `/[1-9]+\.\d+/`,
			want:  false,
		},
		{
			name:  "TestRangeTerm03",
			input: `"dsad 7089"`,
			want:  false,
		},
		{
			name:  "TestRangeTerm04",
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
		want  int
	}

	var testCases = []testCase{
		{
			name:  "TestFuzzines01",
			input: `12313\+90`,
			want:  0,
		},
		{
			name:  "TestFuzzines02",
			input: `/[1-9]+\.\d+/`,
			want:  0,
		},
		{
			name:  "TestFuzzines03",
			input: `"dsad 7089"`,
			want:  0,
		},
		{
			name:  "TestFuzzines04",
			input: `[1 TO 454 ]`,
			want:  0,
		},
		{
			name:  "TestFuzzines05",
			input: `12313\+90~3`,
			want:  3,
		},
		{
			name:  "TestFuzzines06",
			input: `"dsad 7089"~3`,
			want:  3,
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
		want  float64
	}

	var testCases = []testCase{
		{
			name:  "TestBoost01",
			input: `12313\+90`,
			want:  1.0,
		},
		{
			name:  "TestBoost02",
			input: `/[1-9]+\.\d+/`,
			want:  1.0,
		},
		{
			name:  "TestBoost03",
			input: `"dsad 7089"`,
			want:  1.0,
		},
		{
			name:  "TestBoost04",
			input: `[1 TO 454 ]`,
			want:  1.0,
		},
		{
			name:  "TestBoost05",
			input: `12313\+90^1.2`,
			want:  1.2,
		},
		{
			name:  "TestBoost06",
			input: `12313\+90^0.2`,
			want:  0.2,
		},
		{
			name:  "TestBoost07",
			input: `"dsad 7089"^3.8`,
			want:  3.8,
		},
		{
			name:  "TestBoost08",
			input: `"dsad 7089"^0.8`,
			want:  0.8,
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
