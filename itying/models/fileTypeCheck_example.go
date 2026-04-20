package models

/*
使用示例：

// 示例 1: 检查文件类型并自动修正后缀名
success, detectedType, err := CheckFileTypeWithHeader("uploads/test.jpg", "")
if err != nil {
	fmt.Println("错误:", err)
} else {
	fmt.Printf("文件类型: %s, 检查结果: %v\n", detectedType, success)
}

// 示例 2: 检查文件是否为期望的类型 (JPEG)
success, detectedType, err := CheckFileTypeWithHeader("uploads/photo.jpg", "jpeg")
if err != nil {
	fmt.Println("错误:", err)
} else {
	if success {
		fmt.Println("✓ 文件确实是 JPEG 格式")
	} else {
		fmt.Printf("✗ 文件类型不匹配。期望: jpeg, 实际: %s\n", detectedType)
	}
}

// 示例 3: 检查文件是否为期望的类型 (PNG)
success, detectedType, err := CheckFileTypeWithHeader("uploads/image.png", "png")
if success {
	fmt.Println("✓ 文件是 PNG 格式")
}

// 示例 4: 检查文件是否为期望的类型 (PDF)
success, detectedType, err := CheckFileTypeWithHeader("uploads/document.pdf", "pdf")
if !success {
	fmt.Printf("✗ 文件类型不是 PDF，实际类型: %s\n", detectedType)
}

// 示例 5: 打印所有支持的文件类型
PrintSupportedFileTypes()

// 支持的文件类型列表：
// - jpeg (.jpg, .jpeg)
// - png (.png)
// - gif (.gif)
// - bmp (.bmp)
// - pdf (.pdf)
// - zip (.zip)
// - rar (.rar)
// - 7z (.7z)
// - exe (.exe, .dll)
// - txt (.txt)

功能说明：
1. CheckFileTypeWithHeader(filePath, expectedType) 函数会：
   - 读取文件的前512字节获取文件头（magic bytes）
   - 识别实际的文件类型
   - 检查文件后缀名是否匹配
   - 如果后缀名不匹配，自动修改为正确的后缀名
   - 如果指定了 expectedType，验证文件类型是否符合期望

2. 返回值：
   - bool: 当 expectedType 为空时总是返回 true（如果检测成功），
           当 expectedType 不为空时，返回文件类型是否符合期望
   - string: 检测到的实际文件类型
   - error: 如果发生错误会返回错误信息

3. 修改后缀名说明：
   - 函数会自动将错误的后缀名修改为正确的后缀名
   - 例如：test.txt 实际是 PNG 文件，会被改为 test.png
   - 例如：photo.jpg 实际是 GIF 文件，会被改为 photo.gif
*/
