package main

import (
	"GinStudy/respond/res"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	resp := res.Respond{
		Code: 0,
		Msg:  "成功",
		Data: 114514,
	}
	//resp.OkWithData(c, resp.Data)
	resp.FailWithCode(c, 1002)
}
func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/", Index)
	r.Run(":8080")
}
