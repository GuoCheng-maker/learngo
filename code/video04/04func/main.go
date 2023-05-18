package main

import (
	"fmt"
	"reflect"
)

// 函数声明
func printLing(n int) {

	// 打印菱形
	for i := 1; i <= n; i++ {
		for k := 1; k <= n-i; k++ {
			fmt.Print(" ")
		}

		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := n - 1; i >= 1; i-- {
		for k := 1; k <= n-i; k++ {
			fmt.Print(" ")
		}

		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}

func gotoDemo2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
		fmt.Println("123")
	}
	return
	// 标签
breakTag:
	fmt.Println("结束for循环")
}

func sum(nums ...int) {
	//变参函数，可变参数通常要作为函数的最后一个参数。
	fmt.Println("nums", nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func sum2(base int, nums ...int) int {
	// 参数中有...的，就是可变参数。
	// 可变参数接受任意个数的int参数，打包成一个切片([]int） 可变参数必须放在参数列表的最后。
	fmt.Println(base, nums, reflect.TypeOf(base), reflect.TypeOf(nums))
	for index, num := range nums {
		println(index, num)
	}
	sum := base
	for _, v := range nums {
		sum = sum + v
	}
	return sum
}

func sum3(base int, nums ...int) (int, bool) {
	fmt.Println(base, nums)
	for index, num := range nums {
		println(index, num)
	}
	sum := base
	for _, v := range nums {
		sum = sum + v
	}
	return sum, true
}

func sum4(base int, nums ...int) (z int) {
	fmt.Println(base, nums)
	for index, num := range nums {
		println(index, num)
	}
	z = base
	for _, v := range nums {
		z = z + v
	}
	return
}

func foo() {
	var x int
	x = 10
	fmt.Println(x)
}

var x = 999 // 全局变量只能通过var声明

func main() {
	//gotoDemo2()
	//// 打印菱形 函数调用
	//printLing(60)
	//fmt.Println("hello ")
	//// 打印菱形 函数调用
	//printLing(7)
	//sum(12, 23)
	//sum(1, 2, 3, 4)
	ret := sum2(100, 2, 3, 4, 5)
	fmt.Println(ret)
	//ret2, ok := sum3(100, 2, 3, 4, 5)
	//fmt.Println(ret2, ok)
	//
	//sum3(100, 2, 3, 4, 5)
	//x := sum4(100, 2, 3, 4, 5)
	//fmt.Println(x)
	fmt.Println(x)
	foo()
	fmt.Println(x)
	// {}内的都开辟作用域，比如if、for、switch、select、func等
}
