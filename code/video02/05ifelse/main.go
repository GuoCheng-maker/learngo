package main

import "fmt"

func main() {
	var age int
	fmt.Print("请输入年龄：")
	fmt.Scan(&age)
	fmt.Println(age)
	// 单分支语句
	if age > 18 {
		fmt.Println("成年人.......")
	}

	// 双分支语句
	if age >= 18 {
		println("成年人")
	} else {
		println("未成年人")
	}
}
