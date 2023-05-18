package main

import (
	"bytes"
	"fmt"
)

func main() {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')
	fmt.Println(sepIndex)

	// dir1 := path[:sepIndex]  // len 4  cap 14

	// 在末尾加上第二个冒号和一个 sepIndex，表示将 dir1 的容量设置为 sepIndex，但是依旧共享 path 的底层数组，
	// 不过一旦扩容，怎不在共享
	dir1 := path[:sepIndex:sepIndex] // len 4  cap 4

	fmt.Println(len(dir1), cap(dir1))
	dir2 := path[sepIndex+1:]

	ptr1 := &dir1[0]
	ptr2 := &path[0]          // 使用 path 的地址而不是 dir2 的地址
	fmt.Println(ptr1 == ptr2) // 输出 true

	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAA
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => BBBBBBBBB

	dir1 = append(dir1, "suffix"...)

	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAAsuffix
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => uffixBBBB
}
