package main

import (
	"fmt"
	"reflect"
)

func foo() {
	var bar func(int, int) int
	bar = func(x, y int) int {
		a := x + y
		fmt.Println(a)
		return a
	}
	fmt.Println(reflect.TypeOf(bar))
	bar(99, 99)
}

func main() {
	//var f = func() {
	//	fmt.Println("hello world")
	//}
	//f()
	//f()
	//f()

	(func() {
		fmt.Println("hello world")
	})()

	(func(x, y int) {
		c := x + y
		fmt.Println(c)
	})(5, 10)
	foo()
}
