package lexer

import (
	"testing"

	"fmt"

	"github.com/Lumexralph/chimp/token"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
		let ten = 10;
		let add = fn(x, y) {
			x + y;
		};
		let result = add(five, ten);`

	cases := []struct {
		wantedType    token.Type
		wantedLiteral string
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
		{token.SEMICOLON, ";"},
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
		{token.EOF, ""},
	}

	l := New(input)

	for i, tc := range cases {
		t.Run(fmt.Sprint("test for valid token"), func(t *testing.T) {
			tok := l.NextToken()

			if tok.Type != tc.wantedType {
				t.Errorf("l.NextToken(input[%d]) - wrong tokentype, got %v; want %v", i, tok.Type, tc.wantedType)
			}

			if tok.Literal != tc.wantedLiteral {
				t.Errorf("l.NextToken(input[%d]) - wrong literal type, got %v; want %v", i, tok.Literal, tc.wantedLiteral)
			}
		})
	}
}
