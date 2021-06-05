package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateShopProductCate
//@description: 创建ShopProductCate记录
//@param: shopProductCate model.ShopProductCate
//@return: err error

func CreateShopProductCate(shopProductCate model.ShopProductCate) (err error) {
	err = global.GVA_DB.Create(&shopProductCate).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteShopProductCate
//@description: 删除ShopProductCate记录
//@param: shopProductCate model.ShopProductCate
//@return: err error

func DeleteShopProductCate(shopProductCate model.ShopProductCate) (err error) {
	err = global.GVA_DB.Delete(&shopProductCate).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteShopProductCateByIds
//@description: 批量删除ShopProductCate记录
//@param: ids request.IdsReq
//@return: err error

func DeleteShopProductCateByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.ShopProductCate{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateShopProductCate
//@description: 更新ShopProductCate记录
//@param: shopProductCate *model.ShopProductCate
//@return: err error

func UpdateShopProductCate(shopProductCate model.ShopProductCate) (err error) {
	err = global.GVA_DB.Save(&shopProductCate).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetShopProductCate
//@description: 根据id获取ShopProductCate记录
//@param: id uint
//@return: err error, shopProductCate model.ShopProductCate

func GetShopProductCate(id uint) (err error, shopProductCate model.ShopProductCate) {
	err = global.GVA_DB.Where("id = ?", id).First(&shopProductCate).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetShopProductCateInfoList
//@description: 分页获取ShopProductCate记录
//@param: info request.ShopProductCateSearch
//@return: err error, list interface{}, total int64

func GetShopProductCateInfoList(info request.ShopProductCateSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.ShopProductCate{})
    var shopProductCates []model.ShopProductCate
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&shopProductCates).Error
	return err, shopProductCates, total
}