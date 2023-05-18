package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	name string
}

// 值接受者
func (a Animal) eat() {
	fmt.Printf("%s在吃东西\n", a.name)
}

// 指针接受者
func (a *Animal) sleep() {
	fmt.Printf("%s在睡觉\n", a.name)
}

func (d *Dog) bark() {
	fmt.Printf("%s在狂吠\n", d.name)
}

type Dog struct {
	kind string
	Animal
}

type Cat struct {
	Animal
}

func main() {
	/*
		可以调用d.bark()的原因是因为Dog类型定义了一个方法bark()，因此任何Dog类型的值都可以调用该方法。
		在调用d.bark()时，Go语言会自动将d转换为方法接收者类型。
		如果方法定义的接收者是值类型，那么调用时会对d进行值拷贝，如果方法定义的接收者是指针类型，那么调用时会对d进行取地址操作。
		在下面例子中，d是一个Dog类型的值，而bark()方法的接收者是一个指向Dog类型的指针，所以Go语言会自动将d的地址作为方法的接收者进行调用。
		换句话说，d.bark()其实是(&d).bark()的简写形式，其中&d表示d的地址，也就是一个指向Dog类型的指针。
		总结一下，无论是值类型还是指针类型的接收者，Go语言都会自动进行类型转换。
	*/
	d := Dog{kind: "金毛", Animal: Animal{"小黑"}}
	fmt.Println(reflect.TypeOf(d))
	d.eat()   // 使用的是Animal的方法
	d.sleep() // 使用的是Animal的方法，  按理来说应该是(&d).sleep()
	d.bark()  // 使用的是Dog的方法

	c := Cat{Animal: Animal{"小花"}}
	c.eat()   // 使用的是Animal的方法
	c.sleep() // 使用的是Animal的方法

}
