package must_test

import (
	"fmt"

	"github.com/xiegeo/must"
)

func ExampleTrue() {
	trueFunc := func() bool {
		return true
	}

	must.True(trueFunc(), "a no-op")
	fmt.Printf("True:%#v\n", must.Panic(func() {
		must.True(false, "this will panic")
	}))

	// True can also be written using B1
	must.B1(trueFunc())(true)
	fmt.Printf("B1:%#v\n", must.Panic(func() {
		must.B1(false)(true, "this will also panic")
	}))

	// Output:
	// True:[]interface {}{"this will panic", "must be true"}
	// B1:[]interface {}{"this will also panic", false, true, "values must equal"}
}

func ExampleFalse() {
	must.False(false, "a no-op")
	fmt.Printf("False:%#v\n", must.Panic(func() {
		must.False(true, "this will panic")
	}))

	// False can also be written using B1
	must.B1(false)(false)
	fmt.Printf("B1:%#v\n", must.Panic(func() {
		must.B1(true)(false, "this will also panic")
	}))

	// Output:
	// False:[]interface {}{"this will panic", "must be false"}
	// B1:[]interface {}{"this will also panic", true, false, "values must equal"}
}
