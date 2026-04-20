package main

import (
	"GoGIN/itying/routers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	//r.Use(middlewares.InitMiddlewares) // 全局级中间件
	routers.ApiRoutersInit(r)
	routers.UserRoutersInit(r)
	routers.AdminRoutersInit(r)
	fmt.Println("Server is running on port 8080.")
	r.Run(":8080")
}
