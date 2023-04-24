package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

type visitor struct {
	fset *token.FileSet
}

// Visit function checks for constructs of type `<variable> + 42` and prints them to stdout.
func (v *visitor) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return nil
	}

	// Check whether given AST node is a binary expression.
	expr, ok := node.(*ast.BinaryExpr)
	if !ok {
		return v
	}

	// Check what's the operation performed.
	if expr.Op != token.ADD {
		return v
	}

	// Check whether the left operand is an identifier.
	left, ok := expr.X.(*ast.Ident)
	if !ok {
		return v
	}

	// Check whether the right operand is a literal.
	right, ok := expr.Y.(*ast.BasicLit)
	if !ok {
		return v
	}

	// Check the type of the literal.
	if right.Kind != token.INT {
		return v
	}

	// Check the value of the integer literal.
	if right.Value == "42" {
		fmt.Printf("%s: adding 42 to variable %s\n", v.fset.Position(right.Pos()), left.Name)
	}

	return v
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("expected 2 arguments, got %d\n", len(os.Args))
	}

	// file set holds information about the parsed files.
	fset := token.NewFileSet()
	root, err := parser.ParseFile(fset, os.Args[1], nil, parser.ParseComments|parser.AllErrors)
	if err != nil {
		fmt.Printf("could not parse source code: %v\n", err)
		os.Exit(1)
	}

	if root == nil {
		fmt.Println("root is nil!")
		os.Exit(1)
	}

	v := &visitor{fset: fset}
	ast.Walk(v, root)
}
