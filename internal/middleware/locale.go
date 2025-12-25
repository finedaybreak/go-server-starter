package middleware

import (
	"go-server-starter/internal/i18n"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	LocaleContextKey = "locale"
)

// Locale 中间件从请求头中提取语言环境并存储在上下文中
// 按以下顺序检查请求头：
// 1. "locale" - 用于显式指定语言环境的自定义请求头
// 2. "Accept-Language" - 用于语言偏好的标准 HTTP 请求头
func Locale() gin.HandlerFunc {
	return func(c *gin.Context) {
		locale := extractLocale(c)
		c.Set(LocaleContextKey, locale)
		c.Next()
	}
}

// extractLocale 从请求头中提取并规范化语言环境
func extractLocale(c *gin.Context) string {
	// 首先，检查自定义的 "locale" 请求头
	locale := c.GetHeader("locale")
	if locale != "" {
		return i18n.NormalizeLocale(locale)
	}

	// 其次，检查 "Accept-Language" 请求头
	acceptLanguage := c.GetHeader("Accept-Language")
	if acceptLanguage != "" {
		// 解析 Accept-Language 请求头 (例如, "zh-CN,zh;q=0.9,en;q=0.8")
		// 取第一个语言偏好
		languages := strings.Split(acceptLanguage, ",")
		if len(languages) > 0 {
			// 移除质量值（如果存在）(例如, "zh;q=0.9" -> "zh")
			lang := strings.Split(languages[0], ";")[0]
			lang = strings.TrimSpace(lang)
			return i18n.NormalizeLocale(lang)
		}
	}

	// 如果没有找到语言环境信息，默认使用 DefaultLocale
	return i18n.DefaultLocale
}
