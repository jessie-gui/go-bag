package core

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// NewLogger 新建日志对象
func NewLogger() *zap.Logger {
	// 日志文件路径配置
	writer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./logs/app.log", // 日志文件路径
		MaxSize:    128,              // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 30,               // 日志文件最多保存多少个备份
		MaxAge:     7,                // 文件最多保存多少天
		Compress:   true,             // 是否压缩
	})

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)

	// 配置文件编码方式
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 时间格式

	// 设置日志输出格式
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig), // 编码器配置
		writer,                                   // 打印到控制台和文件
		atomicLevel,                              // 日志级别
	)

	logger := zap.New(core)

	log.Println("logger 初始化完成！")

	return logger
}
