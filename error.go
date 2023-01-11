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
// You can use this instead of ignoring errors that never happen by contract.
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

// Value3 panics on error, otherwise returns the first 3 values.
func Value3[T any, U any, V any](out1 T, out2 U, out3 V, err error) (T, U, V) {
	if err != nil {
		panic(err)
	}
	return out1, out2, out3
}

// V3 is short for Value3, which panics on error, otherwise returns the first 3 values.
func V3[T any, U any, V any](out1 T, out2 U, out3 V, err error) (T, U, V) {
	return Value3(out1, out2, out3, err)
}

// Value4 panics on error, otherwise returns the first 4 values.
func Value4[T any, U any, V any, W any](out1 T, out2 U, out3 V, out4 W, err error) (T, U, V, W) {
	if err != nil {
		panic(err)
	}
	return out1, out2, out3, out4
}

// V4 is short for Value4, which panics on error, otherwise returns the first 4 values.
func V4[T any, U any, V any, W any](out1 T, out2 U, out3 V, out4 W, err error) (T, U, V, W) {
	return Value4(out1, out2, out3, out4, err)
}
