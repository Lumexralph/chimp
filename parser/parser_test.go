package parser

import (
	"fmt"
	"github.com/Lumexralph/chimp/ast"
	"github.com/Lumexralph/chimp/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	cases := []struct {
		input              string
		expectedIdentifier string
		expectedValue      interface{}
	}{
		{"let x = 5;", "x", 5},
		{"let y = true;", "y", true},
		{"let cost = y;", "cost", "y"},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("with valid identifier %q", tc.expectedIdentifier), func(t *testing.T) {
			l := lexer.New(tc.input)
			p := New(l)

			program := p.ParseProgram()
			if program == nil {
				t.Fatal("ParseProgram() returned nil")
			}

			if len(program.Statements) != 1 {
				t.Fatalf("program.Statements does not contain 1 statement. got=%d; want=%d", len(program.Statements), 1)
			}

			stmt := program.Statements[0]
			if !testLetStatement(t, stmt, tc.expectedIdentifier) {
				return
			}

			// val := stmt.(*ast.LetStatement).Value
			// if !testLetStatement(t, val, tc.expectedValue) {

			// }
		})
	}
}

func testLetStatement(t *testing.T, stmt ast.Statement, name string) bool {
	if stmt.TokenLiteral() != "let" {
		t.Errorf("TokenLiteral() returned invalid token got=%q; want=%q", stmt.TokenLiteral(), "let")
		return false
	}

	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf("stmt does not have the *ast.LetStatement interface. got=%T", stmt)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value is not %q. got=%q; want=%q", letStmt.Name.Value, letStmt.Name.Value, name)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() is not %q. got=%q; want=%q", letStmt.Name.TokenLiteral(), letStmt.Name.TokenLiteral(), name)
		return false
	}

	return true
}

// func testLiteralExpressions(t *testing.T, exp ast.Expression, expected interface{}) bool {
// 	switch v := expected.(type){
// 	case int:

// 	}
// }