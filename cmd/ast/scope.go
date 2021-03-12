package ast

type ScopeTypeEnum int

const (
	GLOBAL_SCOPE   ScopeTypeEnum = iota
	FUNCTION_SCOPE               = iota
)

// Scope represent a symbol table for a given scope
type Scope struct {
	symbols   map[string][]Node
	Children  []*Scope
	Parent    *Scope
	Data      Node
	ScopeType ScopeTypeEnum
}

func NewScope(scopeType ScopeTypeEnum, data Node, parent *Scope) *Scope {
	scope := new(Scope)
	scope.Parent = parent
	scope.ScopeType = scopeType

	if scopeType == FUNCTION_SCOPE {
		scope.Data = data
	}
	scope.symbols  = make(map[string][]Node)
	scope.Children = make([]*Scope, 0)
	return scope
}

// FindSymbol Finds a symbols in the scope
func (scope *Scope) FindSymbol(name string, justCurrent bool) Node {
	return scope.Data
}

func (scope *Scope) AddChildren(scopeType ScopeTypeEnum, name string, symbol *FunctionDefinition) *Scope {
	scope.symbols[name] = append(scope.symbols[name], symbol)

	childScope := NewScope(scopeType, symbol, scope)
	scope.Children = append(scope.Children, childScope)

	return childScope
}

func (scope *Scope) AddSymbol(name string, symbol Node) {
	scope.symbols[name] = append(scope.symbols[name], symbol)
}
