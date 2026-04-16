package main

import (
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	//c.JSON(200, gin.H{"content-Type": "text/html"})
	/*_, err := c.Writer.WriteString(
		`<h1>
			<p>hello NULLICN！</p>
		</h1>
		`,
	)
	if err != nil {
		return
	}*/
	c.Header("Content-Type", "text/html")
	c.String(200, "<h1><p>hello NULLICN!!!</p></h1>")
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/", Index)
	r.Run(":8080")
}
