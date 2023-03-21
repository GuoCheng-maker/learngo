package main

import (
	"fmt"
	"reflect"
	"time"
)

// 闭包 返回值是一个带有time.Duration返回值类型的匿名函数
func executionTime() func() time.Duration {
	start := time.Now()

	return func() time.Duration {
		return time.Since(start)
	}
}

func main() {
	now := time.Now() // 时间对象
	fmt.Println(now, "|", reflect.TypeOf(now))
	year := now.Year()           // 年
	month := now.Month()         // 月
	day := now.Day()             // 日
	hour := now.Hour()           // 时
	minute := now.Minute()       // 分
	second := now.Second()       // 秒
	timeStamp1 := now.Unix()     // 时间戳（秒数）
	timeStamp2 := now.UnixNano() // 时间戳（纳秒数）
	fmt.Println(year, month, day, hour, minute, second)
	fmt.Println(timeStamp1, timeStamp2)

	// 将时间戳转换为具体的时间
	t := time.Unix(timeStamp1, 0)
	fmt.Println(t, "|", reflect.TypeOf(t))

	//时间间隔
	// n := 2
	// time.Sleep(time.Second * 2)
	// time.Sleep(time.Second * time.Duration(n)) // 这里的time.Duration(n)是将n转换为time.Duration类型，不能直接写time.Second * n

	// 时间add
	t2 := now.Add(time.Hour * 2)
	fmt.Println(t2)

	// 时间sub
	t3 := t2.Sub(now)
	fmt.Println(t3)

	// 时间定时器
	//for tmp := range time.Tick(time.Second) {
	//	fmt.Println(tmp)
	//}
	// 时间格式化 go语言的时间格式化必须是固定的，不能随意写必须是2006-01-02 15:04:05(也可以记为2006，12345，这里的3是下午3点)
	fmt.Println(now.Format("2006-01-02 15:04:05"))         // 2023-03-17 10:24:54
	fmt.Println(now.Format("2006/01/02 15:04:05"))         // 24小时制
	fmt.Println(now.Format("2006/01/02 15:04:05 PM"))      // 12小时制
	fmt.Println(now.Format("2006/01/02 15:04:05 Monday"))  // 加星期
	fmt.Println(now.Format("2006/01/02 15:04:05 January")) // 加月份
	fmt.Println(now.Format("2006/01/02 15:04:05.000"))     // 毫秒 2023/03/17 10:24:54.000

	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(loc)

	timeStr := "2023-03-18 10:24:54"
	// 解析一个字符串格式的时间
	timeObj1, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	//根据时区去解析一个字符串格式的时间

	timeObj2, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(timeObj1) // 2023-03-18 10:24:54 +0000 UTC
	fmt.Println(timeObj2) // 2023-03-18 10:24:54 +0800 CST
}
