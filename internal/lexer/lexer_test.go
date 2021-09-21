package lexer_test

import (
	"lioslang/internal/lexer"
	"lioslang/internal/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
	def ten = 10;
	def five = 5;
	def add = fn(x, y){
		x+y;
	};
	def result = add(five, ten);
	`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.DEF, "def"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.DEF, "def"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
	}
	l := lexer.New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
