package ast

import "strings"

// BaseNode is the Base AST type
type BaseNode struct {
	_NodePrinter
	FileName   string
	LineNumber int
}

type _NodePrinter interface {
	ToString(int) string
}

func GetTabs(tabLevel int) string {
	return strings.Repeat(" ", tabLevel * 4)
}

func GetLineBeginChar(tabLevel int) string {
	backspacesCount := 0

	if tabLevel != 0 {
		backspacesCount = tabLevel * 4 - 2
	}

	backspaces :=  strings.Repeat("\b", backspacesCount)

	return backspaces + ">>"
}
