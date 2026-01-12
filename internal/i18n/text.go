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

// T returns the translated message for the given locale
func (m Text) T(locale string, params ...map[string]string) string {
	var message string
	switch strings.ToLower(locale) {
	case "zh", "zh-cn":
		message = m.Zh
	case "en", "en-us":
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
