# 快速参考 - 文件类型检查函数

## 快速开始

### 基本用法
```go
// 检查文件类型并自动修正后缀名
success, fileType, err := models.CheckFileTypeWithHeader("uploads/myfile.jpg", "")

// 检查文件是否为特定类型
isJPEG, detectedType, err := models.CheckFileTypeWithHeader("uploads/photo.jpg", "jpeg")
```

## 常见场景

### 场景 1️⃣: 安全的文件上传处理
```go
uploadPath := "uploads/" + header.Filename
c.SaveUploadedFile(header, uploadPath)

// 验证并修正
_, detectedType, err := models.CheckFileTypeWithHeader(uploadPath, "")
if err == nil {
    fmt.Printf("上传成功，检测类型: %s\n", detectedType)
}
```

### 场景 2️⃣: 仅允许图片上传
```go
for _, allowedType := range []string{"jpeg", "png", "gif"} {
    ok, _, _ := models.CheckFileTypeWithHeader(uploadPath, allowedType)
    if ok {
        fmt.Println("✓ 允许上传")
        break
    }
}
```

### 场景 3️⃣: 验证特定文件类型
```go
ok, detectedType, err := models.CheckFileTypeWithHeader("file.pdf", "pdf")
if !ok {
    fmt.Printf("警告: 期望 PDF，实际是 %s\n", detectedType)
}
```

## 支持的类型速查

| 类型 | 期望值 | 扩展名 |
|------|--------|--------|
| 图片-JPEG | `"jpeg"` | .jpg, .jpeg |
| 图片-PNG | `"png"` | .png |
| 图片-GIF | `"gif"` | .gif |
| 图片-BMP | `"bmp"` | .bmp |
| 文档-PDF | `"pdf"` | .pdf |
| 压缩包 | `"zip"` | .zip |
| 文本 | `"txt"` | .txt |

## 函数签名
```go
func CheckFileTypeWithHeader(filePath, expectedType string) (bool, string, error)
```

| 参数 | 返回 |
|------|------|
| `filePath`: 文件路径 | `bool`: 是否符合期望 |
| `expectedType`: 期望类型或 `""` | `string`: 检测到的类型 |
|  | `error`: 错误信息 |

## 返回值说明

| expectedType | 返回 true | 返回 false |
|--------------|-----------|-----------|
| `""` (空) | 成功检测到类型 | 无法识别类型 |
| `"jpeg"` | 文件确实是 JPEG | 文件不是 JPEG |

## 工作流程

```
┌─────────────────┐
│ 读取文件头      │
└────────┬────────┘
         │
┌────────▼────────┐
│ 识别文件类型    │
└────────┬────────┘
         │
┌────────▼────────┐
│ 检查后缀名      │
└────────┬────────┘
         │
    ┌────┴─────┐
    │ 匹配？    │
    └┬─────────┘
  不匹配
     │
┌────▼─────────────┐
│ 自动修改文件名   │
└────┬─────────────┘
     │
┌────▼──────────────────┐
│ expectedType 为空?     │
└┬─────────────────────┘
 │
 ├─ 是 → 返回 true (已检测并修正)
 │
 └─ 否 → 检查是否符合期望类型
       → 符合返回 true，不符合返回 false
```

## 错误处理示例

```go
success, detectedType, err := models.CheckFileTypeWithHeader(uploadPath, "")

if err != nil {
    switch {
    case errors.Is(err, os.ErrNotExist):
        fmt.Println("文件不存在")
    default:
        fmt.Printf("错误: %v\n", err)
    }
    return
}

fmt.Printf("✓ 文件类型: %s\n", detectedType)
```

## 性能提示

✅ 做法：
- 为批量文件上传创建类型白名单
- 在过滤后才调用检查函数

❌ 不要：
- 对每个文件都检查所有支持的类型
- 在没有错误处理的情况下使用

## 输出示例

```
✓ 文件 photo.jpg 检测类型: jpeg, 验证结果: true
✓ 文件 image.txt 检测类型: png, 验证结果: false (期望 png)
文件后缀名已从 .txt 修改为 .png
```

## 安全性要点

🔒 防护特性：
- ✅ 读取真实的文件签名（Magic Bytes）
- ✅ 防止伪造的文件后缀名
- ✅ 自动修正错误命名
- ✅ 支持类型白名单验证

---

📚 详见: `FILE_TYPE_CHECK_README.md`

