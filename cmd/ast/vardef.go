package ast

// VariableDefinition represents a variable definition in the code
type VariableDefinition struct {
	BaseNode
	Name                string
	VarType             string
	IsGlobal            bool
	AssignmentStatement *Assignment
}

func (node *VariableDefinition) ToString(tabLevel int) string {
	tabs := GetTabs(tabLevel)
	beginChar := GetLineBeginChar(tabLevel) + " "

	str := node.putBeginChar(tabs, beginChar, "", true)
	str += node.Name + ": " + node.VarType

	if node.AssignmentStatement != nil {
		str += node.putBeginChar(tabs, beginChar, "\n", false)
		str += node.AssignmentStatement.ToString(0)
	}

	return str
}

func (node *VariableDefinition) putBeginChar(tabs, beginChar, char string, putLast bool) string {
	var str string

	if !putLast {
		str += char
	}

	str += tabs

	if node.IsGlobal {
		str += beginChar
	}

	if putLast {
		str += char
	}

	return str
}
