package redis

import (
	"context"
	"github.com/Hanabi-wxl/dlu-design-system/config"
	"github.com/go-redis/redis/v8"
)

var RdJWT *redis.Client
var rdContext = context.Background()

func Init() {
	RdJWT = redis.NewClient(&redis.Options{
		Addr:     config.Conf.Redis.RdAddr,
		Password: config.Conf.Redis.RdPassWord,
		DB:       0,
	})
}
