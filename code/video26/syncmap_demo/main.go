package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var m = make(map[int]int)
var m2 = sync.Map{} // key和value都是空接口类型，不用初始化也不用指定类型

func get(k int) int {
	return m[k]
}

func set(k, v int) {
	m[k] = v
}

//func main() {
//	// fatal error: concurrent map writes
//	for i := 0; i < 10; i++ {
//		wg.Add(1)
//		go func(i int) {
//			set(i, i+100)
//			fmt.Printf("key :%v, value: %v\n", i, get(i))
//			wg.Done()
//		}(i)
//	}
//	wg.Wait()
//}

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			m2.Store(i, i+100)
			value, _ := m2.Load(i)
			fmt.Printf("key :%v, value: %v\n", i, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
