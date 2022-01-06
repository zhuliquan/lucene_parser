package term

import (
	"strconv"
)

// single side range term or double side range and with boost like this [1 TO 2]^2
type RangeTerm struct {
	SRangeTerm  *SRangeTerm `parser:"( @@ " json:"s_range_term"`
	DRangeTerm  *DRangeTerm `parser:"| @@)" json:"d_range_term"`
	BoostSymbol string      `parser:"@(BOOST NUMBER (DOT NUMBER)?)?" json:"boost_symbol"`
}

func (t *RangeTerm) GetTermType() TermType {
	if t == nil {
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

func (t *RangeTerm) Boost() float64 {
	if t == nil {
		return 0.0
	} else if len(t.BoostSymbol) == 0 {
		return 1.0
	} else {
		var res, _ = strconv.ParseFloat(t.BoostSymbol[1:], 64)
		return res
	}
}

// fuzzy term: term can by suffix with fuzzy or boost like this foo^2 / "foo bar"^2 / foo~ / "foo bar"~2
type FuzzyTerm struct {
	SingleTerm  *SingleTerm `parser:"( @@ " json:"single_term"`
	PhraseTerm  *PhraseTerm `parser:"| @@)" json:"phrase_term"`
	FuzzySymbol string      `parser:"( @(FUZZY NUMBER?)  " json:"fuzzy_symbol"`
	BoostSymbol string      `parser:"| @(BOOST NUMBER (DOT NUMBER)?))?" json:"boost_symbol"`
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

func (t *FuzzyTerm) Boost() float64 {
	if t == nil {
		return 0.0
	} else if len(t.BoostSymbol) == 0 {
		return 1.0
	} else {
		var res, _ = strconv.ParseFloat(t.BoostSymbol[1:], 64)
		return res
	}
}

func (t *FuzzyTerm) Fuzziness() int {
	if t == nil || len(t.FuzzySymbol) == 0 {
		return 0
	} else if t.FuzzySymbol == "~" {
		return 1
	} else {
		var v, _ = strconv.Atoi(t.FuzzySymbol[1:])
		return v
	}
}

func (t *FuzzyTerm) String() string {
	if t == nil {
		return ""
	} else if t.SingleTerm != nil {
		return t.SingleTerm.String() + t.FuzzySymbol
	} else if t.PhraseTerm != nil {
		return t.PhraseTerm.String() + t.FuzzySymbol
	} else {
		return ""
	}
}

func (t *FuzzyTerm) haveWildcard() bool {
	if t == nil {
		return false
	} else if t.SingleTerm != nil {
		return t.SingleTerm.haveWildcard()
	} else if t.PhraseTerm != nil {
		return t.PhraseTerm.haveWildcard()
	} else {
		return false
	}
}

func (t *FuzzyTerm) ValueS() string {
	if t == nil {
		return ""
	} else if t.SingleTerm != nil {
		return t.SingleTerm.ValueS()
	} else if t.PhraseTerm != nil {
		return t.PhraseTerm.ValueS()
	} else {
		return ""
	}
}

// term group element
type TermGroupElem struct {
	SingleTerm *SingleTerm `parser:"  @@" json:"single_term"`
	PhraseTerm *PhraseTerm `parser:"| @@" json:"phrase_term"`
	SRangeTerm *SRangeTerm `parser:"| @@" json:"single_range_term"`
	DRangeTerm *DRangeTerm `parser:"| @@" json:"double_range_term"`
}

func (t *TermGroupElem) String() string {
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

func (t *TermGroupElem) GetTermType() TermType {
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
