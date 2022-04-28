package operator

// and operator ("AND" / "and" / "&&")
type AndSymbol struct {
	Symbol string `parser:"  WHITESPACE+ @( AND AND | 'AND' | 'and' ) WHITESPACE+" json:"symbol"`
}

func (o *AndSymbol) String() string {
	if o == nil || o.Symbol == "" {
		return ""
	} else {
		return " AND "
	}
}

func (o *AndSymbol) GetLogicType() LogicOPType {
	if o == nil {
		return UNKNOWN_LOGIC_TYPE
	} else {
		return AND_LOGIC_TYPE
	}
}

// or operator ("OR" / "or" / "||")
type OrSymbol struct {
	Symbol string `parser:"  WHITESPACE+ @( SOR SOR | 'OR' | 'or' ) WHITESPACE+" json:"symbol"`
}

func (o *OrSymbol) String() string {
	if o == nil || o.Symbol == "" {
		return ""
	} else {
		return " OR "
	}
}

func (o *OrSymbol) GetLogicType() LogicOPType {
	if o == nil {
		return UNKNOWN_LOGIC_TYPE
	} else {
		return OR_LOGIC_TYPE
	}
}

// not operator ("NOT" / "not" / "!")
type NotSymbol struct {
	Symbol string `parser:"@( NOT | 'NOT' | 'not') WHITESPACE+" json:"symbol"`
}

func (o *NotSymbol) String() string {
	if o == nil || o.Symbol == "" {
		return ""
	} else {
		return "NOT "
	}
}

func (o *NotSymbol) GetLogicType() LogicOPType {
	if o == nil {
		return UNKNOWN_LOGIC_TYPE
	} else {
		return NOT_LOGIC_TYPE
	}
}
