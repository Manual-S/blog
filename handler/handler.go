package handler

import (
	"blog/global"
	"blog/server"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ApiArticleList(c *gin.Context) {
	// 调用server层的内容
	list, err := server.NewArticleServer(global.MysqlRW, global.RedisRW).
		ArticleList()
	if err != nil {
		replyErr(c, ErrInterval)
		return
	}
	replySucc(c, list)
	return
}

// ApiArticleScoreList 获取文章投票列表 从大到小排序
func ApiArticleScoreList(c *gin.Context) {
	start := c.Query("start")
	stop := c.Query("stop")
	startInt, err := strconv.ParseInt(start, 10, 64)
	stopInt, err := strconv.ParseInt(stop, 10, 64)

	server := server.NewArticleServer(global.MysqlRW, global.RedisRW)
	list, err := server.ArticleScoreList(startInt, stopInt)
	if err != nil {
		replyErr(c, ErrInterval)
		return
	}
	replySucc(c, list)
}

func ApiArticleVote(c *gin.Context) {
	// 向某篇文章投票
	titleID := c.PostForm("title_id")
	if titleID == "" {
		replyErr(c, ErrParams)
		return
	}
	server := server.NewArticleServer(global.MysqlRW, global.RedisRW)
	err := server.ArticleVote(titleID)
	if err != nil {
		replyErr(c, ErrInterval)
		return
	}
	replySucc(c, nil)
}

func ApiArticleCreate(c *gin.Context) {
	//title := c.PostForm("title")
	//summary := c.PostForm("summary")
	//content := c.PostForm("content")

}
