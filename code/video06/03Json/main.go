package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var s = []int{1, 2, 3, 4, 5}
	var d = map[string]string{"name": "guocheng", "age": "18"}

	// 序列化
	data1, _ := json.Marshal(s)
	data2, _ := json.Marshal(d)
	fmt.Println(string(data1))
	fmt.Println(string(data2))

	// 反序列化
	var s1 []int
	var d1 map[string]string
	_ = json.Unmarshal(data1, &s1)
	_ = json.Unmarshal(data2, &d1)
	fmt.Println(s1)
	fmt.Println(d1)

}
