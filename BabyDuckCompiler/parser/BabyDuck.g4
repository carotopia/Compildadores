grammar BabyDuck;
// Lexer Rules

VOID : 'void' ;
INTTYPE : 'int' ;
FLOATTYPE : 'float' ;

LPAREN : '(' ;
RPAREN : ')' ;
LBRACKET : '[' ;
RBRACKET : ']' ;
COLON : ':' ;
COMMA : ',' ;
SEMICOLON : ';' ;



ID : [a-zA-Z_][a-zA-Z_0-9]* ;
INT : [0-9]+ ;
FLOAT : [0-9]+ '.' [0-9]+ ;
STRING : '"' ~('"')* '"';



WS : [ \t\r\n]+ -> skip ;

// Parser Rules

// Main Program
program: 'program' ID ';' (vars)? (funcs)* 'main' body 'end'  ;




// Variable Declaration
vars : 'var' var_decl+ ;
var_decl : ID (',' ID)* COLON type SEMICOLON ;
type : INTTYPE | FLOATTYPE ;

// Body
body : '{' statement* '}' ;

// Statement
statement : assign
          | cycle
          | f_call
          | print_stmt
          | condition ;

//  Assign
assign : ID '=' exp ';' ;

// While
cycle : 'while' '('expression ')' 'do' body ';' ;

// Conditional
condition : 'if' '(' expression ')' body else_part ';' ;
else_part : ('else' body)? ;

// Print
print_stmt : 'print' '(' printexpr (',' printexpr)* ')' ';' ;
printexpr : exp
          | STRING ;

// Constant
constant : INT
        | FLOAT ;


//Expressions
expression : exp (relop exp)? ;
relop : ('>' | '<' ) | '!=' ;

// Arithmetic Expressions
exp : term (addop term)* ;
addop : '+' | '-' ;
term : factor (mulop factor)* ;
mulop : '*' | '/' ;

// Factor
factor : '(' exp ')' | factorsign;
factorsign : (addop)? value;
value: ID | constant;



// Functions
funcs : func;
func : 'void' ID '(' param_list? ')' funcbody ;
param_list : param (',' param)* ;
param : ID COLON type  ;
funcbody : '[' vars? body ']' ';' ;


f_call : ID '(' arg_list? ')' ';' ;
arg_list : exp (',' exp)* ;



