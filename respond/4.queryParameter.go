package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func QueryParameter(c *gin.Context) {
	data := c.DefaultQuery("data", "NULLICN")
	fmt.Println("k=v:", data)
}
func DynamicParamePath(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("path:", id)
}
func PostFormParameter(c *gin.Context) {
	id, ok := c.GetPostForm("id")
	//fmt.Println(id, ok)
	if ok {
		idNum, err := strconv.Atoi(id)
		//fmt.Println(idNum, err)
		if err == nil {
			fmt.Printf("id:%d type:%T", idNum, idNum)
		}
	}
}
func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/", QueryParameter)
	//r.GET("/:id", DynamicParamePath)
	r.POST("/post", PostFormParameter)
	r.Run(":8080")
}
