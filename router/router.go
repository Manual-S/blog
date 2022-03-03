// Package router 处理路由请求
package router

import (
	"blog/handler"
	"blog/middle"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(middle.AccessMiddle)
	apiv1 := r.Group("/api/")
	{
		apiv1.GET("articles/list", handler.ApiArticleList) // 获取文章列表
		apiv1.GET("articles/socrelist", handler.ApiArticleScoreList)
		apiv1.POST("articles/votes", handler.ApiArticleVote) // 给文章投票
	}

	return r
}
