package must_test

import (
	"fmt"

	"github.com/xiegeo/must"
)

func ExamplePanic() {
	// must.Panic caches a panic and returns the thrown value
	fmt.Println(must.Panic(func() { panic("foo") }))

	// must.Panic will panic if function did not panic
	fmt.Println(must.Panic(func() { must.Panic(func() {}) }))

	// use must.Recover instead if function might not panic
	fmt.Println(must.Recover(func() {}))

	// output:
	// foo
	// function expected to panic but did not
	// <nil>
}
