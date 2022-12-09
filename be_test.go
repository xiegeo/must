package must_test

import (
	"fmt"

	"github.com/xiegeo/must"
)

func ExampleB3() {
	// lets say we have an io.RuneReader
	readRune := func() (r rune, size int, err error) {
		return 'a', 1, nil
	}

	// that must return 1 byte runes and never err out
	char, _, _ := must.B3(readRune())(must.Any, 1, nil)
	fmt.Println(string(char))

	// output:
	// a
}
