package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.OpenFile("满江红2", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file)

	// 按字节写或字符串写
	file.Write([]byte("满江红666\n"))
	file.WriteString("满江红！\n")

	// 缓冲写
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(strconv.Itoa(i) + "\n")
	}
	writer.Flush() // 将缓冲区的内容写入文件，这样可以提升写文件的效率

	// 写整个文件
	str := "满江红666\n满江红！\n"
	os.WriteFile("满江红3", []byte(str), 0666)

}
