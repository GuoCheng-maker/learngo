package main

import "fmt"

func main() {

	//for i := 0; i < 100; i++ {
	//	if i == 6 {
	//		break
	//	}
	//	fmt.Println(i) // 到6直接退出
	//}
	//fmt.Println("end")
	//
	//for i := 0; i < 100; i++ {
	//	if i == 6 {
	//		continue
	//	}
	//	fmt.Println(i) // 6 不打印，后续还继续输出
	//}
	//fmt.Println("end")
	var a int8 = 100
	var b int16 = 200
	fmt.Println(int16(a) + b)

	var sum int
	for n := 0; n < 100; n++ {
		if n == 88 {
			continue
		}
		if n%2 == 0 {
			sum -= n
		} else {
			sum += n
		}
		fmt.Println("sum", sum)
	}
	fmt.Println(sum)

	for g := 0; g < 100; g = g + 24 {
		fmt.Println(g)
	}

	var x int
	if x%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}
	cmd := ""
	cmd += "l"
	cmd += "s"
	fmt.Println(cmd)

	var v = 33
	var name = "yuan"
	fmt.Printf("该数字的十进制:%d\n", v)  // 该数字的十进制:33
	fmt.Printf("该数字的二进制:%b\n", v)  // 该数字的二进制:100001   从右往左，依次每一位乘2的N次方 1*2^5+0*2^4+0*2^3+0*2^2+0*2^1+1*2^0=33
	fmt.Printf("该数字的八进制:%o\n", v)  // 该数字的八进制:41  从右往左，依次每一位乘8的N次方 4*8^1+1*8^0=33
	fmt.Printf("该数字的十六进制:%x\n", v) // 该数字的十六进制:21  从右往左，依次每一位乘16的N次方 2*16^1+1*16^0=33

	var salary = 3045.3293243
	fmt.Printf("工资:%.2f\n", salary)
	fmt.Printf("类型:%T\n", salary)

	var q = 2 > 1
	fmt.Printf("b的值：%t\n", q)

	fmt.Printf("任意类型的打印：%#v\n", q)
	fmt.Printf("任意类型的打印：%#v\n", salary)
	fmt.Printf("任意类型的打印：%#v\n", name)
	fmt.Printf("内存地址：%p\n", &name)

}
