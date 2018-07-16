package lexer

import (
	"testing"
	"strconv"
	"monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `=+-*/(){},;!<>`
	tokenAssertions := (*TokenT)(t)
	tests := []ExpectedToken{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.ASTERISK, "*"},
		{token.SLASH, "/"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.LT, "<"},
		{token.GT, ">"},
		{token.EOF, ""},
	}

	line := NewLexer(input)

	for i, expected := range tests {
		actual := line.NextToken()
		tokenAssertions.fatalToken("Test "+strconv.Itoa(i), expected, actual)
	}
}

func TestNextTokenWithCode(t *testing.T) {

	tokenAssertions := (*TokenT)(t)

	lexer := NewLexer(InputLetAndAddCode)

	for i, tt := range InputLetAndAddTokens {
		tok := lexer.NextToken()
		tokenAssertions.fatalToken("Test "+strconv.Itoa(i), tt, tok)
	}
}

func TestInputEqualsNotEquals(t *testing.T) {
	tokenAssertions := (*TokenT)(t)
	lexer := NewLexer(InputEqualsNotEquals)
	for i, tt := range InputEqualsNotEqualsTokens {
		tok := lexer.NextToken()
		tokenAssertions.fatalToken("Test "+strconv.Itoa(i), tt, tok)
	}
}
