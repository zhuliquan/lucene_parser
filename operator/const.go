package operator

type PrefixOPType uint32

const (
	UNKNOWN_PREFIX_TYPE PrefixOPType = iota
	SHOULD_PREFIX_TYPE
	MUST_PREFIX_TYPE
	MUST_NOT_PREFIX_TYPE
)

var prefixOPType_Values = map[PrefixOPType]string{
	UNKNOWN_PREFIX_TYPE:  "",
	SHOULD_PREFIX_TYPE:   " ",
	MUST_PREFIX_TYPE:     " +",
	MUST_NOT_PREFIX_TYPE: " -",
}

func (o PrefixOPType) String() string {
	return prefixOPType_Values[o]
}

type LogicOPType uint32

const (
	UNKNOWN_LOGIC_TYPE LogicOPType = iota
	AND_LOGIC_TYPE
	OR_LOGIC_TYPE
	NOT_LOGIC_TYPE
)

var LogicOPType_Values = map[LogicOPType]string{
	UNKNOWN_LOGIC_TYPE: "",
	AND_LOGIC_TYPE:     " AND ",
	OR_LOGIC_TYPE:      " OR ",
	NOT_LOGIC_TYPE:     "NOT ",
}

func (o LogicOPType) String() string {
	return LogicOPType_Values[o]
}
