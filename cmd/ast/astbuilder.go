package ast

import (
	"LlamaLangCompiler/cmd/antlr"
	"LlamaLangCompiler/cmd/shared"
	antlr4 "github.com/antlr/antlr4/runtime/Go/antlr"
)

// AstBuilder builds the ast by using visitor pattern in the parsing tree
type AstBuilder struct {
	antlr.BaseLlamaLangVisitor
	errors  []shared.Error
	Program *Program
}

func (builder *AstBuilder) Visit(tree antlr4.ParseTree) interface{} {
	return tree.Accept(builder)
}

func (builder *AstBuilder) VisitChildren(node antlr4.RuleNode) interface{}     { return nil }
func (builder *AstBuilder) VisitTerminal(node antlr4.TerminalNode) interface{} { return nil }
func (builder *AstBuilder) VisitErrorNode(node antlr4.ErrorNode) interface{}   { return nil }

func (builder *AstBuilder) VisitSourceFile(ctx *antlr.SourceFileContext) interface{} {
	return builder.Program
}

