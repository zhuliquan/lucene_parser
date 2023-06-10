package lucene_parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhuliquan/lucene_parser/operator"
	"github.com/zhuliquan/lucene_parser/term"
)

func TestLucene(t *testing.T) {
	type testCase struct {
		name    string
		input   string
		want    *Lucene
		wantErr bool
		wantStr string
	}

	var testCases = []testCase{
		{
			name:  "Test_space_not_replace_or_not_01",
			input: `x:1 NOT x:2`,
			want: &Lucene{
				OrQuery: &OrQuery{
					AndQuery: &AndQuery{
						FieldQuery: &FieldQuery{
							Field: &term.Field{Value: []string{"x"}},
							Term: &term.Term{FuzzyTerm: &term.FuzzyTerm{
								SingleTerm: &term.SingleTerm{Begin: "1"},
							}},
						},
					},
					AnSQuery: []*AnSQuery{
						{
							NotSymbol: &operator.NotSymbol{Symbol: "NOT"},
							AndQuery: &AndQuery{
								FieldQuery: &FieldQuery{
									Field: &term.Field{Value: []string{"x"}},
									Term: &term.Term{FuzzyTerm: &term.FuzzyTerm{
										SingleTerm: &term.SingleTerm{Begin: "2"},
									}},
								},
							},
						},
					},
				},
			},
			wantErr: false,
			wantStr: `x:1 AND NOT x:2`,
		},
		{
			name:  "Test_space_not_replace_or_not_02",
			input: `!x:1 NOT x:2`,
			want: &Lucene{
				OrQuery: &OrQuery{
					AndQuery: &AndQuery{
						NotSymbol: &operator.NotSymbol{Symbol: "!"},
						FieldQuery: &FieldQuery{
							Field: &term.Field{Value: []string{"x"}},
							Term: &term.Term{FuzzyTerm: &term.FuzzyTerm{
								SingleTerm: &term.SingleTerm{Begin: "1"},
							}},
						},
					},
					AnSQuery: []*AnSQuery{
						{
							NotSymbol: &operator.NotSymbol{Symbol: "NOT"},
							AndQuery: &AndQuery{
								FieldQuery: &FieldQuery{
									Field: &term.Field{Value: []string{"x"}},
									Term: &term.Term{FuzzyTerm: &term.FuzzyTerm{
										SingleTerm: &term.SingleTerm{Begin: "2"},
									}},
								},
							},
						},
					},
				},
			},
			wantErr: false,
			wantStr: `NOT x:1 AND NOT x:2`,
		},
		{
			name:  "TestLucene01",
			input: `x:1 AND NOT x:2`,
			want: &Lucene{
				OrQuery: &OrQuery{
					AndQuery: &AndQuery{
						FieldQuery: &FieldQuery{
							Field: &term.Field{Value: []string{"x"}},
							Term: &term.Term{FuzzyTerm: &term.FuzzyTerm{
								SingleTerm: &term.SingleTerm{Begin: "1"},
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
										SingleTerm: &term.SingleTerm{Begin: "2"},
									}},
								},
							},
						},
					},
				},
			},
			wantErr: false,
			wantStr: `x:1 AND NOT x:2`,
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
												SingleTerm: &term.SingleTerm{Begin: "1"},
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
														SingleTerm: &term.SingleTerm{Begin: "2"},
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
										SingleTerm: &term.SingleTerm{Begin: "9"},
									}},
								},
							},
						},
					},
				},
			},
			wantErr: false,
			wantStr: `NOT ( x:1 AND y:2 ) OR z:9`,
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
												SingleTerm: &term.SingleTerm{Begin: "1"},
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
														SingleTerm: &term.SingleTerm{Begin: "2"},
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
														SingleTerm: &term.SingleTerm{Begin: "8"},
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
																SingleTerm: &term.SingleTerm{Begin: "90"},
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
			wantStr: `( x:1 AND NOT y:2 ) AND ( NOT x:8 AND k:90 )`,
		},
		{
			name:  "TestLucene04",
			input: `x:(txt OR foo OR bar) AND NOT x-y:"xxx" OR !zz:iopio\ 90`,
			want: &Lucene{
				OrQuery: &OrQuery{
					AndQuery: &AndQuery{
						FieldQuery: &FieldQuery{
							Field: &term.Field{Value: []string{"x"}},
							Term: &term.Term{TermGroup: &term.TermGroup{
								LogicTermGroup: &term.LogicTermGroup{
									OrTermGroup: &term.OrTermGroup{
										AndTermGroup: &term.AndTermGroup{
											FieldTermGroup: &term.FieldTermGroup{
												SingleTerm: &term.SingleTerm{Begin: "txt"},
											},
										},
									},
									OSTermGroup: []*term.OSTermGroup{
										{
											OrSymbol: &operator.OrSymbol{Symbol: "OR"},
											OrTermGroup: &term.OrTermGroup{
												AndTermGroup: &term.AndTermGroup{
													FieldTermGroup: &term.FieldTermGroup{
														SingleTerm: &term.SingleTerm{Begin: "foo"},
													},
												},
											},
										},
										{
											OrSymbol: &operator.OrSymbol{Symbol: "OR"},
											OrTermGroup: &term.OrTermGroup{
												AndTermGroup: &term.AndTermGroup{
													FieldTermGroup: &term.FieldTermGroup{
														SingleTerm: &term.SingleTerm{Begin: "bar"},
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
										PhraseTerm: &term.PhraseTerm{Chars: []string{`xxx`}},
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
								NotSymbol: &operator.NotSymbol{Symbol: "!"},
								FieldQuery: &FieldQuery{
									Field: &term.Field{Value: []string{"zz"}},
									Term: &term.Term{FuzzyTerm: &term.FuzzyTerm{
										SingleTerm: &term.SingleTerm{Begin: `iopio`, Chars: []string{`\ `, `90`}},
									}},
								},
							},
						},
					},
				},
			},
			wantErr: false,
			wantStr: `x:( txt OR foo OR bar ) AND NOT x-y:"xxx" OR NOT zz:iopio\ 90`,
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
														OrTermGroup: &term.OrTermGroup{AndTermGroup: &term.AndTermGroup{FieldTermGroup: &term.FieldTermGroup{SingleTerm: &term.SingleTerm{Begin: "foo"}}}},
														OSTermGroup: []*term.OSTermGroup{
															{
																OrSymbol: &operator.OrSymbol{Symbol: "or"},
																OrTermGroup: &term.OrTermGroup{
																	AndTermGroup: &term.AndTermGroup{FieldTermGroup: &term.FieldTermGroup{SingleTerm: &term.SingleTerm{Begin: "bar"}}},
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
										SingleTerm: &term.SingleTerm{Begin: `you`},
									}},
								},
							},
						},
					},
				},
			},
			wantErr: false,
			wantStr: `( NOT x:( foo OR bar ) ) AND z:you`,
		},
		{
			name:    "parse_wrong_lucene",
			input:   `x:("dsa`,
			want:    nil,
			wantErr: true,
			wantStr: "",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			lucene, err := ParseLucene(tt.input)
			assert.Equal(t, tt.wantErr, (err != nil))
			assert.Equal(t, tt.want, lucene)
			assert.Equal(t, tt.wantStr, lucene.String())
		})
	}
}

func TestType(t *testing.T) {
	type test struct {
		name  string
		input Query
		qType QueryType
	}
	for _, tt := range []test{
		{
			name:  "test_lucene_type",
			input: &Lucene{},
			qType: LUCENE_QUERY,
		},
		{
			name:  "test_or_query_type",
			input: &OrQuery{},
			qType: OR_QUERY,
		},
		{
			name:  "test_os_query_type",
			input: &OSQuery{},
			qType: OS_QUERY,
		},
		{
			name:  "test_and_query_type",
			input: &AndQuery{},
			qType: AND_QUERY,
		},
		{
			name:  "test_ans_query_type",
			input: &AnSQuery{},
			qType: ANS_QUERY,
		},
		{
			name:  "test_paren_query_type",
			input: &ParenQuery{},
			qType: PAREN_QUERY,
		},
		{
			name:  "test_field_query_type",
			input: &FieldQuery{},
			qType: FIELD_QUERY,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.qType, tt.input.GetQueryType())
		})
	}
}

func TestWrongString(t *testing.T) {
	type test struct {
		name  string
		input Query
	}
	for _, tt := range []test{
		{
			name:  "test_lucene_string",
			input: &Lucene{},
		},
		{
			name:  "test_or_query_string",
			input: &OrQuery{},
		},
		{
			name:  "test_os_query_string",
			input: &OSQuery{},
		},
		{
			name:  "test_and_query_string",
			input: &AndQuery{},
		},
		{
			name:  "test_ans_query_string",
			input: &AnSQuery{},
		},
		{
			name:  "test_paren_query_string",
			input: &ParenQuery{},
		},
		{
			name:  "test_field_query_string",
			input: &FieldQuery{},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, "", tt.input.String())
		})
	}
}

func TestWrongLuceneString(t *testing.T) {
	var l *Lucene
	assert.Equal(t, "", l.String())
	var o *OrQuery
	assert.Equal(t, "", o.String())
	var s *OSQuery
	assert.Equal(t, "", s.String())
	var a *AndQuery
	assert.Equal(t, "", a.String())
	var x *AnSQuery
	assert.Equal(t, "", x.String())
	var f *FieldQuery
	assert.Equal(t, "", f.String())
	var p *ParenQuery
	assert.Equal(t, "", p.String())
}
