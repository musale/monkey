package token

import "testing"

func TestLookupIdent(t *testing.T) {
	inputs := []Token{
		{Type: LET, Literal: "let"},
		{Type: IF, Literal: "if"},
		{Type: ELSE, Literal: "else"},
	}

	for _, input := range inputs {
		if LookupIdent(input.Literal) == IDENT {
			t.Errorf("Expected Type %s but got IDENT", input.Type)
		}
	}
	if eq := LookupIdent("EQ"); eq != IDENT {
		t.Errorf("Expected Type EQ but got %s", eq)
	}
}
