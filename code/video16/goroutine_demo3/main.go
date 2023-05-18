package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func a() {
	for i := 0; i < 100; i++ {
		fmt.Println("A:", i)
	}
	wg.Done()
}

func b() {
	for i := 0; i < 100; i++ {
		fmt.Println("B:", i)
	}
	wg.Done()
}

func main() { // 开启一个主goroutine去执行main函数
	runtime.GOMAXPROCS(6) // 设置使用几个核
	wg.Add(2)
	go a()
	go b()
	fmt.Println("hello main")
	wg.Wait() // 阻塞等待计数器变为0
}
