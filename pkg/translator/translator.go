package translator

import (
	"go-server-starter/internal/i18n"
	"go-server-starter/pkg/validator"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	_validator "github.com/go-playground/validator/v10"
	en_trans "github.com/go-playground/validator/v10/translations/en"
	zh_trans "github.com/go-playground/validator/v10/translations/zh"
)

type Translator struct {
	uni *ut.UniversalTranslator
}

// NewTranslator 创建并初始化翻译器
func NewTranslator() (*Translator, error) {
	// 第一个参数是 fallback locale，后续参数是所有支持的 locale
	// 需要同时注册 en 和 zh 才能正常工作
	uni := ut.New(zh.New(), zh.New(), en.New())

	// 注册所有支持的语言翻译
	if v, ok := binding.Validator.Engine().(*_validator.Validate); ok {
		if err := registerTranslations(v, uni); err != nil {
			return nil, err
		}
	}

	return &Translator{uni: uni}, nil
}

// registerTranslations 注册所有语言的翻译
func registerTranslations(v *_validator.Validate, uni *ut.UniversalTranslator) error {
	// 注册中文翻译
	if trans, found := uni.GetTranslator(i18n.LOCALE_ZH); found {
		if err := zh_trans.RegisterDefaultTranslations(v, trans); err != nil {
			return err
		}
		// 注册自定义验证规则的中文翻译
		if err := validator.RegisterTranslations(v, trans, i18n.LOCALE_ZH); err != nil {
			return err
		}
	}

	// 注册英文翻译
	if trans, found := uni.GetTranslator(i18n.LOCALE_EN); found {
		if err := en_trans.RegisterDefaultTranslations(v, trans); err != nil {
			return err
		}
		// 注册自定义验证规则的英文翻译
		if err := validator.RegisterTranslations(v, trans, i18n.LOCALE_EN); err != nil {
			return err
		}
	}

	return nil
}

// GetTranslator 根据 locale 获取对应的翻译器
func (t *Translator) GetTranslator(locale string) (ut.Translator, bool) {
	return t.uni.GetTranslator(locale)
}
