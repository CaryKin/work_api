package product

import (
	"fmt"
	"gin-vue-admin/common"
	"gin-vue-admin/global"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
)

func GetShopProduct(c *gin.Context) {
	if err := common.InitTrans("zh"); err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
		return
	}
	var info request.ShopProductByCateId
	if err := c.ShouldBind(&info); err != nil {

		errs, ok := err.(validator.ValidationErrors)

		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		c.JSON(http.StatusOK, gin.H{
			"msg": common.RemoveTopStruct(errs.Translate(global.GVA_Trans)),
		})
		return
	}

	/*uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")

	var info request.ShopProductByCateId
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册翻译器
		_= zh_translations.RegisterDefaultTranslations(v, trans)

		//注册一个函数，获取struct tag里自定义的label作为字段名
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name:=fld.Tag.Get("label")
			return name
		})
	}


	if err := c.ShouldBindWith(&info, binding.Query); err != nil {
		errs := err.(validator.ValidationErrors)
		fmt.Println(errs.Translate(trans))
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.Translate(trans)})
		//response.FailWithMessage("获取失败", c)
	}*/
	//common.CheckValidate(c.ShouldBindJSON(&info), c)
	/*if err := c.ShouldBindJSON(&info); err != nil {
		global.GVA_LOG.Error("参数校验不通过", zap.Any("err", err))
		response.FailWithMessage("params error", c)
		return
	}*/

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
