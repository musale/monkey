package environment

import (
	"github.com/musale/monkey/pkg/object"
)

// NewEnvironment creates a new Environment with a map
func NewEnvironment() *Environment {
	s := make(map[string]object.Object)
	return &Environment{store: s}
}

// Environment holds the values in global
type Environment struct {
	store map[string]object.Object
}

// Get fetches an object by name
func (e *Environment) Get(name string) (object.Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

// Set creates a new Object in environment
func (e *Environment) Set(name string, val object.Object) object.Object {
	e.store[name] = val
	return val
}
