package main

import "fmt"

// 这种写法会污染全局变量
//var x int
//
//func counter() {
//	x++
//	fmt.Println(x)
//}

func getCounter() func() {
	i := 0
	counter := func() {
		i++
		fmt.Println(i)
	}
	return counter

}

func main() {
	counter1 := getCounter()
	counter1()
	counter1()
}
