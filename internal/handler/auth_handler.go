package handler

import (
	"go-server-starter/internal/ctx"
	"go-server-starter/internal/dto"
	"go-server-starter/internal/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthHandler interface {
	LoginByMobileAndCode(c *gin.Context)
	LoginByEmailAndCode(c *gin.Context)
}

type AuthHandlerImpl struct {
	logger  *zap.Logger
	service service.Service
}

func NewAuthHandler(logger *zap.Logger, service service.Service) AuthHandler {
	return &AuthHandlerImpl{logger: logger, service: service}
}

func (h *AuthHandlerImpl) LoginByMobileAndCode(c *gin.Context) {
	var ctx = ctx.FromGinCtx(c)
	var params dto.AuthLoginByMobileAndCodeReqDto
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.ToError(err)
		return
	}
	res, err := h.service.Auth().LoginByMobileAndCode(ctx, params)
	if err != nil {
		ctx.ToError(err)
		return
	}
	ctx.ToSuccess(res)
}

func (h *AuthHandlerImpl) LoginByEmailAndCode(c *gin.Context) {
	var ctx = ctx.FromGinCtx(c)
	var params dto.AuthLoginByEmailAndCodeReqDto
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.ToError(err)
		return
	}
	res, err := h.service.Auth().LoginByEmailAndCode(ctx, params)
	if err != nil {
		ctx.ToError(err)
		return
	}
	ctx.ToSuccess(res)
}
