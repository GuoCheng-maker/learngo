package main

import "fmt"

func foo() {
	println("foo")

}

func bar() {
	println("bar")

}

func calledNum(f func()) func() {
	var count int
	inner := func() {
		count++
		f()
		fmt.Println("函数被调用了", count, "次")
	}
	return inner

}
func main() {
	foo := calledNum(foo)
	foo()
	foo()
	foo()

	bar := calledNum(bar)
	bar()
	bar()
	bar()

}
