package main

import (
	"fmt"
	"reflect"
)

func main() {

	// 整型
	var x int = 100
	fmt.Println(x)

	var age uint8 = 150
	fmt.Println(age)

	// 浮点型
	var f1 float32 = 3.0123456789123456789 //最多显示7位小数
	var f2 float64 = 3.0123456789123456789 //最多显示16位小数
	fmt.Println(f1)
	fmt.Println(f2)
	var f3 = 2e10 // 科学计数法， float类型
	fmt.Println(f3, reflect.TypeOf(f3))

	// bool类型
	var b1 bool
	b1 = true
	b1 = false
	fmt.Println(b1)
	fmt.Println(2 < 1)

	c := 2 > 1
	fmt.Println(c, reflect.TypeOf(c))
}
