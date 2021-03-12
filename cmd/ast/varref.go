package ast

// VariableReference represents a variable that is being referenced in a statement
type VariableReference struct {
	BaseNode
	VarName string
}

func NewVariableReference(varName string) *VariableReference {
	node := new(VariableReference)
	node.VarName = varName
	return node
}

func (node *VariableReference) ToString(tabLevel int) string {
	tabs := GetTabs(tabLevel)
	str := tabs + node.VarName
	return str
}
