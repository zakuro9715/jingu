package ast

import (
	"fmt"
)

type Visitor interface {
	VisitTerminal(*TerminalAST)
	VisitWhile(*WhileAST)
  Visit([]AST)
  Init(int)
}

type Interpreter struct {
	memory []int
	ptr    int
}

func (it *Interpreter) Init(memorySize int) {
	it.memory = make([]int, memorySize)
  it.ptr = 0
}

func (it *Interpreter) VisitTerminal(tree *TerminalAST) {
	switch tree.Token().Type {
	case PtrInc:
		it.ptr++
	case PtrDec:
		it.ptr--
	case ValInc:
		it.memory[it.ptr]++
	case ValDec:
		it.memory[it.ptr]--
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
    tree.Visit(it)
  }
}
