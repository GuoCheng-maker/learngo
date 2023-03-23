package main

import (
	"fmt"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func Provider(data chan<- int, name string) {
	for i := 0; i < 100; i++ {
		select {
		case data <- i:
			fmt.Println(name, "生产了:", i)
		default:
		}
	}
	wg.Done()
}

func Consumer(data <-chan int, name string) {
	for {
		select {
		case tmp, ok := <-data:
			if !ok {
				wg.Done()
				return
			}
			fmt.Println(name, "读取了", tmp)
		default:
			fmt.Println("============================================================================")
		}
	}
}

func main() {
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
