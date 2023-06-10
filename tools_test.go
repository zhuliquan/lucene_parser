package lucene_parser

import (
	"reflect"
	"testing"

	"github.com/zhuliquan/lucene_parser/operator"
	"github.com/zhuliquan/lucene_parser/term"
)

func TestTermGroupToLucene(t *testing.T) {
	type args struct {
		field     *term.Field
		termGroup *term.TermGroup
	}
	tests := []struct {
		name string
		args args
		want *Lucene
	}{
		{
			name: "test_nil_term_group",
			args: args{field: nil, termGroup: nil},
			want: nil,
		},
		{
			name: "test_nil_term_term_group",
			args: args{field: nil, termGroup: &term.TermGroup{}},
			want: nil,
		},
		{
			name: "test_nil_logic_term_group",
			args: args{field: nil, termGroup: &term.TermGroup{LogicTermGroup: &term.LogicTermGroup{}}},
			want: nil,
		},
		{
			name: "test_nil_or_term_group",
			args: args{field: nil, termGroup: &term.TermGroup{LogicTermGroup: &term.LogicTermGroup{
				OrTermGroup: &term.OrTermGroup{},
			}}},
			want: nil,
		},
		{
			name: "test_nil_and_term_group",
			args: args{field: nil, termGroup: &term.TermGroup{LogicTermGroup: &term.LogicTermGroup{
				OrTermGroup: &term.OrTermGroup{AndTermGroup: &term.AndTermGroup{}},
			}}},
			want: nil,
		},
		{
			name: "test_nil_field_term_group",
			args: args{field: nil, termGroup: &term.TermGroup{LogicTermGroup: &term.LogicTermGroup{
				OrTermGroup: &term.OrTermGroup{AndTermGroup: &term.AndTermGroup{FieldTermGroup: &term.FieldTermGroup{}}},
			}}},
			want: nil,
		},
		{
			name: "test_nil_paren_term_group",
			args: args{field: nil, termGroup: &term.TermGroup{LogicTermGroup: &term.LogicTermGroup{
				OrTermGroup: &term.OrTermGroup{AndTermGroup: &term.AndTermGroup{ParenTermGroup: &term.ParenTermGroup{}}},
			}}},
			want: nil,
		},
		{
			name: "test_single_field_term_group",
			args: args{field: &term.Field{Value: []string{"x"}}, termGroup: &term.TermGroup{LogicTermGroup: &term.LogicTermGroup{
				OrTermGroup: &term.OrTermGroup{AndTermGroup: &term.AndTermGroup{
					FieldTermGroup: &term.FieldTermGroup{
						SingleTerm: &term.SingleTerm{
							Begin: "y",
						},
					},
				}},
			}}},
			want: &Lucene{
				OrQuery: &OrQuery{
					AndQuery: &AndQuery{
						FieldQuery: &FieldQuery{
							Field: &term.Field{Value: []string{"x"}},
							Term: &term.Term{
								FuzzyTerm: &term.FuzzyTerm{
									SingleTerm: &term.SingleTerm{Begin: "y"},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "test_phrase_field_term_group",
			args: args{field: &term.Field{Value: []string{"x"}}, termGroup: &term.TermGroup{LogicTermGroup: &term.LogicTermGroup{
				OrTermGroup: &term.OrTermGroup{AndTermGroup: &term.AndTermGroup{
					FieldTermGroup: &term.FieldTermGroup{
						PhraseTerm: &term.PhraseTerm{
							Chars: []string{"y"},
						},
					},
				}},
			}}},
			want: &Lucene{
				OrQuery: &OrQuery{
					AndQuery: &AndQuery{
						FieldQuery: &FieldQuery{
							Field: &term.Field{Value: []string{"x"}},
							Term: &term.Term{
								FuzzyTerm: &term.FuzzyTerm{
									PhraseTerm: &term.PhraseTerm{Chars: []string{"y"}},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "test_single_range_field_term_group",
			args: args{field: &term.Field{Value: []string{"x"}}, termGroup: &term.TermGroup{LogicTermGroup: &term.LogicTermGroup{
				OrTermGroup: &term.OrTermGroup{AndTermGroup: &term.AndTermGroup{
					FieldTermGroup: &term.FieldTermGroup{
						SRangeTerm: &term.SRangeTerm{
							Symbol: ">",
							Value:  &term.RangeValue{SingleValue: []string{"10"}},
						},
					},
				}},
			}}},
			want: &Lucene{
				OrQuery: &OrQuery{
					AndQuery: &AndQuery{
						FieldQuery: &FieldQuery{
							Field: &term.Field{Value: []string{"x"}},
							Term: &term.Term{
								RangeTerm: &term.RangeTerm{
									SRangeTerm: &term.SRangeTerm{
										Symbol: ">",
										Value:  &term.RangeValue{SingleValue: []string{"10"}},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "test_double_range_field_term_group",
			args: args{field: &term.Field{Value: []string{"x"}}, termGroup: &term.TermGroup{LogicTermGroup: &term.LogicTermGroup{
				OrTermGroup: &term.OrTermGroup{AndTermGroup: &term.AndTermGroup{
					FieldTermGroup: &term.FieldTermGroup{
						DRangeTerm: &term.DRangeTerm{
							LBRACKET: "[",
							LValue:   &term.RangeValue{SingleValue: []string{"10"}},
							RValue:   &term.RangeValue{SingleValue: []string{"20"}},
							RBRACKET: "]",
						},
					},
				}},
			}}},
			want: &Lucene{
				OrQuery: &OrQuery{
					AndQuery: &AndQuery{
						FieldQuery: &FieldQuery{
							Field: &term.Field{Value: []string{"x"}},
							Term: &term.Term{
								RangeTerm: &term.RangeTerm{
									DRangeTerm: &term.DRangeTerm{
										LBRACKET: "[",
										LValue:   &term.RangeValue{SingleValue: []string{"10"}},
										RValue:   &term.RangeValue{SingleValue: []string{"20"}},
										RBRACKET: "]",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "test_paren_term_group",
			args: args{
				field: &term.Field{Value: []string{"x"}},
				termGroup: &term.TermGroup{
					LogicTermGroup: &term.LogicTermGroup{
						OrTermGroup: &term.OrTermGroup{
							AndTermGroup: &term.AndTermGroup{
								ParenTermGroup: &term.ParenTermGroup{
									SubTermGroup: &term.LogicTermGroup{
										OrTermGroup: &term.OrTermGroup{
											AndTermGroup: &term.AndTermGroup{
												FieldTermGroup: &term.FieldTermGroup{
													DRangeTerm: &term.DRangeTerm{
														LBRACKET: "[",
														LValue:   &term.RangeValue{SingleValue: []string{"10"}},
														RValue:   &term.RangeValue{SingleValue: []string{"20"}},
														RBRACKET: "]",
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
			want: &Lucene{
				OrQuery: &OrQuery{
					AndQuery: &AndQuery{
						ParenQuery: &ParenQuery{
							SubQuery: &Lucene{
								OrQuery: &OrQuery{
									AndQuery: &AndQuery{
										FieldQuery: &FieldQuery{
											Field: &term.Field{Value: []string{"x"}},
											Term: &term.Term{
												RangeTerm: &term.RangeTerm{
													DRangeTerm: &term.DRangeTerm{
														LBRACKET: "[",
														LValue:   &term.RangeValue{SingleValue: []string{"10"}},
														RValue:   &term.RangeValue{SingleValue: []string{"20"}},
														RBRACKET: "]",
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
			name: "test_double_and_query",
			args: args{field: &term.Field{Value: []string{"x"}}, termGroup: &term.TermGroup{LogicTermGroup: &term.LogicTermGroup{
				OrTermGroup: &term.OrTermGroup{
					AndTermGroup: &term.AndTermGroup{
						FieldTermGroup: &term.FieldTermGroup{
							SRangeTerm: &term.SRangeTerm{
								Symbol: ">",
								Value:  &term.RangeValue{SingleValue: []string{"10"}},
							},
						},
					},
					AnSTermGroup: []*term.AnSTermGroup{
						{
							AndSymbol: &operator.AndSymbol{Symbol: "AND"},
							AndTermGroup: &term.AndTermGroup{
								FieldTermGroup: &term.FieldTermGroup{
									SRangeTerm: &term.SRangeTerm{
										Symbol: "<",
										Value:  &term.RangeValue{SingleValue: []string{"20"}},
									},
								},
							},
						},
					},
				},
			}}},
			want: &Lucene{
				OrQuery: &OrQuery{
					AndQuery: &AndQuery{
						FieldQuery: &FieldQuery{
							Field: &term.Field{Value: []string{"x"}},
							Term: &term.Term{
								RangeTerm: &term.RangeTerm{
									SRangeTerm: &term.SRangeTerm{
										Symbol: ">",
										Value:  &term.RangeValue{SingleValue: []string{"10"}},
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
									Field: &term.Field{Value: []string{"x"}},
									Term: &term.Term{
										RangeTerm: &term.RangeTerm{
											SRangeTerm: &term.SRangeTerm{
												Symbol: "<",
												Value:  &term.RangeValue{SingleValue: []string{"20"}},
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
			name: "test_double_or_query",
			args: args{field: &term.Field{Value: []string{"x"}}, termGroup: &term.TermGroup{LogicTermGroup: &term.LogicTermGroup{
				OrTermGroup: &term.OrTermGroup{
					AndTermGroup: &term.AndTermGroup{
						FieldTermGroup: &term.FieldTermGroup{
							SRangeTerm: &term.SRangeTerm{
								Symbol: ">",
								Value:  &term.RangeValue{SingleValue: []string{"10"}},
							},
						},
					},
				},
				OSTermGroup: []*term.OSTermGroup{
					{
						OrSymbol: &operator.OrSymbol{Symbol: "OR"},
						OrTermGroup: &term.OrTermGroup{
							AndTermGroup: &term.AndTermGroup{
								FieldTermGroup: &term.FieldTermGroup{
									SRangeTerm: &term.SRangeTerm{
										Symbol: "<",
										Value:  &term.RangeValue{SingleValue: []string{"20"}},
									},
								},
							},
						},
					},
				},
			}}},
			want: &Lucene{
				OrQuery: &OrQuery{
					AndQuery: &AndQuery{
						FieldQuery: &FieldQuery{
							Field: &term.Field{Value: []string{"x"}},
							Term: &term.Term{
								RangeTerm: &term.RangeTerm{
									SRangeTerm: &term.SRangeTerm{
										Symbol: ">",
										Value:  &term.RangeValue{SingleValue: []string{"10"}},
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
									Field: &term.Field{Value: []string{"x"}},
									Term: &term.Term{
										RangeTerm: &term.RangeTerm{
											SRangeTerm: &term.SRangeTerm{
												Symbol: "<",
												Value:  &term.RangeValue{SingleValue: []string{"20"}},
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TermGroupToLucene(tt.args.field, tt.args.termGroup); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TermGroupToLucene() = %v, want %v", got, tt.want)
			}
		})
	}
}
