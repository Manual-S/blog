package server

import (
	"blog/data/elasticsearch"
	"blog/data/mysql"
	rClient "blog/data/redis"
	"blog/global"
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type ArticleServer interface {
	ArticleList() ([]mysql.Article, error)
	ArticleVote(titleID string) error
	ArticleScoreList(start, stop int64) ([]string, error)
	ArticleCreate(title, summary, content string) error
	ArticleInfo(titleId string) (mysql.Article, error)
	ArticleSearch(keyWord string) ([]elasticsearch.Article, error)
}

type articleServer struct {
	ctx         context.Context
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
func (a *articleServer) ArticleList() ([]mysql.Article, error) {
	mysqlClient := mysql.NewArticleClient(global.MysqlRW)
	list, err := mysqlClient.QueryArticleList()
	if err != nil {
		return nil, err
	}
	return list, nil
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

// ArticleCreate 创建文章
func (a *articleServer) ArticleCreate(title, summary, content string) error {
	mysqlClient := mysql.NewArticleClient(global.MysqlRW)
	articleId, err := global.SonyFlake.NextID()
	if err != nil {
		return err
	}
	strconv.FormatUint(articleId, 10)
	err = mysqlClient.CreateArticle(mysql.Article{
		ArticleId: strconv.FormatUint(articleId, 10),
		Title:     title,
		Summary:   summary,
		Content:   content,
	})
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (a *articleServer) ArticleInfo(titleId string) (mysql.Article, error) {
	mysqlClient := mysql.NewArticleClient(global.MysqlRW)
	articleInfo, err := mysqlClient.QueryArticleInfo(titleId)
	if err != nil {
		return mysql.Article{}, err
	}

	return articleInfo, nil
}

func (a *articleServer) ArticleSearch(keyWord string) ([]elasticsearch.Article, error) {
	esClient := elasticsearch.NewTitleClient(global.ElasticRW)
	articleList, err := esClient.QueryTitle(a.ctx, keyWord)
	if err != nil {
		return nil, err
	}
	return articleList, nil
}
