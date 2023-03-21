package main

import "sync"

var icons map[string]string
var loadIconsOnce sync.Once

func loadIcons() {
	icons = map[string]string{
		"left":  "left.png",
		"up":    "up.png",
		"right": "right.png",
		"down":  "down.png",
	}
}

// Icon 被多个goroutine调用时不是并发安全的
func Icon(name string) string {
	// 这个是不安全的，因为多个goroutine可能同时执行到这里，其中某个goroutine刚刚初始化，还没来得及把图片放到map里，其他goroutine就来了
	//此时不是nil, 但是可能部分图片没来得及放入到map,导致其他goroutine可能拿不到图片
	//if icons == nil {
	//	loadIcons()
	//}
	loadIconsOnce.Do(loadIcons) // 用sync.Once保证只执行一次。
	return icons[name]
}

func main() {

}
