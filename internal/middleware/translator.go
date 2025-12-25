package middleware

import (
	"go-server-starter/pkg/translator"

	"github.com/gin-gonic/gin"
)

// Translations 翻译中间件,根据请求头的 locale 设置翻译器
func Translations(t *translator.Translator) gin.HandlerFunc {
	return func(c *gin.Context) {
		locale := c.GetHeader("locale")
		trans, _ := t.GetTranslator(locale)
		c.Set("trans", trans)
		c.Next()
	}
}
