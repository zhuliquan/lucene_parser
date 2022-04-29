package term

import (
	"reflect"
	"testing"

	"github.com/alecthomas/participle"
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
		wildward bool
	}
	var testCases = []testCase{
		{
			name:     "TestSimpleTerm01",
			input:    `\/dsada\/\ dasda80980?`,
			want:     &SingleTerm{Begin: `\/dsada\/\ dasda`, Chars: []string{`80980`, `?`}},
			values:   `\/dsada\/\ dasda80980?`,
			wildward: true,
		},
		{
			name:     "TestSimpleTerm02",
			input:    `\/dsada\/\ dasda80980*`,
			want:     &SingleTerm{Begin: `\/dsada\/\ dasda`, Chars: []string{`80980`, `*`}},
			values:   `\/dsada\/\ dasda80980*`,
			wildward: true,
		},
		{
			name:     "TestSimpleTerm03",
			input:    `\/dsada\/\ dasda8\?0980\*`,
			want:     &SingleTerm{Begin: `\/dsada\/\ dasda`, Chars: []string{`8`, `\?`, `0980`, `\*`}},
			values:   `\/dsada\/\ dasda8\?0980\*`,
			wildward: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var out = &SingleTerm{}
			if err := termParser.ParseString(tt.input, out); err != nil {
				t.Errorf("failed to parse input: %s, err: %+v", tt.input, err)
			} else if !reflect.DeepEqual(tt.want, out) {
				t.Errorf("termParser.ParseString( %s ) = %+v, want: %+v", tt.input, out, tt.want)
			} else if tt.values != out.String() {
				t.Errorf("expect get values: %s, but get values: %+v", tt.values, out.String())
			} else if tt.wildward != out.haveWildcard() {
				t.Errorf("expect get wildcard: %+v, but get wildcard: %+v", tt.wildward, out.haveWildcard())
			}
		})
	}
	var s *SingleTerm
	if s.String() != "" {
		t.Errorf("expect empty")
	}
	if s.haveWildcard() {
		t.Errorf("expect no wildcard")
	}
	if s.GetTermType() != UNKNOWN_TERM_TYPE {
		t.Errorf("expect unknown term type")
	}
	if _, err := s.Value(func(s string) (interface{}, error) { return s, nil }); err != ErrEmptySingleTerm {
		t.Errorf("expect empty single term")
	}
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
		wildward bool
	}
	var testCases = []testCase{
		{
			name:     "TestPhraseTerm01",
			input:    `"dsada 78"`,
			want:     &PhraseTerm{Chars: []string{`dsada`, ` `, `78`}},
			values:   `"dsada 78"`,
			wildward: false,
		},
		{
			name:     "TestPhraseTerm02",
			input:    `"*dsada 78"`,
			want:     &PhraseTerm{Chars: []string{`*`, `dsada`, ` `, `78`}},
			values:   `"*dsada 78"`,
			wildward: true,
		},
		{
			name:     "TestPhraseTerm03",
			input:    `"?dsada 78"`,
			want:     &PhraseTerm{Chars: []string{`?`, `dsada`, ` `, `78`}},
			values:   `"?dsada 78"`,
			wildward: true,
		},
		{
			name:     "TestPhraseTerm04",
			input:    `"dsada* 78"`,
			want:     &PhraseTerm{Chars: []string{`dsada`, `*`, ` `, `78`}},
			values:   `"dsada* 78"`,
			wildward: true,
		},
		{
			name:     "TestPhraseTerm05",
			input:    `"dsada? 78"`,
			want:     &PhraseTerm{Chars: []string{`dsada`, `?`, ` `, `78`}},
			values:   `"dsada? 78"`,
			wildward: true,
		},
		{
			name:     "TestPhraseTerm06",
			input:    `"dsada\* 78"`,
			want:     &PhraseTerm{Chars: []string{`dsada\*`, ` `, `78`}},
			values:   `"dsada\* 78"`,
			wildward: false,
		},
		{
			name:     "TestPhraseTerm07",
			input:    `"dsada\? 78"`,
			want:     &PhraseTerm{Chars: []string{`dsada\?`, ` `, `78`}},
			values:   `"dsada\? 78"`,
			wildward: false,
		},
		{
			name:     "TestPhraseTerm09",
			input:    `"\*dsada 78"`,
			want:     &PhraseTerm{Chars: []string{`\*dsada`, ` `, `78`}},
			values:   `"\*dsada 78"`,
			wildward: false,
		},
		{
			name:     "TestPhraseTerm10",
			input:    `"\?dsada 78"`,
			want:     &PhraseTerm{Chars: []string{`\?dsada`, ` `, `78`}},
			values:   `"\?dsada 78"`,
			wildward: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var out = &PhraseTerm{}
			if err := termParser.ParseString(tt.input, out); err != nil {
				t.Errorf("failed to parse input: %s, err: %+v", tt.input, err)
			} else if !reflect.DeepEqual(tt.want, out) {
				t.Errorf("phraseTermParser.ParseString( %s ) = %+v, want: %+v", tt.input, out, tt.want)
			} else if tt.values != out.String() {
				t.Errorf("expect get values: %s, but get values: %+v", tt.values, out.String())
			} else if tt.wildward != out.haveWildcard() {
				t.Errorf("expect get wildcard: %+v, but get wildcard: %+v", tt.wildward, out.haveWildcard())
			}
		})
	}
	var s *PhraseTerm
	if s.String() != "" {
		t.Errorf("expect empty")
	}
	if s.haveWildcard() {
		t.Errorf("expect no wildcard")
	}
	if s.GetTermType() != UNKNOWN_TERM_TYPE {
		t.Errorf("expect unknown term type")
	}
	if _, err := s.Value(func(s string) (interface{}, error) { return s, nil }); err != ErrEmptyPhraseTerm {
		t.Errorf("expect empty phrase term")
	}
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
			if err := termParser.ParseString(tt.input, out); err != nil {
				t.Errorf("failed to parse input: %s, err: %+v", tt.input, err)
			} else if !reflect.DeepEqual(tt.want, out) {
				t.Errorf("regexpTermParser.ParseString( %s ) = %+v, want: %+v", tt.input, out, tt.want)
			}
		})
	}
	var s *RegexpTerm
	if s.String() != "" {
		t.Errorf("expect empty")
	}
	if s.GetTermType() != UNKNOWN_TERM_TYPE {
		t.Errorf("expect unknown term type")
	}
	if _, err := s.Value(func(s string) (interface{}, error) { return s, nil }); err != ErrEmptyRegexpTerm {
		t.Errorf("expect empty regexp term")
	}
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
			if err := termParser.ParseString(tt.input, out); err != nil {
				t.Errorf("failed to parse input: %s, err: %+v", tt.input, err)
			} else if !reflect.DeepEqual(tt.want, out) {
				t.Errorf("rangeTermParser.ParseString( %s ) = %+v, want: %+v", tt.input, out, tt.want)
			} else if !reflect.DeepEqual(tt.bound, out.GetBound()) {
				t.Errorf("expect get bound: %+v, but get bound: %+v", tt.bound, out.GetBound())
			} else if out.GetTermType() != RANGE_TERM_TYPE {
				t.Errorf("expect range term")
			}
			var b = out.GetBound()
			if b.LeftValue.IsInf(-1) != tt.lninf {
				t.Errorf("expect %v, but %v", tt.lninf, b.LeftValue.IsInf(-1))
			}
			if b.LeftValue.IsInf(1) != tt.lpinf {
				t.Errorf("expect %v, but %v", tt.lninf, b.LeftValue.IsInf(1))
			}
			if b.RightValue.IsInf(-1) != tt.rninf {
				t.Errorf("expect %v, but %v", tt.lninf, b.RightValue.IsInf(-1))
			}
			if b.RightValue.IsInf(1) != tt.rpinf {
				t.Errorf("expect %v, but %v", tt.lninf, b.RightValue.IsInf(1))
			}
		})
	}
	var s *DRangeTerm
	if s.String() != "" {
		t.Errorf("expect empty")
	}
	if s.GetTermType() != UNKNOWN_TERM_TYPE {
		t.Errorf("expect unknown term type")
	}
	if s.GetBound() != nil {
		t.Errorf("expect empty bound")
	}
	s = &DRangeTerm{}
	if s.String() != "" {
		t.Errorf("expect empty")
	}
	if s.GetTermType() != UNKNOWN_TERM_TYPE {
		t.Errorf("expect unknown term type")
	}
	if s.GetBound() != nil {
		t.Errorf("expect empty bound")
	}
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
			want:  &SRangeTerm{Symbol: "<", Value: &RangeValue{PhraseValue: []string{`dsada\\`, ` `, `78`}}},
			bound: &Bound{
				LeftValue:    &RangeValue{InfinityVal: "*", flag: false},
				RightValue:   &RangeValue{PhraseValue: []string{`dsada\\`, ` `, `78`}, flag: true},
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
			want:  &SRangeTerm{Symbol: ">=", Value: &RangeValue{SingleValue: []string{`dsada\ `, `78`}}},
			bound: &Bound{
				LeftValue:    &RangeValue{SingleValue: []string{`dsada\ `, `78`}, flag: false},
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
			want:  &SRangeTerm{Symbol: ">", Value: &RangeValue{SingleValue: []string{`dsada\ `, `78`}}},
			bound: &Bound{
				LeftValue:    &RangeValue{SingleValue: []string{`dsada\ `, `78`}, flag: false},
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
			if err := termParser.ParseString(tt.input, out); err != nil {
				t.Errorf("failed to parse input: %s, err: %+v", tt.input, err)
			} else if !reflect.DeepEqual(tt.want, out) {
				t.Errorf("rangesTermParser.ParseString( %s ) = %+v, want: %+v", tt.input, out, tt.want)
			} else if !reflect.DeepEqual(tt.bound, out.GetBound()) {
				t.Errorf("expect get bound: %+v, but get bound: %+v", tt.bound, out.GetBound())
			} else if out.GetTermType() != RANGE_TERM_TYPE {
				t.Errorf("expect range term")
			}
			var b = out.GetBound()
			if b.LeftValue.IsInf(-1) != tt.lninf {
				t.Errorf("expect %v, but %v", tt.lninf, b.LeftValue.IsInf(-1))
			}
			if b.LeftValue.IsInf(1) != tt.lpinf {
				t.Errorf("expect %v, but %v", tt.lninf, b.LeftValue.IsInf(1))
			}
			if b.RightValue.IsInf(-1) != tt.rninf {
				t.Errorf("expect %v, but %v", tt.lninf, b.RightValue.IsInf(-1))
			}
			if b.RightValue.IsInf(1) != tt.rpinf {
				t.Errorf("expect %v, but %v", tt.lninf, b.RightValue.IsInf(1))
			}
		})
	}
	var s *SRangeTerm
	if s.String() != "" {
		t.Errorf("expect empty")
	}
	if s.GetTermType() != UNKNOWN_TERM_TYPE {
		t.Errorf("expect unknown term type")
	}
	if s.GetBound() != nil {
		t.Errorf("expect empty bound")
	}
	s = &SRangeTerm{}
	if s.String() != "" {
		t.Errorf("expect empty")
	}
	if s.GetTermType() != UNKNOWN_TERM_TYPE {
		t.Errorf("expect unknown term type")
	}
	if s.GetBound() != nil {
		t.Errorf("expect empty bound")
	}
}
