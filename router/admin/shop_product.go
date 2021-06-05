package admin

import (
	"gin-vue-admin/backend/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitShopProductRouter(Router *gin.RouterGroup) {
	ShopProductRouter := Router.Group("shopProduct").Use(middleware.OperationRecord())
	{
		ShopProductRouter.POST("createShopProduct", v1.CreateShopProduct)   // 新建ShopProduct
		ShopProductRouter.DELETE("deleteShopProduct", v1.DeleteShopProduct) // 删除ShopProduct
		ShopProductRouter.DELETE("deleteShopProductByIds", v1.DeleteShopProductByIds) // 批量删除ShopProduct
		ShopProductRouter.PUT("updateShopProduct", v1.UpdateShopProduct)    // 更新ShopProduct
		ShopProductRouter.GET("findShopProduct", v1.FindShopProduct)        // 根据ID获取ShopProduct
		ShopProductRouter.GET("getShopProductList", v1.GetShopProductList)  // 获取ShopProduct列表
	}
}
