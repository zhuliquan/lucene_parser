package standard

import (
	"fmt"

	"github.com/alecthomas/participle"
)

var LuceneParser *participle.Parser

func init() {
	LuceneParser = participle.MustBuild(
		&Lucene{},
		participle.Lexer(Lexer),
		participle.CaseInsensitive("IDENT"),
		participle.UseLookahead(1024),
	)
}

func ShowGrammar() {
	fmt.Println(LuceneParser)
}

// ParseLucene: parse query to Lucene struct
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
