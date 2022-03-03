package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Err struct {
	Code    int    `json:"code"`
	Info    string `json:"info"`
	Message string `json:"message"`
}

// 系统内部的错误应该打印到日志中 而不应该暴露给用户
// 统一错误定义
var (
	// ErrParams _
	ErrParams = Err{1001, "params wrong", "参数错误"}
	// ErrInterval _
	ErrInterval = Err{1002, "interval error", "服务器繁忙，请稍后再试"}
	// ErrPanic _
	ErrPanic = Err{1003, "interval error", "服务器繁忙，请稍后再试"}
)

func reply(c *gin.Context, code int, message string, data interface{}) {
	res := map[string]interface{}{
		"code":    code,
		"message": message,
		"data":    data,
	}
	c.JSON(http.StatusOK, res)
}

func replySucc(c *gin.Context, data interface{}) {
	reply(c, 0, "success", data)
}

func replyErr(c *gin.Context, err Err) {
	reply(c, err.Code, err.Message, nil)
}
