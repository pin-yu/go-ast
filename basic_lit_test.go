package goast

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTest(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "./basic_lit.go", nil, parser.ParseComments)
	if err != nil {
		t.Fatal(err)
	}

	actuals := []*ast.BasicLit{}

	ast.Inspect(f, func(n ast.Node) bool {
		b, ok := n.(*ast.BasicLit)
		if ok {
			actuals = append(actuals, b)
		}

		return true
	})

	assert.Equal(t, "`json:\"id\"`", actuals[0].Value)
	assert.Equal(t, token.STRING, actuals[0].Kind)

	assert.Equal(t, "\"this is an id\"", actuals[1].Value)
	assert.Equal(t, token.STRING, actuals[1].Kind)

	assert.Equal(t, "87", actuals[2].Value)
	assert.Equal(t, token.INT, actuals[2].Kind)
}
