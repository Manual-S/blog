package server

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type ArticleServer interface {
	ArticleList() error
	ArticleVote(titleID string) error
}

type articleServer struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewArticleServer(db *gorm.DB, redis *redis.Client) ArticleServer {
	return &articleServer{
		db:          db,
		redisClient: redis,
	}
}

// ArticleList 获取文章列表
func (a *articleServer) ArticleList() error {
	return nil
}

// ArticleVote 向指定文章投票
func (a *articleServer) ArticleVote(titleID string) error {
	// 查看文章存不存在
	// 不存在返回
	// 存在先写入mysql 再写入redis

	return nil
}
