package controller

import (
	"wordCountServer/api/service"

	"github.com/gin-gonic/gin"
)

type ServerController struct {
	serverImpl service.ServerImpl 
}

func NewServerContrller(service service.ServerImpl) *ServerController {
	return &ServerController{ 
		serverImpl: service,
	}
}

func (sc *ServerController) Query(c *gin.Context) {
	
}