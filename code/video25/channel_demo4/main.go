package main

import (
	"fmt"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func Provider(data chan<- int, name string) {
	for i := 0; i < 100; i++ {
		//保证传入i的操作和print操作"原子运行"
		select {
		case data <- i:
			fmt.Println(name, "生产了:", i)
		}
	}
	wg.Done()
}

func Consumer(data <-chan int, name string) {
	for {
		tmp, ok := <-data
		if !ok {
			break
		}
		fmt.Println(name, "读取了", tmp)
	}
	wg.Done()
}

func main() {
	//定义缓冲区大小
	data := make(chan int, 10)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go Provider(data, strconv.Itoa(i))
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go Consumer(data, strconv.Itoa(i))
	}

	wg.Wait()
	fmt.Println("end")
}
