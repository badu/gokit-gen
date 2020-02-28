// Code generated from antlr.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // antlr

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseantlrListener is a complete listener for a parse tree produced by antlrParser.
type BaseantlrListener struct{}

var _ antlrListener = &BaseantlrListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseantlrListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseantlrListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseantlrListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseantlrListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProto is called when production proto is entered.
func (s *BaseantlrListener) EnterProto(ctx *ProtoContext) {}

// ExitProto is called when production proto is exited.
func (s *BaseantlrListener) ExitProto(ctx *ProtoContext) {}

// EnterSyntax is called when production syntax is entered.
func (s *BaseantlrListener) EnterSyntax(ctx *SyntaxContext) {}

// ExitSyntax is called when production syntax is exited.
func (s *BaseantlrListener) ExitSyntax(ctx *SyntaxContext) {}

// EnterImports is called when production imports is entered.
func (s *BaseantlrListener) EnterImports(ctx *ImportsContext) {}

// ExitImports is called when production imports is exited.
func (s *BaseantlrListener) ExitImports(ctx *ImportsContext) {}

// EnterPackageName is called when production packageName is entered.
func (s *BaseantlrListener) EnterPackageName(ctx *PackageNameContext) {}

// ExitPackageName is called when production packageName is exited.
func (s *BaseantlrListener) ExitPackageName(ctx *PackageNameContext) {}

// EnterOption is called when production option is entered.
func (s *BaseantlrListener) EnterOption(ctx *OptionContext) {}

// ExitOption is called when production option is exited.
func (s *BaseantlrListener) ExitOption(ctx *OptionContext) {}

// EnterOptionName is called when production optionName is entered.
func (s *BaseantlrListener) EnterOptionName(ctx *OptionNameContext) {}

// ExitOptionName is called when production optionName is exited.
func (s *BaseantlrListener) ExitOptionName(ctx *OptionNameContext) {}

// EnterOptionDecl is called when production optionDecl is entered.
func (s *BaseantlrListener) EnterOptionDecl(ctx *OptionDeclContext) {}

// ExitOptionDecl is called when production optionDecl is exited.
func (s *BaseantlrListener) ExitOptionDecl(ctx *OptionDeclContext) {}

// EnterOptionDeclVar is called when production optionDeclVar is entered.
func (s *BaseantlrListener) EnterOptionDeclVar(ctx *OptionDeclVarContext) {}

// ExitOptionDeclVar is called when production optionDeclVar is exited.
func (s *BaseantlrListener) ExitOptionDeclVar(ctx *OptionDeclVarContext) {}

// EnterMessage is called when production message is entered.
func (s *BaseantlrListener) EnterMessage(ctx *MessageContext) {}

// ExitMessage is called when production message is exited.
func (s *BaseantlrListener) ExitMessage(ctx *MessageContext) {}

// EnterMessageDecl is called when production messageDecl is entered.
func (s *BaseantlrListener) EnterMessageDecl(ctx *MessageDeclContext) {}

// ExitMessageDecl is called when production messageDecl is exited.
func (s *BaseantlrListener) ExitMessageDecl(ctx *MessageDeclContext) {}

// EnterEnum is called when production enum is entered.
func (s *BaseantlrListener) EnterEnum(ctx *EnumContext) {}

// ExitEnum is called when production enum is exited.
func (s *BaseantlrListener) ExitEnum(ctx *EnumContext) {}

// EnterEnumDecl is called when production enumDecl is entered.
func (s *BaseantlrListener) EnterEnumDecl(ctx *EnumDeclContext) {}

// ExitEnumDecl is called when production enumDecl is exited.
func (s *BaseantlrListener) ExitEnumDecl(ctx *EnumDeclContext) {}

// EnterEnumField is called when production enumField is entered.
func (s *BaseantlrListener) EnterEnumField(ctx *EnumFieldContext) {}

// ExitEnumField is called when production enumField is exited.
func (s *BaseantlrListener) ExitEnumField(ctx *EnumFieldContext) {}

// EnterKeyValue is called when production keyValue is entered.
func (s *BaseantlrListener) EnterKeyValue(ctx *KeyValueContext) {}

// ExitKeyValue is called when production keyValue is exited.
func (s *BaseantlrListener) ExitKeyValue(ctx *KeyValueContext) {}

// EnterService is called when production service is entered.
func (s *BaseantlrListener) EnterService(ctx *ServiceContext) {}

// ExitService is called when production service is exited.
func (s *BaseantlrListener) ExitService(ctx *ServiceContext) {}

// EnterRpcName is called when production rpcName is entered.
func (s *BaseantlrListener) EnterRpcName(ctx *RpcNameContext) {}

// ExitRpcName is called when production rpcName is exited.
func (s *BaseantlrListener) ExitRpcName(ctx *RpcNameContext) {}

// EnterRpcParams is called when production rpcParams is entered.
func (s *BaseantlrListener) EnterRpcParams(ctx *RpcParamsContext) {}

// ExitRpcParams is called when production rpcParams is exited.
func (s *BaseantlrListener) ExitRpcParams(ctx *RpcParamsContext) {}

// EnterClientStreams is called when production clientStreams is entered.
func (s *BaseantlrListener) EnterClientStreams(ctx *ClientStreamsContext) {}

// ExitClientStreams is called when production clientStreams is exited.
func (s *BaseantlrListener) ExitClientStreams(ctx *ClientStreamsContext) {}

// EnterRpcReturns is called when production rpcReturns is entered.
func (s *BaseantlrListener) EnterRpcReturns(ctx *RpcReturnsContext) {}

// ExitRpcReturns is called when production rpcReturns is exited.
func (s *BaseantlrListener) ExitRpcReturns(ctx *RpcReturnsContext) {}

// EnterServerStreams is called when production serverStreams is entered.
func (s *BaseantlrListener) EnterServerStreams(ctx *ServerStreamsContext) {}

// ExitServerStreams is called when production serverStreams is exited.
func (s *BaseantlrListener) ExitServerStreams(ctx *ServerStreamsContext) {}

// EnterRpc is called when production rpc is entered.
func (s *BaseantlrListener) EnterRpc(ctx *RpcContext) {}

// ExitRpc is called when production rpc is exited.
func (s *BaseantlrListener) ExitRpc(ctx *RpcContext) {}

// EnterReserved is called when production reserved is entered.
func (s *BaseantlrListener) EnterReserved(ctx *ReservedContext) {}

// ExitReserved is called when production reserved is exited.
func (s *BaseantlrListener) ExitReserved(ctx *ReservedContext) {}

// EnterRanges is called when production ranges is entered.
func (s *BaseantlrListener) EnterRanges(ctx *RangesContext) {}

// ExitRanges is called when production ranges is exited.
func (s *BaseantlrListener) ExitRanges(ctx *RangesContext) {}

// EnterRangeDef is called when production rangeDef is entered.
func (s *BaseantlrListener) EnterRangeDef(ctx *RangeDefContext) {}

// ExitRangeDef is called when production rangeDef is exited.
func (s *BaseantlrListener) ExitRangeDef(ctx *RangeDefContext) {}

// EnterFieldNames is called when production fieldNames is entered.
func (s *BaseantlrListener) EnterFieldNames(ctx *FieldNamesContext) {}

// ExitFieldNames is called when production fieldNames is exited.
func (s *BaseantlrListener) ExitFieldNames(ctx *FieldNamesContext) {}

// EnterFieldType is called when production fieldType is entered.
func (s *BaseantlrListener) EnterFieldType(ctx *FieldTypeContext) {}

// ExitFieldType is called when production fieldType is exited.
func (s *BaseantlrListener) ExitFieldType(ctx *FieldTypeContext) {}

// EnterField is called when production field is entered.
func (s *BaseantlrListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseantlrListener) ExitField(ctx *FieldContext) {}

// EnterFieldOpts is called when production fieldOpts is entered.
func (s *BaseantlrListener) EnterFieldOpts(ctx *FieldOptsContext) {}

// ExitFieldOpts is called when production fieldOpts is exited.
func (s *BaseantlrListener) ExitFieldOpts(ctx *FieldOptsContext) {}

// EnterOneof is called when production oneof is entered.
func (s *BaseantlrListener) EnterOneof(ctx *OneofContext) {}

// ExitOneof is called when production oneof is exited.
func (s *BaseantlrListener) ExitOneof(ctx *OneofContext) {}

// EnterOneofField is called when production oneofField is entered.
func (s *BaseantlrListener) EnterOneofField(ctx *OneofFieldContext) {}

// ExitOneofField is called when production oneofField is exited.
func (s *BaseantlrListener) ExitOneofField(ctx *OneofFieldContext) {}

// EnterMapField is called when production mapField is entered.
func (s *BaseantlrListener) EnterMapField(ctx *MapFieldContext) {}

// ExitMapField is called when production mapField is exited.
func (s *BaseantlrListener) ExitMapField(ctx *MapFieldContext) {}

// EnterKeyType is called when production keyType is entered.
func (s *BaseantlrListener) EnterKeyType(ctx *KeyTypeContext) {}

// ExitKeyType is called when production keyType is exited.
func (s *BaseantlrListener) ExitKeyType(ctx *KeyTypeContext) {}

// EnterReservedWord is called when production reservedWord is entered.
func (s *BaseantlrListener) EnterReservedWord(ctx *ReservedWordContext) {}

// ExitReservedWord is called when production reservedWord is exited.
func (s *BaseantlrListener) ExitReservedWord(ctx *ReservedWordContext) {}

// EnterFullIdent is called when production fullIdent is entered.
func (s *BaseantlrListener) EnterFullIdent(ctx *FullIdentContext) {}

// ExitFullIdent is called when production fullIdent is exited.
func (s *BaseantlrListener) ExitFullIdent(ctx *FullIdentContext) {}

// EnterMessageOrEnumType is called when production messageOrEnumType is entered.
func (s *BaseantlrListener) EnterMessageOrEnumType(ctx *MessageOrEnumTypeContext) {}

// ExitMessageOrEnumType is called when production messageOrEnumType is exited.
func (s *BaseantlrListener) ExitMessageOrEnumType(ctx *MessageOrEnumTypeContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseantlrListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseantlrListener) ExitConstant(ctx *ConstantContext) {}
