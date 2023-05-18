package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("worker %d start job %d\n", id, job)
		results <- job * 2
		time.Sleep(time.Millisecond * 500)
		fmt.Printf("worker %d end job %d\n", id, job)
	}
}

func main() {
	job := make(chan int, 100)
	results := make(chan int, 100)

	// 开启3个协程
	for i := 0; i < 3; i++ {
		go worker(i, job, results)
	}

	// 5个任务
	for i := 0; i < 5; i++ {
		job <- i
	}
	close(job)

	for i := 0; i < 5; i++ {
		ret := <-results
		fmt.Println(ret)
	}
}
