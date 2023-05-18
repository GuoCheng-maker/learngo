package main

import "fmt"

func main() {
	a1 := []int{1, 2, 3, 4, 5}
	a2 := a1
	var a3 = make([]int, 3, 3)
	/*
		copy函数的第一个参数是目标切片，第二个参数是源切片，copy函数会把源切片的元素复制到目标切片中,如果目标切片的容量不够，则只复制目标切片的容量个元素
		如果对a1进行修改，a3的值不会改变。
	*/
	copy(a3, a1)

	fmt.Println(a1, a2, a3)

	a1[0] = 100
	fmt.Println(a1, a2, a3) // [100 2 3] [100 2 3] [1 2 3]
}
