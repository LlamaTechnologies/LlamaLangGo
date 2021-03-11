package ast

type ScopeTypeEnum int

const (
	GLOBAL_SCOPE   ScopeTypeEnum = iota
	FUNCTION_SCOPE               = iota
)

// Scope represent a symbol table for a given scope
type Scope struct {
	Symbols   map[string][]*BaseNode
	Children  []*Scope
	Parent    *Scope
	Data      *BaseNode
	ScopeType ScopeTypeEnum
}

func NewScope(scopeType ScopeTypeEnum, data *BaseNode, parent *Scope) *Scope {
	scope := new(Scope)
	scope.Parent = parent
	scope.ScopeType = scopeType

	if scopeType == FUNCTION_SCOPE {
		scope.Data = data
	}

	return scope
}

// FindSymbol Finds a symbols in the scope
func (scope *Scope) FindSymbol(name string, justCurrent bool) *BaseNode {
	return scope.Data
}

func (scope *Scope) AddChildren(scopeType ScopeTypeEnum, name string, symbol *FunctionDefinition) *Scope {
	scope.Symbols[name] = append(scope.Symbols[name], symbol.BaseNode)

	childScope := NewScope(scopeType, symbol.BaseNode, scope)
	scope.Children = append(scope.Children, childScope)

	return childScope
}
