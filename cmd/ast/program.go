package ast

import (
	"strings"
)

// Program is the ast node representing
// the whole program
type Program struct {
	BaseNode
	Name     string
	Children []BaseNode
}

func (node *Program) toString(tabLevel int) string {
	tabs := strings.Repeat("\t", tabLevel)
	str := tabs + node.Name + ":"

	for _, child := range node.Children {
		str += child.toString(tabLevel + 1)
	}

	return str
}
