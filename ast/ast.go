package ast

type TokenType int

const (
	Unknown    TokenType = iota
	PtrInc               // >
	PtrDec               // <
	ValInc               // +
	ValDec               // -
	Print                // .
	Scan                 // ,
	WhileStart           // [
	WhileEnd             // ]
)

type Token struct {
	Type   TokenType
	Line   int
	Column int
	Text   string
}

type TokenImpl struct {
	token Token
}

func (t *TokenImpl) Token() Token {
	return t.token
}

func (t *TokenImpl) SetToken(token Token) {
	t.token = token
}

type AST interface {
	Token() Token
	Accept(Visitor)
}

type TerminalAST struct {
	TokenImpl
}

func (this *TerminalAST) Accept(v Visitor) {
	v.VisitTerminal(this)
}

type WhileAST struct {
	TokenImpl
	Children []AST
}

func (this *WhileAST) Accept(v Visitor) {
	v.VisitWhile(this)
}
