package main

import "fmt"

// 接口中没有定义任何方法，该接口就是一个空接口
// 所以任何类型都实现了空接口 --> 空接口可以存储任意类型的数值

// 空接口一般不需要提前定义，直接使用即可，下面仅为展示
type xxx interface{}

// 空接口类型的应用
// 1. 空接口作为函数的参数（比如Println函数参数就是interface{},可以打印任何值。｜TODO 这里Println函数的参数变为了any
// 2. 空接口作为map的值

func main() {
	var x interface{}
	x = "hello"
	x = 100
	x = true
	fmt.Println(x)

	m := make(map[interface{}]interface{})
	m2 := make(map[any]any) // any是go1.18新加的类型，可以存储任意类型的值 type any = interface{}
	m1 := map[interface{}]interface{}{"a": 1, "b": "2", 3: 3}
	m["name"] = "小王子"
	m["age"] = 18
	m["hobby"] = []interface{}{"篮球", "足球", "双色球", 1, 2, 3, true, false, 1.1, 2.2, 3.3}
	m[20] = 20
	m2["name"] = "小王"
	fmt.Println(m)
	fmt.Println(m1)
	fmt.Println(m2)

	// 类型断言
	ret := x.(bool)
	fmt.Println(ret)    //true
	ret2, ok := x.(int) // 猜的类型不对时候，ok为false，ret2为对应类型的零值.(int类型的零值为0
	if !ok {
		fmt.Println("类型断言失败", ret2)
	} else {
		fmt.Println("类型断言成功", ret2)
	}

	// 使用switch case进行类型断言
	switch x.(type) {
	case string:
		fmt.Println("x是string类型")
	case int:
		fmt.Println("x是int类型")
	case bool:
		fmt.Println("x是bool类型")
	default:
		fmt.Println("x是其他类型")
	}

}
