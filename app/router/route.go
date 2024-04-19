package router

import (
	"context"

	"go-app/app/apis"
	"go-app/app/vali"
	"go-app/app/middleware"
	"go-app/pkg/config"
	"go-app/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// SetupRouter 路由
func SetupRouter() *gin.Engine {
	if err := vali.InitTrans("zh"); err != nil {
		zap.L().Error("init trans failed, err:", zap.Error(err))
		return nil
	}

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
	r.Use(middleware.UnifiedResponseMiddleware())

	// 前端项目静态资源
	r.StaticFile("/", "./static/dist/index.html")
	r.Static("/assets", "./static/dist/assets")
	r.StaticFile("/favicon.ico", "./static/dist/favicon.ico")
	// 其他静态资源
	r.Static("/public", "./static")
	r.Static("/storage", "./storage/app/public")

	// 注册 api 分组路由
	apiGroup := r.Group("/api/v1")
	apis.SetApiGroupRoutes(apiGroup)

	// 注册 register 分组路由
	registerGroup := r.Group("/api/v1")
	apis.SetRegisterGroupROutes(registerGroup)

	return r
}

// 启动http服务器
func RunServer() {
	r := SetupRouter()
	// r.Run(config.Conf.ApplicationConfig.Host + ":" + config.Conf.ApplicationConfig.Port)

	srv := &http.Server{
		Addr:    config.Conf.ApplicationConfig.Host + ":" + config.Conf.ApplicationConfig.Port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Error("listen:", zap.Error(err))
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Error("Server Shutdown:", zap.Error(err))
	}
	zap.L().Info("Server exiting")
}
