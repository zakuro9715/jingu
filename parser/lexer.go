package parser

import (
	"github.com/zakuro9715/jingu/ast"
	"github.com/zakuro9715/jingu/core"
)

type Lexer struct {
	src      []rune
	offset   int
	lineHead int
	line     int
	config   *core.Config
}

const (
	EOF = -1
)

func (l *Lexer) Init(src string, config *core.Config) {
	l.src = []rune(src)
	l.config = config
}

func (l *Lexer) Current() rune {
	if len(l.src) <= l.offset {
		return EOF
	}
	return l.src[l.offset]
}

func (l *Lexer) Consume() {
	l.offset++
}

func (l *Lexer) tokenType() ast.TokenType {
	switch l.Current() {
	case '>':
		return ast.PtrInc
	case '<':
		return ast.PtrDec
	case '+':
		return ast.ValInc
	case '-':
		return ast.ValDec
	case '.':
		return ast.Print
	case ',':
		return ast.Scan
	case '[':
		return ast.WhileStart
	case ']':
		return ast.WhileEnd
	default:
		return ast.Unknown
	}
}

func (l *Lexer) Scan() (ast.Token, error) {
	line := l.line
	column := l.offset - l.lineHead
	tokenType := l.tokenType()
	text := string(l.Current())

	l.Consume()

	return ast.Token{
		Type:   tokenType,
		Line:   line,
		Column: column,
		Text:   text,
	}, nil
}
