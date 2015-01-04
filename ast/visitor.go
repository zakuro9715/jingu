package ast

import (
	"fmt"
)

type Visitor interface {
	VisitTerminal(*TerminalAST)
	VisitWhile(*WhileAST)
}

type Interpreter struct {
	memory []int
	ptr    int
}

func (it *Interpreter) Init(memorySize int) {
	it.memory = make([]int, memorySize)
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
		fmt.Scanf("%c", &it.memory[it.ptr])
	}
}

func (it *Interpreter) VisitWhile(tree *WhileAST) {
	for it.memory[it.ptr] != 0 {
		for _, child := range tree.Children {
			child.Visit(it)
		}
	}
}
