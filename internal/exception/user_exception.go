package exception

import (
	"go-server-starter/internal/i18n"
	"net/http"
)

var (
	UserUniCodeNotFound                   = New(http.StatusBadRequest, 20001, "user uniCode not found", i18n.ExcUserUniCodeNotFound)
	UserMobileVerificationCodeIsIncorrect = New(http.StatusBadRequest, 20002, "the mobile verification code is incorrect", i18n.ExcUserMobileVerificationCodeIsIncorrect)
	UserEmailVerificationCodeIsIncorrect  = New(http.StatusBadRequest, 20003, "the email verification code is incorrect", i18n.ExcUserEmailVerificationCodeIsIncorrect)
	UserMobileHasBeenRegistered           = New(http.StatusBadRequest, 20004, "the mobile number has already been registered", i18n.ExcUserMobileHasBeenRegistered)
	UserEmailHasBeenRegistered            = New(http.StatusBadRequest, 20005, "the email address has already been registered", i18n.ExcUserEmailHasBeenRegistered)
	UserEmailHasNotBeenRegistered         = New(http.StatusBadRequest, 20006, "the email address has not been registered", i18n.ExcUserEmailHasNotBeenRegistered)
	UserMobileHasNotBeenRegistered        = New(http.StatusBadRequest, 20007, "the mobile number has not been registered", i18n.ExcUserMobileHasNotBeenRegistered)
	UserPasswordIsIncorrect               = New(http.StatusBadRequest, 20008, "the password is incorrect", i18n.ExcUserPasswordIsIncorrect)
	UserNotFound                          = New(http.StatusNotFound, 20009, "user not found", i18n.ExcUserNotFound)
	UserUpdateInfoFailed                  = New(http.StatusInternalServerError, 20010, "user update info failed", i18n.ExcUserUpdateInfoFailed)
)
