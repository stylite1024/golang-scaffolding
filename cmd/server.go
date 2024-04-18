package cmd

import (
	"fmt"
	"go-app/pkg/config"
	"go-app/pkg/database"
	"go-app/pkg/logger"
	"go-app/pkg/tools"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	configYml string
	serverCmd = &cobra.Command{
		Use:          "server",
		Short:        "start server",
		Long:         `go-app server -c config/appication.yml`,
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func run() {
	r := gin.Default()

	// 使用zap记录gin日志
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 测试路由
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	// 启动服务器
	r.Run(config.Conf.ApplicationConfig.Host + ":" + config.Conf.ApplicationConfig.Port)
}

func setup() {
	// 0、检查配置文件
	fileExists(configYml)

	// 1、加载配置
	config.Setup(configYml)

	// 2、初始化日志
	logger.Setup(config.Conf.LogConfig, config.Conf.ApplicationConfig.Mode)

	// 3、初始化redis
	// redis.Setup(config.Conf.RedisConfig)
	zap.L().Error(tools.Red("aaa"))

	// 4、初始化mysql
	database.Setup(config.Conf.DatabaseConfig.Driver)
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config.yml", "Start server with provided configuration file")
	rootCmd.AddCommand(serverCmd)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		fmt.Println("config file is not exist!!")
		os.Exit(-1)
	}
	return !info.IsDir()
}
