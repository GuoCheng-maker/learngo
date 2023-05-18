package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"sync"
)

// fetchUrlDemo 并发获取url内容
func fetchUrlDemo() {
	wg := sync.WaitGroup{}
	var urls = []string{
		"http://pkg.go.dev",
		"http://www.liwenzhou.com",
		"http://www.yixieqitawangzhi.com",
	}

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err == nil {
				fmt.Printf("获取%s成功\n", url)
				resp.Body.Close()
			}
			return // 如何将错误返回呢？
		}(url)
	}
	wg.Wait()
	// 如何获取goroutine中可能出现的错误呢？
}

// fetchUrlDemo2 使用errgroup并发获取url内容
func fetchUrlDemo2() error {
	g := new(errgroup.Group) // 创建等待组（类似sync.WaitGroup）
	var urls = []string{
		"http://pkg.go.dev",
		"http://www.liwenzhou.com",
		"http://www.yixieqitawangzhi.com",
		"http://www.baidu.com",
	}
	for _, url := range urls {
		fmt.Println("====", url)
		url := url // 注意此处声明新的变量
		//在 fetchUrlDemo2() 中，由于 url 是在 for 循环的作用域内声明的，因此每次迭代都会创建一个新的变量 url。
		//而在循环内部创建的 goroutine 会在某个时间点异步执行，并且该 goroutine 会在后续的某个时间点才会访问 url 变量的值。
		//在这种情况下，如果不使用 url := url 的方式为每个 goroutine 创建一个新的变量，那么所有 goroutine 将会共享同一个变量 url 的副本，
		//而 url 变量的值将会在 for 循环中被多次修改。
		//假设在第一个 goroutine 执行之前，for 循环已经执行到了第二个迭代，那么此时 url 变量的值已经被更新为第二个 URL。
		//在这种情况下，第一个 goroutine 在访问 url 变量的值时，实际上获得的是第二个 URL 的值，而不是第一个 URL 的值，这将导致执行 http.Get() 请求时访问的 URL 不正确。
		//因此，使用 url := url 的方式为每个 goroutine 创建一个新的变量，可以避免多个 goroutine 之间对 url 变量的竞态条件，保证每个 goroutine 获得的是自己的副本，而不会出现互相干扰的情况。

		// 启动一个goroutine去获取url内容
		g.Go(func() error { // 传入的参数是一个（返回err类型的函数）
			resp, err := http.Get(url)
			if err == nil {
				fmt.Printf("获取%s成功\n", url)
				resp.Body.Close()
			}
			return err // 返回错误
		})
	}
	err := g.Wait()
	if err != nil {
		// 处理可能出现的错误
		fmt.Println(err)
		return err
	}
	fmt.Println("所有goroutine均成功")
	return nil
}

func main() {
	// fetchUrlDemo()
	fetchUrlDemo2()
}
