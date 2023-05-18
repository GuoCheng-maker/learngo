package main

import (
	"fmt"
	"sort"
)

func main() {
	//s1 := []int{1, 2, 3}
	//s2 := s1[1:]
	//s2[1] = 4
	//fmt.Println(s1)
	//
	//var a = []int{1, 2, 3}
	//b := a
	//c := a[1:]
	//a[0] = 100
	//fmt.Println(b)
	//fmt.Println(c)

	//a := []int{10, 2, 3, 100}
	//sort.Ints(a)
	//fmt.Println(a) // [2 3 10 100]
	//
	//b := []string{"melon", "banana", "caomei", "apple"}
	//sort.Strings(b)
	//fmt.Println(b) // [apple banana caomei melon]
	//
	//c := []float64{3.14, 5.25, 1.12, 4, 78}
	//sort.Float64s(c)
	//fmt.Println(c) // [1.12 3.14 4 5.25 78]

	// 注意：如果是一个数组，需要先转成切片再排序  [:]
	a := [4]int{10, 2, 3, 100}
	c := [5]float64{3.14, 5.25, 1.12, 4, 78}

	sort.Sort(sort.Reverse(sort.IntSlice(a[:])))
	sort.Sort(sort.Reverse(sort.Float64Slice(c[:])))
	fmt.Println(a, c)

	var s1 = []int{1, 2, 3, 4, 5}
	var s2 = make([]int, len(s1))
	copy(s2, s1)
	fmt.Println(s2) // [1 2 3 4 5]

	s3 := []int{4, 5}
	s4 := []int{6, 7, 8, 9}
	copy(s4, s3)
	fmt.Println(s4) //[4 5 8 9]

	s5 := []int{6, 7, 8, 9}
	s6 := []int{4, 5}
	copy(s6, s5)
	fmt.Println(s6) // [6, 7]
}
