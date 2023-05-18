package test_demo

import (
	"fmt"
)

func Split(s map[string]interface{}, sep string) []interface{} {
	list := make([]interface{}, 0, 10)
	for k, v := range s {
		fmt.Println(k, v)
		list = append(list, k)
		list = append(list, sep)
	}
	return list
}
