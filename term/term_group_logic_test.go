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
															SingleTerm: &SingleTerm{Value: []string{"quick"}},
														},
													},
													AnSTermGroup: []*AnSTermGroup{
														{
															AndSymbol: &op.AndSymbol{Symbol: "AND"},
															AndTermGroup: &AndTermGroup{
																TermGroupElem: &TermGroupElem{
																	SingleTerm: &SingleTerm{Value: []string{"fox"}},
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
																	SingleTerm: &SingleTerm{Value: []string{"brown"}},
																},
															},
															AnSTermGroup: []*AnSTermGroup{
																{
																	AndSymbol: &op.AndSymbol{Symbol: "AND"},
																	AndTermGroup: &AndTermGroup{
																		TermGroupElem: &TermGroupElem{
																			SingleTerm: &SingleTerm{Value: []string{"fox"}},
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
													SingleTerm: &SingleTerm{Value: []string{"fox"}},
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
									SingleTerm: &SingleTerm{Value: []string{"news"}},
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
			} else if !reflect.DeepEqual(tt.want, out) {
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
		name  string
		input string
		want  *TermGroup
		boost float64
	}
	var testCases = []testCase{
		{
			name:  "TestLogicTermGroup01",
			input: `(((quick AND fox) OR (brown AND fox) OR fox) AND NOT news)^8.78`,
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
																SingleTerm: &SingleTerm{Value: []string{"quick"}},
															},
														},
														AnSTermGroup: []*AnSTermGroup{
															{
																AndSymbol: &op.AndSymbol{Symbol: "AND"},
																AndTermGroup: &AndTermGroup{
																	TermGroupElem: &TermGroupElem{
																		SingleTerm: &SingleTerm{Value: []string{"fox"}},
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
																		SingleTerm: &SingleTerm{Value: []string{"brown"}},
																	},
																},
																AnSTermGroup: []*AnSTermGroup{
																	{
																		AndSymbol: &op.AndSymbol{Symbol: "AND"},
																		AndTermGroup: &AndTermGroup{
																			TermGroupElem: &TermGroupElem{
																				SingleTerm: &SingleTerm{Value: []string{"fox"}},
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
														SingleTerm: &SingleTerm{Value: []string{"fox"}},
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
										SingleTerm: &SingleTerm{Value: []string{"news"}},
									},
								},
							},
						},
					},
				},
				BoostSymbol: "^8.78",
			},
			boost: 8.78,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var out = &TermGroup{}
			if err := termParser.ParseString(tt.input, out); err != nil {
				t.Errorf("failed to parse input: %s, err: %+v", tt.input, err)
			} else if !reflect.DeepEqual(tt.want, out) {
				t.Errorf("termParser.ParseString( %s ) = %+v, want: %+v", tt.input, out, tt.want)
			} else if math.Abs(tt.boost-out.Boost()) > 1E-6 {
				t.Errorf("expect get boost: %+v, but get boost: %+v", tt.boost, out.Boost())
			}
		})
	}
}
