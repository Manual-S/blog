package main

import (
	"blog/global"
	"blog/router"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
)

func init() {
	file, err := os.Create("info.log")
	if err != nil {
		return
	}
	global.Logger = log.New(file, "", log.LstdFlags|log.Llongfile)

	global.RedisRW = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0, // 默认使用0号库
	})
}

func main() {
	r := router.NewRouter()
	s := &http.Server{
		Addr:    "localhost:8080",
		Handler: r,
	}
	s.ListenAndServe()
}
