package controllers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/jessie-gui/go-bag/cache"
	"github.com/jessie-gui/go-bag/constant"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// Rank 排行榜控制器。
type Rank struct {
	*base
}

// NewRank 新建排行榜控制器。
func NewRank(base *base) *Rank {
	return &Rank{
		base,
	}
}

// Top 获取排行榜。
func (r *Rank) Top(c echo.Context) error {
	r.GetContext().GetLogger().Info("获取英雄战力排行榜")

	list := cache.NewRank(r.GetContext().GetCache()).Top(context.Background(), r.GetContext(), constant.HeroCombatRank, 0, 50)

	return c.JSON(http.StatusOK, list)
}

// Set 设置排行。
func (r *Rank) Set(c echo.Context) error {
	r.GetContext().GetLogger().Info("设置玩家英雄战力排行榜积分")

	score := c.FormValue("score")
	uid := "123"

	scoreFloat, err := strconv.ParseFloat(score, 64)
	if err != nil {
		r.GetContext().GetLogger().Error("set rank err:", zap.Error(err))
	}

	err1 := cache.NewRank(r.GetContext().GetCache()).Set(context.Background(), r.GetContext(), uid, scoreFloat, constant.HeroCombatRank)

	return c.JSON(http.StatusOK, map[string]any{
		"ok":  1,
		"err": err1,
	})
}

// GetRanking 获取个人排行。
func (r *Rank) GetRanking(c echo.Context) error {
	r.GetContext().GetLogger().Info("获取玩家英雄战力排行榜积分")

	uid := c.Param("id")

	rScore := cache.NewRank(r.GetContext().GetCache()).GetRanking(context.Background(), r.GetContext(), constant.HeroCombatRank, uid)

	return c.JSON(http.StatusOK, rScore)
}
