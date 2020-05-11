package token

const (
	// ILLEGAL - an unidentifiable token
	ILLEGAL = "ILLEGAL"
	// EOF - end of file
	EOF = "EOF"

	// IDENT - an identifier
	IDENT = "IDENT" // add, foobar, x, y, ...
	// INT - a literal
	INT = "INT" // 1343456

	// ASSIGN - operator
	ASSIGN = "="
	// PLUS - addition operator
	PLUS = "+"

	// COMMA - delimiter
	COMMA = ","
	// SEMICOLON - delimiter
	SEMICOLON = ";"

	// LPAREN - left paranthesis
	LPAREN = "("
	// RPAREN - left paranthesis
	RPAREN = ")"
	// LBRACE - left brace
	LBRACE = "{"
	// RBRACE - left brace
	RBRACE = "}"

	// FUNCTION - keyword for declaring a funtion
	FUNCTION = "FUNCTION"
	// LET - keyword for assigning a variable
	LET = "LET"
)

// TokenType - is the name of a given token
type TokenType string

// Token - holds the name of the token and the literal value
type Token struct {
	Type    TokenType
	Literal string
}
