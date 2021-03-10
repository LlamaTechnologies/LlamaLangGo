// Code generated from ../../grammar/LlamaLang.g4 by ANTLR 4.9.1. DO NOT EDIT.

package antlr // LlamaLang
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by LlamaLangParser.
type LlamaLangVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by LlamaLangParser#sourceFile.
	VisitSourceFile(ctx *SourceFileContext) interface{}

	// Visit a parse tree produced by LlamaLangParser#identifierList.
	VisitIdentifierList(ctx *IdentifierListContext) interface{}

	// Visit a parse tree produced by LlamaLangParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by LlamaLangParser#functionDef.
	VisitFunctionDef(ctx *FunctionDefContext) interface{}

	// Visit a parse tree produced by LlamaLangParser#varDef.
	VisitVarDef(ctx *VarDefContext) interface{}

	// Visit a parse tree produced by LlamaLangParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by LlamaLangParser#statementList.
	VisitStatementList(ctx *StatementListContext) interface{}

	// Visit a parse tree produced by LlamaLangParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by LlamaLangParser#simpleStmt.
	VisitSimpleStmt(ctx *SimpleStmtContext) interface{}

	// Visit a parse tree produced by LlamaLangParser#expressionStmt.
	VisitExpressionStmt(ctx *ExpressionStmtContext) interface{}

	// Visit a parse tree produced by LlamaLangParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by LlamaLangParser#assign_op.
	VisitAssign_op(ctx *Assign_opContext) interface{}

	// Visit a parse tree produced by LlamaLangParser#emptyStmt.
	VisitEmptyStmt(ctx *EmptyStmtContext) interface{}

	// Visit a parse tree produced by LlamaLangParser#returnStmt.
	VisitReturnStmt(ctx *ReturnStmtContext) interface{}

	// Visit a parse tree produced by LlamaLangParser#type_.
	VisitType_(ctx *Type_Context) interface{}

	// Visit a parse tree produced by LlamaLangParser#pointer.
	VisitPointer(ctx *PointerContext) interface{}

	// Visit a parse tree produced by LlamaLangParser#array.
	VisitArray(ctx *ArrayContext) interface{}

	// Visit a parse tree produced by LlamaLangParser#typeName.
	VisitTypeName(ctx *TypeNameContext) interface{}

	// Visit a parse tree produced by LlamaLangParser#signature.
	VisitSignature(ctx *SignatureContext) interface{}

	// Visit a parse tree produced by LlamaLangParser#result.
	VisitResult(ctx *ResultContext) interface{}

	// Visit a parse tree produced by LlamaLangParser#parameters.
	VisitParameters(ctx *ParametersContext) interface{}

	// Visit a parse tree produced by LlamaLangParser#parameterDecl.
	VisitParameterDecl(ctx *ParameterDeclContext) interface{}

	// Visit a parse tree produced by LlamaLangParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by LlamaLangParser#primaryExpr.
	VisitPrimaryExpr(ctx *PrimaryExprContext) interface{}

	// Visit a parse tree produced by LlamaLangParser#unaryExpr.
	VisitUnaryExpr(ctx *UnaryExprContext) interface{}

	// Visit a parse tree produced by LlamaLangParser#conversion.
	VisitConversion(ctx *ConversionContext) interface{}

	// Visit a parse tree produced by LlamaLangParser#operand.
	VisitOperand(ctx *OperandContext) interface{}

	// Visit a parse tree produced by LlamaLangParser#unaryOp.
	VisitUnaryOp(ctx *UnaryOpContext) interface{}

	// Visit a parse tree produced by LlamaLangParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by LlamaLangParser#basicLit.
	VisitBasicLit(ctx *BasicLitContext) interface{}

	// Visit a parse tree produced by LlamaLangParser#integer.
	VisitInteger(ctx *IntegerContext) interface{}

	// Visit a parse tree produced by LlamaLangParser#floatingPoint.
	VisitFloatingPoint(ctx *FloatingPointContext) interface{}

	// Visit a parse tree produced by LlamaLangParser#operandName.
	VisitOperandName(ctx *OperandNameContext) interface{}

	// Visit a parse tree produced by LlamaLangParser#qualifiedIdent.
	VisitQualifiedIdent(ctx *QualifiedIdentContext) interface{}

	// Visit a parse tree produced by LlamaLangParser#literalType.
	VisitLiteralType(ctx *LiteralTypeContext) interface{}

	// Visit a parse tree produced by LlamaLangParser#fieldDecl.
	VisitFieldDecl(ctx *FieldDeclContext) interface{}

	// Visit a parse tree produced by LlamaLangParser#string_.
	VisitString_(ctx *String_Context) interface{}

	// Visit a parse tree produced by LlamaLangParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by LlamaLangParser#methodExpr.
	VisitMethodExpr(ctx *MethodExprContext) interface{}

	// Visit a parse tree produced by LlamaLangParser#receiverType.
	VisitReceiverType(ctx *ReceiverTypeContext) interface{}

	// Visit a parse tree produced by LlamaLangParser#eos.
	VisitEos(ctx *EosContext) interface{}
}
