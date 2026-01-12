package exception

import (
	"go-server-starter/internal/i18n"
	"net/http"
)

var (
	UserUniCodeNotFound                   = User.New(http.StatusBadRequest, "user uniCode not found", i18n.ExcUserUniCodeNotFound)
	UserMobileVerificationCodeIsIncorrect = User.New(http.StatusBadRequest, "the mobile verification code is incorrect", i18n.ExcUserMobileVerificationCodeIsIncorrect)
	UserEmailVerificationCodeIsIncorrect  = User.New(http.StatusBadRequest, "the email verification code is incorrect", i18n.ExcUserEmailVerificationCodeIsIncorrect)
	UserMobileHasBeenRegistered           = User.New(http.StatusBadRequest, "the mobile number has already been registered", i18n.ExcUserMobileHasBeenRegistered)
	UserEmailHasBeenRegistered            = User.New(http.StatusBadRequest, "the email address has already been registered", i18n.ExcUserEmailHasBeenRegistered)
	UserEmailHasNotBeenRegistered         = User.New(http.StatusBadRequest, "the email address has not been registered", i18n.ExcUserEmailHasNotBeenRegistered)
	UserMobileHasNotBeenRegistered        = User.New(http.StatusBadRequest, "the mobile number has not been registered", i18n.ExcUserMobileHasNotBeenRegistered)
	UserPasswordIsIncorrect               = User.New(http.StatusBadRequest, "the password is incorrect", i18n.ExcUserPasswordIsIncorrect)
	UserNotFound                          = User.New(http.StatusNotFound, "user not found", i18n.ExcUserNotFound)
	UserUpdateInfoFailed                  = User.New(http.StatusInternalServerError, "user update info failed", i18n.ExcUserUpdateInfoFailed)
)
