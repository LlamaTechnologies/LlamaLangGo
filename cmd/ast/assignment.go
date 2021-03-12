package ast

// Assignment represents a assignment to a variable
type Assignment struct {
	BaseNode
	Variable VariableReference
	Value    Node
}

func NewAssignmentNode(varRefName string) *Assignment {
	assignmentNode := new(Assignment)
	assignmentNode.Variable = VariableReference{
		VarName: varRefName,
	}
	return assignmentNode
}

func (node *Assignment) ToString(tabLevel int) string {
	tabs := GetTabs(tabLevel)
	str := tabs + node.Variable.VarName + " = " + node.Value.ToString(0)
	return str
}
