package parser_test

import (
	"lioslang/internal/ast"
	"lioslang/internal/lexer"
	"lioslang/internal/parser"
	"testing"
)

func TestLetStatement(t *testing.T) {
	input := `
	def x = 5;
	def y = 5;
	def xy = 10;
	`

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returning nil!")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("Program not contains 3 statment go=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"xy"},
	}
	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testStatement(t *testing.T, s ast.Statment, name string) bool {
	if s.TokenLiteral() != "def" {
		t.Errorf("s.TokenLiteral not def got=%q", s.TokenLiteral())
		return false
	}
	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}
	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}
	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, letStmt.Name)
		return false
	}
	return true
}
