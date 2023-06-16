package cache

import (
	"context"
	"fmt"
	"strconv"

	_interface "github.com/jessie-gui/go-bag/interface"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

// Role 英雄背包缓存对象。
type role struct {
	redis *redis.Client
}

// NewRole 新建英雄背包缓存对象。
func NewRole(redis *redis.Client) *role {
	return &role{
		redis: redis,
	}
}

// GetRoles 获取英雄列表。
func (r *role) GetRoles(ctx context.Context, appCtx _interface.Context, uid string) map[string]any {
	dr := make(map[string]any)
	k1 := "u:gr:urids:" + uid
	k2 := "u:gr:ids:" + uid

	urids, err := r.redis.SMembers(ctx, k1).Result()
	if err != nil {
		appCtx.GetLogger().Error("GetRoles error:", zap.Error(err))
		return nil
	}

	ids, err1 := r.redis.HGetAll(ctx, k2).Result()
	if err1 != nil {
		appCtx.GetLogger().Error("GetRoles error1:", zap.Error(err1))
		return nil
	}

	dr["ids"] = ids

	dts := make(map[string]map[string]string)
	for _, urid := range urids {
		k3 := "u:gr:dt:" + urid

		dt, err2 := r.redis.HGetAll(ctx, k3).Result()
		if err2 != nil {
			appCtx.GetLogger().Error("GetRoles error2:", zap.Error(err2))
			return nil
		}

		dts[urid] = dt
	}

	dr["dt"] = dts
	dr["urids"] = urids

	return dr
}

// AddRole 添加英雄。
func (r *role) AddRole(ctx context.Context, appCtx _interface.Context, uid string, urid string, roleId string) int64 {
	k1 := "u:gr:urids:" + uid
	k2 := "u:gr:ids:" + uid
	k3 := "u:gr:dt:" + urid

	num, err := r.redis.Exists(ctx, k3).Result()
	if err != nil {
		appCtx.GetLogger().Error("AddRole error1:", zap.Error(err))
		return 0
	}

	if num > 0 {
		appCtx.GetLogger().Error("AddRole error2: 英雄ID重复，添加英雄失败！")
		return 0
	}

	_, err1 := r.redis.SAdd(ctx, k1, urid).Result()
	if err1 != nil {
		appCtx.GetLogger().Error("AddRole error3:", zap.Error(err1))
		return 0
	}

	roleNum, err2 := r.redis.HIncrBy(ctx, k2, roleId, 1).Result()
	if err2 != nil {
		appCtx.GetLogger().Error("AddRole error4:", zap.Error(err2))
		return 0
	}

	r.redis.HMSet(ctx, k3, map[string]string{
		"role_id": roleId,
	})

	r.redis.HIncrBy(ctx, k3, "level", 1)

	return roleNum
}

// DelRole 删除英雄。
func (r *role) DelRole(ctx context.Context, appCtx _interface.Context, uid string, urid string, roleId string, num int) int64 {
	k1 := "u:gr:urids:" + uid
	k2 := "u:gr:ids:" + uid
	k3 := "u:gr:dt:" + urid

	uridInt, err := strconv.Atoi(urid)
	if err != nil {
		appCtx.GetLogger().Error("DelRole error:", zap.Error(err))
	}

	roleIdInt, err0 := strconv.Atoi(roleId)
	if err0 != nil {
		appCtx.GetLogger().Error("DelRole error0:", zap.Error(err0))
	}

	if uridInt > 0 {
		d3, err1 := r.redis.HMGet(ctx, k3, "role_id").Result()
		if err1 != nil {
			appCtx.GetLogger().Error("DelRole error1:", zap.Error(err1))
		}

		rid := d3[0]

		r.redis.SRem(ctx, k1, urid)
		r.redis.Del(ctx, k3)

		rNum, err2 := r.redis.HIncrBy(ctx, k2, fmt.Sprintf("%v", rid), -1).Result()
		if err2 != nil {
			appCtx.GetLogger().Error("DelRole error2:", zap.Error(err2))
		}

		if rNum <= 0 {
			r.redis.HDel(ctx, k2, fmt.Sprintf("%v", rid))
		}

		return rNum
	} else if roleIdInt > 0 {
		lNum, err3 := r.redis.HGet(ctx, k2, roleId).Result()
		if err3 != nil {
			appCtx.GetLogger().Error("DelRole error3:", zap.Error(err3))
		}

		lNumInt, err4 := strconv.Atoi(lNum)
		if err4 != nil {
			appCtx.GetLogger().Error("DelRole error4:", zap.Error(err4))
		}

		if lNumInt < num {
			appCtx.GetLogger().Error("DelRole error3: 英雄数量不足")
		}

		d1, err5 := r.redis.SMembers(ctx, k1).Result()
		if err5 != nil {
			appCtx.GetLogger().Error("DelRole error5:", zap.Error(err5))
		}

		n := 0
		for _, u := range d1 {
			d3, err6 := r.redis.HMGet(ctx, "u:gr:dt:"+u, "role_id").Result()
			if err6 != nil {
				appCtx.GetLogger().Error("DelRole error6:", zap.Error(err6))
			}

			rid := d3[0]

			if roleId == rid {
				r.redis.SRem(ctx, k1, u)
				r.redis.Del(ctx, "u:gr:dt:"+u)
				n++
			}
		}

		rNum, err7 := r.redis.HIncrBy(ctx, k2, roleId, int64(-n)).Result()
		if err7 != nil {
			appCtx.GetLogger().Error("DelRole error7:", zap.Error(err7))
		}

		if rNum <= 0 {
			r.redis.HDel(ctx, k2, roleId)
		}

		return rNum
	}

	return 0
}

// UpdateLevel 英雄升级。
func (r *role) UpdateLevel(ctx context.Context, appCtx _interface.Context, uid string, urid string, value int64) int64 {
	k3 := "u:gr:dt:" + urid

	v, err := r.redis.Exists(ctx, k3).Result()
	if err != nil {
		appCtx.GetLogger().Error("UpdateLevel error:", zap.Error(err))
	}

	if v > 0 {
		rNum, err1 := r.redis.HIncrBy(ctx, k3, "level", value).Result()
		if err1 != nil {
			appCtx.GetLogger().Error("UpdateLevel error1:", zap.Error(err1))
		}

		return rNum
	}

	return 0
}
