package main

import (
	"fmt"
	"learngo/code/video33/test2_demo"
	"unsafe"
)

func trace(s string)   { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }

func aoo() {
	trace("a")
	defer untrace("a")
	// do something....
}

var b = 11

func foo(n int) {
	b := 1
	b += n
}

func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
	aoo()
	fmt.Println(b)
	foo(999)
	fmt.Println(b)

	var a, d = int(5), uint(6)
	fmt.Println(a, d)
	var p uintptr = 0x12345678
	fmt.Println("signed integer a's length is", unsafe.Sizeof(a))   // 8
	fmt.Println("unsigned integer b's length is", unsafe.Sizeof(d)) // 8
	fmt.Println("uintptr's length is", unsafe.Sizeof(p))            // 8

	var f1 float32 = 16777216.0
	var f2 float32 = 16777217.0
	fmt.Println(f1 == f2) // true
	fmt.Println("dsad21" > "123456")
	fileSize := test2_demo.ByteSize(4567890)
	fmt.Println(fileSize.String()) // 输出：4.36MB
}
