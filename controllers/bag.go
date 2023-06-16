package controllers

import (
	"net/http"

	"github.com/jessie-gui/go-bag/entity"
	"github.com/labstack/echo/v4"
)

// Bag 背包控制器。
type Bag struct {
	*base
}

// NewBag 新建背包控制器。
func NewBag(base *base) *Bag {
	return &Bag{
		base,
	}
}

// List 获取背包列表。
func (b *Bag) List(c echo.Context) error {
	b.GetContext().GetLogger().Info("获取背包列表")

	uid := "123"

	bag := entity.NewBag(b.GetContext(), uid).ToMap()

	return c.JSON(http.StatusOK, bag)
}
