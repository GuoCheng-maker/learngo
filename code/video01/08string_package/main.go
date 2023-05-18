package main

import (
	"fmt"
	"strings"
)

func main() {
	var name = "Yuan"
	var upperName = strings.ToUpper(name)
	var lowerName = strings.ToLower(name)
	fmt.Println(upperName, lowerName)

	var s = "rain yuan alivin"
	fmt.Println(strings.HasPrefix(s, "rai")) // bool
	fmt.Println(strings.HasSuffix(s, "in"))  // bool
	fmt.Println(strings.Contains(s, "an "))  // bool

	userName := " yu  an   "
	fmt.Println(strings.Trim(userName, " "), "|")       // 去掉字符串两边的空格，不会去除字符串中间的空格
	fmt.Println(strings.TrimLeft(userName, " "), "|")   // 去掉字符串左边的空格
	fmt.Println(strings.TrimRight(userName, " "), "|")  // 去掉字符串右边的空格
	fmt.Println(strings.TrimSpace(userName), "|")       // 去掉字符串的空格
	fmt.Println(strings.TrimLeft(userName, "  y"), "|") // 去掉字符串左边的空格和y
	fmt.Println(strings.TrimRight(userName, "n "), "|") // 去掉字符串右边的n和空格

	// index索引
	fmt.Println(strings.Index(userName, "yu")) // 2,不存在返回-1

	// 小作业： 有一条mysql命令， cmd  = mysql 127.0.0.1 -u root -p 123, 从cmd中获取用户名和密码
	var cmd = "mysql 127.0.0.1 -u root -p 123"
	fmt.Println(cmd[1:3]) // ys 从索引1开始，到索引3结束，不包含索引3
	var index1 = strings.Index(cmd, "-u")
	var index2 = strings.Index(cmd, "-p")
	var user = strings.TrimSpace(cmd[index1+2 : index2])
	var pwd = strings.TrimSpace(cmd[index2+2:])
	fmt.Println(user, pwd)

	// 分割与拼接
	var s3 = "rain yuan alivin"
	fmt.Println(strings.Split(s3, " "))                    // [rain yuan alivin]
	fmt.Println(strings.Join(strings.Split(s3, " "), "-")) // rain-yuan-alivin
}
