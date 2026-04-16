package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ParamsValidate(c *gin.Context) {
	type User struct {
		Name string `json:"name" binding:"required,min=4"` // binding属性条件：字符串不为""，最短长度4字符
		Id   int    `json:"id"`
		List []any  `json:"list"`
	}
	var user User
	err := c.ShouldBindJSON(&user)
	fmt.Println("Json:", user, err)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
			"type":  "Json",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Json bind success",
		"data":    user,
	})
}
func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/paramsValidate", ParamsValidate)
	r.Run(":8080")
}
