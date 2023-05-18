package main

import "fmt"

func foo() {
	fmt.Println("foo")
}

func bar(f func()) {
	fmt.Println(f)
	f()
}

func main() {
	bar(foo)
}
