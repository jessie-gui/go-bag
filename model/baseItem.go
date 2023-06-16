package model

// BaseItem 装备表。
type BaseItem struct {
	ItemId            int
	ItemName          string
	ItemNo            string
	ItemDesc          string
	ItemQuality       int
	ItemStar          int
	ItemUseType       int
	ItemPrice         int
	ItemEffect        string
	ItemSuitsId       int
	IsCompose         int
	IsDrop            int
	ComposeAttr       string
	DropAttr          string
	FuncId            string
	ItemProEffectShow string
	ItemProEffect     string
	ItemProfession    int
	TopItemId         int
	TopItemEffectShow string
	TopItemEffect     string
}

type Item struct {
	Id  string
	Num int
}
