package lucene_parser

import (
	"reflect"
	"testing"

	"github.com/alecthomas/participle"
	"github.com/zhuliquan/lucene_parser/operator"
	"github.com/zhuliquan/lucene_parser/term"
	"github.com/zhuliquan/lucene_parser/token"
)

func TestLucene(t *testing.T) {
	var luceneParser = participle.MustBuild(
		&Lucene{},
		participle.Lexer(token.Lexer),
	)

	type testCase struct {
		name    string
		input   string
		want    *Lucene
		wantErr bool
	}

	var testCases = []testCase{
		{
			name:  "TestLucene01",
			input: `x:1 AND NOT x:2`,
			want: &Lucene{
				OrQuery: &OrQuery{
					AndQuery: &AndQuery{
						FieldQuery: &FieldQuery{
							Field: &term.Field{Value: []string{"x"}},
							Term: &term.Term{FuzzyTerm: &term.FuzzyTerm{
								SingleTerm: &term.SingleTerm{Value: []string{"1"}},
							}},
						},
					},
					AnSQuery: []*AnSQuery{
						{
							AndSymbol: &operator.AndSymbol{Symbol: "AND"},
							AndQuery: &AndQuery{
								NotSymbol: &operator.NotSymbol{Symbol: "NOT"},
								FieldQuery: &FieldQuery{
									Field: &term.Field{Value: []string{"x"}},
									Term: &term.Term{FuzzyTerm: &term.FuzzyTerm{
										SingleTerm: &term.SingleTerm{Value: []string{"2"}},
									}},
								},
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name:  "TestLucene02",
			input: `NOT (x:1 AND y:2) OR z:9`,
			want: &Lucene{
				OrQuery: &OrQuery{
					AndQuery: &AndQuery{
						NotSymbol: &operator.NotSymbol{Symbol: "NOT"},
						ParenQuery: &ParenQuery{
							SubQuery: &Lucene{
								OrQuery: &OrQuery{
									AndQuery: &AndQuery{
										FieldQuery: &FieldQuery{
											Field: &term.Field{Value: []string{"x"}},
											Term: &term.Term{FuzzyTerm: &term.FuzzyTerm{
												SingleTerm: &term.SingleTerm{Value: []string{"1"}},
											}},
										},
									},
									AnSQuery: []*AnSQuery{
										{
											AndSymbol: &operator.AndSymbol{Symbol: "AND"},
											AndQuery: &AndQuery{
												FieldQuery: &FieldQuery{
													Field: &term.Field{Value: []string{"y"}},
													Term: &term.Term{FuzzyTerm: &term.FuzzyTerm{
														SingleTerm: &term.SingleTerm{Value: []string{"2"}},
													}},
												},
											},
										},
									},
								},
							},
						},
					},
				},
				OSQuery: []*OSQuery{
					{
						OrSymbol: &operator.OrSymbol{Symbol: "OR"},
						OrQuery: &OrQuery{
							AndQuery: &AndQuery{
								FieldQuery: &FieldQuery{
									Field: &term.Field{Value: []string{"z"}},
									Term: &term.Term{FuzzyTerm: &term.FuzzyTerm{
										SingleTerm: &term.SingleTerm{Value: []string{"9"}},
									}},
								},
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name:  "TestLucene03",
			input: `(x:1 AND NOT y:2) AND (NOT x:8 AND k:90)`,
			want: &Lucene{
				OrQuery: &OrQuery{
					AndQuery: &AndQuery{
						ParenQuery: &ParenQuery{
							SubQuery: &Lucene{
								OrQuery: &OrQuery{
									AndQuery: &AndQuery{
										FieldQuery: &FieldQuery{
											Field: &term.Field{Value: []string{"x"}},
											Term: &term.Term{FuzzyTerm: &term.FuzzyTerm{
												SingleTerm: &term.SingleTerm{Value: []string{"1"}},
											}},
										},
									},
									AnSQuery: []*AnSQuery{
										{
											AndSymbol: &operator.AndSymbol{Symbol: "AND"},
											AndQuery: &AndQuery{
												NotSymbol: &operator.NotSymbol{Symbol: "NOT"},
												FieldQuery: &FieldQuery{
													Field: &term.Field{Value: []string{"y"}},
													Term: &term.Term{FuzzyTerm: &term.FuzzyTerm{
														SingleTerm: &term.SingleTerm{Value: []string{"2"}},
													}},
												},
											},
										},
									},
								},
							},
						},
					},
					AnSQuery: []*AnSQuery{
						{
							AndSymbol: &operator.AndSymbol{Symbol: "AND"},
							AndQuery: &AndQuery{
								ParenQuery: &ParenQuery{
									SubQuery: &Lucene{
										OrQuery: &OrQuery{
											AndQuery: &AndQuery{
												NotSymbol: &operator.NotSymbol{Symbol: "NOT"},
												FieldQuery: &FieldQuery{
													Field: &term.Field{Value: []string{"x"}},
													Term: &term.Term{FuzzyTerm: &term.FuzzyTerm{
														SingleTerm: &term.SingleTerm{Value: []string{"8"}},
													}},
												},
											},
											AnSQuery: []*AnSQuery{
												{
													AndSymbol: &operator.AndSymbol{Symbol: "AND"},
													AndQuery: &AndQuery{
														FieldQuery: &FieldQuery{
															Field: &term.Field{Value: []string{"k"}},
															Term: &term.Term{FuzzyTerm: &term.FuzzyTerm{
																SingleTerm: &term.SingleTerm{Value: []string{"90"}},
															}},
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
				},
			},
			wantErr: false,
		},
		{
			name:  "TestLucene04",
			input: `x:(txt OR foo OR bar) AND NOT x-y:"xxx" OR NOT zz:iopio\ 90`,
			want: &Lucene{
				OrQuery: &OrQuery{
					AndQuery: &AndQuery{
						FieldQuery: &FieldQuery{
							Field: &term.Field{Value: []string{"x"}},
							Term: &term.Term{TermGroup: &term.TermGroup{
								LogicTermGroup: &term.LogicTermGroup{
									OrTermGroup: &term.OrTermGroup{
										AndTermGroup: &term.AndTermGroup{
											TermGroupElem: &term.TermGroupElem{
												SingleTerm: &term.SingleTerm{Value: []string{"txt"}},
											},
										},
									},
									OSTermGroup: []*term.OSTermGroup{
										{
											OrSymbol: &operator.OrSymbol{Symbol: "OR"},
											OrTermGroup: &term.OrTermGroup{
												AndTermGroup: &term.AndTermGroup{
													TermGroupElem: &term.TermGroupElem{
														SingleTerm: &term.SingleTerm{Value: []string{"foo"}},
													},
												},
											},
										},
										{
											OrSymbol: &operator.OrSymbol{Symbol: "OR"},
											OrTermGroup: &term.OrTermGroup{
												AndTermGroup: &term.AndTermGroup{
													TermGroupElem: &term.TermGroupElem{
														SingleTerm: &term.SingleTerm{Value: []string{"bar"}},
													},
												},
											},
										},
									},
								},
							}},
						},
					},
					AnSQuery: []*AnSQuery{
						{
							AndSymbol: &operator.AndSymbol{Symbol: "AND"},
							AndQuery: &AndQuery{
								NotSymbol: &operator.NotSymbol{Symbol: "NOT"},
								FieldQuery: &FieldQuery{
									Field: &term.Field{Value: []string{"x", "-", "y"}},
									Term: &term.Term{FuzzyTerm: &term.FuzzyTerm{
										PhraseTerm: &term.PhraseTerm{Value: []string{`xxx`}},
									}},
								},
							},
						},
					},
				},
				OSQuery: []*OSQuery{
					{
						OrSymbol: &operator.OrSymbol{Symbol: "OR"},
						OrQuery: &OrQuery{
							AndQuery: &AndQuery{
								NotSymbol: &operator.NotSymbol{Symbol: "NOT"},
								FieldQuery: &FieldQuery{
									Field: &term.Field{Value: []string{"zz"}},
									Term: &term.Term{FuzzyTerm: &term.FuzzyTerm{
										SingleTerm: &term.SingleTerm{Value: []string{`iopio\ `, `90`}},
									}},
								},
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name:  "TestLucene05",
			input: `(NOT x:(foo or bar)) AND z:you`,
			want: &Lucene{
				OrQuery: &OrQuery{
					AndQuery: &AndQuery{
						ParenQuery: &ParenQuery{
							&Lucene{
								OrQuery: &OrQuery{
									AndQuery: &AndQuery{
										NotSymbol: &operator.NotSymbol{Symbol: "NOT"},
										FieldQuery: &FieldQuery{
											Field: &term.Field{Value: []string{"x"}},
											Term: &term.Term{
												TermGroup: &term.TermGroup{
													LogicTermGroup: &term.LogicTermGroup{
														OrTermGroup: &term.OrTermGroup{AndTermGroup: &term.AndTermGroup{TermGroupElem: &term.TermGroupElem{SingleTerm: &term.SingleTerm{Value: []string{"foo"}}}}},
														OSTermGroup: []*term.OSTermGroup{
															{
																OrSymbol: &operator.OrSymbol{Symbol: "or"},
																OrTermGroup: &term.OrTermGroup{
																	AndTermGroup: &term.AndTermGroup{TermGroupElem: &term.TermGroupElem{SingleTerm: &term.SingleTerm{Value: []string{"bar"}}}},
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
						},
					},
					AnSQuery: []*AnSQuery{
						{
							AndSymbol: &operator.AndSymbol{Symbol: "AND"},
							AndQuery: &AndQuery{
								FieldQuery: &FieldQuery{
									Field: &term.Field{Value: []string{"z"}},
									Term: &term.Term{FuzzyTerm: &term.FuzzyTerm{
										SingleTerm: &term.SingleTerm{Value: []string{`you`}},
									}},
								},
							},
						},
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var lucene = &Lucene{}
			if err := luceneParser.ParseString(tt.input, lucene); (err != nil) != tt.wantErr {
				t.Errorf("parser lucene, err: %+v", err)
			} else if !reflect.DeepEqual(lucene, tt.want) {
				t.Errorf("luceneParser.ParseString( %s ) = %+v, but want %+v", tt.input, lucene, tt.want)
			}
		})
	}
}
