package ast

type FunctionDefinition struct {
	*BaseNode
	Name       string
	RetType    string
	Parameters []*VariableDefinition
	Block      []*BaseNode
}

func (node *FunctionDefinition) ToString(tabLevel int) string {
	tabs := GetTabs(tabLevel)
	beginChar := GetLineBeginChar(tabLevel)

	str := tabs + beginChar + " func " + node.Name + ": " + node.RetType + "{"

	childLevelType := tabLevel + 1
	if len(node.Parameters) > 0 {
		for _, param := range node.Parameters {
			str += "\n"
			str += param.ToString(childLevelType)
		}
		str += "\n" + tabs
	}
	str += ") {"

	for _, stmnt := range node.Block {
		str += "\n"
		str += stmnt.ToString(childLevelType)
	}

	str += "\n" + tabs + "}"

	return str
}
