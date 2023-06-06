package lucene_parser

import (
	"fmt"
	"strings"

	"github.com/alecthomas/participle"
	op "github.com/zhuliquan/lucene_parser/operator"
	tm "github.com/zhuliquan/lucene_parser/term"
	tk "github.com/zhuliquan/lucene_parser/token"
)

var LuceneParser *participle.Parser

func init() {
	LuceneParser = participle.MustBuild(
		&Lucene{},
		participle.Lexer(tk.Lexer),
	)
}

func ParseLucene(queryString string) (*Lucene, error) {
	var (
		err error
		lqy = &Lucene{}
	)
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("failed to parse lucene, err: %+v", r)
		}
	}()

	if err = LuceneParser.ParseString(queryString, lqy); err != nil {
		return nil, err
	} else {
		return lqy, nil
	}
}

type Query interface {
	String() string
	GetQueryType() QueryType
}

// lucene: consist of or query and or symbol query
type Lucene struct {
	OrQuery *OrQuery   `parser:"@@" json:"or_query"`
	OSQuery []*OSQuery `parser:"@@*" json:"or_sym_query"`
}

func (q *Lucene) GetQueryType() QueryType {
	return LUCENE_QUERY
}

func (q *Lucene) String() string {
	if q == nil {
		return ""
	} else if q.OrQuery != nil {
		var sl = []string{q.OrQuery.String()}
		for _, x := range q.OSQuery {
			sl = append(sl, x.String())
		}
		return strings.Join(sl, "")
	} else {
		return ""
	}
}

// or query: consist of and query and and_symbol_query
type OrQuery struct {
	AndQuery *AndQuery   `parser:"@@" json:"and_query"`
	AnSQuery []*AnSQuery `parser:"@@*" json:"and_sym_query" `
}

func (q *OrQuery) GetQueryType() QueryType {
	return OR_QUERY
}

func (q *OrQuery) String() string {
	if q == nil || q.AndQuery == nil {
		return ""
	} else {
		var sl = []string{q.AndQuery.String()}
		for _, x := range q.AnSQuery {
			sl = append(sl, x.String())
		}
		return strings.Join(sl, "")
	}
}

// or symbol query: or query is prefix with or symbol
type OSQuery struct {
	OrSymbol *op.OrSymbol `parser:"@@" json:"or_symbol"`
	OrQuery  *OrQuery     `parser:"@@" json:"or_query"`
}

func (q *OSQuery) GetQueryType() QueryType {
	return OS_QUERY
}

func (q *OSQuery) String() string {
	if q == nil || q.OrQuery == nil {
		return ""
	} else {
		return q.OrSymbol.String() + q.OrQuery.String()
	}
}

// and query: consist of not query and paren query and field_query
type AndQuery struct {
	NotSymbol  *op.NotSymbol `parser:"  @@?" json:"not_symbol"`
	ParenQuery *ParenQuery   `parser:"( @@ " json:"paren_query"`
	FieldQuery *FieldQuery   `parser:"| @@)" json:"field_query"`
}

func (q *AndQuery) GetQueryType() QueryType {
	return AND_QUERY
}

func (q *AndQuery) String() string {
	if q == nil {
		return ""
	} else if q.ParenQuery != nil {
		return q.NotSymbol.String() + q.ParenQuery.String()
	} else if q.FieldQuery != nil {
		return q.NotSymbol.String() + q.FieldQuery.String()
	} else {
		return ""
	}
}

// and symbol query: and query is prefix with and symbol
type AnSQuery struct {
	AndSymbol *op.AndSymbol `parser:"( @@ " json:"and_symbol"`
	NotSymbol *op.NotSymbol `parser:"| WHITESPACE+ @@)" json:"not_symbol"`
	AndQuery  *AndQuery     `parser:"@@" json:"and_query"`
}

func (q *AnSQuery) GetQueryType() QueryType {
	return ANS_QUERY
}

func (q *AnSQuery) String() string {
	if q == nil || q.AndQuery == nil {
		return ""
	} else {
		if q.AndSymbol != nil {
			return q.AndSymbol.String() + q.AndQuery.String()
		} else {
			return " AND " + q.NotSymbol.String() + q.AndQuery.String()
		}
	}
}

// paren query: lucene query is surround with paren
type ParenQuery struct {
	SubQuery *Lucene `parser:"LPAREN WHITESPACE* @@ WHITESPACE* RPAREN" json:"sub_query"`
}

func (q *ParenQuery) GetQueryType() QueryType {
	return PAREN_QUERY
}

func (q *ParenQuery) String() string {
	if q == nil || q.SubQuery == nil {
		return ""
	} else {
		return "( " + q.SubQuery.String() + " )"
	}
}

// field query: consit of field and term
type FieldQuery struct {
	Field *tm.Field `parser:"@@ COLON" json:"field"`
	Term  *tm.Term  `parser:"@@" json:"term"`
}

func (q *FieldQuery) GetQueryType() QueryType {
	return FIELD_QUERY
}

func (q *FieldQuery) String() string {
	if q == nil || q.Field == nil || q.Term == nil {
		return ""
	} else {
		return q.Field.String() + ":" + q.Term.String()
	}
}
