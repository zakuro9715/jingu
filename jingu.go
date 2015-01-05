package jingu

import (
	"fmt"
	"github.com/zakuro9715/jingu/core"
	"github.com/zakuro9715/jingu/ast"
	"github.com/zakuro9715/jingu/parser"
	"io"
)

const (
	prompt = ">>> "
)

func main() {
	interactive(new(ast.Interpreter), new(core.Config))
}

func interactive(visitor ast.Visitor, config *core.Config) {
	for {
		var src string
		fmt.Print(prompt)
		_, err := fmt.Scanln(&src)

		if err == io.EOF {
			fmt.Println("\nBye")
			return
		}

		p := parser.Parser{}
		p.Init(src, config)
		asts, errs := p.Parse()

		for _, err := range errs {
			fmt.Println(err)
		}

		if len(errs) == 0 {
			visitor.Init(30000, config)
			visitor.Visit(asts)
		}
	}
}
