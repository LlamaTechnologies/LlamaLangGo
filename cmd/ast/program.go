package ast

// Program is the ast node representing
// the whole program
type Program struct {
	BaseNode
	Name     string
	Children []BaseNode
}

func (node *Program) ToString(tabLevel int) string {
	tabs := GetTabs(tabLevel)
	beginChar := GetLineBeginChar(tabLevel)

	str := tabs + beginChar + " Program '" + node.Name + "':"

	for _, child := range node.Children {
		str += child.ToString(tabLevel + 1)
	}

	return str
}
