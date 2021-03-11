package ast

import (
	"LlamaLangCompiler/cmd/antlr"
	"LlamaLangCompiler/cmd/shared"
	antlr4 "github.com/antlr/antlr4/runtime/Go/antlr"
	"strings"
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

	// Get parameters
	parametersResult := builder.Visit(ctx.Signature())
	parameters, isParameters := parametersResult.([]*VariableDefinition)
	if isParameters {
		astNode.Parameters = parameters
	}

	// Get the result of the last child visited (block)
	childResult := builder.Visit(ctx.Block())
	block, isBlock := childResult.([]*BaseNode)
	if isBlock {
		astNode.Block = block
	}

	return astNode
}

func (builder *AstBuilder) VisitVarDef(ctx *antlr.VarDefContext) interface{} {
	// variable definition
	varDefNode := new(VariableDefinition)
	varDefNode.FileName = builder.Program.FileName
	varDefNode.LineNumber = ctx.GetStart().GetLine()
	varDefNode.Name = ctx.IDENTIFIER().GetText()
	varDefNode.VarType = ctx.Type_().GetText()

	if ctx.ASSIGN() != nil {
		expressionList := ctx.ExpressionList().(*antlr.ExpressionListContext)
		assignmentNode := assignmentNodeConstruct(builder, varDefNode.Name, expressionList)
		varDefNode.AssignmentStatement = assignmentNode
	}

	// Add symbol
	builder.currentScope.AddSymbol(varDefNode.Name, varDefNode.BaseNode)

	return varDefNode
}

func (builder *AstBuilder) VisitParameters(ctx *antlr.ParametersContext) interface{} {
	var parameters []*VariableDefinition
	for i := range ctx.GetChildren() {
		paramCtx := ctx.ParameterDecl(i).(*antlr.ParameterDeclContext)
		paramNode := new(VariableDefinition)
		paramNode.FileName = builder.Program.FileName
		paramNode.LineNumber = paramCtx.GetStart().GetLine()
		paramNode.Name = paramCtx.IDENTIFIER().GetText()
		paramNode.VarType = paramCtx.Type_().GetText()
		parameters = append(parameters, paramNode)
	}
	return parameters
}

func (builder *AstBuilder) VisitStatementList(ctx *antlr.StatementListContext) interface{} {
	var block []*BaseNode
	for i := range ctx.GetChildren() {
		stmntCtx := ctx.Statement(i).(*antlr.StatementContext)
		stmntResult := builder.Visit(stmntCtx)
		stmnt, isBaseNode := stmntResult.(*BaseNode)
		if isBaseNode {
			block = append(block, stmnt)
		}
	}
	return block
}

func (builder *AstBuilder) VisitAssignment(ctx *antlr.AssignmentContext) interface{} {
	expressionList := ctx.ExpressionList().(*antlr.ExpressionListContext)
	varName := ctx.IDENTIFIER().GetText()
	return assignmentNodeConstruct(builder, varName, expressionList)
}

func (builder *AstBuilder) VisitReturnStmt(ctx *antlr.ReturnStmtContext) interface{} {
	retStmnt := new(UnaryExpression)
	retStmnt.ExprID = RETURN
	retStmnt.FileName = builder.Program.FileName
	retStmnt.LineNumber = ctx.GetStart().GetLine()

	valueResult := builder.VisitChildren(ctx)

	if valueResult != nil {
		value := valueResult.(*Expression)
		retStmnt.Value = value
	}

	return retStmnt
}

func (builder *AstBuilder) VisitExpression(ctx *antlr.ExpressionContext) interface{} {
	return builder.VisitChildren(ctx)
}

func (builder *AstBuilder) VisitUnaryExpr(ctx *antlr.UnaryExprContext) interface{} {
	childResult := builder.VisitChildren(ctx)

	constantNode, isConstNode := childResult.(*LiteralConstant)
	unaryOpCtx := ctx.UnaryOp().(*antlr.UnaryOpContext)

	// Returns a negative or positive constant
	if isConstNode {
		if unaryOpCtx.MINUS() != nil {
			constantNode.Value = "-" + constantNode.Value
		}
		return constantNode
	}

	// TODO: Returns a DECREMENT or INCREMENT expression
	unaryExpr := new(UnaryExpression)

	if unaryOpCtx.MINUS_MINUS() != nil {
		unaryExpr.ExprID = DECREMENT
	} else {
		if unaryOpCtx.PLUS_PLUS() != nil {
			unaryExpr.ExprID = INCREMENT
		}
	}

	unaryExpr.Value = childResult.(*Expression)
	return unaryExpr
}

func (builder *AstBuilder) VisitBasicLit(ctx *antlr.BasicLitContext) interface{} {
	constantNode := new(LiteralConstant)
	constantNode.FileName = builder.Program.FileName
	constantNode.LineNumber = ctx.GetStart().GetLine()

	// Is a Integer number
	if ctx.Integer() != nil {
		text := ctx.Integer().GetText()
		constantNode.Value = text

		// is Hex
		if strings.ToLower(string(text[1])) == "x" {
			constantNode.ConstantID = CONST_HEX
			return constantNode
		}

		// is Oct
		if strings.ToLower(string(text[0])) == "0" {
			constantNode.ConstantID = CONST_OCT
			return constantNode
		}

		// Is Decimal
		constantNode.ConstantID = CONST_DEC
		return constantNode
	}

	// Is Floating Point number
	if ctx.FloatingPoint() != nil {
		text := ctx.FloatingPoint().GetText()
		textSize := len(text)
		if textSize < 1 {
			constantNode.ConstantID = CONST_DOUBLE
		}

		indexOfLast := textSize - 1
		fChar := string(text[indexOfLast])

		if strings.ToLower(fChar) == "f" {
			constantNode.ConstantID = CONST_FLOAT
			// pop the F from the end
			text = text[:indexOfLast]
		}

		constantNode.Value = text
		return constantNode
	}

	// Is a Character
	if ctx.RUNE_LIT() != nil {
		constantNode.ConstantID = CONST_CHAR
		constantNode.Value = ctx.String_().GetText()
		return constantNode
	}

	// Is a string
	if ctx.String_() != nil {
		constantNode.ConstantID = CONST_STRING
		constantNode.Value = ctx.String_().GetText()
		return constantNode
	}

	return nil
}

func (builder *AstBuilder) VisitOperandName(ctx *antlr.OperandNameContext) interface{} {
	varName := ctx.IDENTIFIER().GetText()

	varRefNode := new(VariableReference)
	varRefNode.VarName = varName

	return varRefNode
}

// PRIVATE FUNCTIONS

func assignmentNodeConstruct(builder *AstBuilder, name string, epxressionList *antlr.ExpressionListContext) *Assignment {
	assignmentNode := new(Assignment)
	assignmentNode.Variable = VariableReference{
		VarName: name,
	}
	exprResult := builder.Visit(epxressionList)
	exprNode, isExprNode := exprResult.(*Expression)
	if isExprNode {
		assignmentNode.Value = exprNode
	}

	return assignmentNode
}

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
