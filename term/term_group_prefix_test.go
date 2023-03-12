//go:build prefix_group
// +build prefix_group

package term

import (
	"reflect"
	"testing"

	"github.com/alecthomas/participle"
	"github.com/stretchr/testify/assert"
	op "github.com/zhuliquan/lucene_parser/operator"
	"github.com/zhuliquan/lucene_parser/token"
)

func TestPrefixTerm(t *testing.T) {
	var termParser = participle.MustBuild(
		&PrefixTerm{},
		participle.Lexer(token.Lexer),
	)

	type testCase struct {
		name  string
		input string
		want  *PrefixTerm
		oType op.PrefixOPType
	}
	var testCases = []testCase{
		{
			name:  "TestPrefixTerm01",
			input: `"dsada 78"`,
			want:  &PrefixTerm{Elem: &TermGroupElem{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, ` `, `78`}}}},
			oType: op.SHOULD_PREFIX_TYPE,
		},
		{
			name:  "TestPrefixTerm02",
			input: `+"dsada 78"`,
			want:  &PrefixTerm{Symbol: "+", Elem: &TermGroupElem{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, ` `, `78`}}}},
			oType: op.MUST_PREFIX_TYPE,
		},
		{
			name:  "TestPrefixTerm03",
			input: `-"dsada 78"`,
			want:  &PrefixTerm{Symbol: "-", Elem: &TermGroupElem{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, ` `, `78`}}}},
			oType: op.MUST_NOT_PREFIX_TYPE,
		},
		{
			name:  "TestPrefixTerm04",
			input: `\+\/dsada\/\ dasda80980?*`,
			want:  &PrefixTerm{Elem: &TermGroupElem{SingleTerm: &SingleTerm{Begin: `\+`, Chars: []string{`\/`, `dsada`, `\/\ `, `dasda`, `80980`, `?`, `*`}}}},
			oType: op.SHOULD_PREFIX_TYPE,
		},
		{
			name:  "TestPrefixTerm05",
			input: `+\/dsada\/\ dasda80980?*`,
			want: &PrefixTerm{Symbol: "+", Elem: &TermGroupElem{
				SingleTerm: &SingleTerm{Begin: `\/`, Chars: []string{`dsada`, `\/\ `, `dasda`, `80980`, `?`, `*`}},
			}},
			oType: op.MUST_PREFIX_TYPE,
		},
		{
			name:  "TestPrefixTerm06",
			input: `-\-\/dsada\/\ dasda80980?*`,
			want: &PrefixTerm{Symbol: "-", Elem: &TermGroupElem{
				SingleTerm: &SingleTerm{Begin: `\-`, Chars: []string{`\/`, `dsada`, `\/\ `, `dasda`, `80980`, `?`, `*`}},
			}},
			oType: op.MUST_NOT_PREFIX_TYPE,
		},
		{
			name:  "TestPrefixTerm07",
			input: `->890`,
			want: &PrefixTerm{Symbol: "-", Elem: &TermGroupElem{
				SRangeTerm: &SRangeTerm{Symbol: ">", Value: &RangeValue{SingleValue: []string{`890`}}},
			}},
			oType: op.MUST_NOT_PREFIX_TYPE,
		},
		{
			name:  "TestPrefixTerm08",
			input: `>890`,
			want:  &PrefixTerm{Elem: &TermGroupElem{SRangeTerm: &SRangeTerm{Symbol: ">", Value: &RangeValue{SingleValue: []string{`890`}}}}},
			oType: op.SHOULD_PREFIX_TYPE,
		},
		{
			name:  "TestPrefixTerm09",
			input: `+>=890`,
			want:  &PrefixTerm{Symbol: "+", Elem: &TermGroupElem{SRangeTerm: &SRangeTerm{Symbol: ">=", Value: &RangeValue{SingleValue: []string{`890`}}}}},
			oType: op.MUST_PREFIX_TYPE,
		},
		{
			name:  "TestPrefixTerm10",
			input: `+[1 TO 2]`,
			want: &PrefixTerm{Symbol: "+", Elem: &TermGroupElem{DRangeTerm: &DRangeTerm{
				LBRACKET: "[", LValue: &RangeValue{SingleValue: []string{"1"}},
				RValue: &RangeValue{SingleValue: []string{"2"}}, RBRACKET: "]",
			}}},
			oType: op.MUST_PREFIX_TYPE,
		},
		{
			name:  "TestPrefixTerm11",
			input: `-[1 TO 2]`,
			want: &PrefixTerm{Symbol: "-", Elem: &TermGroupElem{DRangeTerm: &DRangeTerm{
				LBRACKET: "[", LValue: &RangeValue{SingleValue: []string{"1"}},
				RValue: &RangeValue{SingleValue: []string{"2"}}, RBRACKET: "]",
			}}},
			oType: op.MUST_NOT_PREFIX_TYPE,
		},
		{
			name:  "TestPrefixTerm12",
			input: `[1 TO 2]`,
			want: &PrefixTerm{Elem: &TermGroupElem{DRangeTerm: &DRangeTerm{
				LBRACKET: "[", LValue: &RangeValue{SingleValue: []string{"1"}},
				RValue: &RangeValue{SingleValue: []string{"2"}}, RBRACKET: "]",
			}}},
			oType: op.SHOULD_PREFIX_TYPE,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var out = &PrefixTerm{}
			if err := termParser.ParseString(tt.input, out); err != nil {
				t.Errorf("failed to parse input: %s, err: %+v", tt.input, err)
			} else if !reflect.DeepEqual(tt.want, out) {
				t.Errorf("termParser.ParseString( %s ) = %+v, want: %+v", tt.input, out, tt.want)
			} else if out.GetPrefixType() != tt.oType {
				t.Errorf("expect get type: %+v, but get type: %+v", tt.oType, out.GetPrefixType())
			}
		})
	}
}

func TestWPrefixTerm(t *testing.T) {
	var termParser = participle.MustBuild(
		&WPrefixTerm{},
		participle.Lexer(token.Lexer),
	)

	type testCase struct {
		name  string
		input string
		want  *WPrefixTerm
		oType op.PrefixOPType
	}
	var testCases = []testCase{
		{
			name:  "TestWPrefixTerm01",
			input: `  "dsada 78"`,
			want:  &WPrefixTerm{Elem: &TermGroupElem{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, ` `, `78`}}}},
			oType: op.SHOULD_PREFIX_TYPE,
		},
		{
			name:  "TestWPrefixTerm02",
			input: `   +"dsada 78"`,
			want:  &WPrefixTerm{Symbol: "+", Elem: &TermGroupElem{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, ` `, `78`}}}},
			oType: op.MUST_PREFIX_TYPE,
		},
		{
			name:  "TestWPrefixTerm03",
			input: `  -"dsada 78"`,
			want:  &WPrefixTerm{Symbol: "-", Elem: &TermGroupElem{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, ` `, `78`}}}},
			oType: op.MUST_NOT_PREFIX_TYPE,
		},
		{
			name:  "TestWPrefixTerm04",
			input: `  \+\/dsada\/\ dasda80980?*`,
			want:  &WPrefixTerm{Elem: &TermGroupElem{SingleTerm: &SingleTerm{Begin: `\+`, Chars: []string{`\/`, `dsada`, `\/\ `, `dasda`, `80980`, `?`, `*`}}}},
			oType: op.SHOULD_PREFIX_TYPE,
		},
		{
			name:  "TestWPrefixTerm05",
			input: `  +\/dsada\/\ dasda80980?*`,
			want:  &WPrefixTerm{Symbol: "+", Elem: &TermGroupElem{SingleTerm: &SingleTerm{Begin: `\/`, Chars: []string{`dsada`, `\/\ `, `dasda`, `80980`, `?`, `*`}}}},
			oType: op.MUST_PREFIX_TYPE,
		},
		{
			name:  "TestWPrefixTerm06",
			input: `  -\-\/dsada\/\ dasda80980?*`,
			want:  &WPrefixTerm{Symbol: "-", Elem: &TermGroupElem{SingleTerm: &SingleTerm{Begin: `\-`, Chars: []string{`\/`, `dsada`, `\/\ `, `dasda`, `80980`, `?`, `*`}}}},
			oType: op.MUST_NOT_PREFIX_TYPE,
		},
		{
			name:  "TestWPrefixTerm07",
			input: `  ->890`,
			want:  &WPrefixTerm{Symbol: "-", Elem: &TermGroupElem{SRangeTerm: &SRangeTerm{Symbol: ">", Value: &RangeValue{SingleValue: []string{`890`}}}}},
			oType: op.MUST_NOT_PREFIX_TYPE,
		},
		{
			name:  "TestWPrefixTerm08",
			input: `  >890`,
			want:  &WPrefixTerm{Elem: &TermGroupElem{SRangeTerm: &SRangeTerm{Symbol: ">", Value: &RangeValue{SingleValue: []string{`890`}}}}},
			oType: op.SHOULD_PREFIX_TYPE,
		},
		{
			name:  "TestWPrefixTerm09",
			input: `  +>=890`,
			want:  &WPrefixTerm{Symbol: "+", Elem: &TermGroupElem{SRangeTerm: &SRangeTerm{Symbol: ">=", Value: &RangeValue{SingleValue: []string{`890`}}}}},
			oType: op.MUST_PREFIX_TYPE,
		},
		{
			name:  "TestWPrefixTerm10",
			input: `   +[1 TO 2]`,
			want: &WPrefixTerm{Symbol: "+", Elem: &TermGroupElem{DRangeTerm: &DRangeTerm{
				LBRACKET: "[", LValue: &RangeValue{SingleValue: []string{"1"}},
				RValue: &RangeValue{SingleValue: []string{"2"}}, RBRACKET: "]",
			}}},
			oType: op.MUST_PREFIX_TYPE,
		},
		{
			name:  "TestWPrefixTerm11",
			input: `  -[1 TO 2]`,
			want: &WPrefixTerm{Symbol: "-", Elem: &TermGroupElem{DRangeTerm: &DRangeTerm{
				LBRACKET: "[", LValue: &RangeValue{SingleValue: []string{"1"}},
				RValue: &RangeValue{SingleValue: []string{"2"}}, RBRACKET: "]",
			}}},
			oType: op.MUST_NOT_PREFIX_TYPE,
		},
		{
			name:  "TestWPrefixTerm12",
			input: `  [1 TO 2]`,
			want: &WPrefixTerm{Elem: &TermGroupElem{DRangeTerm: &DRangeTerm{
				LBRACKET: "[", LValue: &RangeValue{SingleValue: []string{"1"}},
				RValue: &RangeValue{SingleValue: []string{"2"}}, RBRACKET: "]",
			}}},
			oType: op.SHOULD_PREFIX_TYPE,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var out = &WPrefixTerm{}
			if err := termParser.ParseString(tt.input, out); err != nil {
				t.Errorf("failed to parse input: %s, err: %+v", tt.input, err)
			} else if !reflect.DeepEqual(tt.want, out) {
				t.Errorf("termParser.ParseString( %s ) = %+v, want: %+v", tt.input, out, tt.want)
			} else if out.GetPrefixType() != tt.oType {
				t.Errorf("expect get type: %+v, but get type: %+v", tt.oType, out.GetPrefixType())
			}
		})
	}
}

func TestPrefixTermGroup(t *testing.T) {
	var termParser = participle.MustBuild(
		&PrefixTermGroup{},
		participle.Lexer(token.Lexer),
	)

	type testCase struct {
		name  string
		input string
		want  *PrefixTermGroup
	}
	var testCases = []testCase{
		{
			name:  "TestPrefixTermGroup01",
			input: `8908  "dsada 78" +"89080  xxx" -"xx yyyy" +\+dsada\ 7897 -\-\-dsada\-7897`,
			want: &PrefixTermGroup{
				PrefixTerm: &PrefixTerm{Elem: &TermGroupElem{SingleTerm: &SingleTerm{Begin: "8908"}}},
				PrefixTerms: []*WPrefixTerm{
					{Elem: &TermGroupElem{PhraseTerm: &PhraseTerm{Value: []string{`dsada`, ` `, `78`}}}},
					{Symbol: "+", Elem: &TermGroupElem{PhraseTerm: &PhraseTerm{Value: []string{`89080`, ` `, `xxx`}}}},
					{Symbol: "-", Elem: &TermGroupElem{PhraseTerm: &PhraseTerm{Value: []string{`xx`, ` `, `yyyy`}}}},
					{Symbol: "+", Elem: &TermGroupElem{SingleTerm: &SingleTerm{Value: []string{`\+`, `dsada`, `\ `, `7897`}}}},
					{Symbol: "-", Elem: &TermGroupElem{SingleTerm: &SingleTerm{Value: []string{`\-\-`, `dsada`, `\-`, `7897`}}}},
				},
			},
		},
		{
			name:  "TestPrefixTermGroup02",
			input: `8908`,
			want: &PrefixTermGroup{
				PrefixTerm: &PrefixTerm{Elem: &TermGroupElem{SingleTerm: &SingleTerm{Begin: "8908"}}},
			},
		},
		{
			name:  "TestPrefixTermGroup03",
			input: `8908 [ -1 TO 3]`,
			want: &PrefixTermGroup{
				PrefixTerm: &PrefixTerm{Elem: &TermGroupElem{SingleTerm: &SingleTerm{Begin: "8908"}}},
				PrefixTerms: []*WPrefixTerm{
					{
						Elem: &TermGroupElem{DRangeTerm: &DRangeTerm{
							LBRACKET: "[", LValue: &RangeValue{SingleValue: []string{"-", "1"}},
							RValue: &RangeValue{SingleValue: []string{"3"}}, RBRACKET: "]",
						}},
					},
				},
			},
		},
		{
			name:  "TestPrefixTermGroup04",
			input: `+>2021-11-04 +<2021-11-11`,
			want: &PrefixTermGroup{
				PrefixTerm: &PrefixTerm{Symbol: "+", Elem: &TermGroupElem{
					SRangeTerm: &SRangeTerm{
						Symbol: ">",
						Value:  &RangeValue{SingleValue: []string{`2021`, "-", "11", "-", "04"}},
					},
				}},
				PrefixTerms: []*WPrefixTerm{
					{Symbol: "+", Elem: &TermGroupElem{
						SRangeTerm: &SRangeTerm{
							Symbol: "<",
							Value:  &RangeValue{SingleValue: []string{`2021`, "-", "11", "-", "11"}},
						},
					}},
				},
			},
		},
		{
			name:  "TestPrefixTermGroup05",
			input: `[-1 TO 3]`,
			want: &PrefixTermGroup{
				PrefixTerm: &PrefixTerm{Elem: &TermGroupElem{
					DRangeTerm: &DRangeTerm{
						LBRACKET: "[", LValue: &RangeValue{SingleValue: []string{"-", "1"}},
						RValue: &RangeValue{SingleValue: []string{"3"}}, RBRACKET: "]",
					}}},
			},
		},
		{
			name:  "TestPrefixTermGroup06",
			input: `[-1 TO 3] [1 TO 2] +[5 TO 10}  -{8 TO 90]`,
			want: &PrefixTermGroup{
				PrefixTerm: &PrefixTerm{Elem: &TermGroupElem{
					DRangeTerm: &DRangeTerm{
						LBRACKET: "[", LValue: &RangeValue{SingleValue: []string{"-", "1"}},
						RValue: &RangeValue{SingleValue: []string{"3"}}, RBRACKET: "]",
					}}},
				PrefixTerms: []*WPrefixTerm{
					{Elem: &TermGroupElem{
						DRangeTerm: &DRangeTerm{
							LBRACKET: "[", LValue: &RangeValue{SingleValue: []string{"1"}},
							RValue: &RangeValue{SingleValue: []string{"2"}}, RBRACKET: "]",
						},
					}},
					{Symbol: "+", Elem: &TermGroupElem{
						DRangeTerm: &DRangeTerm{
							LBRACKET: "[", LValue: &RangeValue{SingleValue: []string{"5"}},
							RValue: &RangeValue{SingleValue: []string{"10"}}, RBRACKET: "}",
						},
					}},
					{Symbol: "-", Elem: &TermGroupElem{
						DRangeTerm: &DRangeTerm{
							LBRACKET: "{", LValue: &RangeValue{SingleValue: []string{"8"}},
							RValue: &RangeValue{SingleValue: []string{"90"}}, RBRACKET: "]",
						},
					}},
				},
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var out = &PrefixTermGroup{}
			err := termParser.ParseString(tt.input, out)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, out)
		})
	}
}

func TestTermGroup(t *testing.T) {
	var termParser = participle.MustBuild(
		&TermGroup{},
		participle.Lexer(token.Lexer),
	)

	type testCase struct {
		name  string
		input string
		want  *TermGroup
		boost float64
	}
	var testCases = []testCase{
		{
			name:  "TestTermGroup01",
			input: `( 8908  "dsada 78" +"89080  xxx" -"xx yyyy" +\+dsada\ 7897 -\-\-dsada\-7897)^1.8`,
			want: &TermGroup{
				PrefixTermGroup: &PrefixTermGroup{
					PrefixTerm: &PrefixTerm{Elem: &TermGroupElem{SingleTerm: &SingleTerm{Begin: "8908"}}},
					PrefixTerms: []*WPrefixTerm{
						{Elem: &TermGroupElem{PhraseTerm: &PhraseTerm{Chars: []string{`dsada`, ` `, `78`}}}},
						{Symbol: "+", Elem: &TermGroupElem{PhraseTerm: &PhraseTerm{Value: []string{`89080`, ` `, `xxx`}}}},
						{Symbol: "-", Elem: &TermGroupElem{PhraseTerm: &PhraseTerm{Value: []string{`xx`, ` `, `yyyy`}}}},
						{Symbol: "+", Elem: &TermGroupElem{SingleTerm: &SingleTerm{Value: []string{`\+`, `dsada`, `\ `, `7897`}}}},
						{Symbol: "-", Elem: &TermGroupElem{SingleTerm: &SingleTerm{Value: []string{`\-\-`, `dsada`, `\-`, `7897`}}}},
					},
				},
				BoostSymbol: "^1.8",
			},
			boost: 1.8,
		},
		{
			name:  "TestTermGroup02",
			input: `( 8908)^1.8`,
			want: &TermGroup{
				PrefixTermGroup: &PrefixTermGroup{
					PrefixTerm: &PrefixTerm{Elem: &TermGroupElem{SingleTerm: &SingleTerm{Begin: "8908"}}},
				},
				BoostSymbol: "^1.8",
			},
			boost: 1.8,
		},
		{
			name:  "TestTermGroup03",
			input: `(8908 [ -1 TO 3])^1.9`,
			want: &TermGroup{
				PrefixTermGroup: &PrefixTermGroup{
					PrefixTerm: &PrefixTerm{Elem: &TermGroupElem{SingleTerm: &SingleTerm{Begin: "8908"}}},
					PrefixTerms: []*WPrefixTerm{
						{
							Elem: &TermGroupElem{DRangeTerm: &DRangeTerm{
								LBRACKET: "[", LValue: &RangeValue{SingleValue: []string{"-", "1"}},
								RValue: &RangeValue{SingleValue: []string{"3"}}, RBRACKET: "]",
							}},
						},
					},
				},
				BoostSymbol: "^1.9",
			},
			boost: 1.9,
		},
		{
			name:  "TestTermGroup04",
			input: `(+>2021-11-04 +<2021-11-11)^1.9`,
			want: &TermGroup{
				PrefixTermGroup: &PrefixTermGroup{
					PrefixTerm: &PrefixTerm{Symbol: "+", Elem: &TermGroupElem{
						SRangeTerm: &SRangeTerm{
							Symbol: ">",
							Value:  &RangeValue{SingleValue: []string{`2021`, "-", "11", "-", "04"}},
						},
					}},
					PrefixTerms: []*WPrefixTerm{
						{Symbol: "+", Elem: &TermGroupElem{
							SRangeTerm: &SRangeTerm{
								Symbol: "<",
								Value:  &RangeValue{SingleValue: []string{`2021`, "-", "11", "-", "11"}},
							},
						}},
					},
				},
				BoostSymbol: "^1.9",
			},
			boost: 1.9,
		},
		{
			name:  "TestTermGroup05",
			input: `( [-1 TO 3] )`,
			want: &TermGroup{
				PrefixTermGroup: &PrefixTermGroup{
					PrefixTerm: &PrefixTerm{Elem: &TermGroupElem{
						DRangeTerm: &DRangeTerm{
							LBRACKET: "[", LValue: &RangeValue{SingleValue: []string{"-", "1"}},
							RValue: &RangeValue{SingleValue: []string{"3"}}, RBRACKET: "]",
						}}},
				},
			},
			boost: 1.0,
		},
		{
			name:  "TestTermGroup06",
			input: `( [-1 TO 3] [1 TO 2] +[5 TO 10}  -{8 TO 90] )`,
			want: &TermGroup{
				PrefixTermGroup: &PrefixTermGroup{
					PrefixTerm: &PrefixTerm{Elem: &TermGroupElem{
						DRangeTerm: &DRangeTerm{
							LBRACKET: "[", LValue: &RangeValue{SingleValue: []string{"-", "1"}},
							RValue: &RangeValue{SingleValue: []string{"3"}}, RBRACKET: "]",
						}}},
					PrefixTerms: []*WPrefixTerm{
						{Elem: &TermGroupElem{
							DRangeTerm: &DRangeTerm{
								LBRACKET: "[", LValue: &RangeValue{SingleValue: []string{"1"}},
								RValue: &RangeValue{SingleValue: []string{"2"}}, RBRACKET: "]",
							},
						}},
						{Symbol: "+", Elem: &TermGroupElem{
							DRangeTerm: &DRangeTerm{
								LBRACKET: "[", LValue: &RangeValue{SingleValue: []string{"5"}},
								RValue: &RangeValue{SingleValue: []string{"10"}}, RBRACKET: "}",
							},
						}},
						{Symbol: "-", Elem: &TermGroupElem{
							DRangeTerm: &DRangeTerm{
								LBRACKET: "{", LValue: &RangeValue{SingleValue: []string{"8"}},
								RValue: &RangeValue{SingleValue: []string{"90"}}, RBRACKET: "]",
							},
						}},
					},
				},
			},
			boost: 1.0,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var out = &TermGroup{}
			err := termParser.ParseString(tt.input, out)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, out)
			assert.Equal(t, tt.boost, out.Boost())
		})
	}
}
