package ast

import (
	"LlamaLangCompiler/cmd/antlr"
	"LlamaLangCompiler/cmd/shared"
	antlr4 "github.com/antlr/antlr4/runtime/Go/antlr"
)

// AstBuilder builds the ast by using visitor pattern in the parsing tree
type AstBuilder struct {
	antlr.BaseLlamaLangVisitor
	Errors       []shared.Error
	Program      *Program
	GlobalScope  *Scope
	currentScope *Scope
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

func (builder *AstBuilder) VisitErrorNode(errorNode antlr4.ErrorNode) interface{} {
	err := shared.CreateError(errorNode, builder.Program.FileName)
	builder.Errors = append(builder.Errors, err)
	return nil
}

func (builder *AstBuilder) VisitSourceFile(ctx *antlr.SourceFileContext) interface{} {
	// set GlobalScope as the current scope
	builder.currentScope = builder.GlobalScope

	// Visit Children
	for _, childTree := range ctx.GetChildren() {
		child := childTree.(antlr4.RuleNode)
		result := builder.Visit(child)
		if result != nil {
			astNode := result.(*BaseNode)
			builder.Program.Children = append(builder.Program.Children, astNode)
		}
	}

	return builder.Program
}

func (builder *AstBuilder) VisitFunctionDef(ctx *antlr.FunctionDefContext) interface{} {
	astNode := new(FunctionDefinition)
	astNode.FileName = builder.Program.FileName
	astNode.LineNumber = ctx.GetStart().GetLine()
	astNode.Name = ctx.IDENTIFIER().GetText()
	astNode.RetType = ctx.Type_().GetText()

	// Add function scope to parent scope and make it current
	builder.currentScope = builder.currentScope.AddChildren(FUNCTION_SCOPE, astNode.Name, astNode)

	// Get the result of the last child visited (block)
	childResult := builder.VisitChildren(ctx)

	block, isBlock := childResult.([]*BaseNode)
	if isBlock {
		astNode.Block = block
	}

	return astNode
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
	if ok && ctx.GetChildCount() > 0 {
		return true
	}

	return false
}
