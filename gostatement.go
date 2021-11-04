package gostatement

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "gostatement",
	Doc:      "gostatement is an analyzer checking for occurrence of `go` statements",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.GoStmt)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		pass.Reportf(node.Pos(), "go statement found")
	})

	return nil, nil
}
