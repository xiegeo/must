package must

// NoError panics on error.
func NoError(err error) {
	if err != nil {
		panic(err)
	}
}

// Value panics on error, otherwise returns the first value.
func Value[T any](out T, err error) T {
	if err != nil {
		panic(err)
	}
	return out
}

// V is short for Value, which panics on error, otherwise returns the first value.
func V[T any](out T, err error) T {
	if err != nil {
		panic(err)
	}
	return out
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
	if err != nil {
		panic(err)
	}
	return out1, out2
}
