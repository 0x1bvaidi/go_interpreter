package lexer

import (
	"github.com/0x1bvaidi/go_interpreter.git"
	"testing"
)

// so we will write a function in which we will feed the input program as text
// and it will generate the tokens and display them.
// it will expect certain kind of tokentype with respect to the input
// and it will expect a literal that is the exact text of the token that will be in the source code.

func TestingNextToken(t *testing.T) {
	input = `+-*/=$@{}()`

	// i will test the input on the basis of this struct
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ADD, "+"},
		{token.MINUS, "-"},
		{token.MULT, "*"},
		{token.DIVIDE, "/"},
		{token.ASSIGN, "="},
		{token.COMMENT, "$"},
		{token.STOP, "@"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.EOF, ""},

	}
}
