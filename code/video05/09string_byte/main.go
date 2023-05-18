package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var x uint8
	//x = 25
	//fmt.Println(x)

	var y byte                        // byte字节类型：声明字符   (type byte = uint8 byte是uint8的别名) 1个字节
	y = 'a'                           // 双引号表示字符串，单引号表示字符
	fmt.Println(y, reflect.TypeOf(y)) // uint8
	fmt.Printf("%c\n", y)             // \c表示字符
	fmt.Printf("%d\n", y)             // \d表示十进制
	fmt.Printf("%x\n", y)             // \x表示十六进制
	fmt.Printf("%b\n", y)             // \b表示二进制

	fmt.Println("=====================================")
	var z rune // rune字符类型：声明字符   (type rune = int32 rune是int32的别名) 4个字节
	z = '苑'
	fmt.Println(z, reflect.TypeOf(z)) // int32
	fmt.Println(string(z))

	fmt.Printf("字符'苑'unicode的十进制：%d\n", z)  // \d表示十进制
	fmt.Printf("字符'苑'unicode的十六进制：%x\n", z) // \x表示十六进制
	fmt.Printf("字符'苑'unicode的二进制：%b\n", z)  // \b表示二进制

	fmt.Println("=====================================")
	// go语⾔的string是⼀种数据类型，这个数据类型占⽤16字节空间，前8字节是⼀个指针，指向字符串值的地址，后⼋个字节是⼀个整数，标识字符串的长度
	s := "苑abc" //底层数组长啥样 []byte{'\xe8', '\x8b', '\x91', 'a', 'b', 'c'}, 对应的十进制是[232, 139, 145, 97, 98, 99]
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i]) // s中存储了苑abc的utf-8对应的字节，因为中文占3个字节，所以苑打印了3个出来，具体每一个字节的值是根据utf-8的编码规则来的，而英文是1个字节，所以打印出来的是97，98，99
	}

	fmt.Println("=====================================")
	for i, v := range s {
		fmt.Println(i, v) // 打印的是字符的unicode码
		fmt.Println(i, string(v))
	}
	// len(s) 和range(s) 的区别是 len(s)是字节的长度，range(s)是字符的长度
	// rune和byte类型的区别 rune是int32的别名，byte是uint8的别名 ，byte对应的是ASCII码，rune对应的是unicode编码表的符号

	fmt.Println("=====================================")
	var msg = "abc郭成"
	// 字符串转字节串
	b := []byte(msg)
	c := []rune(msg)
	fmt.Println(b) // [97 98 99 233 131 173 230 136 144] 每一个元素都是一个字符的utf-8码，占1个字节
	fmt.Println(c) // [97 98 99 37101 25104] 每一个元素都是一个字符的unicode码，占4个字节

	// 字节串转字符串
	s2 := string(b)
	s3 := []byte{97, 98, 99, 232, 139, 145}
	fmt.Println(s2)         // abc郭成
	fmt.Println(string(s3)) // abc郭成
}
