package proto

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/badu/gokit-gen/pkg/parser"
	"log"
	"strconv"
	"strings"
)

type ProtoListener struct {
	*parser.BaseantlrListener
	Proto Proto
}

type Proto struct {
	Name        string
	Syntax      string
	PackageName string
	Options     []Option
	Imports     []string
	Messages    []Message
	Services    []Service
	Enums       []Enum
	Error       *ErrNode
}

type Service struct {
	Name    string
	Methods []Method
	Options []Option
}

type Method struct {
	Name              string
	In                []Field
	Out               []Field
	Options           []Option
	IsServerStreaming bool
	IsClientStreaming bool
}

type Message struct {
	Name     string
	Fields   []Field
	Options  []Option
	Enums    []Enum
	Messages []Message // nested messages
	IsNested bool
}

type Enum struct {
	Name     string
	Options  []Option
	Fields   []Field
	IsNested bool
}

func (f *Enum) PublicName() string {
	return strings.ToUpper(f.Name[:1]) + f.Name[1:]
}

type Field struct {
	Name       string
	Kind       string
	MapKind    KeyValue
	Pos        string
	Options    []Option
	OneOf      []Field
	IsRepeated bool
	IsMap      bool
	IsOneOf    bool
	IsReserved bool
}

func (f *Field) IsBasic() bool {
	switch f.Kind {
	case "unsafe.Pointer", "bool", "byte",
		"complex64", "complex128",
		"error",
		"float32", "float64", "float",
		"int", "int8", "int16", "int32", "int64",
		"rune", "string",
		"uint", "uint8", "uint16", "uint32", "uint64", "uintptr":
		return true
	}
	return false
}

func getKind(kind string) string {
	switch kind {
	case "bool":
		return "bool"
	case "bytes":
		return "[]byte"
	case "double":
		return "float64"
	case "fixed32":
		return "uint32"
	case "fixed64":
		return "uint64"
	case "float":
		return "float32"
	case "int32":
		return "int32"
	case "int64":
		return "int64"
	case "sfixed32":
		return "int32"
	case "sfixed64":
		return "int64"
	case "sint32":
		return "int32"
	case "sint64":
		return "int64"
	case "string":
		return "string"
	case "uint32":
		return "uint32"
	case "uint64":
		return "uint64"
	}
	return kind
}

func (f *Field) GoKind() string {
	switch f.Kind {
	case "bool":
		return "bool"
	case "bytes":
		return "[]byte"
	case "double":
		return "float64"
	case "fixed32":
		return "uint32"
	case "fixed64":
		return "uint64"
	case "float":
		return "float32"
	case "int32":
		return "int32"
	case "int64":
		return "int64"
	case "map":
		return "map[" + getKind(f.MapKind.Key) + "]" + getKind(f.MapKind.Value)
	case "sfixed32":
		return "int32"
	case "sfixed64":
		return "int64"
	case "sint32":
		return "int32"
	case "sint64":
		return "int64"
	case "string":
		return "string"
	case "uint32":
		return "uint32"
	case "uint64":
		return "uint64"
	}
	return f.Kind
}

func (f *Field) PublicName() string {
	return strings.ToUpper(f.Name[:1]) + f.Name[1:]
}

func (f *Field) IsTimestamp() bool {
	return false
}

type Option struct {
	Name       string
	Value      string
	Values     []KeyValue
	IsConstant bool
}

type KeyValue struct {
	Key   string
	Value string
}

type ErrNode struct {
	Line   int
	Column int
	Text   string
}

type ErrorListener struct {
	*antlr.DefaultErrorListener
	Listener *ProtoListener
}

func (uel *ErrorListener) SyntaxError(rec antlr.Recognizer, offend interface{}, line int, column int, msg string, e antlr.RecognitionException) {
	uel.Listener.Proto.Error = &ErrNode{Line: line, Column: column, Text: "line " + strconv.Itoa(line) + ":" + strconv.Itoa(column) + " " + msg}
}

func getIdent(target *string, node antlr.TerminalNode) {
	if node == nil {
		return
	}
	*target = node.GetText()
}

func extractConstant(ctx parser.IConstantContext) string {
	ct, ok := ctx.(*parser.ConstantContext)
	if !ok {
		log.Print("expecting *ConstantContext")
		return ""
	}
	if ct.MINUS() != nil {
		fmt.Println("minus", ct.MINUS())
	}
	if ct.PLUS() != nil {
		fmt.Println("plus", ct.PLUS())
	}
	if ct.IntLit() != nil {
		return ct.IntLit().GetText()
	}
	if ct.StrLit() != nil {
		result, err := strconv.Unquote(ct.StrLit().GetText())
		if err != nil {
			log.Printf("error unquoting : %v", err)
		}
		return result
	}
	if ct.FloatLit() != nil {
		return ct.FloatLit().GetText()
	}
	if ct.FullIdent() != nil {
		return ct.FullIdent().GetText()
	}
	return ct.GetText()
}

func makeOptionVars(ctx []parser.IOptionDeclVarContext) []KeyValue {
	var result []KeyValue

	for _, dvDecl := range ctx {
		dvCtx, ok := dvDecl.(*parser.OptionDeclVarContext)
		if !ok {
			log.Print("expecting *OptionDeclVarContext")
			continue
		}
		if dvCtx.OptionName() == nil {
			log.Print("expecting OptionName not nil")
			continue
		}
		if dvCtx.Constant() == nil {
			log.Print("expecting Constant not nil")
			continue
		}

		optField := KeyValue{Key: dvCtx.OptionName().GetText(), Value: extractConstant(dvCtx.Constant())}
		result = append(result, optField)
	}
	return result
}

func makeOption(ctx *parser.OptionContext) Option {
	var result Option
	nameCtx, ok := ctx.OptionName().(*parser.OptionNameContext)
	if !ok {
		log.Print("expecting *OptionNameContext")
		return result
	}
	for _, ident := range nameCtx.AllIdent() {
		result.Name += ident.GetText()
	}

	if ctx.OptionDecl() != nil {
		od, ok := ctx.OptionDecl().(*parser.OptionDeclContext)
		if !ok {
			log.Print("expecting *OptionDeclContext")
			return result
		}
		optVals := makeOptionVars(od.AllOptionDeclVar())
		result.Values = append(result.Values, optVals...)
	}

	if ctx.Constant() != nil {
		result.IsConstant = true
		result.Value = extractConstant(ctx.Constant())
	}
	return result
}

func makeOptions(ctx []parser.IOptionContext) []Option {
	var result []Option
	for _, optDecl := range ctx {
		optCtx, ok := optDecl.(*parser.OptionContext)
		if !ok {
			log.Print("expecting *OptionContext")
			continue
		}
		result = append(result, makeOption(optCtx))
	}
	return result
}

func makeKeyValueOptions(ctx []parser.IKeyValueContext) []Option {
	var result []Option
	for _, kvDecl := range ctx {
		kvCtx, ok := kvDecl.(*parser.KeyValueContext)
		if !ok {
			log.Print("expecting *FieldOptContext")
			continue
		}
		if kvCtx.OptionName() == nil {
			log.Print("expecting not nil option name")
			continue
		}
		if kvCtx.Constant() == nil {
			log.Print("expecting not nil constant")
			continue
		}

		opt := Option{Name: kvCtx.OptionName().GetText(), Value: extractConstant(kvCtx.Constant())}
		result = append(result, opt)
	}
	return result
}

func makeEnumFields(ctx []parser.IEnumFieldContext) []Field {
	var result []Field
	for _, enuFieldDecl := range ctx {
		enuFieldCtx, ok := enuFieldDecl.(*parser.EnumFieldContext)
		if !ok {
			log.Print("expecting *EnumFieldContext")
			continue
		}

		field := Field{Name: enuFieldCtx.Ident().GetText(), Kind: enuFieldCtx.IntLit().GetText()}

		options := makeKeyValueOptions(enuFieldCtx.AllKeyValue())
		field.Options = append(field.Options, options...)

		result = append(result, field)
	}
	return result
}

func makeEnums(ctx []parser.IEnumContext) []Enum {
	var result []Enum
	for _, enuDecl := range ctx {
		enumCtx, ok := enuDecl.(*parser.EnumContext)
		if !ok {
			log.Print("expecting *EnumContext")
			continue
		}

		enum := Enum{}
		getIdent(&enum.Name, enumCtx.Ident())

		declCtx, ok := enumCtx.EnumDecl().(*parser.EnumDeclContext)
		if !ok {
			log.Print("expecting *EnumContext")
			continue
		}

		enuOpts := makeOptions(declCtx.AllOption())
		enum.Options = append(enum.Options, enuOpts...)

		enuFields := makeEnumFields(declCtx.AllEnumField())
		enum.Fields = append(enum.Fields, enuFields...)

		result = append(result, enum)
	}
	return result
}

func makeFields(ctx []parser.IFieldContext) []Field {
	var result []Field
	for _, afCtx := range ctx {
		fieldCtx, ok := afCtx.(*parser.FieldContext)
		if !ok {
			log.Print("expecting *FieldContext")
			continue
		}
		if fieldCtx.Ident() == nil {
			log.Print("no ident provided for field")
			continue
		}

		if fieldCtx.IntLit() == nil {
			log.Print("no pos provided for field")
			continue
		}
		if fieldCtx.FieldType() == nil {
			log.Print("no field type provided")
			continue
		}

		field := Field{Name: fieldCtx.Ident().GetText(), Kind: fieldCtx.FieldType().GetText(), Pos: fieldCtx.IntLit().GetText(), IsRepeated: fieldCtx.REPEATED() != nil}

		if fieldCtx.FieldOpts() != nil {
			fieldOptsCtx, ok := fieldCtx.FieldOpts().(*parser.FieldOptsContext)
			if !ok {
				log.Printf("expecting *FieldOptsContext got %T", fieldCtx.FieldOpts())
				continue
			}
			options := makeKeyValueOptions(fieldOptsCtx.AllKeyValue())
			field.Options = append(field.Options, options...)
		}

		result = append(result, field)
	}
	return result
}

func makeMapFields(ctx []parser.IMapFieldContext) []Field {
	var result []Field
	for _, mapDecl := range ctx {
		fieldCtx, ok := mapDecl.(*parser.MapFieldContext)
		if !ok {
			log.Print("expecting *MapFieldContext")
			continue
		}

		if fieldCtx.Ident() == nil {
			log.Print("map with no ident provided")
			continue
		}

		if fieldCtx.IntLit() == nil {
			log.Print("no pos provided")
			continue
		}

		if fieldCtx.FieldType() == nil {
			log.Printf("map has no value type")
			continue
		}

		if fieldCtx.KeyType() == nil {
			log.Printf("map has no key type")
			continue
		}

		mapInfo := KeyValue{Key: fieldCtx.KeyType().GetText(), Value: fieldCtx.FieldType().GetText()}
		field := Field{Name: fieldCtx.Ident().GetText(), Pos: fieldCtx.IntLit().GetText(), MapKind: mapInfo, IsMap: true}

		if fieldCtx.FieldOpts() != nil {
			fieldOptsCtx, ok := fieldCtx.FieldOpts().(*parser.FieldOptsContext)
			if !ok {
				log.Printf("expecting *FieldOptsContext got %T", fieldCtx.FieldOpts())
				continue
			}
			options := makeKeyValueOptions(fieldOptsCtx.AllKeyValue())
			field.Options = append(field.Options, options...)
		}

		result = append(result, field)
	}
	return result
}

func makeOneOfFields(ctx []parser.IOneofContext) (string, []Field) {
	var (
		result    []Field
		fieldName string
	)
	for _, ooDecl := range ctx {
		oneOfCtx, ok := ooDecl.(*parser.OneofContext)
		if !ok {
			log.Print("expecting *OneofContext")
			continue
		}
		if oneOfCtx.Ident() == nil {
			log.Print("one of with no ident")
			continue
		}
		fieldName = oneOfCtx.Ident().GetText()
		for _, s := range oneOfCtx.AllOneofField() {
			oofCtx, ok := s.(*parser.OneofFieldContext)
			if !ok {
				log.Print("expecting *OneofFieldContext")
			}
			if oofCtx.Ident() == nil {
				log.Print("one of with no ident provided")
				continue
			}

			if oofCtx.IntLit() == nil {
				log.Print("no pos provided")
				continue
			}

			if oofCtx.FieldType() == nil {
				log.Printf("one of has no value type")
				continue
			}
			field := Field{Name: oofCtx.Ident().GetText(), Kind: oofCtx.FieldType().GetText(), Pos: oofCtx.IntLit().GetText()}
			if oofCtx.FieldOpts() != nil {
				fieldOptsCtx, ok := oofCtx.FieldOpts().(*parser.FieldOptsContext)
				if !ok {
					log.Printf("expecting *FieldOptsContext got %T", oofCtx.FieldOpts())
					continue
				}
				options := makeKeyValueOptions(fieldOptsCtx.AllKeyValue())
				field.Options = append(field.Options, options...)
			}
			result = append(result, field)
		}
	}
	return fieldName, result
}

func makeMessage(ctx *parser.MessageContext, proto *Proto) Message {
	var result Message
	getIdent(&result.Name, ctx.Ident())

	declCtx, ok := ctx.MessageDecl().(*parser.MessageDeclContext)
	if !ok {
		log.Print("expecting *MessageDeclContext")
		return result
	}
	for _, rDecl := range declCtx.AllReserved() {
		resCtx, ok := rDecl.(*parser.ReservedContext)
		if !ok {
			log.Print("expecting *ReservedContext")
			continue
		}
		if resCtx.Ranges() != nil {
			rangeCtx, ok := resCtx.Ranges().(*parser.RangesContext)
			if !ok {
				log.Print("expecting *RangesContext")
				continue
			}
			for _, rangeDef := range rangeCtx.AllRangeDef() {
				rCtx, ok := rangeDef.(*parser.RangeDefContext)
				if !ok {
					log.Print("expecting *RangeDefContext")
					continue
				}
				for _, intlit := range rCtx.AllIntLit() {
					result.Fields = append(result.Fields, Field{IsReserved: true, Pos: intlit.GetText()})
				}
			}
		}

		if resCtx.FieldNames() != nil {
			fnCtx, ok := resCtx.FieldNames().(*parser.FieldNamesContext)
			if !ok {
				log.Print("expecting *FieldNamesContext")
				continue
			}
			for _, n := range fnCtx.AllStrLit() {
				result.Fields = append(result.Fields, Field{IsReserved: true, Name: n.GetText()})
			}
		}
	}
	msgOpts := makeOptions(declCtx.AllOption())
	result.Options = append(result.Options, msgOpts...)

	msgEnums := makeEnums(declCtx.AllEnum())
	result.Enums = append(result.Enums, msgEnums...)
	proto.Enums = append(proto.Enums, msgEnums...)

	msgFields := makeFields(declCtx.AllField())
	result.Fields = append(result.Fields, msgFields...)

	for _, nestedMsg := range declCtx.AllMessage() {
		nestMsgCtx, ok := nestedMsg.(*parser.MessageContext)
		if !ok {
			log.Print("expecting *MessageContext")
			continue
		}
		nestedMsg := makeMessage(nestMsgCtx, proto)
		nestedMsg.IsNested = true
		result.Messages = append(result.Messages, nestedMsg)
		proto.Messages = append(proto.Messages, nestedMsg)
	}

	mapFields := makeMapFields(declCtx.AllMapField())
	result.Fields = append(result.Fields, mapFields...)
	if declCtx.AllOneof() != nil {
		fieldName, oneOfFields := makeOneOfFields(declCtx.AllOneof())
		if len(oneOfFields) > 0 {
			result.Fields = append(result.Fields, Field{Name: fieldName, IsOneOf: true, OneOf: oneOfFields})
		}
	}
	return result
}

func makeNameAndKinds(idents []antlr.TerminalNode) []Field {
	var result []Field
	for _, id := range idents {
		result = append(result, Field{Kind: id.GetText()})
	}
	return result
}

func makeMethods(ctx []parser.IRpcContext) []Method {
	var result []Method
	for _, rpcDecl := range ctx {
		rpcCtx, ok := rpcDecl.(*parser.RpcContext)
		if !ok {
			log.Print("expecting *RpcContext")
			continue
		}
		if rpcCtx.RpcName() == nil {
			log.Print("rpcCtx.RpcName()")
			continue
		}
		serverStream := false
		if rpcCtx.ServerStreams() != nil {
			s, ok := rpcCtx.ServerStreams().(*parser.ServerStreamsContext)
			if ok {
				serverStream = s.STREAM() != nil
			} else {
				log.Print("expecting *ServerStreamsContext")
			}
		}
		clientStream := false
		if rpcCtx.ClientStreams() != nil {
			s, ok := rpcCtx.ClientStreams().(*parser.ClientStreamsContext)
			if ok {
				clientStream = s.STREAM() != nil
			} else {
				log.Print("expecting *ClientStreamsContext")
			}
		}
		rpc := Method{Name: rpcCtx.RpcName().GetText(), IsServerStreaming: serverStream, IsClientStreaming: clientStream}

		rpcOpts := makeOptions(rpcCtx.AllOption())
		rpc.Options = append(rpc.Options, rpcOpts...)

		rpcParams, ok := rpcCtx.RpcParams().(*parser.RpcParamsContext)
		if !ok {
			log.Print("expecting *RpcParamsContext")
			continue
		}
		rpc.In = append(rpc.In, makeNameAndKinds(rpcParams.AllIdent())...)

		rpcRets, ok := rpcCtx.RpcReturns().(*parser.RpcReturnsContext)
		if !ok {
			log.Print("expecting *RpcReturnsContext")
			continue
		}
		rpc.Out = append(rpc.Out, makeNameAndKinds(rpcRets.AllIdent())...)

		result = append(result, rpc)
	}
	return result
}

func (s *ProtoListener) ExitProto(ctx *parser.ProtoContext) {
	protoOpts := makeOptions(ctx.AllOption())
	s.Proto.Options = append(s.Proto.Options, protoOpts...)

	enums := makeEnums(ctx.AllEnum())
	s.Proto.Enums = append(s.Proto.Enums, enums...)

	for _, msgDecl := range ctx.AllMessage() {
		msgCtx, ok := msgDecl.(*parser.MessageContext)
		if !ok {
			log.Print("expecting *MessageContext")
			return
		}

		msg := makeMessage(msgCtx, &s.Proto)
		s.Proto.Messages = append(s.Proto.Messages, msg)
	}

	for _, impDecl := range ctx.AllImports() {
		impCtx, ok := impDecl.(*parser.ImportsContext)
		if !ok {
			log.Print("expecting *ImportsContext")
			return
		}

		importStr := ""
		getIdent(&importStr, impCtx.StrLit())
		unquotedImport, err := strconv.Unquote(importStr)
		if err != nil {
			log.Printf("error unquoting : %v", err)
			continue
		}
		s.Proto.Imports = append(s.Proto.Imports, unquotedImport)
	}

	for _, pkgDecl := range ctx.AllPackageName() {
		pkgCtx, ok := pkgDecl.(*parser.PackageNameContext)
		if !ok {
			log.Print("expecting *PackageNameContext")
			return
		}
		fiCtx, ok := pkgCtx.FullIdent().(*parser.FullIdentContext)
		if !ok {
			log.Print("expecting *FullIdentContext")
			return
		}
		packageNames := ""
		for idx, idCtx := range fiCtx.AllIdent() {
			packageName := ""
			getIdent(&packageName, idCtx)
			if idx > 0 {
				packageNames += "."
			}
			packageNames += packageName
		}
		s.Proto.PackageName = packageNames
	}

	for _, svcDecl := range ctx.AllService() {
		svcCtx, ok := svcDecl.(*parser.ServiceContext)
		if !ok {
			log.Print("expecting *ServiceContext")
			return
		}
		svc := Service{}
		getIdent(&svc.Name, svcCtx.Ident())
		svcOpts := makeOptions(svcCtx.AllOption())
		svc.Options = append(svc.Options, svcOpts...)
		svcMethods := makeMethods(svcCtx.AllRpc())
		svc.Methods = append(svc.Methods, svcMethods...)
		s.Proto.Services = append(s.Proto.Services, svc)
	}
}
