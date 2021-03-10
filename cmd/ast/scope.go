package ast

// Scope represent a symbol table for a given scope
type Scope struct {
	symbols  []BaseNode
	children []Scope
	data     BaseNode
}

// FindSymbol Finds a symbols in the scope
func (current *Scope) FindSymbol(name string, justCurrent bool) *BaseNode {
	return &current.data
}
