package term

type Term struct {
	RegexpTerm *RegexpTerm `parser:"  @@" json:"regexp_term"`
	FuzzyTerm  *FuzzyTerm  `parser:"| @@" json:"fuzzy_term"`
	RangeTerm  *RangeTerm  `parser:"| @@" json:"range_term"`
	TermGroup  *TermGroup  `parser:"| @@" json:"term_group"`
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

func (t *Term) GetTermType() TermType {
	if t == nil {
		return UNKNOWN_TERM_TYPE
	} else if t.RegexpTerm != nil {
		return t.RegexpTerm.GetTermType()
	} else if t.FuzzyTerm != nil {
		return t.FuzzyTerm.GetTermType()
	} else if t.RangeTerm != nil {
		return t.RangeTerm.GetTermType()
	} else if t.TermGroup != nil {
		return t.TermGroup.GetTermType()
	} else {
		return UNKNOWN_TERM_TYPE
	}
}

func (t *Term) Value(f func(string) (interface{}, error)) (interface{}, error) {
	if t == nil {
		return nil, ErrEmptyTerm
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

func (t *Term) GetBound() *Bound {
	if t == nil || t.RangeTerm == nil {
		return nil
	} else {
		return t.RangeTerm.GetBound()
	}
}

func (t *Term) Fuzziness() Fuzziness {
	if t == nil || t.FuzzyTerm == nil {
		return NoFuzzy
	} else {
		return t.FuzzyTerm.Fuzzy()
	}
}

func (t *Term) Boost() BoostValue {
	if t == nil {
		return NoBoost
	} else if t.FuzzyTerm != nil {
		return t.FuzzyTerm.Boost()
	} else if t.RangeTerm != nil {
		return t.RangeTerm.Boost()
	} else if t.TermGroup != nil {
		return t.TermGroup.Boost()
	} else if t.RegexpTerm != nil {
		return t.RegexpTerm.Boost()
	} else {
		return NoBoost
	}
}
