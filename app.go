package main

import "github.com/jessie-gui/go-bag/core"

/**
 *
 *
 * @author        Jessie Gui <guijiaxian@gmail.com>
 * @version       1.0.0
 * @copyright (c) 2022, Jessie Gui
 */
func main() {
	// 初始化配置。
	config := core.NewConfig()

	// 构建应用上下文。
	context := core.NewContext(
		core.SetLogger(core.NewLogger()),
		core.SetCache(core.NewRedis(
			core.SetAddr(config.Redis.Address),
		)),
		core.SetMysql(core.NewMysql(
			core.SetUser(config.Mysql.User),
			core.SetAddress(config.Mysql.Address),
			core.SetPwd(config.Mysql.Pwd),
			core.SetDbBase(config.Mysql.DbBase)).Init(),
		),
	)

	// 启动应用。
	core.NewApp(core.SetContext(context)).Run()
}
