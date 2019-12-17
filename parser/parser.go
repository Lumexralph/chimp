// Package parser is the implementation of the
// program that converts the tokens from the lexers to AST.
package parser

import (
	"github.com/Lumexralph/chimp/ast"
	"github.com/Lumexralph/chimp/lexer"
	"github.com/Lumexralph/chimp/token"
)

type Parser struct {
	l            *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
}

// New - create a new parser
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read token tokens twice to set currentToken and peekToken
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram - starts the construction of the tokens as
// AST nodes
func (p *Parser) ParseProgram() *ast.Program {
	// root of the AST
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.currentTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

// construct the ast.Statement node,
// handle the let x = 5; statements
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.currentToken}

	// x
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{
		Token: p.currentToken,
		Value: p.currentToken.Literal,
	}

	// =
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: skipping the expressions until we
	// get to a semicolon
	for !p.currentTokenIs(token.SEMICOLON) {
		// keep searching
		p.nextToken()
	}

	return stmt
}

func(p *Parser) currentTokenIs(t token.Type) bool {
	return p.currentToken.Type == t
}

func(p *Parser) peekTokenIs(t token.Type) bool {
	return p.peekToken.Type == t
}

func(p *Parser) expectPeek(t token.Type) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	return false
}