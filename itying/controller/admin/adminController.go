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

func (admin AdminController) AdminUploadFiles(c *gin.Context) {
	form, _ := c.MultipartForm()
	var filesName []string
	for key, headers := range form.File {
		fmt.Println(key, headers)
		for _, header := range headers {
			c.SaveUploadedFile(header, "uploads/"+header.Filename)
			filesName = append(filesName, header.Filename)
		}
	}
	admin.Success(c, filesName)
}
