package product

import (
	"gin-vue-admin/common/validate"
	"gin-vue-admin/global"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetShopProduct(c *gin.Context) {

	var info request.ShopProductByCateId
	//验证参数
	if err := c.ShouldBind(&info); err != nil {
		errs := validate.Translate(err, info, c)
		response.FailWithDetailed(errs, "参数验证失败", c)
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
