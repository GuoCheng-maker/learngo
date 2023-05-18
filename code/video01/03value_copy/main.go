package main

import (
	"fmt"
)

func main() {
	// 值拷贝
	var x = 10
	var y = x // 值拷贝, x, y对应的内存地址不同，而Python中的变量是引用, x, y对应的内存地址相同
	x = 20
	fmt.Println(x)
	fmt.Println(y)

	//
	var a = 1 + 1
	fmt.Println(a)
	var b = x * y
	fmt.Println(b)

	var x1, y1 = 10, 20
	fmt.Println(x1 + y1)
	fmt.Println(x1 - y1)
	fmt.Println(x1 * y1)
	fmt.Println(x1 / y1)
	fmt.Println(x1 % y1)

	fmt.Println("================")
	// 这种写法是不能交换x2,y2的值
	var x2 = 10
	var y2 = 20
	//x2 = y2
	//y2 = x2
	//fmt.Println(x2, y2)

	//正确的交换
	//x2, y2 = y2, x2
	//fmt.Println(x2, y2)

	// 临时变量实现交换也可以
	tmp := x2
	x2 = y2
	y2 = tmp
	fmt.Println(x2, y2)
}
