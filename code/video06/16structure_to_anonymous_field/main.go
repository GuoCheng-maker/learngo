package main

import "fmt"

type Address struct {
	province, city string
}

type Person struct {
	string // 光写类型默认其实是 string string
	int
	Address
}

type Student struct {
	name string // 光写类型默认其实是 string string
	age  int
	addr Address // 可以将结构体作为另一个结构体字段的类型
}

func main() {
	p1 := Person{"小明", 18, Address{"山西省", "太原市"}}
	fmt.Println(p1.string)
	fmt.Println(p1.int)
	fmt.Println(p1.city) // 匿名函数如果使用自己构造的类型，则可以p1.city直接使用。等同于p1.Address.city, 当然Person结构里不能用city这个字段

	s1 := Student{"小明", 18, Address{"北京", "海淀区"}}
	fmt.Println(s1.addr.city)
	// fmt.Println(s1.city)  // 这么用不对
}
