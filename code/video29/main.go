package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func redirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://www.liwenzhou.com")
}

func b(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "b"})
}

func main() {
	r := gin.Default()
	r.GET("/redirect", redirect) // http重定向
	r.GET("/a", func(c *gin.Context) {
		c.Request.URL.Path = "/b"
		r.HandleContext(c)
	}) // 路由重定向
	r.GET("/b", b)
	r.Run()
}
