# lucene_parser
## Introduction:
This package can parse lucene query used by ES (ElasticSearch), this package is pure go package, lex and yacc in this package doesn't follow standard lucene parser, and this package is used to convert lucene to other special DSL(domain special language) mainly, such as LuceneToSQL / LuceneToEQL (Used in ES). If you want to parse standard lucene query, you can use sub package **standard** in this repository
## Features
- 1、support phrase term query, for instance `x:"foo bar"`.
- 2、support regexp term query, for instance `x:/\d+\\.?\d+/`.
- 3、support bool operator （i.e. `AND`, `OR`, `NOT`, `&&`, `||`, `!`） join sub query, for instance `x:1 AND y:2`, `x:1 || y:2`, we also support lower case bool operator (i.e. `and`, `or`, `not`).
- 4、support bound range query,  for instance `x:[1 TO 2]`, `x:[1 TO 2}`.
- 5、support side range query, for instance `x:>1` , `x:>=1` , `x:<1` , `x:<=1`.
- 6、support boost modifier, for instance `x:1^2` , `x:"dsada 8908"^3`.
- 7、support [fuzzy query](https://www.elastic.co/guide/en/elasticsearch/guide/current/fuzzy-query.html) with default [fuzziness](https://www.elastic.co/guide/en/elasticsearch/guide/current/fuzziness.html) or specific fuzziness, for instance `x:for~1.0`, `x;foo~`.
- 8、support proximity query, for instance `x:"foo bar"~2`.
- 9、support term group query, for instance `x:(foo OR bar)`, `x:(>1 && <2)`.
- 10、support not operator be used with just one term (i.g. `not x:y`), this feature is differs from [the definition of `NOT` in standard lucene syntax](https://lucene.apache.org/core/2_9_4/queryparsersyntax.html#NOT).
- 11、support ignore `AND` operator when it behind with `NOT` operator (i.e. you can write `x:y and not x2:y2` as `x:y not x2:y2`).
- 12、support prefix operator `("+", "-", "!")` is ahead of field term, for instance `-foo:bar +foo1:bar1 foo2:bar2 !foo3:bar3`.

## Limitations
- 1、only support lucene query with **field name**, instead of query without **field name** (i.e. this project can't parse query like `foo OR bar`, `foo AND bar`, but can parse `foo:bar`, `foo:(bar1 AND bar2)`).
- 2、prefix and bool operator cannot be supported at the same time. on the other hand, you can't parse query which consist bool operator (`AND`/`OR`/`OR`/`NOT`/`&&`/`||`/`!`) and prefix operator (`+`/`-`) at same time.
- 3、don't support fuzziness of similarity (float number between 0 and 1), instead of fuzziness of maximum edit distance (i.e. Levenshtein Edit Distance — the number of one character changes that need to be made to one string to make it the same as another string.).
- 4、don't support space is regard as `OR` operator (i.g. `x1:y1 x2:y2`). (I don't know how to handle expression which includes both `or` token and space token (i.g. `x y or z`) . If you have good idea, please contact me)

## Note
- 1、If similarity is not specified in the fuzzy query, and you will get `-1` by invoking function `Fuzziness` of term, which is allow the user to customize the default fuzziness or parameter of [AUTO fuzziness](https://www.elastic.co/guide/en/elasticsearch/reference/8.4/common-options.html#fuzziness). For example, when -1 is returned, you can specify the maximum and minimum term length of the AUTO parameter according to the fuzziness definition.

- 2、according to definition of fuzziness, specific fuzziness must to be integer. if you input float fuzziness, we will round this number. For example: input query `x:foo~1.2`, you will get fuzziness `1`; input query `x:foo~1.6` you will get fuzziness `2`.

- 3、if you input boost symbol but value, you will get 1.0 by invoking function `Boost` of term. for instance query `foo:bar^`.

## Usage
### basic lucene parser
```golang
package main

import (
    "fmt"
    "github.com/zhuliquan/lucene_parser"
)

func main() {
    if lucene, err := lucene_parser.ParseLucene("x:foo AND y:bar"); err != nil {
        panic(err)
    } else {
        fmt.Println(lucene)
    }
}
```
### prefix operator lucene parser
You also can parse lucene query with prefix operator by using `prefix` package, as below:
```golang
package main

import (
    "fmt"
    "github.com/zhuliquan/lucene_parser/prefix"
)

func main() {
    if lucene, err := prefix.ParseLucene("x:foo AND y:bar"); err != nil {
        panic(err)
    } else {
        fmt.Println(lucene)
    }
}
```

### standard syntax lucene parser
You also can parse lucene which follows standard syntax by using **standard** package, as below:
```golang
package main

import (
    "fmt"
    "github.com/zhuliquan/lucene_parser/standard"
)

func main() {
    if lucene, err := standard.ParseLucene("foo^10 bar AND yacc"); err != nil {
        panic(err)
    } else {
        fmt.Println(lucene)
    }
}
```

## EBNF of Lucene
lucene parser will convert string of lucene query to ast, according to EBNF of lucene. EBNF of lucene is below.

```ebnf
(* lucene expression *)
lucene = or_query, { or_sym_query } ;
or_sym_query  = or_symbol, or_query ;
or_query      = and_query, { and_sym_query } ;
and_sym_query =  ( and_symbol | whitespace, not_symbol ), and_query ;
and_query     = [ not_symbol ], ( '(', [ whitespace ], lucene, [ whitespace ], ')' | ( field, term) ) ;

(* field and term *)
field_char       = identifier | '-' | number | dot ;
field            = field_char, { field_char }, ':' ;
term = range_term | fuzzy_term | regexp_term | term_group ;

(* term group *)
term_group = '(', logic_term_group, ')', [ boost_modifier ] ;
logic_term_group   = or_term_group, { or_sym_term_group } ;
or_sym_term_group  = ( or_symbol | whitespace, not_symbol ), or_term_group ;
or_term_group      = and_term_group, { and_sym_term_group } ;
and_sym_term_group = and_symbol, and_term_group ;
and_term_group     = [ not_symbol ], ( '(', [whitespace] , logic_term_group, [whitespace], ')'  | group_elem );
group_elem = simple_term | phrase_term | single_range_term | double_range_term ;

(* compound term *)
range_term = ( double_range_term | single_range_term ), [ boost_modifier ] ;
fuzzy_term = ( simple_term | phrase_term ), [ fuzzy_modifier | boost_modifier ] ;

(* simple term *)
double_range_term = ('[' | '{' ), [whitespace], range_value, whitespace, 'TO', whitespace, range_value, [whitespace], ( ']' | '}' ) ;
single_range_term = [ ('>' | '<'), ['='] ], range_value ;
range_value       = phrase_term | (identifier | number | '.' | '+' | '-' | '|' | '/' | ':') { (identifier | '+' | '-' | dot | ) } | '*' ;
simple_term      = (identifier | number | '+' | '-'), { simple_term_char } ;
phrase_term      = quote, phrase_term_char, {phrase_term_char}, quote ;
regexp_term      = '/', regexp_term_char, { regexp_term_char }, '/' ;
phrase_term_char = ( -quote | '\\', quote ) ;
simple_term_char = identifier | number | dot | '?' | '*' | '-' | '+' | '|' | '/' ;
regexp_term_char = ( -'/' | '\\', '/') ;

(* bool operator *)
and_symbol = whitespace , (( '&', '&' ) | 'AND' | 'and' ), whitespace ;
or_symbol  = whitespace , (( '|', '|' ) | 'OR' | 'or' ), whitespace ;
not_symbol = ('!' , [whitespace] ) | (('NOT' | 'not' ), whitespace ;

(* modifier *)
fuzzy_modifier = '~', [ float ] ;
boost_modifier = '^', [ float ] ;

(* basic element *)
identifier = ident_char , { ident_char } ;
number     = digit , { digit } ;
float      = digit , { digit }, [ dot, digit, { digit } ] ;
escape     = '-' | '+' | '!' | '&' | '|' | '?' | '*' | '\\' | '(' | ')' | '[' | ']' | '{' | '}' | '/' | '<' | '>' | '=' | '~' | '^'  | ':' ;
compare    = ('<' | '>')，[ '=' ] ;
ident_char = ( -( escape | digit | dot | whitespace_char | quote ) | '\\' , (escape | whitespace_char) ) ;
digit      = '0' ... '9' ;
whitespace = whitespace_char , { whitespace_char };
whitespace_char = '\t' | '\r' | '\f' | ' ' ;
quote      = '"' ;
eol        = '\n' ;
dot        = '.' ;
```
