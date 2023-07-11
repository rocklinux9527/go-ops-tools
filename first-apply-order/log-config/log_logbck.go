package log_config

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"time"
)

func LogBck() *lumberjack.Logger {
	currentTime := time.Now()
	formattedTime := currentTime.Format("20060102")
	logRotation := &lumberjack.Logger{
		Filename:   "./logs/" + formattedTime + "-app.log", // 日志文件名
		MaxSize:    10,                                     // 单个日志文件的最大大小（单位：MB）
		MaxBackups: 5,                                      // 最多保留的旧日志文件数
		LocalTime:  true,                                   // 使用本地时间作为日志的文件名和轮转
		Compress:   false,                                  // 是否压缩旧的日志文件
	}
	// fmt.Printf("logRotation type: %T\n", logRotation) // 打印日志对象的类型
	return logRotation
}
