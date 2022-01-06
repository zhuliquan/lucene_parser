package term

import "fmt"

type TermType uint64

const (
	UNKNOWN_TERM_TYPE TermType = 1 << iota
	SINGLE_TERM_TYPE
	PHRASE_TERM_TYPE
	REGEXP_TERM_TYPE
	RANGE_TERM_TYPE
	PREFIX_TERM_TYPE
	WILDCARD_TERM_TYPE
	GROUP_TERM_TYPE
	FUZZY_TERM_TYPE
	BOOST_TERM_TYPE
)

type FieldType uint32

const (
	TEXT_FIELD_TYPE FieldType = iota
	KEYWORD_FIELD_TYPE
	INTEGER_FIELD_TYPE
	BOOLEAN_FIELD_TYPE
	DATE_FIELD_TYPE
	NESTED_FIELD_TYPE
	IP_FIELD_TYPE
	GEO_FIELD_TYPE
)

var (
	Inf           = &RangeValue{InfinityVal: "*"}
	ErrEmptyValue = fmt.Errorf("empty value")
)

type BoundType uint16

const (
	UNKNOWN_BOUND_TYPE         BoundType = iota
	LEFT_EXCLUDE_RIGHT_INCLUDE BoundType = iota
	LEFT_EXCLUDE_RIGHT_EXCLUDE
	LEFT_INCLUDE_RIGHT_INCLUDE
	LEFT_INCLUDE_RIGHT_EXCLUDE
)
