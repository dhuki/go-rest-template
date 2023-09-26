package cache

import (
	"github.com/dhuki/go-rest-template/internal/infra/configloader"
	"github.com/go-redis/redis"
)

func InitRedis(conf *configloader.RedisConfig) (err error) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     conf.Host,
		Password: conf.Password,
		DB:       conf.DB,
	})

	if _, err = RedisClient.Ping().Result(); err != nil {
		return
	}
	return
}
