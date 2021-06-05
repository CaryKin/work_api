package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateShopProduct
//@description: 创建ShopProduct记录
//@param: shopProduct model.ShopProduct
//@return: err error

func CreateShopProduct(shopProduct model.ShopProduct) (err error) {
	err = global.GVA_DB.Create(&shopProduct).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteShopProduct
//@description: 删除ShopProduct记录
//@param: shopProduct model.ShopProduct
//@return: err error

func DeleteShopProduct(shopProduct model.ShopProduct) (err error) {
	err = global.GVA_DB.Delete(&shopProduct).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteShopProductByIds
//@description: 批量删除ShopProduct记录
//@param: ids request.IdsReq
//@return: err error

func DeleteShopProductByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.ShopProduct{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateShopProduct
//@description: 更新ShopProduct记录
//@param: shopProduct *model.ShopProduct
//@return: err error

func UpdateShopProduct(shopProduct model.ShopProduct) (err error) {
	err = global.GVA_DB.Save(&shopProduct).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetShopProduct
//@description: 根据id获取ShopProduct记录
//@param: id uint
//@return: err error, shopProduct model.ShopProduct

func GetShopProduct(id uint) (err error, shopProduct model.ShopProduct) {
	err = global.GVA_DB.Where("id = ?", id).First(&shopProduct).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetShopProductInfoList
//@description: 分页获取ShopProduct记录
//@param: info request.ShopProductSearch
//@return: err error, list interface{}, total int64

func GetShopProductInfoList(info request.ShopProductSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.ShopProduct{})
    var shopProducts []model.ShopProduct
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&shopProducts).Error
	return err, shopProducts, total
}