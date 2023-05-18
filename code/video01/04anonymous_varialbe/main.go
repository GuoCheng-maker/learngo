package main

import "fmt"

func foo() (int, int) {
	return 1, 2
}

func main() {
	var a, _ = foo() // 匿名变量一半用函数引用返回的值，比如上面foo()函数, 我只想使用第一个值，则第二个值可以用匿名变量隐掉
	// _表示匿名变量，匿名变量不占用命名空间，不会分配内存，匿名变量之间也不会因为重复声明而无法使用
	fmt.Println(a)

	firstName := "guo" // 命名规范：驼峰命名法
	secondName := "cheng"
	fmt.Println(firstName, secondName)
}
