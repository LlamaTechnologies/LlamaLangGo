package ast

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
	BINEXPR_ADD BinaryExprEnum = iota
	BINEXPR_SUB                = iota
	BINEXPR_MUL                = iota
	BINEXPR_DIV                = iota
	BINEXPR_MOD                = iota
)

var binaryExprEnumNames = map[BinaryExprEnum]string{
	BINEXPR_ADD: "ADD",
	BINEXPR_SUB: "SUB",
	BINEXPR_MUL: "MUL",
	BINEXPR_DIV: "DIV",
	BINEXPR_MOD: "MOD",
}

// Expression is the base node to represent expressions
type Expression struct {
	BaseNode
}

func (node *Expression) ToString(_ int) string {
	return ""
}

// UnaryExpression are those that have only one operand
type UnaryExpression struct {
	Expression
	ExprID UnaryExprEnum
	Value  Node
}

// BinaryExpression are those that have two operands
type BinaryExpression struct {
	Expression
	ExprID BinaryExprEnum
	Right  Node
	Left   Node
}

func (node *UnaryExpression) ToString(tabLevel int) string {
	tabs := GetTabs(tabLevel) + GetLineBeginChar() + " "
	exprName := unaryExprEnumNames[node.ExprID]
	str := tabs + exprName + " " + node.Value.ToString(0)
	return str
}

func (node *BinaryExpression) toString(tabLevel int) string {
	tabs := GetTabs(tabLevel)
	exprName := binaryExprEnumNames[node.ExprID]
	str := tabs + exprName + " " + node.Left.ToString(0) + ", " + node.Right.ToString(0)
	return str
}
