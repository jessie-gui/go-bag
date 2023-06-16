package entity

import (
	"context"

	"github.com/jessie-gui/go-bag/cache"
	_interface "github.com/jessie-gui/go-bag/interface"
)

// Bag 背包对象。
type Bag struct {
	items map[string]string
	roles map[string]any
}

func NewBag(appCtx _interface.Context, uid string) *Bag {
	bag := &Bag{}
	ctx := context.Background()

	bag.items = cache.NewItem(appCtx.GetCache()).GetItems(ctx, appCtx, uid)
	bag.roles = cache.NewRole(appCtx.GetCache()).GetRoles(ctx, appCtx, uid)

	return bag
}

func (b *Bag) ToMap() map[string]map[string]any {
	list := make(map[string]map[string]any)

	// 装备
	list["items"] = make(map[string]any)
	for itemId, items := range b.items {
		list["items"][itemId] = items
	}

	// 英雄
	list["roles"] = b.roles

	return list
}
