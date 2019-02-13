package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type Result struct {
	Package string
	Types   []string
}

type visitor struct {
	lookFor string
	result  *Result
}

func Parse(file string, structName string) (result Result, err error) {
	fileSet := token.NewFileSet()
	f, err := parser.ParseFile(fileSet, file, nil, parser.ParseComments)
	if err != nil {
		return
	}

	ast.Walk(&visitor{result: &result, lookFor: structName}, f)

	fmt.Printf("%+v", result)
	return result, nil
}

func (v *visitor) Visit(n ast.Node) (w ast.Visitor) {
	switch n := n.(type) {
	case *ast.Package:
		return v
	case *ast.File:
		v.result.Package = n.Name.String()
		return v
	case *ast.GenDecl:
		return v
	case *ast.TypeSpec:
		if v.lookFor == "" || v.lookFor == n.Name.String() {
			v.result.Types = append(v.result.Types, n.Name.String())
		}
		return nil
	}
	return nil
}
