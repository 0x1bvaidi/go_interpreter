package lexer

import (
	"github.com/0x1bvaidi/go_interpreter/token"
	"testing"
)

// so we will write a function in which we will feed the input program as text
// and it will generate the tokens and display them.
// it will expect certain kind of tokentype with respect to the input
// and it will expect a literal that is the exact text of the token that will be in the source code.

func TestingNextToken(t *testing.T) {
	input := `let five = 5@
			  let ten = 10@
			  let add = phi(x, y) {
			  x + y@
 			}
		 let result = add(five,ten)@
 	`

	// i will test the input on the basis of this struct
	tests := []struct {
		expectedType    token.TokenType // the TokenType type string defintion from token package
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.STOP, "@"},

		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.STOP, "@"},

		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "phi"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.ADD, "+"},
		{token.IDENT, "y"},
		{token.STOP, "@"},
		{token.RBRACE, "}"},
		{token.STOP, "@"},

		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.STOP, "@"},
		{token.EOF, ""},
	}

	// New is the user defined function
	l := New(input)

	for i, tt := range tests {
		tok := l.ConsumeToken_Advance()

		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d] - tokentype is wrong. expected: %q, received: %q\n", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("test[%d] - literal is wrong. expected: %q, received: %q\n", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
