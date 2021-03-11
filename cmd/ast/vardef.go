package ast

import "strings"

// VariableDefinition represents a variable definition in the code
type VariableDefinition struct {
	*BaseNode
	Name                string
	VarType             string
	AssignmentStatement *Assignment
}

func (node *VariableDefinition) ToString(tabLevel int) string {
	tabs := strings.Repeat("\t", tabLevel)
	str := tabs + node.Name + ":" + node.VarType

	if node.AssignmentStatement != nil {
		str += node.AssignmentStatement.ToString(0)
	}

	return str
}
