package standard

// Lucene ::= Query <EOF>
type Lucene struct {
	Query *Query `parser:"@@"`
}

// Query ::= DisjQuery ( DisjQuery )*
type Query struct {
	DisjQueries []*DisjQuery `parser:"@@ ( WHITESPACE @@ )*"`
}

// DisjQuery ::= ConjQuery ( OR ConjQuery )*
type DisjQuery struct {
	ConjQueries []*ConjQuery `parser:"@@ ( WHITESPACE ('OR' | 'or' | '|' '|' ) WHITESPACE @@ )*"`
}

// ConjQuery ::= ModClause ( AND ModClause )*
type ConjQuery struct {
	ModClauses []*ModClause `parser:"@@ ( WHITESPACE ('AND' | 'and' | '&' '&' ) WHITESPACE @@ )*"`
}

// ModClause ::= (Modifier)? Clause
type ModClause struct {
	Modifier string  `parser:"@(PLUS | MINUS | NOT)?"`
	Clause   *Clause `parser:"@@"`
}

// Clause ::= (FieldName ':')? (TermExpr | GroupingExpr | Range | PhraseExpr | Regex) Boost?
type Clause struct {
	Field      *FieldName  `parser:"@@?"`
	TermExpr   *TermExpr   `parser:"( @@ "`
	PhraseExpr *PhraseExpr `parser:"| @@ "`
	GroupExpr  *GroupExpr  `parser:"| @@ "`
	RegexpExpr *Regexp     `parser:"| @@ "`
	RangeExpr  *Range      `parser:"| @@)"`
	Boost      *Boost      `parser:"@@?"`
}

// TermExpr ::= TERM Fuzzy?
type TermExpr struct {
	Term  *TERM  `parser:"@@"`
	Fuzzy *Fuzzy `parser:"@@?"`
}

// Range ::= SingleRange | DoubleRange
type Range struct {
	SingleRange *SingleRange `parser:"( @@ "`
	DoubleRange *DoubleRange `parser:"| @@)"`
}

// PhraseExpr ::= Phrase Boost?
type PhraseExpr struct {
	Phrase *Phrase `parser:"@@"`
	Fuzzy  *Fuzzy  `parser:"@@?"`
}

// GroupExpr ::= '(' Query ')' Boost?
type GroupExpr struct {
	Query *Query `parser:"LPAREN WHITESPACE? @@ WHITESPACE? RPAREN"`
}

type TERM struct {
	Token []string `parser:"@(IDENT | NUMBER | ESCAPE | DOT | MINUS | PLUS | WILDCARD)+"`
}

type Phrase struct {
	Token []string `parser:"QUOTE @( REVERSE QUOTE | !QUOTE )+ QUOTE"`
}

type Regexp struct {
	Token []string `parser:"SLASH @( REVERSE SLASH | !SLASH )+ SLASH"`
}

type SingleRange struct {
	Compare    string      `parser:"@COMPARE"`
	RangeValue *RangeValue `parser:"@@"`
}

type DoubleRange struct {
	LParen string     `parser:"@( LBRACE | LBRACK ) WHITESPACE?"`
	Left   *RangeNode `parser:"@@"`
	TO     string     `parser:"WHITESPACE @'TO' WHITESPACE"`
	Right  *RangeNode `parser:"@@"`
	RParen string     `parser:"WHITESPACE? @( RBRACE | RBRACK )"`
}

type RangeValue struct {
	Term   *TERM   `parser:"  @@"`
	Phrase *Phrase `parser:"| @@"`
	Number *Number `parser:"| @@"`
}

type RangeNode struct {
	RangeValue *RangeValue `parser:"  @@"`
	Infinite   *string     `parser:"| @'*'"`
}

type FieldName struct {
	FieldName *TERM `parser:"@@ COLON"`
}

//  Boost ::= ('^' NUMBER?)
type Boost struct {
	Number *Number `parser:"CARAT @@?"`
}

// Fuzzy :: = ( '~' NUMBER?)
type Fuzzy struct {
	Number *Number `parser:"TILDE @@?"`
}

type Number struct {
	Integer int `parser:"@NUMBER"`         // the part of integer
	Decimal int `parser:"@( DOT NUMBER)?"` // the part of decimal
}

type AND struct {
	AND string `parser:"@('AND' | 'and' | '&' '&')"`
}

type OR struct {
	OR string `parser:"@('OR' | 'or' | '|' '|')"`
}

type NOT struct {
	NOT string `parser:"@('NOT' | 'not' | NOT)"`
}
