package parser

import (
	"github.com/0x1bvaidi/go_interpreter/ast"
	"github.com/0x1bvaidi/go_interpreter/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x 5@
	let  = 10@
	let  383838@
	`

	l := lexer.New(input)
	p := New(l) // New() from parser package, same package so no need for parser.New()

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does ot contain 3 statements. got %d", len(program.Statements))
	}

	// slice of structs with each having expectedIdentifier field of type string
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser){
	errors := p.Errors()
	if len(errors) == 0 {
		return 
	}

	t.Errorf("parser has %d errors", len(errors))
	for _,msg := range errors{
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool{
	if s.TokenLiteral() != "let"{
		t.Errorf("s.TokenLiteral not 'let'. got %q", s.TokenLiteral())
		return false
	}

		letStmt, ok := s.(*ast.LetStatement) // LetStatement a struct in ast
		if !ok{
			t.Errorf("s not *ast.LetStatement. got %T", s)
			return false
		}

		if letStmt.Name.Value != name{
			// letStmt.Name.Value:= Name field of LetStatement is of type *Identifier
			// Identifier is a struct in ast with filed Value of type string
			// so Name.Value is a string
			// letStmt is an instance of LetStatement
			// and letStmt.Name.Value is a string henceforth.
			t.Errorf("letStmt.Name.Value is not %s. got %s", name, letStmt.Name.Value)
			return false
		}

		if letStmt.Name.TokenLiteral() !=  name{
			t.Errorf("s.Name is not %s. got %s", name, letStmt.Name)
			return false
		}

		return true	
}
