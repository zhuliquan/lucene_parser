package term

import (
	"math"
	"reflect"
	"testing"

	"github.com/alecthomas/participle"
	"github.com/zhuliquan/lucene_parser/operator"
	"github.com/zhuliquan/lucene_parser/token"
)

func TestTerm(t *testing.T) {
	var termParser = participle.MustBuild(
		&Term{},
		participle.Lexer(token.Lexer),
	)

	type testCase struct {
		name  string
		input string
		want  *Term
	}
	var testCases = []testCase{
		{
			name:  "TestTerm01",
			input: `"dsada 78"`,
			want:  &Term{FuzzyTerm: &FuzzyTerm{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, ` `, `78`}}}},
		},
		{
			name:  "TestTerm02",
			input: `"dsada 78"^08`,
			want:  &Term{FuzzyTerm: &FuzzyTerm{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, ` `, `78`}}, BoostSymbol: "^08"}},
		},
		{
			name:  "TestTerm03",
			input: `"dsada 78"~8`,
			want:  &Term{FuzzyTerm: &FuzzyTerm{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, ` `, `78`}}, FuzzySymbol: "~8"}},
		},
		{
			name:  "TestTerm05",
			input: `/dsada 78/`,
			want:  &Term{RegexpTerm: &RegexpTerm{Chars: []string{`dsada`, ` `, `78`}}},
		},
		{
			name:  "TestTerm06",
			input: `\/dsada\/\ dasda80980?*`,
			want:  &Term{FuzzyTerm: &FuzzyTerm{SingleTerm: &SingleTerm{Begin: `\/dsada\/\ dasda`, Chars: []string{`80980`, `?`, `*`}}}},
		},
		{
			name:  "TestTerm07",
			input: `\/dsada\/\ dasda80980?*\^\^^08`,
			want:  &Term{FuzzyTerm: &FuzzyTerm{SingleTerm: &SingleTerm{Begin: `\/dsada\/\ dasda`, Chars: []string{`80980`, `?`, `*`, `\^\^`}}, BoostSymbol: `^08`}},
		},
		{
			name:  "TestTerm08",
			input: `\/dsada\/\ dasda80980?*\^\^~8`,
			want:  &Term{FuzzyTerm: &FuzzyTerm{SingleTerm: &SingleTerm{Begin: `\/dsada\/\ dasda`, Chars: []string{`80980`, `?`, `*`, `\^\^`}}, FuzzySymbol: `~8`}},
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
		},
		{
			name:  "TestTerm21",
			input: `<=dsada\ 78`,
			want: &Term{
				RangeTerm: &RangeTerm{
					SRangeTerm: &SRangeTerm{
						Symbol: "<=",
						Value:  &RangeValue{SingleValue: []string{`dsada\ `, `78`}},
					},
				},
			},
		},
		{
			name:  "TestTerm22",
			input: `<dsada\ 78^08`,
			want: &Term{
				RangeTerm: &RangeTerm{
					SRangeTerm: &SRangeTerm{
						Symbol: "<",
						Value:  &RangeValue{SingleValue: []string{`dsada\ `, `78`}},
					},
					BoostSymbol: "^08",
				},
			},
		},
		{
			name:  "TestTerm24",
			input: `>dsada\ 78^080`,
			want: &Term{
				RangeTerm: &RangeTerm{
					SRangeTerm: &SRangeTerm{
						Symbol: ">",
						Value:  &RangeValue{SingleValue: []string{`dsada\ `, `78`}},
					},
					BoostSymbol: "^080",
				},
			},
		},
		{
			name:  "TestTerm25",
			input: `/\d+\d+\.\d+.+/`,
			want:  &Term{RegexpTerm: &RegexpTerm{Chars: []string{`\`, `d`, `+`, `\`, `d`, `+`, `\`, `.`, `\`, `d`, `+`, `.`, `+`}}},
		},
		{
			name:  "TestTerm26",
			input: `(foo OR bar)`,
			want: &Term{TermGroup: &TermGroup{
				LogicTermGroup: &LogicTermGroup{
					OrTermGroup: &OrTermGroup{
						AndTermGroup: &AndTermGroup{
							TermGroupElem: &TermGroupElem{SingleTerm: &SingleTerm{Begin: "foo"}},
						},
					},
					OSTermGroup: []*OSTermGroup{
						{
							OrSymbol: &operator.OrSymbol{Symbol: "OR"},
							OrTermGroup: &OrTermGroup{
								AndTermGroup: &AndTermGroup{
									TermGroupElem: &TermGroupElem{SingleTerm: &SingleTerm{Begin: "bar"}},
								},
							},
						},
					},
				},
			}},
		},
		{
			name:  "TestTerm27",
			input: `(foo OR bar or [1 TO 2])^7`,
			want: &Term{TermGroup: &TermGroup{
				LogicTermGroup: &LogicTermGroup{
					OrTermGroup: &OrTermGroup{
						AndTermGroup: &AndTermGroup{
							TermGroupElem: &TermGroupElem{SingleTerm: &SingleTerm{Begin: "foo"}},
						},
					},
					OSTermGroup: []*OSTermGroup{
						{
							OrSymbol: &operator.OrSymbol{Symbol: "OR"},
							OrTermGroup: &OrTermGroup{
								AndTermGroup: &AndTermGroup{
									TermGroupElem: &TermGroupElem{SingleTerm: &SingleTerm{Begin: "bar"}},
								},
							},
						},
						{
							OrSymbol: &operator.OrSymbol{Symbol: "or"},
							OrTermGroup: &OrTermGroup{
								AndTermGroup: &AndTermGroup{
									TermGroupElem: &TermGroupElem{DRangeTerm: &DRangeTerm{
										LBRACKET: "[",
										LValue:   &RangeValue{SingleValue: []string{"1"}},
										RValue:   &RangeValue{SingleValue: []string{"2"}},
										RBRACKET: "]",
									}},
								},
							},
						},
					},
				},
				BoostSymbol: "^7",
			}},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var out = &Term{}
			if err := termParser.ParseString(tt.input, out); err != nil {
				t.Errorf("failed to parse input: %s, err: %+v", tt.input, err)
			} else if !reflect.DeepEqual(tt.want, out) {
				t.Errorf("termParser.ParseString( %s ) = %+v, want: %+v", tt.input, out, tt.want)
			}
		})
	}
}

func TestTerm_isRegexp(t *testing.T) {
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
			if err := termParser.ParseString(tt.input, out); err != nil {
				t.Errorf("failed to parse input: %s, err: %+v", tt.input, err)
			} else if (out.GetTermType()&REGEXP_TERM_TYPE == REGEXP_TERM_TYPE) != tt.want {
				t.Errorf("isRegexp() = %+v, want: %+v", (out.GetTermType()&REGEXP_TERM_TYPE == REGEXP_TERM_TYPE), tt.want)
			}
		})
	}

}

func TestTerm_isWildcard(t *testing.T) {

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
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var out = &Term{}
			if err := termParser.ParseString(tt.input, out); err != nil {
				t.Errorf("failed to parse input: %s, err: %+v", tt.input, err)
			} else if (out.GetTermType()&WILDCARD_TERM_TYPE == WILDCARD_TERM_TYPE) != tt.want {
				t.Errorf("haveWildcard() = %+v, want: %+v", (out.GetTermType()&WILDCARD_TERM_TYPE == WILDCARD_TERM_TYPE), tt.want)
			}
		})
	}
}

func TestTerm_isRange(t *testing.T) {
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
			if err := termParser.ParseString(tt.input, out); err != nil {
				t.Errorf("failed to parse input: %s, err: %+v", tt.input, err)
			} else if (out.GetTermType()&RANGE_TERM_TYPE == RANGE_TERM_TYPE) != tt.want {
				t.Errorf("isRange() = %+v, want: %+v", (out.GetTermType()&RANGE_TERM_TYPE == RANGE_TERM_TYPE), tt.want)
			}
		})
	}
}

func TestTerm_fuzziness(t *testing.T) {

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
			if err := termParser.ParseString(tt.input, out); err != nil {
				t.Errorf("failed to parse input: %s, err: %+v", tt.input, err)
			} else if out.Fuzziness() != tt.want {
				t.Errorf("fuzziness() = %+v, want: %+v", out.Fuzziness(), tt.want)
			}
		})
	}

}

func TestTerm_boost(t *testing.T) {

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
			if err := termParser.ParseString(tt.input, out); err != nil {
				t.Errorf("failed to parse input: %s, err: %+v", tt.input, err)
			} else if math.Abs(out.Boost()-tt.want) > 1E-8 {
				t.Errorf("boost() = %+v, want: %+v", out.Boost(), tt.want)
			}
		})
	}

}
