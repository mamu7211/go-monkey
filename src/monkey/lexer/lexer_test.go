package lexer

import (
	"testing"
	"strconv"
	"monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `=+-(){},;`
	tokenAssertions := (*TokenT)(t)
	tests := []ExpectedToken{
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

	for i, expected := range tests {
		actual := line.NextToken()
		tokenAssertions.fatalToken("Test "+strconv.Itoa(i), expected, actual)
	}
}

func TestNextTokenWithCode(t *testing.T) {

	tokenAssertions := (*TokenT)(t)

	line := New(InputLetAndAddCode)

	for i, tt := range InputLetAndAddTokens {
		tok := line.NextToken()
		tokenAssertions.fatalToken("Test " + strconv.Itoa(i),tt,tok)
	}
}
