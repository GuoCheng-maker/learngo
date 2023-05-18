package main

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
)

func read(readPath string) (io.Reader, error) {
	r, err := os.Open(readPath)
	//错误包装
	if err != nil {
		return nil, errors.Wrap(err, "read failed")
	}
	return r, nil
}

func main() {
	_, err := read("a")
	if err != nil {
		fmt.Println(err)
	}
}
