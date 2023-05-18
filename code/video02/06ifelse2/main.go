package main

import "fmt"

func gotoDemo() {
	/*
		goto用的比较少，和fallthrough一样，一般不建议使用
	*/
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
		fmt.Println("123")
	}
	return
	// 标签
breakTag:
	fmt.Println("结束for循环")
}

func main() {
	//var score int
	//println("请输入成绩：")
	//fmt.Scan(&score)

	//if score < 60 {
	//	fmt.Println("不及格")
	//} else if score < 80 {
	//	fmt.Println("良好")
	//} else if score < 90 {
	//	fmt.Println("优秀")
	//} else if score <= 100 {
	//	fmt.Println("完美")
	//} else {
	//	fmt.Println("成绩输入有误")
	//}

	//var week int
	//fmt.Scan(&week)
	//switch week {
	//case 1:
	//	fmt.Println("星期一")
	//case 2:
	//	fmt.Println("星期二")
	//case 3:
	//	fmt.Println("星期三")
	//case 4:
	//	fmt.Println("星期四")
	//case 5:
	//	fmt.Println("星期五")
	//case 6:
	//	fmt.Println("星期六")
	//case 7:
	//	fmt.Println("星期日")
	//default:
	//	fmt.Println("输入有误")
	//}

	//switch n := 7; n {
	//case 1, 3, 5, 7, 9:
	//	println("奇数")
	//case 2, 4, 6, 8, 10:
	//	println("偶数")
	//default:
	//	println("输入有误")
	//}

	//var age int
	//fmt.Scan(&age)
	//switch {
	//case age < 25:
	//	println("年轻")
	//case age < 40:
	//	println("中年")
	//case age < 60:
	//	println("中老年")
	//case age < 80:
	//	println("老年")
	//default:
	//	println("老寿星")
	//}
	gotoDemo()

}
