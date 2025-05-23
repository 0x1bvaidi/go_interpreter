package parser

import (
	"github.com/0x1bvaidi/go_interpreter/ast"
	"github.com/0x1bvaidi/go_interpreter/lexer"
	"github.com/0x1bvaidi/go_interpreter/token"
	"fmt"
)

// this one struct has another struct as its field
type Parser struct {
	l *lexer.Lexer // pointer to Lexer struct in lexer package
	errors []string // for error handling
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	temp := Parser{
		l: l,
		errors: []string{},
		}

	p := &temp // p stores the address of the temp, so p is of type *Parser

	p.nextToken()
	p.nextToken()

	return p
}

// error handling
func (p *Parser) Errors() []string{
	 return	p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := 	fmt.Sprintf("Expected next token to be %s. got %s", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.ConsumeToken_Advance()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
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
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement{
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT){
		return nil
	}


	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN){
		return nil
	}	

	 // TODO: We're skipping the expressions until we
	 // encounter a stop '@'

	 for !p.curTokenIs(token.STOP){
		 p.nextToken()
	 }

	 return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool{
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool{
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool{
	if p.peekTokenIs(t){
		p.nextToken()
		return true
	}else{
		p.peekError(t)
		return false
	}
}


