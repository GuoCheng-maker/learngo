package main

func foo1() {
	// 这个函数中第一次i=0，只能走到case ch <- i，第二次在循环i=1时候，case x := <-ch: 会被执行，因为ch中有值=0，循环下去：2，4，6，8
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			println("received from ch", x)
		case ch <- i:
			println("sent to ch", i)
		default:
			println("default")
		}
	}
}

func foo2() {
	ch := make(chan int, 10) // 开了10的缓冲区大小以后，select中，如果符合多个case条件，会随机选择一个case执行，所以foo2()每次结果都不同
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			println("received from ch", x)
		case ch <- i:
			println("sent to ch", i)
		default:
			println("default")
		}
	}
}

func main() {
	// foo1()
	foo2()
}
