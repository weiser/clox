package value

import "fmt"

type ValueType int

const (
	VAL_BOOL ValueType = iota
	VAL_NIL
	VAL_NUMBER
)

type Value struct {
	Type ValueType
	// on pg 329 this is defined as a union, but go doesn't have unions.
	// therefore I'm just going to create a data member per type we will support
	Vbool   bool
	Vnumber float64
}

// a facsimile of BOOL_VAL(value) on pg 330
func BoolVal(b bool) Value {
	return Value{Type: VAL_BOOL, Vbool: b}
}

// a facsimile of NIL_VAL on pg 330
func BoolNil() Value {
	return Value{Type: VAL_NIL, Vnumber: 0}
}

// a facsimile of NUMBER_VAL(value) on pg 330
func NumberVal(f float64) Value {
	return Value{Type: VAL_NUMBER, Vnumber: f}
}

// a facsimile of AS_BOOL(value) on pg 330
func AsBool(v Value) bool {
	return v.Vbool
}

// a facsimile of AS_NUMBER(value) on pg 330
func AsNumber(v Value) float64 {
	return v.Vnumber
}

// a facsimile of IS_BOOL(value) on pg 330
func IsBool(v Value) bool {
	return v.Type == VAL_BOOL
}

// a facsimile of IS_NIL(value) on pg 330
func IsNil(v Value) bool {
	return v.Type == VAL_NIL
}

// a facsimile of IS_NUMBER(value) on pg 330
func IsNumber(v Value) bool {
	return v.Type == VAL_NUMBER
}

type ValueArray []Value

func PrintValue(v Value) {
	fmt.Printf("%f", AsNumber(v))
}
