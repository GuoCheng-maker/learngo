package main

import (
	"fmt"
	"reflect"
)

// 类型包含属性和方法

// Student 声明结构体， 结构体是自定义数据类型，也是一种数据类型
type Student struct {
	sid, age int //字段名必须唯一
	name     string
	course   []string
	xxx      int
}

func main() {
	var s1 = new(Student)           // &Student .new()开辟了空间，并将这个空间的地址返回给了s1
	fmt.Println(reflect.TypeOf(s1)) //*main.Student s是一个指针类型
	s1.name = "rain"
	fmt.Println(*s1)     //{0 0 rain [] 0} 注意这里想要看结构体的内容，需要*s1,因为s1是一个指针类型，并不是一个结构体类型
	fmt.Println(s1.name) // rain

	var s2 = make([]int, 3)         //  make函数开辟了空间，并初始化了一个切片，返回的是一个切片
	fmt.Println(reflect.TypeOf(s2)) //[]int s2是一个切片类型

}
