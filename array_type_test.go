package goast

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayType(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "./array_type.go", nil, parser.ParseComments)
	if err != nil {
		t.Fatal(err)
	}

	actuals := []*ast.ArrayType{}

	ast.Inspect(f, func(n ast.Node) bool {
		a, ok := n.(*ast.ArrayType)
		if ok {
			actuals = append(actuals, a)
		}
		return true
	})

	assert.Equal(t, "int", actuals[0].Elt.(*ast.Ident).Name)
	assert.Equal(t, token.INT, actuals[0].Len.(*ast.BasicLit).Kind)
	assert.Equal(t, "10", actuals[0].Len.(*ast.BasicLit).Value)

	assert.Equal(t, "int", actuals[1].Elt.(*ast.Ident).Name)
	assert.Nil(t, actuals[1].Len)
}
