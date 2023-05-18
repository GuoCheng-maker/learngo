package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [1024]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Printf("error:%v\n", err)
			break
		}
		receiver := string(buf[:n])
		fmt.Printf("receiver info: %v\n", receiver)
		conn.Write([]byte("ok"))
	}

}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("error:%v\n", err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("error:%v\n", err)
			continue
		}
		go process(conn)

	}
}
