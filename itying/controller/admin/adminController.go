package admin

import (
	"GoGIN/itying/models"
	"fmt"

	"github.com/gin-contrib/sessions"
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
	fmt.Println("form Values：", form.Value)
	var filesName []string
	for key, headers := range form.File {
		fmt.Println(key, headers)
		for _, header := range headers {
			uploadPath := "uploads/" + header.Filename
			if err := c.SaveUploadedFile(header, uploadPath); err != nil {
				fmt.Printf("保存文件失败: %v\n", err)
				continue
			}

			// 检查文件类型（检查所有常见的图片和文档格式）
			// 设置期望类型为空，只验证和修正后缀名
			success, detectedType, err := models.CheckFileTypeWithHeader(uploadPath, "")
			if err != nil {
				fmt.Printf("文件类型检查失败: %v\n", err)
			} else {
				fmt.Printf("文件 %s 检测类型: %s, 验证结果: %v\n", header.Filename, detectedType, success)
			}

			filesName = append(filesName, header.Filename)
		}
	}
	admin.Success(c, filesName)
}

func (admin AdminController) AdminSession(c *gin.Context) {
	session := sessions.Default(c)

	if session.Get("hello") != "world" {
		session.Set("hello", "world")
		session.Save()
	}

	respondData := map[string]interface{}{
		"hello": session.Get("hello"),
	}

	admin.Success(c, respondData)
}
