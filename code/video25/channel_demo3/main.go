package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	NumProducers = 50
	NumConsumers = 100
	BufferSize   = 100
)

func main() {
	buffer := make(chan int, BufferSize)
	var wg sync.WaitGroup

	// Launch producers
	for i := 0; i < NumProducers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 200; j++ {
				buffer <- j
			}
		}(i)
	}

	// Launch consumers
	for i := 0; i < NumConsumers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				val, ok := <-buffer
				if !ok {
					return
				}
				fmt.Println(val)
			}
		}(i)
	}

	// Wait for all producers and consumers to finish
	wg.Wait()
	close(buffer)
	time.Sleep(time.Second)
	fmt.Println("123")
}
