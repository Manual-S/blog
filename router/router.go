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
		apiv1.GET("articles/list", handler.ApiArticleList)           // 获取文章列表
		apiv1.GET("articles/socrelist", handler.ApiArticleScoreList) // 获取文章的排序列表
		apiv1.POST("articles/votes", handler.ApiArticleVote)         // 给文章投票
		apiv1.POST("article/crate", handler.ApiArticleCreate)        // 创建文章
		apiv1.GET("article/info", handler.ApiArticleInfo)            // 获取文章详情
		apiv1.GET("article/search", handler.ApiArticleSearch)
		//apiv1.POST("user/create", handler.)

	}

	return r
}
