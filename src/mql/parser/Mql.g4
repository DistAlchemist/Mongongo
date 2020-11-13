// Mql.g4
grammar Mql;

// Tokens 
GET: 'GET';
SET: 'SET';
WHITESPACE: [ \r\n\t]+ -> skip;
ASSOC: '=>';
COMMA: ',';
LEFT_BRACE: '{';
RIGHT_BRACE: '}';
SEMICOLON: ';';

fragment 
Letter
    : 'a'..'z'
    | 'A'..'Z'
    ;

fragment
Digit
    : '0'..'9'
    ;

Identifier
    : Letter ( Letter | Digit | '_')*
    ;

StringLiteral 
    : '\'' (~'\'')* '\'' ( '\'' (~'\'')* '\'' )*
    ;

IntegerLiteral
    : Digit+
    ;

stringVal
    : '?'
    | StringLiteral
    ;

// Rules
stmt
    : getStmt
    | setStmt
    ;

getStmt
    : GET columnSpec
    ;

setStmt
    : SET columnSpec '=' valueExpr
    ;

columnSpec
    : tableName '.' columnFamilyName '[' rowKey ']'
        ( '[' a+=columnOrSuperColumnKey ']'
            ('[' a+=columnOrSuperColumnKey ']')?
        )?
    ;

tableName: Identifier;

columnFamilyName: Identifier;

valueExpr
    : cellValue
    | columnMapValue
    | superColumnMapValue
    ;

cellValue
    : stringVal
    ;

columnMapValue
    : LEFT_BRACE columnMapEntry (COMMA columnMapEntry)* RIGHT_BRACE
    ;

superColumnMapValue
    : LEFT_BRACE superColumnMapEntry (COMMA superColumnMapEntry)* RIGHT_BRACE
    ;

columnMapEntry
    : columnKey ASSOC cellValue
    ;

superColumnMapEntry
    : superColumnKey ASSOC columnMapValue
    ;

columnOrSuperColumnName: Identifier;

rowKey: stringVal;
columnOrSuperColumnKey: stringVal;
columnKey: stringVal;
superColumnKey: stringVal;

