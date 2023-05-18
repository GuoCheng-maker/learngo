package main

import "fmt"

func main() {
	var s []int
	fmt.Println(len(s), cap(s))
	s1 := append(s, 1)
	s2 := append(s1, 1, 2, 3, 4)
	fmt.Println(s1)
	fmt.Println(s2)

	var t = []int{9, 8, 7, 6}
	s3 := append(s2, t...)
	fmt.Println(s3)

	var s4 = make([]int, 3, 10)
	s5 := append(s4, 100)
	fmt.Println(s5) // [[0 0 0 100]
}
