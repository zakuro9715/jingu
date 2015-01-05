package parser

import (
	"github.com/zakuro9715/jingu/ast"
	"github.com/zakuro9715/jingu/core"
)

type Parser struct {
	lexer Lexer
}

func (p *Parser) Init(src string, config *core.Config) {
	p.lexer.Init(src, config)
}

func (p *Parser) Config() *core.Config {
	return p.lexer.config
}

func (p *Parser) ParseStmt(isInWhile bool) ([]ast.AST, []error) {
	asts := []ast.AST{}
	errs := []error{}

	for p.lexer.Current() != EOF {
		token, err := p.lexer.Scan()

		if err != nil {
			errs = append(errs, err)
		}

		if token.Type == ast.WhileStart {
			var innerErrs []error
			tree := &ast.WhileAST{}
			tree.SetToken(token)
			tree.Children, innerErrs = p.ParseStmt(true)
			asts = append(asts, tree)
			errs = append(errs, innerErrs...)
		} else if token.Type == ast.WhileEnd {
			if !isInWhile {
				errs = append(errs, NewSyntaxError("Found unexpected ]", token.Line, token.Column))
			}
			return asts, errs
		} else {
			tree := &ast.TerminalAST{}
			tree.SetToken(token)
			asts = append(asts, tree)
		}
	}
	if isInWhile {
		errs = append(errs, NewSyntaxError("] is expected but EOF found.", 0, 0))
	}
	return asts, errs
}

func (p *Parser) Parse() ([]ast.AST, []error) {
	return p.ParseStmt(false)
}
