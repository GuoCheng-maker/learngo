package calc

import "fmt"

func SayHelloSub() {
	println("SayHellosub!")
}

func Sub(x, y int) int {
	fmt.Println(x - y)
	return x - y
}
