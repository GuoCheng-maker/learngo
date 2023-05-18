package main

import (
	"fmt"
	"time"
)

func foo() {
	time.Sleep(time.Second * 2)
	println("foo")

}

func bar() {
	time.Sleep(time.Second * 3)
	println("bar")

}

func getProTimer(f func()) func() {
	inner := func() {
		startTime := time.Now().Unix()
		f()
		endTime := time.Now().Unix()
		fmt.Println("函数执行时间为", endTime-startTime, "秒")
	}
	return inner

}
func main() {
	foo := getProTimer(foo)
	foo()

	bar := getProTimer(bar)
	bar()
}
