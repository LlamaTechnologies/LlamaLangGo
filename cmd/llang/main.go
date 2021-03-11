package main

import (
	"LlamaLangCompiler/cmd/antlr"
	"LlamaLangCompiler/cmd/ast"
	"LlamaLangCompiler/cmd/shared"
	"fmt"
	antlr4 "github.com/antlr/antlr4/runtime/Go/antlr"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var separator = strings.Repeat("=", 96)

/**
* Main entry for the compiler
 */
func main() {
	fmt.Println("Starting...")

	config := shared.Config{}
	srcFileURL, _ := filepath.Abs("./examples/test.llang")
	//srcFileURL, _ :=  filepath.Abs("./examples/test_syntax_errors.llang")
	//srcFileURL, _ :=  filepath.Abs("./examples/test_semantic_errors.llang")

	executableName := filepath.Base(srcFileURL)

	// get command line arguments
	args := os.Args[1:]

	fmt.Println("Args:")
	for _, arg := range args {
		if arg == "-v" {
			config.IsVerbose = true
		}
		fmt.Println(arg)
	}

	// open source code
	srcFile, err := ioutil.ReadFile(srcFileURL)

	if err != nil {
		panic(err)
	}

	strFile := string(srcFile)

	// print source file
	if config.IsVerbose {
		fmt.Println(separator)
		fmt.Println("Source file:")
		srcLines := strings.Split(strFile, "\n")
		for _, line := range srcLines {
			fmt.Println(">> " + line)
		}
	}

	// input stream
	srcInputStream := antlr4.NewInputStream(strFile)

	// lexer
	lexer := antlr.NewLlamaLangLexer(srcInputStream)

	// TokenStream
	tokenStream := antlr4.NewCommonTokenStream(lexer, antlr4.LexerDefaultTokenChannel)

	// parser
	errorListener := CompilerErrorListener{}
	parser := antlr.NewLlamaLangParser(tokenStream)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(&errorListener)
	tree := parser.SourceFile()

	// build ast
	astBuilder := ast.AstBuilder{
		Program:     ast.NewProgramNode(executableName),
		GlobalScope: ast.NewScope(ast.GLOBAL_SCOPE, nil, nil),
		Errors:      errorListener.Errors,
	}
	astree := astBuilder.Visit(tree).(*ast.Program)

	if config.IsVerbose {
		fmt.Println(separator)
		fmt.Println("Abstract Syntax Tree:")
		fmt.Println(astree.ToString(0))
	}

	// check semantics

	// print errors
	errCount := len(astBuilder.Errors)
	if errCount > 0 {
		fmt.Println(separator)
		fmt.Println(" Errors found (" + strconv.Itoa(errCount) + "):")
		for _, err := range astBuilder.Errors {
			errStr := fmt.Errorf("%v\n", err)
			fmt.Println(errStr.Error())
		}

		fmt.Println(separator)
		fmt.Println("Finished with errors")
		os.Exit(1)
	}

	// print errors
	// exit if errors
	// ast -> IR
	// call clang and compile IR

	fmt.Println(separator)
	fmt.Println("Finished compiling!")
}
