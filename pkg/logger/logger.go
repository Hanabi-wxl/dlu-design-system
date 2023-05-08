package logger

import (
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var Logger = &lumberjack.Logger{
	Filename:   "./log/log.txt",
	MaxSize:    10, // megabytes
	MaxBackups: 3,
	MaxAge:     28, //days
}

// Init logrus同时写文件和终端
func Init() {
	// 设置输出样式，自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006/01/02 15:04:05",
	})
	// 同时写文件和屏幕
	fileAndStdoutWriter := io.MultiWriter(os.Stdout, Logger)
	logrus.SetOutput(fileAndStdoutWriter)
	// 设置最低loglevel
	logrus.SetLevel(logrus.InfoLevel)
}
