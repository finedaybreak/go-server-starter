package i18n

import (
	"strings"
)

const (
	DefaultLocale  = "zh" // 默认语言
	FallbackLocale = "en" // 回退语言
)

// NormalizeLocale 标准化语言字符串到支持的语言
// 示例: "zh-CN", "zh-TW" -> "zh", "en-US", "en-GB" -> "en"
func NormalizeLocale(locale string) string {
	if locale == "" {
		return DefaultLocale
	}
	// 转换为小写并提取语言代码
	locale = strings.ToLower(locale)
	// 映射到支持的语言
	switch locale {
	case "zh", "zh-cn":
		return "zh"
	case "en", "en-us":
		return "en"
	default:
		return DefaultLocale
	}
}
