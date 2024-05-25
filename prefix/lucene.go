package prefix

import (
	"fmt"
	"strings"

	"github.com/alecthomas/participle"
	"github.com/zhuliquan/lucene_parser"
	"github.com/zhuliquan/lucene_parser/term"
	"github.com/zhuliquan/lucene_parser/token"
)

var LuceneParser *participle.Parser

func init() {
	LuceneParser = participle.MustBuild(
		&Lucene{},
		participle.Lexer(token.Lexer),
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

// lucene: consist of list of prefix clauses
type Lucene struct {
	Clauses []*PrefixClause `parser:"@@*" json:"or_query"`
}

func (q *Lucene) GetQueryType() lucene_parser.QueryType {
	return lucene_parser.LUCENE_QUERY
}

func (q *Lucene) String() string {
	if q == nil {
		return ""
	} else {
		var sl = []string{}
		for _, x := range q.Clauses {
			sl = append(sl, x.String())
		}
		return strings.Join(sl, " ")
	}
}

// PrefixClause: prefix operator is prefix operator paren query and field_query
type PrefixClause struct {
	PrefixOp   string      `parser:"WHITESPACE* @( PLUS | MINUS | '!')?" json:"prefix_op"`
	ParenQuery *ParenQuery `parser:"( @@ " json:"paren_query"`
	FieldQuery *FieldQuery `parser:"| @@)" json:"field_query"`
}

func (q *PrefixClause) String() string {
	if q == nil {
		return ""
	}
	if q.FieldQuery != nil {
		return q.PrefixOp + q.FieldQuery.String()
	} else if q.ParenQuery != nil {
		return q.PrefixOp + q.ParenQuery.String()
	} else {
		return ""
	}
}

func (q *PrefixClause) GetQueryType() lucene_parser.QueryType {
	switch q.PrefixOp {
	case "+":
		return lucene_parser.AND_QUERY
	case "-":
		return lucene_parser.NOT_QUERY
	default:
		return lucene_parser.OR_QUERY
	}
}

// ParenQuery: lucene query is surround with paren
type ParenQuery struct {
	SubQuery *Lucene `parser:"LPAREN WHITESPACE* @@ WHITESPACE* RPAREN" json:"sub_query"`
}

func (q *ParenQuery) GetQueryType() lucene_parser.QueryType {
	return lucene_parser.PAREN_QUERY
}

func (q *ParenQuery) String() string {
	if q == nil || q.SubQuery == nil {
		return ""
	} else {
		return "( " + q.SubQuery.String() + " )"
	}
}

// FieldQuery: consist of field and term
type FieldQuery struct {
	Field *term.Field `parser:"@@ COLON" json:"field"`
	Term  *Term       `parser:"@@" json:"term"`
}

func (q *FieldQuery) GetQueryType() lucene_parser.QueryType {
	return lucene_parser.FIELD_QUERY
}

func (q *FieldQuery) String() string {
	if q == nil || q.Field == nil || q.Term == nil {
		return ""
	} else {
		return q.Field.String() + ":" + q.Term.String()
	}
}
