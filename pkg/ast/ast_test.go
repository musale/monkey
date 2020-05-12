package ast

import (
	"testing"

	"github.com/musale/monkey/pkg/token"
)

func TestAst(t *testing.T) {
	t.Run("Test TokenLiteral", func(t *testing.T) {
		stmts := []Statement{}
		program := &Program{Statements: stmts}
		literal := program.TokenLiteral()
		if literal != "" {
			t.Errorf("Expected an empty literal but got %s", literal)
		}

		// add a let statement
		letToken := token.Token{Type: token.LET, Literal: "let"}
		program.Statements = append(program.Statements, &LetStatement{Token: letToken})
		literal = program.TokenLiteral()
		if literal != "let" {
			t.Errorf("Expected 'let' but got %s", literal)
		}

		// Create an identifier
		intToken := token.Token{Type: token.INT, Literal: "5"}
		identifier := &Identifier{Token: intToken}
		literal = identifier.TokenLiteral()
		if literal != "5" {
			t.Errorf("Expected '5' but got %s", literal)
		}
	})
}
