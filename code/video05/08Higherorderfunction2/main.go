package main

import (
	"fmt"
	"time"
)

func foo() {
	fmt.Println("foo start")
	time.Sleep(time.Second * 3)
	fmt.Println("foo end")
}

func bar() {
	fmt.Println("bar start")
	time.Sleep(time.Second * 3)
	fmt.Println("bar end")

}

func timeCost(f func()) {
	start := time.Now().Unix()
	f()
	end := time.Now().Unix()
	fmt.Printf("time cost: %d", end-start)
}

func main() {
	//f1 := time.Now().Unix()
	//fmt.Println(f1, reflect.TypeOf(f1))
	//
	//time.Sleep(time.Second * 3)
	//
	//f2 := time.Now().Unix()
	//fmt.Println(f2, reflect.TypeOf(f2))

	//foo()
	//bar()
	timeCost(foo)
	timeCost(bar) // timeCost是一个高阶函数，它接收一个函数作为参数
}
