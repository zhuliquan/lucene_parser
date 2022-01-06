package lucene_parser

type QueryType uint32

const (
	LUCENE_QUERY QueryType = iota
	OR_QUERY
	OS_QUERY
	AND_QUERY
	ANS_QUERY
	FIELD_QUERY
	PAREN_QUERY
)
