package main

import (
	"fmt"
)

func main() {
	// 方式1.通过数组获取一个切片
	//var names = [3]string{"张三", "李四", "王五"}
	//var slices = names[1:3]
	//fmt.Printf("%T\n", slices)

	s := []string{"1", "2"}
	ss := make([]int, 5, 10)
	k := []int{8, 9, 10}
	b := []int{8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	ss[0] = 10000
	m := append(ss, k...)
	ss[1] = 20000
	v := append(ss, b...)
	ss[2] = 30000
	fmt.Println(ss, m, v)

	s[0] = "333"
	fmt.Println(s)

	var a = [5]int{1, 2, 3, 4, 5}
	// 切片是对数组的一个引用，切片是一个引用类型， 切片中的start和end是对数组的索引, 存储了 （起始地址, 长度, 容量）
	var slices = a[:]
	fmt.Println(slices)
	newSlices := slices[1:3]

	newSlices[1] = 1000
	fmt.Println(a)
	fmt.Println(len(a), cap(a)) // 5 5
	fmt.Println(slices)
	fmt.Println(len(slices), cap(slices)) // 5 (end-start=5-0=5)  5 (len(arr) - start = 5-0=5)
	fmt.Println(newSlices)
	fmt.Println(len(newSlices), cap(newSlices)) // 2(end-start=3-1=2)   4 (len(arr) - start = 5-1=4)

	// 声明一个切片，但是没有初始化，切片的值为nil
	var array = [5]int{1, 2, 3, 4, 5}
	// 方式2.直接声明切片,区别于数组，切片没有长度
	var slice = []int{1, 2, 3, 4, 5} // 等同于 var array = [5]int{1, 2, 3, 4, 5}; var slice = array[:]
	fmt.Println(array, slice)

	s1 := slice[1:4]
	fmt.Println(len(s1), cap(s1)) // 3  4
	s2 := slice[3:]
	fmt.Println(len(s2), cap(s2)) // 2  2

	s3 := s1[1:2]                 // s3 = [3]
	fmt.Println(len(s3), cap(s3)) // 1(2-1=1)  (len(slice)-s3中初始元素在slice中的位置 = 5-2=3)

	var arrr = [...]int{1, 2, 3, 4, 5, 6}
	a1 := arrr[0:3]
	a2 := arrr[0:5]
	a3 := arrr[1:5] // [2, 3, 4, 5]
	a4 := arrr[1:]
	a5 := arrr[:]
	a6 := a3[1:2]                                  // [3]
	fmt.Printf("a1的长度%d，容量%d\n", len(a1), cap(a1)) // 3 6
	fmt.Printf("a2的长度%d，容量%d\n", len(a2), cap(a2)) // 5 6
	fmt.Printf("a3的长度%d，容量%d\n", len(a3), cap(a3)) // 4 5
	fmt.Printf("a4的长度%d，容量%d\n", len(a4), cap(a4)) // 5 5
	fmt.Printf("a5的长度%d，容量%d\n", len(a5), cap(a5)) // 6 6
	fmt.Printf("a6的长度%d，容量%d\n", len(a6), cap(a6)) // 1 4

}
