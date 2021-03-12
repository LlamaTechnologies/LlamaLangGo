package ast

// Program is the ast node representing
// the whole program
type Program struct {
	BaseNode
	Name     string
	Children []Node
}

func NewProgramNode(executableName string) *Program {
	node := new(Program)
	node.Name = executableName
	node.FileName = node.Name
	node.LineNumber = 0
	node.Children = make([]Node, 0)
	return node
}


func (node *Program) ToString(tabLevel int) string {
	tabs := GetTabs(tabLevel)
	beginChar := GetLineBeginChar()

	str := tabs + beginChar + " Program '" + node.Name + "':"

	for _, child := range node.Children {
		str += "\n"
		str += child.ToString(tabLevel + 1)
	}

	return str
}
