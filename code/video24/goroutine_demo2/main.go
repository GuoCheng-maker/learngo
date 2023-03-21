package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() { // 开启一个主goroutine去执行main函数
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go func(i int) {
			fmt.Printf("hello %d\n", i)
			wg.Done()
		}(i)
	}
	fmt.Println("hello main")
	wg.Wait() // 阻塞等待计数器变为0
}
