package ast

type Visitor interface {
	VisitTerminal(*TerminalAST)
	VisitWhile(*WhileAST)
}
