package ast

import (
	"LlamaLangCompiler/cmd/antlr"
	"LlamaLangCompiler/cmd/shared"
	antlr4 "github.com/antlr/antlr4/runtime/Go/antlr"
)

// AstBuilder builds the ast by using visitor pattern in the parsing tree
type AstBuilder struct {
	antlr.BaseLlamaLangVisitor
	Errors  []shared.Error
	Program *Program
}

func (builder *AstBuilder) Visit(tree antlr4.ParseTree) interface{} {
	return tree.Accept(builder)
}

func (builder *AstBuilder) VisitChildren(node antlr4.RuleNode) interface{} {
	var result interface{}
	for _, childTree := range node.GetChildren() {
		child := childTree.(antlr4.RuleNode)
		if !shouldVisitNextChild(child) {
			continue
		}
		childResult := child.Accept(builder)
		aggregateResult(result, childResult)
	}

	return result
}

func (builder *AstBuilder) VisitErrorNode(errorNode antlr4.ErrorNode) interface{}   {
	err := shared.CreateError(errorNode, builder.Program.FileName)
	builder.Errors = append(builder.Errors, err)
	return nil
}

func (builder *AstBuilder) VisitSourceFile(ctx *antlr.SourceFileContext) interface{} {
	return builder.Program
}

// PRIVATE FUNCTIONS

func aggregateResult(result interface{}, nextResult interface{}) interface{} {
	if result != nil && nextResult == nil {
		return result
	}
	return nextResult
}

func shouldVisitNextChild(node antlr4.ParseTree) bool {
	_, ok := interface{}(node).(antlr4.TerminalNode)
	if !ok {
		return false
	}

	_, ok = interface{}(node).(antlr4.ErrorNode)
	if !ok {
		return false
	}

	ctx, ok := interface{}(node).(antlr4.RuleNode)
	if ok && ctx.GetChildCount() > 0{
		return true
	}

	return false
}