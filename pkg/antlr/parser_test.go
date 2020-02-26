package antlr

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/badu/stroo/halp"
)

func TestListFiles(t *testing.T) {
	wd, _ := os.Getwd()
	for _, each := range []string{
		"/testdata/struct.proto",
		"/testdata/test_messages_proto3.proto",
		"/testdata/type.proto",
		"/testdata/test.proto",
	} {
		file, err := os.Open(wd + each)
		if err != nil {
			t.Fatalf("%v : %s", err, wd+each)
		}
		b, err := ioutil.ReadAll(file)
		if err != nil {
			t.Fatalf(each, err)
		}
		is := antlr.NewInputStream(string(b))
		// Create the Lexer
		lexer := NewantlrLexer(is)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		// Create the Parser
		p := NewantlrParser(stream)

		// Finally parse the expression (by walking the tree)
		var listener ProtoListener
		antlr.ParseTreeWalkerDefault.Walk(&listener, p.Proto())

		t.Logf("Proto:\n%s", halp.SPrint(listener.Proto))
	}
}

func TestParser(t *testing.T) {
	type testCase struct {
		id     int
		input  string
		output Proto
	}
	var testCases = []testCase{
		{
			id: 1,
			input: `syntax = "proto3";
					package google.protobuf;`,
			output: Proto{
				PackageName: "google.protobuf",
			},
		},
		{
			id: 2,
			input: `syntax = "bad_syntax";
					package google.protobuf;`,
			output: Proto{
				PackageName: "google.protobuf",
				Error: ErrNode{
					Line:   1,
					Column: 9,
					Text:   `line 1:9 mismatched input '"bad_syntax"' expecting {'"proto3"', ''proto3''}`,
				},
			},
		},
		{
			id: 3,
			input: `syntax = "proto3";
					package google.protobuf;
					option java_multiple_files = true;
					option java_package = "com.google.protobuf";`,
			output: Proto{
				PackageName: "google.protobuf",
				Options: []Option{
					{
						Name:       "java_multiple_files",
						Value:      "true",
						IsConstant: true,
					},
					{
						Name:       "java_package",
						Value:      "com.google.protobuf",
						IsConstant: true,
					},
				},
			},
		},
		{
			id: 4,
			input: `syntax = "proto3";
					import`,
			output: Proto{
				Imports: []string{
					"",
				},
				Error: ErrNode{
					Line:   2,
					Column: 11,
					Text:   `line 2:11 mismatched input '<EOF>' expecting {'public', 'weak', StrLit}`,
				},
			},
		},
		{
			id: 5,
			input: `syntax = "proto3";
					import bad_import_name`,
			output: Proto{
				Imports: []string{
					"",
				},
				Error: ErrNode{
					Line:   2,
					Column: 12,
					Text:   "line 2:12 mismatched input 'bad_import_name' expecting {'public', 'weak', StrLit}",
				},
			},
		},
		{
			id: 6,
			input: `syntax = "proto3";
					package google.protobuf;
					message Struct {
  						map<string, Value> fields = 1;
					}`,
			output: Proto{
				PackageName: "google.protobuf",
				Messages: []Message{
					{
						Name: "Struct",
						Fields: []Field{
							{
								Name: "fields",
								MapKind: KeyValue{
									Key:   "string",
									Value: "Value",
								},
								Pos:   "1",
								IsMap: true,
							},
						},
					},
				},
			},
		},
		{
			id: 7,
			input: `syntax = "proto3";
					package google.protobuf;
					message Struct {`,
			output: Proto{
				PackageName: "google.protobuf",
				Messages: []Message{
					Message{
						Name: "Struct",
					},
				},
				Error: ErrNode{
					Line:   3,
					Column: 21,
					Text:   "line 3:21 mismatched input '<EOF>' expecting {'bool', 'bytes', 'double', 'enum', 'fixed32', 'fixed64', 'float', 'int32', 'int64', 'map', 'message', 'oneof', 'option', 'package', 'repeated', 'reserved', 'rpc', 'service', 'sfixed32', 'sfixed64', 'sint32', 'sint64', 'stream', 'string', 'syntax', 'uint32', 'uint64', 'weak', '}', ';', '.', Ident}",
				},
			},
		},
		{
			id: 8,
			input: `syntax = "proto3";
					package google.protobuf;
					message Struct {
  						map<string, Value> fields = 1 [packed = true, weird_prop = "something"];
					}`,
			output: Proto{
				PackageName: "google.protobuf",
				Messages: []Message{
					{
						Name: "Struct",
						Fields: []Field{
							{
								Name: "fields",
								MapKind: KeyValue{
									Key:   "string",
									Value: "Value",
								},
								Pos: "1",
								Options: []Option{
									Option{
										Name:  "packed",
										Value: "true",
									},
									Option{
										Name:  "weird_prop",
										Value: "something",
									},
								},
								IsMap: true,
							},
						},
					},
				},
			},
		},
		{
			id: 9,
			input: `syntax = "proto3";
					message TestMessage {
					  int32 optional_int32 = 1;
					  int64 optional_int64 = 2;
					  uint32 optional_uint32 = 3;
					  uint64 optional_uint64 = 4;
					  sint32 optional_sint32 = 5;
					  sint64 optional_sint64 = 6;
					  fixed32 optional_fixed32 = 7;
					  fixed64 optional_fixed64 = 8;
					  sfixed32 optional_sfixed32 = 9;
					  sfixed64 optional_sfixed64 = 10;
					  float optional_float = 11;
					  double optional_double = 12;
					  bool optional_bool = 13;
					  string optional_string = 14;
					  bytes optional_bytes = 15;
					}`,
			output: Proto{
				Messages: []Message{
					Message{
						Name: "TestMessage",
						Fields: []Field{
							Field{
								Name: "optional_int32",
								Kind: "int32",
								Pos:  "1",
							},
							Field{
								Name: "optional_int64",
								Kind: "int64",
								Pos:  "2",
							},
							Field{
								Name: "optional_uint32",
								Kind: "uint32",
								Pos:  "3",
							},
							Field{
								Name: "optional_uint64",
								Kind: "uint64",
								Pos:  "4",
							},
							Field{
								Name: "optional_sint32",
								Kind: "sint32",
								Pos:  "5",
							},
							Field{
								Name: "optional_sint64",
								Kind: "sint64",
								Pos:  "6",
							},
							Field{
								Name: "optional_fixed32",
								Kind: "fixed32",
								Pos:  "7",
							},
							Field{
								Name: "optional_fixed64",
								Kind: "fixed64",
								Pos:  "8",
							},
							Field{
								Name: "optional_sfixed32",
								Kind: "sfixed32",
								Pos:  "9",
							},
							Field{
								Name: "optional_sfixed64",
								Kind: "sfixed64",
								Pos:  "10",
							},
							Field{
								Name: "optional_float",
								Kind: "float",
								Pos:  "11",
							},
							Field{
								Name: "optional_double",
								Kind: "double",
								Pos:  "12",
							},
							Field{
								Name: "optional_bool",
								Kind: "bool",
								Pos:  "13",
							},
							Field{
								Name: "optional_string",
								Kind: "string",
								Pos:  "14",
							},
							Field{
								Name: "optional_bytes",
								Kind: "bytes",
								Pos:  "15",
							},
						},
					},
				},
			},
		},
		{
			id: 10,
			input: `syntax = "proto3";
					message TestMessage {
						repeated    int32 repeated_int32    = 31;
						repeated    int64 repeated_int64    = 32;
						repeated   uint32 repeated_uint32   = 33;
						repeated   uint64 repeated_uint64   = 34;
						repeated   sint32 repeated_sint32   = 35;
						repeated   sint64 repeated_sint64   = 36;
						repeated  fixed32 repeated_fixed32  = 37;
						repeated  fixed64 repeated_fixed64  = 38;
						repeated sfixed32 repeated_sfixed32 = 39;
						repeated sfixed64 repeated_sfixed64 = 40;
						repeated    float repeated_float    = 41;
						repeated   double repeated_double   = 42;
						repeated     bool repeated_bool     = 43;
						repeated   string repeated_string   = 44;
						repeated    bytes repeated_bytes    = 45;
					}`,
			output: Proto{
				Messages: []Message{
					Message{
						Name: "TestMessage",
						Fields: []Field{
							Field{
								Name:       "repeated_int32",
								Kind:       "int32",
								Pos:        "31",
								IsRepeated: true,
							},
							Field{
								Name:       "repeated_int64",
								Kind:       "int64",
								Pos:        "32",
								IsRepeated: true,
							},
							Field{
								Name:       "repeated_uint32",
								Kind:       "uint32",
								Pos:        "33",
								IsRepeated: true,
							},
							Field{
								Name:       "repeated_uint64",
								Kind:       "uint64",
								Pos:        "34",
								IsRepeated: true,
							},
							Field{
								Name:       "repeated_sint32",
								Kind:       "sint32",
								Pos:        "35",
								IsRepeated: true,
							},
							Field{
								Name:       "repeated_sint64",
								Kind:       "sint64",
								Pos:        "36",
								IsRepeated: true,
							},
							Field{
								Name:       "repeated_fixed32",
								Kind:       "fixed32",
								Pos:        "37",
								IsRepeated: true,
							},
							Field{
								Name:       "repeated_fixed64",
								Kind:       "fixed64",
								Pos:        "38",
								IsRepeated: true,
							},
							Field{
								Name:       "repeated_sfixed32",
								Kind:       "sfixed32",
								Pos:        "39",
								IsRepeated: true,
							},
							Field{
								Name:       "repeated_sfixed64",
								Kind:       "sfixed64",
								Pos:        "40",
								IsRepeated: true,
							},
							Field{
								Name:       "repeated_float",
								Kind:       "float",
								Pos:        "41",
								IsRepeated: true,
							},
							Field{
								Name:       "repeated_double",
								Kind:       "double",
								Pos:        "42",
								IsRepeated: true,
							},
							Field{
								Name:       "repeated_bool",
								Kind:       "bool",
								Pos:        "43",
								IsRepeated: true,
							},
							Field{
								Name:       "repeated_string",
								Kind:       "string",
								Pos:        "44",
								IsRepeated: true,
							},
							Field{
								Name:       "repeated_bytes",
								Kind:       "bytes",
								Pos:        "45",
								IsRepeated: true,
							},
						},
					},
				},
			},
		},
		{
			id: 11,
			input: `syntax = "proto3";
					package foo;
					message TestMessage {
						oneof my_oneof {
							int32  oneof_int32    = 51;
							int64  oneof_int64    = 52;
							uint32 oneof_uint32   = 53;
							uint64 oneof_uint64   = 54;
							uint32 oneof_sint32   = 55;
							uint64 oneof_sint64   = 56;
							uint32 oneof_fixed32  = 57;
							uint64 oneof_fixed64  = 58;
							uint32 oneof_sfixed32 = 59;
							uint64 oneof_sfixed64 = 60;
							double oneof_double   = 61;
							float  oneof_float    = 62;
							bool   oneof_bool     = 63 [required=true];
							string oneof_string   = 64;
							bytes  oneof_bytes    = 65;
							TestEnum oneof_enum   = 66;
							Sub    oneof_message  = 67;
						}
					}`,
			output: Proto{
				PackageName: "foo",
				Messages: []Message{
					Message{
						Name: "TestMessage",
						Fields: []Field{
							Field{
								Name: "my_oneof",
								OneOf: []Field{
									Field{
										Name: "oneof_int32",
										Kind: "int32",
										Pos:  "51",
									},
									Field{
										Name: "oneof_int64",
										Kind: "int64",
										Pos:  "52",
									},
									Field{
										Name: "oneof_uint32",
										Kind: "uint32",
										Pos:  "53",
									},
									Field{
										Name: "oneof_uint64",
										Kind: "uint64",
										Pos:  "54",
									},
									Field{
										Name: "oneof_sint32",
										Kind: "uint32",
										Pos:  "55",
									},
									Field{
										Name: "oneof_sint64",
										Kind: "uint64",
										Pos:  "56",
									},
									Field{
										Name: "oneof_fixed32",
										Kind: "uint32",
										Pos:  "57",
									},
									Field{
										Name: "oneof_fixed64",
										Kind: "uint64",
										Pos:  "58",
									},
									Field{
										Name: "oneof_sfixed32",
										Kind: "uint32",
										Pos:  "59",
									},
									Field{
										Name: "oneof_sfixed64",
										Kind: "uint64",
										Pos:  "60",
									},
									Field{
										Name: "oneof_double",
										Kind: "double",
										Pos:  "61",
									},
									Field{
										Name: "oneof_float",
										Kind: "float",
										Pos:  "62",
									},
									Field{
										Name: "oneof_bool",
										Kind: "bool",
										Pos:  "63",
										Options: []Option{
											Option{
												Name:  "required",
												Value: "true",
											},
										},
									},
									Field{
										Name: "oneof_string",
										Kind: "string",
										Pos:  "64",
									},
									Field{
										Name: "oneof_bytes",
										Kind: "bytes",
										Pos:  "65",
									},
									Field{
										Name: "oneof_enum",
										Kind: "TestEnum",
										Pos:  "66",
									},
									Field{
										Name: "oneof_message",
										Kind: "Sub",
										Pos:  "67",
									},
								},
								IsOneOf: true,
							},
						},
					},
				},
			},
		},
		{
			id: 12,
			input: `syntax = "proto3";
					package google.protobuf;
					message TestMessage {
						message Sub {
    						int32 a = 1;
    						repeated int32 b = 2;
						}
  					}`,
			output: Proto{
				PackageName: "google.protobuf",
				Messages: []Message{
					Message{
						Name: "TestMessage",
						Messages: []Message{
							Message{
								Name: "Sub",
								Fields: []Field{
									Field{
										Name: "a",
										Kind: "int32",
										Pos:  "1",
									},
									Field{
										Name:       "b",
										Kind:       "int32",
										Pos:        "2",
										IsRepeated: true,
									},
								},
								IsNested: true,
							},
						},
					},
				},
			},
		},
		{
			id: 13,
			input: `syntax = "proto3";
					package google.protobuf;
					message TestMessage {
						enum TestEnum {
							option allow_alias = true;
							ZERO = 0;
							ONE  = 1;
							TWO  = 2 [ (custom_option) = "hello world" ];
							ECHO = 3;  // Test reserved name.
						}
  					}`,
			output: Proto{
				PackageName: "google.protobuf",
				Messages: []Message{
					Message{
						Name: "TestMessage",
						Enums: []Enum{
							Enum{
								Name: "TestEnum",
								Options: []Option{
									Option{
										Name:       "allow_alias",
										Value:      "true",
										IsConstant: true,
									},
								},
								Fields: []Field{
									Field{
										Name: "ZERO",
										Kind: "0",
									},
									Field{
										Name: "ONE",
										Kind: "1",
									},
									Field{
										Name: "TWO",
										Kind: "2",
										Options: []Option{
											Option{
												Name:  "(custom_option)",
												Value: "hello world",
											},
										},
									},
									Field{
										Name: "ECHO",
										Kind: "3",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			id: 14,
			input: `syntax = "proto3";
					package google.protobuf;
					message TestMessage {
						option (my_option).a = true;
						reserved "Foo";
						reserved 111,111,111;
						reserved 120 to 333;
						map<string, google.protobuf.Any> map_string_any = 122;
  					}`,
			output: Proto{
				PackageName: "google.protobuf",
				Messages: []Message{
					Message{
						Name: "TestMessage",
						Fields: []Field{
							Field{
								Name:       "\"Foo\"",
								IsReserved: true,
							},
							Field{
								Pos:        "111",
								IsReserved: true,
							},
							Field{
								Pos:        "111",
								IsReserved: true,
							},
							Field{
								Pos:        "111",
								IsReserved: true,
							},
							Field{
								Pos:        "120",
								IsReserved: true,
							},
							Field{
								Pos:        "333",
								IsReserved: true,
							},
							Field{
								Name: "map_string_any",
								MapKind: KeyValue{
									Key:   "string",
									Value: "google.protobuf.Any",
								},
								Pos:   "122",
								IsMap: true,
							},
						},
						Options: []Option{
							Option{
								Name:       "a",
								Value:      "true",
								IsConstant: true,
							},
						},
					},
				},
			},
		},
		{
			id: 15,
			input: `syntax = "proto3";
					package google.protobuf;
					enum TestEnum {
					  ZERO = 0 [ (custom_option) = "hello world" ];
					  ONE  = 1;
					  TWO  = 2;
					  ECHO = 3;  // Test reserved name.
					}`,
			output: Proto{
				PackageName: "google.protobuf",
				Enums: []Enum{
					Enum{
						Name: "TestEnum",
						Fields: []Field{
							Field{
								Name: "ZERO",
								Kind: "0",
								Options: []Option{
									Option{
										Name:  "(custom_option)",
										Value: "hello world",
									},
								},
							},
							Field{
								Name: "ONE",
								Kind: "1",
							},
							Field{
								Name: "TWO",
								Kind: "2",
							},
							Field{
								Name: "ECHO",
								Kind: "3",
							},
						},
					},
				},
			},
		},
		{
			id: 16,
			input: `syntax = "proto3";
					package google.protobuf;
					service myService {
						option (my_option).a = true;
						rpc Search (SearchRequest) returns (SearchResponse){};
						rpc SearchStream (stream SearchStreamRequest) returns (SearchResponse){};
						rpc SearchStream (SearchRequest) returns (stream SearchStreamResponse){};
						rpc SearchStream (stream SearchStreamRequest) returns (stream SearchStreamResponse){
							option (google.api.http) = {
            					post: "/errorclass"
            					body: "*"
        					};
						};
					}`,
			output: Proto{
				PackageName: "google.protobuf",
				Services: []Service{
					Service{
						Name: "myService",
						RPCs: []RPC{
							RPC{
								Name: "Search",
								In: []Field{
									Field{
										Kind: "SearchRequest",
									},
								},
								Out: []Field{
									Field{
										Kind: "SearchResponse",
									},
								},
							},
							RPC{
								Name: "SearchStream",
								In: []Field{
									Field{
										Kind: "SearchStreamRequest",
									},
								},
								Out: []Field{
									Field{
										Kind: "SearchResponse",
									},
								},
								IsClientStreaming: true,
							},
							RPC{
								Name: "SearchStream",
								In: []Field{
									Field{
										Kind: "SearchRequest",
									},
								},
								Out: []Field{
									Field{
										Kind: "SearchStreamResponse",
									},
								},
								IsServerStreaming: true,
							},
							RPC{
								Name: "SearchStream",
								In: []Field{
									Field{
										Kind: "SearchStreamRequest",
									},
								},
								Out: []Field{
									Field{
										Kind: "SearchStreamResponse",
									},
								},
								Options: []Option{
									Option{
										Values: []KeyValue{
											KeyValue{
												Key:   "post",
												Value: "/errorclass",
											},
											KeyValue{
												Key:   "body",
												Value: "*",
											},
										},
									},
								},
								IsServerStreaming: true,
								IsClientStreaming: true,
							},
						},
						Options: []Option{
							Option{
								Name:       "a",
								Value:      "true",
								IsConstant: true,
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Logf("running test id %d", tc.id)
		is := antlr.NewInputStream(tc.input)
		// Create the Lexer
		lexer := NewantlrLexer(is)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		// Create the Parser
		parser := NewantlrParser(stream)
		// remove ConsoleErrorListener
		parser.RemoveErrorListeners()

		var listener ProtoListener
		// add our own error listener
		parser.AddErrorListener(&ErrorListener{Listener: &listener})

		// Finally parse the expression (by walking the tree)
		antlr.ParseTreeWalkerDefault.Walk(&listener, parser.Proto())

		// check result
		if !reflect.DeepEqual(tc.output, listener.Proto) {
			cmp := halp.Equal(listener.Proto, tc.output)
			t.Logf("Compare : %#v", cmp)
			t.Fatalf("[%d] Expected:\n%s\nGot:\n%s", tc.id, halp.SPrint(tc.output), halp.SPrint(listener.Proto))

		}
	}
}
