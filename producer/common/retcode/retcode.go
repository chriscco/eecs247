package retcode

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS = 1
	ERROR = 2
)

var errorMessage = map[int]string {
	SUCCESS: "ok",
	ERROR: "fatal error",
}

type ErrorCodeGet interface {
	GetCode() int
	GetMessage() string
}
type Error struct {
	code int
	msg  string
}

func (e *Error) GetCode() int {
	return e.code
}
func (e *Error) GetMessage() string {
	return e.msg
}

func NewError(code int, msg string) *Error {
	return &Error{
		code: code,
		msg:  msg,
	}
}

func Ok(c *gin.Context, data interface{}) {
	renderReply(c, data)
}
func Fatal(c *gin.Context, err error, msg string) {
	renderErrReply(c, ERROR, msg)
}

func renderReply(c *gin.Context, data interface{}) {
	render(c, SUCCESS, data, nil) 
}
func renderErrReply(c *gin.Context, code int, msg string) {
	render(c, code, nil, errors.New(msg))
}

func render(c *gin.Context, code int, data interface{}, err error) {
	var msg string 
	if err != nil {
		msg = err.Error() 
	} else if defaultMsg, ok := errorMessage[code]; ok {
		msg = defaultMsg 		
	}
	r := gin.H {
		"code": code,
		"msg": msg,
		"data": data,
	}
	c.Set("return_code", code)
	c.JSON(http.StatusOK, r)
}
