package main

import "fmt"

func main() {
	//var a = []int{11, 22, 33}
	//var c = append(a, 44)
	//a[0] = 100
	//fmt.Println(a)                 // 100,22,33
	//fmt.Println(c, len(c), cap(c)) // [11,22,33,44] len(c) = 4 怎么算的呢？c = c[0:4] 4-0=4,cap为6是由于原来的a扩容一倍得到的6

	// 如果容量够呢？
	a := make([]int, 3, 10)
	fmt.Println(a) // [0,0,0]
	b := append(a, 11, 22)
	fmt.Println(a) // [0,0,0] 小心a等于多少，因为切片固定是初始指向，长度，容量。 a的长度是3，所以肯定是0，0，0，并不会把append的打印出来
	fmt.Println(b) // [0,0,0,11,22]
	a[0] = 100
	fmt.Println(a) // [100,0,0]
	fmt.Println(b) // [100,0,0,11,22]

	//
	arr := [4]int{10, 20, 30, 40}
	s1 := arr[0:2] // [10, 20]
	s2 := s1
	s3 := append(append(append(s1, 1), 2), 3)
	s1[0] = 1000
	fmt.Println(s1)  // [1000, 20]
	fmt.Println(s2)  // [1000, 20]
	fmt.Println(s3)  // [10,20,1,2,3]
	fmt.Println(arr) // [1000,20,1,2]

}
