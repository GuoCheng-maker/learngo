package main

import (
	"fmt"
	"strconv"
)

func main() {
	var stu01 = map[string]string{"name": "yuan", "age": "18"}
	fmt.Println(stu01)
	// 增加
	stu01["gender"] = "female"
	fmt.Println(stu01)
	// 修改
	stu01["name"] = "guo"
	fmt.Println(stu01)
	// 删除
	delete(stu01, "age")
	fmt.Println(stu01)
	// 查询
	fmt.Println(stu01["name"])

	// 基于make函数声明初始化
	stu2 := make(map[string]string) // make(map[string]interface{}) ,如果不知道value的类型，可以用interface{}
	stu2["name"] = "test"
	stu2["age"] = "18"
	stu2["gender"] = "male"
	// 如果不使用interface，则对age进行+1，需要进行类型转换
	age, _ := strconv.Atoi(stu2["age"])
	age++
	stu2["age"] = strconv.Itoa(age)
	fmt.Println(stu2)

	// 回顾遍历一个数组
	//var arr = []int{1, 2, 3, 4, 5}
	//for i, v := range stu2 {
	//	fmt.Println(i, v)
	//}

	// 遍历一个map
	for k, v := range stu2 {
		if k == "age" && v == "19" {
			continue
		}
		fmt.Println(k, v)
	}
}
