package resource

// MysqlSetting 声明数据库配置结构体
type MysqlSetting struct {
	UserName     string
	Password     string
	Host         string
	DBName       string
	MaxIdleConns int
	MaxOpenConns int
}

type RedisConfig struct {
	Host     string
	Password string
	DB       int
}

type ElasticSearchConfig struct {
	Host     string
	User     string
	Password string
}
