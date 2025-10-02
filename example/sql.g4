

program
    : statement^ ;

statement
    : selectstmt ';\n'
    | createtablestmt ';\n'
    | insertstmt ';\n' ;

selectstmt
    : 'SELECT ' selectlist ' FROM ' identifier ( ' WHERE ' condition )? ;

selectlist
    : '*'
    | column (', ' column)* ;

column
    : identifier ;

createtablestmt
    : 'CREATE TABLE ' identifier '(' columnlist ')' ;

columnlist
    : columndef (', ' columndef)* ;

columndef
    : identifier ' ' datatype ;

insertstmt
    : 'INSERT INTO ' identifier '(' columnlistshort ')' ' VALUES(' valuelist ')' ;

columnlistshort
    : identifier (', ' identifier)* ;

valuelist
    : value (', ' value)* ;

value
    : integer
    | float
    | string
    | boolean ;

condition
    : identifier comparator value (logicalop identifier comparator value)* ;

comparator
    : ' = ' | ' < ' | ' > ' | ' <= ' | ' >= ' | ' != ' ;

logicalop
    : ' AND ' | ' OR ' ;

datatype
    : 'INT'
    | 'FLOAT'
    | 'TEXT'
    | 'BOOL' ;

boolean
    : 'TRUE'
    | 'FALSE' ;

string
    : [A-Z0-9_ ] ;

identifier
    : [a-z] ;

integer
    : [0-9]+ ;

float
    : [0-9]+ '.' [0-9]+ ;
