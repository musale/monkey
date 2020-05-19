package object

import "fmt"

const (
	// IntegerObj name
	IntegerObj = "INTEGER"
	// BooleanObj name
	BooleanObj = "BOOLEAN"
	// NullObj name
	NullObj = "NULL"
)

// Type represents every value during AST evaluation
type Type string

//Object interface
type Object interface {
	Type() Type
	Inspect() string
}

// Integer data type
type Integer struct {
	Value int64
}

// Inspect returns the integer value string format
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

// Type returns the integer object name
func (i *Integer) Type() Type { return IntegerObj }

// Boolean data type
type Boolean struct {
	Value bool
}

// Inspect returns the Boolean value string format
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

// Type returns the Boolean object name
func (b *Boolean) Type() Type { return BooleanObj }

// Null data type
type Null struct{}

// Inspect returns the Null value string format
func (b *Null) Inspect() string { return "<null>" }

// Type returns the Null object name
func (b *Null) Type() Type { return NullObj }
