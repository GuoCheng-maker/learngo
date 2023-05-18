package main

import (
	"fmt"
	"reflect"
)

func main() {

	// 三部曲1. 获取地址
	var x = 10 // x是整型变量
	fmt.Printf("%p\n", &x)
	//x = 100
	//fmt.Printf("%p\n", &x)

	// 三部曲2. 地址赋值：17，18行可以写为： var p = &x
	var p *int     // p是一个整型指针变量
	p = &x         // p指向x的内存地址
	fmt.Println(p) // p表示取p指向的内存地址

	// 三部曲3. 取值操作： *指针变量，  表示取指针变量指向的内存地址的值
	fmt.Println(*p, reflect.TypeOf(*p)) // *p表示取p指向的内存地址的值， type是int
	*p = 199
	fmt.Println(x) // 199

	// 小作业
	var a = 50
	var b = &a
	var c *int
	c = b
	fmt.Println(a, *b, *c)
	*c = 100
	fmt.Println(a, *b, *c)
	*b = 90
	fmt.Println(a, *b, *c)
	a = 80
	fmt.Println(a, *b, *c)

	// 指针应用1
	var x1 = 10
	var y1 = x1
	var z1 = &x1
	x1 = 20
	fmt.Println(x1)
	fmt.Println(y1)
	fmt.Println(*z1)
	*z1 = 30
	fmt.Println(x1)

	// 指针应用2
	var x2 = 10
	var y2 = &x2 //  *int 类型
	var z2 = *y2 // int 类型
	x2 = 20
	fmt.Println(x2)  // 20
	fmt.Println(*y2) // 20
	fmt.Println(z2)  // 10  因为z2是x2的值，而不是x2的地址

	*y2 = 30
	fmt.Println(z2) // 依旧是10

	// 指针应用3
	var a3 = 100
	var b3 = &a3 // *int
	var c3 = &b3 // **int
	// 打印a3的内存地址
	fmt.Println(&a3) // a的内存地址
	fmt.Println(*c3) // b3存的值，也就是a的内存地址
	fmt.Println(c3)  // b的内存地址
	fmt.Println(reflect.TypeOf(c3))
	**c3 = 200
	fmt.Println(a3)

	// 指针应用4
	p1 := 1
	p2 := &p1
	*p2++  // 虽然优先级中，++比*高，但是p2是一个地址的值，不能直接++，所以先执行*p2，再执行++
	fmt.Println(p1)  //2
	fmt.Println(*p2) //2
}
