package parser

import (
	"testing"

	"github.com/musale/monkey/pkg/ast"
	"github.com/musale/monkey/pkg/lexer"
)

func TestParse(t *testing.T) {
	t.Run("Test let Statements", func(t *testing.T) {
		input := `
		let x = 5;
		let y = 10;
		let foo = 12345;
		`
		l := lexer.NewLexer(input)
		p := NewParser(l)

		program := p.ParseProgram()
		checkParserErrors(t, p)
		if program == nil {
			t.Fatalf("Program is nil!")
		}

		if len(program.Statements) != 3 {
			t.Fatalf("Expected 3 statements but got %d", len(program.Statements))
		}
		tests := []struct {
			expectedIdentifier string
		}{
			{"x"},
			{"y"},
			{"foo"},
		}

		for i, tt := range tests {
			stmt := program.Statements[i]
			if !testLetStatement(t, stmt, tt.expectedIdentifier) {
				return
			}
		}
	})
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func testLetStatement(t *testing.T, stmt ast.Statement, name string) bool {
	t.Helper()
	if stmt.TokenLiteral() != "let" {
		t.Errorf("stmt.TokenLiteral() not 'let' got %q", stmt.TokenLiteral())
		return false
	}
	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf("stmt not *ast.LetStatement. got=%T", stmt)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s'. got=%s",
			name, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}
