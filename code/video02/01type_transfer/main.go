package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// 整形之间的转换 int8 int16 int32
	var x int8 = 10
	var y int16 = 20
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(int8(y))
	fmt.Println(x + int8(y))

	// 字符串与整型之间的转换
	var ageStr = "32"
	AgeInt, _ := strconv.Atoi(ageStr) // 字符串转整型
	fmt.Println(AgeInt, reflect.TypeOf(AgeInt))

	priceInt := 998
	priceStr := strconv.Itoa(priceInt) // 整型转字符串
	fmt.Println(priceStr, reflect.TypeOf(priceStr))

	// strconv Parse系列函数
	b, _ := strconv.ParseInt("123", 10, 64) // 10进制， 64位
	fmt.Println(b, reflect.TypeOf(b))

	c, _ := strconv.ParseFloat("3.1415926", 64) // 64位
	fmt.Println(c, reflect.TypeOf(c))

	d, _ := strconv.ParseBool("true")
	fmt.Println(d, reflect.TypeOf(d))

	d1, err := strconv.ParseBool("0")
	// 返回字符串表示的bool值。它接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE；否则返回错误。
	fmt.Println(d1, err, reflect.TypeOf(d1))

}
