program
   : statement+
   ;

statement
   : 'if' paren_expr statement
   | 'if' paren_expr statement 'else' statement
   | 'while' paren_expr statement
   | 'do' statement 'while' paren_expr ';'
   | '{' statement* '}'
   | expr ';'
   | ';'
   ;

paren_expr
   : '(' expr ')'
   ;

expr
   : test
   | id '=' expr
   ;

test
   : sum
   | sum '<' sum
   ;

sum
   : term
   | sum '+' term
   | sum '-' term
   ;

term
   : id
   | integer
   | paren_expr
   ;

id
   : STRING
   ;

integer
   : INT
   ;


STRING
   : [a-z]+
   ;

INT
   : [0-9]+
   ;


parenthesized expressions are highest priority,
followed by -> and . (dot)
followed by * and /
folled by + and -
followed by &&
followed by ||, which in this simplified dialect of C has the least precedence
