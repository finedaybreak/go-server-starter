package exception

import (
	"go-server-starter/internal/i18n"
	"net/http"
)

var (
	InternalServerError  = New(http.StatusInternalServerError, 1000, "internal server error", i18n.ExcInternalServerError)
	NotFound             = New(http.StatusNotFound, 1001, "not found", i18n.ExcNotFound)
	BadRequest           = New(http.StatusBadRequest, 1002, "bad request", i18n.ExcBadRequest)
	InvalidParam         = New(http.StatusBadRequest, 1003, "invalid param", i18n.ExcInvalidParam)
	InvalidPathParamID   = New(http.StatusBadRequest, 1004, "invalid path param id", i18n.ExcInvalidPathParamID)
	Unauthorized         = New(http.StatusUnauthorized, 1005, "unauthorized", i18n.ExcUnauthorized)
	TokenInvalid         = New(http.StatusUnauthorized, 1006, "token invalid", i18n.ExcTokenInvalid)
	TokenNotFound        = New(http.StatusUnauthorized, 1007, "token not found", i18n.ExcTokenNotFound)
	TokenUsed            = New(http.StatusUnauthorized, 1008, "token used", i18n.ExcTokenUsed)
	TokenRevoked         = New(http.StatusUnauthorized, 1009, "token revoked", i18n.ExcTokenRevoked)
	TokenExpired         = New(http.StatusUnauthorized, 1010, "token expired", i18n.ExcTokenExpired)
	TokenHasBeenAttacked = New(http.StatusUnauthorized, 1011, "token has been attacked", i18n.ExcTokenHasBeenAttacked)
	TokenGenerateFailed  = New(http.StatusUnauthorized, 1012, "token generate failed", i18n.ExcTokenGenerateFailed)
	Forbidden            = New(http.StatusForbidden, 1013, "forbidden", i18n.ExcForbidden)
	TooManyRequests      = New(http.StatusTooManyRequests, 1014, "too many requests", i18n.ExcTooManyRequests)
	BadGateway           = New(http.StatusBadGateway, 1015, "bad gateway", i18n.ExcBadGateway)
	ServiceUnavailable   = New(http.StatusServiceUnavailable, 1016, "service unavailable", i18n.ExcServiceUnavailable)
	GatewayTimeout       = New(http.StatusGatewayTimeout, 1017, "gateway timeout", i18n.ExcGatewayTimeout)
	NotImplemented       = New(http.StatusNotImplemented, 1018, "not implemented", i18n.ExcNotImplemented)
	ServiceError         = New(http.StatusServiceUnavailable, 1019, "service error", i18n.ExcServiceError)
	ServiceTimeout       = New(http.StatusServiceUnavailable, 1020, "service timeout", i18n.ExcServiceTimeout)
	DatabaseError        = New(http.StatusInternalServerError, 1021, "database error", i18n.ExcDatabaseError)
)
