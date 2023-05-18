package main

import (
	"fmt"
)

// 类型包含属性和方法

// Student 声明结构体， 结构体是自定义数据类型，也是一种数据类型
type Student struct {
	sid, age int //字段名必须唯一
	name     string
	course   []string
	xxx      int
}

type Person struct {
	name, gender, addr string
	age                uint8
}

func main() {
	var s Student  //上面声明不开辟空间，这里实例化的时候才去真正的按照结构体里面的字段去开辟空间
	fmt.Println(s) // {0 0  []}
	s.sid = 1
	s.age = 18
	s.name = "小明"
	// s.course = []string{"语文", "数学", "英语"}
	fmt.Println(s)      //{1 18 小明 [语文 数学 英语]}
	fmt.Println(s.name) //小明
	s.course = append(s.course, "语文")
	s.course = append(s.course, "数学")
	fmt.Println(s.course)
	fmt.Println(s.course[0])

	// 结构体的内存布局
	fmt.Printf("%p\n", &s)          //0x14000120040  结构体的地址,也就是第一个成员变量的地址，所以和下面的地址一样
	fmt.Printf("%p\n", &(s.sid))    // 0x14000120040 8个字节 int
	fmt.Printf("%p\n", &(s.name))   // 0x14000120050 16个字节 string
	fmt.Printf("%p\n", &(s.age))    // 0x14000120048 8个字节 int
	fmt.Printf("%p\n", &(s.course)) // 0x14000120060 切片24个字节
	fmt.Printf(" \n", &(s.xxx))     // 0x14000120078 8个字节 int 这个是为了看出来上面的切片占用的内存

}
