package main

import "fmt"

func main() {

	var s = 0
	for count := 0; count <= 100; count++ {
		if count%2 == 0 {
			fmt.Println(count)
			s += count
		}

	}
	fmt.Println(s)
	fmt.Println("========")

	//var num int
	//fmt.Scan(&num)
	//if num > 100 {
	//	// 打印1~numm
	//	for i := 1; i <= num; i++ {
	//		fmt.Println(i)
	//	}
	//} else {
	//	// 打印num~1
	//	for i := num; i > 0; i-- {
	//		fmt.Println(i)
	//	}
	//}

	// 类似于while True
	//for true {
	//	var num int
	//	fmt.Scan(&num)
	//	if num > 100 {
	//		// 打印1~numm
	//		for i := 1; i <= num; i++ {
	//			fmt.Println(i)
	//		}
	//	} else {
	//		// 打印num~1
	//		for i := num; i > 0; i-- {
	//			fmt.Println(i)
	//		}
	//	}
	//}

	fmt.Println("==================================")
	// goto可以直接跳转到标签处，在先面这个双层循环中可以直接跳出2层循环，而break只能跳出一层循环.但是goto会导致程序的可读性变差，不建议使用
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		for n := 0; n < 6; n++ {
			fmt.Println(n)
			if n == 5 {
				goto end
			}
		}
	}
end:
	fmt.Println("end!!!")

}
