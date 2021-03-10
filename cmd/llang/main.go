package main

import (
	"LlamaLangCompiler/cmd/shared"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const srcFileURL = "../../tests/test.llang"

/**
* Main entry for the compiler
 */
func main() {
	fmt.Println("Starting...")

	config := shared.Config{}

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

	// print source file
	if config.IsVerbose {
		fmt.Println("Source file:")
		srcLines := strings.Split(string(srcFile), "\n")
		for _, line := range srcLines {
			fmt.Println(">> " + line)
		}
	}

	// lexer
	// parser
	// build ast
	// check semantics
	// print errors
	// exit if errors
	// ast -> IR
	// call clang and compile IR

	fmt.Println("Finished compiling!")
}
