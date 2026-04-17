package admin

import "github.com/gin-gonic/gin"

type BaseController struct{}

func (base BaseController) Success(c *gin.Context, data interface{}) {
	//c.JSON(200, gin.H{"message": "Success", "data": data})
	c.JSON(200, "成功")
}

func (base BaseController) Fail(c *gin.Context, data interface{}) {
	//c.JSON(200, gin.H{"message": "Fail", "data": data})
	c.JSON(200, "失败")
}
