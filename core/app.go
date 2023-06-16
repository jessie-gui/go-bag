package core

import (
	"context"

	"github.com/jessie-gui/go-bag/routers"
	"github.com/jessie-gui/x/xserver/xhttp"
	"go.uber.org/zap"
)

type AppOption func(a *App)

// App 应用对象。
type App struct {
	context *Context
}

// SetContext 配置应用上下文对象。
func SetContext(context *Context) AppOption {
	return func(a *App) {
		a.context = context
	}
}

// NewApp 新建应用。
func NewApp(opts ...AppOption) *App {
	app := &App{}
	for _, opt := range opts {
		opt(app)
	}

	return app
}

// Run 启动应用。
func (a *App) Run() {
	server := xhttp.NewServer(
		xhttp.Address(":8080"),
		xhttp.Handler(routers.NewEcho(a.context)),
	)

	if err := server.Start(context.Background()); err != nil {
		a.context.GetLogger().Fatal("服务启动失败:", zap.Error(err))
	}
}
