// 中间件
package middle

import (
	"blog/global"

	"github.com/gin-gonic/gin"
)

func AccessMiddle(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()

	defer func() {
		global.Logger.Printf("[%s] [%s]", c.Request.RemoteAddr, c.Request.URL)
	}()
}
