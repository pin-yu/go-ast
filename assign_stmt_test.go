package goast

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssignStmt(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "./assign_stmt.go", nil, parser.ParseComments)
	if err != nil {
		t.Fatal(err)
	}

	actuals := []*ast.AssignStmt{}

	ast.Inspect(f, func(n ast.Node) bool {
		a, ok := n.(*ast.AssignStmt)
		if ok {
			actuals = append(actuals, a)
		}

		return true
	})

	// lhs
	assert.Equal(t, "a", actuals[0].Lhs[0].(*ast.Ident).Name)
	assert.Equal(t, "b", actuals[0].Lhs[1].(*ast.Ident).Name)

	// token
	assert.Equal(t, token.DEFINE, actuals[0].Tok)

	// rhs
	assert.Equal(t, token.INT, actuals[0].Rhs[0].(*ast.BasicLit).Kind)
	assert.Equal(t, "10", actuals[0].Rhs[0].(*ast.BasicLit).Value)
	assert.Equal(t, "input", actuals[0].Rhs[1].(*ast.Ident).Name)
}
