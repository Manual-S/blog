package main

import (
	"blog/global"
	"blog/middle"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	file, err := os.Create("info.log")
	if err != nil {
		return
	}
	global.Logger = log.New(file, "", log.LstdFlags|log.Llongfile)
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
