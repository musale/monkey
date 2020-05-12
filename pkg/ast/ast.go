package ast

import (
	"github.com/musale/monkey/pkg/token"
)

// Node is implemented by all nodes in the program
type Node interface {
	TokenLiteral() string
}

// Statement defines what declares a value
type Statement interface {
	Node
	statementNode()
}

// Expression defines what produces a value
type Expression interface {
	Node
	expressionNode()
}

// Program is the root of every AST
type Program struct {
	Statements []Statement
}

// TokenLiteral returns the string value of the token
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		s := p.Statements[0]
		return s.TokenLiteral()
	}
	return ""
}

// LetStatement expresses a let <identinfier> = <expression>
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral ...
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// Identifier expresses identifier binding
// in let <identinfier> = <expression>
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral ...
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// ReturnStatement describes return <expression>
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral of a return statement
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
