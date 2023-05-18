package main

import (
	"fmt"
	"time"
)

// 取1-100的数字，放入到一个ch1中
func f1(ch1 chan<- int) {
	for i := 0; i < 100; i++ {
		ch1 <- i
		// x := <-ch1 不能这么做，因为这是一个单向通道，只能往通道send
	}
	close(ch1)
}

// 从ch1中取出数据，并计算其平方，放入到ch2中
func f2(ch1 <-chan int, ch2 chan<- int) {
	for {
		/*
			f2函数在循环中使用tmp和ok变量来检查通道是否已关闭。在这种情况下，如果通道已关闭，则ok变量将为false，
			而tmp变量将为通道元素类型的零值在这种情况下，循环将退出，因为f1不再向通道中写入任何值。
		*/
		tmp, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- tmp * tmp
	}
	close(ch2)
}

func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)

	go f1(ch1)
	go f2(ch1, ch2)
	time.Sleep(time.Second * 4)
	// 从ch2中取出结果并打印
	for ret := range ch2 {
		fmt.Println(ret)
	}
}
