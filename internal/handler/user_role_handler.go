package handler

import (
	"go-server-starter/internal/service"

	"go.uber.org/zap"
)

type UserRoleHandler interface {
}

type UserRoleHandlerImpl struct {
	logger  *zap.Logger
	service service.Service
}

func NewUserRoleHandler(logger *zap.Logger, service service.Service) UserRoleHandler {
	return &UserRoleHandlerImpl{logger: logger, service: service}
}
