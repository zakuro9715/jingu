package main

import (
	"flag"
	"fmt"
	"github.com/zakuro9715/jingu/ast"
	"github.com/zakuro9715/jingu/core"
	"github.com/zakuro9715/jingu/parser"
	"io"
	"os"
)

const (
	prompt = ">>> "
)

const (
  UnknownCode = -1
	successCode = iota
	invalidFormatCode
	notImplementedCode
)

func usage(exitCode int) {
	fmt.Println("Usage: jingu [flags] [programfile]")
	flag.PrintDefaults()
	os.Exit(exitCode)
}

func notImplemented() {
	os.Exit(notImplementedCode)
}

func main() {
	isInteractive := flag.Bool("i", false, "enable interactive mode")
	isCompile := flag.Bool("c", false, "enable compile")
	isUsage := flag.Bool("h", false, "display this help and exit")

	flag.Parse()

	var visitor ast.Visitor

	if *isUsage {
		usage(successCode)
	}

	if *isCompile {
		fmt.Println("Sorry, I will implement soon.")
		notImplemented()
	} else {
		visitor = new(ast.Interpreter)
	}

	config := new(core.Config)

	if *isInteractive {
		if len(flag.Args()) > 0 {
			usage(invalidFormatCode)
		}
		runInteractive(visitor, config)
	} else {
		if len(flag.Args()) != 1 {
			usage(invalidFormatCode)
		}
		runFromFile(visitor, config, flag.Arg(0))
	}
}

func runFromFile(visitor ast.Visitor, config *core.Config, path string) {
	notImplemented()
}

func runInteractive(visitor ast.Visitor, config *core.Config) {
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
