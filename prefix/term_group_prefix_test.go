package prefix

import (
	"testing"

	"github.com/alecthomas/participle"
	"github.com/stretchr/testify/assert"
	op "github.com/zhuliquan/lucene_parser/operator"
	"github.com/zhuliquan/lucene_parser/term"
	"github.com/zhuliquan/lucene_parser/token"
)

func TestTermGroup(t *testing.T) {
	var termParser = participle.MustBuild(
		&TermGroup{},
		participle.Lexer(token.Lexer),
	)

	type testCase struct {
		name     string
		input    string
		want     *TermGroup
		boost    term.BoostValue
		termType term.TermType
	}
	var testCases = []testCase{
		{
			name:  "term_four_term_group",
			input: `( +8908  "dsada 78" +"89080  xxx" -"xx yyyy" +\+dsada\ 7897 -\-\-dsada\-7897  [-1 TO 3] +>2021-11-04 +<2021-11-11 +(  ![-1 TO 3] [1 TO 2] +[5 TO 10} -dsadad\ dsad\+789 +"dsad xx") +{8 TO 90])^1.8`,
			want: &TermGroup{
				PrefixTermGroup: &PrefixTermGroup{
					PrefixTerms: []*PrefixOperatorTerm{
						{PrefixOp: "+", FieldTermGroup: &term.FieldTermGroup{SingleTerm: &term.SingleTerm{Begin: `8908`}}},
						{PrefixOp: "", FieldTermGroup: &term.FieldTermGroup{PhraseTerm: &term.PhraseTerm{Chars: []string{`dsada`, ` `, `78`}}}},
						{PrefixOp: "+", FieldTermGroup: &term.FieldTermGroup{PhraseTerm: &term.PhraseTerm{Chars: []string{`89080`, `  `, `xxx`}}}},
						{PrefixOp: "-", FieldTermGroup: &term.FieldTermGroup{PhraseTerm: &term.PhraseTerm{Chars: []string{`xx`, ` `, `yyyy`}}}},
						{PrefixOp: "+", FieldTermGroup: &term.FieldTermGroup{SingleTerm: &term.SingleTerm{Begin: `\+`, Chars: []string{`dsada`, `\ `, `7897`}}}},
						{PrefixOp: "-", FieldTermGroup: &term.FieldTermGroup{SingleTerm: &term.SingleTerm{Begin: `\-\-`, Chars: []string{`dsada`, `\-`, `7897`}}}},
						{PrefixOp: "", FieldTermGroup: &term.FieldTermGroup{DRangeTerm: &term.DRangeTerm{LBRACKET: "[", LValue: &term.RangeValue{SingleValue: []string{"-", "1"}}, RValue: &term.RangeValue{SingleValue: []string{"3"}}, RBRACKET: "]"}}},
						{PrefixOp: "+", FieldTermGroup: &term.FieldTermGroup{SRangeTerm: &term.SRangeTerm{Symbol: ">", Value: &term.RangeValue{SingleValue: []string{`2021`, "-", "11", "-", "04"}}}}},
						{PrefixOp: "+", FieldTermGroup: &term.FieldTermGroup{SRangeTerm: &term.SRangeTerm{Symbol: "<", Value: &term.RangeValue{SingleValue: []string{`2021`, "-", "11", "-", "11"}}}}},
						{PrefixOp: "+", ParenTermGroup: &PrefixTermGroup{
							PrefixTerms: []*PrefixOperatorTerm{
								{PrefixOp: "!", FieldTermGroup: &term.FieldTermGroup{DRangeTerm: &term.DRangeTerm{LBRACKET: "[", LValue: &term.RangeValue{SingleValue: []string{"-", "1"}}, RValue: &term.RangeValue{SingleValue: []string{"3"}}, RBRACKET: "]"}}},
								{PrefixOp: "", FieldTermGroup: &term.FieldTermGroup{DRangeTerm: &term.DRangeTerm{LBRACKET: "[", LValue: &term.RangeValue{SingleValue: []string{"1"}}, RValue: &term.RangeValue{SingleValue: []string{"2"}}, RBRACKET: "]"}}},
								{PrefixOp: "+", FieldTermGroup: &term.FieldTermGroup{DRangeTerm: &term.DRangeTerm{LBRACKET: "[", LValue: &term.RangeValue{SingleValue: []string{"5"}}, RValue: &term.RangeValue{SingleValue: []string{"10"}}, RBRACKET: "}"}}},
								{PrefixOp: "-", FieldTermGroup: &term.FieldTermGroup{SingleTerm: &term.SingleTerm{Begin: "dsadad", Chars: []string{`\ `, `dsad`, `\+`, `789`}}}},
								{PrefixOp: "+", FieldTermGroup: &term.FieldTermGroup{PhraseTerm: &term.PhraseTerm{Chars: []string{`dsad`, ` `, `xx`}}}},
							},
						}},
						{PrefixOp: "+", FieldTermGroup: &term.FieldTermGroup{DRangeTerm: &term.DRangeTerm{LBRACKET: "{", LValue: &term.RangeValue{SingleValue: []string{"8"}}, RValue: &term.RangeValue{SingleValue: []string{"90"}}, RBRACKET: "]"}}},
					},
				},
				BoostSymbol: "^1.8",
			},
			boost:    term.BoostValue(1.8),
			termType: term.GROUP_TERM_TYPE | term.BOOST_TERM_TYPE,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var out = &TermGroup{}
			err := termParser.ParseString(tt.input, out)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, out)
			assert.Equal(t, tt.boost, out.Boost())
			assert.Equal(t, tt.termType, out.GetTermType())
		})
	}
}

func TestPrefixOperatorTerm(t *testing.T) {
	var tt *PrefixOperatorTerm
	assert.Equal(t, "", tt.String())
	tt = &PrefixOperatorTerm{}
	assert.Equal(t, op.SHOULD_PREFIX_TYPE, tt.GetPrefixType())
	assert.Equal(t, "", tt.String())
	tt = &PrefixOperatorTerm{PrefixOp: "+"}
	assert.Equal(t, op.MUST_PREFIX_TYPE, tt.GetPrefixType())
	assert.Equal(t, "", tt.String())
	tt = &PrefixOperatorTerm{PrefixOp: "-"}
	assert.Equal(t, op.MUST_NOT_PREFIX_TYPE, tt.GetPrefixType())
	assert.Equal(t, "", tt.String())
}

func TestPrefixTermGroup(t *testing.T) {
	var tt *PrefixTermGroup
	assert.Equal(t, "", tt.String())
	assert.Equal(t, term.UNKNOWN_TERM_TYPE, tt.GetTermType())
	tt = &PrefixTermGroup{}
	assert.Equal(t, "", tt.String())
	assert.Equal(t, term.UNKNOWN_TERM_TYPE, tt.GetTermType())
}
