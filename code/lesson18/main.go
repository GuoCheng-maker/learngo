package main

import (
	"fmt"
	"gin_demo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
	"time"
)

func indexHandler(c *gin.Context) {
	time.Sleep(time.Second * 1)
	fmt.Println("index")
	name, _ := c.Get("name")
	c.JSON(http.StatusOK, gin.H{"msg": "OK", "name": name})
}

func timeCostMiddleware1(c *gin.Context) {
	start := time.Now()
	c.Set("name", "m1")
	go func(c *gin.Context) {
		c.Set("name", "update!!!")
	}(c.Copy()) // 这里如果用c,那么goroutine就会把c的set的name值给改掉，这样是并发不安全的，必须用c.Copy()，只能读，不能进行修改操作
	time.Sleep(time.Millisecond * 10)
	fmt.Println("m1 in....")
	// num := rand.Intn(100) // 0-100随机一个数字
	// fmt.Println(num)
	c.Next()
	//if num > 5 {
	//	c.Next() // 调用后续的处理函数
	//} else {
	//	c.Abort() // 阻止调用后续函数
	//}
	cost := time.Since(start)
	fmt.Printf("耗时：%v\n", cost)
	fmt.Println("m1 out....")
}
func timeCostMiddleware2(c *gin.Context) {
	fmt.Println("m2 in....")
	c.Set("name", "guocheng")
	c.Next() // 调用后续的处理函数
	fmt.Println("m2 out....")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 设置多核运行
	fmt.Println(runtime.NumCPU())
	gin.SetMode(gin.ReleaseMode) // 设置gin的模式
	r := gin.Default()           // 默认gin使用了 Logger() Recovery()中间件，如果不想使用可以直接 r := gin.New()
	r.Use(timeCostMiddleware1)   // 全局使用中间件 如果2个中间件，按照顺序执行，但是如果中间件有next或者abord则执行
	// 比如说上面的例子会1. m1 in... 2.m2 in... 3. 执行index 4.m2 out... 5.m1 out...
	// r.GET("/index", timeCostMiddleware, indexHandler)  // 也可以在指定路由中使用中间件
	r.GET("/index", indexHandler)

	xxGroup := r.Group("/shop")
	xxGroup.Use(timeCostMiddleware1)
	{ // 此处括号可加可不加，增加提高阅读性
		xxGroup.GET("/s1", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "OK", "name": "s1"})
		})
		xxGroup.GET("s2", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "OK", "name": "s2", "time": utils.GetTime()})
		})
	}
	r.Run()
}
