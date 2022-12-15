package must

// TBSubset is a subset of testing.TB, TBSubset may add or remove public methods that are defined in testing.TB
// without updating major version. TBSubset should only used to accept testing.TB without depending on testing in this package.
// It also severs as documentation to what features of testing.TB must.*T# functions depends on.
type TBSubset interface {
	// The method comments are copied from *testing.T

	// Helper marks the calling function as a test helper function.
	// When printing file and line information, that function will be skipped.
	// Helper may be called simultaneously from multiple goroutines.
	Helper()
	// Fatal is equivalent to Log followed by FailNow.
	//
	// Log formats its arguments using default formatting, analogous to Println, and records the text in the error log. For tests, the text will be printed only if the test fails or the -test.v flag is set. For benchmarks, the text is always printed to avoid having performance depend on the value of the -test.v flag.
	//
	// FailNow marks the function as having failed and stops its execution by calling runtime.Goexit (which then runs all deferred calls in the current goroutine). Execution will continue at the next test or benchmark. FailNow must be called from the goroutine running the test or benchmark function, not from other goroutines created during the test. Calling FailNow does not stop those other goroutines.
	Fatal(args ...any)
}

// NoErrorT and alike *T# functions provide test fail instead of panic version of must.
//
// NoErrorT's signature follow other test fail functions for consistency.
func NoErrorT(err error) func(TBSubset) {
	return func(t TBSubset) {
		if err != nil {
			t.Helper()
			t.Fatal(err)
		}
	}
}

// ValueT and alike *T# functions provide test fail instead of panic version of must.
//
// The API has accepts TBSubset in a separate function call to allow function chaining,
// It can not go first because of limitations in go type inference. It can not be encapsulated
// in an interface because interface methods can not introduce generic parameters.
func ValueT[T any](out T, err error) func(TBSubset) T {
	return func(t TBSubset) T {
		if err != nil {
			t.Helper()
			t.Fatal(err)
		}
		return out
	}
}

// VT is short for ValueT.
//
// ValueT and alike *T# functions provide test fail instead of panic version of must.
func VT[T any](out T, err error) func(TBSubset) T {
	return ValueT(out, err)
}

// ValueT2 and alike *T# functions provide test fail instead of panic version of must.
func ValueT2[T any, U any](out1 T, out2 U, err error) func(TBSubset) (T, U) {
	return func(t TBSubset) (T, U) {
		if err != nil {
			t.Helper()
			t.Fatal(err)
		}
		return out1, out2
	}
}

// VT2 is short for ValueT2.
//
// ValueT2 and alike *T# functions provide test fail instead of panic version of must.
func VT2[T any, U any](out1 T, out2 U, err error) func(TBSubset) (T, U) {
	return ValueT2(out1, out2, err)
}
