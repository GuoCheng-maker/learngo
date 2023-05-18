package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readByBytes(file *os.File) string {
	var data = make([]byte, 9)
	file.Read(data)
	return string(data)
}

func readByLine1(file *os.File) {
	reader := bufio.NewReader(file)
	for true {
		lineContent, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(lineContent)
	}
}

func readByLine2(file *os.File) {
	reader := bufio.NewReader(file)
	for {
		line, isPrefix, err := reader.ReadLine()
		if err != nil {
			break
		}
		fmt.Println(string(line))
		// 如果读取到的行没有以换行符结束，则继续读取
		for isPrefix {
			line, isPrefix, err = reader.ReadLine()
			if err != nil {
				break
			}
			fmt.Println(string(line))
		}
	}

}

func main() {
	// 1. 打开文件
	file, err := os.Open("满江红")
	if err != nil {
		fmt.Println("open file failed, err:", err)
	}
	fmt.Println(file)

	// 1.按字节读（用的少）
	//ret := readByBytes(file)
	//fmt.Println(ret)

	// 2.按行读
	// readByLine1(file)

	// readByLine2(file)

	// 3.读整个文件
	content, _ := os.ReadFile("满江红")
	fmt.Println(string(content))

}
