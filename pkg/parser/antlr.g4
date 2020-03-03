/**
 * A Protocol Buffers 3 grammar for ANTLR v4.
 *
 * Derived and adapted from:
 * https://developers.google.com/protocol-buffers/docs/reference/proto3-spec
 *
 * @author Marco Willemart
 *
 * Taken from https://github.com/antlr/grammars-v4/blob/master/protobuf3/Protobuf3.g4
 */

grammar antlr;

//
// Proto file
//

proto
    :   syntax (   imports
               |   packageName
               |   option
               |   message
               |   enum
               |   service
               |   extend
               |   SEMI
               )*
        EOF
    ;

//
// Syntax
//

syntax : SYNTAX EQ (PROTO3_DOUBLE | PROTO3_SINGLE ) SEMI ;

//
// Import Statement
//

imports : IMPORT (WEAK | PUBLIC)? StrLit SEMI ;

//
// Package
//

packageName : PACKAGE fullIdent SEMI ;

//
// Option
//

option : OPTION optionName EQ (constant | optionDecl)  SEMI ;

optionName : (Ident | LPAREN fullIdent RPAREN ) (DOT (Ident | reservedWord))* ;

optionDecl : LBRACE (optionDeclVar)* RBRACE ;

optionDeclVar : optionName ':' constant ;

//
// Extend
//

extend : EXTEND extendName extendDecl ;

extendDecl :   LBRACE ( field )* RBRACE ;

extendName : ( Ident | fullIdent );

// Message definition

message : MESSAGE Ident messageDecl ;

messageDecl
    :   LBRACE (    field
                |   enum
                |   message
                |   option
                |   oneof
                |   mapField
                |   reserved
                |   SEMI
                )*
       RBRACE
    ;

// Enum definition

enum : ENUM Ident enumDecl ;

enumDecl :   
        LBRACE (   option
            |   enumField
            |   SEMI
            )*
        RBRACE
    ;

enumField : Ident EQ MINUS? IntLit (LBRACK keyValue (COMMA  keyValue)* RBRACK)? SEMI ;

keyValue : optionName EQ constant ;

// Service definition

service
    :   SERVICE Ident LBRACE (   option
                                  |   rpc
                                  // not defined in the protobuf specification
                                  //|   stream
                                  |   SEMI
                                  )*
        RBRACE
    ;

rpcName : Ident ;
rpcParams : (Ident DOT)* Ident ;
clientStreams : STREAM? ;
rpcReturns : (Ident DOT)* Ident ;
serverStreams : STREAM? ;

rpc
    :   RPC rpcName LPAREN clientStreams DOT? rpcParams RPAREN
        RETURNS LPAREN serverStreams DOT? rpcReturns RPAREN ((LBRACE (option | SEMI)* RBRACE) | SEMI)
    ;

//
// Reserved
//

reserved : RESERVED (ranges | fieldNames) SEMI ;

ranges : rangeDef (COMMA rangeDef)* ;

rangeDef
    :   IntLit
    |   IntLit TO IntLit
    ;

fieldNames : StrLit (COMMA StrLit)* ;

//s
// Fields
//

fieldType
    :   (   DOUBLE
        |   FLOAT
        |   INT32
        |   INT64
        |   UINT32
        |   UINT64
        |   SINT32
        |   SINT64
        |   FIXED32
        |   FIXED64
        |   SFIXED32
        |   SFIXED64
        |   BOOL
        |   STRING
        |   BYTES
        )
    |   messageOrEnumType
    ;

// Normal field

field : REPEATED? fieldType (Ident | reservedWord) EQ IntLit (LBRACK fieldOpts RBRACK)? SEMI ;

fieldOpts : keyValue (COMMA  keyValue)* ;

// Oneof and oneof field

oneof : ONEOF Ident LBRACE (oneofField | SEMI)* RBRACE ;

oneofField : fieldType (Ident | reservedWord) EQ IntLit (LBRACK fieldOpts RBRACK)? SEMI ;

// Map field

mapField : MAP LCHEVR keyType COMMA fieldType RCHEVR Ident EQ IntLit (LBRACK fieldOpts RBRACK)? SEMI ;

keyType
    :   INT32
    |   INT64
    |   UINT32
    |   UINT64
    |   SINT32
    |   SINT64
    |   FIXED32
    |   FIXED64
    |   SFIXED32
    |   SFIXED64
    |   BOOL
    |   STRING
    ;

reservedWord
    :   MESSAGE
    |   OPTION
    |   PACKAGE
    |   SERVICE
    |   STREAM
    |   STRING
    |   SYNTAX
    |   WEAK
    |   RPC
    |   EXTEND
    ;
//
// Lexical elements
//

// Keywords and literals

BOOL            : 'bool';
BYTES           : 'bytes';
DOUBLE          : 'double';
ENUM            : 'enum';
EXTEND          : 'extend';
FIXED32         : 'fixed32';
FIXED64         : 'fixed64';
FLOAT           : 'float';
IMPORT          : 'import';
INT32           : 'int32';
INT64           : 'int64';
MAP             : 'map';
MESSAGE         : 'message';
ONEOF           : 'oneof';
OPTION          : 'option';
PACKAGE         : 'package';
PROTO3_DOUBLE   : '"proto3"';
PROTO3_SINGLE   : '\'proto3\'';
PUBLIC          : 'public';
REPEATED        : 'repeated';
RESERVED        : 'reserved';
RETURNS         : 'returns';
RPC             : 'rpc';
SERVICE         : 'service';
SFIXED32        : 'sfixed32';
SFIXED64        : 'sfixed64';
SINT32          : 'sint32';
SINT64          : 'sint64';
STREAM          : 'stream';
STRING          : 'string';
SYNTAX          : 'syntax';
TO              : 'to';
UINT32          : 'uint32';
UINT64          : 'uint64';
WEAK            : 'weak';
LPAREN          : '(';
RPAREN          : ')';
LBRACE          : '{';
RBRACE          : '}';
LBRACK          : '[';
RBRACK          : ']';
LCHEVR          : '<';
RCHEVR          : '>';
SEMI            : ';';
COMMA           : ',';
DOT             : '.';
MINUS           : '-';
PLUS            : '+';
SLASH           : '\\';
EQ              : '=';

// Letters and digits

fragment
Letter  :   [A-Za-z_] ;

fragment
DecimalDigit :   [0-9] ;

fragment
OctalDigit :   [0-7] ;

fragment
HexDigit :   [0-9A-Fa-f] ;

// Identifiers

Ident : Letter (Letter | DecimalDigit)* ;

fullIdent : Ident (DOT Ident)* ;

messageOrEnumType : DOT? ( (Ident | reservedWord) DOT)* Ident ;

// Integer literals

IntLit
    :   DecimalLit
    |   OctalLit
    |   HexLit
    ;

fragment
DecimalLit    :   [1-9] DecimalDigit* ;

fragment
OctalLit    :   '0' OctalDigit* ;

fragment
HexLit    :   '0' ('x' | 'X') HexDigit+ ;

// Floating-point literals

FloatLit
    :   (   Decimals DOT Decimals? Exponent?
        |   Decimals Exponent
        |   DOT Decimals Exponent?
        )
    |   'inf'
    |   'nan'
    ;

fragment
Decimals    :   DecimalDigit+ ;

fragment
Exponent    :   ('e' | 'E') (PLUS | MINUS)? Decimals ;

// Boolean

BoolLit
    :   'true'
    |   'false'
    ;

// String literals

StrLit
    :   '\'' CharValue* '\''
    |   '"' CharValue* '"'
    ;

// Constant

constant
    :   fullIdent
        |   (MINUS | PLUS)? IntLit
        |   (MINUS | PLUS)? FloatLit
        |   (   StrLit
            |   BoolLit
            )
    ;


fragment
CharValue
    :   HexEscape
    |   OctEscape
    |   CharEscape
    |   ~[\u0000\n\\]
    ;

fragment
HexEscape
    :   SLASH ('x' | 'X') HexDigit HexDigit
    ;

fragment
OctEscape
    :   SLASH OctalDigit OctalDigit OctalDigit
    ;

fragment
CharEscape
    :   SLASH [abfnrtv\\'"]
    ;

// Whitespace and comments

WS  :   [ \t\r\n\u000C]+ -> skip
    ;

COMMENT
    :   '/*' .*? '*/' -> skip
    ;

LINE_COMMENT
    :   '//' ~[\r\n]* -> skip
    ;