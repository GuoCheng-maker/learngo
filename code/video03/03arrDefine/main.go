package main

import "fmt"

func main() {
	// 数组必须限制长度
	var arr [3]int
	fmt.Print(arr)

	// 索引取值
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])

	// 索引赋值
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	fmt.Println(arr)

	// 数组的声明并赋值
	var names = [3]string{"张三", "李四", "王五"}
	fmt.Println(names)
	var ages = [3]int{10, 20, 30}
	fmt.Println(ages)

	// 省略长度并赋值
	var scores = [...]int{100, 99, 98}
	fmt.Println(scores)

	// 索引赋值
	var names2 = [...]string{0: "李四", 2: "王五"}
	fmt.Println(names2)

	// len()函数获取数组长度
	fmt.Println(len(names2)) // 3个，有一个是空字符串
	fmt.Println(len("test")) // 4个
}
