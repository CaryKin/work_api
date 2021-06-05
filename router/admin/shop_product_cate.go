package admin

import (
	"gin-vue-admin/backend/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitShopProductCateRouter(Router *gin.RouterGroup) {
	ShopProductCateRouter := Router.Group("shopProductCate").Use(middleware.OperationRecord())
	{
		ShopProductCateRouter.POST("createShopProductCate", v1.CreateShopProductCate)   // 新建ShopProductCate
		ShopProductCateRouter.DELETE("deleteShopProductCate", v1.DeleteShopProductCate) // 删除ShopProductCate
		ShopProductCateRouter.DELETE("deleteShopProductCateByIds", v1.DeleteShopProductCateByIds) // 批量删除ShopProductCate
		ShopProductCateRouter.PUT("updateShopProductCate", v1.UpdateShopProductCate)    // 更新ShopProductCate
		ShopProductCateRouter.GET("findShopProductCate", v1.FindShopProductCate)        // 根据ID获取ShopProductCate
		ShopProductCateRouter.GET("getShopProductCateList", v1.GetShopProductCateList)  // 获取ShopProductCate列表
	}
}
