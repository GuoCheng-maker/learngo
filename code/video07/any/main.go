package main

import "fmt"

// reverse 只支持int类型，如果想要支持更多类型，可以使用下面的泛型
func reverse(s []int) []int {
	l := len(s)
	r := make([]int, l)

	for i, e := range s {
		r[l-i-1] = e
	}
	fmt.Println(r)
	return r
}

// reverseIF 支持int和float64类型
func reverseIF[T int | float64](s []T) []T {
	l := len(s)
	r := make([]T, l)
	for i, e := range s {
		r[l-i-1] = e
	}
	fmt.Println(r)
	return r
}

// reverseAny 支持任意类型
func reverseAny[T any](s []T) []T {
	l := len(s)
	r := make([]T, l)
	for i, e := range s {
		r[l-i-1] = e
	}
	fmt.Println(r)
	return r
}

func MapKeys[Key comparable, Val any](m map[Key]Val) []Key {
	s := make([]Key, 0, len(m))
	for k := range m {
		s = append(s, k)
	}
	return s
}

func MapKeysInt[Key int, Val any](m map[Key]Val) []Key {
	n := make([]Key, 0, len(m))
	for k := range m {
		n = append(n, k)
	}
	return n
}

func main() {
	reverse([]int{1, 2, 3, 4, 5})
	reverseIF([]int{1, 2, 3, 4, 5})
	reverseIF([]float64{1, 2, 3, 4, 5})
	// reverseIF([]float32{1, 2, 3, 4, 5})  // 泛型中只对int和float64类型做了处理，所以这里会报错，也可以使用any类型
	reverseAny([]string{"1", "2", "3", "4", "5"})
	reverseAny([]float32{1.1, 2.2, 3.3, 4.4, 5.5})
	m := map[int]interface{}{1: true, 2: "2", 3: "3.00"}
	ret := MapKeys(m)
	fmt.Println(ret)

	ret2 := MapKeysInt(m)
	fmt.Println(ret2)
}
