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

// @Tags ShopProduct
// @Summary 创建ShopProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShopProduct true "创建ShopProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shopProduct/createShopProduct [post]
func CreateShopProduct(c *gin.Context) {
	var shopProduct model.ShopProduct
	_ = c.ShouldBindJSON(&shopProduct)
	if err := service.CreateShopProduct(shopProduct); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags ShopProduct
// @Summary 删除ShopProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShopProduct true "删除ShopProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shopProduct/deleteShopProduct [delete]
func DeleteShopProduct(c *gin.Context) {
	var shopProduct model.ShopProduct
	_ = c.ShouldBindJSON(&shopProduct)
	if err := service.DeleteShopProduct(shopProduct); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ShopProduct
// @Summary 批量删除ShopProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ShopProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /shopProduct/deleteShopProductByIds [delete]
func DeleteShopProductByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteShopProductByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags ShopProduct
// @Summary 更新ShopProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShopProduct true "更新ShopProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /shopProduct/updateShopProduct [put]
func UpdateShopProduct(c *gin.Context) {
	var shopProduct model.ShopProduct
	_ = c.ShouldBindJSON(&shopProduct)
	if err := service.UpdateShopProduct(shopProduct); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags ShopProduct
// @Summary 用id查询ShopProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShopProduct true "用id查询ShopProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /shopProduct/findShopProduct [get]
func FindShopProduct(c *gin.Context) {
	var shopProduct model.ShopProduct
	_ = c.ShouldBindQuery(&shopProduct)
	if err, reshopProduct := service.GetShopProduct(shopProduct.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reshopProduct": reshopProduct}, c)
	}
}

// @Tags ShopProduct
// @Summary 分页获取ShopProduct列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ShopProductSearch true "分页获取ShopProduct列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shopProduct/getShopProductList [get]
func GetShopProductList(c *gin.Context) {
	var pageInfo request.ShopProductSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetShopProductInfoList(pageInfo); err != nil {
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
