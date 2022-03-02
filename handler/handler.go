package handler

import (
	"blog/global"
	"blog/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiArticleList(c *gin.Context) {
	// 调用server层的内容
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
	})
}

func ApiArticleVote(c *gin.Context) {
	// 向某篇文章投票
	titleID := c.PostForm("title_id")
	if titleID == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    1001,
			"message": "param error",
		})
		return
	}
	server := server.NewArticleServer(global.MysqlRW, global.RedisRW)
	err := server.ArticleVote(titleID)
	if err != nil {
		replyErr(c, 1001, "internal error")
		return
	}
	replySucc(c, nil)
}
