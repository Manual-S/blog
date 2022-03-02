package global

import (
	"log"

	"github.com/go-redis/redis"

	"github.com/olivere/elastic"
	"gorm.io/gorm"
)

var (
	Logger    *log.Logger     // 日志组件
	MysqlRW   *gorm.DB        // mysql
	ElasticRW *elastic.Client // elastic
	RedisRW   *redis.Client   // redis
)
