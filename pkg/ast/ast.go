package ast

import (
	"bytes"
	"fmt"

	"github.com/musale/monkey/pkg/token"
)

// Node is implemented by all nodes in the program
type Node interface {
	TokenLiteral() string
	String() string
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

// String repr of a Program
func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
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

func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("%s ", ls.TokenLiteral()))
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")

	return out.String()
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

func (i *Identifier) String() string {
	return i.Value
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

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("%s ", rs.Token.Literal))
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

// ExpressionStatement is a statement with one Expression
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral of the ExpressionStatement Token
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()

	}
	return ""
}
