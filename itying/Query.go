package main

import "github.com/gin-gonic/gin"

func GetQuery(c *gin.Context) {
	v := c.Query("username")
	c.JSON(200, gin.H{"username": v})
}

func Query() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/getQuery", GetQuery) // 获取GET数据 /?k=v&k1=v1
	r.Run(":8080")
}
