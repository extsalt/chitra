program
   : statement+
   ;

statement
   : __if__ paren_expr statement
   | __if__ paren_expr statement __else__ statement
   | __while__ paren_expr statement
   | __do__ statement __while__ paren_expr __;__
   | __{__ statement* __}__
   | expr __;__
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
