package conn

import (
	"aschool/settings"
	"github.com/go-redis/redis"
)

// 保存redis连接
var (
	RedisDb *redis.Client
)

// 创建到redis的连接
func InitRedis(cfg *settings.RedisConfig) (err error) {

	RedisDb = redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password, // no password set
		DB:       cfg.DB,       // use default DB
	})

	_, err = RedisDb.Ping().Result()
	if err != nil {
		return err
	}
	return nil

}

func CloseRedis() {
	RedisDb.Close()
}
