package request

import "gin-vue-admin/model"

type ShopProductSearch struct{
    model.ShopProduct
    PageInfo
}