package controllers

import (
	_interface "github.com/jessie-gui/go-bag/interface"
)

// 基础控制器。
type base struct {
	context _interface.Context
}

// NewBase 新建基础控制器。
func NewBase(context _interface.Context) *base {
	return &base{
		context: context,
	}
}

// GetContext 获取应用上下文对象。
func (b *base) GetContext() _interface.Context {
	return b.context
}
