package prefix

import "github.com/zhuliquan/lucene_parser/term"

type Term struct {
	RegexpTerm *term.RegexpTerm `parser:"  @@" json:"regexp_term"`
	FuzzyTerm  *term.FuzzyTerm  `parser:"| @@" json:"fuzzy_term"`
	RangeTerm  *term.RangeTerm  `parser:"| @@" json:"range_term"`
	TermGroup  *TermGroup       `parser:"| @@" json:"term_group"`
}

func (t *Term) String() string {
	if t == nil {
		return ""
	} else if t.RegexpTerm != nil {
		return t.RegexpTerm.String()
	} else if t.FuzzyTerm != nil {
		return t.FuzzyTerm.String()
	} else if t.RangeTerm != nil {
		return t.RangeTerm.String()
	} else if t.TermGroup != nil {
		return t.TermGroup.String()
	} else {
		return ""
	}
}

func (t *Term) GetTermType() term.TermType {
	if t == nil {
		return term.UNKNOWN_TERM_TYPE
	} else if t.RegexpTerm != nil {
		return t.RegexpTerm.GetTermType()
	} else if t.FuzzyTerm != nil {
		return t.FuzzyTerm.GetTermType()
	} else if t.RangeTerm != nil {
		return t.RangeTerm.GetTermType()
	} else if t.TermGroup != nil {
		return t.TermGroup.GetTermType()
	} else {
		return term.UNKNOWN_TERM_TYPE
	}
}

func (t *Term) Value(f func(string) (interface{}, error)) (interface{}, error) {
	if t == nil {
		return nil, term.ErrEmptyTerm
	} else if t.FuzzyTerm != nil {
		return t.FuzzyTerm.Value(f)
	} else if t.RegexpTerm != nil {
		return t.RegexpTerm.Value(f)
	} else if t.RangeTerm != nil {
		return t.RangeTerm.GetBound(), nil
	} else if t.TermGroup != nil {
		return t.TermGroup.Value(f)
	} else {
		return f("")
	}
}

func (t *Term) GetBound() *term.Bound {
	if t == nil || t.RangeTerm == nil {
		return nil
	} else {
		return t.RangeTerm.GetBound()
	}
}

func (t *Term) Fuzziness() term.Fuzziness {
	if t == nil || t.FuzzyTerm == nil {
		return term.NoFuzzy
	} else {
		return t.FuzzyTerm.Fuzzy()
	}
}

func (t *Term) Boost() term.BoostValue {
	if t == nil {
		return term.NoBoost
	} else if t.FuzzyTerm != nil {
		return t.FuzzyTerm.Boost()
	} else if t.RangeTerm != nil {
		return t.RangeTerm.Boost()
	} else if t.TermGroup != nil {
		return t.TermGroup.Boost()
	} else if t.RegexpTerm != nil {
		return t.RegexpTerm.Boost()
	} else {
		return term.NoBoost
	}
}
