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
	Options     Options
	Imports     []string
	Messages    []Message
	Services    Services
	Enums       []Enum
	Error       *ErrNode
}

type Service struct {
	Name    string
	Methods []Method
	Options Options
}

type Services []Service

func (s Services) AnyOptionNamedValueEq(name, value string) bool {
	for _, svc := range s {
		for _, meth := range svc.Methods {
			if meth.Options.NamedValueEq(name, value) {
				return true
			}
		}
	}
	return false
}

func (s Services) HasOptionNamed(name string) bool {
	for _, svc := range s {
		for _, meth := range svc.Methods {
			if meth.Options.HasOneNamed(name) {
				return true
			}
		}
	}
	return false
}

type Method struct {
	Name              string
	InKind            string
	OutKind           string
	Options           Options
	IsServerStreaming bool
	IsClientStreaming bool
}

func (m *Method) NoStreaming() bool {
	return !m.IsServerStreaming && !m.IsClientStreaming
}

func (m *Method) GetHTTPVerb() string {
	for _, opt := range m.Options {
		if opt.Name == "google.api.http" {
			for _, val := range opt.Values {
				switch val.Key {
				case "get":
					return "http.MethodGet"
				case "put":
					return "http.MethodPut"
				case "delete":
					return "http.MethodDelete"
				case "post":
					return "http.MethodPost"
				case "patch":
					return "http.MethodPatch"
				}
			}
		}
	}
	return "UNSPECIFIED"
}

func (m *Method) GetHTTPRoute() string {
	for _, opt := range m.Options {
		if opt.Name == "google.api.http" {
			for _, val := range opt.Values {
				switch val.Key {
				case "get", "put", "delete", "post", "patch":
					return val.Value
				}
			}
		}
	}
	return "UNSPECIFIED"
}

type Message struct {
	Name     string
	Fields   Fields
	Options  Options
	Enums    []Enum
	Messages []Message // nested messages
	IsNested bool
}

type Enum struct {
	Name       string
	ParentName string
	Options    Options
	Fields     Fields
	IsEmbedded bool
}

func (e *Enum) PublicName() string {
	return e.ParentName + strings.ToUpper(e.Name[:1]) + e.Name[1:]
}

type Field struct {
	Name        string
	Kind        string
	MapKind     KeyValue
	Pos         string
	Options     Options
	OneOf       Fields
	IsEnumField bool // belongs to enum fields
	IsRepeated  bool
	IsMap       bool
	IsOneOf     bool
	IsReserved  bool
}

type Fields []Field

func (coll Fields) FirstPublicName() string {
	if len(coll) >= 2 {
		return coll[0].PublicName()
	}
	return ""
}

func (coll Fields) L() error {
	log.Printf("%d fields", len(coll))
	return nil
}

func (f *Field) IsBasic() bool {
	switch f.Kind {
	case "bool", "bytes", "double", "fixed32", "fixed64", "float", "int32", "int64", "sfixed32", "sfixed64", "sint32", "sint64", "string", "uint32", "uint64", "google.protobuf.Timestamp":
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
	case "google.protobuf.Timestamp":
		return "timestamp.Timestamp"
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
	case "google.protobuf.Timestamp":
		return "timestamp.Timestamp"
	}
	return f.Kind
}

func (f *Field) ZeroValue() string {
	switch f.Kind {
	case "bool":
		return "false"
	case "bytes":
		return "[]byte{}"
	case "double":
		return "0.0"
	case "fixed32":
		return "0"
	case "fixed64":
		return "0"
	case "float":
		return "0.0"
	case "int32":
		return "0"
	case "int64":
		return "0"
	case "map":
		return "make(map[" + getKind(f.MapKind.Key) + "]" + getKind(f.MapKind.Value) + ")"
	case "sfixed32":
		return "0"
	case "sfixed64":
		return "0"
	case "sint32":
		return "0"
	case "sint64":
		return "0"
	case "string":
		return "\"\""
	case "uint32":
		return "0"
	case "uint64":
		return "0"
	case "google.protobuf.Timestamp":
		return "&timestamp.Timestamp{}"
	}
	return "nil"
}

func (f *Field) PublicName() string {
	if len(f.Name) == 0 {
		return ""
	}
	name := ""
	if strings.Contains(f.Name, "_") {
		// clean enum _
		nameParts := strings.Split(f.Name, "_")
		for _, part := range nameParts {
			name += strings.ToUpper(part[:1]) + strings.ToLower(part[1:])
		}
		return name
	}
	return strings.ToUpper(f.Name[:1]) + f.Name[1:]
}

func (f *Field) IsTimestamp() bool {
	if f.Kind == "google.protobuf.Timestamp" {
		return true
	}
	return false
}

type Option struct {
	Name       string
	Value      string
	Values     []KeyValue
	IsConstant bool
}

type Options []Option

func (o Options) HasOneNamed(name string) bool {
	for _, opt := range o {
		if opt.Name == name {
			return true
		}
	}
	return false
}

func (o Options) NamedValueEq(name, value string) bool {
	for _, opt := range o {
		if opt.Name == name && opt.Value == value {
			return true
		}
	}
	return false
}

func (o Options) NamedValue(name string) string {
	for _, opt := range o {
		if opt.Name == name {
			return opt.Value
		}
	}
	return ""
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
		strLit := ct.StrLit().GetText()
		result, err := strconv.Unquote(strLit)
		if err != nil {
			log.Printf("error unquoting %q : %v", strLit, err)
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
	willPrintIdents := false
	if len(nameCtx.AllIdent()) > 1 {
		willPrintIdents = true
	}
	for _, ident := range nameCtx.AllIdent() {
		if willPrintIdents {
			log.Printf("%d option idents found : %q", len(nameCtx.AllIdent()), ident.GetText())
		}
		result.Name += ident.GetText()
	}
	if nameCtx.FullIdent() != nil {
		result.Name = nameCtx.FullIdent().GetText()
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

func makeOptions(ctx []parser.IOptionContext) Options {
	var result Options
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

func makeKeyValueOptions(ctx []parser.IKeyValueContext) Options {
	var result Options
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

func makeEnumFields(ctx []parser.IEnumFieldContext) Fields {
	var result Fields
	for _, enuFieldDecl := range ctx {
		enuFieldCtx, ok := enuFieldDecl.(*parser.EnumFieldContext)
		if !ok {
			log.Print("expecting *EnumFieldContext")
			continue
		}

		field := Field{Name: enuFieldCtx.Ident().GetText(), Kind: enuFieldCtx.IntLit().GetText(), IsEnumField: true}

		options := makeKeyValueOptions(enuFieldCtx.AllKeyValue())
		field.Options = append(field.Options, options...)

		result = append(result, field)
	}
	return result
}

func makeEnums(ctx []parser.IEnumContext, parentName string) []Enum {
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
		enum.ParentName = parentName
		enum.IsEmbedded = parentName != ""

		result = append(result, enum)
	}
	return result
}

func makeFields(ctx []parser.IFieldContext, enumsMap map[string]string) Fields {
	var result Fields
	for _, afCtx := range ctx {
		fieldCtx, ok := afCtx.(*parser.FieldContext)
		if !ok {
			log.Print("expecting *FieldContext")
			continue
		}
		name := ""
		if fieldCtx.Ident() == nil {
			if fieldCtx.ReservedWord() == nil {
				log.Printf("no ident provided for field (and no reserved word)")
				continue
			}
			res, ok := fieldCtx.ReservedWord().(*parser.ReservedWordContext)
			if !ok {
				log.Printf("expecting *ReservedWordContext")
				continue
			}
			name = res.GetText()
		} else {
			name = fieldCtx.Ident().GetText()
		}

		if fieldCtx.IntLit() == nil {
			log.Print("no pos provided for field")
			continue
		}

		if fieldCtx.FieldType() == nil {
			log.Print("no field type provided")
			continue
		}

		kind := fieldCtx.FieldType().GetText()
		field := Field{Name: name, Pos: fieldCtx.IntLit().GetText(), IsRepeated: fieldCtx.REPEATED() != nil}

		// fix embed enums fields
		if alterName, has := enumsMap[kind]; has {
			field.Kind = alterName
		} else {
			field.Kind = kind
		}

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

func makeMapFields(ctx []parser.IMapFieldContext) Fields {
	var result Fields
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

func makeOneOfFields(ctx []parser.IOneofContext) (string, Fields) {
	var (
		result    Fields
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

	msgEnums := makeEnums(declCtx.AllEnum(), result.Name)
	enumsMap := make(map[string]string)
	for _, enum := range msgEnums {
		enumsMap[enum.Name] = enum.ParentName + enum.Name
	}
	result.Enums = append(result.Enums, msgEnums...)
	proto.Enums = append(proto.Enums, msgEnums...)

	msgFields := makeFields(declCtx.AllField(), enumsMap)

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

func makeNameAndKinds(idents []antlr.TerminalNode) Fields {
	var result Fields
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
		if len(rpcParams.AllIdent()) != 1 {
			log.Printf("error : found %d params in rpc", len(rpcParams.AllIdent()))
			continue
		}
		rpc.InKind = rpcParams.AllIdent()[0].GetText()

		rpcRets, ok := rpcCtx.RpcReturns().(*parser.RpcReturnsContext)
		if !ok {
			log.Print("expecting *RpcReturnsContext")
			continue
		}
		if len(rpcRets.AllIdent()) != 1 {
			log.Printf("error : found %d returns in rpc", len(rpcRets.AllIdent()))
			continue
		}
		rpc.OutKind = rpcRets.AllIdent()[0].GetText()

		result = append(result, rpc)
	}
	return result
}

func (s *ProtoListener) ExitProto(ctx *parser.ProtoContext) {
	protoOpts := makeOptions(ctx.AllOption())
	s.Proto.Options = append(s.Proto.Options, protoOpts...)

	enums := makeEnums(ctx.AllEnum(), "")
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
			s.Proto.Imports = append(s.Proto.Imports, importStr)
			log.Printf("error unquoting %q : %v", importStr, err)
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
