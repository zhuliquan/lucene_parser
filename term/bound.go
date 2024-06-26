package term

import (
	"strings"
)

// range bound like this [1, 2] [1, 2) (1, 2] (1, 2)
type Bound struct {
	LeftValue    *RangeValue `json:"left_value,omitempty"`
	RightValue   *RangeValue `json:"right_Value,omitempty"`
	LeftInclude  bool        `json:"left_include,omitempty"`
	RightInclude bool        `json:"right_include,omitempty"`
}

func (n *Bound) GetBoundType() BoundType {
	if n == nil {
		return UNKNOWN_BOUND_TYPE
	} else if n.LeftInclude && n.RightInclude {
		return LEFT_INCLUDE_RIGHT_INCLUDE
	} else if n.LeftInclude && !n.RightInclude {
		return LEFT_INCLUDE_RIGHT_EXCLUDE
	} else if !n.LeftInclude && n.RightInclude {
		return LEFT_EXCLUDE_RIGHT_INCLUDE
	} else {
		return LEFT_EXCLUDE_RIGHT_EXCLUDE
	}
}

type RangeValue struct {
	SideFlag    bool     // 表示左右的 左为false 右为true
	InfinityVal string   `parser:"  @('*')" json:"infinity_val"`
	PhraseValue []string `parser:"| QUOTE @( REVERSE QUOTE | !QUOTE )* QUOTE" json:"phrase_value"`
	SingleValue []string `parser:"| @(IDENT|ESCAPE|NUMBER|DOT|PLUS|MINUS|SOR|SLASH|COLON)+" json:"simple_value"`
}

func (v *RangeValue) String() string {
	if v == nil {
		return ""
	} else if len(v.PhraseValue) != 0 {
		return "\"" + strings.Join(v.PhraseValue, "") + "\""
	} else if len(v.InfinityVal) != 0 {
		return v.InfinityVal
	} else if len(v.SingleValue) != 0 {
		return strings.Join(v.SingleValue, "")
	} else {
		return ""
	}
}

func (v *RangeValue) Value(f func(string) (interface{}, error)) (interface{}, error) {
	if v == nil {
		return nil, ErrEmptyValue
	} else if len(v.InfinityVal) != 0 {
		return f(v.InfinityVal)
	} else if len(v.PhraseValue) != 0 {
		return f(strings.Join(v.PhraseValue, ""))
	} else if len(v.SingleValue) != 0 {
		return f(strings.Join(v.SingleValue, ""))
	} else {
		return f("")
	}
}

// sign == 0 only check v is infinite
// sign > 0 check v is positive infinite
// sign < 0 check v is negative infinite
func (v *RangeValue) IsInf(sign int) bool {
	if sign == 0 {
		return v != nil && len(v.InfinityVal) != 0
	} else if sign > 0 {
		return v.IsInf(0) && v.SideFlag
	} else {
		return v.IsInf(0) && !v.SideFlag
	}
}
