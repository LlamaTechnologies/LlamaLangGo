package ast

import (
	"LlamaLangCompiler/cmd/shared"
)

// ASTbuilder builds the ast by using visitor pattern in the parsing tree
type ASTbuilder struct {
	errors  []shared.Error
	program *Program
}
