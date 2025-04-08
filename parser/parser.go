package parser

import (
	"github.com/0x1bvaidi/go_interpreter/ast"
	"github.com/0x1bvaidi/go_interpreter/lexer"
	"github.com/0x1bvaidi/go_interpreter/token"
)

// this one struct has another struct as its field
type Parser struct {
	l *lexer.Lexer // pointer to Lexer struct in lexer package

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	temp := Parser{l: l}
	p := &temp // p stores the address of the temp, so p is of type *Parser

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.ConsumeToken_Advance()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != Token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case Token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}
