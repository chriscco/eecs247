package initialize

import (
	"fmt"
	"wordCountServer/global"

	"github.com/go-redis/redis"
)

func initRedis() *redis.Client {
	redis_config := global.Config.Redis 
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redis_config.Host, redis_config.Port),
		Password: redis_config.Password, // no password set
		DB:       redis_config.DataBase, // use default DB
	})
	ping := client.Ping() 
	err := ping.Err() 
	if err != nil {
		panic(err.Error()) 
	}
	return client 
}