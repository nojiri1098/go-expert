package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "analyzer",
	Doc:      "reports functions with high cyclomatic complexity",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	i := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}
	i.Preorder(nodeFilter, func(n ast.Node) {
		fn := n.(*ast.FuncDecl)
		c := Count(fn)

		if c > 10 {
			fn.Pos().IsValid()
			pass.Reportf(fn.Pos(), "function %s has cyclomatic complexity of %d", fn.Name.Name, c)
		}
	})

	return nil, nil
}
