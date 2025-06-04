package router

import (
	"wordCountServer/api/controller"
	"wordCountServer/api/service"

	"github.com/gin-gonic/gin"
)

type ServerRouter struct {
	service service.ServerImpl 
}
func (sr *ServerRouter) ApiRouterInit(router *gin.RouterGroup) {
	r := router.Group("/") 
	sr.service = service.NewServerImpl() 
	serverController := controller.NewServerContrller(sr.service) 
	{
		r.GET("/query", serverController.Query) 
	}
}
