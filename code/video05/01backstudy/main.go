package main

import (
	"fmt"
)

func main() {
	var names = [3]string{"yuan", "rain", "alvin"}
	var names2 = map[string]string{"name": "yuan", "age": "31"}
	fmt.Println(names)
	fmt.Println(names2)
	names3 := names2 // map是一个引用类型，赋值后指向同一个内存地址，所以names2的改变会影响names3
	names2["gender"] = "male"
	fmt.Println(names3) // [age:31 gender:male name:yuan]

	// map增删改查
	data := make(map[string]string)
	data["name"] = "alex" // 添加键值对
	data["age"] = "18"
	n, exist := data["name2"]
	fmt.Println(n, exist)

	data["name"] = "guo" // 修改键值对
	fmt.Println(data)

	for k, v := range data { // 遍历map
		fmt.Println(k, v)
	}

	// map嵌套
	data2 := make(map[string]map[string]string)
	data2["user1"] = make(map[string]string)
	data2["user1"]["name"] = "yuan"
	data2["user1"]["age"] = "31"
	fmt.Println(data2)

	var x = [3]int{1, 2, 3}
	var y = [4]string{"a", "b", "c", "d"}
	fmt.Println(x, y)

	var d = make(map[string][]string)
	d["shanxi"] = []string{"长治", "太原"}
	d["henan"] = []string{"郑州", "洛阳"}
	d["beijing"] = []string{"朝阳", "海淀"}
	fmt.Println(d)

	// 案例1
	var x1 = 10
	fmt.Printf("x的地址%p\n", &x1)
	y1 := x1
	fmt.Printf("y的地址%p\n", &y1)
	x1 = 100
	fmt.Println(y1)

	// 案例2
	var a = []int{1, 2, 3}
	b := a
	a[0] = 100
	fmt.Println(b)

	// 案例3
	var m1 = map[string]string{"name": "yuan"}
	var m2 = m1
	m2["age"] = "22"
	fmt.Println(m1)
}
