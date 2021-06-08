package product

import (
	"gin-vue-admin/api/v1/product"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitShopProductRouter(Router *gin.RouterGroup) {
	ShopProductCateRouter := Router.Group("product").Use(middleware.OperationRecord())
	{
		//ShopProductCateRouter.POST("createShopProductCate", product.CreateShopProductCate)   // 新建ShopProductCate
		//ShopProductCateRouter.DELETE("deleteShopProductCate", product.DeleteShopProductCate) // 删除ShopProductCate
		//ShopProductCateRouter.DELETE("deleteShopProductCateByIds", product.DeleteShopProductCateByIds) // 批量删除ShopProductCate
		//ShopProductCateRouter.PUT("updateShopProductCate", product.UpdateShopProductCate)    // 更新ShopProductCate
		//ShopProductCateRouter.GET("findShopProductCate", product.FindShopProductCate)        // 根据ID获取ShopProductCate
		ShopProductCateRouter.GET("get", product.GetShopProduct) // 获取ShopProductCate列表
	}
}
