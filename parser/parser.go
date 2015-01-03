package parser

import (
	"fmt"
	"github.com/zakuro9715/jingu/ast"
)

type Parser struct {
	lexer Lexer
}

func (p *Parser) Init(src string) {
	p.lexer.Init(src)
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

    if token.Type == ast.Unknown {
      errs = append(errs, NewSyntaxError(fmt.Sprintf("Found unexpected %v", token.Text), token.Line, token.Column))
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
