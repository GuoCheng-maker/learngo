package main

import "fmt"

func setAge(age int) {
	age++
}

func setAge2(age *int) {
	*age++
}

func foo(s []int) {
	s[0] = 100
}

func main() {
	// 值传递
	age := 20
	age2 := 100
	setAge(age)
	setAge2(&age2)
	fmt.Println(age)  // 20 , x的值没有改变，因为传递的是值
	fmt.Println(age2) // 101 , x的值改变了,因为传递的是地址

	var s = []int{1, 2, 3}
	foo(s)
	fmt.Println(s)
}
