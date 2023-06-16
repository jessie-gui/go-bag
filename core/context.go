package core

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ContextOption func(c *Context)

// Context 应用上下文对象。
type Context struct {
	cache  *redis.Client
	logger *zap.Logger
	mysql  *gorm.DB
}

// SetCache 配置缓存对象。
func SetCache(cache *redis.Client) ContextOption {
	return func(c *Context) {
		c.cache = cache
	}
}

// SetLogger 配置日志处理器。
func SetLogger(logger *zap.Logger) ContextOption {
	return func(c *Context) {
		c.logger = logger
	}
}

// SetMysql 配置数据库对象。
func SetMysql(mysql *gorm.DB) ContextOption {
	return func(c *Context) {
		c.mysql = mysql
	}
}

// NewContext 新建上下文对象。
func NewContext(opts ...ContextOption) *Context {
	context := &Context{}
	for _, opt := range opts {
		opt(context)
	}

	return context
}

// GetCache 获取缓存对象。
func (c *Context) GetCache() *redis.Client {
	return c.cache
}

// GetLogger 获取日志处理器。
func (c *Context) GetLogger() *zap.Logger {
	return c.logger
}

// GetMysql 获取数据库对象。
func (c *Context) GetMysql() *gorm.DB {
	return c.mysql
}
