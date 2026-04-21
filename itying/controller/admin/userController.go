package admin

import (
	"GoGIN/itying/models"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (con UserController) Add(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)

	fmt.Println("Json:", user, err)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
			"type":  "Json",
		})
		return
	}
	// 后端自动添加当前时间
	user.AddTime = time.Now()

	DB := models.DB
	err = DB.Create(&user).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
			"type":  "Database",
		})
		return
	}
	checkData := []models.User{} // 这里可能会返回多个同id的数据
	DB.Where("id = ?", user.Id).Find(&checkData)
	c.JSON(200, gin.H{
		"message": "Json bind success",
		"data":    checkData,
	})
}

func (con UserController) Index(c *gin.Context) {
	DB := models.DB
	userList := models.User{}
	DB.Find(&userList)
	fmt.Println("userList:", userList)
	c.JSON(200, gin.H{"result": userList})
}
