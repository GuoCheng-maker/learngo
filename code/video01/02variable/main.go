package main

import (
	"fmt"
)

func main() {
	// 声明变量 var 变量名 类型
	var age int
	fmt.Println(age) // 声明未使用， 默认为零值
	// 变量赋值
	age = 22
	fmt.Println(age)

	// 声明并赋值一行实现
	var name = 123
	fmt.Println(name)

	// 声明多个变量
	var x, y int
	fmt.Println(x, y)

	var (
		a int    // 默认为 0
		b string // 默认为 ""
		c bool   // 默认为 false
	)
	fmt.Println(a, b, c)

	// 声明并赋值的简洁写法
	name2 := "小王子" // 全局变量不可以使用 :=这种语法，这种写法仅限于函数内部使用
	fmt.Println(name2)

	// 一行声明赋值多个变量
	name3, age3, isMarried := "guocheng3", 26, false
	fmt.Println(name3, age3, isMarried)
}
