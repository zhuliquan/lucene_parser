package term

import (
	"strings"

	op "github.com/zhuliquan/lucene_parser/operator"
)

// logic term group: join sum term elem by OR / AND / NOT
type LogicTermGroup struct {
	OrTermGroup *OrTermGroup   `parser:"@@ " json:"or_term_group"`
	OSTermGroup []*OSTermGroup `parser:"@@*" json:"or_symbol_term_group"`
}

func (t *LogicTermGroup) String() string {
	if t == nil || t.OrTermGroup == nil {
		return ""
	} else {
		var sl = []string{t.OrTermGroup.String()}
		for _, x := range t.OSTermGroup {
			sl = append(sl, x.String())
		}
		return strings.Join(sl, "")
	}
}

type OrTermGroup struct {
	AndTermGroup *AndTermGroup   `parser:"@@ " json:"and_term_group"`
	AnSTermGroup []*AnSTermGroup `parser:"@@*" json:"and_symbol_term_group"`
}

func (t *OrTermGroup) String() string {
	if t == nil || t.AndTermGroup == nil {
		return ""
	} else {
		var sl = []string{t.AndTermGroup.String()}
		for _, x := range t.AnSTermGroup {
			sl = append(sl, x.String())
		}
		return strings.Join(sl, "")
	}
}

// "or" | " !" | " not "
type OSTermGroup struct {
	OrSymbol    *op.OrSymbol  `parser:"( @@ " json:"or_symbol"`
	NotSymbol   *op.NotSymbol `parser:"| WHITESPACE+ @@)" json:"not_symbol"`
	OrTermGroup *OrTermGroup  `parser:"@@ " json:"or_term_group"`
}

func (t *OSTermGroup) String() string {
	if t == nil || t.OrTermGroup == nil {
		return ""
	} else {
		if t.OrSymbol != nil {
			return t.OrSymbol.String() + t.OrTermGroup.String()
		} else {
			return " OR " + t.NotSymbol.String() + t.OrTermGroup.String()
		}
	}
}

type AndTermGroup struct {
	NotSymbol      *op.NotSymbol   `parser:"@@?" json:"not_symbol"`
	ParenTermGroup *ParenTermGroup `parser:"( @@ " json:"paren_term_group"`
	TermGroupElem  *TermGroupElem  `parser:"| @@)" json:"term_group_elem"`
}

func (t *AndTermGroup) String() string {
	if t == nil {
		return ""
	} else if t.ParenTermGroup != nil {
		return t.NotSymbol.String() + t.ParenTermGroup.String()
	} else if t.TermGroupElem != nil {
		return t.NotSymbol.String() + t.TermGroupElem.String()
	} else {
		return ""
	}
}

type AnSTermGroup struct {
	AndSymbol    *op.AndSymbol `parser:"@@" json:"and_symbol"`
	AndTermGroup *AndTermGroup `parser:"@@" json:"and_term_group"`
}

func (t *AnSTermGroup) String() string {
	if t == nil || t.AndTermGroup == nil {
		return ""
	} else {
		return t.AndSymbol.String() + t.AndTermGroup.String()
	}
}

type ParenTermGroup struct {
	SubTermGroup *LogicTermGroup `parser:"LPAREN WHITESPACE* @@ WHITESPACE* RPAREN" json:"sub_term_group"`
}

func (t *ParenTermGroup) String() string {
	if t == nil || t.SubTermGroup == nil {
		return ""
	} else {
		return "( " + t.SubTermGroup.String() + " )"
	}
}

// term group: join sum prefix term group together
type TermGroup struct {
	LogicTermGroup *LogicTermGroup `parser:"LPAREN WHITESPACE* @@ WHITESPACE* RPAREN" json:"logic_term_group"`
	BoostSymbol    string          `parser:"@(BOOST NUMBER? (DOT NUMBER)?)?" json:"boost_symbol"`
}

func (t *TermGroup) String() string {
	if t == nil || t.LogicTermGroup == nil {
		return ""
	} else {
		return "( " + t.LogicTermGroup.String() + " )" + t.BoostSymbol
	}
}

func (t *TermGroup) Boost() float64 {
	if t == nil || t.LogicTermGroup == nil {
		return 0.0
	} else {
		return getBoostValue(t.BoostSymbol)
	}
}

func (t *TermGroup) GetTermType() TermType {
	if t == nil || t.LogicTermGroup == nil {
		return UNKNOWN_TERM_TYPE
	} else if len(t.BoostSymbol) == 0 {
		return GROUP_TERM_TYPE
	} else {
		return GROUP_TERM_TYPE | BOOST_TERM_TYPE
	}
}

func (t *TermGroup) Value(f func(string) (interface{}, error)) (interface{}, error) {
	if t == nil || t.LogicTermGroup == nil {
		return nil, ErrEmptyGroupTerm
	} else {
		return f(t.LogicTermGroup.String())
	}
}
