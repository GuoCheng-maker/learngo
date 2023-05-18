package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	// 我是不知道别人调用接口传入的参数是什么类型
	// 1.通过之前学的类型断言
	// 2.方法反射
	obj := reflect.TypeOf(x)
	fmt.Println(obj, "|", obj.Name(), "|", obj.Kind())
	fmt.Printf("%T\n", obj)
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("%v,%T\n", v, v) // 10,reflect.Value
	k := v.Kind()               // 拿到值对应的类型种类
	switch k {
	case reflect.Int64:
		ret := v.Int()
		fmt.Printf("type is int64, value is %d\n", ret)
	case reflect.Float32:
		ret := v.Float()
		fmt.Printf("type is float32, value is %f\n", ret)
	}
}

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	// Elem() 取指针对应的值
	k := v.Elem().Kind()
	fmt.Println(k)
	switch k {
	case reflect.Int64:
		v.Elem().SetInt(200)
	case reflect.Float32:
		v.Elem().SetFloat(3.1415926)
	}
}

type person struct{}
type dog struct{}

func main() {
	reflectType(10)
	var p person
	reflectType(p)
	var d dog
	reflectType(d)
	var e []int
	var f []string
	reflectType(e)
	reflectType(f)

	// valueOf
	//var x float32 = 3.14
	//reflectValue(x)

	// set value
	var a int64 = 99
	reflectSetValue(&a) // 传入的是指针才能修改值，不然修改不了。
	fmt.Println(a)

}
