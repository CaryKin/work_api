package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags ShopProductCate
// @Summary 创建ShopProductCate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShopProductCate true "创建ShopProductCate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shopProductCate/createShopProductCate [post]
func CreateShopProductCate(c *gin.Context) {
	var shopProductCate model.ShopProductCate
	_ = c.ShouldBindJSON(&shopProductCate)
	if err := service.CreateShopProductCate(shopProductCate); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags ShopProductCate
// @Summary 删除ShopProductCate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShopProductCate true "删除ShopProductCate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shopProductCate/deleteShopProductCate [delete]
func DeleteShopProductCate(c *gin.Context) {
	var shopProductCate model.ShopProductCate
	_ = c.ShouldBindJSON(&shopProductCate)
	if err := service.DeleteShopProductCate(shopProductCate); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ShopProductCate
// @Summary 批量删除ShopProductCate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ShopProductCate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /shopProductCate/deleteShopProductCateByIds [delete]
func DeleteShopProductCateByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteShopProductCateByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags ShopProductCate
// @Summary 更新ShopProductCate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShopProductCate true "更新ShopProductCate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /shopProductCate/updateShopProductCate [put]
func UpdateShopProductCate(c *gin.Context) {
	var shopProductCate model.ShopProductCate
	_ = c.ShouldBindJSON(&shopProductCate)
	if err := service.UpdateShopProductCate(shopProductCate); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags ShopProductCate
// @Summary 用id查询ShopProductCate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShopProductCate true "用id查询ShopProductCate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /shopProductCate/findShopProductCate [get]
func FindShopProductCate(c *gin.Context) {
	var shopProductCate model.ShopProductCate
	_ = c.ShouldBindQuery(&shopProductCate)
	if err, reshopProductCate := service.GetShopProductCate(shopProductCate.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reshopProductCate": reshopProductCate}, c)
	}
}

// @Tags ShopProductCate
// @Summary 分页获取ShopProductCate列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ShopProductCateSearch true "分页获取ShopProductCate列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shopProductCate/getShopProductCateList [get]
func GetShopProductCateList(c *gin.Context) {
	var pageInfo request.ShopProductCateSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetShopProductCateInfoPage(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
