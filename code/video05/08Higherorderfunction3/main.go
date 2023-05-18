package main

import "fmt"

// foo也是一个高阶函数， 它返回一个匿名函数
func foo() func() int {
	// 声明匿名函数
	var inner = func() int {
		fmt.Println("hello world")
		return 100
	}
	return inner
}

func main() {
	foo()() // 调用foo函数，返回一个匿名函数，再调用匿名函数

	// 这个写法和上面的调用意思是一样的
	var f func() int
	f = foo()
	f()
}
