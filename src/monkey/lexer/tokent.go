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
	t.Fatalf("%s - type wrong. Expected Type=%q, got Type=%q, Literal=%q", description, expected.Type, actual.Type, actual.Literal)
	t.Fatalf("%s - literal wrong. Expected Type=%q, got Type=%q, Literal=%q", description, expected.Literal, actual.Type, actual.Literal)
}
