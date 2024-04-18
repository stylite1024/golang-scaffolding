package database

import (
	"go-app/pkg/common"
	"go-app/pkg/config"
	"go-app/pkg/tools"

	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type SqLite struct {
}

func (e *SqLite) Setup() {
	var err error

	common.DB, err = e.Open(e.GetConnect(), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
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
