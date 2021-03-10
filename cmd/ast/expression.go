package ast

import "strings"

// UnaryExprEnum enumType
type UnaryExprEnum int

// UNNARY_EXPR enum values
const (
	INCREMENT UnaryExprEnum = iota
	DECREMENT               = iota
	RETURN                  = iota
)

var unaryExprEnumNames = map[UnaryExprEnum]string{
	INCREMENT: "INC",
	DECREMENT: "DEC",
	RETURN:    "RET",
}

// BinaryExprEnum enumType
type BinaryExprEnum int

// BINARY_EXPR enum values
const (
	ADD BinaryExprEnum = iota
	SUB                = iota
	MUL                = iota
	DIV                = iota
	MOD                = iota
)

var binaryExprEnumNames = map[BinaryExprEnum]string{
	ADD: "ADD",
	SUB: "SUB",
	MUL: "MUL",
	DIV: "DIV",
	MOD: "MOD",
}

// Expression is the base node to represent expressions
type Expression struct {
	BaseNode
}

// UnaryExpression are those that have only one operand
type UnaryExpression struct {
	Expression
	ExprID UnaryExprEnum
	Value  Expression
}

// BinaryExpression are those that have two operands
type BinaryExpression struct {
	Expression
	ExprID BinaryExprEnum
	Right  Expression
	Left   Expression
}

func (node *UnaryExpression) toString(tabLevel int) string {
	tabs := strings.Repeat("\t", tabLevel)
	exprName := unaryExprEnumNames[node.ExprID]
	str := tabs + exprName + " " + node.Value.toString(0)
	return str
}

func (node *BinaryExpression) toString(tabLevel int) string {
	tabs := strings.Repeat("\t", tabLevel)
	exprName := binaryExprEnumNames[node.ExprID]
	str := tabs + exprName + " " + node.Left.toString(0) + ", " + node.Right.toString(0)
	return str
}
