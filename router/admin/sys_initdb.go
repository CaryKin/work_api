package admin

import (
	"gin-vue-admin/backend/v1"
	"github.com/gin-gonic/gin"
)

func InitInitRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("init")
	{
		ApiRouter.POST("initdb", v1.InitDB)   // 创建Api
		ApiRouter.POST("checkdb", v1.CheckDB) // 创建Api
	}
}
