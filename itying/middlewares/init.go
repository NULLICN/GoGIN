package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// 此层为中间层
func InitMiddlewares(c *gin.Context) {
	fmt.Println("InitMiddlewares from package middlewares")
	fmt.Println("Request path is:", c.Request.URL.Path)
	c.Next()
	c.Set("mid2", "来自二层中间件传递的值")

	cCp := c.Copy()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("InitMiddlewares from goroutine:" + cCp.Request.URL.Path)
	}()
}
