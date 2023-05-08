package term

import (
	"testing"

	"github.com/alecthomas/participle"
	"github.com/stretchr/testify/assert"
	"github.com/zhuliquan/lucene_parser/token"
)

func TestSingleTerm(t *testing.T) {
	var termParser = participle.MustBuild(
		&SingleTerm{},
		participle.Lexer(token.Lexer),
	)

	type testCase struct {
		name     string
		input    string
		want     *SingleTerm
		values   string
		wildcard bool
	}
	var testCases = []testCase{
		{
			name:     "TestSimpleTerm01",
			input:    `\/dsada\/\ dasda80980?`,
			want:     &SingleTerm{Begin: `\/`, Chars: []string{`dsada`, `\/\ `, `dasda`, `80980`, `?`}},
			values:   `\/dsada\/\ dasda80980?`,
			wildcard: true,
		},
		{
			name:     "TestSimpleTerm02",
			input:    `\/dsada\/\ dasda80980*`,
			want:     &SingleTerm{Begin: `\/`, Chars: []string{`dsada`, `\/\ `, `dasda`, `80980`, `*`}},
			values:   `\/dsada\/\ dasda80980*`,
			wildcard: true,
		},
		{
			name:     "TestSimpleTerm03",
			input:    `\/dsada\/\ dasda8\?0980\*`,
			want:     &SingleTerm{Begin: `\/`, Chars: []string{`dsada`, `\/\ `, `dasda`, `8`, `\?`, `0980`, `\*`}},
			values:   `\/dsada\/\ dasda8\?0980\*`,
			wildcard: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var out = &SingleTerm{}
			err := termParser.ParseString(tt.input, out)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, out)
			assert.Equal(t, tt.values, out.String())
			assert.Equal(t, tt.wildcard, out.haveWildcard())
		})
	}
	var s *SingleTerm
	assert.Empty(t, s.String())
	assert.False(t, s.haveWildcard())
	assert.Equal(t, UNKNOWN_TERM_TYPE, s.GetTermType())
	_, err := s.Value(func(s string) (interface{}, error) { return s, nil })
	assert.Equal(t, ErrEmptySingleTerm, err)
}

func TestPhraseTerm(t *testing.T) {
	var termParser = participle.MustBuild(
		&PhraseTerm{},
		participle.Lexer(token.Lexer),
	)

	type testCase struct {
		name     string
		input    string
		want     *PhraseTerm
		values   string
		wildcard bool
	}
	var testCases = []testCase{
		{
			name:     "TestPhraseTerm01",
			input:    `"dsada 78"`,
			want:     &PhraseTerm{Chars: []string{`dsada`, ` `, `78`}},
			values:   `"dsada 78"`,
			wildcard: false,
		},
		{
			name:     "TestPhraseTerm02",
			input:    `"*dsada 78"`,
			want:     &PhraseTerm{Chars: []string{`*`, `dsada`, ` `, `78`}},
			values:   `"*dsada 78"`,
			wildcard: true,
		},
		{
			name:     "TestPhraseTerm03",
			input:    `"?dsada 78"`,
			want:     &PhraseTerm{Chars: []string{`?`, `dsada`, ` `, `78`}},
			values:   `"?dsada 78"`,
			wildcard: true,
		},
		{
			name:     "TestPhraseTerm04",
			input:    `"dsada* 78"`,
			want:     &PhraseTerm{Chars: []string{`dsada`, `*`, ` `, `78`}},
			values:   `"dsada* 78"`,
			wildcard: true,
		},
		{
			name:     "TestPhraseTerm05",
			input:    `"dsada? 78"`,
			want:     &PhraseTerm{Chars: []string{`dsada`, `?`, ` `, `78`}},
			values:   `"dsada? 78"`,
			wildcard: true,
		},
		{
			name:     "TestPhraseTerm06",
			input:    `"dsada\* 78"`,
			want:     &PhraseTerm{Chars: []string{`dsada`, `\*`, ` `, `78`}},
			values:   `"dsada\* 78"`,
			wildcard: false,
		},
		{
			name:     "TestPhraseTerm07",
			input:    `"dsada\? 78"`,
			want:     &PhraseTerm{Chars: []string{`dsada`, `\?`, ` `, `78`}},
			values:   `"dsada\? 78"`,
			wildcard: false,
		},
		{
			name:     "TestPhraseTerm09",
			input:    `"\*dsada 78"`,
			want:     &PhraseTerm{Chars: []string{`\*`, `dsada`, ` `, `78`}},
			values:   `"\*dsada 78"`,
			wildcard: false,
		},
		{
			name:     "TestPhraseTerm10",
			input:    `"\?dsada 78"`,
			want:     &PhraseTerm{Chars: []string{`\?`, `dsada`, ` `, `78`}},
			values:   `"\?dsada 78"`,
			wildcard: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var out = &PhraseTerm{}
			err := termParser.ParseString(tt.input, out)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, out)
			assert.Equal(t, tt.values, out.String())
			assert.Equal(t, tt.wildcard, out.haveWildcard())
		})
	}
	var s *PhraseTerm
	assert.Empty(t, s.String())
	assert.False(t, s.haveWildcard())
	assert.Equal(t, UNKNOWN_TERM_TYPE, s.GetTermType())
	_, err := s.Value(func(s string) (interface{}, error) { return s, nil })
	assert.Equal(t, ErrEmptyPhraseTerm, err)
}

func TestRegexpTerm(t *testing.T) {
	var termParser = participle.MustBuild(
		&RegexpTerm{},
		participle.Lexer(token.Lexer),
	)

	type testCase struct {
		name  string
		input string
		want  *RegexpTerm
	}
	var testCases = []testCase{
		{
			name:  "RegexpTerm01",
			input: `/dsada 78/`,
			want:  &RegexpTerm{Chars: []string{`dsada`, ` `, `78`}},
		},
		{
			name:  "RegexpTerm02",
			input: `/\d+\/\d+\.\d+.+/`,
			want:  &RegexpTerm{Chars: []string{`\`, `d`, `+`, `\/`, `\`, `d`, `+`, `\`, `.`, `\`, `d`, `+`, `.`, `+`}},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var out = &RegexpTerm{}
			err := termParser.ParseString(tt.input, out)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, out)
		})
	}
	var s *RegexpTerm
	assert.Empty(t, s.String())
	assert.Equal(t, UNKNOWN_TERM_TYPE, s.GetTermType())
	_, err := s.Value(func(s string) (interface{}, error) { return s, nil })
	assert.Equal(t, ErrEmptyRegexpTerm, err)
}

func TestDRangeTerm(t *testing.T) {
	var termParser = participle.MustBuild(
		&DRangeTerm{},
		participle.Lexer(token.Lexer),
	)

	type testCase struct {
		name  string
		input string
		want  *DRangeTerm
		bound *Bound
		lninf bool
		lpinf bool
		rninf bool
		rpinf bool
	}
	var testCases = []testCase{
		{
			name:  "DRangeTerm01",
			input: `[1 TO 2]`,
			want: &DRangeTerm{
				LBRACKET: "[",
				LValue:   &RangeValue{SingleValue: []string{"1"}},
				RValue:   &RangeValue{SingleValue: []string{"2"}},
				RBRACKET: "]",
			},
			bound: &Bound{
				LeftValue:    &RangeValue{SingleValue: []string{"1"}, flag: false},
				RightValue:   &RangeValue{SingleValue: []string{"2"}, flag: true},
				LeftInclude:  true,
				RightInclude: true,
			},
			lninf: false,
			lpinf: false,
			rninf: false,
			rpinf: false,
		},
		{
			name:  "DRangeTerm02",
			input: `[1 TO 2 }`,
			want: &DRangeTerm{
				LBRACKET: "[",
				LValue:   &RangeValue{SingleValue: []string{"1"}},
				RValue:   &RangeValue{SingleValue: []string{"2"}},
				RBRACKET: "}",
			},
			bound: &Bound{
				LeftValue:    &RangeValue{SingleValue: []string{"1"}, flag: false},
				RightValue:   &RangeValue{SingleValue: []string{"2"}, flag: true},
				LeftInclude:  true,
				RightInclude: false,
			},
			lninf: false,
			lpinf: false,
			rninf: false,
			rpinf: false,
		},
		{
			name:  `DRangeTerm03`,
			input: `{ 1 TO 2}`,
			want: &DRangeTerm{
				LBRACKET: "{",
				LValue:   &RangeValue{SingleValue: []string{"1"}},
				RValue:   &RangeValue{SingleValue: []string{"2"}},
				RBRACKET: "}",
			},
			bound: &Bound{
				LeftValue:    &RangeValue{SingleValue: []string{"1"}, flag: false},
				RightValue:   &RangeValue{SingleValue: []string{"2"}, flag: true},
				LeftInclude:  false,
				RightInclude: false,
			},
			lninf: false,
			lpinf: false,
			rninf: false,
			rpinf: false,
		},
		{
			name:  `DRangeTerm04`,
			input: `{ 1 TO 2]`,
			want: &DRangeTerm{
				LBRACKET: "{",
				LValue:   &RangeValue{SingleValue: []string{"1"}},
				RValue:   &RangeValue{SingleValue: []string{"2"}},
				RBRACKET: "]",
			},
			bound: &Bound{
				LeftValue:    &RangeValue{SingleValue: []string{"1"}, flag: false},
				RightValue:   &RangeValue{SingleValue: []string{"2"}, flag: true},
				LeftInclude:  false,
				RightInclude: true,
			},
			lninf: false,
			lpinf: false,
			rninf: false,
			rpinf: false,
		},
		{
			name:  `DRangeTerm05`,
			input: `[10 TO *]`,
			want: &DRangeTerm{
				LBRACKET: "[",
				LValue:   &RangeValue{SingleValue: []string{"10"}},
				RValue:   &RangeValue{InfinityVal: "*"},
				RBRACKET: "]",
			},
			bound: &Bound{
				LeftValue:    &RangeValue{SingleValue: []string{"10"}, flag: false},
				RightValue:   &RangeValue{InfinityVal: "*", flag: true},
				LeftInclude:  true,
				RightInclude: false,
			},
			lninf: false,
			lpinf: false,
			rninf: false,
			rpinf: true,
		},
		{
			name:  `DRangeTerm06`,
			input: `{* TO 2012-01-01}`,
			want: &DRangeTerm{
				LBRACKET: "{",
				LValue:   &RangeValue{InfinityVal: "*"},
				RValue:   &RangeValue{SingleValue: []string{"2012", "-", "01", "-", "01"}},
				RBRACKET: "}",
			},
			bound: &Bound{
				LeftValue:    &RangeValue{InfinityVal: "*", flag: false},
				RightValue:   &RangeValue{SingleValue: []string{"2012", "-", "01", "-", "01"}, flag: true},
				LeftInclude:  false,
				RightInclude: false,
			},
			lninf: true,
			lpinf: false,
			rninf: false,
			rpinf: false,
		},
		{
			name:  `DRangeTerm07`,
			input: `{* TO "2012-01-01 09:08:16"}`,
			want: &DRangeTerm{
				LBRACKET: "{",
				LValue:   &RangeValue{InfinityVal: "*"},
				RValue:   &RangeValue{PhraseValue: []string{"2012", "-", "01", "-", "01", " ", "09", ":", "08", ":", "16"}},
				RBRACKET: "}",
			},
			bound: &Bound{
				LeftValue:    &RangeValue{InfinityVal: "*", flag: false},
				RightValue:   &RangeValue{PhraseValue: []string{"2012", "-", "01", "-", "01", " ", "09", ":", "08", ":", "16"}, flag: true},
				LeftInclude:  false,
				RightInclude: false,
			},
			lninf: true,
			lpinf: false,
			rninf: false,
			rpinf: false,
		},
		{
			name:  `DRangeTerm08`,
			input: `{* TO 2012/01/01T09:08.16||8d/M }`,
			want: &DRangeTerm{
				LBRACKET: "{",
				LValue:   &RangeValue{InfinityVal: "*"},
				RValue:   &RangeValue{SingleValue: []string{"2012", "/", "01", "/", "01", "T", "09", ":", "08", ".", "16", "|", "|", "8", "d", "/", "M"}},
				RBRACKET: "}",
			},
			bound: &Bound{
				LeftValue:    &RangeValue{InfinityVal: "*", flag: false},
				RightValue:   &RangeValue{SingleValue: []string{"2012", "/", "01", "/", "01", "T", "09", ":", "08", ".", "16", "|", "|", "8", "d", "/", "M"}, flag: true},
				LeftInclude:  false,
				RightInclude: false,
			},
			lninf: true,
			lpinf: false,
			rninf: false,
			rpinf: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var out = &DRangeTerm{}
			err := termParser.ParseString(tt.input, out)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, out)
			assert.Equal(t, tt.bound, out.GetBound())
			assert.Equal(t, RANGE_TERM_TYPE, out.GetTermType())
			var b = out.GetBound()
			assert.Equal(t, tt.lninf, b.LeftValue.IsInf(-1))
			assert.Equal(t, tt.lpinf, b.LeftValue.IsInf(1))
			assert.Equal(t, tt.rninf, b.RightValue.IsInf(-1))
			assert.Equal(t, tt.rpinf, b.RightValue.IsInf(1))
		})
	}
	var s *DRangeTerm
	assert.Empty(t, s.String())
	assert.Equal(t, UNKNOWN_TERM_TYPE, s.GetTermType())
	assert.Nil(t, s.GetBound())
	s = &DRangeTerm{}
	assert.Empty(t, s.String())
	assert.Equal(t, UNKNOWN_TERM_TYPE, s.GetTermType())
	assert.Nil(t, s.GetBound())
}

func TestSRangeTerm(t *testing.T) {
	var termParser = participle.MustBuild(
		&SRangeTerm{},
		participle.Lexer(token.Lexer),
	)

	type testCase struct {
		name  string
		input string
		want  *SRangeTerm
		bound *Bound
		lninf bool
		lpinf bool
		rninf bool
		rpinf bool
	}
	var testCases = []testCase{
		{
			name:  "SRangeTerm01",
			input: `<="dsada\455 78"`,
			want:  &SRangeTerm{Symbol: "<=", Value: &RangeValue{PhraseValue: []string{`dsada`, `\`, `455`, ` `, `78`}}},
			bound: &Bound{
				LeftValue:    &RangeValue{InfinityVal: "*", flag: false},
				RightValue:   &RangeValue{PhraseValue: []string{`dsada`, `\`, `455`, ` `, `78`}, flag: true},
				LeftInclude:  false,
				RightInclude: true,
			},
			lninf: true,
			lpinf: false,
			rninf: false,
			rpinf: false,
		},
		{
			name:  "SRangeTerm02",
			input: `<"dsada\\ 78"`,
			want:  &SRangeTerm{Symbol: "<", Value: &RangeValue{PhraseValue: []string{`dsada`, `\\`, ` `, `78`}}},
			bound: &Bound{
				LeftValue:    &RangeValue{InfinityVal: "*", flag: false},
				RightValue:   &RangeValue{PhraseValue: []string{`dsada`, `\\`, ` `, `78`}, flag: true},
				LeftInclude:  false,
				RightInclude: false,
			},
			lninf: true,
			lpinf: false,
			rninf: false,
			rpinf: false,
		},
		{
			name:  "SRangeTerm03",
			input: `>=dsada\ 78`,
			want:  &SRangeTerm{Symbol: ">=", Value: &RangeValue{SingleValue: []string{`dsada`, `\ `, `78`}}},
			bound: &Bound{
				LeftValue:    &RangeValue{SingleValue: []string{`dsada`, `\ `, `78`}, flag: false},
				RightValue:   &RangeValue{InfinityVal: "*", flag: true},
				LeftInclude:  true,
				RightInclude: false,
			},
			lninf: false,
			lpinf: false,
			rninf: false,
			rpinf: true,
		},
		{
			name:  "SRangeTerm04",
			input: `>dsada\ 78`,
			want:  &SRangeTerm{Symbol: ">", Value: &RangeValue{SingleValue: []string{`dsada`, `\ `, `78`}}},
			bound: &Bound{
				LeftValue:    &RangeValue{SingleValue: []string{`dsada`, `\ `, `78`}, flag: false},
				RightValue:   &RangeValue{InfinityVal: "*", flag: true},
				LeftInclude:  false,
				RightInclude: false,
			},
			lninf: false,
			lpinf: false,
			rninf: false,
			rpinf: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var out = &SRangeTerm{}
			err := termParser.ParseString(tt.input, out)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, out)
			assert.Equal(t, tt.bound, out.GetBound())
			assert.Equal(t, RANGE_TERM_TYPE, out.GetTermType())
			var b = out.GetBound()
			assert.Equal(t, tt.lninf, b.LeftValue.IsInf(-1))
			assert.Equal(t, tt.lpinf, b.LeftValue.IsInf(1))
			assert.Equal(t, tt.rninf, b.RightValue.IsInf(-1))
			assert.Equal(t, tt.rpinf, b.RightValue.IsInf(1))
		})
	}
	var s *SRangeTerm
	assert.Empty(t, s.String())
	assert.Equal(t, UNKNOWN_TERM_TYPE, s.GetTermType())
	assert.Nil(t, s.GetBound())
	s = &SRangeTerm{}
	assert.Empty(t, s.String())
	assert.Equal(t, UNKNOWN_TERM_TYPE, s.GetTermType())
	assert.Nil(t, s.GetBound())
}
