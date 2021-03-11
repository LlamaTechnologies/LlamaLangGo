package main

import (
	"LlamaLangCompiler/cmd/shared"
	antlr4 "github.com/antlr/antlr4/runtime/Go/antlr"
)

type CompilerErrorListener struct {
	antlr4.ErrorListener
	Errors []shared.Error
}

func (listener *CompilerErrorListener) SyntaxError(_ antlr4.Recognizer, _ interface{}, line, column int, msg string, e antlr4.RecognitionException) {
	err := shared.Error{
		Line: line,
		Column: column,
		FileName: "file_name.llang",
		Message: msg,
		ErrorType: "ERROR",
	}
	listener.Errors = append(listener.Errors, err)
}

func (listener *CompilerErrorListener) ReportAmbiguity(recognizer antlr4.Parser, dfa *antlr4.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr4.BitSet, configs antlr4.ATNConfigSet) {

}

func (listener *CompilerErrorListener) ReportAttemptingFullContext(recognizer antlr4.Parser, dfa *antlr4.DFA, startIndex, stopIndex int, conflictingAlts *antlr4.BitSet, configs antlr4.ATNConfigSet) {

}

func (listener *CompilerErrorListener) ReportContextSensitivity(recognizer antlr4.Parser, dfa *antlr4.DFA, startIndex, stopIndex, prediction int, configs antlr4.ATNConfigSet) {

}