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

// Student类型的方法接收器
func (s Student) read(book string) {
	fmt.Printf("%s学生正在读%s书\n", s.name, book)
}

// Student类型的方法接收器
func (s Student) lean() {
	fmt.Printf("%s学生正在学习\n", s.name)
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
	s1.lean()
	s1.read("go语言")

	s2 := newStudent(1003, 28, "rain", []string{"语文", "数学", "英语"})
	fmt.Println(s2.name)
	s2.lean()
	s2.read("C语言")
}
