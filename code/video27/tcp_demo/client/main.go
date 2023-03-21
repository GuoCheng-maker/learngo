package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("error:%v\n", err)
	}
	defer conn.Close()
	input := bufio.NewReader(os.Stdin)
	for {
		s, _ := input.ReadString('\n')
		s = strings.TrimSpace(s)
		_, err := conn.Write([]byte(s))
		if err != nil {
			fmt.Printf("send error:%v\n", err)
			return
		}
		var buf [1024]byte
		n, err := conn.Read(buf[:]) // 将从conn读取到到数据放到buf中，并返回读取到到字节多少，也就是n
		if err != nil {
			fmt.Printf("recv error:%v\n", err)
		}
		fmt.Println("收到服务端的消息：", string(buf[:n])) // 取读取到的n个字节的切片（字节流），转换为字符串

	}
}
