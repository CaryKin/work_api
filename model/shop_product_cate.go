// 自动生成模板ShopProductCate
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type ShopProductCate struct {
	global.GVA_MODEL
	Title string `json:"title" form:"title" gorm:"column:title;comment:标题;type:varchar(50);size:50;"`
	Cover string `json:"cover" form:"cover" gorm:"column:cover;comment:封面图;type:varchar(255);size:255;"`
	Sort  int    `json:"sort" form:"sort" gorm:"column:sort;comment:排序;type:int;size:10;"`
	Pid   int    `json:"pid" form:"pid" gorm:"column:pid;comment:上级id;type:int;size:10;"`
	Tree  string `json:"tree" form:"tree" gorm:"column:tree;comment:树;type:text;"`
}

func (ShopProductCate) TableName() string {
	return "shop_product_cate"
}
