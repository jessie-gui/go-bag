package _interface

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Context 应用上下文接口
type Context interface {
	GetCache() *redis.Client
	GetLogger() *zap.Logger
	GetMysql() *gorm.DB
}
