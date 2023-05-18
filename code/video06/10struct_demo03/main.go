package main

import "fmt"

// 类型包含属性和方法

// Student 声明结构体， 结构体是自定义数据类型，也是一种数据类型
type Student struct {
	sid, age int //字段名必须唯一
	name     string
	course   []string
	xxx      int
}

func main() {
	// k-v赋值 ，字段xxx不赋值也没关系，而且kv形式不用按照顺序
	var s1 = Student{name: "郭成", sid: 1001, course: []string{"语文", "数学", "英语"}, age: 26}
	// 多值赋值 这种方式必须按照声明的顺序,且所有的变量都必须赋值
	var s2 = Student{1001, 26, "郭成", []string{"语文", "数学", "英语"}, 99}
	fmt.Println(s1)
	fmt.Println(s2)
}
