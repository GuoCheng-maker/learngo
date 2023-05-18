package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getParams(c *gin.Context) {
	username := c.Param("username")
	password := c.Param("password")
	c.JSON(http.StatusOK, gin.H{"user_name": username, "pass_word": password})
}

func main() {
	r := gin.Default()
	// 注意这里的url的匹配不要冲突，如果有2个url 一个是/:username/:password，一个是/blog/:year/:month，则就会发生冲突,可以改成/user/:username/:password
	r.GET("/user/:username/:password", getParams) // localhost:8080/user/hahaha/123 { "pass_word": "123", "user_name": "hahaha"}
	r.Run()
}
