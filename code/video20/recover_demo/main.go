package main

import (
	"fmt"
	"time"
)

func f1() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("recover panic:%v\n", e)
		}
	}()
	// 开启一个goroutine执行任务
	go func() {
		fmt.Println("in goroutine....")
		// 只能触发当前goroutine中的defer
		panic("panic in goroutine")
	}()

	time.Sleep(time.Second)
	fmt.Println("exit")
}

func f2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recover outer panic:%v\n", r)
		}
	}()
	// 开启一个goroutine执行任务
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("recover inner panic:%v\n", r)
			}
		}()
		fmt.Println("in goroutine....")
		// 只能触发当前goroutine中的defer
		panic("panic in goroutine")
	}()

	time.Sleep(time.Second)
	fmt.Println("exit")
}

func main() {
	// f1()
	// in goroutine....
	// panic: panic in goroutine
	// goroutine 18 [running]:
	// main.f1.func2()
	// /Users/guocheng/GoLandProjects/learngo/code/video28/recover_demo/main.go:18 +0x68
	// created by main.f1
	// /Users/guocheng/GoLandProjects/learngo/code/video28/recover_demo/main.go:15 +0x40

	f2()
	// in goroutine....
	// recover inner panic:panic in goroutine
	// exit
}
