package exception

import (
	"fmt"
	"go-server-starter/internal/i18n"
)

type Exception struct {
	StatusCode int       `json:"-"`
	Code       int       `json:"code"`
	Message    string    `json:"message"` // Will be set during translation
	Details    []string  `json:"details"`
	I18nMsg    i18n.Text `json:"-"` // i18n message for translation
}

var codes = map[int]struct{}{}

// register is the internal function to register an exception code
func register(statusCode int, code int, message string, i18nMsg i18n.Text) *Exception {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("Exception code %d already exists", code))
	}
	codes[code] = struct{}{}
	return &Exception{StatusCode: statusCode, Code: code, Message: message, Details: []string{}, I18nMsg: i18nMsg}
}

func NewException(statusCode int, code int, message string, i18nMsg i18n.Text) *Exception {
	return register(statusCode, code, message, i18nMsg)
}

func (e *Exception) clone() *Exception {
	ne := *e
	return &ne
}

func (e *Exception) Append(details ...string) *Exception {
	e = e.clone()
	e.Details = append(e.Details, details...)
	return e
}

func (e *Exception) Is(err *Exception) bool {
	if e == nil || err == nil {
		return false
	}
	return e.Code == err.Code
}
