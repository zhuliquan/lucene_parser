package term

import (
	"fmt"
	"strings"
)

// simple term: is a single term without escape char and whitespace
type SingleTerm struct {
	First    string   `parser:@(IDENT|NUMBER|MINUS|PLUS)`
	Value    []string `parser:"@(IDENT|NUMBER|DOT|WILDCARD|MIUNS|PLUS|MINUS|SOR|SLASH)*" json:"value"`
	wildcard int8
}

func (t *SingleTerm) GetTermType() TermType {
	if t == nil {
		return UNKNOWN_TERM_TYPE
	}
	var res = SINGLE_TERM_TYPE
	if t.haveWildcard() {
		res |= WILDCARD_TERM_TYPE
	}
	return res
}

func (t *SingleTerm) String() string {
	if t == nil {
		return ""
	} else {
		return strings.Join(t.Value, "")
	}
}

func (t *SingleTerm) haveWildcard() bool {
	if t == nil {
		return false
	} else if t.wildcard == -1 {
		return false
	} else if t.wildcard == 1 {
		return true
	} else {
		for _, tk := range t.Value {
			if tk == "*" || tk == "?" {
				t.wildcard = 1
				return true
			}
		}
		t.wildcard = -1
		return false
	}

}

// phrase term: a series of terms be surrounded with quotation, for instance "foo bar".
type PhraseTerm struct {
	Value    []string `parser:"QUOTE @( REVERSE QUOTE | !QUOTE )* QUOTE" json:"value"`
	wildcard int8
}

func (t *PhraseTerm) GetTermType() TermType {
	if t == nil {
		return UNKNOWN_TERM_TYPE
	}
	var res = PHRASE_TERM_TYPE
	if t.haveWildcard() {
		res |= WILDCARD_TERM_TYPE
	}
	return res
}

func (t *PhraseTerm) ValueS() string {
	if t == nil || len(t.Value) == 0 {
		return ""
	} else {
		return strings.Join(t.Value, "")
	}
}

func (t *PhraseTerm) String() string {
	if t == nil {
		return ""
	} else {
		return "\"" + strings.Join(t.Value, "") + "\""
	}
}

func (t *PhraseTerm) haveWildcard() bool {
	if t == nil {
		return false
	} else if t.wildcard == -1 {
		return false
	} else if t.wildcard == 1 {
		return true
	} else {
		for _, tk := range t.Value {
			if tk == "*" || tk == "?" {
				t.wildcard = 1
				return true
			}
		}
		t.wildcard = -1
		return false
	}

}

// a regexp term is surrounded be slash, for instance /\d+\.?\d+/ in here if you want present '/' you should type '\/'
type RegexpTerm struct {
	Value []string `parser:"SLASH @( REVERSE SLASH | !SLASH )+ SLASH" json:"value"`
}

func (t *RegexpTerm) GetTermType() TermType {
	if t == nil {
		return UNKNOWN_TERM_TYPE
	}
	return REGEXP_TERM_TYPE
}

func (t *RegexpTerm) ValuesS() string {
	if t == nil || len(t.Value) == 0 {
		return ""
	} else {
		return strings.Join(t.Value, "")
	}
}

func (t *RegexpTerm) String() string {
	if t == nil {
		return ""
	} else {
		return "/" + strings.Join(t.Value, "") + "/"
	}
}

//double side of range term: a term is surrounded by brace / bracket, for instance [1 TO 2] / [1 TO 2} / {1 TO 2] / {1 TO 2}
type DRangeTerm struct {
	LBRACKET string      `parser:"@(LBRACE|LBRACK) WHITESPACE*" json:"left_bracket"`
	LValue   *RangeValue `parser:"@@ WHITESPACE+ 'TO'" json:"left_value"`
	RValue   *RangeValue `parser:"WHITESPACE+ @@" json:"right_value"`
	RBRACKET string      `parser:"WHITESPACE* @(RBRACK|RBRACE)" json:"right_bracket"`
}

func (t *DRangeTerm) GetTermType() TermType {
	if t == nil {
		return UNKNOWN_TERM_TYPE
	}
	return RANGE_TERM_TYPE
}

func (t *DRangeTerm) GetBound() *Bound {
	var res *Bound
	if t == nil {
		return nil
	} else if t.LBRACKET == "[" && t.RBRACKET == "]" {
		res = &Bound{LeftValue: t.LValue, RightValue: t.RValue, LeftInclude: true, RightInclude: true}
	} else if t.LBRACKET == "[" && t.RBRACKET == "}" {
		res = &Bound{LeftValue: t.LValue, RightValue: t.RValue, LeftInclude: true, RightInclude: false}
	} else if t.LBRACKET == "{" && t.RBRACKET == "]" {
		res = &Bound{LeftValue: t.LValue, RightValue: t.RValue, LeftInclude: false, RightInclude: true}
	} else if t.LBRACKET == "{" && t.RBRACKET == "}" {
		res = &Bound{LeftValue: t.LValue, RightValue: t.RValue, LeftInclude: false, RightInclude: false}
	} else {
		return nil
	}
	res.LeftInclude = res.LeftInclude && !t.LValue.IsInf()
	res.RightInclude = res.RightInclude && !t.RValue.IsInf()
	return res
}

func (t *DRangeTerm) String() string {
	if t == nil {
		return ""
	} else {
		return fmt.Sprintf("%s %s TO %s %s", t.LBRACKET, t.LValue.String(), t.RValue.String(), t.RBRACKET)
	}
}

// single side of range term: a term is behind of symbol ('>' / '<' / '>=' / '<=')
type SRangeTerm struct {
	Symbol string      `parser:"@COMPARE" json:"symbol"`
	Value  *RangeValue `parser:"@@" json:"value"`
	drange *DRangeTerm
}

func (t *SRangeTerm) GetTermType() TermType {
	if t == nil {
		return UNKNOWN_TERM_TYPE
	}
	return RANGE_TERM_TYPE
}

func (t *SRangeTerm) toDRangeTerm() *DRangeTerm {
	if t == nil {
		return nil
	} else if t.drange != nil {
		return t.drange
	} else {
		if t.Symbol == ">" && t.Value != nil {
			t.drange = &DRangeTerm{LBRACKET: "{", LValue: t.Value, RValue: &RangeValue{InfinityVal: "*"}, RBRACKET: "}"}
		} else if t.Symbol == ">=" && t.Value != nil {
			t.drange = &DRangeTerm{LBRACKET: "[", LValue: t.Value, RValue: &RangeValue{InfinityVal: "*"}, RBRACKET: "}"}
		} else if t.Symbol == "<" && t.Value != nil {
			t.drange = &DRangeTerm{LBRACKET: "{", LValue: &RangeValue{InfinityVal: "*"}, RValue: t.Value, RBRACKET: "}"}
		} else if t.Symbol == "<=" && t.Value != nil {
			t.drange = &DRangeTerm{LBRACKET: "{", LValue: &RangeValue{InfinityVal: "*"}, RValue: t.Value, RBRACKET: "]"}
		}
	}
	return t.drange
}

func (t *SRangeTerm) GetBound() *Bound {
	return t.toDRangeTerm().GetBound()
}

func (t *SRangeTerm) String() string {
	return t.toDRangeTerm().String()
}
