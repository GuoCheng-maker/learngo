package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x      int
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwlock sync.RWMutex
)

func read() {
	defer wg.Done()
	// rwlock.RLock()  // 读锁使用读写锁就可以（其他读也可以，写不行）
	lock.Lock()
	time.Sleep(time.Millisecond)
	// rwlock.RUnlock()
	lock.Unlock()
}

func write() {
	defer wg.Done()
	// rwlock.Lock() // 写锁使用互斥锁就可以（其他读也不行，写也不行）
	lock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond * 5)
	// rwlock.Unlock()
	lock.Unlock()
}

func main() {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	wg.Wait()
	fmt.Println(x)
	fmt.Println(time.Now().Sub(start)) // 不使用读写锁 1.223551667s
	//fmt.Println(time.Now().Sub(start)) // 使用读写锁 60.176833ms

}
