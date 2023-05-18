package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func upload(c *gin.Context) {
	fileReader, err := c.FormFile("f1")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "Failed"})
	} else {
		dst := fmt.Sprintf("./%s", fileReader.Filename)
		err := c.SaveUploadedFile(fileReader, dst)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"status": "Failed"})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": "OK"})
		}

	}
}

func main() {
	r := gin.Default()
	r.POST("/upload", upload)
	r.Run()
}
