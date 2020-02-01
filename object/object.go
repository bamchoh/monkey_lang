package object

import "fmt"

// ObjectType is object type.
type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ    = "NULL"
)

// Object is object in this lang
type Object interface {
	Type() ObjectType
	Inspect() string
}

// Integer is integer object
type Integer struct {
	Value int64
}

// Type returns integer type
func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

// Inspect returns integer value
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

// Boolean represents bool object
type Boolean struct {
	Value bool
}

// Type returns boolean type
func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

// Inspect returns boolean value
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

// Null resresents null object
type Null struct{}

// Type returns null type
func (n *Null) Type() ObjectType {
	return NULL_OBJ
}

// Inspect returns null value
func (n *Null) Inspect() string {
	return "null"
}
