package redis

import (
	"context"
	"fmt"
	"go-app/pkg/common"
	"go-app/pkg/config"
	"go-app/pkg/tools"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var (
	Nil = redis.Nil
	ctx = context.Background()
)

// 初始化redis
func Setup(cfg *config.RedisConfig) (err error) {
	common.RDS = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:     cfg.Password, // no password set
		DB:           cfg.DB,       // use default DB
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})
	_, err = common.RDS.Ping(ctx).Result()
	if err != nil {
		zap.L().Error(tools.Red("init redis failied"), zap.Any("err", err))
		return
	}
	zap.L().Info(tools.Green("init redis success"))
	return nil
}
