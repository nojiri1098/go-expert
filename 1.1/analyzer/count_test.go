package analyzer

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestCount(t *testing.T) {
	tests := []struct {
		name       string
		code       string
		complexity int
	}{
		{
			name: "simple",
			code: `package main
func Double(n int) int {
	return n * 2
}`,
			complexity: 1,
		},
		{
			name: "if statement",
			code: `package main
func Double(n int) int {
	if n%2 == 0 {
		return 0
	}
	return n
}`,
			complexity: 2,
		},
		{
			name: "for statement",
			code: `package main
func Sum(n int) int {
	c := 0
	for i := 0; i < n; i++ {
		c += i
	}
	return c
}`,
			complexity: 2,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			a := GetFuncNode(t, test.code)
			c := Count(a)

			if c != test.complexity {
				t.Errorf("got %d, want %d", c, test.complexity)
			}
		})
	}
}

func GetFuncNode(t *testing.T, code string) ast.Node {
	t.Helper()

	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, "", code, 0)
	if err != nil {
		t.Fatal(err)
	}
	for _, decl := range file.Decls {
		if fd, ok := decl.(*ast.FuncDecl); ok {
			return fd
		}
	}
	t.Fatal("no function declare found")
	return nil
}
