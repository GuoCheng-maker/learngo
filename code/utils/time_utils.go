package utils

import (
	"fmt"
	"time"
)

func GetTime() string {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	t := time.Now().Format("2006-01-02 15:04:05")
	return t
}
