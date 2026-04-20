package models

import (
	"fmt"
)

/*
这个文件展示了如何在实际项目中使用 CheckFileTypeWithHeader 函数

在 AdminUploadFiles 中的集成示例：

func (admin AdminController) AdminUploadFiles(c *gin.Context) {
	form, _ := c.MultipartForm()
	var filesName []string
	var validFiles []string

	for key, headers := range form.File {
		fmt.Println(key, headers)
		for _, header := range headers {
			uploadPath := "uploads/" + header.Filename
			c.SaveUploadedFile(header, uploadPath)

			// 方案 1: 只检查和修正后缀名，不验证特定类型
			success, detectedType, err := CheckFileTypeWithHeader(uploadPath, "")
			if err != nil {
				fmt.Printf("文件类型检查失败: %v\n", err)
				continue
			}
			fmt.Printf("✓ 文件 %s 检测类型: %s\n", header.Filename, detectedType)

			// 方案 2: 检查文件是否为允许的图片格式
			allowedTypes := []string{"jpeg", "png", "gif", "bmp"}
			isAllowed := false
			for _, allowedType := range allowedTypes {
				ok, _, _ := CheckFileTypeWithHeader(uploadPath, allowedType)
				if ok {
					isAllowed = true
					validFiles = append(validFiles, header.Filename)
					break
				}
			}
			if !isAllowed {
				fmt.Printf("✗ 文件 %s 不是允许的图片格式\n", header.Filename)
				continue
			}

			filesName = append(filesName, header.Filename)
		}
	}
	admin.Success(c, filesName)
}
*/

// DemonstrateUsage 展示文件类型检查功能的用法
func DemonstrateUsage() {
	fmt.Println("=== 文件类型检查功能演示 ===\n")

	// 打印支持的文件类型
	PrintSupportedFileTypes()

	fmt.Println("\n=== 使用场景 ===")
	fmt.Println("1. 检查并自动修正文件后缀名")
	fmt.Println("   CheckFileTypeWithHeader(\"uploads/test.jpg\", \"\")")
	fmt.Println("   - 读取文件头检测实际类型")
	fmt.Println("   - 如果后缀名不匹配，自动修改")
	fmt.Println()

	fmt.Println("2. 验证文件是否为期望的类型")
	fmt.Println("   CheckFileTypeWithHeader(\"uploads/photo.jpg\", \"jpeg\")")
	fmt.Println("   - 检测文件类型是否为 JPEG")
	fmt.Println("   - 返回 true 如果是 JPEG，否则 false")
	fmt.Println()

	fmt.Println("3. 在文件上传时进行验证和修正")
	fmt.Println("   - 保存上传的文件")
	fmt.Println("   - 调用 CheckFileTypeWithHeader 检查类型")
	fmt.Println("   - 自动修正错误的后缀名")
	fmt.Println("   - 可选：验证文件是否为允许的类型")
	fmt.Println()

	fmt.Println("=== 安全性提示 ===")
	fmt.Println("✓ 不仅检查文件后缀名，而是读取文件头（magic bytes）")
	fmt.Println("✓ 防止用户上传恶意文件（例如改后缀名的病毒文件）")
	fmt.Println("✓ 自动修正错误命名的文件")
	fmt.Println("✓ 支持多种常见文件格式")
}
