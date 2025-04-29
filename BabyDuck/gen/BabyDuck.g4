grammar BabyDuck;

// Parser Rules

// Main Program
program : 'program' ID ';' vars funcs 'main' body 'end' ';' ;


// Variable Declaration
vars : ('var' var_list ';')? ;
var_list : var_decl(var_decl)*;
var_decl : id_list ':' type;
id_list : ID (',' ID)*;
type : 'int' | 'float' ;

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
cycle : 'while' '(' exp ')' 'do' body ;

// Conditional
condition : 'if' '(' expression ')' body else_part ;
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
relop : ('>' | '<' )? '!=' ;

// Aritmetic Expressions
exp : term (addop term)* ;
addop : '+' | '-' ;
term : factor (mulop factor)* ;
mulop : '*' | '/' ;

// Factor
factor : '(' exp ')' factorsign
      | ID factorsign
      | INT factorsign
      | FLOAT factorsign ;

factorsign : '*' factor
           | '/' factor
           | ;

// Functions
funcs : func funcs
      | ;
func : functype ID '(' param_list? ')' funcbody ;
functype : 'void'
         | type ;

// Params
param_list : param (',' param)* ;
param : type ID ;

funcbody : '{' vars? statement* '}' ;

f_call : ID '(' arg_list? ')' ';' ;

arg_list : exp (',' exp)* ;


// Lexer Rules


ID : [a-zA-Z_][a-zA-Z_0-9]* ;
INT : [0-9]+ ;
FLOAT : [0-9]+ '.' [0-9]+ ;
STRING : '"' (~["\n])* '"' ;


WS : [ \t\r\n]+ -> skip ;

