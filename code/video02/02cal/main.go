package main

import (
	"fmt"
)

func main() {
	x, y := 10, 20
	// 算术运算符
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(x % y)

	// 比较运算符 返回bool值
	fmt.Println(x > y)
	fmt.Println(x == y)
	fmt.Println(x != y)
	fmt.Println(x >= y)
	fmt.Println(x <= y)

	// 逻辑运算符
	// && 与运算（and）真与真为真 真与假为假 假与真为假  假与假为假
	// || 或运算（Or） 真或真为真 真或假为真 真或真为真  假或假为假
	// != 非运算（!=） 非真为假   非假为真
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(true != false)

	// 赋值运算符
	var n = 10
	n += 1
	fmt.Println(n)

	var y1 = 11
	y1++ // 12
	fmt.Println(y1)

	// 优先级
	var a, b, c = 1, 2, 3
	fmt.Println(a, b, c)
	var t = a + b*c
	fmt.Println(t)
}
