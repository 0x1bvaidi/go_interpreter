package ast

import(
	"github.com/0x1bvaidi/go_interpreter/token"
)

// starting point of every ast
type Node interface{
	TokenLiteral() string // return the token that starts this ast
}

//ast nodes that are statements
type Statement interface{
	Node
	statementNode()
}

// ast nodes that are expressions
type Expression interface{
	Node
	expressionNode()
}

// my program that will be parsed
type Program struct{
	Statements []Statement // it is the slice of all the ast that are statements	
}

func (p *Program) TokenLiteral() string{
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral() // returns the first token in the []Statement as a string
	}else{
		return ""
	}	
}

// now it is the time to parse let statement
// example let x = 5@
type LetStatement struct{
	Token token.Token // token.LET
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode(){

}

func (ls *LetStatement) TokenLiteral() string{
	return ls.Token.Literal // "LET"
}

type Identifier struct{
	Token token.Token  // token.IDENT
	Value string
}

func (i *Identifier) expressionNode(){

} 

func (i *Identifier) TokenLiteral() string{
	return i.Token.Literal // "IDENT"
}
