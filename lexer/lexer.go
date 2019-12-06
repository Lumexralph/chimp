// Package lexer implements the lexer, it will do the lexical analysis
// of the source code given to it as input, tokenize it and return
// tokens that will be sent to the parser for further processing.
// lexer supports ASCII characters instead of the full Unicode range.
package lexer

import "github.com/Lumexralph/chimp/token"

/*
TODO

- add filename to the token
- add line number and column number
- create full support for unicode UTF-8 and emojis
- use isDigit to create support for floats,  hexadecimal, Octal notation
*/

// Lexer - the attributes of our lexer
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	char         byte // current char under examination
}

// New - creates a new Lexer.
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // start up the lexer.
	return l
}

// readChar - give us the next character and advance our
// position in the input string.
func (l *Lexer) readChar() {
	// either we haven't read anything or it is the end of file (EOF).
	if l.readPosition >= len(l.input) {
		// sets char to 0, which is the ASCII code for the "NUL".
		l.char = 0
	} else {
		l.char = l.input[l.readPosition]
	}

	// set the current position of the character.
	l.position = l.readPosition

	// proceed to the next character.
	l.readPosition++
}

// NextToken - return a token depending on which character it is
// advance to the next character in the input
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	// whitespace only acts as a separator of tokens and
	// doesn’t have meaning in the language
	// so we need to skip over it entirely
	l.skipWhiteSpace()

	switch l.char {
	case '=':
		tok = newToken(token.ASSIGN, l.char)
	case ';':
		tok = newToken(token.SEMICOLON, l.char)
	case '(':
		tok = newToken(token.LPAREN, l.char)
	case ')':
		tok = newToken(token.RPAREN, l.char)
	case '+':
		tok = newToken(token.PLUS, l.char)
	case ',':
		tok = newToken(token.COMMA, l.char)
	case '{':
		tok = newToken(token.LBRACE, l.char)
	case '}':
		tok = newToken(token.RBRACE, l.char)
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	default:
		if isLetter(l.char) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		}
		if isDigit(l.char) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		}
		tok = newToken(token.ILLEGAL, l.char)
	}

	l.readChar()

	return tok
}

func newToken(t token.Type, char byte) token.Token {
	return token.Token{
		Type:    t,
		Literal: string(char),
	}
}

// readIdentifier - reads in an identifier and advances our
// lexer’s positions until it encounters a non-letter-character.
func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.char) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.char) {
		l.readChar()
	}
	return l.input[position:l.position]
}

/*
This little helper function is found in a lot of parsers.
Sometimes it’s called eatWhitespace and sometimes consumeWhitespace
and sometimes something entirely different. Which characters these
functions actually skip depends on the language being lexed.
*/
func (l *Lexer) skipWhiteSpace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}
