package main

import (
	gin "github.com/gin-gonic/gin"
	"net/http"
)

func sayHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hello golang!"})
}

func getJson(c *gin.Context) {
	// 方式1. 使用map
	// data := map[string]interface{}{"name": "guocheng", "age": 18, "gender": false}

	// 方式2. 使用gin.H  其实type H map[string]any // H is a shortcut for map[string]interface{}
	// data := gin.H{"name": "guocheng", "age": 18, "gender": false}

	// 方式3. 结构体（结构体里面的字段如果是小写则不会被访问到） 灵活使用tag对字节定制化
	type Data struct {
		Name   string `json:"name"`
		Age    int
		Gender bool
	}
	data := Data{
		Name:   "guocheng",
		Age:    99,
		Gender: false,
	}
	c.JSON(http.StatusOK, data)
}

func getQueryString(c *gin.Context) {
	// name := c.Query("query") // 郭成  http://127.0.0.1:8080/query_string?query=郭成
	name := c.DefaultQuery("query", "哈哈哈") // 哈哈哈  http://127.0.0.1:8080/query_string?qery=郭
	age := c.Query("age")
	c.JSON(http.StatusOK, gin.H{"name": name, "age": age})
}

func main() {
	r := gin.Default()
	r.GET("/hello", sayHello)
	r.GET("/json", getJson)
	r.GET("/query_string", getQueryString)

	r.Run()
}
