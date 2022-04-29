package term

import (
	"math"
	"reflect"
	"testing"

	"github.com/alecthomas/participle"
	op "github.com/zhuliquan/lucene_parser/operator"
	"github.com/zhuliquan/lucene_parser/token"
)

func TestLogicTermGroup(t *testing.T) {
	var termParser = participle.MustBuild(
		&LogicTermGroup{},
		participle.Lexer(token.Lexer),
	)

	type testCase struct {
		name  string
		input string
		want  *LogicTermGroup
	}
	var testCases = []testCase{
		{
			name:  "TestLogicTermGroup01",
			input: `((quick AND fox) OR (brown AND fox) OR fox) AND NOT news`,
			want: &LogicTermGroup{
				OrTermGroup: &OrTermGroup{
					AndTermGroup: &AndTermGroup{
						ParenTermGroup: &ParenTermGroup{
							SubTermGroup: &LogicTermGroup{
								OrTermGroup: &OrTermGroup{
									AndTermGroup: &AndTermGroup{
										ParenTermGroup: &ParenTermGroup{
											SubTermGroup: &LogicTermGroup{
												OrTermGroup: &OrTermGroup{
													AndTermGroup: &AndTermGroup{
														TermGroupElem: &TermGroupElem{
															SingleTerm: &SingleTerm{Begin: "quick"},
														},
													},
													AnSTermGroup: []*AnSTermGroup{
														{
															AndSymbol: &op.AndSymbol{Symbol: "AND"},
															AndTermGroup: &AndTermGroup{
																TermGroupElem: &TermGroupElem{
																	SingleTerm: &SingleTerm{Begin: "fox"},
																},
															},
														},
													},
												},
											},
										},
									},
								},
								OSTermGroup: []*OSTermGroup{
									{
										OrSymbol: &op.OrSymbol{Symbol: "OR"},
										OrTermGroup: &OrTermGroup{
											AndTermGroup: &AndTermGroup{
												ParenTermGroup: &ParenTermGroup{
													SubTermGroup: &LogicTermGroup{
														OrTermGroup: &OrTermGroup{
															AndTermGroup: &AndTermGroup{
																TermGroupElem: &TermGroupElem{
																	SingleTerm: &SingleTerm{Begin: "brown"},
																},
															},
															AnSTermGroup: []*AnSTermGroup{
																{
																	AndSymbol: &op.AndSymbol{Symbol: "AND"},
																	AndTermGroup: &AndTermGroup{
																		TermGroupElem: &TermGroupElem{
																			SingleTerm: &SingleTerm{Begin: "fox"},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
									{
										OrSymbol: &op.OrSymbol{Symbol: "OR"},
										OrTermGroup: &OrTermGroup{
											AndTermGroup: &AndTermGroup{
												TermGroupElem: &TermGroupElem{
													SingleTerm: &SingleTerm{Begin: "fox"},
												},
											},
										},
									},
								},
							},
						},
					},
					AnSTermGroup: []*AnSTermGroup{
						{
							AndSymbol: &op.AndSymbol{Symbol: "AND"},
							AndTermGroup: &AndTermGroup{
								NotSymbol: &op.NotSymbol{Symbol: "NOT"},
								TermGroupElem: &TermGroupElem{
									SingleTerm: &SingleTerm{Begin: "news"},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var out = &LogicTermGroup{}
			if err := termParser.ParseString(tt.input, out); err != nil {
				t.Errorf("failed to parse input: %s, err: %+v", tt.input, err)
			}
			if !reflect.DeepEqual(tt.want, out) {
				t.Errorf("termParser.ParseString( %s ) = %+v, want: %+v", tt.input, out, tt.want)
			}
		})
	}
}

func TestTermGroup(t *testing.T) {
	var termParser = participle.MustBuild(
		&TermGroup{},
		participle.Lexer(token.Lexer),
	)

	type testCase struct {
		name     string
		input    string
		want     *TermGroup
		boost    float64
		termType TermType
		wantStr  string
	}
	var testCases = []testCase{
		{
			name:  "TestLogicTermGroup01",
			input: `(((quick and fox) OR (brown AND fox) OR fox) AND NOT news)^8.78`,
			want: &TermGroup{
				LogicTermGroup: &LogicTermGroup{
					OrTermGroup: &OrTermGroup{
						AndTermGroup: &AndTermGroup{
							ParenTermGroup: &ParenTermGroup{
								SubTermGroup: &LogicTermGroup{
									OrTermGroup: &OrTermGroup{
										AndTermGroup: &AndTermGroup{
											ParenTermGroup: &ParenTermGroup{
												SubTermGroup: &LogicTermGroup{
													OrTermGroup: &OrTermGroup{
														AndTermGroup: &AndTermGroup{
															TermGroupElem: &TermGroupElem{
																SingleTerm: &SingleTerm{Begin: "quick"},
															},
														},
														AnSTermGroup: []*AnSTermGroup{
															{
																AndSymbol: &op.AndSymbol{Symbol: "and"},
																AndTermGroup: &AndTermGroup{
																	TermGroupElem: &TermGroupElem{
																		SingleTerm: &SingleTerm{Begin: "fox"},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
									OSTermGroup: []*OSTermGroup{
										{
											OrSymbol: &op.OrSymbol{Symbol: "OR"},
											OrTermGroup: &OrTermGroup{
												AndTermGroup: &AndTermGroup{
													ParenTermGroup: &ParenTermGroup{
														SubTermGroup: &LogicTermGroup{
															OrTermGroup: &OrTermGroup{
																AndTermGroup: &AndTermGroup{
																	TermGroupElem: &TermGroupElem{
																		SingleTerm: &SingleTerm{Begin: "brown"},
																	},
																},
																AnSTermGroup: []*AnSTermGroup{
																	{
																		AndSymbol: &op.AndSymbol{Symbol: "AND"},
																		AndTermGroup: &AndTermGroup{
																			TermGroupElem: &TermGroupElem{
																				SingleTerm: &SingleTerm{Begin: "fox"},
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
										{
											OrSymbol: &op.OrSymbol{Symbol: "OR"},
											OrTermGroup: &OrTermGroup{
												AndTermGroup: &AndTermGroup{
													TermGroupElem: &TermGroupElem{
														SingleTerm: &SingleTerm{Begin: "fox"},
													},
												},
											},
										},
									},
								},
							},
						},
						AnSTermGroup: []*AnSTermGroup{
							{
								AndSymbol: &op.AndSymbol{Symbol: "AND"},
								AndTermGroup: &AndTermGroup{
									NotSymbol: &op.NotSymbol{Symbol: "NOT"},
									TermGroupElem: &TermGroupElem{
										SingleTerm: &SingleTerm{Begin: "news"},
									},
								},
							},
						},
					},
				},
				BoostSymbol: "^8.78",
			},
			boost:    8.78,
			termType: GROUP_TERM_TYPE | BOOST_TERM_TYPE,
			wantStr:  `( ( ( quick AND fox ) OR ( brown AND fox ) OR fox ) AND NOT news )^8.78`,
		},
		{
			name:  "TestLogicTermGroup02",
			input: `(((quick and fox) OR (brown AND fox) OR fox) AND NOT news)`,
			want: &TermGroup{
				LogicTermGroup: &LogicTermGroup{
					OrTermGroup: &OrTermGroup{
						AndTermGroup: &AndTermGroup{
							ParenTermGroup: &ParenTermGroup{
								SubTermGroup: &LogicTermGroup{
									OrTermGroup: &OrTermGroup{
										AndTermGroup: &AndTermGroup{
											ParenTermGroup: &ParenTermGroup{
												SubTermGroup: &LogicTermGroup{
													OrTermGroup: &OrTermGroup{
														AndTermGroup: &AndTermGroup{
															TermGroupElem: &TermGroupElem{
																SingleTerm: &SingleTerm{Begin: "quick"},
															},
														},
														AnSTermGroup: []*AnSTermGroup{
															{
																AndSymbol: &op.AndSymbol{Symbol: "and"},
																AndTermGroup: &AndTermGroup{
																	TermGroupElem: &TermGroupElem{
																		SingleTerm: &SingleTerm{Begin: "fox"},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
									OSTermGroup: []*OSTermGroup{
										{
											OrSymbol: &op.OrSymbol{Symbol: "OR"},
											OrTermGroup: &OrTermGroup{
												AndTermGroup: &AndTermGroup{
													ParenTermGroup: &ParenTermGroup{
														SubTermGroup: &LogicTermGroup{
															OrTermGroup: &OrTermGroup{
																AndTermGroup: &AndTermGroup{
																	TermGroupElem: &TermGroupElem{
																		SingleTerm: &SingleTerm{Begin: "brown"},
																	},
																},
																AnSTermGroup: []*AnSTermGroup{
																	{
																		AndSymbol: &op.AndSymbol{Symbol: "AND"},
																		AndTermGroup: &AndTermGroup{
																			TermGroupElem: &TermGroupElem{
																				SingleTerm: &SingleTerm{Begin: "fox"},
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
										{
											OrSymbol: &op.OrSymbol{Symbol: "OR"},
											OrTermGroup: &OrTermGroup{
												AndTermGroup: &AndTermGroup{
													TermGroupElem: &TermGroupElem{
														SingleTerm: &SingleTerm{Begin: "fox"},
													},
												},
											},
										},
									},
								},
							},
						},
						AnSTermGroup: []*AnSTermGroup{
							{
								AndSymbol: &op.AndSymbol{Symbol: "AND"},
								AndTermGroup: &AndTermGroup{
									NotSymbol: &op.NotSymbol{Symbol: "NOT"},
									TermGroupElem: &TermGroupElem{
										SingleTerm: &SingleTerm{Begin: "news"},
									},
								},
							},
						},
					},
				},
			},
			boost:    1.0,
			termType: GROUP_TERM_TYPE,
			wantStr:  `( ( ( quick AND fox ) OR ( brown AND fox ) OR fox ) AND NOT news )`,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var out = &TermGroup{}
			if err := termParser.ParseString(tt.input, out); err != nil {
				t.Errorf("failed to parse input: %s, err: %+v", tt.input, err)
			}
			if !reflect.DeepEqual(tt.want, out) {
				t.Errorf("termParser.ParseString( %s ) = %+v, want: %+v", tt.input, out, tt.want)
			}
			if math.Abs(tt.boost-out.Boost()) > 1E-6 {
				t.Errorf("expect get boost: %+v, but get boost: %+v", tt.boost, out.Boost())
			}
			if out.GetTermType() != tt.termType {
				t.Errorf("expect term type: %v but got type: %v", tt.termType, out.GetTermType())
			}
			if out.String() != tt.wantStr {
				t.Errorf("expect %s, but %s", tt.wantStr, out.String())
			}
		})
	}
	// test empty term group
	var out *TermGroup
	if out.String() != "" {
		t.Errorf("expect empty")
	}
	if out.Boost() != 0.0 {
		t.Errorf("expect no boost")
	}
	if out.GetTermType() != UNKNOWN_TERM_TYPE {
		t.Errorf("expect to unknown term type")
	}

	out = &TermGroup{}
	if out.String() != "" {
		t.Errorf("expect empty")
	}
	if out.Boost() != 0.0 {
		t.Errorf("expect no boost")
	}
	if out.GetTermType() != UNKNOWN_TERM_TYPE {
		t.Errorf("expect to unknown term type")
	}
	if _, err := out.Value(func(s string) (interface{}, error) { return s, nil }); err != ErrEmptyGroupTerm {
		t.Errorf("expect got empty term error, err: %v", err)
	}

	var l *LogicTermGroup
	if l.String() != "" {
		t.Error("expect empty")
	}
	l = &LogicTermGroup{}
	if l.String() != "" {
		t.Error("expect empty")
	}
	if _, err := out.Value(func(s string) (interface{}, error) { return s, nil }); err != ErrEmptyGroupTerm {
		t.Errorf("expect got empty term error, err: %v", err)
	}

	var a *AndTermGroup
	if a.String() != "" {
		t.Error("expect empty")
	}
	a = &AndTermGroup{}
	if a.String() != "" {
		t.Error("expect empty")
	}

	var as *AnSTermGroup
	if as.String() != "" {
		t.Error("expect empty")
	}
	as = &AnSTermGroup{}
	if as.String() != "" {
		t.Error("expect empty")
	}

	var o *OrTermGroup
	if o.String() != "" {
		t.Error("expect empty")
	}
	o = &OrTermGroup{}
	if o.String() != "" {
		t.Error("expect empty")
	}

	var os *OSTermGroup
	if os.String() != "" {
		t.Error("expect empty")
	}
	os = &OSTermGroup{}
	if os.String() != "" {
		t.Error("expect empty")
	}

	var p *ParenTermGroup
	if p.String() != "" {
		t.Error("expect empty")
	}
	p = &ParenTermGroup{}
	if p.String() != "" {
		t.Errorf("expect empty")
	}

}
