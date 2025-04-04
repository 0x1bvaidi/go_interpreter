package token

type TokenType string

// type = "string which is the value of the token of that type"
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// literals and identifiers
	IDENT = "IDENT"
	INT   = "INT"

	//delimiters
	COMMENT = "$"
	STOP    = "@"
	RBRACE  = "{"
	LBRACE  = "}"
	RPAREN  = "("
	LPAREN  = ")"
	COMMA   = ","

	//Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	RETURN   = "RETURN"

	//operators
	ASSIGN = "="
	MULT   = "*"
	ADD    = "+"
	MINUS  = "-"
	DIVIDE = "/"
)
