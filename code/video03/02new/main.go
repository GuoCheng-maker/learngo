package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x = 10     // 基本数据类型（整型,浮点型,布尔类型,字符串类型,数组,结构体）都是值类型，指针类型属于引用类型（之后还有切片，map, channel）
	fmt.Println(x) // 值类型特点： 当声明未赋值之前存在一个默认值，而引用类型不存在默认值的，声明后不开辟空间

	//var p *int
	//*p = 100 // 指针类型的默认值是nil，如果不开辟空间，直接赋值会报错
	//fmt.Println(*p)

	var p *int
	p = new(int)   // new函数的作用是开辟一块内存空间，将new括号里面的类型的零值赋值给这块内存空间，并且返回这块内存空间的地址
	fmt.Println(p) // 0x14000122010
	*p = 100
	fmt.Println(*p) // 100

	// new函数的使用
	// new函数和make函数的区别
	// 1.new 和 make 是 Go 语言中用于内存分配的原语。简单来说，new 只分配内存，make 用于初始化 slice、map 和 channel
	// 3.new函数是在堆上分配内存，make函数是在栈上分配内存
	// 4.new函数的返回值是指针，make函数的返回值是引用类型
	// 5.new函数的返回值是类型的零值，make函数的返回值是类型的初始化值

	var sl1 []int
	fmt.Println(sl1, reflect.TypeOf(sl1))
	sl1 = append(sl1, 100)
	sl1 = append(sl1, []int{1, 2, 3, 4, 5}...)
	fmt.Println(sl1)
	var m1 map[string]int
	fmt.Println(m1, reflect.TypeOf(m1))
	m1["key1"] = 10
}

