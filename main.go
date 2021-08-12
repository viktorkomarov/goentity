package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
)

type StructFinder struct {
	strts []*ast.StructType
}

func (s *StructFinder) Visit(node ast.Node) ast.Visitor {
	strct, ok := node.(*ast.StructType)
	if !ok {
		return s
	}

	s.strts = append(s.strts, strct)
	return s
}

func main() {
	src := os.Getenv("GOFILE")

	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, src, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	walker := &StructFinder{
		strts: make([]*ast.StructType, 0),
	}

	ast.Walk(walker, f)

}
