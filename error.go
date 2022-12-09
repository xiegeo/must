package must

// NoError panics on error.
//
// Use this because test coverage.
func NoError(err error) {
	if err != nil {
		panic(err)
	}
}

// Value panics on error, otherwise returns the first value.
//
// Use this instead of writing a Must version of your function.
func Value[T any](out T, err error) T {
	if err != nil {
		panic(err)
	}
	return out
}

// V is short for Value, which panics on error, otherwise returns the first value.
func V[T any](out T, err error) T {
	return Value(out, err)
}

// Value2 panics on error, otherwise returns the first 2 values.
func Value2[T any, U any](out1 T, out2 U, err error) (T, U) {
	if err != nil {
		panic(err)
	}
	return out1, out2
}

// V2 is short for Value2, which panics on error, otherwise returns the first 2 values.
func V2[T any, U any](out1 T, out2 U, err error) (T, U) {
	return Value2(out1, out2, err)
}
