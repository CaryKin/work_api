package product

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetShopProduct(c *gin.Context) {
	var info request.ShopProductByCateId

	if err := c.ShouldBindJSON(&info); err != nil {
		global.GVA_LOG.Error("参数校验不通过", zap.Any("err", err))
		response.FailWithMessage("params error", c)
		return
	}

	if err, list, total := service.GetShopProductInfoByCate(info); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     info.Page,
			PageSize: info.PageSize,
		}, "获取成功", c)
	}
}
