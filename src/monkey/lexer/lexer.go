package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (lexer *Lexer) NextToken() token.Token {
	var result token.Token
	lexer.skipWhiteSpace()
	switch lexer.ch {
	case '=':
		if lexer.peekChar() == '=' {
			ch := lexer.ch
			lexer.readChar()
			result = token.Token{
				Type:    token.EQ,
				Literal: string(ch) + string(lexer.ch),
			}
		} else {
			result = newToken(token.ASSIGN, lexer.ch)
		}
	case '+':
		result = newToken(token.PLUS, lexer.ch)
	case '-':
		result = newToken(token.MINUS, lexer.ch)
	case '*':
		result = newToken(token.ASTERISK, lexer.ch)
	case '/':
		result = newToken(token.SLASH, lexer.ch)
	case '(':
		result = newToken(token.LPAREN, lexer.ch)
	case ')':
		result = newToken(token.RPAREN, lexer.ch)
	case '{':
		result = newToken(token.LBRACE, lexer.ch)
	case '}':
		result = newToken(token.RBRACE, lexer.ch)
	case ',':
		result = newToken(token.COMMA, lexer.ch)
	case ';':
		result = newToken(token.SEMICOLON, lexer.ch)
	case '<':
		result = newToken(token.LT, lexer.ch)
	case '>':
		result = newToken(token.GT, lexer.ch)
	case '!':
		if lexer.peekChar() == '=' {
			ch := lexer.ch
			lexer.readChar()
			result = token.Token{
				Type:    token.NOT_EQ,
				Literal: string(ch) + string(lexer.ch),
			}
		} else {
			result = newToken(token.BANG, lexer.ch)
		}
	case 0:
		result.Type = token.EOF
		result.Literal = ""
	default:
		if isLetter(lexer.ch) {
			result.Literal = lexer.readIdentifier()
			result.Type = token.LookupIdentifier(result.Literal)
			return result
		} else if isDigit(lexer.ch) {
			result.Literal = lexer.readNumber()
			result.Type = token.INT
			return result
		} else {
			result = newToken(token.ILLEGAL, lexer.ch)
		}
	}

	lexer.readChar()
	return result
}

func (lexer *Lexer) readIdentifier() string {
	position := lexer.position
	for isLetter(lexer.ch) {
		lexer.readChar()
	}
	return lexer.input[position:lexer.position]
}

func (lexer *Lexer) readNumber() string {
	position := lexer.position
	for isDigit(lexer.ch) {
		lexer.readChar()
	}
	return lexer.input[position:lexer.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (lexer *Lexer) readChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.ch = 0
	} else {
		lexer.ch = lexer.input[lexer.readPosition]
	}
	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func (lexer *Lexer) peekChar() byte {
	if lexer.readPosition >= len(lexer.input) {
		return 0
	} else {
		return lexer.input[lexer.readPosition]
	}
}

func (lexer *Lexer) skipWhiteSpace() {
	for lexer.ch == ' ' || lexer.ch == '\t' || lexer.ch == '\n' || lexer.ch == '\r' {
		lexer.readChar()
	}
}
