// 自动生成模板ShopProduct
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type ShopProduct struct {
	global.GVA_MODEL
	CateId       int    `json:"cateId" form:"cateId" gorm:"column:cate_id;comment:商品分类编号;type:int;size:10;"`
	Sketch       string `json:"sketch" form:"sketch" gorm:"column:sketch;comment:简述;type:varchar(200);size:200;"`
	Intro        string `json:"intro" form:"intro" gorm:"column:intro;comment:商品描述;type:text;"`
	TotalSales   int    `json:"totalSales" form:"totalSales" gorm:"column:total_sales;comment:总销量;type:int;size:10;"`
	WarningStock int    `json:"warningStock" form:"warningStock" gorm:"column:warning_stock;comment:库存警告;type:int;size:10;"`
	Covers       string `json:"covers" form:"covers" gorm:"column:covers;comment:幻灯片;type:text;"`
	Status       *bool  `json:"status" form:"status" gorm:"column:status;comment:商品状态 [0下架，1正常];type:tinyint;"`
}

func (ShopProduct) TableName() string {
	return "shop_product"
}
