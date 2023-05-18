package main

import "fmt"

// 接口

type dog struct{}

func (d dog) say() {
	fmt.Println("汪汪汪")
}

type cat struct{}

func (c cat) say() {
	fmt.Println("喵喵喵")
}

type person struct {
	name string
}

func (p person) say() {
	fmt.Println("嘿嘿嘿")
}

// 接口不管你是什么类型，只管你要实现哪个方法
// 定义一个类型，一个抽象的类型，只要实现了say方法，都可以称为sayer类型
type sayer interface {
	say()
}

// 接受一个sayer类型的变量（）
func beat(arg sayer) {
	arg.say()
}

func main() {
	c1 := cat{}
	beat(c1)
	d1 := dog{}
	beat(d1)

	var s1 sayer
	p1 := person{name: "小王子"}
	s1 = p1 // sayer类型的变量可以接收person类型的变量
	fmt.Println(s1)
	s1.say()

}
