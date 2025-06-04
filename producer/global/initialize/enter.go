package initialize

import (
	"wordCountServer/config"
	"wordCountServer/global"

	"github.com/gin-gonic/gin"
)

func GlobalInit() *gin.Engine {
	global.Config = config.ConfigInit() 
	global.Redis = initRedis() 
	router := RouterInit() 
	return router 
}