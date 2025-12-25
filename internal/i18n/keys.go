package i18n

// Message keys for type-safe access to translations
// These constants should match the keys in locale JSON files

const (
	EchoHello   = "echo.hello"
	RespSuccess = "response.success"
	// Exception messages - Common
	ExcInternalServerError  = "exception.internal_server_error"
	ExcNotFound             = "exception.not_found"
	ExcBadRequest           = "exception.bad_request"
	ExcInvalidParam         = "exception.invalid_param"
	ExcInvalidPathParamID   = "exception.invalid_path_param_id"
	ExcUnauthorized         = "exception.unauthorized"
	ExcTokenInvalid         = "exception.token_invalid"
	ExcTokenNotFound        = "exception.token_not_found"
	ExcTokenUsed            = "exception.token_used"
	ExcTokenRevoked         = "exception.token_revoked"
	ExcTokenExpired         = "exception.token_expired"
	ExcTokenHasBeenAttacked = "exception.token_has_been_attacked"
	ExcTokenGenerateFailed  = "exception.token_generate_failed"
	ExcForbidden            = "exception.forbidden"
	ExcTooManyRequests      = "exception.too_many_requests"
	ExcBadGateway           = "exception.bad_gateway"
	ExcServiceUnavailable   = "exception.service_unavailable"
	ExcGatewayTimeout       = "exception.gateway_timeout"
	ExcNotImplemented       = "exception.not_implemented"
	ExcServiceError         = "exception.service_error"
	ExcServiceTimeout       = "exception.service_timeout"
	ExcDatabaseError        = "exception.database_error"

	// Exception messages - User
	ExcUserUniCodeNotFound                   = "exception.user_uni_code_not_found"
	ExcUserMobileVerificationCodeIsIncorrect = "exception.user_mobile_verification_code_is_incorrect"
	ExcUserEmailVerificationCodeIsIncorrect  = "exception.user_email_verification_code_is_incorrect"
	ExcUserMobileHasBeenRegistered           = "exception.user_mobile_has_been_registered"
	ExcUserEmailHasBeenRegistered            = "exception.user_email_has_been_registered"
	ExcUserEmailHasNotBeenRegistered         = "exception.user_email_has_not_been_registered"
	ExcUserMobileHasNotBeenRegistered        = "exception.user_mobile_has_not_been_registered"
	ExcUserPasswordIsIncorrect               = "exception.user_password_is_incorrect"
	ExcUserNotFound                          = "exception.user_not_found"
	ExcUserUpdateInfoFailed                  = "exception.user_update_info_failed"

	// Exception messages - User Role
	ExcUserRoleNotFound      = "exception.user_role_not_found"
	ExcUserRoleAlreadyExists = "exception.user_role_already_exists"
)
