package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func UserGroup(c *gin.RouterGroup) {
	c.GET("checkPath", UserView)
}

func UserView(c *gin.Context) {
	path := c.Request.URL.Path
	fmt.Println(path)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	apiGroup := r.Group("api")
	UserGroup(apiGroup)
	r.Run(":8080")
}
