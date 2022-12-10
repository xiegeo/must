package must

import (
	"fmt"
	"reflect"
)

type allowAll struct{}

// Any implements a Checker that always returns true.
var Any = allowAll{}

const anyDebugMessage = "Any allows all values"

func (allowAll) MustBe(v any) bool {
	return true
}

func (allowAll) Debug(v any) any {
	return anyDebugMessage
}

// Checker can be used by B# functions to check if a value is allowed.
type Checker interface {
	// MustBe returns true iff value is allowed.
	MustBe(v any) bool
	// Debug is used to explain why the value is not allowed.
	Debug(v any) any
}

// B1 asserts on one value.
//
// v is the value to be checked. If check is passed v is returned, otherwise panics.
//
// b checks v, there are a few cases:
//
//   - A Checker (such as must.Any for all values); or
//   - a literal value to test for equality using `==`,
//     if types are the same and literal is comparable; or
//   - a nil value, which will match nil of any type.
//   - All other cases panic, room for future expansion.
//
// v and b are separated to two function calls for clarity,
// v has to go first so the compiler can auto type it.
//
// debug does not participate in functionality. If the assert need a line comment,
// then use it as the debug message is suggested.
func B1[V any](v V) func(b any, debug ...any) V {
	return func(b any, debug ...any) V {
		checker, ok := b.(Checker)
		if ok {
			if !checker.MustBe(v) {
				panicWith(debug, checker.Debug(v))
			}
			return v
		}
		if b == nil {
			value := reflect.ValueOf(v)
			if value.IsValid() && !value.IsNil() {
				panicWith(append(debug, v), "value must be nil")
			}
			return v
		}
		bAsV, ok := b.(V)
		if ok {
			var equal bool
			r := Recover(func() { equal = any(bAsV) == any(v) })
			if r != nil {
				debug = append(debug, fmt.Sprintf("type(%T) is not comparable", bAsV), r)
			}
			if !equal {
				panicWith(append(debug, v, b), "values must equal")
			}
			return v
		}
		panicWith(debug, fmt.Sprintf("must use Checker interface or the same comparable type as input. expected %T got %T", v, b))
		return v // unreachable
	}
}

// B2 is 2 valued version of B1. To only Check one value, use must.Any on the other value to skip it.
func B2[V1 any, V2 any](v1 V1, v2 V2) func(b1, b2 any, debug ...any) (V1, V2) {
	return func(b1, b2 any, debug ...any) (V1, V2) {
		B1(v1)(b1, append(debug, "#1")...)
		B1(v2)(b2, append(debug, "#2")...)
		return v1, v2
	}
}

// B3 is 3 valued version of B1. To only Check some value, use must.Any on the other values to skip it.
func B3[V1 any, V2 any, V3 any](v1 V1, v2 V2, v3 V3) func(b1, b2, b3 any, debug ...any) (V1, V2, V3) {
	return func(b1, b2, b3 any, debug ...any) (V1, V2, V3) {
		B1(v1)(b1, append(debug, "#1")...)
		B1(v2)(b2, append(debug, "#2")...)
		B1(v3)(b3, append(debug, "#3")...)
		return v1, v2, v3
	}
}
