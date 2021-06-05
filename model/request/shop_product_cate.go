package request

import "gin-vue-admin/model"

type ShopProductCateSearch struct{
    model.ShopProductCate
    PageInfo
}