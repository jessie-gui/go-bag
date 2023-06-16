package cache

import (
	"context"

	_interface "github.com/jessie-gui/go-bag/interface"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

// Item 装备缓存对象。
type Item struct {
	redis *redis.Client
}

func NewItem(redis *redis.Client) *Item {
	return &Item{
		redis: redis,
	}
}

func (i *Item) GetItems(ctx context.Context, appCtx _interface.Context, uid string) map[string]string {
	key := "u:gi" + uid

	data, err := i.redis.HGetAll(ctx, key).Result()
	if err != nil {
		appCtx.GetLogger().Error("GetItems error:", zap.Error(err))
	}

	return data
}

func (i *Item) AddItem(ctx context.Context, appCtx _interface.Context, uid string, id string, value int64) int64 {
	if value < 1 {
		appCtx.GetLogger().Error("AddItem 数量不能小于1！")
	}

	key := "u:gi" + uid

	num, err := i.redis.HIncrBy(ctx, key, id, value).Result()
	if err != nil {
		appCtx.GetLogger().Error("AddItem error:", zap.Error(err))
	}

	return num
}

func (i *Item) DelItem(ctx context.Context, appCtx _interface.Context, uid string, id string, value int64) int64 {
	if value < 1 {
		appCtx.GetLogger().Error("AddItem 数量不能小于1！")
	}

	key := "u:gi" + uid

	num, err := i.redis.HIncrBy(ctx, key, id, -value).Result()
	if err != nil {
		appCtx.GetLogger().Error("AddItem error:", zap.Error(err))
	}

	return num
}

func (i *Item) GetNum(ctx context.Context, appCtx _interface.Context, uid string, id string) string {
	key := "u:gi" + uid

	num, err := i.redis.HGet(ctx, key, id).Result()
	if err != nil {
		appCtx.GetLogger().Error("Item GetNum error:", zap.Error(err))
	}

	return num
}
