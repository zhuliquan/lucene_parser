package prefix

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhuliquan/lucene_parser"
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
			name:  "Test_(x:1 -x:2)",
			input: `x:1 -x:2`,
			want: &Lucene{
				Clauses: []*PrefixClause{
					{
						FieldQuery: &FieldQuery{
							Field: &term.Field{Value: []string{"x"}},
							Term: &Term{FuzzyTerm: &term.FuzzyTerm{
								SingleTerm: &term.SingleTerm{Begin: "1"},
							}},
						},
					},
					{
						PrefixOp: "-",
						FieldQuery: &FieldQuery{
							Field: &term.Field{Value: []string{"x"}},
							Term: &Term{FuzzyTerm: &term.FuzzyTerm{
								SingleTerm: &term.SingleTerm{Begin: "2"},
							}},
						},
					},
				},
			},
			wantErr: false,
			wantStr: `x:1 -x:2`,
		},
		{
			name:  "Test_(-x:1 -x:2)",
			input: `-x:1 -x:2`,
			want: &Lucene{
				Clauses: []*PrefixClause{
					{
						PrefixOp: "-",
						FieldQuery: &FieldQuery{
							Field: &term.Field{Value: []string{"x"}},
							Term: &Term{FuzzyTerm: &term.FuzzyTerm{
								SingleTerm: &term.SingleTerm{Begin: "1"},
							}},
						},
					},
					{
						PrefixOp: "-",
						FieldQuery: &FieldQuery{
							Field: &term.Field{Value: []string{"x"}},
							Term: &Term{FuzzyTerm: &term.FuzzyTerm{
								SingleTerm: &term.SingleTerm{Begin: "2"},
							}},
						},
					},
				},
			},
			wantErr: false,
			wantStr: `-x:1 -x:2`,
		},
		{
			name:  "test_(-(x:1 +y:2) z:9)",
			input: `-(x:1 +y:2) z:9`,
			want: &Lucene{
				Clauses: []*PrefixClause{
					{
						PrefixOp: "-",
						ParenQuery: &ParenQuery{
							SubQuery: &Lucene{
								Clauses: []*PrefixClause{
									{
										FieldQuery: &FieldQuery{
											Field: &term.Field{Value: []string{"x"}},
											Term: &Term{FuzzyTerm: &term.FuzzyTerm{
												SingleTerm: &term.SingleTerm{Begin: "1"},
											}},
										},
									},
									{
										PrefixOp: "+",
										FieldQuery: &FieldQuery{
											Field: &term.Field{Value: []string{"y"}},
											Term: &Term{FuzzyTerm: &term.FuzzyTerm{
												SingleTerm: &term.SingleTerm{Begin: "2"},
											}},
										},
									},
								},
							},
						},
					},
					{
						FieldQuery: &FieldQuery{
							Field: &term.Field{Value: []string{"z"}},
							Term: &Term{FuzzyTerm: &term.FuzzyTerm{
								SingleTerm: &term.SingleTerm{Begin: "9"},
							}},
						},
					},
				},
			},
			wantErr: false,
			wantStr: `-( x:1 +y:2 ) z:9`,
		},
		{
			name:  "test_((x:1   -y:2)   +(  -x:8 +k:90))",
			input: `(x:1   -y:2)   +(  -x:8 +k:90)`,
			want: &Lucene{
				Clauses: []*PrefixClause{
					{
						ParenQuery: &ParenQuery{
							SubQuery: &Lucene{
								Clauses: []*PrefixClause{
									{
										FieldQuery: &FieldQuery{
											Field: &term.Field{Value: []string{"x"}},
											Term: &Term{FuzzyTerm: &term.FuzzyTerm{
												SingleTerm: &term.SingleTerm{Begin: "1"},
											}},
										},
									},
									{
										PrefixOp: "-",
										FieldQuery: &FieldQuery{
											Field: &term.Field{Value: []string{"y"}},
											Term: &Term{FuzzyTerm: &term.FuzzyTerm{
												SingleTerm: &term.SingleTerm{Begin: "2"},
											}},
										},
									},
								},
							},
						},
					},
					{
						PrefixOp: "+",
						ParenQuery: &ParenQuery{
							SubQuery: &Lucene{
								Clauses: []*PrefixClause{
									{
										PrefixOp: "-",
										FieldQuery: &FieldQuery{
											Field: &term.Field{Value: []string{"x"}},
											Term: &Term{FuzzyTerm: &term.FuzzyTerm{
												SingleTerm: &term.SingleTerm{Begin: "8"},
											}},
										},
									},
									{
										PrefixOp: "+",
										FieldQuery: &FieldQuery{
											Field: &term.Field{Value: []string{"k"}},
											Term: &Term{FuzzyTerm: &term.FuzzyTerm{
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
			wantErr: false,
			wantStr: `( x:1 -y:2 ) +( -x:8 +k:90 )`,
		},
		{
			name:  "test_(x:(txt foo bar) -x-y:\"xxx\" !zz:iopio\\ 90)",
			input: `x:(txt foo bar) -x-y:"xxx" !zz:iopio\ 90`,
			want: &Lucene{
				Clauses: []*PrefixClause{
					{
						FieldQuery: &FieldQuery{
							Field: &term.Field{Value: []string{"x"}},
							Term: &Term{
								TermGroup: &TermGroup{
									PrefixTermGroup: &PrefixTermGroup{
										PrefixTerms: []*PrefixOperatorTerm{
											{
												FieldTermGroup: &term.FieldTermGroup{
													SingleTerm: &term.SingleTerm{Begin: "txt"},
												},
											},
											{
												FieldTermGroup: &term.FieldTermGroup{
													SingleTerm: &term.SingleTerm{Begin: "foo"},
												},
											},
											{
												FieldTermGroup: &term.FieldTermGroup{
													SingleTerm: &term.SingleTerm{Begin: "bar"},
												},
											},
										},
									},
								},
							},
						},
					},
					{
						PrefixOp: "-",
						FieldQuery: &FieldQuery{
							Field: &term.Field{Value: []string{"x", "-", "y"}},
							Term: &Term{FuzzyTerm: &term.FuzzyTerm{
								PhraseTerm: &term.PhraseTerm{Chars: []string{`xxx`}},
							}},
						},
					},
					{
						PrefixOp: "!",
						FieldQuery: &FieldQuery{
							Field: &term.Field{Value: []string{"zz"}},
							Term: &Term{FuzzyTerm: &term.FuzzyTerm{
								SingleTerm: &term.SingleTerm{Begin: `iopio`, Chars: []string{`\ `, `90`}},
							}},
						},
					},
				},
			},
			wantErr: false,
			wantStr: `x:( txt foo bar ) -x-y:"xxx" !zz:iopio\ 90`,
		},
		{
			name:  "test_((-x:(foo bar))  +z:you)",
			input: `(-x:(foo bar))  +z:you`,
			want: &Lucene{
				Clauses: []*PrefixClause{
					{
						ParenQuery: &ParenQuery{
							SubQuery: &Lucene{
								Clauses: []*PrefixClause{
									{
										PrefixOp: "-",
										FieldQuery: &FieldQuery{
											Field: &term.Field{Value: []string{"x"}},
											Term: &Term{
												TermGroup: &TermGroup{
													PrefixTermGroup: &PrefixTermGroup{
														PrefixTerms: []*PrefixOperatorTerm{
															{
																FieldTermGroup: &term.FieldTermGroup{
																	SingleTerm: &term.SingleTerm{Begin: "foo"},
																},
															},
															{
																FieldTermGroup: &term.FieldTermGroup{
																	SingleTerm: &term.SingleTerm{Begin: "bar"},
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
					{
						PrefixOp: "+",
						FieldQuery: &FieldQuery{
							Field: &term.Field{Value: []string{"z"}},
							Term: &Term{FuzzyTerm: &term.FuzzyTerm{
								SingleTerm: &term.SingleTerm{Begin: `you`},
							}},
						},
					},
				},
			},
			wantErr: false,
			wantStr: `( -x:( foo bar ) ) +z:you`,
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
		input lucene_parser.Query
		qType lucene_parser.QueryType
	}
	for _, tt := range []test{
		{
			name:  "test_lucene_type",
			input: &lucene_parser.Lucene{},
			qType: lucene_parser.LUCENE_QUERY,
		},
		{
			name:  "test_not_query_type",
			input: &PrefixClause{PrefixOp: "-"},
			qType: lucene_parser.NOT_QUERY,
		},
		{
			name:  "test_and_query_type",
			input: &PrefixClause{PrefixOp: "+"},
			qType: lucene_parser.AND_QUERY,
		},
		{
			name:  "test_or_query_type",
			input: &PrefixClause{PrefixOp: ""},
			qType: lucene_parser.OR_QUERY,
		},
		{
			name:  "test_paren_query_type",
			input: &ParenQuery{},
			qType: lucene_parser.PAREN_QUERY,
		},
		{
			name:  "test_field_query_type",
			input: &FieldQuery{},
			qType: lucene_parser.FIELD_QUERY,
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
		input lucene_parser.Query
	}
	for _, tt := range []test{
		{
			name:  "test_lucene_string",
			input: &Lucene{},
		},
		{
			name:  "test_prefix_clause_string",
			input: &PrefixClause{},
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
	var x *PrefixClause
	assert.Equal(t, "", x.String())
	var f *FieldQuery
	assert.Equal(t, "", f.String())
	var p *ParenQuery
	assert.Equal(t, "", p.String())
}
