package handler

import (
	"go-server-starter/internal/ctx"
	"go-server-starter/internal/i18n"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HelloHandler interface {
	Hello(c *gin.Context)
}

type HelloHandlerImpl struct {
	logger *zap.Logger
}

func NewHelloHandler(logger *zap.Logger) HelloHandler {
	return &HelloHandlerImpl{logger: logger}
}

func (h *HelloHandlerImpl) Hello(c *gin.Context) {
	ctx := ctx.FromGinCtx(c)
	name := c.Query("name")
	if name == "" {
		ctx.ToSuccess("Hello, World!")
		return
	}
	ctx.ToSuccess(ctx.T(i18n.EchoHello, map[string]string{"name": name}))
}
