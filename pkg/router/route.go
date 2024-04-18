package router

import (
	"go-app/app/admin/apis"
	"go-app/pkg/config"
	"go-app/pkg/logger"

	"github.com/gin-gonic/gin"
)

// SetupRouter 路由
func SetupRouter() *gin.Engine {
	mode := config.Conf.Mode
	if mode == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else if mode == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	r := gin.New()

	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册 api 分组路由
	apiGroup := r.Group("/api/v1")
	apis.SetApiGroupRoutes(apiGroup)

	return r
}

// 启动http服务器
func RunServer() {
	r := SetupRouter()
	r.Run(config.Conf.ApplicationConfig.Host + ":" + config.Conf.ApplicationConfig.Port)
}
