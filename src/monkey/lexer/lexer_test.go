package lexer

import (
	"testing"
	"monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `=+-(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	line := New(input)

	for i, tt := range tests {
		tok := line.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("Test tests[%d] - tokentype wrong. Expected '%q', got '%q'", i+1, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("Test tests[%d] - literal wrong. Expected '%q', got '%q", i+1, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextTokenWithCode(t *testing.T) {
	input := `
let five = 5
let ten = 10

let add = fn(x,y) {
	x + y
};

let result = add(five, ten)
`
	tests := []struct {
		expectedType token.TokenType,
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RBRACE, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, "let"},
	}
}
