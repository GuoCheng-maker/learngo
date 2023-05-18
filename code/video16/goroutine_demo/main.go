package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// goroutine
func hello(i int) {
	fmt.Printf("hello %d\n", i)
	wg.Done() // 计数器-1
}

func main() { // 开启一个主goroutine去执行main函数
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go hello(i) // 计数器+1
	}
	// wg.Add(1)
	// go hello() // 开启一个单独的goroutine去执行hello函数
	fmt.Println("hello main")
	wg.Wait() // 阻塞等待计数器变为0
}
