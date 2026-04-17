package respond

import (
	"github.com/gin-gonic/gin"
)

func ResponseString(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.String(200, "<h1><p>hello NULLICN!!</p></h1>")
}
func ResponseJson(c *gin.Context) {
	c.JSON(200, map[string]interface{}{
		"code": 0,
		"msg":  "success",
	})
}

type UserInfo struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

func ResponseJsonWithStruct(c *gin.Context) {
	c.JSON(200, UserInfo{
		"NULLICN",
		114514,
	})
}
func ResponseHTML(c *gin.Context) {
	content := c.Param("content")
	c.HTML(200, "index.html", gin.H{"title": "NULLICN", "content": content})
}
func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/string", ResponseString)            // 响应文本
	r.GET("/json", ResponseJson)                // 响应json
	r.GET("jsonStruct", ResponseJsonWithStruct) // 用结构体充当响应json
	r.GET("/html/:content", ResponseHTML)       // 响应html模板
	r.Run(":8080")
}
