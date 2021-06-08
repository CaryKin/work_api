package request

import "gin-vue-admin/model"

type ShopProductSearch struct {
	model.ShopProduct
	PageInfo
}

type ShopProductByCateId struct {
	ID uint `json:"id" binding:"required"`
	PageInfo
}

func AttributeLabels() {

}
