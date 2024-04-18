package database

import (
	"database/sql"
	"go-app/pkg/common"
	"go-app/pkg/config"
	"go-app/pkg/tools"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	common.DB, err = e.Open(sqlDB, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
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
