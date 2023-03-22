package main

import (
	"fmt"
	"sync"
)

const (
	numProducer = 50
	numConsumer = 50
	numTasks    = 1000
)

var wg sync.WaitGroup

func producer(taskChan chan<- int) {
	defer wg.Done()
	for i := 0; i < numTasks; i++ {
		taskChan <- i
	}
}

func consumer(id int, taskChan <-chan int) {
	defer wg.Done()
	for {
		task, ok := <-taskChan
		if !ok {
			break
		}
		fmt.Printf("Consumer %d received task %d\n", id, task)
	}
}

func main() {
	taskChan := make(chan int, 10000)

	for i := 0; i < numProducer; i++ {
		wg.Add(1)
		go producer(taskChan)
	}

	for i := 0; i < numConsumer; i++ {
		wg.Add(1)
		go consumer(i, taskChan)
	}
	wg.Wait()
	close(taskChan)

	fmt.Println("end!")

}
