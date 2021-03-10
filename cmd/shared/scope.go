package shared

import ast "LlamaLangCompiler/cmd/ast"

// Scope represent a symbol table for a given scope
type Scope struct {
	symbols  []ast.BaseNode
	children []Scope
	data     ast.BaseNode
}

// FindSymbol Finds a symbols in the scope
func (current *Scope) FindSymbol(name string, justCurrent bool) *ast.BaseNode {
	return &current.data
}
