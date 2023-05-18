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

func newStudent(sid, age int, name string, course []string) Student {
	s1 := Student{
		name:   name,
		sid:    sid,
		course: course,
		age:    age,
	}
	return s1
}

func main() {
	s1 := newStudent(1002, 26, "郭成", []string{"语文", "数学", "英语"})
	fmt.Println(s1.name)
}
