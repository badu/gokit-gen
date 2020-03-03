// Code generated from antlr.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // antlr

import "github.com/antlr/antlr4/runtime/Go/antlr"

// antlrListener is a complete listener for a parse tree produced by antlrParser.
type antlrListener interface {
	antlr.ParseTreeListener

	// EnterProto is called when entering the proto production.
	EnterProto(c *ProtoContext)

	// EnterSyntax is called when entering the syntax production.
	EnterSyntax(c *SyntaxContext)

	// EnterImports is called when entering the imports production.
	EnterImports(c *ImportsContext)

	// EnterPackageName is called when entering the packageName production.
	EnterPackageName(c *PackageNameContext)

	// EnterOption is called when entering the option production.
	EnterOption(c *OptionContext)

	// EnterOptionName is called when entering the optionName production.
	EnterOptionName(c *OptionNameContext)

	// EnterOptionDecl is called when entering the optionDecl production.
	EnterOptionDecl(c *OptionDeclContext)

	// EnterOptionDeclVar is called when entering the optionDeclVar production.
	EnterOptionDeclVar(c *OptionDeclVarContext)

	// EnterExtend is called when entering the extend production.
	EnterExtend(c *ExtendContext)

	// EnterExtendDecl is called when entering the extendDecl production.
	EnterExtendDecl(c *ExtendDeclContext)

	// EnterExtendName is called when entering the extendName production.
	EnterExtendName(c *ExtendNameContext)

	// EnterMessage is called when entering the message production.
	EnterMessage(c *MessageContext)

	// EnterMessageDecl is called when entering the messageDecl production.
	EnterMessageDecl(c *MessageDeclContext)

	// EnterEnum is called when entering the enum production.
	EnterEnum(c *EnumContext)

	// EnterEnumDecl is called when entering the enumDecl production.
	EnterEnumDecl(c *EnumDeclContext)

	// EnterEnumField is called when entering the enumField production.
	EnterEnumField(c *EnumFieldContext)

	// EnterKeyValue is called when entering the keyValue production.
	EnterKeyValue(c *KeyValueContext)

	// EnterService is called when entering the service production.
	EnterService(c *ServiceContext)

	// EnterRpcName is called when entering the rpcName production.
	EnterRpcName(c *RpcNameContext)

	// EnterRpcParams is called when entering the rpcParams production.
	EnterRpcParams(c *RpcParamsContext)

	// EnterClientStreams is called when entering the clientStreams production.
	EnterClientStreams(c *ClientStreamsContext)

	// EnterRpcReturns is called when entering the rpcReturns production.
	EnterRpcReturns(c *RpcReturnsContext)

	// EnterServerStreams is called when entering the serverStreams production.
	EnterServerStreams(c *ServerStreamsContext)

	// EnterRpc is called when entering the rpc production.
	EnterRpc(c *RpcContext)

	// EnterReserved is called when entering the reserved production.
	EnterReserved(c *ReservedContext)

	// EnterRanges is called when entering the ranges production.
	EnterRanges(c *RangesContext)

	// EnterRangeDef is called when entering the rangeDef production.
	EnterRangeDef(c *RangeDefContext)

	// EnterFieldNames is called when entering the fieldNames production.
	EnterFieldNames(c *FieldNamesContext)

	// EnterFieldType is called when entering the fieldType production.
	EnterFieldType(c *FieldTypeContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterFieldOpts is called when entering the fieldOpts production.
	EnterFieldOpts(c *FieldOptsContext)

	// EnterOneof is called when entering the oneof production.
	EnterOneof(c *OneofContext)

	// EnterOneofField is called when entering the oneofField production.
	EnterOneofField(c *OneofFieldContext)

	// EnterMapField is called when entering the mapField production.
	EnterMapField(c *MapFieldContext)

	// EnterKeyType is called when entering the keyType production.
	EnterKeyType(c *KeyTypeContext)

	// EnterReservedWord is called when entering the reservedWord production.
	EnterReservedWord(c *ReservedWordContext)

	// EnterFullIdent is called when entering the fullIdent production.
	EnterFullIdent(c *FullIdentContext)

	// EnterMessageOrEnumType is called when entering the messageOrEnumType production.
	EnterMessageOrEnumType(c *MessageOrEnumTypeContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// ExitProto is called when exiting the proto production.
	ExitProto(c *ProtoContext)

	// ExitSyntax is called when exiting the syntax production.
	ExitSyntax(c *SyntaxContext)

	// ExitImports is called when exiting the imports production.
	ExitImports(c *ImportsContext)

	// ExitPackageName is called when exiting the packageName production.
	ExitPackageName(c *PackageNameContext)

	// ExitOption is called when exiting the option production.
	ExitOption(c *OptionContext)

	// ExitOptionName is called when exiting the optionName production.
	ExitOptionName(c *OptionNameContext)

	// ExitOptionDecl is called when exiting the optionDecl production.
	ExitOptionDecl(c *OptionDeclContext)

	// ExitOptionDeclVar is called when exiting the optionDeclVar production.
	ExitOptionDeclVar(c *OptionDeclVarContext)

	// ExitExtend is called when exiting the extend production.
	ExitExtend(c *ExtendContext)

	// ExitExtendDecl is called when exiting the extendDecl production.
	ExitExtendDecl(c *ExtendDeclContext)

	// ExitExtendName is called when exiting the extendName production.
	ExitExtendName(c *ExtendNameContext)

	// ExitMessage is called when exiting the message production.
	ExitMessage(c *MessageContext)

	// ExitMessageDecl is called when exiting the messageDecl production.
	ExitMessageDecl(c *MessageDeclContext)

	// ExitEnum is called when exiting the enum production.
	ExitEnum(c *EnumContext)

	// ExitEnumDecl is called when exiting the enumDecl production.
	ExitEnumDecl(c *EnumDeclContext)

	// ExitEnumField is called when exiting the enumField production.
	ExitEnumField(c *EnumFieldContext)

	// ExitKeyValue is called when exiting the keyValue production.
	ExitKeyValue(c *KeyValueContext)

	// ExitService is called when exiting the service production.
	ExitService(c *ServiceContext)

	// ExitRpcName is called when exiting the rpcName production.
	ExitRpcName(c *RpcNameContext)

	// ExitRpcParams is called when exiting the rpcParams production.
	ExitRpcParams(c *RpcParamsContext)

	// ExitClientStreams is called when exiting the clientStreams production.
	ExitClientStreams(c *ClientStreamsContext)

	// ExitRpcReturns is called when exiting the rpcReturns production.
	ExitRpcReturns(c *RpcReturnsContext)

	// ExitServerStreams is called when exiting the serverStreams production.
	ExitServerStreams(c *ServerStreamsContext)

	// ExitRpc is called when exiting the rpc production.
	ExitRpc(c *RpcContext)

	// ExitReserved is called when exiting the reserved production.
	ExitReserved(c *ReservedContext)

	// ExitRanges is called when exiting the ranges production.
	ExitRanges(c *RangesContext)

	// ExitRangeDef is called when exiting the rangeDef production.
	ExitRangeDef(c *RangeDefContext)

	// ExitFieldNames is called when exiting the fieldNames production.
	ExitFieldNames(c *FieldNamesContext)

	// ExitFieldType is called when exiting the fieldType production.
	ExitFieldType(c *FieldTypeContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitFieldOpts is called when exiting the fieldOpts production.
	ExitFieldOpts(c *FieldOptsContext)

	// ExitOneof is called when exiting the oneof production.
	ExitOneof(c *OneofContext)

	// ExitOneofField is called when exiting the oneofField production.
	ExitOneofField(c *OneofFieldContext)

	// ExitMapField is called when exiting the mapField production.
	ExitMapField(c *MapFieldContext)

	// ExitKeyType is called when exiting the keyType production.
	ExitKeyType(c *KeyTypeContext)

	// ExitReservedWord is called when exiting the reservedWord production.
	ExitReservedWord(c *ReservedWordContext)

	// ExitFullIdent is called when exiting the fullIdent production.
	ExitFullIdent(c *FullIdentContext)

	// ExitMessageOrEnumType is called when exiting the messageOrEnumType production.
	ExitMessageOrEnumType(c *MessageOrEnumTypeContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)
}
