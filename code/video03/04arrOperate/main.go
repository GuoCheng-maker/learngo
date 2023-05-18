package main

import (
	"fmt"
	"reflect"
)

func main() {
	var names = [...]string{"张三", "李四", "王五"}
	// 索引的操作
	fmt.Println(names[2])
	names[2] = "赵六"
	fmt.Println(names)

	var arr = [...]int{1, 2, 3, 4, 5}
	// 切片操作 arr[start索引:end索引]  左闭右开
	s := arr[1:3]
	fmt.Println(s, reflect.TypeOf(s)) // [2 3] []int

	fmt.Println(arr[1:])
	fmt.Println(arr[:3])

	// 遍历数组
	for i := 0; i < len(names); i++ {
		println(names[i])
	}

	// range循环
	for i, v := range names {
		v = "test"
		fmt.Println(i, v)
	}
	fmt.Println(names)
}
