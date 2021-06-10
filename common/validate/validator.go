package validate

import (
	"fmt"
	"gin-vue-admin/global"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"

	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

// InitTrans 初始化翻译器
func InitTrans(locale string) (err error) {

	// 修改gin框架中的Validator引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器

		// 第一个参数是备用（fallback）的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		// uni := ut.New(zhT, zhT) 也是可以的
		uni := ut.New(enT, zhT, enT)

		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			label := field.Tag.Get("label")
			if label == "" {
				return field.Tag.Get("json")
			}
			return label
		})

		// locale 通常取决于 http 请求头的 'Accept-Language'
		var ok bool
		// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
		global.GVA_Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		// 注册翻译器
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, global.GVA_Trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, global.GVA_Trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, global.GVA_Trans)
		}
		return
	}

	return
}

// 删除返回的model方法
func RemoveTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

type ValidationErrorsTranslations map[string]string

// 自定义验证错误返回方法
func Translate(err error, request interface{}) (ValidationErrorsTranslations, string) {
	errs, ok := err.(validator.ValidationErrors)

	var errList = make(ValidationErrorsTranslations, len(errs))

	if !ok {
		// 非validator.ValidationErrors类型错误直接返回
		return errList, err.Error()

	} else {
		for _, err := range errs {
			// 获取原来的标签 - json
			fieldName := err.StructField()
			t := reflect.TypeOf(request)
			field, _ := t.FieldByName(fieldName)
			j := field.Tag.Get("json")

			// 修改
			errList[j] = err.Translate(global.GVA_Trans)
		}

		return errList, "验证失败"
	}

}
