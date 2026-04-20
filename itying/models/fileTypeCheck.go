package models

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// FileTypeInfo 存储文件类型信息
type FileTypeInfo struct {
	FileType    string   // 文件类型（如 "jpg", "png", "pdf"）
	MagicBytes  [][]byte // 文件头特征字节
	Extensions  []string // 对应的文件扩展名
	Description string   // 文件类型描述
}

// 支持的文件类型定义
var supportedFileTypes = map[string]FileTypeInfo{
	"jpeg": {
		FileType:    "jpeg",
		MagicBytes:  [][]byte{{0xFF, 0xD8, 0xFF}},
		Extensions:  []string{".jpg", ".jpeg"},
		Description: "JPEG Image",
	},
	"png": {
		FileType:    "png",
		MagicBytes:  [][]byte{{0x89, 0x50, 0x4E, 0x47}},
		Extensions:  []string{".png"},
		Description: "PNG Image",
	},
	"gif": {
		FileType:    "gif",
		MagicBytes:  [][]byte{{0x47, 0x49, 0x46, 0x38}}, // GIF87a or GIF89a
		Extensions:  []string{".gif"},
		Description: "GIF Image",
	},
	"bmp": {
		FileType:    "bmp",
		MagicBytes:  [][]byte{{0x42, 0x4D}},
		Extensions:  []string{".bmp"},
		Description: "BMP Image",
	},
	"pdf": {
		FileType:    "pdf",
		MagicBytes:  [][]byte{{0x25, 0x50, 0x44, 0x46}}, // %PDF
		Extensions:  []string{".pdf"},
		Description: "PDF Document",
	},
	"zip": {
		FileType:    "zip",
		MagicBytes:  [][]byte{{0x50, 0x4B, 0x03, 0x04}, {0x50, 0x4B, 0x05, 0x06}, {0x50, 0x4B, 0x07, 0x08}},
		Extensions:  []string{".zip"},
		Description: "ZIP Archive",
	},
	"rar": {
		FileType:    "rar",
		MagicBytes:  [][]byte{{0x52, 0x61, 0x72, 0x21}}, // Rar!
		Extensions:  []string{".rar"},
		Description: "RAR Archive",
	},
	"7z": {
		FileType:    "7z",
		MagicBytes:  [][]byte{{0x37, 0x7A, 0xBC, 0xAF, 0x27, 0x1C}}, // 7z¼¯'
		Extensions:  []string{".7z"},
		Description: "7Z Archive",
	},
	"exe": {
		FileType:    "exe",
		MagicBytes:  [][]byte{{0x4D, 0x5A}}, // MZ
		Extensions:  []string{".exe", ".dll"},
		Description: "Executable File",
	},
	"txt": {
		FileType:    "txt",
		MagicBytes:  [][]byte{}, // 文本文件没有固定的文件头
		Extensions:  []string{".txt"},
		Description: "Text File",
	},
}

// CheckFileTypeWithHeader 检查文件类型，基于文件头和期望的类型
// filePath: 文件路径
// expectedType: 期望的文件类型（如 "jpg", "png" 等），为空时只检查并修正后缀名，不验证类型
// 返回: (是否符合期望类型, 实际检测到的类型, 错误信息)
func CheckFileTypeWithHeader(filePath string, expectedType string) (bool, string, error) {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return false, "", fmt.Errorf("打开文件失败: %w", err)
	}
	defer file.Close()

	// 读取文件头（最多读前512字节用于检测）
	headerBytes := make([]byte, 512)
	n, err := file.Read(headerBytes)
	if err != nil {
		return false, "", fmt.Errorf("读取文件头失败: %w", err)
	}
	headerBytes = headerBytes[:n]

	// 获取当前文件扩展名
	currentExt := strings.ToLower(filepath.Ext(filePath))
	if currentExt == "" {
		currentExt = ".unknown"
	}

	// 检测文件类型
	detectedType := detectFileType(headerBytes)

	// 如果没有检测到类型，返回错误
	if detectedType == "" {
		return false, "", fmt.Errorf("无法识别的文件类型")
	}

	// 检查和修正后缀名
	correctExt := getCorrectExtension(detectedType)
	if !strings.EqualFold(currentExt, correctExt) {
		// 后缀名不匹配，自动修改
		newFilePath := strings.TrimSuffix(filePath, currentExt) + correctExt
		err := renameFileWithRetry(filePath, newFilePath, 3)
		if err != nil {
			return false, detectedType, fmt.Errorf("修改文件后缀名失败: %w", err)
		}
		fmt.Printf("文件后缀名已从 %s 修改为 %s\n", currentExt, correctExt)
	}

	// 如果指定了期望的类型，进行验证
	if expectedType != "" {
		expectedType = strings.ToLower(expectedType)
		detectedType = strings.ToLower(detectedType)

		// 检查实际类型是否符合期望类型
		if detectedType == expectedType {
			return true, detectedType, nil
		}
		return false, detectedType, nil
	}

	// 如果没有指定期望类型，检查成功
	return true, detectedType, nil
}

// detectFileType 根据文件头检测文件类型
func detectFileType(headerBytes []byte) string {
	if len(headerBytes) == 0 {
		return ""
	}

	// 遍历所有支持的文件类型
	for _, fileTypeInfo := range supportedFileTypes {
		for _, magicBytes := range fileTypeInfo.MagicBytes {
			if len(headerBytes) >= len(magicBytes) {
				// 比较文件头
				if bytesMatch(headerBytes[:len(magicBytes)], magicBytes) {
					return fileTypeInfo.FileType
				}
			}
		}
	}

	// 如果文件头为空或无特征，尝试根据内容判断（如纯文本文件）
	// 检查是否为纯文本文件
	if isPureText(headerBytes) {
		return "txt"
	}

	return ""
}

// bytesMatch 比较字节切片
func bytesMatch(data, pattern []byte) bool {
	if len(data) < len(pattern) {
		return false
	}
	for i := range pattern {
		if data[i] != pattern[i] {
			return false
		}
	}
	return true
}

// isPureText 判断文件是否为纯文本
func isPureText(headerBytes []byte) bool {
	if len(headerBytes) == 0 {
		return true
	}

	// 检查是否包含过多的非文本字符
	nonTextCount := 0
	for _, b := range headerBytes {
		// 检查是否为控制字符（除了常见的文本控制字符）
		if b < 32 && b != 9 && b != 10 && b != 13 { // tab, newline, carriage return
			nonTextCount++
		}
		if b > 127 { // 扩展ASCII
			// 可能是UTF-8或其他编码
			continue
		}
	}

	// 如果非文本字符少于5%，认为是文本文件
	threshold := len(headerBytes) / 20 // 5%
	return nonTextCount < threshold
}

// getCorrectExtension 获取正确的文件扩展名
func getCorrectExtension(fileType string) string {
	fileType = strings.ToLower(fileType)
	if info, exists := supportedFileTypes[fileType]; exists {
		if len(info.Extensions) > 0 {
			return info.Extensions[0] // 返回主要扩展名
		}
	}
	return ".unknown"
}

// GetSupportedFileTypes 获取所有支持的文件类型列表
func GetSupportedFileTypes() map[string]FileTypeInfo {
	return supportedFileTypes
}

// PrintSupportedFileTypes 打印所有支持的文件类型
func PrintSupportedFileTypes() {
	fmt.Println("=== 支持的文件类型 ===")
	for _, fileTypeInfo := range supportedFileTypes {
		fmt.Printf("类型: %s | 扩展名: %v | 描述: %s\n",
			fileTypeInfo.FileType,
			fileTypeInfo.Extensions,
			fileTypeInfo.Description)
	}
}

// renameFileWithRetry 通过读写方式重命名文件，带重试机制
// 这种方式避免了 os.Rename 可能遇到的文件锁定问题
func renameFileWithRetry(oldPath, newPath string, maxRetries int) error {
	var lastErr error

	for attempt := 0; attempt < maxRetries; attempt++ {
		err := renameFileByReadWrite(oldPath, newPath)
		if err == nil {
			return nil
		}
		lastErr = err

		// 如果不是最后一次尝试，等待后重试
		if attempt < maxRetries-1 {
			time.Sleep(time.Millisecond * 100 * time.Duration(attempt+1))
		}
	}

	return lastErr
}

// renameFileByReadWrite 通过读取原文件、写入新文件的方式重命名
// 原文件异步删除，避免被占用问题
func renameFileByReadWrite(oldPath, newPath string) error {
	// 1. 打开原文件进行读取
	oldFile, err := os.Open(oldPath)
	if err != nil {
		return fmt.Errorf("打开原文件失败: %w", err)
	}
	defer oldFile.Close()

	// 2. 创建新文件
	newFile, err := os.Create(newPath)
	if err != nil {
		return fmt.Errorf("创建新文件失败: %w", err)
	}
	defer newFile.Close()

	// 3. 复制文件内容
	_, err = io.Copy(newFile, oldFile)
	if err != nil {
		os.Remove(newPath)
		return fmt.Errorf("复制文件内容失败: %w", err)
	}

	// 4. 确保新文件内容已完全写入磁盘
	err = newFile.Sync()
	if err != nil {
		os.Remove(newPath)
		return fmt.Errorf("文件同步失败: %w", err)
	}

	// 5. 关闭文件
	oldFile.Close()
	newFile.Close()

	// 6. 异步删除原文件（不阻塞主流程）
	go func() {
		time.Sleep(time.Second)
		for i := 0; i < 5; i++ {
			if os.Remove(oldPath) == nil {
				fmt.Printf("原文件已删除: %s\n", oldPath)
				return
			}
			time.Sleep(time.Second)
		}
	}()

	return nil
}
