package controllers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/jessie-gui/go-bag/cache"
	"github.com/jessie-gui/go-bag/model"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// Item 装备控制器。
type Item struct {
	*base
}

// NewItem 新建装备控制器。
func NewItem(base *base) *Item {
	return &Item{
		base,
	}
}

// ItemList 装备背包列表。
func (i *Item) ItemList(c echo.Context) error {
	i.base.GetContext().GetLogger().Info("获取装备列表")

	uid := "123"
	itemList := cache.NewItem(i.base.GetContext().GetCache()).GetItems(context.Background(), i.base.GetContext(), uid)

	return c.JSON(http.StatusOK, itemList)
}

// AddItem 添加装备。
func (i *Item) AddItem(c echo.Context) error {
	id := c.Param("id")
	uid := "123"

	i.base.GetContext().GetLogger().Info("添加装备:", zap.String("item_id", id))

	itemId, err := strconv.Atoi(id)
	if err != nil {
		i.base.GetContext().GetLogger().Error("AddItem err:", zap.Error(err))
		return nil
	}

	item := &model.BaseItem{}
	i.base.GetContext().GetMysql().Where("item_id = ?", itemId).Find(item)

	if item != nil {
		cache.NewItem(i.base.GetContext().GetCache()).AddItem(context.Background(), i.base.GetContext(), uid, id, 1)
	}

	return c.JSON(http.StatusOK, item)
}

// DeleteItem 删除装备。
func (i *Item) DeleteItem(c echo.Context) error {
	itemId := c.FormValue("id")
	value := c.FormValue("value")
	uid := "123"

	i.GetContext().GetLogger().Info("消耗装备:", zap.String("item", itemId))

	itemNum, err := strconv.Atoi(value)
	if err != nil {
		i.GetContext().GetLogger().Error("消耗装备错误:", zap.Error(err))
	}

	num := cache.NewItem(i.GetContext().GetCache()).DelItem(context.Background(), i.GetContext(), uid, itemId, int64(itemNum))

	return c.JSON(http.StatusOK, map[string]int64{
		itemId: num,
	})
}

// GetItem 获取装备详情。
func (i *Item) GetItem(c echo.Context) error {
	id := c.Param("id")
	uid := "123"

	i.GetContext().GetLogger().Info("获取装备:", zap.String("item_id", id))

	num := cache.NewItem(i.GetContext().GetCache()).GetNum(context.Background(), i.GetContext(), uid, id)

	return c.JSON(http.StatusOK, map[string]string{
		id: num,
	})
}
