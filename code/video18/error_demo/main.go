package main

import (
	"errors"
	"fmt"
)

type Info interface{}

var EOF = errors.New("EOF")

// OpError 自定义结构体类型
type OpError struct {
	Op string
}

// Error OpError 类型实现error接口
func (e *OpError) Error() string {
	return fmt.Sprintf("无权执行%s操作", e.Op)
}

func queryById(id int64) (*Info, error) {
	if id <= 0 {
		// return nil, errors.New("无效的id")
		return nil, EOF
	}
	return nil, nil
}
func main() {
	s, a := queryById(-10)
	fmt.Println(s, a) // <nil> 无效的id

	err := fmt.Errorf("查询数据库失败，err:%w", a)
	fmt.Println(err) // 查询数据库失败，err:EOF

	e := OpError{Op: "delete"}
	s2 := e.Error()
	fmt.Println(s2) // 无权执行delete操作
}
