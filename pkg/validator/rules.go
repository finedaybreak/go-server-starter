package validator

import "github.com/go-playground/validator/v10"

type Rule struct {
	Tag  string
	Fn   validator.Func
	I18n map[string]string // 多语言翻译
}

var Rules = []Rule{
	{
		Tag: "mobile",
		Fn:  validateMobile,
		I18n: map[string]string{
			"zh": "{0}必须是有效的手机号码",
			"en": "{0} must be a valid mobile number",
		},
	},
	{
		Tag: "username",
		Fn:  validateUsername,
		I18n: map[string]string{
			"zh": "{0}必须是4-20位的字母、数字或下划线",
			"en": "{0} must be 4-20 characters of letters, numbers or underscores",
		},
	},
}

func validateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	if len(mobile) != 11 {
		return false
	}
	return mobile[0] == '1'
}

func validateUsername(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	if len(username) < 4 || len(username) > 20 {
		return false
	}
	for _, r := range username {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') ||
			(r >= '0' && r <= '9') || r == '_') {
			return false
		}
	}
	return true
}
