package ast

import "strings"

// BaseNode is the Base AST type
type BaseNode struct {
	FileName   string
	LineNumber int
}

type Node interface {
	ToString(int) string
}

func GetTabs(tabLevel int) string {
	return strings.Repeat(" ", tabLevel * 4)
}

const beginLine = "\b\b>>"
func GetLineBeginChar() string {
	return beginLine
}
