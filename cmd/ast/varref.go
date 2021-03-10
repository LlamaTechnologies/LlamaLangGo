package ast

import "strings"

// VariableReference represents a variable that is being referenced in a statement
type VariableReference struct {
	BaseNode
	VarName string
}

func (node *VariableReference) toString(tabLevel int) string {
	tabs := strings.Repeat("\t", tabLevel)
	str := tabs + node.VarName
	return str
}
