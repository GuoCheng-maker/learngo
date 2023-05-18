package main

import (
	"fmt"
	calc2 "learngo/utils/calc"
	"learngo/utils/tools"
	"learngo/utils/utils"
	"learngo/utils/weather"
)

func main() {
	tools.Sum(1, 9)
	utils.Sum(10, 20)
	weather.Snow()
	// 先执行包的init函数,按照顺序，如果包A里引用包B,则现执行B的init函数
	calc2.Add(1, 2)
	calc2.SayHelloAdd()
	calc2.Sub(5, 1)
	calc2.SayHelloSub()
	fmt.Println(calc2.Name)
}
