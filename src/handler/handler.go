package handler

import (
	"net/http"
	"api-server-demo/src/errno"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	// write response data to log
	log.Infof("code:%s message:%s data:%s", code, message, data)

	// always return http.StatusOK
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}