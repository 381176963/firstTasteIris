package validates

import (
	"fmt"
	"reflect"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

// 表单验证
// https://github.com/go-playground/validator
var (
	uni           *ut.UniversalTranslator
	Validate      *validator.Validate
	ValidateTrans ut.Translator
)

func init() {
	zh2 := zh.New()
	uni = ut.New(zh2, zh2)
	ValidateTrans, _ = uni.GetTranslator("zh")
	Validate = validator.New()

	// 收集结构体中的comment标签，用于替换英文字段名称，这样返回错误就能展示中文字段名称了
	Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get("comment")
	})

	_ = Validate.RegisterValidation("customTag", customFunc) // hechengyu自己加的自定义验证
	_ = Validate.RegisterTranslation(
		"customTag",
		ValidateTrans,
		func(ut ut.Translator) error {
			return ut.Add("customTag", "{0} 不能包括HTML标签!", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("customTag", fe.Field())
			return t
		},
	) // hechengyu自己加的自定义验证 这几个参数我也没看懂就是套上用 这篇文章里搜到的https://go.gitbaidu.com/gopkg/binding/src/branch/master/default_validator.go

	if err := zhTranslations.RegisterDefaultTranslations(Validate, ValidateTrans); err != nil {
		fmt.Println(fmt.Sprintf("RegisterDefaultTranslations %v", err))
	}
}

func customFunc(fl validator.FieldLevel) bool {
	fmt.Println("打印结果页面输入的用户名: ", fl.Field())
	if fl.Field().String() == "invalid" {
		return false
	}

	return false
}

// 验证表单
func BaseValid(alr interface{}) string {
	var formErrs string
	err := Validate.Struct(alr)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(ValidateTrans) {
			if len(e) > 0 {
				formErrs += e + ";"
			}
		}
	}
	return formErrs
}
