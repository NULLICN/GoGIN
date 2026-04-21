package routers

import (
	"GoGIN/itying/controller/admin"

	"github.com/gin-gonic/gin"
)

func UserRoutersInit(r *gin.Engine) {
	userGroup := r.Group("user")
	userGroup.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	userGroup.GET("/index", admin.UserController{}.Index)
	userGroup.POST("/add", admin.UserController{}.Add)
}
