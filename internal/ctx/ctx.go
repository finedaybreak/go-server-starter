package ctx

import (
	"context"
	"errors"
	"go-server-starter/internal/constant"
	"go-server-starter/internal/enum"
	"go-server-starter/internal/exception"
	"go-server-starter/internal/i18n"
	"go-server-starter/internal/repo"
	"go-server-starter/pkg/utils"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type Context struct {
	Ctx        context.Context
	Gtx        *gin.Context
	translator ut.Translator
}

// FromGinCtx 从 Gin 上下文创建一个新的 Context
func FromGinCtx(c *gin.Context) *Context {
	var translator ut.Translator
	trans, ok := c.Get(constant.CTX_KEY_OF_TRANSLATOR)
	if !ok {
		translator = nil
	} else {
		translator = trans.(ut.Translator)
	}
	return &Context{Ctx: c.Request.Context(), Gtx: c, translator: translator}
}

func (c *Context) ShouldBind(obj any) *exception.Exception {
	var errs []string
	if err := c.Gtx.ShouldBind(obj); err != nil {
		switch {
		case errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF) || err.Error() == "EOF":
			return exception.BadRequest.Append("EOF body is empty")
		default:
			if vErrors, ok := err.(validator.ValidationErrors); ok && c.translator != nil {
				var verrors = make([]string, 0, len(vErrors))
				for _, err := range vErrors {
					var dtoPath = strings.SplitN(err.Namespace(), ".", 2)
					if len(dtoPath) > 1 {
						var translate = strings.Replace(err.Translate(c.translator), err.Field(), "", 1)
						verrors = append(verrors, dtoPath[1]+translate)
					} else {
						verrors = append(verrors, err.Translate(c.translator))
					}
				}
				errs = append(errs, verrors...)
				break
			}
			errs = append(errs, err.Error())
		}
	}
	if len(errs) > 0 {
		return exception.InvalidParam.Append(errs...)
	}
	return nil
}

func (c *Context) GetUserUniCode() (string, *exception.Exception) {
	code := c.Gtx.GetString(constant.CTX_KEY_OF_USER_UNI_CODE)
	if code == "" {
		return "", exception.UserUniCodeNotFound
	}
	return code, nil
}

func (c *Context) SetUserUniCode(code string) {
	c.Gtx.Set(constant.CTX_KEY_OF_USER_UNI_CODE, code)
}

func (c *Context) GetUserID(dbRepo repo.Repo) (uint64, *exception.Exception) {
	code, exc := c.GetUserUniCode()
	if exc != nil {
		return 0, exc
	}
	var ID, err = dbRepo.User().GetIDByUniCode(c.Ctx, code)
	if err != nil {
		return 0, exception.InternalServerError.Append(err.Error())
	}
	return ID, nil
}

func (c *Context) GetPathParamID(key string) (uint64, *exception.Exception) {
	stringID := c.Gtx.Param(key)
	ID := utils.StrToUint64(stringID)
	if ID == 0 {
		return 0, exception.InvalidPathParamID
	}
	return ID, nil
}

func (c *Context) ToError(err *exception.Exception) {
	// Translate exception message if i18nKey is set
	locale := c.GetLocale()
	message := err.I18nMsg.T(locale)
	if message == "" {
		message = err.Message
	}

	c.Gtx.JSON(err.StatusCode, gin.H{
		"code":    err.Code,
		"message": message,
		"details": err.Details,
	})
	c.Gtx.Abort()
}

func (c *Context) ToSuccess(data any) {
	locale := c.GetLocale()
	message := i18n.RespSuccess.T(locale)
	c.Gtx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": message,
		"data":    data,
	})
	c.Gtx.Abort()
}

// GetLocale 从上下文中获取当前的语言环境
func (c *Context) GetLocale() string {
	locale, exists := c.Gtx.Get(constant.CTX_KEY_OF_LOCALE)
	if !exists {
		return "zh"
	}
	if localeStr, ok := locale.(string); ok {
		return localeStr
	}
	return "zh"
}

// GetDeviceType 从gin header 字段判断设备类型
func (c *Context) GetDeviceType() enum.DeviceType {
	deviceType := c.Gtx.GetHeader("Device-Type")
	deviceTypeEnum, err := enum.ParseDeviceType(deviceType)
	if err != nil {
		return enum.DeviceTypeUnknown
	}
	return deviceTypeEnum
}

// 获取设备ID
func (c *Context) GetDeviceID() string {
	deviceID := c.Gtx.GetHeader("Device-ID")
	return deviceID
}

// 获取设备名称
func (c *Context) GetDeviceName() string {
	deviceName := c.Gtx.GetHeader("Device-Name")
	return deviceName
}

// 获取设备信息
func (c *Context) GetDeviceInfo() string {
	deviceInfo := c.Gtx.GetHeader("Device-Info")
	return deviceInfo
}

// 获取用户代理
func (c *Context) GetUserAgent() string {
	userAgent := c.Gtx.GetHeader("User-Agent")
	return userAgent
}
