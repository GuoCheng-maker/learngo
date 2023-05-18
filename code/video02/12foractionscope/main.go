package main

import "fmt"

func foo() {
	var y = 100
	fmt.Println(y)
}

func main() {
	var name = "yuan"
	for i := 0; i < 100; i++ {
		fmt.Println("000")
	}
	// fmt.Println(i) // undefined: i for循环作用域仅限于for循环中
	fmt.Println(name) // undefined: i for循环作用域仅限于for循环中
}
