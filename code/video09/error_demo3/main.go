package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type Reader struct {
	r   io.Reader
	err error
}

type Point struct {
	Longitude     float64
	Latitude      float64
	Distance      float64
	ElevationGain float64
	ElevationLoss float64
}

func (r *Reader) read(data interface{}) {
	if r.err == nil {
		r.err = binary.Read(r.r, binary.BigEndian, data)
	}
}

func parse(input io.Reader) (*Point, error) {
	var p Point
	r := Reader{r: input, err: nil}

	r.read(&p.Longitude)
	r.read(&p.Latitude)
	r.read(&p.Distance)
	r.read(&p.ElevationGain)
	r.read(&p.ElevationLoss)

	if r.err != nil {
		return nil, r.err
	}

	return &p, nil
}

func main() {
	// 构造字节缓冲区
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, float64(123.456)) // 写入经度
	binary.Write(&buf, binary.BigEndian, float64(45.678))  // 写入纬度
	binary.Write(&buf, binary.BigEndian, float64(32.234))  // 写入距离
	binary.Write(&buf, binary.BigEndian, float64(65.32))   // 写入上升高度
	binary.Write(&buf, binary.BigEndian, float64(75))      // 写入下降高度

	point, err := parse(&buf)
	fmt.Println(point)
	if err != nil {
		println(err)
		return
	}
	println(point)
}
