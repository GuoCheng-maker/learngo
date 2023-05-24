// package main
//
// import (
//
//	"crypto/sha256"
//	"fmt"
//	"math"
//	"runtime"
//	"strconv"
//	"strings"
//	"sync"
//	"time"
//
// )
//
// var wg sync.WaitGroup
//
//	func main() {
//		runtime.GOMAXPROCS(20) // 设置使用几个核
//
//		data := "guocheng"
//		n := 1
//		start := time.Now()
//
//		index := int(math.Pow(2, 32))
//		for n < index { // (1 << 32)也表示2的32次方
//			wg.Add(1)
//			inputStr := data + strconv.Itoa(n)
//			hashValue := sha256.Sum256([]byte(inputStr))
//			hashString := fmt.Sprintf("%x", hashValue)
//			hashValue = sha256.Sum256([]byte(hashString))
//			hashString = fmt.Sprintf("%x", hashValue)
//			// fmt.Println(hashString)
//
//			go func(hashString string) {
//				defer wg.Done()
//				if strings.HasPrefix(hashString, "000") {
//					fmt.Println(inputStr, hashString)
//					fmt.Println(time.Since(start).Seconds())
//				}
//				n = n + 1
//			}(hashString)
//			wg.Wait() // 阻塞等待计数器变为0
//		}
//	}
package main

import (
	"crypto/sha256"
	"fmt"
	"math"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var found int32 // 用于标记是否已经找到匹配的哈希值

func main() {
	runtime.GOMAXPROCS(10) // 设置使用几个核

	data := []byte("guocheng")
	n := 1
	start := time.Now()

	index := int(math.Pow(2, 32))
	prefix := "00000" // 哈希前缀
	for n < index {
		if atomic.LoadInt32(&found) == 1 {
			break // 如果已经找到匹配的哈希值，立即退出循环
		}
		wg.Add(1)
		inputStr := append(data, []byte(strconv.Itoa(n))...)
		hashValue := sha256.Sum256(inputStr)
		hashValue = sha256.Sum256(hashValue[:])
		hashString := fmt.Sprintf("%x", hashValue)

		go func(inputStr []byte, hashString string) {
			defer wg.Done()
			if strings.HasPrefix(hashString, prefix) {
				atomic.StoreInt32(&found, 1) // 设置 found 为 1，表示已经找到匹配的哈希值
				fmt.Println(string(inputStr), hashString)
				fmt.Println(time.Since(start).Seconds())
			}
		}(inputStr, hashString)
		n++
	}

	wg.Wait()
}
