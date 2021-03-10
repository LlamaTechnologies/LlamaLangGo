package ast

// BaseNode is the Base AST type
type BaseNode struct {
	_NodePrinter
	FileName   string
	LineNumber int
}

type _NodePrinter interface {
	toString(int) string
}
