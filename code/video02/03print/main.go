package main

import "fmt"

func main() {
	// 输出函数 Print和Println
	// fmt.Println("hello", "world", "!")
	// fmt.Print("hello", "world", "!")
	// fmt.Printf("%s", "hello world!")

	var name, age, isMarried = "yuan", 32, false
	// 标准化输出
	fmt.Printf("name:%s, age:%d, isMarried: %t \n", name, age, isMarried)

	// Sprintf
	s := fmt.Sprintf("name:%s, age:%d, isMarried: %t", name, age, isMarried)
	fmt.Println(s)

	// 输入函数
	var (
		name2 string
		age2  int
	)
	fmt.Print("请输入姓名和年龄：")
	fmt.Scan(&name2, &age2)
	fmt.Println(name2, age2)
	fmt.Println("end")

	const pai float64 = 3.1415926
	const URL string = "https://www.baidu.com"
	fmt.Println(pai, URL)

}
