package main

import "fmt"

type myInt int

func (m myInt) hello() {
	fmt.Println("我是一个int")
}

func main() {
	m := myInt(100)
	m.hello() // 调用自定义类型的方法,相当于给int类型额外添加了一个方法
}
