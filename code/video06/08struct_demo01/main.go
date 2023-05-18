package main

// 类型包含属性和方法

// Student 声明结构体， 结构体是自定义数据类型，也是一种数据类型
type Student struct {
	sid, age int //字段名必须唯一
	name     string
	course   []string
}

type Person struct {
	name, gender, addr string
	age                uint8
}

func main() {

}
