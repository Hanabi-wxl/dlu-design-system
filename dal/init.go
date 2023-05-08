package dal

import (
	"fmt"
	"github.com/Hanabi-wxl/dlu-design-system/config"
	gen "github.com/Hanabi-wxl/dlu-design-system/dal/db"
	myLogger "github.com/Hanabi-wxl/dlu-design-system/pkg/logger"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io"
	"log"
	"os"
	"time"
)

type MyWriter struct {
	mlog *logrus.Logger
}

func (m *MyWriter) Printf(format string, v ...interface{}) {
	logstr := fmt.Sprintf(format, v...)
	//利用logrus记录日志
	m.mlog.Info(logstr)
}

func NewMyWriter() *MyWriter {
	logg := logrus.New()
	// 设置输出样式，自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	logg.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006/01/02 15:04:05",
	})
	// 同时写文件和屏幕
	fileAndStdoutWriter := io.MultiWriter(os.Stdout, myLogger.Logger)
	logg.SetOutput(fileAndStdoutWriter)
	// 设置最低loglevel
	logg.SetLevel(logrus.InfoLevel)
	return &MyWriter{mlog: logg}
}

func Init() {
	db, err := gorm.Open(mysql.Open(config.Conf.Mysql.DSN), &gorm.Config{
		Logger: logger.New(
			//设置Logger
			NewMyWriter(),
			logger.Config{
				//慢SQL阈值
				SlowThreshold: time.Millisecond,
				//设置日志级别，打印sql
				LogLevel: logger.Info,
			},
		),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 默认不加复数
		}})
	if err != nil {
		log.Fatal("数据库链接失败: ", err)
	}
	sqlDB, _ := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(20)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(200)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	gen.SetDefault(db)
}
