package parser_test

import (
	"lioslang/internal/ast"
	"lioslang/internal/lexer"
	"lioslang/internal/parser"
	"testing"
)

func TestDefStatement(t *testing.T) {
	input := `
	def x=  5;
	def  y= 5;
	def xy=10;
	`

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)
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
	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.returnStatement. got=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q", returnStmt.TokenLiteral())
		}
	}
}

func TestReturnStatement(t *testing.T) {
	input := `
	return 5;
	return 10;
	`
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)
	if len(program.Statements) != 2 {
		t.Fatalf("Program not contains 3 statment go=%d", len(program.Statements))
	}
}

func checkParseErrors(t *testing.T, p *parser.Parser) {
	errors := p.Errrors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func testStatement(t *testing.T, s ast.Statment, name string) bool {
	if s.TokenLiteral() != "def" {
		t.Errorf("s.TokenLiteral not def got=%q", s.TokenLiteral())
		return false
	}
	letStmt, ok := s.(*ast.DefStatement)
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
