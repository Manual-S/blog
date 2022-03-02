package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func replyErr(c *gin.Context, errCode int, errMessage string) {
	reply(c, errCode, errMessage, nil)
}
