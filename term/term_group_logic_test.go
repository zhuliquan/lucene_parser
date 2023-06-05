package term

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/alecthomas/participle"
	"github.com/stretchr/testify/assert"
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
		name     string
		input    string
		want     *TermGroup
		boost    float64
		termType TermType
		wantStr  string
	}
	var testCases = []testCase{
		{
			name:  "Test_Space_not_as_or_not",
			input: `( x not y !z and x1 and not x2 not x3 OR not x4)`,
			want: &TermGroup{
				LogicTermGroup: &LogicTermGroup{
					OrTermGroup: &OrTermGroup{
						AndTermGroup: &AndTermGroup{
							TermGroupElem: &TermGroupElem{
								SingleTerm: &SingleTerm{Begin: "x"},
							},
						},
					},
					OSTermGroup: []*OSTermGroup{
						{
							NotSymbol: &op.NotSymbol{Symbol: "not"},
							OrTermGroup: &OrTermGroup{
								AndTermGroup: &AndTermGroup{
									TermGroupElem: &TermGroupElem{
										SingleTerm: &SingleTerm{Begin: "y"},
									},
								},
							},
						},
						{
							NotSymbol: &op.NotSymbol{Symbol: "!"},
							OrTermGroup: &OrTermGroup{
								AndTermGroup: &AndTermGroup{
									TermGroupElem: &TermGroupElem{
										SingleTerm: &SingleTerm{Begin: "z"},
									},
								},
								AnSTermGroup: []*AnSTermGroup{
									{
										AndSymbol: &op.AndSymbol{Symbol: "and"},
										AndTermGroup: &AndTermGroup{
											TermGroupElem: &TermGroupElem{
												SingleTerm: &SingleTerm{Begin: "x", Chars: []string{"1"}},
											},
										},
									},
									{
										AndSymbol: &op.AndSymbol{Symbol: "and"},
										AndTermGroup: &AndTermGroup{
											NotSymbol: &op.NotSymbol{Symbol: "not"},
											TermGroupElem: &TermGroupElem{
												SingleTerm: &SingleTerm{Begin: "x", Chars: []string{"2"}},
											},
										},
									},
								},
							},
						},
						{
							NotSymbol: &op.NotSymbol{Symbol: "not"},
							OrTermGroup: &OrTermGroup{
								AndTermGroup: &AndTermGroup{
									TermGroupElem: &TermGroupElem{
										SingleTerm: &SingleTerm{Begin: "x", Chars: []string{"3"}},
									},
								},
							},
						},
						{
							OrSymbol: &op.OrSymbol{Symbol: "OR"},
							OrTermGroup: &OrTermGroup{
								AndTermGroup: &AndTermGroup{
									NotSymbol: &op.NotSymbol{Symbol: "not"},
									TermGroupElem: &TermGroupElem{
										SingleTerm: &SingleTerm{Begin: "x", Chars: []string{"4"}},
									},
								},
							},
						},
					},
				},
			},
			boost:    1.0,
			termType: GROUP_TERM_TYPE,
			wantStr:  `( x OR NOT y OR NOT z AND x1 AND NOT x2 OR NOT x3 OR NOT x4 )`,
		},
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
		{
			name:  "TestLogicTermGroup03",
			input: `(((quick and fox) OR (brown AND fox) OR fox) AND NOT news)^`,
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
				BoostSymbol: "^",
			},
			boost:    1.0,
			termType: GROUP_TERM_TYPE | BOOST_TERM_TYPE,
			wantStr:  `( ( ( quick AND fox ) OR ( brown AND fox ) OR fox ) AND NOT news )^`,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var out = &TermGroup{}
			err := termParser.ParseString(tt.input, out)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, out)
			if !reflect.DeepEqual(tt.want, out) {
				x1, _ := json.Marshal(tt.want)
				x2, _ := json.Marshal(out)
				t.Logf("want: %s, out: %s\n", x1, x2)
			}
			assert.Equal(t, tt.boost, out.Boost())
			assert.Equal(t, tt.termType, out.GetTermType())
			assert.Equal(t, tt.wantStr, out.String())
		})
	}
	// test empty term group
	var out *TermGroup
	assert.Equal(t, "", out.String())
	assert.Equal(t, 0.0, out.Boost())
	assert.Equal(t, UNKNOWN_TERM_TYPE, out.GetTermType())

	out = &TermGroup{}
	assert.Equal(t, "", out.String())
	assert.Equal(t, 0.0, out.Boost())
	assert.Equal(t, UNKNOWN_TERM_TYPE, out.GetTermType())
	_, err := out.Value(func(s string) (interface{}, error) { return s, nil })
	assert.Equal(t, ErrEmptyGroupTerm, err)

	var l *LogicTermGroup
	assert.Equal(t, "", l.String())
	l = &LogicTermGroup{}
	assert.Equal(t, "", l.String())
	_, err = out.Value(func(s string) (interface{}, error) { return s, nil })
	assert.Equal(t, ErrEmptyGroupTerm, err)

	var a *AndTermGroup
	assert.Equal(t, "", a.String())
	a = &AndTermGroup{}
	assert.Equal(t, "", a.String())

	var as *AnSTermGroup
	assert.Equal(t, "", as.String())
	as = &AnSTermGroup{}
	assert.Equal(t, "", as.String())

	var o *OrTermGroup
	assert.Equal(t, "", o.String())
	o = &OrTermGroup{}
	assert.Equal(t, "", o.String())

	var os *OSTermGroup
	assert.Equal(t, "", os.String())
	os = &OSTermGroup{}
	assert.Equal(t, "", os.String())

	var p *ParenTermGroup
	assert.Equal(t, "", p.String())
	p = &ParenTermGroup{}
	assert.Equal(t, "", p.String())
}
