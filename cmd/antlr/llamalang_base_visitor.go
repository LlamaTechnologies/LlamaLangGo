// Code generated from ../../grammar/LlamaLang.g4 by ANTLR 4.9.1. DO NOT EDIT.

package antlr // LlamaLang
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseLlamaLangVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseLlamaLangVisitor) VisitSourceFile(ctx *SourceFileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitIdentifierList(ctx *IdentifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitFunctionDef(ctx *FunctionDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitVarDef(ctx *VarDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitStatementList(ctx *StatementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitSimpleStmt(ctx *SimpleStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitExpressionStmt(ctx *ExpressionStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitAssign_op(ctx *Assign_opContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitEmptyStmt(ctx *EmptyStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitReturnStmt(ctx *ReturnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitType_(ctx *Type_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitPointer(ctx *PointerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitArray(ctx *ArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitTypeName(ctx *TypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitSignature(ctx *SignatureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitResult(ctx *ResultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitParameters(ctx *ParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitParameterDecl(ctx *ParameterDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitPrimaryExpr(ctx *PrimaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitUnaryExpr(ctx *UnaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitConversion(ctx *ConversionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitOperand(ctx *OperandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitUnaryOp(ctx *UnaryOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitBasicLit(ctx *BasicLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitInteger(ctx *IntegerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitFloatingPoint(ctx *FloatingPointContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitOperandName(ctx *OperandNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitQualifiedIdent(ctx *QualifiedIdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitLiteralType(ctx *LiteralTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitFieldDecl(ctx *FieldDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitString_(ctx *String_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitMethodExpr(ctx *MethodExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitReceiverType(ctx *ReceiverTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLlamaLangVisitor) VisitEos(ctx *EosContext) interface{} {
	return v.VisitChildren(ctx)
}
