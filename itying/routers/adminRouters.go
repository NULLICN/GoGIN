package routers

import (
	"GoGIN/itying/controller/admin"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	adminGroup := r.Group("admin")

	// 添加路由级的session中间件
	//store := cookie.NewStore([]byte("secret")) // secret是加密密钥，可以自定义
	store, err := redis.NewStore(10, "tcp", "47.109.80.234:6379", "", "040412", []byte("secret"))
	if err != nil {
		fmt.Println(err)
		return
	}

	adminGroup.Use(sessions.Sessions("mysession", store))

	//adminGroup.Use(middlewares.InitMiddlewares) // 配置路由级别的中间件
	adminGroup.GET("/account", admin.AdminController{}.AdminAccount)
	adminGroup.POST("/upload", admin.AdminController{}.AdminUploadFiles)
	adminGroup.GET("/session", admin.AdminController{}.AdminSession)
}

func middleware(c *gin.Context) {
	fmt.Println("---Enter middleware---")
	c.Set("mid", "来自全局中间件传递的值")
	c.Next()
	fmt.Println("---Left Context---")
	mid2Value := c.GetString("mid2")
	fmt.Println("mid2Value:", mid2Value)
}

/*func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(middleware) // 全局中间件
	AdminRoutersInit(r)
	fmt.Println("Server is running on port 8080.")
	r.Run(":8080")
}*/
