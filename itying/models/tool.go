package models

import (
	"time"
)

// 时间戳转换成日期
func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

// 日期转换成时间戳 2020-05-02 15:04:05
func DateToUnix(str string) int64 {
	template := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(template, str, time.Local)
	if err != nil {
		return 0
	}
	return t.Unix()
}

// 获取时间戳
func GetUnix() int64 {
	// 补充：获取当前时间的时间戳（秒级）
	return time.Now().Unix()
}

// 可选：获取当前时间的毫秒级时间戳（扩展功能）
func GetUnixMilli() int64 {
	return time.Now().UnixMilli()
}

// 可选：获取当前时间的微秒级时间戳（扩展功能）
func GetUnixMicro() int64 {
	return time.Now().UnixMicro()
}
