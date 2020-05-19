package object

// NewEnclosedEnvironment creates a new Enclosed env
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// NewEnvironment creates a new Environment with a map
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

// Environment holds the values in global
type Environment struct {
	store map[string]Object
	outer *Environment
}

// Get fetches an object by name
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set creates a new Object in environment
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
