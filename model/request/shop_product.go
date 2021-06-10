package request

import "gin-vue-admin/model"

type ShopProductSearch struct {
	model.ShopProduct
	PageInfo
}

type ShopProductByCateId struct {
	ID uint `form:"id" json:"id" binding:"required" label:"主键"`
	PageInfo
}
