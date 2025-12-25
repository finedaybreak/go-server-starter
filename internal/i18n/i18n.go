package i18n

import (
	"embed"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
)

//go:embed locales/*.json
var localesFS embed.FS

const (
	DefaultLocale  = "zh" // 默认语言
	FallbackLocale = "en" // 回退语言
)

// I18n 国际化管理器
type I18n struct {
	messages map[string]map[string]string // locale -> key -> message
	mu       sync.RWMutex
}

var (
	instance *I18n
	once     sync.Once
)

// Init 初始化 i18n 实例（单例模式）
func Init() (*I18n, error) {
	var err error
	once.Do(func() {
		instance = &I18n{
			messages: make(map[string]map[string]string),
		}
		err = instance.loadMessages()
	})
	return instance, err
}

// GetInstance 返回 I18n 的单例实例
func GetInstance() *I18n {
	if instance == nil {
		panic("i18n not initialized, call Init() first")
	}
	return instance
}

// loadMessages 从嵌入的文件系统中加载所有翻译文件
func (i *I18n) loadMessages() error {
	// 加载支持的语言
	locales := []string{"zh", "en"}

	for _, locale := range locales {
		filename := fmt.Sprintf("locales/%s.json", locale)
		data, err := localesFS.ReadFile(filename)
		if err != nil {
			return fmt.Errorf("failed to read locale file %s: %w", filename, err)
		}

		var messages map[string]string
		if err := json.Unmarshal(data, &messages); err != nil {
			return fmt.Errorf("failed to parse locale file %s: %w", filename, err)
		}

		i.mu.Lock()
		i.messages[locale] = messages
		i.mu.Unlock()
	}

	return nil
}

// T 翻译消息键，支持可选参数替换
// params 应该是一个占位符名称到值的映射
// 示例: T("zh", "user.registered", map[string]string{"name": "Alice"})
func (i *I18n) T(locale, key string, params ...map[string]string) string {
	i.mu.RLock()
	defer i.mu.RUnlock()

	// 获取请求语言的消息
	message, ok := i.matchMessage(locale, key)
	if !ok {
		// 如果未找到，回退到默认语言
		message, ok = i.matchMessage(FallbackLocale, key)
		if !ok {
			// 如果没有找到翻译，返回 key 本身
			return key
		}
	}

	// 如果提供了参数，替换占位符
	if len(params) > 0 && params[0] != nil {
		for k, v := range params[0] {
			placeholder := fmt.Sprintf("{%s}", k)
			message = strings.ReplaceAll(message, placeholder, v)
		}
	}

	return message
}

// matchMessage 获取特定语言和键的消息
func (i *I18n) matchMessage(locale, key string) (string, bool) {
	localeMessages, ok := i.messages[locale]
	if !ok {
		return "", false
	}

	message, ok := localeMessages[key]
	return message, ok
}

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
