package operator

type PrefixOPType uint32

const (
	UNKNOWN_PREFIX_TYPE PrefixOPType = iota
	SHOULD_PREFIX_TYPE
	MUST_PREFIX_TYPE
	MUST_NOT_PREFIX_TYPE
)

type LogicOPType uint32

const (
	UNKNOWN_LOGIC_TYPE LogicOPType = iota
	AND_LOGIC_TYPE
	OR_LOGIC_TYPE
	NOT_LOGIC_TYPE
)
