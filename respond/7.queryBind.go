package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func QueryBind(c *gin.Context) {
	type User struct {
		Name string `form:"name"`
		Id   int    `form:"id"`
	}
	var user User
	err := c.ShouldBindQuery(&user)
	fmt.Println("Query:", user, err)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
			"type":  "Query",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Query bind success",
		"data":    user,
	})
}
func PathBind(c *gin.Context) {
	type User struct {
		Name string `uri:"name"`
		Id   int    `uri:"id"`
	}
	var user User
	err := c.ShouldBindUri(&user)
	fmt.Println("Path:", user, err)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
			"type":  "Path",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Path bind success",
		"data":    user,
	})
}
func FormBind(c *gin.Context) {
	type User struct {
		Name string `form:"name"`
		Id   int    `form:"id"`
	}
	var user User
	err := c.ShouldBind(&user)
	fmt.Println("Json:", user, err)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
			"type":  "Json",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Form bind success",
		"data":    user,
	})
}
func JsonBind(c *gin.Context) {
	type User struct {
		Name string `json:"name"`
		Id   int    `json:"id"`
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
func HeadersBind(c *gin.Context) {
	type User struct {
		Name      string `header:"name"`
		Id        int    `header:"id"`
		UserAgent string `header:"User-Agent"`
	}
	var user User
	err := c.ShouldBindHeader(&user)
	fmt.Println("Json:", user, err)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
			"type":  "Json",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Headers bind success",
		"data":    user,
	})
}
func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/", QueryBind)              // 查询参数
	r.GET("/path/:id/:name", PathBind) // 路径参数
	r.POST("/form", FormBind)          // 表单参数
	r.POST("/json", JsonBind)          // json参数
	r.GET("/headers", HeadersBind)     // headers参数
	r.Run(":8080")
}
