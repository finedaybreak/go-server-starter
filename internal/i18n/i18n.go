package i18n

import (
	"fmt"
	"strings"
)

// Text represents a translatable message with multiple locale support
type Text struct {
	En string
	Zh string
}

const (
	LOCALE_ZH      = "zh"
	LOCALE_EN      = "en"
	LOCALE_ZH_CN   = "zh-cn"
	LOCALE_EN_US   = "en-us"
	DEFAULT_LOCALE = LOCALE_ZH
)

func NormalizeLocale(locale string) string {
	switch strings.ToLower(locale) {
	case LOCALE_ZH, LOCALE_ZH_CN:
		return LOCALE_ZH
	case LOCALE_EN, LOCALE_EN_US:
		return LOCALE_EN
	}
	return DEFAULT_LOCALE
}

// T returns the translated message for the given locale
func (m Text) T(locale string, params ...map[string]string) string {
	var message string
	switch NormalizeLocale(locale) {
	case LOCALE_ZH:
		message = m.Zh
	case LOCALE_EN:
		message = m.En
	default:
		message = m.Zh // default to Chinese
	}

	// Replace placeholders with params
	if len(params) > 0 && params[0] != nil {
		for k, v := range params[0] {
			placeholder := fmt.Sprintf("{%s}", k)
			message = strings.ReplaceAll(message, placeholder, v)
		}
	}

	return message
}

// Common non-exception messages
var (
	EchoHello   = Text{En: "Hello, {name}!", Zh: "你好, {name}!"}
	RespSuccess = Text{En: "Success", Zh: "成功"}
)
