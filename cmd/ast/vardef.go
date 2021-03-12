package ast

// VariableDefinition represents a variable definition in the code
type VariableDefinition struct {
	BaseNode
	Name                string
	VarType             string
	IsGlobal            bool
	IsParam             bool
	AssignmentStatement *Assignment
}

func (node *VariableDefinition) ToString(tabLevel int) string {
	tabs := GetTabs(tabLevel)
	beginChar := GetLineBeginChar() + " "
	if node.IsParam {
		beginChar = ""
	}

	str := tabs + beginChar
	str += node.Name + ": " + node.VarType

	if node.AssignmentStatement != nil {
		str += "\n" + tabs + beginChar
		str += node.AssignmentStatement.ToString(0)
	}

	return str
}
