package must_test

import (
	"errors"
	"fmt"
	"io"

	"github.com/xiegeo/must"
)

func ExampleNoError() {
	var err error // create nil error
	must.NoError(err)
	fmt.Printf("NoError:%#v\n", must.Panic(func() {
		must.NoError(io.EOF)
	}))

	// NoError can also be written using B1
	_ = must.B1(err)(nil)
	fmt.Printf("B1:%v\n", must.Panic(func() { // use %v because %#v prints out pointer address
		_ = must.B1(io.EOF)(nil)
	}))

	// Output:
	// NoError:&errors.errorString{s:"EOF"}
	// B1:[EOF value must be nil]
}

func ExampleValue() {
	var err error // create nil error
	fmt.Printf("Value nil error:%#v\n", must.Value(5, err))
	fmt.Printf("Value with error:%#v\n", must.Panic(func() {
		_ = must.V(6, io.EOF)
	}))

	// Value can also be written using B2
	value, _ := must.B2(7, err)(must.Any, nil)
	fmt.Printf("B2 nil error:%#v\n", value)
	fmt.Printf("B2 with error:%v\n", must.Panic(func() { // use %v because %#v prints out pointer address
		_, _ = must.B2(8, io.EOF)(must.Any, nil)
	}))

	// Output:
	// Value nil error:5
	// Value with error:&errors.errorString{s:"EOF"}
	// B2 nil error:7
	// B2 with error:[#2 EOF value must be nil]
}

func ExampleValue2() {
	fmt.Print("Value2 nil error: ")
	fmt.Println(must.Value2(1, 2, nil))

	fmt.Println("Panic message:", must.Panic(func() {
		must.Value2(1, 2, errors.New("P"))
	}).(error).Error())

	// Output:
	// Value2 nil error: 1 2
	// Panic message: P
}

func ExampleValue3() {
	fmt.Print("Value3 nil error: ")
	fmt.Println(must.Value3(1, 2, 3, nil))

	fmt.Println("Panic message:", must.Panic(func() {
		must.Value3(1, 2, 3, errors.New("P"))
	}).(error).Error())

	// Output:
	// Value3 nil error: 1 2 3
	// Panic message: P
}

func ExampleValue4() {
	fmt.Print("Value4 nil error: ")
	fmt.Println(must.Value4(1, 2, 3, 4, nil))

	fmt.Println("Panic message:", must.Panic(func() {
		must.Value4(1, 2, 3, 4, errors.New("P"))
	}).(error).Error())

	// Output:
	// Value4 nil error: 1 2 3 4
	// Panic message: P
}
