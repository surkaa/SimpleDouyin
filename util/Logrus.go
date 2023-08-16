package util

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

// logger 创建全局的logrus实例
var logger *logrus.Logger

func InitLogger(logFileName string) {
	logger = logrus.New()

	// 设置日志级别
	logger.SetLevel(logrus.InfoLevel)

	// 可选：将日志输出到文件
	file, err := os.Create(logFileName)
	if err != nil {
		log.Fatal("Failed to create log file:", err)
	}
	defer file.Close()
	logger.SetOutput(file)

	// 可选：设置日志格式
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:  "06-01-02 15:04:05.000",
		DisableTimestamp: true,
	})
}

func Log() *logrus.Logger {
	return logger
}
