package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	model "wordCountServer/api/model"
	"wordCountServer/api/service"

	"github.com/gin-gonic/gin"
)

// ServerController
type ServerController struct {
	serverImpl service.ServerImpl 
}

// NewServerContrller 
//	@param service 
//	@return *ServerController 
func NewServerContrller(service service.ServerImpl) *ServerController {
	return &ServerController{ 
		serverImpl: service,
	}
}

// Query 
//	@param c 
func (sc *ServerController) Query(c *gin.Context) {
	start := time.Now() 
	var request model.Request 
	err := c.ShouldBindBodyWithJSON(&request) 
	if err != nil {
		log.Fatal("unable to bind request: ", err)
		c.String(http.StatusBadRequest, "unable to bind request")  
	}
	result := sc.serverImpl.Count(c, request.Text) 

	timeCost := time.Since(start)

	c.String(http.StatusOK, fmt.Sprintf("WordCount: %s\nTime Cost: %s\n", 
		strconv.Itoa(result), timeCost.String()))
}