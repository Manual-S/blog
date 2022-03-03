package server

import (
	rClient "blog/data/redis"
	"blog/global"
	"fmt"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type ArticleServer interface {
	ArticleList() error
	ArticleVote(titleID string) error
	ArticleScoreList(start, stop int64) ([]string, error)
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
	key := "blog:article:votes"
	err := rClient.NewArticleRedis(global.RedisRW).
		AddVotes(key, titleID, 1)
	if err != nil {
		fmt.Println("AddVotes error", err)
		return err
	}
	return nil
}

func (a *articleServer) ArticleScoreList(start, stop int64) ([]string, error) {
	key := "blog:article:votes"
	vals, err := rClient.NewArticleRedis(global.RedisRW).TitleScoreList(key, start, stop)
	if err != nil {
		return nil, err
	}
	return vals, nil
}
