package prefix

import (
	"fmt"
	"strings"

	op "github.com/zhuliquan/lucene_parser/operator"
	"github.com/zhuliquan/lucene_parser/term"
)

// prefix operator term: a term is behind of prefix operator symbol ("+" / "-" / '!')
type PrefixOperatorTerm struct {
	PrefixOp       string               `parser:"WHITESPACE* @( PLUS | MINUS | '!')?" json:"prefix_op"`
	FieldTermGroup *term.FieldTermGroup `parser:"( @@" json:"field_term_group"`
	ParenTermGroup *PrefixTermGroup     `parser:"| LPAREN WHITESPACE* @@ WHITESPACE* RPAREN)" json:"paren_term_group"`
}

func (t *PrefixOperatorTerm) String() string {
	if t == nil {
		return ""
	} else if t.FieldTermGroup != nil {
		return fmt.Sprintf("%s%s", t.PrefixOp, t.FieldTermGroup)
	} else if t.ParenTermGroup != nil {
		return fmt.Sprintf("%s( %s )", t.PrefixOp, t.ParenTermGroup)
	}
	return ""
}

func (t *PrefixOperatorTerm) GetPrefixType() op.PrefixOPType {
	if t == nil {
		return op.UNKNOWN_PREFIX_TYPE
	} else if t.PrefixOp == "+" {
		return op.MUST_PREFIX_TYPE
	} else if t.PrefixOp == "-" {
		return op.MUST_NOT_PREFIX_TYPE
	} else {
		return op.SHOULD_PREFIX_TYPE
	}
}

type PrefixTermGroup struct {
	PrefixTerms []*PrefixOperatorTerm `parser:"@@*" json:"prefix_terms"`
}

func (t *PrefixTermGroup) String() string {
	if t == nil {
		return ""
	} else {
		var sl = []string{}
		for _, x := range t.PrefixTerms {
			sl = append(sl, x.String())
		}
		return strings.Join(sl, " ")
	}
}

func (t *PrefixTermGroup) GetTermType() term.TermType {
	if t == nil || len(t.PrefixTerms) == 0 {
		return term.UNKNOWN_TERM_TYPE
	} else {
		return term.GROUP_TERM_TYPE
	}
}

type TermGroup struct {
	PrefixTermGroup *PrefixTermGroup `parser:"LPAREN WHITESPACE* @@ WHITESPACE* RPAREN" json:"prefix_term_group"`
	BoostSymbol     string           `parser:"@(BOOST NUMBER? (DOT NUMBER)?)?" json:"boost_symbol"`
}

func (t *TermGroup) String() string {
	if t == nil || t.PrefixTermGroup == nil {
		return ""
	} else {
		return "( " + t.PrefixTermGroup.String() + " )" + t.BoostSymbol
	}
}

func (t *TermGroup) Boost() term.BoostValue {
	if t == nil || t.PrefixTermGroup == nil {
		return term.DefaultBoost
	} else {
		return term.GetBoostValue(t.BoostSymbol)
	}
}

func (t *TermGroup) GetTermType() term.TermType {
	if t == nil || t.PrefixTermGroup == nil {
		return term.UNKNOWN_TERM_TYPE
	} else {
		var res = t.PrefixTermGroup.GetTermType()
		if len(t.BoostSymbol) != 0 {
			res |= term.BOOST_TERM_TYPE
		}
		return res
	}
}

func (t *TermGroup) Value(f func(string) (interface{}, error)) (interface{}, error) {
	if t == nil || t.PrefixTermGroup == nil {
		return nil, term.ErrEmptyGroupTerm
	} else {
		return f(t.PrefixTermGroup.String())
	}
}
