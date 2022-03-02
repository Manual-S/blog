package main

import (
	"blog/middle"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	file, err := os.Create("test.log")
	if err != nil {
		return
	}
	global.logger = log.New(file, "", log.LstdFlags|log.Llongfile)
}

func main() {
	r := gin.Default()
	r.Use(middle.AccessMiddle)
	r.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	r.Run()

}
