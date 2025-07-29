package initialize

import (
	"context"
	"fmt"

	"github.com/quangdat385/holiday-ticket/communications-service/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func InitRedis() {
	r := global.Config.Redis
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		Password: r.Password, // no password set
		DB:       r.Database, // use default DB
		PoolSize: 10,
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Redis initialization Error:", zap.Error(err))
	}
	fmt.Println("InitRedis is running")
	global.Rdb = rdb
}
