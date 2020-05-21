package object

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"strings"

	"github.com/musale/monkey/pkg/ast"
)

const (
	// IntegerObj name
	IntegerObj = "INTEGER"
	// BooleanObj name
	BooleanObj = "BOOLEAN"
	// NullObj name
	NullObj = "NULL"
	// ReturnValueObj name
	ReturnValueObj = "RETURN_VALUE"
	// ErrorObj name
	ErrorObj = "ERROR"
	// FunctionObj name
	FunctionObj = "FUNCTION"
	//StringObj name
	StringObj = "STRING"
	//BuiltinObj name
	BuiltinObj = "BUILTIN"
	// ArrayObj name
	ArrayObj = "ARRAY"
	// HashObj name
	HashObj = "HASH"
)

// BuiltinFunction defines an inbuilt function
type BuiltinFunction func(args ...Object) Object

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
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

// Type returns the integer object name
func (i *Integer) Type() Type { return IntegerObj }

// Boolean data type
type Boolean struct {
	Value bool
}

// Inspect returns the Boolean value string format
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

// Type returns the Boolean object name
func (b *Boolean) Type() Type { return BooleanObj }

// Null data type
type Null struct{}

// Inspect returns the Null value string format
func (b *Null) Inspect() string { return "<null>" }

// Type returns the Null object name
func (b *Null) Type() Type { return NullObj }

// ReturnValue data type
type ReturnValue struct {
	Value Object
}

// Inspect returns the Return value string format
func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

// Type returns the Return object name
func (rv *ReturnValue) Type() Type { return ReturnValueObj }

// Error data type
type Error struct {
	Message string
}

// Inspect returns the Return value string format
func (e *Error) Inspect() string {
	return fmt.Sprintf("%s: %s", ErrorObj, e.Message)
}

// Type returns the Return object name
func (e *Error) Type() Type { return ErrorObj }

// Function data type
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

// Inspect returns the Return value string format
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n\t")
	out.WriteString(f.Body.String())
	out.WriteString("\n\t}")

	return out.String()
}

// Type returns the Return object name
func (f *Function) Type() Type { return FunctionObj }

// String object
type String struct {
	Value string
}

// Type returns the String Object name
func (s *String) Type() Type { return StringObj }

// Inspect returns the string value
func (s *String) Inspect() string { return s.Value }

// Builtin object
type Builtin struct {
	Fn BuiltinFunction
}

// Type returns the String Object name
func (b *Builtin) Type() Type { return BuiltinObj }

// Inspect returns the string value
func (b *Builtin) Inspect() string { return "builtin function" }

// Array Obj
type Array struct {
	Elements []Object
}

// Type of array object
func (ao *Array) Type() Type { return ArrayObj }

// Inspect returns the array value
func (ao *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

// Hashable defines methods for elements that can
// be used as HashKeys
type Hashable interface {
	HashKey() HashKey
}

// HashKey ...
type HashKey struct {
	Type  Type
	Value uint64
}

// HashKey generator for a boolean key
func (b *Boolean) HashKey() HashKey {
	var value uint64

	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return HashKey{Type: b.Type(), Value: value}
}

// HashKey generator for an integer key
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

// HashKey generator for a string key
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))

	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

// HashPair is K:V
type HashPair struct {
	Key   Object
	Value Object
}

// Hash ...
type Hash struct {
	Pairs map[HashKey]HashPair
}

// Type returns the Hash type
func (h *Hash) Type() Type { return HashObj }

// Inspect returns the Hash representation
func (h *Hash) Inspect() string {
	var out bytes.Buffer

	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s",
			pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}
