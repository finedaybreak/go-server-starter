package i18n

// Common exception messages
var (
	ExcInternalServerError  = Text{En: "Internal server error", Zh: "服务器内部错误"}
	ExcNotFound             = Text{En: "Not found", Zh: "资源不存在"}
	ExcBadRequest           = Text{En: "Bad request", Zh: "请求错误"}
	ExcInvalidParam         = Text{En: "Invalid parameters", Zh: "参数错误"}
	ExcInvalidPathParamID   = Text{En: "Invalid path param id", Zh: "路径参数ID错误"}
	ExcUnauthorized         = Text{En: "Unauthorized", Zh: "未授权"}
	ExcTokenInvalid         = Text{En: "Token invalid", Zh: "令牌无效"}
	ExcTokenNotFound        = Text{En: "Token not found", Zh: "未找到令牌"}
	ExcTokenUsed            = Text{En: "Token used", Zh: "令牌已使用"}
	ExcTokenRevoked         = Text{En: "Token revoked", Zh: "令牌已撤销"}
	ExcTokenExpired         = Text{En: "Token expired", Zh: "令牌已过期"}
	ExcTokenHasBeenAttacked = Text{En: "Token has been attacked", Zh: "令牌已被攻击"}
	ExcTokenGenerateFailed  = Text{En: "Token generation failed", Zh: "令牌生成失败"}
	ExcForbidden            = Text{En: "Forbidden", Zh: "禁止访问"}
	ExcTooManyRequests      = Text{En: "Too many requests", Zh: "请求过于频繁"}
	ExcBadGateway           = Text{En: "Bad gateway", Zh: "网关错误"}
	ExcServiceUnavailable   = Text{En: "Service unavailable", Zh: "服务不可用"}
	ExcGatewayTimeout       = Text{En: "Gateway timeout", Zh: "网关超时"}
	ExcNotImplemented       = Text{En: "Not implemented", Zh: "功能未实现"}
	ExcServiceError         = Text{En: "Service error", Zh: "服务错误"}
	ExcServiceTimeout       = Text{En: "Service timeout", Zh: "服务超时"}
	ExcDatabaseError        = Text{En: "Database error", Zh: "数据库错误"}
)
