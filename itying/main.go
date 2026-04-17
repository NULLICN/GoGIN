package main

import (
	"GoGIN/itying/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	routers.ApiRoutersInit(r)
	routers.UserRoutersInit(r)
	r.Run(":8080")
}
