package handler

import (
	"go-server-starter/internal/ctx"
	"go-server-starter/internal/dto"
	"go-server-starter/internal/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandler interface {
	GetInfo(c *gin.Context)
	UpdateInfo(c *gin.Context)
	GetTable(c *gin.Context)
}

type UserHandlerImpl struct {
	logger  *zap.Logger
	service service.Service
}

func NewUserHandler(logger *zap.Logger, service service.Service) UserHandler {
	return &UserHandlerImpl{logger: logger, service: service}
}

func (h *UserHandlerImpl) GetInfo(c *gin.Context) {
	var ctx = ctx.FromGinCtx(c)
	uniCode, err := ctx.GetUserUniCode()
	if err != nil {
		ctx.ToError(err)
		return
	}
	user, err := h.service.User().GetInfoByUniCode(ctx, uniCode)
	if err != nil {
		ctx.ToError(err)
		return
	}
	ctx.ToSuccess(user)
}

func (h *UserHandlerImpl) UpdateInfo(c *gin.Context) {
	var ctx = ctx.FromGinCtx(c)
	var params dto.UserUpdateInfoReqDto
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.ToError(err)
		return
	}
	res, err := h.service.User().UpdateInfo(ctx, params)
	if err != nil {
		ctx.ToError(err)
		return
	}
	ctx.ToSuccess(res)
}

func (h *UserHandlerImpl) GetTable(c *gin.Context) {
	var ctx = ctx.FromGinCtx(c)
	var params dto.UserTableQueryReqDto
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.ToError(err)
		return
	}
	res, err := h.service.User().GetTable(ctx, params)
	if err != nil {
		ctx.ToError(err)
		return
	}
	ctx.ToSuccess(res)
}
