package main

import "fmt"

func main() {
	// 初始化创建切片
	var s = make([]int, 5, 10) // s是一个切片，切片是一个引用类型。  new()和make()是给引用类型申请内存空间。
	fmt.Println(s, len(s), cap(s))
	s[0] = 100
	fmt.Println(s)

	arr := new([2]int) // arr是一个指针 ,指针是一个引用类型。  new()和make()是给引用类型申请内存空间。
	fmt.Println(*arr)

}
