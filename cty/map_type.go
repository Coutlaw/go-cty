package cty

import (
	"fmt"
)

// TypeList instances represent specific list types. Each distinct ElementType
// creates a distinct, non-equal list type.
type typeMap struct {
	typeImplSigil
	elementType Type
}

// Map creates a map type with the given element Type.
//
// Map types are CollectionType implementations.
func Map(elem Type) Type {
	return Type{
		typeMap{
			elementType: elem,
		},
	}
}

// Equals returns true if the other Type is a map whose element type is
// equal to that of the receiver.
func (t typeMap) Equals(other Type) bool {
	ot, isMap := other.typeImpl.(typeMap)
	if !isMap {
		return false
	}

	return t.elementType.Equals(ot.elementType)
}

func (t typeMap) FriendlyName() string {
	return "map of " + t.elementType.FriendlyName()
}

func (t typeMap) ElementType() Type {
	return t.elementType
}

func (t typeMap) GoString() string {
	return fmt.Sprintf("cty.Map(%#v)", t.elementType)
}

// IsMapType returns true if the given type is a list type, regardless of its
// element type.
func (t Type) IsMapType() bool {
	_, ok := t.typeImpl.(typeMap)
	return ok
}

// MapElementType is a convenience method that checks if the given type is
// a map type, returning a pointer to its element type if so and nil
// otherwise. This is intended to allow convenient conditional branches,
// like so:
//
//     if et := t.MapElementType(); et != nil {
//         // Do something with *et
//     }
func (t Type) MapElementType() *Type {
	if lt, ok := t.typeImpl.(typeMap); ok {
		return &lt.elementType
	}
	return nil
}