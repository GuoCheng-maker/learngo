package main

import (
	"fmt"
	"strconv"
)

func main() {
	// map嵌套slice
	var data = make(map[string][]string)
	data["beijing"] = []string{"chaoyang", "haidian", "fengtai"}
	data["shanghai"] = []string{"pudong", "minhang", "baoshan"}
	data["guangzhou"] = []string{"tianhe", "yuexiu", "haizhu"}
	fmt.Println(data["guangzhou"][1])
	for province, cities := range data {
		fmt.Println(province, len(cities))
		for _, city := range cities {
			fmt.Printf(city)
		}
		fmt.Println()
	}
	// 删除北京的key-value
	delete(data, "beijing")

	// 格式化输出
	//stu01 := map[string]map[string]string{"name": {"a": "q"}}
	//fmt.Println(stu01)
	//marshal, _ := json.Marshal(stu01)
	//fmt.Println(string(marshal)) // {"name":{"a":"q"}}

	// map嵌套map
	//stu01 := map[string]string{"name": "rain", "age": "32"}
	//stu02 := map[string]string{"name": "elric", "age": "33"}
	//stu03 := map[string]string{"name": "yuan", "age": "34"}
	//var stus = make(map[int]map[string]string)
	//stus[1000] = stu01
	//stus[1001] = stu02
	//stus[1002] = stu03
	//fmt.Println(stus)
	//fmt.Println(stus[1002]["age"])

	stu01 := map[string]string{"name": "rain", "age": "32"}
	stu02 := map[string]string{"name": "elric", "age": "33"}
	stu03 := map[string]string{"name": "yuan", "age": "34"}
	// 01
	//var slice1 = make([]map[string]string, 3)
	//slice1[0] = stu01
	//slice1[1] = stu02
	//slice1[2] = stu03

	// 02
	slice2 := []map[string]string{stu01, stu02, stu03}
	slice3 := []map[string]string{{"name": "rain", "age": "32"}, {"name": "elric", "age": "33"}, {"name": "yuan", "age": "34"}}
	fmt.Println(slice2, slice3)

	// 新增一个学生map
	newMap := map[string]string{"name": "new", "age": "18"}
	slice2 = append(slice2, newMap)
	fmt.Println(slice2)

	// 删除一个学生map
	//slice2 = append(slice2[:1], slice2[2:]...)
	// fmt.Println(slice2)

	// 删除学生elric的map
	var deleteIndex = 0
	for i, stuMap := range slice2 {
		if stuMap["name"] == "elric" {
			deleteIndex = i
		}
	}
	fmt.Println(deleteIndex)
	slice2 = append(slice2[:deleteIndex], slice2[deleteIndex+1:]...)
	fmt.Println(slice2)

	// 学生rain的age加1
	for _, stuMap := range slice2 {
		if stuMap["name"] == "rain" {
			age, _ := strconv.Atoi(stuMap["age"])
			age++
			stuMap["age"] = strconv.Itoa(age)
		}
	}
	fmt.Println(slice2)

}
