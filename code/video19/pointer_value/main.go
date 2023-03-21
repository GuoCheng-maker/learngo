package main

import "fmt"

//使用值接受者实现接口和使用指针接受者实现接口的区别

// 接口的嵌套
type animal interface {
	mover
	sayer
}

type mover interface {
	move()
}

type sayer interface {
	say()
}

type person struct {
	name string
	age  int8
}

// 使用值接受者实现接口 类型的值和类型的指针都能保存到接口变量中
//func (p person) move() {
//	fmt.Printf("%s在跑步\n", p.name)
//}

// 使用指针接受者实现接口 只有类型的指针能保存到接口变量中
func (p *person) move() {
	fmt.Printf("%s在跑步\n", p.name)
}

func (p *person) say() {
	fmt.Printf("%s在叫\n", p.name)
}

func (p *person) heihei() {
	fmt.Printf("%s在叫\n", p.name)
}

func main() {
	var m mover // 定义了一个mover类型的变量
	// p1 := person{name: "小王子", age: 18}  // person类型的值
	p2 := &person{name: "大王子", age: 18} // person类型的指针
	// m = p1  // 无法赋值 因为p1是person类型的值，没有实现mover接口
	m = p2
	m.move()

	var s sayer // 定义了一个sayer类型的变量
	s = p2
	s.say()
	// 一个类型（person）可以实现多个接口，多个类型(猫和狗都能叫)可以实现同一个接口

	var a animal // 定义了一个animal类型的变量
	a = p2
	a.say()
	a.move()
}
