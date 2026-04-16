package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func FileUpload(c *gin.Context) {
	/*	fileHeader, err := c.FormFile("file")
		fmt.Println(fileHeader.Filename)
		fmt.Println(fileHeader.Size)
		c.SaveUploadedFile(fileHeader, "uploads/"+fileHeader.Filename)
		fmt.Println(err)*/

	form, _ := c.MultipartForm()
	for key, headers := range form.File {
		fmt.Println(key, headers)
		for _, header := range headers {
			c.SaveUploadedFile(header, "uploads/"+header.Filename)
		}
	}
}
func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/", FileUpload)
	r.Run(":8080")
}
