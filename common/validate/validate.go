package validate

import (
	"gin-vue-admin/model/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strings"
)

func CheckValidate(data interface{}, c *gin.Context) {

	validate := validator.New()
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			response.FailWithMessage(err.Field(), c)
			//fmt.Println(err)//Key: 'Users.Name' Error:Field validation for 'Name' failed on the 'CustomValidationErrors' tag
			return
		}
	}
	//response.FailWithMessage("获取失败", c)
}

func TransTagName(libTans, err interface{}) interface{} {
	switch err.(type) {
	case validator.ValidationErrorsTranslations:
		var errs map[string]string
		errs = make(map[string]string, 0)
		for k, v := range err.(validator.ValidationErrorsTranslations) {
			for key, value := range libTans.(map[string]string) {
				v = strings.Replace(v, key, value, -1)
			}
			errs[k] = v
		}
		return errs
	case string:
		var errs string
		for key, value := range libTans.(map[string]string) {
			errs = strings.Replace(errs, key, value, -1)
		}
		return errs
	default:
		return err
	}
}
