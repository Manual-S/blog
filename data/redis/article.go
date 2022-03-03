package redis

import "github.com/go-redis/redis"

type ArticleRedis interface {
	AddVotes(key string, titleID string, score int) error
	TitleScoreList(key string, start int64, stop int64) ([]string, error)
}

type articleRedis struct {
	redisClient *redis.Client
}

func NewArticleRedis(client *redis.Client) ArticleRedis {
	return &articleRedis{
		redisClient: client,
	}
}

func (a *articleRedis) AddVotes(key string, titleID string, score int) error {
	err := a.redisClient.ZIncrBy(key, float64(score), titleID).Err()
	return err
}

func (a *articleRedis) TitleScoreList(key string, start int64, stop int64) ([]string, error) {
	vals, err := a.redisClient.ZRevRange(key, start, stop).Result()
	if err != nil {
		return nil, err
	}
	return vals, nil
}
