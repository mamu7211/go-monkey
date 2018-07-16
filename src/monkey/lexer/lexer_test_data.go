package lexer

import "monkey/token"


// Equals and Not Equals

const InputEqualsNotEquals = `
let a = true == true;
let b = false != true;
`

var InputEqualsNotEqualsTokens = []ExpectedToken {
	{token.LET,"let"},
	{token.IDENT,"a"},
	{token.ASSIGN,"="},
	{token.TRUE,"true"},
	{token.EQ,"=="},
	{token.TRUE,"true"},
	{token.SEMICOLON,";"},
	{token.LET,"let"},
	{token.IDENT,"b"},
	{token.ASSIGN,"="},
	{token.FALSE,"false"},
	{token.NOT_EQ,"!="},
	{token.TRUE,"true"},
	{token.SEMICOLON,";"},
}

// Long code Snippet
const InputLetAndAddCode = `let five = 5;
let ten = 10;

let add = fn(x,y) {
	x + y
};

let result = add(five, ten);

if (5 < 10) {
	return true;
} else {
	return false;
}
`

var InputLetAndAddTokens = []ExpectedToken{
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
	{token.RPAREN, ")"},
	{token.SEMICOLON, ";"},
	{token.IF, "if"},
	{token.LPAREN, "("},
	{token.INT, "5"},
	{token.LT, "<"},
	{token.INT, "10"},
	{token.RPAREN, ")"},
	{token.LBRACE, "{"},
	{token.RETURN, "return"},
	{token.TRUE, "true"},
	{token.SEMICOLON, ";"},
	{token.RBRACE, "}"},
	{token.ELSE, "else"},
	{token.LBRACE, "{"},
	{token.RETURN, "return"},
	{token.FALSE, "false"},
	{token.SEMICOLON, ";"},
	{token.RBRACE, "}"},

	{token.EOF, ""},
}
