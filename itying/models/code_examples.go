package models

import (
	"fmt"
	"os"
	"path/filepath"
)

/*
实际代码示例 - 复制即用

================================================
📌 示例 1: 最简单的使用方式
================================================
*/
func Example1_SimpleUsage() {
	// 检查文件类型并自动修正后缀名
	success, fileType, err := CheckFileTypeWithHeader("uploads/photo.jpg", "")
	if err != nil {
		// 处理错误
		return
	}
	// success=true, fileType="jpeg", 如果后缀名错误已自动修正
	fmt.Printf("success: %v, fileType: %s\n", success, fileType)
}

/*
================================================
📌 示例 2: 验证特定文件类型
================================================
*/
func Example2_VerifySpecificType() {
	// 验证文件是否为 JPEG
	isJPEG, detectedType, err := CheckFileTypeWithHeader("uploads/photo.jpg", "jpeg")
	if err != nil {
		return
	}

	if isJPEG {
		// ✓ 文件确实是 JPEG
	} else {
		// ✗ 文件不是 JPEG，实际类型是 detectedType
	}
	fmt.Printf("detectedType: %s\n", detectedType)
}

/*
================================================
📌 示例 3: 仅允许特定文件类型上传
================================================
*/
func Example3_TypeWhitelist() {
	filePath := "uploads/document.pdf"
	allowedTypes := []string{"jpeg", "png", "gif"}

	isAllowed := false
	for _, allowedType := range allowedTypes {
		if ok, _, _ := CheckFileTypeWithHeader(filePath, allowedType); ok {
			isAllowed = true
			break
		}
	}

	if isAllowed {
		// ✓ 文件类型允许
	} else {
		// ✗ 文件类型不允许
	}
}

/*
================================================
📌 示例 4: 完整的上传处理流程
================================================
*/
func Example4_CompleteUploadFlow(uploadPath string) error {
	// 第 1 步：检查文件是否存在
	if _, err := os.Stat(uploadPath); err != nil {
		return fmt.Errorf("文件不存在: %v", err)
	}

	// 第 2 步：检查文件类型
	success, detectedType, err := CheckFileTypeWithHeader(uploadPath, "")
	if err != nil {
		return fmt.Errorf("文件类型检查失败: %v", err)
	}

	// 第 3 步：记录日志
	fmt.Printf("✓ 文件检查成功，类型: %s, success: %v\n", detectedType, success)

	// 第 4 步：可选 - 验证特定类型
	allowedTypes := []string{"jpeg", "png"}
	typeAllowed := false
	for _, allowed := range allowedTypes {
		if ok, _, _ := CheckFileTypeWithHeader(uploadPath, allowed); ok {
			typeAllowed = true
			break
		}
	}

	if !typeAllowed {
		return fmt.Errorf("文件类型不被允许")
	}

	return nil
}

/*
================================================
📌 示例 5: 批量文件检查
================================================
*/
func Example5_BatchFileCheck(filePaths []string) {
	for _, filePath := range filePaths {
		success, detectedType, err := CheckFileTypeWithHeader(filePath, "")
		if err != nil {
			fmt.Printf("❌ %s - 错误: %v\n", filePath, err)
			continue
		}

		if success {
			fmt.Printf("✓ %s - 类型: %s\n", filePath, detectedType)
		}
	}
}

/*
================================================
📌 示例 6: 错误处理最佳实践
================================================
*/
func Example6_BestPracticeErrorHandling(filePath string) {
	success, detectedType, err := CheckFileTypeWithHeader(filePath, "pdf")

	if err != nil {
		switch {
		case err.Error() == "打开文件失败":
			// 文件打开错误
		case err.Error() == "无法识别的文件类型":
			// 无法识别的文件类型
		default:
			// 其他错误
		}
		return
	}

	if success {
		fmt.Printf("✓ 文件是 PDF\n")
	} else {
		fmt.Printf("✗ 文件不是 PDF，实际类型: %s\n", detectedType)
	}
}

/*
================================================
📌 示例 7: 在 Gin 控制器中使用
================================================

import (
	"GoGIN/itying/models"
	"github.com/gin-gonic/gin"
)

type FileController struct {}

func (fc FileController) Upload(c *gin.Context) {
	// 获取上传的文件
	file, _ := c.FormFile("file")

	// 保存文件
	uploadPath := "uploads/" + file.Filename
	c.SaveUploadedFile(file, uploadPath)

	// 检查文件类型
	success, fileType, err := models.CheckFileTypeWithHeader(uploadPath, "")
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 返回结果
	c.JSON(200, gin.H{
		"success": success,
		"fileType": fileType,
		"path": uploadPath,
	})
}

================================================
*/

/*
================================================
📌 示例 8: 响应结构体
================================================
*/
type FileCheckResponse struct {
	Success      bool   `json:"success"`
	DetectedType string `json:"detectedType"`
	OriginalName string `json:"originalName"`
	SavedPath    string `json:"savedPath"`
	Message      string `json:"message"`
	ErrorDetails string `json:"errorDetails,omitempty"`
}

func Example8_ResponseFormat(filePath string) FileCheckResponse {
	success, fileType, err := CheckFileTypeWithHeader(filePath, "")

	resp := FileCheckResponse{
		Success:      success,
		DetectedType: fileType,
		OriginalName: filepath.Base(filePath),
		SavedPath:    filePath,
	}

	if err != nil {
		resp.Message = "检查失败"
		resp.ErrorDetails = err.Error()
	} else if success {
		resp.Message = "文件检查成功"
	}

	return resp
}

/*
================================================
📌 示例 9: 支持的文件类型查询
================================================
*/
func Example9_ListSupportedTypes() {
	types := GetSupportedFileTypes()

	fmt.Println("支持的文件类型:")
	for typeName, typeInfo := range types {
		fmt.Printf("  %s: %v -> %s\n",
			typeName,
			typeInfo.Extensions,
			typeInfo.Description)
	}

	// 或者直接打印
	PrintSupportedFileTypes()
}

/*
================================================
📌 示例 10: 带有重试逻辑的使用
================================================
*/
func Example10_WithRetryLogic(filePath string, maxRetries int) (bool, string, error) {
	var lastErr error

	for i := 0; i < maxRetries; i++ {
		success, fileType, err := CheckFileTypeWithHeader(filePath, "")

		if err == nil {
			return success, fileType, nil
		}

		lastErr = err

		// 可选: 在重试前等待
		// time.Sleep(time.Second)
	}

	return false, "", fmt.Errorf("最多重试 %d 次后仍失败: %v", maxRetries, lastErr)
}

//================================================
//使用提示：
//
//1. 导入模块
//   import "GoGIN/itying/models"
//
//2. 基本调用
//   success, fileType, err := models.CheckFileTypeWithHeader(path, "")
//
//3. 错误处理
//   if err != nil { /* 处理错误 */  }
//
//4. 类型验证
//   ok, detected, _ := models.CheckFileTypeWithHeader(path, "jpeg")
//   if ok { /* 是 JPEG */ }
//
//5. 支持的类型参数值
//   "jpeg", "png", "gif", "bmp", "pdf", "zip", "rar", "7z", "exe", "txt"
//
//================================================
