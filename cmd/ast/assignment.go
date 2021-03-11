package ast

import "strings"

// Assignment represents a assignment to a variable
type Assignment struct {
	BaseNode
	Variable VariableReference
	Value    *Expression
}

func (node *Assignment) ToString(tabLevel int) string {
	tabs := strings.Repeat("\t", tabLevel)
	str := tabs + node.Variable.VarName + " = " + node.Value.ToString(0)
	return str
}
