package lexer

import "github.com/0x1bvaidi/go_interpreter/token"

type Lexer struct {
	input        string
	Position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// the purpose of readChar() is to give us the next character & advance our postion in the input string
// l.readPosition always points to the next position and l.position points to the one we just read
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // if input string ends or about to begin, it sets l.ch to 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.Position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
