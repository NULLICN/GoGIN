package main

import (
	"bytes"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func readOriginalContent(c *gin.Context) {
	byteData, _ := io.ReadAll(c.Request.Body)
	name := c.PostForm("name")
	fmt.Println(string(byteData))
	c.Request.Body = io.NopCloser(bytes.NewBuffer(byteData))
	fmt.Println(name)
}
func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/", readOriginalContent)
	r.Run(":8080")
}
