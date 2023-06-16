package model

// Config 配置。
type Config struct {
	Redis Redis
	Mysql Mysql
}

// Redis 缓存。
type Redis struct {
	Address string
}

// Mysql 数据库。
type Mysql struct {
	User    string `yaml:"user"`
	Pwd     string `yaml:"pwd"`
	Address string `yaml:"address"`
	DbBase  string `yaml:"dbBase"`
}
