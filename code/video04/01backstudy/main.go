package main

import "fmt"

func main() {
	//var x = 10
	//p := &x
	//*p = 20
	//fmt.Println(x)
	//fmt.Println(*p)

	var x = [4]int{1, 2, 3, 4}
	fmt.Printf("%p\n", &x) // 数组地址其实是指向第一个元素的地址，然后根据连续性，后面的元素地址是连续的
	fmt.Println(&x[0])
	fmt.Println(&x[1])
	fmt.Println(&x[2])
	fmt.Println(&x[3])

	var arr = []int{10, 11, 12, 13, 14}
	s1 := arr[0:3]
	s2 := arr[2:5]
	s3 := s2[0:2]

	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(s3, len(s3), cap(s3))

	var emps = make([]string, 3, 5)
	emps[0] = "张三"
	emps[1] = "李四"
	emps[2] = "王五"
	fmt.Println(emps) // [z,l,w]
	emps2 := append(emps, "rain")
	fmt.Println(emps2) // [z,l,w,r]
	emps3 := append(emps, "elric")
	fmt.Println(emps3) // [z,l,w,e]
	fmt.Println(emps2) // [z,l,w,e]
	emps3 = append(emps3, "elric2")
	emps3 = append(emps3, "elric3")
	fmt.Println(emps3, len(emps3), cap(emps3))
	emps3[0] = "zhangsan"
	fmt.Println(emps3)
	fmt.Println(emps)
}
