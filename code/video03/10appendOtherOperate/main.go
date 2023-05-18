package main

import "fmt"

func main() {
	var s = []int{1, 2, 3}
	fmt.Printf("%p\n", &s)
	s = append(s, 4)
	fmt.Printf("%p\n", &s) // 这里s的地址是不变的，虽然append进行了扩容，但是扩容后再次赋给了s. 所以s的地址是不变的

	// 向开头插入元素
	var a = []int{1, 2, 3}
	n := append([]int{0}, a...) // 如果在开头插入，则生成一个只有0一个元素的切片，把a...(a...的意思是把切片解成一个个的元素，append进新切片)追加到切片后面
	fmt.Println(n)

	// 任意位置插入元素
	var b = []int{1, 2, 3, 4, 5}
	i := 2
	n1 := append(b[:i], append([]int{100}, b[i:]...)...) // 在索引为2的位置插入100 变为[1 2 100 3 4 5]
	fmt.Println(n1)
	// ...的意思是把切片解成一个个的元素，append进新切片

	// 删除元素
	var c = []int{1, 2, 3, 4, 5}
	c = append(c[:2], c[2+1:]...) // 删除索引为2的元素
	fmt.Println(c)
}
