package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/badu/gokit-gen/pkg/parser"
	"github.com/badu/gokit-gen/pkg/proto"
	_ "github.com/badu/gokit-gen/statik"
	"github.com/rakyll/statik/fs"
	"go/format"
	"golang.org/x/tools/imports"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"text/template"
)

func loadProtoFile(wd, src string) ([]byte, error) {
	var err error

	fileName := filepath.Base(src)
	if filepath.Ext(fileName) != ".proto" {
		return nil, fmt.Errorf("input must be a .proto file (bad extension %q)", filepath.Ext(fileName))
	}
	relativePathLookup := fileName == src

	fileDir := ""
	if !relativePathLookup {
		// not absolute path
		if fileDir, err = filepath.Abs(src); err != nil {
			// attempting relative
			if fileDir, err = filepath.Abs(wd + src); err != nil {
				log.Printf("folder not found while looking for file descriptor : %v", err)
				return nil, err
			}
		}
	} else {
		if fileDir, err = filepath.Abs(wd + src); err != nil {
			log.Printf("folder not found while looking for file descriptor : %v", err)
			return nil, err
		}
	}

	if _, err := os.Stat(fileDir); err != nil {
		log.Printf("error finding %q ", fileDir)
		return nil, err
	}

	var protoFile *os.File
	if protoFile, err = os.Open(fileDir); err != nil {
		log.Printf("error opening %q ", fileDir)
		return nil, err
	}

	var result []byte
	if result, err = ioutil.ReadAll(protoFile); err != nil {
		log.Printf("error reading %q ", fileDir)
		return nil, err
	}
	return result, nil
}

func parseProto(wd, src string) (*proto.Proto, error) {
	var (
		protoFileContent []byte
		err              error
	)
	// load the file (with checks)
	if protoFileContent, err = loadProtoFile(wd, src); err != nil {
		return nil, err
	}

	is := antlr.NewInputStream(string(protoFileContent))
	// Create the Lexer
	lexer := parser.NewantlrLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewantlrParser(stream)
	// remove ConsoleErrorListener
	p.RemoveErrorListeners()

	var listener proto.ProtoListener
	// add our own error listener
	p.AddErrorListener(&proto.ErrorListener{Listener: &listener})

	// Finally parse the expression (by walking the tree)
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Proto())

	if listener.Proto.Error != nil {
		return nil, fmt.Errorf("%s", listener.Proto.Error.Text)
	}
	// TODO : analyse imports and load what is possible
	return &listener.Proto, nil
}

func loadSharedProto(wd, src string, pro *proto.Proto) error {
	var (
		protoFileContent []byte
		err              error
	)
	// load the file (with checks)
	if protoFileContent, err = loadProtoFile(wd, src); err != nil {
		return err
	}

	is := antlr.NewInputStream(string(protoFileContent))
	// Create the Lexer
	lexer := parser.NewantlrLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewantlrParser(stream)
	// remove ConsoleErrorListener
	p.RemoveErrorListeners()

	var listener proto.ProtoListener
	// add our own error listener
	p.AddErrorListener(&proto.ErrorListener{Listener: &listener})

	// Finally parse the expression (by walking the tree)
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Proto())

	if listener.Proto.Error != nil {
		return fmt.Errorf("%s", listener.Proto.Error.Text)
	}
	pro.Enums = append(pro.Enums, listener.Proto.Enums...)
	pro.Messages = append(pro.Messages, listener.Proto.Messages...)
	for _, opt := range listener.Proto.Options {
		// skip Go package option
		if opt.Name == "go_package" {
			continue
		}
		pro.Options = append(pro.Options, opt)
	}

	pro.Services = append(pro.Services, listener.Proto.Services...)
	return nil
}

func templateFileNames(templateDir string) ([]string, error) {
	if stat, err := os.Stat(templateDir); err != nil || !stat.IsDir() {
		if err != nil {
			return nil, err
		}
		if !stat.IsDir() {
			return nil, fmt.Errorf("%q is not a folder", templateDir)
		}

	}
	var fileNames []string
	err := filepath.Walk(templateDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".tmpl" {
			return nil
		}
		rel, err := filepath.Rel(templateDir, path)
		if err != nil {
			return err
		}
		fileNames = append(fileNames, rel)
		return nil
	})
	if len(fileNames) < 1 {
		return nil, fmt.Errorf("no templates found in %q", templateDir)
	}
	return fileNames, err
}

func provideFile(statikFS http.FileSystem, path string) (string, error) {
	r, err := statikFS.Open(path)
	if err != nil {
		log.Printf("error finding file : %q", path)
		return "", err
	}
	defer r.Close()
	contents, err := ioutil.ReadAll(r)
	if err != nil {
		log.Printf("error reading file : %q", path)
		return "", err
	}
	return string(contents), nil
}

func buildContent(statikFS http.FileSystem, templateDir, templateFilename string, funcMap template.FuncMap, prt *proto.Proto) ([]byte, error) {
	//log.Printf("processing file %q", templateFilename)
	var (
		tmpl *template.Template
		err  error
	)
	// initialize template engine
	fullPath := filepath.Join(templateDir, templateFilename)
	templateName := filepath.Base(fullPath)
	if templateDir == "default" {
		content, err := provideFile(statikFS, templateFilename)
		if err != nil {
			return nil, err
		}
		tmpl, err = template.New(templateName).Funcs(funcMap).Parse(content)
	} else {
		// template init
		tmpl, err = template.New(templateName).Funcs(funcMap).ParseFiles(fullPath)
		if err != nil {
			return nil, err
		}
	}

	// generate the content
	buffer := new(bytes.Buffer)
	if err := tmpl.Execute(buffer, prt); err != nil {
		return nil, err
	}
	// optimize imports
	optImports, err := imports.Process("", buffer.Bytes(), nil)
	if err != nil {
		return buffer.Bytes(), err
	}

	// format file
	formatted, err := format.Source(optImports)
	if err != nil {
		return optImports, err

	}
	//log.Printf("%q ok", templateFilename)
	return formatted, nil
}

func run(args []string, stdout io.Writer) error {
	var err error
	if len(args) < 4 {
		log.Print("Usage")
		fmt.Fprint(stdout, "Usage: \n gokit-gen <protofile.proto> <templates_folder> <target_package_path> [optional <output_folder>] \n"+
			"<protofile> can be a relative path\n"+
			"<templates_folder> can be a relative path\n"+
			"<target_package_path> must be provided for test imports\n"+
			"<output_folder> must be absolute")

		return errors.New("not enough valid arguments received")
	}

	wd, _ := os.Getwd()
	wd += "/"
	log.Printf("working in folder folder : %q", wd)

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	templateDir := args[2]
	if templateDir != "default" {
		relativePathTemplatesLookup := filepath.Base(args[2]) == args[2]
		if !relativePathTemplatesLookup {
			templateDir, err = filepath.Abs(wd + args[2])
			if err != nil {
				log.Fatalf("folder not found while looking error reading templates dir : %v", err)
			}
		}
	}

	var templates []string
	if templateDir != "default" {
		if templates, err = templateFileNames(templateDir); err != nil {
			return err
		}
	} else {
		err = fs.Walk(statikFS, "/", func(path string, fi os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if path != "/" {
				templates = append(templates, path)
			}
			return nil
		})
	}

	fname := filepath.Base(args[1])
	fdir := filepath.Dir(args[1]) + "/"
	var result *proto.Proto
	if result, err = parseProto(wd, args[1]); err != nil {
		return err
	}
	for _, imp := range result.Imports {
		if filepath.Base(imp) == imp {
			log.Printf("import from the same folder : %q", imp)
			err := loadSharedProto(fdir, imp, result)
			if err != nil {
				return err
			}
		}
	}

	deployTo := ""
	debug := false
	if len(args) == 5 {
		deployTo = args[4] // target deployment provided
		if deployTo == "debug" {
			debug = true
		} else if !strings.HasSuffix(deployTo, "/") {
			deployTo += "/"
		}
	}

	var enumNamesMap = make(map[string]struct{})
	// make enum names map
	for _, enum := range result.Enums {
		enumNamesMap[enum.PublicName()] = struct{}{}
	}
	// check has streaming
	hasStreaming := false
	for _, svc := range result.Services {
		for _, rpc := range svc.Methods {
			if rpc.IsClientStreaming || rpc.IsServerStreaming {
				hasStreaming = true
				goto out
			}
		}
	}
out:
	goPackage := ""
	for _, opt := range result.Options {
		if opt.Name == "go_package" && opt.IsConstant {
			goPackage = opt.Value
			break
		}
	}
	if len(result.Services) > 1 {
		return fmt.Errorf("one service at a time please. found %d services", len(result.Services))
	}

	if goPackage == "" {
		return errors.New("undefined go package")
	}

	var funcMap = template.FuncMap{
		"string": func(i interface {
			String() string
		}) string {
			return i.String()
		},
		"json": func(v interface{}) string {
			a, err := json.Marshal(v)
			if err != nil {
				return err.Error()
			}
			return string(a)
		},
		"prettyjson": func(v interface{}) string {
			a, err := json.MarshalIndent(v, "", "  ")
			if err != nil {
				return err.Error()
			}
			return string(a)
		},
		"splitArray": func(sep string, s string) []interface{} {
			var r []interface{}
			t := strings.Split(s, sep)
			for i := range t {
				if t[i] != "" {
					r = append(r, t[i])
				}
			}
			return r
		},
		"first": func(a []string) string {
			return a[0]
		},
		"last": func(a []string) string {
			return a[len(a)-1]
		},
		"concat": func(a string, b ...string) string {
			return strings.Join(append([]string{a}, b...), "")
		},
		"join": func(sep string, a ...string) string {
			return strings.Join(a, sep)
		},
		"upperFirst": func(s string) string {
			return strings.ToUpper(s[:1]) + s[1:]
		},
		"lowerFirst": func(s string) string {
			return strings.ToLower(s[:1]) + s[1:]
		},
		"contains": func(sub, s string) bool {
			return strings.Contains(s, sub)
		},
		"trimstr": func(cutset, s string) string {
			return strings.Trim(s, cutset)
		},
		"index": func(array interface{}, i int) interface{} {
			slice := reflect.ValueOf(array)
			if slice.Kind() != reflect.Slice {
				panic("Error in index(): given a non-slice type")
			}
			if i < 0 || i >= slice.Len() {
				panic("Error in index(): index out of bounds")
			}
			return slice.Index(i).Interface()
		},

		"projectName": func() string {
			return strings.Replace(fname, filepath.Ext(fname), "", -1)
		},
		"goPkg": func() string {
			return args[3]
		},
		"deployTo": func() string {
			return deployTo
		},
		"isEnumField": func(field *proto.Field) bool {
			_, has := enumNamesMap[field.Kind]
			return has
		},
		"hasStreaming": func() bool {
			return hasStreaming
		},
		"go_package": func() string {
			return goPackage
		},
		"pb_package": func() string {
			return result.PackageName
		},
		"messageByKind": func(kind string) (*proto.Message, error) {
			for _, message := range result.Messages {
				if message.Name == kind {
					return &message, nil
				}
			}
			return nil, errors.New("message " + kind + " not found")
		},
		"usesHTTP": func() bool {
			for _, imp := range result.Imports {
				if imp == "google/api/annotations.proto" {
					return true
				}
			}
			return false
		},
	}

	for _, tpl := range templates {
		content, err := buildContent(statikFS, templateDir, tpl, funcMap, result)
		if err != nil {
			log.Printf("error processing %q : %v", tpl, err)
			if len(content) > 0 {
				log.Printf("partial result : %s", string(content))
			}
			//continue
		}
		if !debug {
			if len(content) == 0 {
				continue
			}
			tpl = strings.Replace(tpl, ".tmpl", "", -1)
			log.Printf("writing %q into %q", filepath.Base(tpl), deployTo)
			srcFile, err := os.Create(deployTo + filepath.Base(tpl))
			if err != nil {
				log.Printf("error creating file : %v", err)
				continue
			}
			if _, err := srcFile.Write(content); err != nil {
				log.Printf("error writing to file : %v", err)
				continue
			}
		} else {
			log.Printf("Generated file : %q\n%s", tpl, content)
		}
	}
	return nil
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("gokit-gen: ")

	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
