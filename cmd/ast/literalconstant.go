package ast

import "strings"

// LiteralTypeEnum enumType
type LiteralTypeEnum int

// LiteralTypeEnum values
const (
	NUMBER LiteralTypeEnum = iota
	FLOAT                  = iota
	STRING                 = iota
)

// LiteralConstant represents a literalConstant
type LiteralConstant struct {
	BaseNode
	ConstantID LiteralTypeEnum
	Value      string
}

func (node *LiteralConstant) ToString(tabLevel int) string {
	tabs := strings.Repeat("\t", tabLevel)
	str := tabs + node.Value
	return str
}
