package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name" ini:"s_name"`
	Age  int    `json:"age" ini:"s_age"`
}

// Study 给student添加两个方法 Study和Sleep(注意首字母大写)
func (p Person) Study(x int) string {
	fmt.Println(x)
	msg := p.Name + "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func (p Person) Sleep(x int) string {
	fmt.Println(x)
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

func (p Person) Run(x int) string {
	fmt.Println(x)
	msg := "好好跑步，强身健体。"
	fmt.Println(msg)
	return msg
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{reflect.ValueOf(i)}
		v.Method(i).Call(args)
	}

	methodObj := v.MethodByName("Run")
	fmt.Println(methodObj.Type())
	methodObj.Call([]reflect.Value{reflect.ValueOf(199)})

}

func main() {
	p := Person{Name: "郭成", Age: 198}
	//
	stu1 := reflect.TypeOf(p)
	fmt.Printf("name:%v kind:%v\n", stu1.Name(), stu1.Kind())
	for i := 0; i < stu1.NumField(); i++ {
		// 通过索引取出结构体的每一个字段
		fieldObj := stu1.Field(i)
		fmt.Printf("name:%v type:%v tag:%v\n", fieldObj.Name, fieldObj.Type, fieldObj.Tag)
		fmt.Printf("json tag:%v\n", fieldObj.Tag.Get("json"))
		fmt.Printf("ini tag:%v\n", fieldObj.Tag.Get("ini"))
	}
	// 根据字段名字获取
	fieldObj2, _ := stu1.FieldByName("Age")
	fmt.Printf("name:%v type:%v tag:%v\n", fieldObj2.Name, fieldObj2.Type, fieldObj2.Tag)

	fmt.Println("=====================================")
	printMethod(p)
}
