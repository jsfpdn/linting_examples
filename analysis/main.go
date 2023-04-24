package main

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/analysis/singlechecker"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "plus42",
	Doc:  "Check whether code contains constructs of form '<variable> + 42'",
	// inspect.Analyzer is required to filter the uninteresting nodes.
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (any, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.BinaryExpr)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		expr := node.(*ast.BinaryExpr)
		if expr.Op != token.ADD {
			return
		}

		// Check whether the left operand is an identifier.
		left, ok := expr.X.(*ast.Ident)
		if !ok {
			return
		}

		// Check whether the right operand is a literal.
		right, ok := expr.Y.(*ast.BasicLit)
		if !ok {
			return
		}

		// Check the type of the literal.
		if right.Kind != token.INT {
			return
		}

		// Check the value of the integer literal.
		if right.Value == "42" {
			pass.Report(analysis.Diagnostic{
				Pos:     left.Pos(),
				End:     right.End(),
				Message: "found '<variable> + 42'",
				SuggestedFixes: []analysis.SuggestedFix{
					{
						Message: "use 43 instead?",
						TextEdits: []analysis.TextEdit{
							{
								Pos:     right.Pos(),
								End:     right.End(),
								NewText: []byte("43"),
							},
						},
					},
				},
			})
		}
	})

	return nil, nil
}

func main() {
	singlechecker.Main(Analyzer)
}
