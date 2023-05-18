package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
)

func queryIP(ip string, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()

	url := "http://qifu-api.baidubce.com/ip/geo/v1/district"
	resp, err := http.Get(fmt.Sprintf("%s?ip=%s", url, ip))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	ch <- string(body)
}

func testGin(wg *sync.WaitGroup) {
	defer wg.Done()
	url := "http://127.0.0.1:8080/shop/s2"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

//func main() {
//	rand.Seed(time.Now().UnixNano())
//
//	var wg sync.WaitGroup
//	ch := make(chan string, 5000)
//
//	for i := 0; i < 5000; i++ {
//		wg.Add(1)
//		go queryIP(fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256)), &wg, ch)
//	}
//
//	go func() {
//		wg.Wait()
//		close(ch)
//	}()
//
//	for res := range ch {
//		fmt.Println("返回信息：" + res)
//	}
//}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go testGin(&wg)
	}
	wg.Wait()
}
