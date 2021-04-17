program
   : statement+
   

statement
   : _if_ paren_expr statement
   | _if_ paren_expr statement _else_ statement
   | _while_ paren_expr statement
   | _do_ statement _while_ paren_expr _;_
   | _{_ statement* _}_
   | expr _;_
   | _;_
   

paren_expr
   : _(_ expr _)_
   

expr
   : test
   | id _=_ expr
   

test
   : sum
   | sum _<_ sum
   

sum
   : term
   | sum _+_ term
   | sum _-_ term
   

term
   : id
   | integer
   | paren_expr
   
   
id
   : STRING
   

integer
   : INT
   


STRING
   : [a-z]+
   
INT
   : [0-9]+
   
