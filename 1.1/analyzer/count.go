package analyzer

import (
	"fmt"
	"go/ast"
)

func main() {
	s := "Hello, world!"
	fmt.Printf("%s", s)
}

func Count(node ast.Node) int {
	count := 1

	ast.Inspect(node, func(n ast.Node) bool {
		switch n.(type) {
		case *ast.IfStmt:
			count++
		case *ast.ForStmt:
			count++
		}
		return true
	})

	return count
}
