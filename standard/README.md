# standard lucene syntax

This package gives standard lucene parser. syntax is refer [standard lucene syntax](https://lucene.apache.org/core/2_9_4/queryparsersyntax.html). and EBNF is refer to [StandardSyntaxParser](https://github.com/apache/lucene/blob/main/lucene/queryparser/src/java/org/apache/lucene/queryparser/flexible/standard/parser/StandardSyntaxParser.jj) and [anltr-v4-lucene](https://github.com/antlr/grammars-v4/tree/master/lucene)

## EBNF

```ebnf
Lucene = Query ;
Query = DisjQuery, { WHITESPACE DisjQuery } ;
DisjQuery = ConjQuery, { WHITESPACE （ "OR" | "or" | "||" ） WHITESPACE ConjQuery } ;
ConjQuery = ModClause, { WHITESPACE （"AND" | "and" | "&&"） WHITESPACE ModClause } ;
ModClause = ('+' | '-' | '!')? Clause ;
Clause = FieldName? ( TermExpr | PhraseExpr | GroupExpr | Regexp | Range ) Boost? ;
FieldName = TERM, ':' ;
TEMR_CHAR = IDENT | NUMBER | ESCAPE | '.' | '-' | '+' | '?' | '*'; 
TERM = TERM_CHAR, { TERM_CHAR };
TermExpr = TERM Fuzzy? ;
Fuzzy = '~', NUMBER? ;
NUMBER = NUMBER_CHAR, [ '.', NUMBER_CHAR ] ;
PhraseExpr = Phrase Fuzzy? ;
Phrase = '"' ( '\\', '"' | !'"' ), { ( '\\', '"' | !'"' ) } '"' ;
GroupExpr = '(', WHITESPACE? Query , WHITESPACE? ')' ;
Regexp = '/' ( '\\' '/' | !'/' ), { ( '\\' '/' | !'/' ) } '/' ;
Range = SingleRange | DoubleRange ;
SingleRange = COMPARE RangeValue ;
RangeValue = TERM | Phrase | Number ;
DoubleRange = ( '{' | '[' ), WHITESPACE? RangeNode WHITESPACE "TO" WHITESPACE RangeNode WHITESPACE? ( '}' | ']' ) ;     
RangeNode = RangeValue | '*' ;
IDENT = IDENT_CHAR , { IDENT_CHAR } ;
ESCAPE     = '-' | '+' | '!' | '&' | '|' | '?' | '*' | '\\' | '(' | ')' | '[' | ']' | '{' | '}' | '/' | '<' | '>' | '=' | '~' | '^'  | ':' ;
COMPARE    = ('<' | '>')，[ '=' ] ;
IDENT_CHAR = ( -( ESCAPE | NUMBER_CHAR | '.' | WHITESPACE | '"' ) | '\\' , (ESCAPE | WHITESPACE) ) ;
NUMBER_CHAR      = '0' ... '9' ;
WHITESPACE = WHITESPACE_CHAR , { WHITESPACE_CHAR };
WHITESPACE_CHAR = '\t' | '\r' | '\f' | ' ' | '0x3000'; （* 0x3000 is Full width whitespace *）
```
