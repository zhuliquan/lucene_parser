# lucene_parser
## Introduction:
This package can parse lucene query used by ES (ElasticSearch), this package is pure go package.
## Features
- 1、support phrase term query, for instance `x:"foo bar"`.
- 2、support regexp term query, for instance `x:/\d+\\.?\d+/`.
- 3、support bool operator （i.e. `AND`, `OR`, `NOT`, `&&`, `||`, `!`） join sub query, for instance `x:1 AND y:2`, `x:1 || y:2`.
- 4、support bound range query,  for instance `x:[1 TO 2]`, `x:[1 TO 2}`.
- 5、support side range query, for instance `x:>1` , `x:>=1` , `x:<1` , `x:<=1`.
- 6、support boost modifier, for instance `x:1^2` , `x:"dsada 8908"^3`.
- 7、support fuzzy query, for instance `x:for~2` , `x:"foo bar"~2`.
- 8、support term group query, for instance `x:(foo OR bar)`, `x:(>1 && <2)`.

## Limitations
- 1、only support lucene query with **field name**, instead of query without **field name** (i.e. this project can't parse query like `foo OR bar`, `foo AND bar`).
- 2、don't support prefix operator `'+'` / `'-'`, for instance `+fo0 -bar`.

## Usage
```golang

```

## EBNF of Lucene
lucene parser will convert string of lucene query to ast, according to EBNF of lucene. EBNF of lucene is below.

```ebnf
(* lucene expression *)
lucene = or_query, { or_sym_query } ;
or_sym_query  = or_symbol, [ not_symbol ], or_query ;
or_query      = and_query, { and_sym_query } ;
and_sym_query = and_symbol, [ not_symbol ], and_query ;
and_query     = ( not_symbol, lucene ) | ( '(', lucene, ')' ) | ( field, term ) ;

(* field and term *)
field_char       = identifier | '-' | number | dot ;
field            = field_char, { field_char }, ':' ;
term = range_term | fuzzy_term | regexp_term | term_group ;

(* term group *)
term_group = logic_term_group, [ boost_modifier ] ;
logic_term_group   = or_term_group, { or_sym_term_group } ;
or_sym_term_group  = or_symbol, [ not_symbol ], or_term_group ;
or_term_group      = and_term_group, { and_sym_term_group } ;
and_sym_term_group = and_symbol, [ not_symbol ], and_term_group ;
and_term_group     = ( not_symbol，logic_term_group ) | ( '(' , logic_term_group, ')' ) | group_elem ;
group_elem = simple_term | phrase_term | single_range_term | double_range_term ;

(* compound term *)
range_term = ( double_range_term | single_range_term ) [ boost_modifier ] ;
fuzzy_term = ( simple_term | phrase_term ), [ fuzzy_modifier | boost_modifier ] ;

(* simple term *)
double_range_term = ('[' | '{' ), [whitespace], range_value, whitespace, 'TO', whitespace, range_value, [whitespace], ( ']' | '}' ) ;
single_range_term = [ ('>' | '<'), ['='] ], range_value ;
range_value       = phrase_term | (identifier | '+' | '-' | dot | ), { (identifier | '+' | '-' | dot | ) } | '*' ;
simple_term      = simple_term_char, { simple_term_char } ;
phrase_term      = quote, phrase_term_char, {phrase_term_char}, quote ;
regexp_term      = '/', regexp_term_char, { regexp_term_char }, '/' ;
phrase_term_char = ( -quote | '\\', quote ) ;
simple_term_char = identifier | number | dot | '?' | '*' ;
regexp_term_char = ( -'/' | '\\', '/') ;

(* bool operator *)
and_symbol = whitespace, (( '&'，'&' ) | 'AND' | 'and' ), whitespace ;
or_symbol  = whitespace, (( '|'，'|' ) | 'OR' | 'or' ), whitespace ;
not_symbol = ('!' | 'NOT' | 'not' ), whitespace ;

(* modifier *)
fuzzy_modifier = '~', [ number ] ;
boost_modifier = '^', float ;

(* basic element *)
identifier = ident_char, { ident_char } ;
number     = digit , { digit } ;
float      = digit , { digit }, [ dot, digit, { digit } ] ;
escape     = '-' | '+' | '!' | '&' | '|' | '?' | '*' | '\\' | '(' | ')' | '[' | ']' | '{' | '}' | '/' | '<' | '>' | '=' | '~' | '^'  | ':' ;
compare    = ('<' | '>')，[ '=' ] ;
ident_char = ( -( escape | digit | dot | whitespace | quote ) | '\\' , (escape | whitespace) ) ;
digit      = '0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9' ;
whitespace = '\t' | '\r' | '\f' | ' ' ;
quote      = '"' ;
eol        = '\n' ;
dot        = '.' ;
```

