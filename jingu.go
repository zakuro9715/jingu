package main

import (
	"fmt"
	"github.com/zakuro9715/jingu/parser"
	"github.com/zakuro9715/jingu/ast"
)

func main() {
  var src string
  fmt.Scanln(&src)
  p := parser.Parser{}
  p.Init(src)
  asts, errs := p.Parse()

  for _, err := range errs {
    fmt.Println(err)
  }

  visitor := new(ast.Interpreter)
  visitor.Init(30000)
  for _, node := range asts {
    node.Visit(visitor)
  }
  fmt.Println()
}
