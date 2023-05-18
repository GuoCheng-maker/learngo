package calc

import "fmt"

// SayHello2 首字母大写
func SayHelloAdd() {
	println("SayHelloAdd!")
}

func Add(x, y int) int {
	fmt.Println(x + y)
	return x + y
}

// Name 全局变量
var Name = "add"

// init函数在包倒入时候自动执行
// init函数没有参数也没有返回值
func init() {
	println("add.init函数")
}
