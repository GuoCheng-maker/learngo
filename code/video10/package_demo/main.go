package main

import (
	"fmt"
	"learngo/utils/tools"
	"learngo/utils/utils"
)

func foo(x, y int) (a, b int) {
	a = x + y
	b = x - y
	return a, b
}

func main() {
	fmt.Println(tools.Name)
	// fmt.Println(tools.name) // 首字母小写不能访问
	sum1 := utils.Sum(1, 2)
	sum2 := tools.Sum(1, 2)
	println(sum1, sum2)
	var s = make([]int, 10)
	fmt.Println(s, len(s), cap(s))
	arr := [4]int{10, 20, 30, 40}
	s1 := arr[0:2] // [10, 20]
	s2 := s1       //  [10, 20]
	s3 := append(append(append(s1, 1), 2), 3)
	s1[0] = 1000

	fmt.Println(s1)  //[1000, 20]
	fmt.Println(s2)  //[1000, 20]
	fmt.Println(s3)  // [10, 20, 1, 2, 3]
	fmt.Println(arr) // [1000,20,1,2]

	a, b := foo(10, 1)
	fmt.Println(a, b)

}
