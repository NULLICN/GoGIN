package routers

import (
	"GoGIN/itying/controller/admin"
	"GoGIN/itying/middlewares"
	"fmt"

	"github.com/gin-gonic/gin"
)

func adminRoutersInit(r *gin.Engine) {
	adminGroup := r.Group("admin")
	adminGroup.Use(middlewares.InitMiddlewares) // 配置路由的中间件
	adminGroup.GET("/account", admin.AdminController{}.AdminAccount)
}

func middleware(c *gin.Context) {
	fmt.Println("---Enter middleware---")
	c.Set("mid", "来自全局中间件传递的值")
	c.Next()
	fmt.Println("---Left Context---")
	mid2Value := c.GetString("mid2")
	fmt.Println("mid2Value:", mid2Value)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(middleware) // 全局中间件
	adminRoutersInit(r)
	fmt.Println("Server is running on port 8080.")
	r.Run(":8080")
}
