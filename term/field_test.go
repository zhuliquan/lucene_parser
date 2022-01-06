package term

import (
	"reflect"
	"testing"

	"github.com/alecthomas/participle"
	"github.com/zhuliquan/lucene_parser/token"
)

func TestField(t *testing.T) {
	var termParser = participle.MustBuild(
		&Field{},
		participle.Lexer(token.Lexer),
	)

	type testCase struct {
		name  string
		input string
		want  *Field
	}

	var testCases = []testCase{
		{
			name:  "TestField01",
			input: `1\+1`,
			want:  &Field{Value: []string{`1`, `\+`, `1`}},
		},
		{
			name:  "TestField02",
			input: `ls`,
			want:  &Field{Value: []string{`ls`}},
		},
		{
			name:  "TestField03",
			input: `x.y`,
			want:  &Field{Value: []string{`x`, `.`, `y`}},
		},
		{
			name:  "TestField04",
			input: `x.y-z`,
			want:  &Field{Value: []string{`x`, `.`, `y`, `-`, `z`}},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var out = &Field{}
			if err := termParser.ParseString(tt.input, out); err != nil {
				t.Errorf("failed to parse input: %s, err: %+v", tt.input, err)
			} else if !reflect.DeepEqual(tt.want, out) {
				t.Errorf("termParser.ParseString( %s ) = %+v, want: %+v", tt.input, out, tt.want)
			}
		})
	}

}
