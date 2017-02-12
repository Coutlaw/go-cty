package cty

// primitiveType is the hidden implementation of the various primitive types
// that are exposed as variables in this package.
type primitiveType struct {
	typeImplSigil
	Kind primitiveTypeKind
}

type primitiveTypeKind byte

const (
	primitiveTypeBool   primitiveTypeKind = 'B'
	primitiveTypeNumber primitiveTypeKind = 'N'
	primitiveTypeString primitiveTypeKind = 'S'
)

func (t primitiveType) Equals(other Type) bool {
	if otherP, ok := other.typeImpl.(primitiveType); ok {
		return otherP.Kind == t.Kind
	}
	return false
}

func (t primitiveType) FriendlyName() string {
	switch t.Kind {
	case primitiveTypeBool:
		return "bool"
	case primitiveTypeNumber:
		return "number"
	case primitiveTypeString:
		return "string"
	default:
		// should never happen
		panic("invalid primitive type")
	}
}

func (t primitiveType) GoString() string {
	switch t.Kind {
	case primitiveTypeBool:
		return "cty.Bool"
	case primitiveTypeNumber:
		return "cty.Number"
	case primitiveTypeString:
		return "cty.String"
	default:
		// should never happen
		panic("invalid primitive type")
	}
}

// Number is the numeric type. Number values are arbitrary-precision
// decimal numbers, which can then be converted into Go's various numeric
// types only if they are in the appropriate range.
var Number = Type{
	primitiveType{Kind: primitiveTypeNumber},
}

// String is the string type. String values are sequences of unicode codepoints
// encoded internally as UTF-8.
var String = Type{
	primitiveType{Kind: primitiveTypeString},
}

// Bool is the boolean type. The two values of this type are True and False.
var Bool = Type{
	primitiveType{Kind: primitiveTypeBool},
}

// True is the truthy value of type Bool
var True = trueValue
var trueValue = Value{
	ty: Bool,
	v:  true,
}

// False is the falsey value of type Bool
var False = falseValue
var falseValue = Value{
	ty: Bool,
	v:  false,
}

// IsPrimitiveType returns true if and only if the reciever is a primitive
// type, which means it's either number, string, or bool. Any two primitive
// types can be safely compared for equality using the standard == operator
// without panic, which is not a guarantee that holds for all types. Primitive
// types can therefore also be used in switch statements.
func (t Type) IsPrimitiveType() bool {
	_, ok := t.typeImpl.(*primitiveType)
	return ok
}