package exception

import "fmt"

type Exception struct {
	StatusCode int      `json:"-"`
	Code       int      `json:"code"`
	Message    string   `json:"message"` // Will be set during translation
	Details    []string `json:"details"`
	I18nKey    string   `json:"-"` // i18n translation key
}

var codes = map[int]struct{}{}

// New creates a new exception with an i18n key
func New(statusCode int, code int, message string, i18nKey string) *Exception {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("Exception code %d already exists", code))
	}
	codes[code] = struct{}{}
	return &Exception{StatusCode: statusCode, Code: code, Message: message, Details: []string{}, I18nKey: i18nKey}
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
