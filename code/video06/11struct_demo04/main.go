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

func initSid(s Student) {
	s.sid = 0
}

func trueInitSid(p *Student) {
	fmt.Printf("%p\n", &*p) //0x14000120040
	(*p).sid = 0            // 这里必须加括号提升优先级，否则不能间接调用，会报错

	/*
		在Go语言中，结构体指针的变量可以继续使用'.'
		这是因为Go语言为了方便开发者访问结构体指针的成员变量可以像访问结构体的成员变量一样简单
		使用了语法糖（Syntactic sugar）技术，将 instance.Name 形式转换为 (*instance).Name。
		意味这上面的代码可以简写为 p.sid = 0
	*/

}

func main() {
	// k-v赋值 ，字段xxx不赋值也没关系，而且kv形式不用按照顺序
	var s1 = Student{name: "郭成", sid: 1001, course: []string{"语文", "数学", "英语"}}
	fmt.Println(s1)
	s2 := s1 // 值拷贝
	s2.age = 30
	fmt.Println(s1.age) // 0

	initSid(s1)         // 这里把s1拷贝给了形参s，形参的sid改为了0，但是s1的sid还是1001，s1和s是两个不同的变量
	fmt.Println(s1.sid) // 1001

	fmt.Printf("%p\n", &s1) //0x14000120040
	trueInitSid(&s1)        // 这里把s1的地址传给了指针p，*p的sid改为了0，s1的sid也改为了0，s1和*p是同一个地址
	fmt.Println(s1.sid)

	// 第三种实例化方式
	var sp = &Student{name: "郭成", sid: 1001, course: []string{"语文", "数学", "英语"}}
	fmt.Println(sp) // 直接把地址拿出来作为指针，后续直接操作指针，这个用于修改同一块数据时候使用。相当于把上面的trueInitSid(&s1)简化了
}
