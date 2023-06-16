package cache

import (
	"context"
	"fmt"

	_interface "github.com/jessie-gui/go-bag/interface"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

// Rank 排行榜缓存对象。
type Rank struct {
	redis *redis.Client
}

// RankMember 排行榜成员缓存对象。
type RankMember struct {
	Uid   string
	Score float64
}

// NewRank 新建排行榜缓存对象。
func NewRank(redis *redis.Client) *Rank {
	return &Rank{
		redis: redis,
	}
}

// Set 设置排行。
func (r *Rank) Set(ctx context.Context, appCtx _interface.Context, uid string, score float64, ty string) error {
	k := "u:rank:" + ty

	fmt.Println(score, uid)
	return r.redis.ZAdd(ctx, k, redis.Z{Score: score, Member: uid}).Err()
}

// Top 获取排行列表。
func (r *Rank) Top(ctx context.Context, appCtx _interface.Context, ty string, start int64, end int64) []RankMember {
	k := "u:rank:" + ty

	res, err := r.redis.ZRevRangeWithScores(ctx, k, start, end).Result()
	if err != nil {
		appCtx.GetLogger().Error("rank top error:", zap.Error(err))
		return nil
	}

	members := make([]RankMember, 0, len(res))
	for _, member := range res {
		mem := RankMember{
			Uid:   member.Member.(string),
			Score: member.Score,
		}
		members = append(members, mem)
	}

	return members
}

// GetRanking 获取个人排行。
func (r *Rank) GetRanking(ctx context.Context, appCtx _interface.Context, ty string, uid string) int64 {
	k := "u:rank:" + ty

	score, err := r.redis.ZScore(ctx, k, uid).Result()
	if err != nil {
		appCtx.GetLogger().Error("rank GetRanking error:", zap.Error(err))
	}

	if score > 0 {
		rank, err1 := r.redis.ZRevRank(ctx, k, uid).Result()
		if err1 != nil {
			appCtx.GetLogger().Error("rank GetRanking error1:", zap.Error(err1))
			return 0
		}

		return rank
	}

	return 0
}
