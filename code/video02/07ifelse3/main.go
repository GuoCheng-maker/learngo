package main

import "fmt"

func main() {
	//for i := 0; i < 101; i++ {
	//	fmt.Println(i) // println和fmt.Println的区别
	//}
	//var count = 100
	//for count > 0 {
	//	fmt.Println(count)
	//	count--
	//}

	// 三要素for循环
	var sum int
	for count := 0; count <= 100; count++ {
		sum += count

	}
	fmt.Println(sum)
}
