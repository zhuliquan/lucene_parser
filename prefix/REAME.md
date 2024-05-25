# EBNF

```ebnf
(* lucene expression with prefix operator *)
lucene = { prefix_operator_clause } ;
prefix_operator_clause = { WHITESPACE }, ['+' | '-' | '!'], ( (field, term) | '(', [WHITESPACE], lucene, [WHITESPACE], ')' )
term = ramge_term | fuzzy_term | regexp | term_group

( * term group with prefix operator *)
term_group = '(', prefix_term_group, ')', [ boost_modifier ] ;
prefix_term_group = prefix_operator_term, { prefix_operator_term } ;
prefix_operator_term = [ WHITESPACE ], ['+' | '-' | '!'], ( simple_term | phrase_term | single_range_term | double_range_term | '(', [WHITESPACE], prefix_operator_term, [WHITESPACE], ')' )

(* field and term *)
field_char       = identifier | '-' | number | dot ;
field            = field_char, { field_char }, ':' ;
term = range_term | fuzzy_term | regexp_term | term_group ;

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
