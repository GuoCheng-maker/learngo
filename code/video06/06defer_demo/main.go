package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readByLine1(file *os.File) {
	reader := bufio.NewReader(file)
	for true {
		lineContent, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(lineContent)
	}
}

func test1() int {
	x := 10
	defer func() {
		x = 100
	}()
	return x // rval=10.   x = 100, ret rval
}

func test2() []int {
	x := []int{1, 2, 3}
	defer func() {
		x[0] = 999
	}()
	return x // rval=切片->（根据首地址，长度，容量找到真正的数组）   x[0] = 999，修改了数组, ret rval [999, 2,3]
}

func main() {
	//fmt.Println("1")
	//defer fmt.Println("2")
	//fmt.Println("3")

	// 案例2
	//f, _ := os.Open("/Users/guocheng/goWork/src/goStudyProject/day6/03fileWrite/满江红")
	//defer f.Close() // 延迟关闭文件
	//readByLine1(f)

	// 案例3
	//fmt.Println("1")
	//defer fmt.Println("2") // 相比于34行，这里的defer语句会最后执行，规则是谁先注册，谁最后执行
	//fmt.Println("3")
	//defer fmt.Println("4")
	//fmt.Println("5")

	// 案例4
	//foo := func() {
	//	fmt.Println("i am foo1")
	//}
	//defer foo() // 这里foo函数保留的是上面的函数体，下面对foo修改没用，相当于把foo函数体复制了一份出来，等待执行
	//
	//foo = func() {
	//	fmt.Println("i am foo2")
	//}

	// 案例5
	//x := 10
	//defer func(a int) {
	//	fmt.Println(a)  // 保留a的地址，这个时候a = x = 10
	//}(x)  // 10  这里的形参a开辟了空间，将x=10赋值了进去，等待执行，后面x++不影响函数最后结果
	//x++

	// 案例6
	//x := 10
	//defer func() {
	//	fmt.Println(x) //保留x的地址
	//}() // 11 这里函数体执行的x时候，地址指向了外面的x，等待延迟执行，函数外的x++对x这个地址的值进行了修改，所以最后结果是11
	//x++

	//rval = xxx
	//defer_func
	//ret rval

	// return在汇编语言中分为了2步，第一步是将返回值放到寄存器rval中，第二步是跳转到ret指令返回rval
	// defer是在这两句中间执行，如果这个时候修改了xxx,那么不会影响返回值，因为值已经给了rval.
	fmt.Println(test1()) // 10
	fmt.Println(test2()) // [999, 2, 3]
}
