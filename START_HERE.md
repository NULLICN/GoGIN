# ⚡ 立即开始 - 30 秒快速指南

## 🎯 你需要知道的一切

### 最简单的使用方式
```go
import "GoGIN/itying/models"

// 检查文件类型（自动修正错误后缀名）
success, fileType, err := models.CheckFileTypeWithHeader("uploads/file.jpg", "")

// 验证特定类型
isJPEG, detected, err := models.CheckFileTypeWithHeader("uploads/photo.jpg", "jpeg")
```

## 🚀 3 行代码开始

```go
// 第 1 行: 导入
import "GoGIN/itying/models"

// 第 2 行: 调用
success, fileType, err := models.CheckFileTypeWithHeader("uploads/myfile.jpg", "")

// 第 3 行: 使用结果
fmt.Printf("文件类型: %s\n", fileType)
```

## 📌 关键参数

| 参数 | 值 | 说明 |
|------|-----|------|
| filePath | 文件路径 | "uploads/photo.jpg" |
| expectedType | 期望类型或空 | `""` = 只检查，`"jpeg"` = 验证 JPEG |

## ✅ 返回值

```go
(bool, string, error)
// true/false  - 符合期望的结果
// "jpeg"      - 检测到的文件类型
// nil/error   - 是否出错
```

## 💡 常用模式

### 模式 1: 只检查类型（推荐用于上传）
```go
_, fileType, err := models.CheckFileTypeWithHeader(path, "")
// 错误的后缀名会自动修正!
```

### 模式 2: 验证特定类型（推荐用于验证）
```go
ok, _, _ := models.CheckFileTypeWithHeader(path, "jpeg")
if ok { /* 是 JPEG */ }
```

### 模式 3: 严格的类型白名单
```go
for _, t := range []string{"jpeg", "png"} {
    if ok, _, _ := models.CheckFileTypeWithHeader(path, t); ok {
        // ✓ 允许
        break
    }
}
```

## 📋 支持的类型参数值

```
"jpeg", "png", "gif", "bmp", "pdf", "zip", "rar", "7z", "exe", "txt"
```

## ✨ 10 秒理解工作原理

```
文件头(Magic Bytes) → 识别真实类型 → 检查后缀名 → 自动修正 → 返回结果
```

## 🎯 实际项目例子

### Gin 框架中的文件上传
```go
func Upload(c *gin.Context) {
    file, _ := c.FormFile("file")
    path := "uploads/" + file.Filename
    c.SaveUploadedFile(file, path)
    
    // 检查并修正
    _, fileType, err := models.CheckFileTypeWithHeader(path, "")
    if err != nil {
        c.JSON(400, gin.H{"error": "检查失败"})
        return
    }
    
    c.JSON(200, gin.H{"type": fileType})
}
```

## 🔒 安全防护案例

```go
// 用户上传 virus.exe 但改名为 photo.jpg
ok, detected, _ := models.CheckFileTypeWithHeader(path, "jpeg")
if !ok && detected == "exe" {
    fmt.Println("🚨 拒绝: 检测到可执行文件!")
    os.Remove(path)
}
```

## 📚 需要更多帮助？

| 需求 | 文件 | ⏱️ |
|------|------|-----|
| 快速参考 | QUICK_REFERENCE.md | 2 分钟 |
| 代码示例 | code_examples.go | 5 分钟 |
| 详细文档 | FILE_TYPE_CHECK_README.md | 10 分钟 |

## ✅ 一句话总结

> 使用 `CheckFileTypeWithHeader()` 读取文件头识别真实类型，自动修正错误后缀名，防止恶意文件上传。

---

**就这么简单！** 🎉 

立即使用:
```go
models.CheckFileTypeWithHeader("uploads/file.jpg", "")
```

