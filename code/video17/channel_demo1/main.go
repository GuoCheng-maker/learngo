package main

import "fmt"

func main() {
	s := make([]int, 4, 10)
	s = append(s, 1)
	fmt.Println(s)

	//var ch1 chan int // 引用类型，需要初始化以后才能使用
	//ch1 = make(chan int, 1)

	ch1 := make(chan int, 5) // 带缓冲区的channel
	// ch2 := make(chan int) // 不带缓冲区的channel， 又称为同步channel
	ch1 <- 10
	fmt.Println(len(ch1), cap(ch1)) // 取出缓冲区的长度和容量
	x := <-ch1
	fmt.Println(x)
	fmt.Println(len(ch1), cap(ch1))
	close(ch1)
}
