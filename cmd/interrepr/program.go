package interrepr

import (
	ast "LlamaLangCompiler/cmd/ast"
	"fmt"

	"github.com/llir/llvm/ir"
)

// GenIRCode generates IR code from a program node
func GenIRCode(node *ast.Program) error {
	// Create a new LLVM IR module.
	m := ir.NewModule()

	for _, child := range node.Children {
		err := genBaseNodeIRCode(m, child)
		if err != nil {
			return err
		}
	}

	// Print the LLVM IR assembly of the module.
	fmt.Println(m)
	return nil
}
