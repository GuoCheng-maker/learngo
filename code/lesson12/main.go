package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getForm(c *gin.Context) {
	// 方式1
	//username := c.PostForm("username")
	//password := c.PostForm("password")

	// 方式2 获取不到key,则把username赋值为default_un， password赋值为default_pwd
	username := c.DefaultPostForm("username", "default_un")
	password := c.DefaultPostForm("password", "default_pwd")

	// 方式3. 判断是否可以获取到，然后根据"_"（应该用ok）做一些判定
	//username, _ := c.GetPostForm("username")
	//password, _ := c.GetPostForm("password")
	c.JSON(http.StatusOK, gin.H{"user_name": username, "pass_word": password})
}

func main() {
	r := gin.Default()
	r.POST("/form", getForm) // {"pass_word": "123456", "user_name": "guocheng"}
	r.Run()
}
