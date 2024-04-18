package database

import (
	"go-app/pkg/tools"

	"go.uber.org/zap"
)

// Setup 配置数据库
func Setup(driver string) {
	dbType := driver
	switch dbType {
	case "mysql":
		var db = new(Mysql)
		db.Setup()
	case "sqlite3":
		var db = new(SqLite)
		db.Setup()
	default:
		zap.L().Error(tools.Red("database driver err or not exist!!"))
		return
	}
}
