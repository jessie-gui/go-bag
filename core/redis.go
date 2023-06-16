package core

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

type RedisOption func(r *Redis)

// Redis 缓存对象。
type Redis struct {
	addr     string
	passWord string
	db       int
	redis    *redis.Client
}

func SetAddr(addr string) RedisOption {
	return func(r *Redis) {
		r.addr = addr
	}
}

func SetPassWord(passWord string) RedisOption {
	return func(r *Redis) {
		r.passWord = passWord
	}
}

func SetDb(db int) RedisOption {
	return func(r *Redis) {
		r.db = db
	}
}

func NewRedis(opts ...RedisOption) *redis.Client {
	r := &Redis{}
	for _, opt := range opts {
		opt(r)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     r.addr,     // use default Addr
		Password: r.passWord, // no password set
		DB:       r.db,       // use default DB
	})

	ctx := context.Background()

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal("redis 连接错误:", err)
	}

	r.redis = rdb

	log.Println("redis初始化完成:", pong)

	return rdb
}
