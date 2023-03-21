package main

import (
	"fmt"
	"learngo/code/calc"
	"learngo/code/tools"
	"learngo/code/utils"
	"learngo/code/weather"
)

func main() {
	tools.Sum(1, 9)
	utils.Sum(10, 20)
	weather.Snow()
	// 先执行包的init函数,按照顺序，如果包A里引用包B,则现执行B的init函数
	calc.Add(1, 2)
	calc.SayHelloAdd()
	calc.Sub(5, 1)
	calc.SayHelloSub()
	fmt.Println(calc.Name)
}
