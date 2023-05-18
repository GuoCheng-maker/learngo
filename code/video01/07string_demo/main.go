package main

import "fmt"

func main() {
	var s string
	s = "hello yuan"

	// 索引操作
	fmt.Println(string(s[0]))

	// 切片操作
	fmt.Println(s[0:5])
	fmt.Println(s[:5])
	fmt.Println(s[6:10])
	fmt.Println(s[6:])
	// fmt.Println(s[6:13]) // ：后面的索引不能超过字符串的长度，也就是最大是[6:10], 如果为11则报错，而Python不会

	// 字符串拼接
	s1 := "hello"
	s2 := "world"
	s3 := s1 + s2
	fmt.Println(s3)

	// 转意符号 \n \r \t
	fmt.Println("hello \nworld\n!")

	// 取消转义
	fmt.Println("D:\\Go\\src\\day1\\main.go") // D:\Go\src\day1\main.go
	fmt.Println("his name is \"rain\"")

	// 多行打印
	info := `1.package main
2.import "fmt"
3.func main()`
	fmt.Println(info)
}
