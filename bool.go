package must

// True panics on false.
func True(b bool, debug ...any) {
	if !b {
		panicWith(debug, "must be true")
	}
}

// False panics on true.
func False(b bool, debug ...any) {
	if b {
		panicWith(debug, "must be false")
	}
}
