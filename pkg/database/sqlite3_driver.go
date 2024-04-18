package database

import (
	"go-app/pkg/common"
	"go-app/pkg/config"
	"go-app/pkg/tools"
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type SqLite struct {
}

var err error

func (e *SqLite) Setup() {
	if config.Conf.EnableSqlLog {
		common.DB, err = e.Open(e.GetConnect(), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			Logger: getGormLogger(),
		})
	} else {
		common.DB, err = e.Open(e.GetConnect(), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})
	}
	if config.Conf.Mode == "dev" {
		common.DB, err = e.Open(e.GetConnect(), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
				SlowThreshold:             200 * time.Millisecond, // 慢 SQL 阈值
				LogLevel:                  logger.Info,            // 日志级别
				IgnoreRecordNotFoundError: false,                  // 忽略ErrRecordNotFound（记录未找到）错误
				Colorful:                  true,                   // 禁用彩色打印
			}),
		})
	}
	if err != nil {
		zap.L().Error("gorm.Open Err", zap.Any("err", err))
		return
	} else {
		zap.L().Info(tools.Green(e.GetDriver() + " init success !"))
	}
	if common.DB.Error != nil {
		zap.L().Error("database Err", zap.Any("err", err))
	}
}

// 打开数据库连接
func (*SqLite) Open(conn string, cfg *gorm.Config) (db *gorm.DB, err error) {
	return gorm.Open(sqlite.Open(conn), cfg)
}

// GetConnect 获取数据库连接
func (e *SqLite) GetConnect() string {
	return config.Conf.DatabaseConfig.Source
}

// GetDriver 获取连接
func (e *SqLite) GetDriver() string {
	return config.Conf.DatabaseConfig.Driver
}
