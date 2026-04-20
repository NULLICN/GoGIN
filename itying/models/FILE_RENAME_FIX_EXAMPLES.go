package models

// ============ 文件重命名修复示例 ============
// 这个文件演示了修复后的文件重命名逻辑如何工作

/*

示例 1: 基本用法 - 检查并修正文件后缀名
=========================================

上传文件: photo.jpg （但实际是 GIF 格式）

// 在 adminController.go 中：
uploadPath := "uploads/photo.jpg"
success, detectedType, err := models.CheckFileTypeWithHeader(uploadPath, "")

结果:
- success: true
- detectedType: "gif"
- err: nil
- 文件已自动从 photo.jpg 重命名为 photo.gif

原理:
1. 检测文件的 Magic Bytes（文件头）
2. 发现实际是 GIF 格式
3. 调用 renameFileWithRetry() 进行重命名
4. renameFileWithRetry 自动重试最多 3 次
5. 通过 renameFileByReadWrite() 完成重命名


示例 2: 类型验证 - 检查文件是否符合期望类型
==========================================

验证上传的 PDF 文件:

uploadPath := "uploads/document.pdf"
isCorrect, detectedType, err := models.CheckFileTypeWithHeader(uploadPath, "pdf")

结果:
- isCorrect: true
- detectedType: "pdf"
- err: nil
- 文件符合期望类型


示例 3: 错误处理 - 处理文件被占用的情况
========================================

旧代码（容易失败）:
err := os.Rename(filePath, newFilePath)
if err != nil {
    // Windows 上经常会失败: "文件被另一个进程使用"
    fmt.Printf("重命名失败: %v\n", err)
}

新代码（稳健处理）:
err := renameFileWithRetry(filePath, newFilePath, 3)
if err != nil {
    // 已经重试 3 次，才返回错误
    fmt.Printf("重命名失败（3次重试后）: %v\n", err)
}

重试流程:
- 第1次: 立即尝试
  ↓ 失败
- 等待 100ms
- 第2次: 重试
  ↓ 失败
- 等待 200ms
- 第3次: 重试
  ↓ 成功或最终失败


示例 4: 检查支持的文件类型
==========================

// 查看所有支持的文件类型
models.PrintSupportedFileTypes()

输出:
=== 支持的文件类型 ===
类型: jpeg | 扩展名: [.jpg .jpeg] | 描述: JPEG Image
类型: png | 扩展名: [.png] | 描述: PNG Image
类型: gif | 扩展名: [.gif] | 描述: GIF Image
类型: bmp | 扩展名: [.bmp] | 描述: BMP Image
类型: pdf | 扩展名: [.pdf] | 描述: PDF Document
类型: zip | 扩展名: [.zip] | 描述: ZIP Archive
类型: rar | 扩展名: [.rar] | 描述: RAR Archive
类型: 7z | 扩展名: [.7z] | 描述: 7Z Archive
类型: exe | 扩展名: [.exe .dll] | 描述: Executable File
类型: txt | 扩展名: [.txt] | 描述: Text File


示例 5: 完整的上传流程
======================

// 在 adminController.go 中的 AdminUploadFiles 方法：

form, _ := c.MultipartForm()
var filesName []string

for key, headers := range form.File {
    for _, header := range headers {
        uploadPath := "uploads/" + header.Filename

        // 1. 保存文件
        if err := c.SaveUploadedFile(header, uploadPath); err != nil {
            fmt.Printf("保存文件失败: %v\n", err)
            continue
        }

        // 2. 检查并修正文件类型
        success, detectedType, err := models.CheckFileTypeWithHeader(uploadPath, "")
        if err != nil {
            fmt.Printf("文件类型检查失败: %v\n", err)
            // 即使检查失败，文件仍然保存了
        } else {
            fmt.Printf("文件 %s 检测类型: %s, 验证结果: %v\n",
                header.Filename, detectedType, success)
        }

        filesName = append(filesName, header.Filename)
    }
}

admin.Success(c, filesName)

流程说明:
- SaveUploadedFile: 将上传的文件保存到磁盘
- CheckFileTypeWithHeader: 检查文件类型并自动修正后缀名
- 即使文件被占用导致修正失败，文件仍然被保存
- 返回文件名列表给客户端


示例 6: 文件被占用场景的重试演示
=================================

场景: 上传大文件，系统还在进行杀毒扫描，导致文件被占用

流程:
1. 保存文件: "uploads/bigfile.jpg"
2. 读文件头并检测到实际类型是 GIF
3. 尝试重命名为 "uploads/bigfile.gif"
4. 调用 os.Rename() -> 失败: "文件被另一个进程使用"
5. 等待 100ms
6. 重试: 调用 io.Copy 方式重命名 -> 成功！
7. 删除原文件 "uploads/bigfile.jpg"
8. 返回成功

结果: 文件已正确重命名为 "uploads/bigfile.gif"


示例 7: 错误恢复演示
====================

场景: 复制文件时磁盘空间不足

流程:
1. 打开原文件: "uploads/file.jpg"
2. 创建新文件: "uploads/file.gif"
3. 调用 io.Copy() -> 失败: 磁盘空间不足
4. 自动删除已创建的新文件 "uploads/file.gif"
5. 返回错误信息

结果: 原文件 "uploads/file.jpg" 保持完整，不会有孤立的新文件


关键改进点总结
=============

✅ 避免文件锁定
   - 不再使用容易失败的 os.Rename()
   - 改用 io.Copy() 方式处理

✅ 自动重试
   - 最多重试 3 次
   - 指数退避延迟（100ms, 200ms）

✅ 完整错误恢复
   - 复制失败自动清理
   - 同步失败自动回滚
   - 删除失败清理新文件

✅ 跨平台兼容
   - Windows 友好
   - Linux/macOS 也支持

✅ 显式资源管理
   - 显式关闭文件句柄
   - 调用 Sync() 确保数据持久化

*/
