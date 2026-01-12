package exception

import (
	"go-server-starter/internal/i18n"
	"net/http"
)

var (
	InternalServerError  = Common.New(http.StatusInternalServerError, "internal server error", i18n.ExcInternalServerError)
	NotFound             = Common.New(http.StatusNotFound, "not found", i18n.ExcNotFound)
	BadRequest           = Common.New(http.StatusBadRequest, "bad request", i18n.ExcBadRequest)
	InvalidParam         = Common.New(http.StatusBadRequest, "invalid param", i18n.ExcInvalidParam)
	InvalidPathParamID   = Common.New(http.StatusBadRequest, "invalid path param id", i18n.ExcInvalidPathParamID)
	Unauthorized         = Common.New(http.StatusUnauthorized, "unauthorized", i18n.ExcUnauthorized)
	TokenInvalid         = Common.New(http.StatusUnauthorized, "token invalid", i18n.ExcTokenInvalid)
	TokenNotFound        = Common.New(http.StatusUnauthorized, "token not found", i18n.ExcTokenNotFound)
	TokenUsed            = Common.New(http.StatusUnauthorized, "token used", i18n.ExcTokenUsed)
	TokenRevoked         = Common.New(http.StatusUnauthorized, "token revoked", i18n.ExcTokenRevoked)
	TokenExpired         = Common.New(http.StatusUnauthorized, "token expired", i18n.ExcTokenExpired)
	TokenHasBeenAttacked = Common.New(http.StatusUnauthorized, "token has been attacked", i18n.ExcTokenHasBeenAttacked)
	TokenGenerateFailed  = Common.New(http.StatusUnauthorized, "token generate failed", i18n.ExcTokenGenerateFailed)
	Forbidden            = Common.New(http.StatusForbidden, "forbidden", i18n.ExcForbidden)
	TooManyRequests      = Common.New(http.StatusTooManyRequests, "too many requests", i18n.ExcTooManyRequests)
	BadGateway           = Common.New(http.StatusBadGateway, "bad gateway", i18n.ExcBadGateway)
	ServiceUnavailable   = Common.New(http.StatusServiceUnavailable, "service unavailable", i18n.ExcServiceUnavailable)
	GatewayTimeout       = Common.New(http.StatusGatewayTimeout, "gateway timeout", i18n.ExcGatewayTimeout)
	NotImplemented       = Common.New(http.StatusNotImplemented, "not implemented", i18n.ExcNotImplemented)
	ServiceError         = Common.New(http.StatusServiceUnavailable, "service error", i18n.ExcServiceError)
	ServiceTimeout       = Common.New(http.StatusServiceUnavailable, "service timeout", i18n.ExcServiceTimeout)
	DatabaseError        = Common.New(http.StatusInternalServerError, "database error", i18n.ExcDatabaseError)
)
