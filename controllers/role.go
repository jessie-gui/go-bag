package controllers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/jessie-gui/go-bag/cache"
	"github.com/jessie-gui/go-bag/library"
	"github.com/jessie-gui/go-bag/model"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// Role 英雄背包控制器。
type Role struct {
	*base
}

// NewRole 新建英雄控制器。
func NewRole(base *base) *Role {
	return &Role{
		base,
	}
}

// RoleList 获取英雄背包列表。
func (r *Role) RoleList(c echo.Context) error {
	r.GetContext().GetLogger().Info("获取英雄列表")

	uid := "123"
	roleList := cache.NewRole(r.GetContext().GetCache()).GetRoles(context.Background(), r.GetContext(), uid)

	return c.JSON(http.StatusOK, roleList)
}

// AddRole 添加英雄。
func (r *Role) AddRole(c echo.Context) error {
	id := c.Param("id")
	uid := "123"

	r.GetContext().GetLogger().Info("添加英雄:", zap.String("role_id", id))

	roleId, err := strconv.Atoi(id)
	if err != nil {
		r.GetContext().GetLogger().Error("AddRole err:", zap.Error(err))
		return nil
	}

	urid := library.NewSnowflake(0).NextId()

	role := &model.BaseRole{}
	r.GetContext().GetMysql().Where("role_id = ?", roleId).Find(role)

	if role != nil {
		cache.NewRole(r.GetContext().GetCache()).AddRole(context.Background(), r.GetContext(), uid, strconv.FormatUint(urid, 10), id)
	}

	return c.JSON(http.StatusOK, role)
}

// DelRole 删除英雄。
func (r *Role) DelRole(c echo.Context) error {
	roleId := c.FormValue("roleId")
	urid := c.FormValue("urid")
	uid := "123"

	num := cache.NewRole(r.GetContext().GetCache()).DelRole(context.Background(), r.GetContext(), uid, urid, roleId, 1)

	return c.JSON(http.StatusOK, num)
}

// UpdateRoleLevel 英雄升级。
func (r *Role) UpdateRoleLevel(c echo.Context) error {
	value := c.FormValue("value")
	urid := c.FormValue("urid")
	uid := "123"

	valueInt, err := strconv.Atoi(value)
	if err != nil {
		r.GetContext().GetLogger().Error("UpdateRoleLevel err:", zap.Error(err))
	}

	num := cache.NewRole(r.GetContext().GetCache()).UpdateLevel(context.Background(), r.GetContext(), uid, urid, int64(valueInt))

	return c.JSON(http.StatusOK, num)
}
