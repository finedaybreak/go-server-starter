package exception

import (
	"go-server-starter/internal/i18n"
)

type ExceptionScope struct {
	name     string
	baseCode int
	counter  int
}

var (
	Common   = &ExceptionScope{name: "common", baseCode: 1000}
	User     = &ExceptionScope{name: "user", baseCode: 20000}
	UserRole = &ExceptionScope{name: "user_role", baseCode: 21000}
)

// New creates a new exception with auto-incrementing code within the module
func (m *ExceptionScope) New(statusCode int, message string, i18nMsg i18n.Text) *Exception {
	m.counter++
	code := m.baseCode + m.counter
	return register(statusCode, code, message, i18nMsg)
}
