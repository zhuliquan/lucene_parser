package term

import "fmt"

var (
	ErrEmptyTerm          = fmt.Errorf("term is nil")
	ErrEmptyFuzzyTerm     = fmt.Errorf("single/phrase term is nil")
	ErrEmptySingleTerm    = fmt.Errorf("single term is nil")
	ErrEmptyPhraseTerm    = fmt.Errorf("phrase term is nil")
	ErrEmptyRegexpTerm    = fmt.Errorf("regexp term is nil")
	ErrEmptyGroupTerm     = fmt.Errorf("group term is nil")
	ErrEmptyTermGroupElem = fmt.Errorf("term group element is nil")
)
