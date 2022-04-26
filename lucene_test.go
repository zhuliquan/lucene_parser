package lucene_parser

import (
	"reflect"
	"testing"

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
												SingleTerm: &term.SingleTerm{Begin: "txt"},
											},
										},
									},
									OSTermGroup: []*term.OSTermGroup{
										{
											OrSymbol: &operator.OrSymbol{Symbol: "OR"},
											OrTermGroup: &term.OrTermGroup{
												AndTermGroup: &term.AndTermGroup{
													TermGroupElem: &term.TermGroupElem{
														SingleTerm: &term.SingleTerm{Begin: "foo"},
													},
												},
											},
										},
										{
											OrSymbol: &operator.OrSymbol{Symbol: "OR"},
											OrTermGroup: &term.OrTermGroup{
												AndTermGroup: &term.AndTermGroup{
													TermGroupElem: &term.TermGroupElem{
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
								NotSymbol: &operator.NotSymbol{Symbol: "NOT"},
								FieldQuery: &FieldQuery{
									Field: &term.Field{Value: []string{"zz"}},
									Term: &term.Term{FuzzyTerm: &term.FuzzyTerm{
										SingleTerm: &term.SingleTerm{Begin: `iopio\ `, Chars: []string{`90`}},
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
														OrTermGroup: &term.OrTermGroup{AndTermGroup: &term.AndTermGroup{TermGroupElem: &term.TermGroupElem{SingleTerm: &term.SingleTerm{Begin: "foo"}}}},
														OSTermGroup: []*term.OSTermGroup{
															{
																OrSymbol: &operator.OrSymbol{Symbol: "or"},
																OrTermGroup: &term.OrTermGroup{
																	AndTermGroup: &term.AndTermGroup{TermGroupElem: &term.TermGroupElem{SingleTerm: &term.SingleTerm{Begin: "bar"}}},
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
			if lucene, err := ParseLucene(tt.input); (err != nil) != tt.wantErr {
				t.Errorf("parser lucene, err: %+v", err)
			} else if !reflect.DeepEqual(lucene, tt.want) {
				t.Errorf("luceneParser.ParseString( %s ) = %+v, but want %+v", tt.input, lucene, tt.want)
			} else if lucene.String() != tt.wantStr {
				t.Errorf("luceneParser.ParseString( %s ) = %s, but want %s", tt.input, lucene, tt.wantStr)
			}
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
			if tt.input.GetQueryType() != tt.qType {
				t.Errorf("expect to %v, but got %v", tt.qType, tt.input.GetQueryType())
			}
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
			if tt.input.String() != "" {
				t.Errorf("expect to empty")
			}
		})
	}
}

func TestWrongLuceneString(t *testing.T) {
	var l *Lucene
	if l.String() != "" {
		t.Error("expect to empty")
	}
	var o *OrQuery
	if o.String() != "" {
		t.Error("expect to empty")
	}
	var s *OSQuery
	if s.String() != "" {
		t.Error("expect to empty")
	}
	var a *AndQuery
	if a.String() != "" {
		t.Error("expect to empty")
	}
	var x *AnSQuery
	if x.String() != "" {
		t.Error("expect to empty")
	}
	var f *FieldQuery
	if f.String() != "" {
		t.Error("expect to empty")
	}
	var p *ParenQuery
	if p.String() != "" {
		t.Error("expect to empty")
	}
}
