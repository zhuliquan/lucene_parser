// +build prefix_group

package term

import (
	"strconv"
	"strings"

	op "github.com/zhuliquan/lucene_parser/operator"
)

// prefix term: a term is behind of prefix operator symbol ("+" / "-")
type PrefixTerm struct {
	Symbol string         `parser:"@( PLUS | MINUS)?" json:"symbol"`
	Elem   *TermGroupElem `parser:"@@" json:"elem"`
}

func (t *PrefixTerm) String() string {
	if t == nil {
		return ""
	} else if t.Elem != nil {
		return t.Symbol + t.Elem.String()
	} else {
		return ""
	}
}

func (t *PrefixTerm) GetTermType() TermType {
	if t == nil {
		return UNKNOWN_TERM_TYPE
	} else if t.Symbol != "" {
		return t.Elem.GetTermType() | PREFIX_TERM_TYPE
	} else {
		return t.Elem.GeTermType()
	}
}

func (t *PrefixTerm) GetPrefixType() op.PrefixOPType {
	if t == nil {
		return op.UNKNOWN_PREFIX_TYPE
	} else if t.Symbol == "+" {
		return op.MUST_PREFIX_TYPE
	} else if t.Symbol == "-" {
		return op.MUST_NOT_PREFIX_TYPE
	} else {
		return op.SHOULD_PREFIX_TYPE
	}
}

// whitespace is prefix with prefix term
type WPrefixTerm struct {
	Symbol string         `parser:"WHITESPACE @(PLUS|MINUS)?" json:"symbol"`
	Elem   *TermGroupElem `parser:"@@" json:"elem"`
}

func (t *WPrefixTerm) String() string {
	if t == nil {
		return ""
	} else if t.Elem != nil {
		return " " + t.Symbol + t.Elem.String()
	} else {
		return ""
	}
}

func (t *WPrefixTerm) GetTermType() TermType {
	if t == nil {
		return UNKNOWN_TERM_TYPE
	} else if t.Symbol != "" {
		return t.Elem.GetTermType() | PREFIX_TERM_TYPE
	} else {
		return t.Elem.GeTermType()
	}
}

func (t *WPrefixTerm) GetPrefixType() op.PrefixOPType {
	if t == nil {
		return op.UNKNOWN_PREFIX_TYPE
	} else if t.Symbol == "+" {
		return op.MUST_PREFIX_TYPE
	} else if t.Symbol == "-" {
		return op.MUST_NOT_PREFIX_TYPE
	} else {
		return op.SHOULD_PREFIX_TYPE
	}
}

type PrefixTermGroup struct {
	PrefixTerm  *PrefixTerm    `parser:"@@ " json:"prefix_term"`
	PrefixTerms []*WPrefixTerm `parser:"@@*" json:"prefix_terms"`
}

func (t *PrefixTermGroup) String() string {
	if t == nil {
		return ""
	} else if t.PrefixTerm != nil {
		var sl = []string{t.PrefixTerm.String()}
		for _, x := range t.PrefixTerms {
			sl = append(sl, x.String())
		}
		return strings.Join(sl, "")
	} else {
		return ""
	}
}

func (t *PrefixTermGroup) GetTermType() TermType {
	if t == nil {
		return UNKNOWN_TERM_TYPE
	} else {
		return GROUP_TERM_TYPE
	}
}

type TermGroup struct {
	PrefixTermGroup *PrefixTermGroup `parser:"LPAREN WHITESPACE* @@ WHITESPACE* RPAREN" json:"prefix_term_group"`
	BoostSymbol     string           `parser:"@(BOOST NUMBER (DOT NUMBER)?)?" json:"boost_symbol"`
}

func (t *TermGroup) String() string {
	return "( " + t.PrefixTermGroup.String() + " )" + t.BoostSymbol
}

func (t *TermGroup) Boost() float64 {
	if t == nil {
		return 0.0
	} else if len(t.BoostSymbol) == 0 {
		return 1.0
	} else {
		var res, _ = strconv.ParseFloat(t.BoostSymbol[1:], 64)
		return res
	}
}

func (t *TermGroup) GetTermType() TermType {
	if t == nil {
		return UNKNOWN_TERM_TYPE
	} else {
		var res = t.PrefixTermGroup.GetTermType()
		if len(t.BoostSymbol) != 0 {
			res |= BOOST_TERM_TYPE
		}
		return res
	}
}
