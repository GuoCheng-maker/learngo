package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	UserName string `form:"username" json:"user_name"`
	PassWord string `form:"password" json:"pass_word"`
}

func getParams(c *gin.Context) {
	var u UserInfo
	err := c.ShouldBind(&u)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		fmt.Printf("%v", u) // {guocheng 123456}
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	}

}

func getJson(c *gin.Context) {
	var u UserInfo
	// ShouldBind会按照下面的顺序解析请求中的数据完成绑定：
	//如果是 GET 请求，只使用 Form 绑定引擎（query）。
	//如果是 POST 请求，首先检查 content-type 是否为 JSON 或 XML，然后再使用 Form（form-data）。
	err := c.ShouldBind(&u)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		fmt.Printf("%v", u) // {guocheng 123456}
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	}

}

func main() {
	r := gin.Default()
	r.GET("/user", getParams)
	r.POST("/form", getParams)
	r.POST("/json", getJson)
	r.Run()
}
