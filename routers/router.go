package routers

import (
	"github.com/jessie-gui/go-bag/controllers"
	_interface "github.com/jessie-gui/go-bag/interface"
	"github.com/labstack/echo/v4"
)

// NewEcho 构建echo框架路由。
func NewEcho(ctx _interface.Context) *echo.Echo {
	e := echo.New()

	// 基础控制器。
	base := controllers.NewBase(ctx)

	// 背包。
	bagController := controllers.NewBag(base)
	g := e.Group("bag")
	g.GET("/", bagController.List)

	// 装备。
	itemController := controllers.NewItem(base)
	g1 := e.Group("item")
	g1.GET("/list", itemController.ItemList)
	g1.PUT("/add/:id", itemController.AddItem)
	g1.DELETE("/delete", itemController.DeleteItem)
	g1.GET("/get/:id", itemController.GetItem)

	// 英雄。
	roleController := controllers.NewRole(base)
	g2 := e.Group("role")
	g2.GET("/list", roleController.RoleList)
	g2.PUT("/add/:id", roleController.AddRole)
	g2.DELETE("/delete", roleController.DelRole)
	g2.POST("/levelUp", roleController.UpdateRoleLevel)

	// 排行。
	rankController := controllers.NewRank(base)
	g3 := e.Group("rank")
	g3.GET("/list", rankController.Top)
	g3.PUT("/add", rankController.Set)
	g3.GET("/get/:id", rankController.GetRanking)

	return e
}
