package global

import (
	"wordCountServer/config"
	"github.com/go-redis/redis"
)

var (
	Config *config.AllConfig 
	Redis *redis.Client
)
