package term

import (
	"strings"
)

type Field struct {
	Value []string `parser:"@(IDENT|ESCAPE|MINUS|NUMBER|DOT)+"`
}

func (f *Field) String() string {
	if f == nil {
		return ""
	} else {
		return strings.Join(f.Value, "")
	}
}
