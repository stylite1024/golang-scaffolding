package database

import (
	"database/sql"
	"go-app/pkg/common"
	"go-app/pkg/config"
	"go-app/pkg/tools"
	"io"
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// Mysql mysql配置结构体
type Mysql struct {
}

// Setup 配置步骤
func (e *Mysql) Setup() {
	common.Source = e.GetConnect()
	sqlDB, err := sql.Open("mysql", common.Source)
	if err != nil {
		zap.L().Error("sql.Open Err", zap.Any("err", err))
		return
	}

	if config.Conf.EnableSqlLog {
		common.DB, err = e.Open(sqlDB, &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			Logger: getGormLogger(),
		})
	} else {
		common.DB, err = e.Open(sqlDB, &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})
	}
	if config.Conf.Mode == "dev" {
		common.DB, err = e.Open(sqlDB, &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
				SlowThreshold:             200 * time.Millisecond, // 慢 SQL 阈值
				LogLevel:                  logger.Info,            // 日志级别
				IgnoreRecordNotFoundError: false,                  // 忽略ErrRecordNotFound（记录未找到）错误
				Colorful:                  true,                  // 禁用彩色打印
			}),
		})
	}
	if err != nil {
		zap.L().Error("gorm.Open Err", zap.Any("err", err))
		return
	} else {
		zap.L().Info(tools.Green(e.GetDriver() + " mysql success !"))
	}
	if common.DB.Error != nil {
		zap.L().Error("database Err", zap.Any("err", err))
	}
}

// Open 打开数据库连接
func (e *Mysql) Open(db *sql.DB, cfg *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(mysql.New(mysql.Config{Conn: db}), cfg)
}

// GetConnect 获取数据库连接
func (e *Mysql) GetConnect() string {
	return config.Conf.DatabaseConfig.Source
}

// GetDriver 获取连接
func (e *Mysql) GetDriver() string {
	return config.Conf.DatabaseConfig.Driver
}

func getGormLogger() logger.Interface {
	var logMode logger.LogLevel

	switch config.Conf.SqlLogLevel {
	case "silent":
		logMode = logger.Silent
	case "error":
		logMode = logger.Error
	case "warn":
		logMode = logger.Warn
	case "info":
		logMode = logger.Info
	default:
		logMode = logger.Info
	}
	return logger.New(getGormLogWriter(), logger.Config{
		SlowThreshold:             200 * time.Millisecond, // 慢 SQL 阈值
		LogLevel:                  logMode,                // 日志级别
		IgnoreRecordNotFoundError: false,                  // 忽略ErrRecordNotFound（记录未找到）错误
		Colorful:                  false,                  // 禁用彩色打印
	})
}

// 自定义 gorm Writer
func getGormLogWriter() logger.Writer {
	var writer io.Writer

	writer = os.Stdout

	// 是否启用日志文件
	if config.Conf.EnableSqlLog {
		// 自定义 Writer
		writer = &lumberjack.Logger{
			Filename:   config.Conf.SqlLogFilename,
			MaxSize:    config.Conf.MaxSize,
			MaxBackups: config.Conf.MaxBackups,
			MaxAge:     config.Conf.MaxAge,
			Compress:   config.Conf.Compress,
		}
	}
	return log.New(writer, "\r\n", log.LstdFlags)
}

// logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
// 	SlowThreshold:             200 * time.Millisecond, // 慢 SQL 阈值
// 	LogLevel:                  logger.Info,                // 日志级别
// 	IgnoreRecordNotFoundError: false,                  // 忽略ErrRecordNotFound（记录未找到）错误
// 	Colorful:                  false,                  // 禁用彩色打印
// })
