package routers

import "github.com/gin-gonic/gin"

func UserRoutersInit(r *gin.Engine) {
	userGroup := r.Group("user")
	userGroup.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
}
