package admin

import (
	"GoGIN/itying/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

// 此层为控制层
type AdminController struct {
	BaseController
}

func (admin AdminController) AdminAccount(c *gin.Context) {
	admin.Success(c, nil)
	fmt.Println("Access Admin Account")
	models.Tool1()
	midValue := c.GetString("mid")
	fmt.Println("midValue:", midValue)
}
