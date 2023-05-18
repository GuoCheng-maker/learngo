package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

func parse(r io.Reader) (*Point, error) {
	// 在parse函数中，如果在调用多个read函数时出现了多个错误，那么在函数返回时，err变量将保存最近一次出现的错误信息。
	//这意味着返回的错误信息将是最近一次出现的错误，而不是最后一个解析失败的字段的错误信息。
	var p Point
	var err error
	read := func(data interface{}) {
		// 打印传入的参数
		fmt.Printf("%T\n", data)
		if err != nil {
			return
		}
		err = binary.Read(r, binary.BigEndian, data)
		fmt.Println("123", err)
	}

	read(&p.Longitude)
	read(&p.Latitude)
	read(&p.Distance)
	read(&p.ElevationGain)
	read(&p.ElevationLoss)

	if err != nil {
		return &p, err
	}
	return &p, nil
}

type Point struct {
	Longitude     float64
	Latitude      float64
	Distance      float64
	ElevationGain float64
	ElevationLoss float64
}

func main() {
	// 构造字节缓冲区
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, float64(123.456)) // 写入经度
	binary.Write(&buf, binary.BigEndian, float64(45.678))  // 写入纬度
	//binary.Write(&buf, binary.BigEndian, float64(32.234))  // 写入距离
	//binary.Write(&buf, binary.BigEndian, float64(65.32)) // 写入上升高度
	//binary.Write(&buf, binary.BigEndian, float64(75))    // 写入下降高度

	point, err := parse(&buf)
	fmt.Println(point)
	if err != nil {
		fmt.Printf("err: %s", err)
		return
	}
}
