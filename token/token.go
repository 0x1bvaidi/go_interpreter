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

var keywords = map[string]TokenType{
	"phi": FUNCTION,
	"let": LET,
	"return": RETURN,
}

//  lookupIdent checks the keywords table to see whether the given identifier is in fact a keyword.
//  If it is, it returns the keyword’s TokenType constant. If it isn’t, we just get back token.IDENT
func LookUpIdent(ident string) TokenType{
	if tok, ok := keywords[ident]; ok{
		return tok
	}
	return IDENT
}
