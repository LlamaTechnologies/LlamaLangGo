package ast

// LiteralTypeEnum enumType
type LiteralTypeEnum int

// LiteralTypeEnum values
const (
	CONST_DEC    LiteralTypeEnum = iota
	CONST_OCT                    = iota
	CONST_HEX                    = iota
	CONST_FLOAT                  = iota
	CONST_DOUBLE                 = iota
	CONST_CHAR                   = iota
	CONST_STRING                 = iota
)

// LiteralConstant represents a literalConstant
type LiteralConstant struct {
	BaseNode
	ConstantID LiteralTypeEnum
	Value      string
}

func (node *LiteralConstant) ToString(tabLevel int) string {
	tabs := GetTabs(tabLevel)
	str := tabs + node.Value
	return str
}
