package lexer

import (
	"testing"
	"monkey/token"
)

type ExpectedToken struct {
	Type    token.TokenType
	Literal string
}

type TokenT testing.T

func (t *TokenT) fatalToken(description string, expected ExpectedToken, actual token.Token) {
	if expected.Type != actual.Type {
		t.Fatalf("%s - type wrong. Expected %q/%q, got %q/%q", description, expected.Type, expected.Literal, actual.Type, actual.Literal)
	}

	if expected.Literal != actual.Literal {
		t.Fatalf("%s - literal wrong. Expected %q/%q, got %q/%q", description, expected.Literal, expected.Literal, actual.Type, actual.Literal)
	}
}
