package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func test(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodPost:
		c.JSON(http.StatusOK, gin.H{"method": "POST"})
	case http.MethodGet:
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	}
}
func noRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "404"})
}

func main() {
	r := gin.Default()
	r.Any("/user", test) // 接受任何method的请求，内部可以进行区分处理
	r.NoRoute(noRoute)   // 没有路由，返回404

	// Group
	videoGroup := r.Group("/video")
	videoGroup.GET("/a", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"video": "a"})
	})
	videoGroup.POST("/b", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"video": "b"})
	})

	shopGroup := r.Group("/shop")
	shopGroup.GET("/a", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"shop": "a"})
	})
	shopGroup.POST("/b", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"shop": "b"})
	})
	r.Run()
}
