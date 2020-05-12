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
		intIdent := &Identifier{Token: intToken}
		literal = intIdent.TokenLiteral()
		if literal != "5" {
			t.Errorf("Expected '5' but got %s", literal)
		}
		// Create a return
		returnToken := token.Token{Type: token.RETURN, Literal: "return"}
		retIdent := &Identifier{Token: returnToken}
		literal = retIdent.TokenLiteral()
		if literal != "return" {
			t.Errorf("Expected 'return' but got %s", literal)
		}
	})
}
