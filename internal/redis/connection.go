package redis

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var (
	redisConnection *redis.Client
)

func GetConnection() (*redis.Client, error) {
	if redisConnection == nil {
		rdb := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", viper.GetString("redis.host"), viper.GetString("redis.port")),
			Password: viper.GetString("redis.password"),
			DB:       viper.GetInt("redis.database"),
		})

		redisConnection = rdb
	}

	return redisConnection, nil
}
