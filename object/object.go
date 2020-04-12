package object

import (
	"bytes"
	"fmt"
	"monkey/ast"
	"strings"
)

const (
	// INTEGEROBJ represents an integer object
	INTEGEROBJ = "INTEGER"
	// BOOLEANOBJ represents an boolean object
	BOOLEANOBJ = "BOOLEAN"
	// NULLOBJ represents an nil object
	NULLOBJ = "NULL"
	// NANOBJ represents an nil object
	NANOBJ = "NAN"
	// RETURNVALUEOBJ represents a return object
	RETURNVALUEOBJ = "RETURN_VALUE"
	// ERROROBJ represents an error object
	ERROROBJ = "ERROR"
	// FUNCTIONOBJ represents a function object
	FUNCTIONOBJ = "FUNCTION"
)

// Type represents the type of an object
type Type string

// Object wraps every value of the language
type Object interface {
	Type() Type
	Inspect() string
}

// Integer the int type
type Integer struct {
	Value int64
}

// Type returns the object type of this value
func (i *Integer) Type() Type { return INTEGEROBJ }

// Inspect returns a readable string of the integer value
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

// Boolean the bool type
type Boolean struct {
	Value bool
}

// Type returns the object type of this value
func (i *Boolean) Type() Type { return BOOLEANOBJ }

// Inspect returns a readable string of the boolean value
func (i *Boolean) Inspect() string { return fmt.Sprintf("%t", i.Value) }

// Null the nil type
type Null struct{}

// Type returns the object type of this value
func (i *Null) Type() Type { return NULLOBJ }

// Inspect returns a readable string of the Null value
func (i *Null) Inspect() string { return "null" }

// Nan represents not-a-number
type Nan struct{}

// Type returns the object type of this value
func (i *Nan) Type() Type { return NANOBJ }

// Inspect returns a readable string of the Nan value
func (i *Nan) Inspect() string { return "NAN" }

// ReturnValue represents a return value
type ReturnValue struct {
	Value Object
}

// Type returns the object type of this value
func (rv *ReturnValue) Type() Type { return RETURNVALUEOBJ }

// Inspect returns a readable string of the return value
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }

// Error represents an error in our program
type Error struct {
	Message string
}

// Type returns the object type of this value
func (e *Error) Type() Type { return ERROROBJ }

// Inspect returns a readable string of the error
func (e *Error) Inspect() string { return "ERROR: " + e.Message }

// Function represents a function in our program
type Function struct {
	Parameters ast.Identifiers
	Body       *ast.BlockStatement
	Env        *Environment
}

// Type returns the object type of this value
func (f *Function) Type() Type { return FUNCTIONOBJ }

// Inspect returns a readable string of the function
func (f *Function) Inspect() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}
