package lexer

import (
	"github.com/0x1bvaidi/go_interpreter/token"
)

type Lexer struct {
	input        string
	position     int
	nextPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.advance()
	return l
}

// the struct name in token is Token, confusion due to same name but different way of writing
// the method in the book is NextToken()
func (l *Lexer) consumeToken_Advance() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '@':
		tok = newToken(token.STOP, l.ch)
	case '+':
		tok = newToken(token.ADD, l.ch)
	case '$':
		tok = newToken(token.COMMENT, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.MULT, l.ch)
	case '/':
		tok = newToken(token.DIVIDE, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch){
			tok.Literal = l.readIdentifier_advance()
			return tok
		}else{
			tok = newToken(tok.ILLEGAL, l.ch)
		}
	}
	l.advance()
	return tok
}

func (l *Lexer) readIdentifier_advance() string{
	position := l.position
	for isLetter(l.ch){
		l.advance()
	}
	return l.input[position: l.position]
}

func isLetter(ch byte) bool{
	return 'a' <= ch &&  ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func newToken(tokenType token.TokenType, ch byte) token.Token{
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// this method advances the current symbols in the input string
// this method in the book is readChar()
func (l *Lexer) advance() {
	if l.nextPosition >= len(l.input) {
		l.ch = 0 // 0 represents EOF
	} else {
		l.ch = l.input[l.nextPosition]
	}
	l.position = l.nextPosition
	l.nextPosition++
}
