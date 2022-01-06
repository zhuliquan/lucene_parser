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

func (t *Term) GetBound() *Bound {
	if t == nil || t.RangeTerm == nil {
		return nil
	} else {
		return t.RangeTerm.GetBound()
	}
}

func (t *Term) Fuzziness() int {
	if t == nil {
		return 0
	} else if t.FuzzyTerm != nil {
		return t.FuzzyTerm.Fuzziness()
	} else {
		return 0
	}
}

func (t *Term) Boost() float64 {
	if t == nil {
		return 0.0
	} else if t.FuzzyTerm != nil {
		return t.FuzzyTerm.Boost()
	} else if t.RangeTerm != nil {
		return t.RangeTerm.Boost()
	} else if t.TermGroup != nil {
		return t.TermGroup.Boost()
	} else {
		return 1.0
	}
}
