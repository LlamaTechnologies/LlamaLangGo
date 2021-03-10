package interrepr

import (
	ast "LlamaLangCompiler/cmd/ast"

	"github.com/llir/llvm/ir"
)

// genIRCode generates IR code for a BaseNode in a program m
func genBaseNodeIRCode(m *ir.Module, node interface{}) error {
	switch node.(type) {
	case *ast.BinaryExprEnum:

	}
	return nil
}
