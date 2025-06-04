package initialize

import (
	"wordCountServer/router"

	"github.com/gin-gonic/gin"
)

func RouterInit() *gin.Engine {
	r := gin.Default()
	allRouter := router.AllRouter 
	base := r.Group("/")
	{
		allRouter.ServerRouter.ApiRouterInit(base) 
	}
	return r
}