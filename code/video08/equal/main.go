package main

import (
	"fmt"
	"reflect"
)

type data struct {
}

func main() {
	// 在 Go 中，== 操作符通常用于比较两个值是否相等。它对于基本类型（如 int、float、bool 等）的值比较非常直观，但对于复合类型（如切片、映射、接口等）则需要注意其特殊的行为。
	//对于切片、映射和接口类型，== 操作符只有在它们都是 nil 或者它们共享相同的底层结构时才会返回 true，否则都会返回 false。这种行为是由 Go 语言规范所定义的。
	//而 reflect.DeepEqual 函数则可以对任意两个值进行深度比较，它会递归地比较它们的成员变量或元素，直到找到基本类型的数据，然后对它们进行相等性比较。因此，它可以比较复杂类型的值，而不仅仅是判断它们是否共享相同的底层结构。
	//需要注意的是，由于 reflect.DeepEqual 函数需要进行深度比较，因此在性能上可能会比 == 操作符慢，尤其是对于大型的数据结构。因此，在需要进行简单的值比较时，应该优先考虑使用 == 操作符。而在需要进行复杂的值比较时，可以使用 reflect.DeepEqual 函数。
	v1 := data{}
	v2 := data{}
	fmt.Println("v1 == v2:", reflect.DeepEqual(v1, v2)) //prints: v1 == v2: true
	m1 := map[string]string{"one": "a", "two": "b"}
	m2 := map[string]string{"two": "b", "one": "a"}
	fmt.Println("m1 == m2:", reflect.DeepEqual(m1, m2)) //prints: m1 == m2: true
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	fmt.Println("s1 == s2:", reflect.DeepEqual(s1, s2)) //prints: s1 == s2: true

	// reflect.DeepEqual 函数可以用于深度比较两个变量的值，它会递归地比较它们的成员变量或元素，直到找到基本类型的数据，然后对它们进行相等性比较。
	//具体来说，它会比较两个变量的类型是否相同，如果类型不同，则不相等；如果类型相同，则按以下规则比较它们的值：
	//如果是指针类型，则比较它们的地址是否相同。
	//如果是接口类型，则比较它们的动态类型和动态值是否相同。
	//如果是数组、切片、映射类型，则比较它们的长度和每个元素的值是否相同。
	//如果是结构体类型，则比较它们的成员变量是否相同。
	//如果是基本类型（如 int、float、bool 等），则直接比较它们的值是否相同。
	//因此，对于上述代码中的每个变量，reflect.DeepEqual 函数都会比较它们的值是否相等，而在这些例子中，变量的值都是相等的，因此函数的返回值都是 true。
	n1 := 100
	n2 := 100
	fmt.Println("n1 == n2:", reflect.DeepEqual(n1, n2), n1 == n2) //prints: n1 == n2: true
}
