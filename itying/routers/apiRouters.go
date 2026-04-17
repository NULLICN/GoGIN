package routers

import "github.com/gin-gonic/gin"

func ApiRoutersInit(r *gin.Engine) {
	apiGroup := r.Group("api")
	apiGroup.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
}
