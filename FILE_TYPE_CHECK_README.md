# 文件类型检查功能说明

## 功能介绍

这是一个强大的文件类型检查和验证系统，通过**读取文件头（Magic Bytes）**而不仅仅是文件后缀名来判断文件的真实类型。该系统能够自动修正文件后缀名错误，提高上传文件的安全性。

## 核心函数

### CheckFileTypeWithHeader(filePath, expectedType)

**函数签名：**
```go
func CheckFileTypeWithHeader(filePath string, expectedType string) (bool, string, error)
```

**参数说明：**
- `filePath` (string): 文件的完整路径
- `expectedType` (string): 期望的文件类型
  - 如为空字符串 `""`: 只检查并修正后缀名，不验证特定类型
  - 如指定类型 (如 `"jpeg"`, `"png"`): 验证文件是否为该类型

**返回值：**
- `bool`: 验证结果
  - 当 `expectedType` 为空时: 如果检测成功则返回 `true`，失败返回 `false`
  - 当 `expectedType` 不为空时: 返回文件类型是否符合期望
- `string`: 检测到的实际文件类型
- `error`: 如有错误，返回错误信息

## 支持的文件类型

| 文件类型 | 扩展名 | 文件签名 (Hex) | 描述 |
|---------|--------|----------------|------|
| jpeg | .jpg, .jpeg | FF D8 FF | JPEG 图像 |
| png | .png | 89 50 4E 47 | PNG 图像 |
| gif | .gif | 47 49 46 38 | GIF 图像 |
| bmp | .bmp | 42 4D | BMP 图像 |
| pdf | .pdf | 25 50 44 46 | PDF 文档 |
| zip | .zip | 50 4B 03 04 | ZIP 压缩包 |
| rar | .rar | 52 61 72 21 | RAR 压缩包 |
| 7z | .7z | 37 7A BC AF 27 1C | 7Z 压缩包 |
| exe | .exe, .dll | 4D 5A | 可执行文件 |
| txt | .txt | (无固定签名) | 文本文件 |

## 使用示例

### 示例 1: 只检查和修正文件后缀名

```go
success, detectedType, err := models.CheckFileTypeWithHeader("uploads/test.jpg", "")
if err != nil {
    fmt.Println("错误:", err)
} else {
    fmt.Printf("文件类型: %s, 检查结果: %v\n", detectedType, success)
}
```

**说明：**
- 如果 `test.jpg` 实际是 PNG 文件，函数会自动将其改名为 `test.png`
- 返回检测到的真实类型

### 示例 2: 验证文件是否为特定类型

```go
success, detectedType, err := models.CheckFileTypeWithHeader("uploads/photo.jpg", "jpeg")
if err != nil {
    fmt.Println("错误:", err)
} else {
    if success {
        fmt.Println("✓ 文件确实是 JPEG 格式")
    } else {
        fmt.Printf("✗ 文件类型不匹配。期望: jpeg, 实际: %s\n", detectedType)
    }
}
```

**说明：**
- 验证文件是否为 JPEG 格式
- 如果文件后缀名不匹配，会自动修正

### 示例 3: 在文件上传时使用

```go
func (admin AdminController) AdminUploadFiles(c *gin.Context) {
    form, _ := c.MultipartForm()
    var filesName []string
    
    for key, headers := range form.File {
        fmt.Println(key, headers)
        for _, header := range headers {
            uploadPath := "uploads/" + header.Filename
            c.SaveUploadedFile(header, uploadPath)
            
            // 检查文件类型
            success, detectedType, err := models.CheckFileTypeWithHeader(uploadPath, "")
            if err != nil {
                fmt.Printf("文件类型检查失败: %v\n", err)
                continue
            }
            
            fmt.Printf("✓ 文件 %s 检测类型: %s\n", header.Filename, detectedType)
            filesName = append(filesName, header.Filename)
        }
    }
    admin.Success(c, filesName)
}
```

### 示例 4: 仅允许特定格式的上传

```go
func (admin AdminController) AdminUploadImages(c *gin.Context) {
    form, _ := c.MultipartForm()
    var validImages []string
    var rejectedFiles []string
    
    for key, headers := range form.File {
        for _, header := range headers {
            uploadPath := "uploads/" + header.Filename
            c.SaveUploadedFile(header, uploadPath)
            
            // 检查是否为允许的图片格式
            allowedTypes := []string{"jpeg", "png", "gif", "bmp"}
            isAllowed := false
            
            for _, allowedType := range allowedTypes {
                ok, _, _ := models.CheckFileTypeWithHeader(uploadPath, allowedType)
                if ok {
                    isAllowed = true
                    validImages = append(validImages, header.Filename)
                    break
                }
            }
            
            if !isAllowed {
                rejectedFiles = append(rejectedFiles, header.Filename)
                fmt.Printf("✗ 拒绝非图片格式文件: %s\n", header.Filename)
            }
        }
    }
    
    admin.Success(c, gin.H{
        "validImages": validImages,
        "rejected": rejectedFiles,
    })
}
```

## 工作原理

1. **读取文件头**
   - 打开文件并读取前 512 字节
   - 这包含了大多数文件类型的特征字节（Magic Bytes）

2. **识别文件类型**
   - 比对文件头与已知的文件类型签名
   - 识别出真实的文件类型

3. **验证后缀名**
   - 检查当前文件的扩展名
   - 与识别的文件类型比对

4. **自动修正**
   - 如果后缀名不匹配，使用 `os.Rename()` 修改文件名
   - 将错误的后缀名改为正确的后缀名

5. **类型验证**
   - 如果指定了期望的类型，进行验证
   - 返回是否符合期望

## 安全性优势

| 防护措施 | 说明 |
|---------|------|
| 文件头检查 | 不依赖容易伪造的文件后缀名 |
| Magic Bytes 识别 | 基于真实的文件签名判断 |
| 自动修正 | 纠正用户错误命名的文件 |
| 恶意文件防护 | 防止用户将病毒文件改成 `.jpg` 后缀上传 |
| 类型验证 | 可以严格限制允许的文件类型 |

## 注意事项

1. **文件权限**
   - 函数需要读取文件的权限
   - 在修改文件名时需要写入权限

2. **性能考虑**
   - 函数会进行文件 I/O 操作
   - 在处理大量文件时可能需要考虑性能优化

3. **支持的类型扩展**
   - 可以在 `supportedFileTypes` 中添加新的文件类型
   - 添加对应的 Magic Bytes 特征

4. **文本文件检测**
   - 对于纯文本文件，由于没有固定的 Magic Bytes
   - 函数会检查是否包含过多的非文本字符来判断

## 函数列表

- `CheckFileTypeWithHeader(filePath, expectedType)` - 主函数，检查文件类型
- `detectFileType(headerBytes)` - 根据文件头检测类型
- `bytesMatch(data, pattern)` - 比较字节切片
- `isPureText(headerBytes)` - 判断是否为文本文件
- `getCorrectExtension(fileType)` - 获取正确的扩展名
- `GetSupportedFileTypes()` - 获取所有支持的文件类型
- `PrintSupportedFileTypes()` - 打印支持的文件类型列表
- `DemonstrateUsage()` - 演示功能使用方式

## 集成指南

1. **在模型层使用**
   ```go
   // 在 models/fileTypeCheck.go 中已实现
   ```

2. **在控制器中调用**
   ```go
   // 在 controller/admin/adminController.go 中已集成示例
   ```

3. **自定义扩展**
   ```go
   // 修改 supportedFileTypes 变量添加新的文件类型支持
   ```

这个功能提供了一个健壮的文件验证系统，确保上传的文件真正符合其声称的类型。

