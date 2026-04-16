package main

import (
	"github.com/gin-gonic/gin"
)

func IndexFILE(c *gin.Context) {

}
func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/", IndexFILE)
	r.Run(":8080")
}
