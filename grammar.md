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
   | __;__
   ;

paren_expr
   : __(__ expr __)__
   ;

expr
   : test
   | id __=__ expr
   ;

test
   : sum
   | sum __<__ sum
   ;

sum
   : term
   | sum __+__ term
   | sum __-__ term
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
