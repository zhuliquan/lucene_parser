package term

// single side range term or double side range and with boost like this [1 TO 2]^2
type RangeTerm struct {
	SRangeTerm  *SRangeTerm `parser:"( @@ " json:"s_range_term"`
	DRangeTerm  *DRangeTerm `parser:"| @@)" json:"d_range_term"`
	BoostSymbol string      `parser:"@(BOOST NUMBER? (DOT NUMBER)?)?" json:"boost_symbol"`
}

func (t *RangeTerm) GetTermType() TermType {
	if t == nil || (t.DRangeTerm == nil && t.SRangeTerm == nil) {
		return UNKNOWN_TERM_TYPE
	}
	var res = RANGE_TERM_TYPE
	if len(t.BoostSymbol) != 0 {
		res |= BOOST_TERM_TYPE
	}
	return res
}

func (t *RangeTerm) String() string {
	if t == nil {
		return ""
	} else if t.SRangeTerm != nil {
		return t.SRangeTerm.String() + t.BoostSymbol
	} else if t.DRangeTerm != nil {
		return t.DRangeTerm.String() + t.BoostSymbol
	} else {
		return ""
	}
}

func (t *RangeTerm) GetBound() *Bound {
	if t == nil {
		return nil
	} else if t.SRangeTerm != nil {
		return t.SRangeTerm.GetBound()
	} else if t.DRangeTerm != nil {
		return t.DRangeTerm.GetBound()
	} else {
		return nil
	}
}

func (t *RangeTerm) Boost() BoostValue {
	if t == nil || (t.DRangeTerm == nil && t.SRangeTerm == nil) {
		return NoBoost
	} else {
		return GetBoostValue(t.BoostSymbol)
	}
}

// fuzzy term: term can by suffix with fuzzy or boost like this foo^2 / "foo bar"^2 / foo~ / "foo bar"~2
type FuzzyTerm struct {
	SingleTerm  *SingleTerm `parser:"( @@ " json:"single_term"`
	PhraseTerm  *PhraseTerm `parser:"| @@)" json:"phrase_term"`
	FuzzySymbol string      `parser:"( @(FUZZY NUMBER? (DOT NUMBER)?)  " json:"fuzzy_symbol"`
	BoostSymbol string      `parser:"| @(BOOST NUMBER? (DOT NUMBER)?))?" json:"boost_symbol"`
}

func (t *FuzzyTerm) GetTermType() TermType {
	var res TermType
	if t == nil {
		return UNKNOWN_TERM_TYPE
	} else if t.SingleTerm != nil {
		res = t.SingleTerm.GetTermType()
	} else if t.PhraseTerm != nil {
		res = t.PhraseTerm.GetTermType()
	} else {
		return UNKNOWN_TERM_TYPE
	}
	if len(t.BoostSymbol) != 0 {
		res |= BOOST_TERM_TYPE
	}
	if len(t.FuzzySymbol) != 0 {
		res |= FUZZY_TERM_TYPE
	}
	return res
}

func (t *FuzzyTerm) Boost() BoostValue {
	if t == nil || (t.SingleTerm == nil && t.PhraseTerm == nil) {
		return NoBoost
	} else {
		return GetBoostValue(t.BoostSymbol)
	}
}

func (t *FuzzyTerm) Fuzzy() Fuzziness {
	if t == nil || len(t.FuzzySymbol) == 0 || (t.SingleTerm == nil && t.PhraseTerm == nil) {
		return NoFuzzy
	} else {
		return getFuzzyValue(t.FuzzySymbol)
	}
}

func (t *FuzzyTerm) String() string {
	if t == nil {
		return ""
	} else if t.SingleTerm != nil {
		return t.SingleTerm.String() + t.FuzzySymbol + t.BoostSymbol
	} else if t.PhraseTerm != nil {
		return t.PhraseTerm.String() + t.FuzzySymbol + t.BoostSymbol
	} else {
		return ""
	}
}

func (t *FuzzyTerm) haveWildcard() bool {
	if t == nil {
		return false
	} else if t.SingleTerm != nil {
		return t.SingleTerm.haveWildcard()
	} else {
		return false
	}
}

func (t *FuzzyTerm) Value(f func(string) (interface{}, error)) (interface{}, error) {
	if t == nil {
		return nil, ErrEmptyFuzzyTerm
	} else if t.SingleTerm != nil {
		return t.SingleTerm.Value(f)
	} else if t.PhraseTerm != nil {
		return t.PhraseTerm.Value(f)
	} else {
		return f("")
	}
}

// term group element
type FieldTermGroup struct {
	SingleTerm *SingleTerm `parser:"  @@" json:"single_term"`
	PhraseTerm *PhraseTerm `parser:"| @@" json:"phrase_term"`
	SRangeTerm *SRangeTerm `parser:"| @@" json:"single_range_term"`
	DRangeTerm *DRangeTerm `parser:"| @@" json:"double_range_term"`
}

func (t *FieldTermGroup) String() string {
	if t == nil {
		return ""
	} else if t.SingleTerm != nil {
		return t.SingleTerm.String()
	} else if t.PhraseTerm != nil {
		return t.PhraseTerm.String()
	} else if t.SRangeTerm != nil {
		return t.SRangeTerm.String()
	} else if t.DRangeTerm != nil {
		return t.DRangeTerm.String()
	} else {
		return ""
	}
}

func (t *FieldTermGroup) GetTermType() TermType {
	if t == nil {
		return UNKNOWN_TERM_TYPE
	} else if t.SingleTerm != nil {
		return SINGLE_TERM_TYPE
	} else if t.PhraseTerm != nil {
		return PHRASE_TERM_TYPE
	} else if t.SRangeTerm != nil {
		return RANGE_TERM_TYPE
	} else if t.DRangeTerm != nil {
		return RANGE_TERM_TYPE
	} else {
		return UNKNOWN_TERM_TYPE
	}
}

func (t *FieldTermGroup) Value(f func(string) (interface{}, error)) (interface{}, error) {
	if t == nil {
		return nil, ErrEmptyTermGroupElem
	} else if t.SingleTerm != nil {
		return t.SingleTerm.Value(f)
	} else if t.PhraseTerm != nil {
		return t.PhraseTerm.Value(f)
	} else if t.SRangeTerm != nil {
		return t.SRangeTerm.GetBound(), nil
	} else if t.DRangeTerm != nil {
		return t.DRangeTerm.GetBound(), nil
	} else {
		return f("")
	}
}
