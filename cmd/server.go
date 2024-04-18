package cmd

import (
	"context"
	"fmt"
	"go-app/pkg/common"
	"go-app/pkg/config"
	"go-app/pkg/database"
	"go-app/pkg/logger"
	"go-app/pkg/redis"
	"os"

	"github.com/spf13/cobra"
)

var (
	configYml string
	ctx       = context.Background()
	serverCmd = &cobra.Command{
		Use:          "server",
		Short:        "start server",
		Long:         `go-app server -c config/appication.yml`,
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		Run: func(cmd *cobra.Command, args []string) {
			db := common.DB
			// 自动迁移
			db.AutoMigrate(&UserInfo{})
			u1 := UserInfo{1, "七米", "男", "篮球"}
			u2 := UserInfo{2, "沙河娜扎", "女", "足球"}
			// 创建记录
			db.Create(&u1)
			db.Create(&u2)
			// 查询
			var u = new(UserInfo)
			db.First(u)
			fmt.Printf("%#v\n", u)

			// 测试redis
			// 设置一个key，过期时间为0，意思就是永远不过期
			err := common.RDS.Set(ctx, "user", "1", 0).Err()
			if err != nil {
				panic(err)
			}
			// 根据key查询缓存，通过Result函数返回两个值
			val, err := common.RDS.Get(ctx, "user").Result()
			// 检测，查询是否出错
			if err != nil {
				panic(err)
			}
			fmt.Println("user", val)
		},
	}
)

// UserInfo 用户信息
type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func setup() {
	// 0、检查配置文件
	fileExists(configYml)

	// 1、加载配置
	config.Setup(configYml)

	// 2、初始化日志
	logger.Setup(config.Conf.LogConfig, config.Conf.ApplicationConfig.Mode)

	// 3、初始化redis
	redis.Setup(config.Conf.RedisConfig)

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
