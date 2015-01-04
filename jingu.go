package main

import (
	"fmt"
	"github.com/zakuro9715/jingu/ast"
	"github.com/zakuro9715/jingu/parser"
	"io"
)

const (
	prompt = ">>> "
)

func main() {
	interactive(new(ast.Interpreter))
}

func interactive(visitor ast.Visitor) {
	for {
		var src string
		fmt.Print(prompt)
		_, err := fmt.Scanln(&src)

		if err == io.EOF {
			fmt.Println("\nBye")
			return
		}

		p := parser.Parser{}
		p.Init(src)
		asts, errs := p.Parse()

		for _, err := range errs {
			fmt.Println(err)
		}

		if len(errs) == 0 {
			visitor.Init(30000)
			visitor.Visit(asts)
		}
	}
}
