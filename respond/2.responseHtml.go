package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "NULLICN"})
}
func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.LoadHTMLGlob("respond/templates/*")
	r.GET("/", IndexHTML)
	r.Run(":8080")
}
