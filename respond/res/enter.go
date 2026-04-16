package res

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Respond struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

var failCode = map[int]string{
	1001: "权限错误",
	1002: "角色错误",
}

func response(c *gin.Context, code int, msg string, data any) {
	c.JSON(200, Respond{code, msg, data})
}
func (resp *Respond) OK(c *gin.Context, msg string, data any) {
	response(c, 0, msg, data)
}

func (resp *Respond) OkWithData(c *gin.Context, data any) {
	resp.OK(c, "成功", data)
}

func (resp *Respond) OkWithMsg(c *gin.Context, msg string) {
	resp.OK(c, msg, gin.H{})
}

// fail
func (resp *Respond) Fail(c *gin.Context, code int, msg string, data any) {
	response(c, code, msg, data)
}

func (resp *Respond) FailWithCode(c *gin.Context, code int) {
	msg, ok := failCode[code]

	fmt.Println(msg, ok)
	if ok {
		response(c, code, msg, nil)
	}
}
func (resp *Respond) FailWithMsg(c *gin.Context, msg string) {
	response(c, 1001, msg, nil)
}
