package must

// Panic expects function f to panic and return the recovered value.
// Panic will panic if the expected panic did not happen.
func Panic(f func(), debug ...any) (r any) {
	defer func() {
		r = recover()
		if r == nil {
			panicWith(debug, "function expected to panic but did not")
		}
	}()
	f()
	return
}

func panicWith(debug []any, last any) {
	if len(debug) > 0 {
		if last == nil {
			panic(debug)
		}
		panic(append(debug, last))
	}
	panic(last)
}

// Recover returns the recovered value, or nil if function f never panicked.
func Recover(f func()) (r any) {
	defer func() {
		r = recover()
	}()
	f()
	return
}
