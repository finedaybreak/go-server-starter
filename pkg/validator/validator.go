package validator

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// Init 初始化验证器配置
func Init() error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册 tag name 函数
		v.RegisterTagNameFunc(jsonTagNameFunc)

		// 注册自定义验证规则
		if err := RegisterCustomRules(v); err != nil {
			return err
		}

		return nil
	}
	return nil
}

// jsonTagNameFunc 从 json 或 form tag 中提取字段名
func jsonTagNameFunc(fld reflect.StructField) string {
	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
	if name == "-" || name == "" {
		name = strings.SplitN(fld.Tag.Get("form"), ",", 2)[0]
	}
	if name == "-" {
		return ""
	}
	return name
}

// RegisterCustomRules 注册所有自定义验证规则
func RegisterCustomRules(v *validator.Validate) error {
	for _, rule := range Rules {
		if err := v.RegisterValidation(rule.Tag, rule.Fn); err != nil {
			return err
		}
	}
	return nil
}

// RegisterTranslations 注册自定义验证规则的翻译
func RegisterTranslations(v *validator.Validate, trans ut.Translator, locale string) error {
	for _, rule := range Rules {
		message, ok := rule.I18n[locale]
		if !ok {
			message = rule.I18n["en"]
		}

		// 避免闭包陷阱
		tag := rule.Tag
		msg := message

		err := v.RegisterTranslation(
			tag,
			trans,
			func(ut ut.Translator) error {
				return ut.Add(tag, msg, false)
			},
			func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T(fe.Tag(), fe.Field())
				return t
			},
		)
		if err != nil {
			return err
		}
	}
	return nil
}
