package ast

import (
	"fmt"
	"github.com/zakuro9715/jingu/core"
)

type Visitor interface {
	VisitTerminal(*TerminalAST)
	VisitWhile(*WhileAST)
	Visit([]AST)
	Init(*core.Config)
	Config() *core.Config
}

type Interpreter struct {
	memory []int
	ptr    int
	config *core.Config
}

func (it *Interpreter) Init(config *core.Config) {
	it.memory = make([]int, config.MemorySize)
	it.config = config
	it.ptr = 0
}

func (it *Interpreter) Config() *core.Config {
  return it.config
}

func (it *Interpreter) VisitTerminal(tree *TerminalAST) {
	switch tree.Token().Type {
	case PtrInc:
		it.ptr++
	case PtrDec:
		it.ptr--
	case ValInc:
    it.memory[it.ptr]++
		if it.memory[it.ptr] > 255 {
      it.memory[it.ptr] -= 256
    }
	case ValDec:
    it.memory[it.ptr]--
		if it.memory[it.ptr] < 0 {
      it.memory[it.ptr] += 256
    }
	case Print:
		fmt.Printf("%c", it.memory[it.ptr])
	case Scan:
		var str string
		fmt.Scanln(&str)
		fmt.Sscanf(str, "%c", &it.memory[it.ptr])
	}
}

func (it *Interpreter) VisitWhile(tree *WhileAST) {
	for it.memory[it.ptr] != 0 {
		it.Visit(tree.Children)
	}
}

func (it *Interpreter) Visit(asts []AST) {
	for _, tree := range asts {
		tree.Accept(it)
	}
}
