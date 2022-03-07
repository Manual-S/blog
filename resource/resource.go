package resource

import (
	"blog/global"
	"blog/internal"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Init 读取配置 初始化global中的变量
func Init(configPath string) error {
	var err error

	// 读取配置
	conClient, err := ReadConfig()
	if err != nil {
		return err
	}

	// 初始化log
	global.Logger, err = InitLog()
	if err != nil {
		return err
	}

	// 初始化redis
	redisConfig := RedisConfig{}
	conClient.UnmarshalKey("Redis", &redisConfig)
	global.RedisRW = InitRedis(redisConfig)

	// 初始化mysql
	mysqlSet := MysqlSetting{}
	conClient.UnmarshalKey("Mysql", &mysqlSet)
	global.MysqlRW, err = InitMysql(mysqlSet)
	if err != nil {
		return err
	}

	// 初始化全局id生成器
	global.SonyFlake, err = internal.InitID(0)
	if err != nil {
		return err
	}
	return nil
}

func InitLog() (*log.Logger, error) {
	file, err := os.Create("info.log")
	if err != nil {
		return nil, err
	}
	logClient := log.New(file, "", log.LstdFlags|log.Llongfile)

	return logClient, nil
}

func InitMysql(mysqlSet MysqlSetting) (*gorm.DB, error) {
	//dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlSet.UserName,
		mysqlSet.Password,
		mysqlSet.Host,
		mysqlSet.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

func InitRedis(c RedisConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     c.Host,
		Password: c.Password,
		DB:       c.DB, // 默认使用0号库
	})

	return client
}

// ReadConfig 从yaml中读取配置
func ReadConfig() (*viper.Viper, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	vp.AddConfigPath("./configs")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return vp, err
}
