package main

import (
	"encoding/json"
	"fmt"
	"goStudyProject/day7/calc"
)

type Addr struct {
	Province string
	City     string
}

// Stu 标识符: 变量名、函数名、结构体名、接口名、常量名、包名
// go语言中的标识符是区分大小写的，如果标识符首字母是大写，则表示对外可见（暴露的，公有的），否则表示对外不可见
type Stu struct { // 结构体中变量如果首字母大写，则可以被外部包访问，否则只能在本包访问
	Name string `json:"name"` // 结构体的标签
	Age  int    `json:"-"`    // 表示不参与序列化
	Addr Addr
}

func main() {

	var stuMap = map[string]interface{}{"name": "yuan", "age": 32, "addr": "beijing"}
	var stuStruct = Stu{Name: "yuan", Age: 18, Addr: Addr{Province: "Hebei", City: "langFang"}}

	// 序列化
	jsonStuMap, _ := json.Marshal(stuMap)
	jsonStuStruct, _ := json.Marshal(stuStruct)

	fmt.Println(string(jsonStuMap))
	fmt.Println(string(jsonStuStruct))

	// 反序列化
	// var x  = make(map[int]string)
	var StuMap map[string]interface{}
	err1 := json.Unmarshal(jsonStuMap, &StuMap)
	if err1 != nil {
		return
	}
	fmt.Println("StuMap", StuMap, StuMap["name"])

	var StuStruct Stu
	err2 := json.Unmarshal(jsonStuStruct, &StuStruct)
	if err2 != nil {
		return
	}
	fmt.Println(StuStruct)
	fmt.Println(StuStruct.Name)
	fmt.Println(StuStruct.Addr.City)

	var s1 = Stu{Name: "s1", Age: 18, Addr: Addr{Province: "Hebei", City: "langFang"}}
	var s2 = Stu{Name: "s2", Age: 18, Addr: Addr{Province: "Hebei", City: "langFang"}}
	var s3 = Stu{Name: "s3", Age: 18, Addr: Addr{Province: "Hebei", City: "langFang"}}

	data := []Stu{s1, s2, s3}
	// 序列化
	jsonData, _ := json.Marshal(data)
	fmt.Println(string(jsonData))

	var data2 = make([]Stu, 0)
	err3 := json.Unmarshal(jsonData, &data2)
	if err3 != nil {
		return
	}
	fmt.Println(data2, data2[0].Name)

	calc.Sum(11, 2)

}
