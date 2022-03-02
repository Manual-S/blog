// 中间件
package middle

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AccessMiddle(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()

	defer func() {
		fmt.Println(c.Request.RemoteAddr)
	}()
}
